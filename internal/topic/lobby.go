package topic

import (
	"log"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ Send Updated Lobby ^
func (tm *TopicManager) Refresh() {
	tm.topicHandler.OnRefresh(tm.Lobby)
}

// ^ handleTopicEvents: listens to Pubsub Events for topic  ^
func (tm *TopicManager) handleTopicEvents(p *md.Peer) {
	// @ Loop Events
	for {
		// Get next event
		lobEvent, err := tm.eventHandler.NextPeerEvent(tm.ctx)
		if err != nil {
			tm.eventHandler.Cancel()
			return
		}

		if lobEvent.Type == pubsub.PeerJoin {
			pbuf, err := p.Buffer()
			if err != nil {
				log.Println(err)
				continue
			}
			lbuf, err := tm.Lobby.Buffer()
			if err != nil {
				log.Println(err)
				continue
			}
			err = tm.Exchange(lobEvent.Peer, pbuf, lbuf)
			if err != nil {
				continue
			}
		}

		if lobEvent.Type == pubsub.PeerLeave {
			tm.Lobby.Delete(lobEvent.Peer)
		}

		md.GetState().NeedsWait()
	}
}

// ^ handleTopicMessages: listens for messages on pubsub topic subscription ^
func (tm *TopicManager) handleTopicMessages(p *md.Peer) {
	for {
		// Get next msg from pub/sub
		msg, err := tm.subscription.Next(tm.ctx)
		if err != nil {
			return
		}

		// Only forward messages delivered by others
		if p.IsPeerID(msg.ReceivedFrom) {
			continue
		}

		// Construct message
		m := &md.LobbyEvent{}
		err = proto.Unmarshal(msg.Data, m)
		if err != nil {
			continue
		}

		// Validate Peer in Lobby
		if tm.HasPeer(m.Id) {
			tm.Messages <- m
		}
		md.GetState().NeedsWait()
	}
}

// ^ processTopicMessages: pulls messages from channel that have been handled ^
func (tm *TopicManager) processTopicMessages(p *md.Peer) {
	for {
		select {
		// @ when we receive a message from the lobby room
		case m := <-tm.Messages:
			// Update Circle by event
			if m.Event == md.LobbyEvent_UPDATE {
				// Update Peer Data
				tm.Lobby.Add(m.From)
				tm.Refresh()
			} else if m.Event == md.LobbyEvent_MESSAGE {
				// Check is Message For Self
				if p.IsPeerIDString(m.To) {
					// Call Event
					tm.topicHandler.OnEvent(m)
				}
			}
		case <-tm.ctx.Done():
			return
		}
		md.GetState().NeedsWait()
	}
}