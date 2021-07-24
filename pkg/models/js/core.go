package js

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	msg "github.com/libp2p/go-msgio"
	"github.com/sonr-io/core/pkg/util"
)

// ** ─── CALLBACK MANAGEMENT ────────────────────────────────────────────────────────
type OnConnected func(r *ConnectionResponse)
type HTTPHandler func(http.ResponseWriter, *http.Request)
type SetStatus func(s Status)
type OnProtobuf func([]byte)
type OnError func(err *SonrError)
type Callback struct {
	// OnInvite      OnProtobuf
	// OnConnected   OnProtobuf
	OnEvent OnProtobuf
	// OnMail        OnProtobuf
	// OnReply       OnProtobuf
	// OnProgress    OnProtobuf
	OnRequest  OnProtobuf
	OnResponse OnProtobuf
	OnError    OnError
	SetStatus  SetStatus
}

// ** ─── State MANAGEMENT ────────────────────────────────────────────────────────
type state struct {
	flag uint64
	chn  chan bool
}

var (
	instance *state
	once     sync.Once
)

func GetState() *state {
	once.Do(func() {
		chn := make(chan bool)
		close(chn)

		instance = &state{chn: chn}
	})

	return instance
}

// Checks rather to wait or does not need
func (c *state) NeedsWait() {
	<-c.chn
}

// Says all of goroutines to resume execution
func (c *state) Resume() {
	if atomic.LoadUint64(&c.flag) == 1 {
		close(c.chn)
		atomic.StoreUint64(&c.flag, 0)
	}
}

// Says all of goroutines to pause execution
func (c *state) Pause() {
	if atomic.LoadUint64(&c.flag) == 0 {
		atomic.StoreUint64(&c.flag, 1)
		c.chn = make(chan bool)
	}
}

// ** ─── Transfer MANAGEMENT ────────────────────────────────────────────────────────
// Returns Transfer for URLLink
func (u *URLLink) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_Url{
			Url: u,
		},
	}
}

// Returns URLLink as Transfer_Url Data
func (u *URLLink) ToData() *Transfer_Url {
	return &Transfer_Url{
		Url: u,
	}
}

// Returns Transfer for SFile
func (f *SFile) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_File{
			File: f,
		},
	}
}

// Returns SFile as Transfer_File Data
func (f *SFile) ToData() *Transfer_File {
	return &Transfer_File{
		File: f,
	}
}

// Returns Transfer for Contact
func (c *Contact) GetTransfer() *Transfer {
	return &Transfer{
		Data: &Transfer_Contact{
			Contact: c,
		},
	}
}

// Returns Contact as Transfer_Contact Data
func (c *Contact) ToData() *Transfer_Contact {
	return &Transfer_Contact{
		Contact: c,
	}
}

// ** ─── MIME MANAGEMENT ────────────────────────────────────────────────────────
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

// ** ─── SFile MANAGEMENT ────────────────────────────────────────────────────────
// Checks if File contains single item
func (f *SFile) IsSingle() bool {
	return len(f.Items) == 1
}

// Checks if Single File is Media
func (f *SFile) IsMedia() bool {
	return f.Payload == Payload_MEDIA || (f.IsSingle() && f.Single().Mime.IsMedia())
}

// Checks if File contains multiple items
func (f *SFile) IsMultiple() bool {
	return len(f.Items) > 1
}

// Returns Index of Item from File
func (f *SFile) IndexOf(item *SFile_Item) int {
	for i, v := range f.Items {
		if v == item {
			return i
		}
	}
	return -1
}

// Method Returns Single if Applicable
func (f *SFile) Single() *SFile_Item {
	if f.IsSingle() {
		return f.Items[0]
	} else {
		return nil
	}
}

// ** ─── Session MANAGEMENT ────────────────────────────────────────────────────────
type Session struct {
	// Inherited Properties
	file      *SFile
	owner     *Peer
	receiver  *Peer
	pid       protocol.ID
	user      *User
	direction CompleteEvent_Direction

	// Management
	call SessionHandler
}

type SessionHandler interface {
	OnCompleted(stream network.Stream, pid protocol.ID, completeEvent *CompleteEvent)
	OnProgress(buf []byte)
	OnError(err *SonrError)
}

