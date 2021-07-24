package js

import (
	"errors"
	"fmt"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/textileio/go-threads/core/thread"

	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/multiformats/go-multiaddr"
)

var (
	NoPubKey = errors.New("Public Key not found from Peer Protobuf.")
)

// ** ─── Topic MANAGEMENT ────────────────────────────────────────────────────────
func (t *Topic) IsLocal() bool {
	return t.Type == Topic_LOCAL
}

// @ Local Lobby Topic Protocol ID
func (r *User) NewLocalTopic() *Topic {
	name := fmt.Sprintf("/sonr/topic/%s", r.Location.OLC())
	return &Topic{
		Name: name,
		Type: Topic_LOCAL,
	}
}

// ** ─── Peer MANAGEMENT ────────────────────────────────────────────────────────
// ^ Create New Peer from Connection Request and Host ID ^ //
func (u *User) NewPeer(id peer.ID, maddr multiaddr.Multiaddr) *SonrError {
	u.Peer = &Peer{
		SName: u.SName,
		Id: &Peer_ID{
			Peer:      id.String(),
			Device:    u.DeviceID(),
			MultiAddr: maddr.String(),
			PublicKey: u.KeyPair().PubKeyBase64(),
			PushToken: u.GetPushToken(),
		},
		Profile:  u.Profile(),
		Platform: u.Device.Platform,
		Model:    u.Device.Model,
	}
	return nil
}

// ** ─── Local Event MANAGEMENT ────────────────────────────────────────────────────────
// Creates New Exit Local Event
func NewJoinLocalEvent(peer *Peer) *TopicEvent {
	return &TopicEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: TopicEvent_JOIN,
	}
}

// Creates New Exit Local Event
func NewUpdateLocalEvent(peer *Peer, topic *Topic) *TopicEvent {
	return &TopicEvent{
		Id:      peer.Id.Peer,
		Peer:    peer,
		Subject: TopicEvent_UPDATE,
		Topic:   topic,
	}
}

// Creates New Exit Local Event
func NewExitLocalEvent(id string, topic *Topic) *TopicEvent {
	return &TopicEvent{
		Id:      id,
		Subject: TopicEvent_EXIT,
		Topic:   topic,
	}
}

// ^ Returns Peer as Buffer ^ //
func (p *Peer) Buffer() []byte {
	return p.Marshal()
}

// ^ Returns Peer User ID ^ //
func (p *Peer) DeviceID() string {
	return string(p.Id.GetDevice())
}

// ^ Returns Peer ID String Value
func (p *Peer) PeerID() string {
	return p.Id.Peer
}

// ^ Returns Peer Public Key ^ //
func (p *Peer) PublicKey() (crypto.PubKey, *SonrError) {
	// Get ID from Public Key
	buf, err := crypto.ConfigDecodeKey(p.GetId().GetPublicKey())
	if err != nil {
		return nil, NewError(err, ErrorMessage_PEER_PUBKEY_DECODE)
	}

	// Unmarshal Public Key
	pubKey, err := crypto.UnmarshalPublicKey(buf)
	if err != nil {
		return nil, NewError(err, ErrorMessage_PEER_PUBKEY_UNMARSHAL)
	}
	return pubKey, nil
}

// ^ Checks if Two Peers are the Same by Device ID and Peer ID
func (p *Peer) IsSame(other *Peer) bool {
	return p.PeerID() == other.PeerID() && p.DeviceID() == other.DeviceID() && p.GetSName() == other.GetSName()
}

// ^ Checks if PeerDeviceIDID is the Same
func (p *Peer) IsSameDeviceID(other *Peer) bool {
	return p.DeviceID() == other.DeviceID()
}

// ^ Checks if PeerID is the Same
func (p *Peer) IsSamePeerID(pid peer.ID) bool {
	return p.PeerID() == pid.String()
}

// ^ Checks if Two Peers are NOT the Same by Device ID and Peer ID
func (p *Peer) IsNotSame(other *Peer) bool {
	return p.PeerID() != other.PeerID() && p.DeviceID() != other.DeviceID() && p.GetSName() != other.GetSName()
}

// ^ Checks if DeviceID is NOT the Same
func (p *Peer) IsNotSameDeviceID(other *Peer) bool {
	return p.DeviceID() == other.DeviceID()
}

// ^ Checks if PeerID is NOT the Same
func (p *Peer) IsNotSamePeerID(pid peer.ID) bool {
	return p.PeerID() != pid.String()
}

// ^ Converts Peer Public Key into Thread Key
func (p *Peer) ThreadKey() (thread.PubKey, *SonrError) {
	// Get Pub Key
	pubKey, err := p.PublicKey()
	if err != nil {
		return nil, err
	}

	// Create Thread Pub Key
	threadKey := thread.NewLibp2pPubKey(pubKey)
	return threadKey, nil
}

// ^ Returns Peer Push Token
func (p *Peer) PushToken() (string, *SonrError) {
	if p.Id.GetPushToken() == "" {
		return "", NewError(nil, ErrorMessage_PEER_PUSH_TOKEN_EMPTY)
	}
	return p.Id.GetPushToken(), nil
}