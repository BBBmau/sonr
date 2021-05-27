// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: remote.proto

package models

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

type RemoteJoinResponse_Status int32

const (
	RemoteJoinResponse_None       RemoteJoinResponse_Status = 0
	RemoteJoinResponse_Approved   RemoteJoinResponse_Status = 1
	RemoteJoinResponse_Declined   RemoteJoinResponse_Status = 2
	RemoteJoinResponse_InProgress RemoteJoinResponse_Status = 3
	RemoteJoinResponse_Complete   RemoteJoinResponse_Status = 4
)

// Enum value maps for RemoteJoinResponse_Status.
var (
	RemoteJoinResponse_Status_name = map[int32]string{
		0: "None",
		1: "Approved",
		2: "Declined",
		3: "InProgress",
		4: "Complete",
	}
	RemoteJoinResponse_Status_value = map[string]int32{
		"None":       0,
		"Approved":   1,
		"Declined":   2,
		"InProgress": 3,
		"Complete":   4,
	}
)

func (x RemoteJoinResponse_Status) Enum() *RemoteJoinResponse_Status {
	p := new(RemoteJoinResponse_Status)
	*p = x
	return p
}

func (x RemoteJoinResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RemoteJoinResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_remote_proto_enumTypes[0].Descriptor()
}

func (RemoteJoinResponse_Status) Type() protoreflect.EnumType {
	return &file_remote_proto_enumTypes[0]
}

func (x RemoteJoinResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RemoteJoinResponse_Status.Descriptor instead.
func (RemoteJoinResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{3, 0}
}

// Remote Request to Join/Create Remote Transfer
type RemoteCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File *SonrFile `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Ttl  int32     `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *RemoteCreateRequest) Reset() {
	*x = RemoteCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteCreateRequest) ProtoMessage() {}

func (x *RemoteCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteCreateRequest.ProtoReflect.Descriptor instead.
func (*RemoteCreateRequest) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteCreateRequest) GetFile() *SonrFile {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *RemoteCreateRequest) GetTtl() int32 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

// Remote Request to Join/Create Remote Transfer
type RemoteJoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicWords []string `protobuf:"bytes,1,rep,name=topicWords,proto3" json:"topicWords,omitempty"`
}

func (x *RemoteJoinRequest) Reset() {
	*x = RemoteJoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteJoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteJoinRequest) ProtoMessage() {}

func (x *RemoteJoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteJoinRequest.ProtoReflect.Descriptor instead.
func (*RemoteJoinRequest) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteJoinRequest) GetTopicWords() []string {
	if x != nil {
		return x.TopicWords
	}
	return nil
}

// Response from Remote Request
type RemoteCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsJoin  bool     `protobuf:"varint,1,opt,name=isJoin,proto3" json:"isJoin,omitempty"`
	Count   int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Display string   `protobuf:"bytes,3,opt,name=display,proto3" json:"display,omitempty"`
	Topic   string   `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Words   []string `protobuf:"bytes,5,rep,name=words,proto3" json:"words,omitempty"`
}

func (x *RemoteCreateResponse) Reset() {
	*x = RemoteCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteCreateResponse) ProtoMessage() {}

func (x *RemoteCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteCreateResponse.ProtoReflect.Descriptor instead.
func (*RemoteCreateResponse) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{2}
}

func (x *RemoteCreateResponse) GetIsJoin() bool {
	if x != nil {
		return x.IsJoin
	}
	return false
}

func (x *RemoteCreateResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *RemoteCreateResponse) GetDisplay() string {
	if x != nil {
		return x.Display
	}
	return ""
}

func (x *RemoteCreateResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *RemoteCreateResponse) GetWords() []string {
	if x != nil {
		return x.Words
	}
	return nil
}

// Response from Remote Request
type RemoteJoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  RemoteJoinResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=RemoteJoinResponse_Status" json:"status,omitempty"`
	IsValid bool                      `protobuf:"varint,2,opt,name=isValid,proto3" json:"isValid,omitempty"`
	Lobby   *Lobby                    `protobuf:"bytes,3,opt,name=lobby,proto3" json:"lobby,omitempty"`
}

