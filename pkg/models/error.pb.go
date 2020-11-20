// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: error.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Level of Error Issue
type Error_Level int32

const (
	Error_NONE     Error_Level = 0
	Error_MINOR    Error_Level = 1
	Error_WARNING  Error_Level = 2
	Error_MAJOR    Error_Level = 3
	Error_CRITICAL Error_Level = 4
	Error_PANIC    Error_Level = 5
)

// Enum value maps for Error_Level.
var (
	Error_Level_name = map[int32]string{
		0: "NONE",
		1: "MINOR",
		2: "WARNING",
		3: "MAJOR",
		4: "CRITICAL",
		5: "PANIC",
	}
	Error_Level_value = map[string]int32{
		"NONE":     0,
		"MINOR":    1,
		"WARNING":  2,
		"MAJOR":    3,
		"CRITICAL": 4,
		"PANIC":    5,
	}
)

func (x Error_Level) Enum() *Error_Level {
	p := new(Error_Level)
	*p = x
	return p
}

func (x Error_Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Level) Descriptor() protoreflect.EnumDescriptor {
	return file_error_proto_enumTypes[0].Descriptor()
}

func (Error_Level) Type() protoreflect.EnumType {
	return &file_error_proto_enumTypes[0]
}

func (x Error_Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Level.Descriptor instead.
func (Error_Level) EnumDescriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0, 0}
}

// Kind of Error Issue
type Error_Kind int32

const (
	Error_NETWORK   Error_Kind = 0
	Error_LOBBY     Error_Kind = 1
	Error_DISCOVERY Error_Kind = 2
	Error_PROTO     Error_Kind = 3
	Error_BYTES     Error_Kind = 4
	Error_JSON      Error_Kind = 5
	Error_PEER      Error_Kind = 6
	Error_BADGER    Error_Kind = 7
	Error_INFO      Error_Kind = 8
	Error_BUFFER    Error_Kind = 9
)

// Enum value maps for Error_Kind.
var (
	Error_Kind_name = map[int32]string{
		0: "NETWORK",
		1: "LOBBY",
		2: "DISCOVERY",
		3: "PROTO",
		4: "BYTES",
		5: "JSON",
		6: "PEER",
		7: "BADGER",
		8: "INFO",
		9: "BUFFER",
	}
	Error_Kind_value = map[string]int32{
		"NETWORK":   0,
		"LOBBY":     1,
		"DISCOVERY": 2,
		"PROTO":     3,
		"BYTES":     4,
		"JSON":      5,
		"PEER":      6,
		"BADGER":    7,
		"INFO":      8,
		"BUFFER":    9,
	}
)

func (x Error_Kind) Enum() *Error_Kind {
	p := new(Error_Kind)
	*p = x
	return p
}

func (x Error_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_error_proto_enumTypes[1].Descriptor()
}

func (Error_Kind) Type() protoreflect.EnumType {
	return &file_error_proto_enumTypes[1]
}

func (x Error_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error_Kind.Descriptor instead.
func (Error_Kind) EnumDescriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0, 1}
}

// Define Error Occurance
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind    Error_Kind  `protobuf:"varint,1,opt,name=kind,proto3,enum=Error_Kind" json:"kind,omitempty"`
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Level   Error_Level `protobuf:"varint,4,opt,name=level,proto3,enum=Error_Level" json:"level,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetKind() Error_Kind {
	if x != nil {
		return x.Kind
	}
	return Error_NETWORK
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetLevel() Error_Level {
	if x != nil {
		return x.Level
	}
	return Error_NONE
}

var File_error_proto protoreflect.FileDescriptor

var file_error_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x02,
	0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4b, 0x69,
	0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x4d, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x49, 0x4e,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x41, 0x4a, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x41,
	0x4e, 0x49, 0x43, 0x10, 0x05, 0x22, 0x79, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a,
	0x07, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x4f,
	0x42, 0x42, 0x59, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56, 0x45,
	0x52, 0x59, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x10, 0x03, 0x12,
	0x09, 0x0a, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x53,
	0x4f, 0x4e, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x45, 0x45, 0x52, 0x10, 0x06, 0x12, 0x0a,
	0x0a, 0x06, 0x42, 0x41, 0x44, 0x47, 0x45, 0x52, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e,
	0x46, 0x4f, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x55, 0x46, 0x46, 0x45, 0x52, 0x10, 0x09,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func file_error_proto_rawDescGZIP() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_error_proto_goTypes = []interface{}{
	(Error_Level)(0), // 0: Error.Level
	(Error_Kind)(0),  // 1: Error.Kind
	(*Error)(nil),    // 2: Error
}
var file_error_proto_depIdxs = []int32{
	1, // 0: Error.kind:type_name -> Error.Kind
	0, // 1: Error.level:type_name -> Error.Level
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_error_proto_init() }
func file_error_proto_init() {
	if File_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_error_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_proto_goTypes,
		DependencyIndexes: file_error_proto_depIdxs,
		EnumInfos:         file_error_proto_enumTypes,
		MessageInfos:      file_error_proto_msgTypes,
	}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}