// ^ Prepare for Outgoing Session ^ //
func NewOutSession(u *User, req *InviteRequest, sh SessionHandler) *Session {
	return &Session{
		file:      req.GetFile(),
		owner:     req.GetFrom(),
		receiver:  req.GetTo(),
		pid:       req.ProtocolID(),
		user:      u,
		direction: CompleteEvent_Outgoing,
		call:      sh,
	}
}

// ^ Prepare for Incoming Session ^ //
func NewInSession(u *User, inv *InviteRequest, sh SessionHandler) *Session {
	// Return Session
	return &Session{
		file:      inv.GetFile(),
		owner:     inv.GetFrom(),
		receiver:  inv.GetTo(),
		pid:       inv.ProtocolID(),
		user:      u,
		direction: CompleteEvent_Incoming,
		call:      sh,
	}
}

// ^ Returns SFile as TransferCard given Receiver and Owner
func (s *Session) Event() *CompleteEvent {
	return &CompleteEvent{
		Direction: s.direction,
		Transfer: &Transfer{
			// SQL Properties
			Payload:  s.file.Payload,
			Received: int32(time.Now().Unix()),

			// Owner Properties
			Owner:    s.owner.GetProfile(),
			Receiver: s.receiver.GetProfile(),

			// Data Properties
			Data: s.file.ToData(),
		},
	}
}

// ^ read buffers sent on stream and save to file ^ //
func (s *Session) ReadFromStream(stream network.Stream) {
	// Concurrent Function
	go func(rs msg.ReadCloser) {
		// Read All Files
		for i, m := range s.file.Items {
			r := m.NewReader(s.user.Device, i, int(s.file.GetCount()), s.call)
			err := r.ReadFrom(rs)
			if err != nil {
				s.call.OnError(NewError(err, ErrorMessage_INCOMING))
			}
		}
		// Set Status
		s.call.OnCompleted(stream, s.pid, s.Event())
	}(msg.NewReader(stream))
}

// ^ write file as Base64 in Msgio to Stream ^ //
func (s *Session) WriteToStream(stream network.Stream) {
	// Concurrent Function
	go func(ws msg.WriteCloser) {
		// Write All Files
		for i, m := range s.file.Items {
			w := m.NewWriter(s.user.Device, i, int(s.file.GetCount()), s.call)
			err := w.WriteTo(ws)
			if err != nil {
				s.call.OnError(NewError(err, ErrorMessage_OUTGOING))
			}
		}
		// Handle Complete
		s.call.OnCompleted(stream, s.pid, s.Event())
	}(msg.NewWriter(stream))
}

// ** ─── SFile_Item MANAGEMENT ────────────────────────────────────────────────────────

func (i *SFile_Item) NewReader(d *Device, index int, total int, c SessionHandler) ItemReader {
	// Return Reader
	return &itemReader{
		item:     i,
		device:   d,
		size:     0,
		callback: c,
		index:    index,
		total:    total,
	}
}

func (m *SFile_Item) NewWriter(d *Device, index int, total int, sh SessionHandler) ItemWriter {
	return &itemWriter{
		item:     m,
		size:     0,
		device:   d,
		callback: sh,
		index:    index,
		total:    total,
	}
}

func (i *SFile_Item) SetPath(d *Device) string {
	// Set Path
	if i.Mime.IsMedia() {
		// Check for Desktop
		if d.IsDesktop() {
			i.Path = filepath.Join(d.FileSystem.GetDownloads().GetPath(), i.Name)
		} else {
			i.Path = filepath.Join(d.FileSystem.GetTemporary().GetPath(), i.Name)
		}
	} else {
		i.Path = filepath.Join(d.FileSystem.GetDownloads().GetPath(), i.Name)
	}
	return i.Path
}

// ** ─── Transfer (Reader) MANAGEMENT ────────────────────────────────────────────────────────
type ItemReader interface {
	Progress() []byte
	ReadFrom(reader msg.ReadCloser) error
}
type itemReader struct {
	ItemReader
	callback SessionHandler
	mutex    sync.Mutex
	item     *SFile_Item
	device   *Device
	index    int
	size     int
	total    int
}

