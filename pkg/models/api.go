package models

import (
	"fmt"

	olc "github.com/google/open-location-code/go"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
)

// ************************** //
// ** MIME Info Management ** //
// ************************** //
// Method adjusts extension for JPEG
func (m *MIME) Ext() string {
	if m.Subtype == "jpg" || m.Subtype == "jpeg" {
		return "jpeg"
	}
	return m.Subtype
}

// Checks if Mime is Audio
func (m *MIME) IsAudio() bool {
	return m.Type == MIME_AUDIO
}

// Checks if Mime is any media
func (m *MIME) IsMedia() bool {
	return m.Type == MIME_AUDIO || m.Type == MIME_IMAGE || m.Type == MIME_VIDEO
}

// Checks if Mime is Image
func (m *MIME) IsImage() bool {
	return m.Type == MIME_IMAGE
}

// Checks if Mime is Video
func (m *MIME) IsVideo() bool {
	return m.Type == MIME_VIDEO
}

// ** ─── AuthInvite MANAGEMENT ────────────────────────────────────────────────────────
func (r *AuthReply) HasAcceptedTransfer() bool {
	return r.Decision && r.Type == AuthReply_Transfer
}

// ** ─── AuthInvite MANAGEMENT ────────────────────────────────────────────────────────
func (i *AuthInvite) GetContact() *Contact {
	return i.GetData().GetContact()
}

func (i *AuthInvite) GetFile() *SonrFile {
	return i.GetData().GetFile()
}

func (i *AuthInvite) GetUrl() *URLLink {
	return i.GetData().GetUrl()
}

func (i *AuthInvite) IsFlat() bool {
	return i.Data.Properties.IsFlat
}

func (i *AuthInvite) IsRemote() bool {
	return i.Data.Properties.IsRemote
}

// ** ─── Location MANAGEMENT ────────────────────────────────────────────────────────
func (l *Location) MinorOLC() string {
	lat := l.GetLatitude()
	lon := l.GetLongitude()
	return olc.Encode(lat, lon, 6)
}

func (l *Location) MajorOLC() string {
	lat := l.GetLatitude()
	lon := l.GetLongitude()
	return olc.Encode(lat, lon, 2)
}

func (l *Location) OLC() string {
	return olc.Encode(float64(l.GetLatitude()), float64(l.GetLongitude()), 5)
}

// ** ─── Router MANAGEMENT ────────────────────────────────────────────────────────
// @ Local Lobby Topic Protocol ID
func (r *User) LocalIPTopic() string {
	return fmt.Sprintf("/sonr/topic/%s", r.Location.OLC())
}

// @ Transfer Controller Data Protocol ID
func (r *User_Router) Transfer(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/transfer/%s", id.Pretty()))
}

// @ Transfer Controller Data Protocol ID
func (r *User_Router) Remote(id peer.ID) protocol.ID {
	return protocol.ID(fmt.Sprintf("/sonr/remote/%s", id.Pretty()))
}

// @ Lobby Topic Protocol ID
func (r *User_Router) Topic(name string) string {
	return fmt.Sprintf("/sonr/topic/%s", name)
}

// @ Major Rendevouz Advertising Point
func (u *User) GetRouter() *User_Router {
	return u.GetConnection().GetRouter()
}

// ** ─── Status MANAGEMENT ────────────────────────────────────────────────────────
// Update Connected Connection Status
func (u *User) SetConnected(value bool) *StatusUpdate {
	// Set Value
	u.Connection.HasConnected = value

	// Update Status
	if value {
		u.Connection.Status = Status_CONNECTED
	} else {
		u.Connection.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Bootstrap Connection Status
func (u *User) SetBootstrapped(value bool) *StatusUpdate {
	// Set Value
	u.Connection.HasBootstrapped = value

	// Update Status
	if value {
		u.Connection.Status = Status_BOOTSTRAPPED
	} else {
		u.Connection.Status = Status_FAILED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Bootstrap Connection Status
func (u *User) SetJoinedLocal(value bool) *StatusUpdate {
	// Set Value
	u.Connection.HasJoinedLocal = value

	// Update Status
	if value {
		u.Connection.Status = Status_AVAILABLE
	} else {
		u.Connection.Status = Status_BOOTSTRAPPED
	}

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Update Node Status
func (u *User) SetStatus(ns Status) *StatusUpdate {
	// Set Value
	u.Connection.Status = ns

	// Returns Status Update
	return &StatusUpdate{Value: u.Connection.GetStatus()}
}

// Checks if Status is Given Value
func (u *User) IsStatus(gs Status) bool {
	return u.GetConnection().GetStatus() == gs
}

// Checks if Status is Not Given Value
func (u *User) IsNotStatus(gs Status) bool {
	return u.GetConnection().GetStatus() != gs
}
