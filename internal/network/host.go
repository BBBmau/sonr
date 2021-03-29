package network

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/crypto"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	dsc "github.com/libp2p/go-libp2p-discovery"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	psub "github.com/libp2p/go-libp2p-pubsub"
	swr "github.com/libp2p/go-libp2p-swarm"
	md "github.com/sonr-io/core/internal/models"
)

type HostNode struct {
	ctx       context.Context
	Discovery *dsc.RoutingDiscovery
	Host      host.Host
	KDHT      *dht.IpfsDHT
	Point     string
	Pubsub    *psub.PubSub
}

// ^ Start Begins Assigning Host Parameters ^
func NewHost(ctx context.Context, point string, key crypto.PrivKey) (*HostNode, error) {
	// IP Address
	ip4 := IPv4()
	ip6 := IPv6()

	// Start Host
	h, err := libp2p.New(
		ctx,
		libp2p.Identity(key),
		// Add listening Addresses
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/%s/tcp/0", ip4),
			fmt.Sprintf("/ip6/%s/tcp/0", ip6),
		),
		libp2p.NATPortMap(),

		libp2p.EnableAutoRelay(),
	)

	// Create DHT
	kdht, err := dht.New(ctx, h)
	if err != nil {
		return nil, err
	}

	// Set Host for Node
	if err != nil {
		return nil, err
	}
	return &HostNode{
		ctx:   ctx,
		Host:  h,
		KDHT:  kdht,
		Point: point,
	}, nil
}

// ^ Bootstrap begins bootstrap with peers ^
func (h *HostNode) Bootstrap() error {
	// Create Bootstrapper Info
	bootstrappers, err := GetBootstrapAddrInfo()
	if err != nil {
		return err
	}

	// Bootstrap DHT
	if err := h.KDHT.Bootstrap(h.ctx); err != nil {
		return err
	}

	// Connect to bootstrap nodes, if any
	for _, pi := range bootstrappers {
		if err := h.Host.Connect(h.ctx, pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Set Routing Discovery, Find Peers
	routingDiscovery := dsc.NewRoutingDiscovery(h.KDHT)
	dsc.Advertise(h.ctx, routingDiscovery, h.Point, dscl.TTL(time.Second*4))
	h.Discovery = routingDiscovery

	// Create Pub Sub
	ps, err := psub.NewGossipSub(h.ctx, h.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		return err
	}
	h.Pubsub = ps
	go h.handleDHTPeers(routingDiscovery)
	return nil
}

// ^ Return Host ID ^
func (h *HostNode) ID() peer.ID {
	return h.Host.ID()
}

// ^ Set Stream Handler for Host ^
func (h *HostNode) HandleStream(pid protocol.ID, handler network.StreamHandler) {
	h.Host.SetStreamHandler(pid, handler)
}

// ^ Start Stream for Host ^
func (h *HostNode) StartStream(p peer.ID, pid protocol.ID) (network.Stream, error) {
	return h.Host.NewStream(h.ctx, p, pid)
}

// ^ Join New Topic with Name ^
func (h *HostNode) Join(name string) (*psub.Topic, *psub.Subscription, *psub.TopicEventHandler, error) {
	// Join Topic
	topic, err := h.Pubsub.Join(name)
	if err != nil {
		return nil, nil, nil, err
	}

	// Subscribe to Topic
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, nil, nil, err
	}

	// Create Topic Handler
	handler, err := topic.EventHandler()
	if err != nil {
		return nil, nil, nil, err
	}
	return topic, sub, handler, nil
}

// @ handleDHTPeers: Connects to Peers in DHT
func (h *HostNode) handleDHTPeers(routingDiscovery *dsc.RoutingDiscovery) {
	for {
		// Find peers in DHT
		peersChan, err := routingDiscovery.FindPeers(
			h.ctx,
			h.Point,
			dscl.TTL(time.Second*4),
		)
		if err != nil {
			return
		}

		// Iterate over Channel
		for pi := range peersChan {
			// Validate not Self
			if pi.ID != h.Host.ID() {
				// Connect to Peer
				if err := h.Host.Connect(h.ctx, pi); err != nil {
					// Remove Peer Reference
					h.Host.Peerstore().ClearAddrs(pi.ID)
					if sw, ok := h.Host.Network().(*swr.Swarm); ok {
						sw.Backoff().Clear(pi.ID)
					}
				}
			}
		}
		md.GetState().NeedsWait()
	}
}