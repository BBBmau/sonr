package transmit

import (
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/common"
	"github.com/sonr-io/core/node"
	"github.com/sonr-io/core/types/go/node/motor/v1"

	transmitV1 "github.com/sonr-io/core/types/go/protocols/transmit/v1"
)

// NewInSession creates a new Session from the given payload with Incoming direction.
func NewInSession(payload *common.Payload, from *common.Peer, to *common.Peer) *transmitV1.Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &transmitV1.Session{
		Direction:    common.Direction_INCOMING,
		Payload:      payload,
		From:         from,
		To:           to,
		LastUpdated:  int64(time.Now().Unix()),
		Items:        CreatePayloadItems(sessionPayload, common.Direction_INCOMING),
		CurrentIndex: 0,
		Results:      make(map[int32]bool),
	}
}

// NewOutSession creates a new Session from the given payload with Outgoing direction.
func NewOutSession(payload *common.Payload, to *common.Peer, from *common.Peer) *transmitV1.Session {
	// Create Session Items
	sessionPayload := NewSessionPayload(payload)
	return &transmitV1.Session{
		Direction:    common.Direction_OUTGOING,
		Payload:      payload,
		To:           to,
		From:         from,
		LastUpdated:  int64(time.Now().Unix()),
		Items:        CreatePayloadItems(sessionPayload, common.Direction_OUTGOING),
		CurrentIndex: 0,
		Results:      make(map[int32]bool),
	}
}

// FinalIndex returns the final index of the session.
func SessionFinalIndex(s *transmitV1.Session) int {
	return len(s.Items) - 1
}

// HasRead returns true if all files have been read.
func SessionHasRead(s *transmitV1.Session) bool {
	return SessionIsIn(s) && SessionIsDone(s)
}

// HasWrote returns true if all files have been written.
func SessionHasWrote(s *transmitV1.Session) bool {
	return SessionIsOut(s) && SessionIsDone(s)
}

// IsDone returns true if all files have been read or written.
func SessionIsDone(s *transmitV1.Session) bool {
	return int(s.GetCurrentIndex()) >= SessionFinalIndex(s)
}

// IsOut returns true if the session is outgoing.
func SessionIsOut(s *transmitV1.Session) bool {
	return s.Direction == common.Direction_OUTGOING
}

// IsIn returns true if the session is incoming.
func SessionIsIn(s *transmitV1.Session) bool {
	return s.Direction == common.Direction_INCOMING
}

// Event returns the complete event for the session.
func SessionEvent(s *transmitV1.Session) *motor.OnTransmitCompleteResponse {
	return &motor.OnTransmitCompleteResponse{
		From:       s.GetFrom(),
		To:         s.GetTo(),
		Direction:  s.GetDirection(),
		Payload:    s.GetPayload(),
		CreatedAt:  s.GetPayload().GetCreatedAt(),
		ReceivedAt: int64(time.Now().Unix()),
		Results:    s.GetResults(),
	}
}

// RouteStream is used to route the given stream to the given peer.
func RouteSessionStream(s *transmitV1.Session, stream network.Stream, n node.CallbackImpl) (*motor.OnTransmitCompleteResponse, error) {
	// Initialize Params
	logger.Debugf("Beginning %s Transmit Stream", s.Direction.String())
	doneChan := make(chan bool)

	// Check for Incoming
	if SessionIsIn(s) {
		// Handle incoming stream
		go func(stream network.Stream, dchan chan bool) {
			// Create reader
			rs := msgio.NewReader(stream)

			// Read all items
			for _, v := range s.GetItems() {
				// Read Stream to File
				if err := ReadItemFromStream(v, n, rs); err != nil {
					logger.Errorf("Error reading stream: %v", err)
					dchan <- false
				} else {
					dchan <- true
				}
			}

			// Close Stream on Done Reading
			stream.Close()
		}(stream, doneChan)
	}

	// Check for Outgoing
	if SessionIsOut(s) {
		// Handle outgoing stream
		go func(stream network.Stream, dchan chan bool) {
			// Create writer
			wc := msgio.NewWriter(stream)

			// Write all items
			for _, v := range s.GetItems() {
				// Write File to Stream
				if err := WriteItemToStream(v, n, wc); err != nil {
					logger.Errorf("Error writing file: %v", err)
					dchan <- false
				} else {
					dchan <- true
				}
			}
		}(stream, doneChan)
	}

	// Wait for all files to be written
	for {
		select {
		case r := <-doneChan:
			// Set Result
			if complete := UpdateCurrent(s, r); !complete {
				continue
			} else {
				return SessionEvent(s), nil
			}
		}
	}
}

// UpdateCurrent updates the current index of the session.
func UpdateCurrent(s *transmitV1.Session, result bool) bool {
	logger.Debugf("Item (%v) transmit result: %v", s.GetCurrentIndex(), result)
	s.Results[s.GetCurrentIndex()] = result
	s.CurrentIndex = s.GetCurrentIndex() + 1
	return int(s.GetCurrentIndex()) >= SessionFinalIndex(s)
}