func (x *RemoteJoinResponse) Reset() {
	*x = RemoteJoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remote_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteJoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteJoinResponse) ProtoMessage() {}

func (x *RemoteJoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_remote_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteJoinResponse.ProtoReflect.Descriptor instead.
func (*RemoteJoinResponse) Descriptor() ([]byte, []int) {
	return file_remote_proto_rawDescGZIP(), []int{3}
}

func (x *RemoteJoinResponse) GetStatus() RemoteJoinResponse_Status {
	if x != nil {
		return x.Status
	}
	return RemoteJoinResponse_None
}

func (x *RemoteJoinResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

func (x *RemoteJoinResponse) GetLobby() *Lobby {
	if x != nil {
		return x.Lobby
	}
	return nil
}

var File_remote_proto protoreflect.FileDescriptor

var file_remote_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x53, 0x6f, 0x6e, 0x72, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22, 0x33, 0x0a, 0x11, 0x52,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x57, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x57, 0x6f, 0x72, 0x64, 0x73,
	0x22, 0x8a, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x4a,
	0x6f, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x4a, 0x6f, 0x69,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xce, 0x01,
	0x0a, 0x12, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x4a, 0x6f, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x12, 0x1c, 0x0a, 0x05, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x06, 0x2e, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x52, 0x05, 0x6c, 0x6f, 0x62, 0x62, 0x79,
	0x22, 0x4c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x6c, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x02,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x04, 0x42, 0x09,
	0x5a, 0x07, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_remote_proto_rawDescOnce sync.Once
	file_remote_proto_rawDescData = file_remote_proto_rawDesc
)

func file_remote_proto_rawDescGZIP() []byte {
	file_remote_proto_rawDescOnce.Do(func() {
		file_remote_proto_rawDescData = protoimpl.X.CompressGZIP(file_remote_proto_rawDescData)
	})
	return file_remote_proto_rawDescData
}

var file_remote_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_remote_proto_goTypes = []interface{}{
	(RemoteJoinResponse_Status)(0), // 0: RemoteJoinResponse.Status
	(*RemoteCreateRequest)(nil),    // 1: RemoteCreateRequest
	(*RemoteJoinRequest)(nil),      // 2: RemoteJoinRequest
	(*RemoteCreateResponse)(nil),   // 3: RemoteCreateResponse
	(*RemoteJoinResponse)(nil),     // 4: RemoteJoinResponse
	(*SonrFile)(nil),               // 5: SonrFile
	(*Lobby)(nil),                  // 6: Lobby
}
var file_remote_proto_depIdxs = []int32{
	5, // 0: RemoteCreateRequest.file:type_name -> SonrFile
	0, // 1: RemoteJoinResponse.status:type_name -> RemoteJoinResponse.Status
	6, // 2: RemoteJoinResponse.lobby:type_name -> Lobby
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_remote_proto_init() }
func file_remote_proto_init() {
	if File_remote_proto != nil {
		return
	}
	file_data_proto_init()
	file_topic_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_remote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteCreateRequest); i {
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
		file_remote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteJoinRequest); i {
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
		file_remote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteCreateResponse); i {
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
		file_remote_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteJoinResponse); i {
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
			RawDescriptor: file_remote_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_remote_proto_goTypes,
		DependencyIndexes: file_remote_proto_depIdxs,
		EnumInfos:         file_remote_proto_enumTypes,
		MessageInfos:      file_remote_proto_msgTypes,
	}.Build()
	File_remote_proto = out.File
	file_remote_proto_rawDesc = nil
	file_remote_proto_goTypes = nil
	file_remote_proto_depIdxs = nil
}
