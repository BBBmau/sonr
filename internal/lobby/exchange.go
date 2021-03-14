package lobby

import (
	"context"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	gorpc "github.com/libp2p/go-libp2p-gorpc"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ****************** //
// ** GRPC Service ** //
// ****************** //
// ExchangeArgs is Peer protobuf
type ExchangeArgs struct {
	Data []byte
}

// ExchangeResponse is also Peer protobuf
type ExchangeResponse struct {
	Data []byte
}

// Service Struct
type ExchangeService struct {
	getUser    md.ReturnPeer
	updatePeer md.UpdatePeer
}

// ^ Calls Invite on Remote Peer ^ //
func (ps *ExchangeService) ExchangeWith(ctx context.Context, args ExchangeArgs, reply *ExchangeResponse) error {
	// Peer Data
	remotePeer := &md.Peer{}
	err := proto.Unmarshal(args.Data, remotePeer)
	if err != nil {
		return err
	}

	// Update Peers
	ps.updatePeer(remotePeer)

	// Set Current Message
	userPeer := ps.getUser()

	// Convert Protobuf to bytes
	replyData, err := proto.Marshal(userPeer)
	if err != nil {
		return err
	}

	// Set Message data and call done
	reply.Data = replyData
	return nil
}

// ^ Calls Invite on Remote Peer ^ //
func (lob *Lobby) Exchange(id peer.ID) {
	// Get Peer Data
	userPeer := lob.call.Peer()
	msgBytes, err := proto.Marshal(userPeer)
	if err != nil {
		lob.call.Error(err, "Exchange")
	}

	// Initialize RPC
	rpcClient := gorpc.NewClient(lob.host, protocol.ID("/sonr/lobby/exchange"))
	var reply ExchangeResponse
	var args ExchangeArgs
	args.Data = msgBytes

	// Call to Peer
	err = rpcClient.Call(id, "ExchangeService", "ExchangeWith", args, &reply)
	if err != nil {
		lob.call.Error(err, "Exchange")
	}

	// Received Message
	remotePeer := &md.Peer{}
	err = proto.Unmarshal(reply.Data, remotePeer)
	if err != nil {
		// Send Error
		lob.call.Error(err, "Exchange")
	}

	// Update Peers
	lob.updatePeer(remotePeer)
}