// Returns Progress of File, Given the written number of bytes
func (p *itemReader) Progress() []byte {
	// Calculate Tracking
	currentProgress := float32(p.size) / float32(p.item.Size)

	// Create Update
	update := &ProgressEvent{
		Progress: float64(currentProgress),
		Current:  int32(p.index),
		Total:    int32(p.total),
	}

	return update.ToGeneric()
}

func (ir *itemReader) ReadFrom(reader msg.ReadCloser) error {
	// Return Created File
	f, err := os.Create(ir.item.SetPath(ir.device))
	if err != nil {
		return err
	}

	// Route Data from Stream
	for i := 0; ; i++ {
		// Read Length Fixed Bytes
		buffer, err := reader.ReadMsg()
		if err != nil {
			return err
		}

		done, buf, err := decodeChunk(buffer)
		if err != nil {
			return err
		}

		if !done {
			ir.mutex.Lock()
			n, err := f.Write(buf)
			if err != nil {
				return err
			}
			ir.size = ir.size + n
			ir.mutex.Unlock()

			if i%10 == 0 {
				go ir.callback.OnProgress(ir.Progress())
			}
		} else {
			// Flush File Data
			if err := f.Sync(); err != nil {
				return err
			}

			// Close File
			if err := f.Close(); err != nil {
				return err
			}
			return nil
		}
		GetState().NeedsWait()
	}
}

// ** ─── Transfer (Writer) MANAGEMENT ────────────────────────────────────────────────────────
type ItemWriter interface {
	Progress() []byte
	WriteTo(writer msg.WriteCloser) error
}

type itemWriter struct {
	ItemWriter
	callback SessionHandler
	item     *SFile_Item
	device   *Device
	index    int
	size     int
	total    int
}

// Returns Progress of File, Given the written number of bytes
func (p *itemWriter) Progress() []byte {
	// Calculate Tracking
	currentProgress := float32(p.size) / float32(p.item.Size)

	// Create Update
	update := &ProgressEvent{
		Progress: float64(currentProgress),
		Current:  int32(p.index),
		Total:    int32(p.total),
	}
	return update.ToGeneric()
}

func (iw *itemWriter) WriteTo(writer msg.WriteCloser) error {
	// Write Item to Stream
	// @ Open Os File
	f, err := os.Open(iw.item.Path)
	if err != nil {
		return errors.New(fmt.Sprintf("Error to read Item, %s", err.Error()))
	}

	// @ Initialize Chunk Data
	r := bufio.NewReader(f)
	buf := make([]byte, 0, util.TRANSFER_CHUNK_SIZE)

	// @ Loop through File
	for i := 0; ; i++ {
		// Initialize
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		// Bytes read is zero
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			return err
		}

		// Create Block Protobuf from Chunk
		data := encodeChunk(buf, false)

		// Write Message Bytes to Stream
		err = writer.WriteMsg(data)
		if err != nil {
			return err
		}

		// Unexpected Error
		if err != nil && err != io.EOF {
			return err
		}

		if i%10 == 0 {
			go iw.callback.OnProgress(iw.Progress())
		}
	}

	// Create Block Protobuf from Chunk
	data := encodeChunk(nil, true)

	// Write Message Bytes to Stream
	err = writer.WriteMsg(data)
	if err != nil {
		return err
	}
	return nil
}

func decodeChunk(data []byte) (bool, []byte, error) {
	// Unmarshal Bytes into Proto
	c, err := new(Chunk).Unmarshal(data)
	if err != nil {
		return false, nil, err
	}

	if c.IsComplete {
		return true, nil, nil
	} else {
		return false, c.Buffer, nil
	}
}

func encodeChunk(buffer []byte, completed bool) []byte {
	if !completed {
		// Create Block Protobuf from Chunk
		chunk := &Chunk{
			Size:       int32(len(buffer)),
			Buffer:     buffer,
			IsComplete: false,
		}
		return chunk.Marshal()
	} else {
		// Create Block Protobuf from Chunk
		chunk := &Chunk{
			IsComplete: true,
		}

		return chunk.Marshal()
	}
}