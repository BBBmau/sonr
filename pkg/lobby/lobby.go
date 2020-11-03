package lobby

import (
	"context"
	"encoding/json"

	"github.com/libp2p/go-libp2p-core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

// SonrCallback returns message from lobby
type SonrCallback interface {
	OnMessage(s string)
	OnNewPeer(s string)
}

// Lobby represents a subscription to a single PubSub topic. Messages
// can be published to the topic with Lobby.Publish, and received
// messages are pushed to the Messages channel.
type Lobby struct {
	// Messages is a channel of messages received from other peers in the chat room
	messages chan *Message
	Callback SonrCallback

	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	Code   string
	selfID peer.ID
}

// Message gets converted to/from JSON and sent in the body of pubsub messages.
type Message struct {
	Value    string
	Event    string
	SenderID string
}

// Enter tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func Enter(ctx context.Context, call SonrCallback, ps *pubsub.PubSub, hostID peer.ID, olcCode string) (*Lobby, error) {
	// join the pubsub topic
	topic, err := ps.Join(olcName(olcCode))
	if err != nil {
		return nil, err
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}

	// Create Lobby Type
	lob := &Lobby{
		ctx:      ctx,
		ps:       ps,
		topic:    topic,
		sub:      sub,
		selfID:   hostID,
		Code:     olcCode,
		Callback: call,
		messages: make(chan *Message, ChatRoomBufSize),
	}

	// start reading messages from the subscription in a loop
	go lob.readLoop()
	return lob, nil
}

// Publish sends a message to the pubsub topic.
func (lob *Lobby) Publish(value string, event string) error {
	// Create Message
	m := Message{
		Event:    event,
		Value:    value,
		SenderID: lob.selfID.Pretty(),
	}

	// Convert to JSON
	msgBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}

	// Publish to Topic
	err = lob.topic.Publish(lob.ctx, msgBytes)
	if err != nil {
		return err
	}
	return nil
}

// ListPeers returns peerids in room
func (lob *Lobby) ListPeers() []peer.ID {
	return lob.ps.ListPeers(olcName(lob.Code))
}

// readLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (lob *Lobby) readLoop() {
	for {
		msg, err := lob.sub.Next(lob.ctx)
		if err != nil {
			close(lob.messages)
			return
		}

		//println(string(msg.Data))
		lob.Callback.OnMessage(string(msg.Data))

		// only forward messages delivered by others
		if msg.ReceivedFrom == lob.selfID {
			continue
		}
		cm := new(Message)
		err = json.Unmarshal(msg.Data, cm)
		if err != nil {
			continue
		}

		// send valid messages onto the Messages channel
		lob.messages <- cm
	}
}

func olcName(code string) string {
	return "olc=" + code
}
