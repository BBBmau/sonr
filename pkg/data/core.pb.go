// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: core.proto

package data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Sonr Default Protocols
type SonrProtocol int32

const (
	SonrProtocol_Unspecified    SonrProtocol = 0 // Default Value
	SonrProtocol_Authorize      SonrProtocol = 1 // Auth Protocol for Invite/Response
	SonrProtocol_Devices        SonrProtocol = 2 // Device Protocol for User Linked Devices
	SonrProtocol_Linker         SonrProtocol = 3 // Protocol for Linking Device to User
	SonrProtocol_LocalTransfer  SonrProtocol = 4 // Protocol for Local Transfer to Peer
	SonrProtocol_RemoteTransfer SonrProtocol = 5 // Protocol for Remote Transfer to Peer
	SonrProtocol_Groups         SonrProtocol = 6 // Protocol for Group Chat
)

// Enum value maps for SonrProtocol.
var (
	SonrProtocol_name = map[int32]string{
		0: "Unspecified",
		1: "Authorize",
		2: "Devices",
		3: "Linker",
		4: "LocalTransfer",
		5: "RemoteTransfer",
		6: "Groups",
	}
	SonrProtocol_value = map[string]int32{
		"Unspecified":    0,
		"Authorize":      1,
		"Devices":        2,
		"Linker":         3,
		"LocalTransfer":  4,
		"RemoteTransfer": 5,
		"Groups":         6,
	}
)

func (x SonrProtocol) Enum() *SonrProtocol {
	p := new(SonrProtocol)
	*p = x
	return p
}

func (x SonrProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SonrProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_core_proto_enumTypes[0].Descriptor()
}

func (SonrProtocol) Type() protoreflect.EnumType {
	return &file_core_proto_enumTypes[0]
}

func (x SonrProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SonrProtocol.Descriptor instead.
func (SonrProtocol) EnumDescriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

// Sent on Data Transfer to Add piece of File - Binary
type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Buffer     []byte `protobuf:"bytes,1,opt,name=buffer,proto3" json:"buffer,omitempty"`          // Binary Value of Chunk
	Size       int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`             // Size of this Chunk
	IsComplete bool   `protobuf:"varint,3,opt,name=isComplete,proto3" json:"isComplete,omitempty"` // Whether this is final chunk
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{0}
}

func (x *Chunk) GetBuffer() []byte {
	if x != nil {
		return x.Buffer
	}
	return nil
}

func (x *Chunk) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Chunk) GetIsComplete() bool {
	if x != nil {
		return x.IsComplete
	}
	return false
}

type PushMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peer *Peer             `protobuf:"bytes,1,opt,name=peer,proto3" json:"peer,omitempty"`                                                                                         // Peer to send to
	Data map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Data to send
}

func (x *PushMessage) Reset() {
	*x = PushMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessage) ProtoMessage() {}

func (x *PushMessage) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessage.ProtoReflect.Descriptor instead.
func (*PushMessage) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{1}
}

func (x *PushMessage) GetPeer() *Peer {
	if x != nil {
		return x.Peer
	}
	return nil
}

func (x *PushMessage) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

type LinkPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Primary   *Device   `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`     // Primary Device
	Secondary *Device   `protobuf:"bytes,2,opt,name=secondary,proto3" json:"secondary,omitempty"` // Secondary Device
	KeyChain  *KeyChain `protobuf:"bytes,3,opt,name=keyChain,proto3" json:"keyChain,omitempty"`   // KeyChain for Linking
}

func (x *LinkPacket) Reset() {
	*x = LinkPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkPacket) ProtoMessage() {}

func (x *LinkPacket) ProtoReflect() protoreflect.Message {
	mi := &file_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkPacket.ProtoReflect.Descriptor instead.
func (*LinkPacket) Descriptor() ([]byte, []int) {
	return file_core_proto_rawDescGZIP(), []int{2}
}

func (x *LinkPacket) GetPrimary() *Device {
	if x != nil {
		return x.Primary
	}
	return nil
}

func (x *LinkPacket) GetSecondary() *Device {
	if x != nil {
		return x.Secondary
	}
	return nil
}

func (x *LinkPacket) GetKeyChain() *KeyChain {
	if x != nil {
		return x.KeyChain
	}
	return nil
}

var File_core_proto protoreflect.FileDescriptor

var file_core_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x0a, 0x70, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x05, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x22,
	0x97, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x70, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x12,
	0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8c, 0x01, 0x0a, 0x0a, 0x4c, 0x69,
	0x6e, 0x6b, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79,
	0x12, 0x2a, 0x0a, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x09, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x61, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x08,
	0x6b, 0x65, 0x79, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x2a, 0x7a, 0x0a, 0x0c, 0x53, 0x6f, 0x6e, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x10,
	0x03, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x73, 0x10, 0x06, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_proto_rawDescOnce sync.Once
	file_core_proto_rawDescData = file_core_proto_rawDesc
)

func file_core_proto_rawDescGZIP() []byte {
	file_core_proto_rawDescOnce.Do(func() {
		file_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_proto_rawDescData)
	})
	return file_core_proto_rawDescData
}

var file_core_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_core_proto_goTypes = []interface{}{
	(SonrProtocol)(0),   // 0: data.SonrProtocol
	(*Chunk)(nil),       // 1: data.Chunk
	(*PushMessage)(nil), // 2: data.PushMessage
	(*LinkPacket)(nil),  // 3: data.LinkPacket
	nil,                 // 4: data.PushMessage.DataEntry
	(*Peer)(nil),        // 5: data.Peer
	(*Device)(nil),      // 6: data.Device
	(*KeyChain)(nil),    // 7: data.KeyChain
}
var file_core_proto_depIdxs = []int32{
	5, // 0: data.PushMessage.peer:type_name -> data.Peer
	4, // 1: data.PushMessage.data:type_name -> data.PushMessage.DataEntry
	6, // 2: data.LinkPacket.primary:type_name -> data.Device
	6, // 3: data.LinkPacket.secondary:type_name -> data.Device
	7, // 4: data.LinkPacket.keyChain:type_name -> data.KeyChain
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_core_proto_init() }
func file_core_proto_init() {
	if File_core_proto != nil {
		return
	}
	file_peer_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkPacket); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_proto_goTypes,
		DependencyIndexes: file_core_proto_depIdxs,
		EnumInfos:         file_core_proto_enumTypes,
		MessageInfos:      file_core_proto_msgTypes,
	}.Build()
	File_core_proto = out.File
	file_core_proto_rawDesc = nil
	file_core_proto_goTypes = nil
	file_core_proto_depIdxs = nil
}