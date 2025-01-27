// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bucket/where_is.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BucketRole is the role of the creator of the bucket.
type BucketRole int32

const (
	// Bucket does not have role set.
	BucketRole_NONE BucketRole = 0
	// Bucket is visible to anyone.
	BucketRole_APPLICATION BucketRole = 1
	// Bucket is visible to anyone who has access token.
	BucketRole_USER BucketRole = 2
)

var BucketRole_name = map[int32]string{
	0: "NONE",
	1: "APPLICATION",
	2: "USER",
}

var BucketRole_value = map[string]int32{
	"NONE":        0,
	"APPLICATION": 1,
	"USER":        2,
}

func (x BucketRole) String() string {
	return proto.EnumName(BucketRole_name, int32(x))
}

func (BucketRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{0}
}

// BucketVisibility is the visibility of the bucket to authorized users of an application
type BucketVisibility int32

const (
	// Bucket does not have visibility set.
	BucketVisibility_UNSPECIFIED BucketVisibility = 0
	// Bucket is visible to anyone.
	BucketVisibility_PUBLIC BucketVisibility = 1
	// Bucket is visible to anyone who has access token.
	BucketVisibility_PRIVATE BucketVisibility = 2
)

var BucketVisibility_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "PUBLIC",
	2: "PRIVATE",
}

var BucketVisibility_value = map[string]int32{
	"UNSPECIFIED": 0,
	"PUBLIC":      1,
	"PRIVATE":     2,
}

func (x BucketVisibility) String() string {
	return proto.EnumName(BucketVisibility_name, int32(x))
}

func (BucketVisibility) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{1}
}

type ResourceIdentifier int32

const (
	ResourceIdentifier_None ResourceIdentifier = 0
	ResourceIdentifier_DID  ResourceIdentifier = 1
	ResourceIdentifier_CID  ResourceIdentifier = 2
)

var ResourceIdentifier_name = map[int32]string{
	0: "None",
	1: "DID",
	2: "CID",
}

var ResourceIdentifier_value = map[string]int32{
	"None": 0,
	"DID":  1,
	"CID":  2,
}

func (x ResourceIdentifier) String() string {
	return proto.EnumName(ResourceIdentifier_name, int32(x))
}

func (ResourceIdentifier) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{2}
}

type WhereIs struct {
	// DID of the created bucket.
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Creator of the new bucket
	Creator string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	// Label of the new bucket.
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	// Visibility of the new bucket.
	Visibility BucketVisibility `protobuf:"varint,4,opt,name=visibility,proto3,enum=sonrio.sonr.bucket.BucketVisibility" json:"visibility,omitempty"`
	// Role of the creator of the new bucket.
	Role BucketRole `protobuf:"varint,5,opt,name=role,proto3,enum=sonrio.sonr.bucket.BucketRole" json:"role,omitempty"`
	// IsActive flag of the new bucket.
	IsActive bool `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Content of the new bucket map of DIDs to CIDs.
	Content []*BucketItem `protobuf:"bytes,7,rep,name=content,proto3" json:"content,omitempty"`
	// JWKs of the new bucket.
	ContentAcl map[string]*AclJwks `protobuf:"bytes,8,rep,name=content_acl,json=contentAcl,proto3" json:"content_acl,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Timestamp of the new bucket.
	Timestamp int64 `protobuf:"varint,9,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *WhereIs) Reset()         { *m = WhereIs{} }
func (m *WhereIs) String() string { return proto.CompactTextString(m) }
func (*WhereIs) ProtoMessage()    {}
func (*WhereIs) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{0}
}
func (m *WhereIs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhereIs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhereIs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhereIs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhereIs.Merge(m, src)
}
func (m *WhereIs) XXX_Size() int {
	return m.Size()
}
func (m *WhereIs) XXX_DiscardUnknown() {
	xxx_messageInfo_WhereIs.DiscardUnknown(m)
}

var xxx_messageInfo_WhereIs proto.InternalMessageInfo

func (m *WhereIs) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *WhereIs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *WhereIs) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *WhereIs) GetVisibility() BucketVisibility {
	if m != nil {
		return m.Visibility
	}
	return BucketVisibility_UNSPECIFIED
}

func (m *WhereIs) GetRole() BucketRole {
	if m != nil {
		return m.Role
	}
	return BucketRole_NONE
}

func (m *WhereIs) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *WhereIs) GetContent() []*BucketItem {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *WhereIs) GetContentAcl() map[string]*AclJwks {
	if m != nil {
		return m.ContentAcl
	}
	return nil
}

func (m *WhereIs) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// AclJwks is a map of JWKs for a given Schema
type AclJwks struct {
	EncryptedJwks map[string][]byte `protobuf:"bytes,1,rep,name=encrypted_jwks,json=encryptedJwks,proto3" json:"encrypted_jwks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *AclJwks) Reset()         { *m = AclJwks{} }
func (m *AclJwks) String() string { return proto.CompactTextString(m) }
func (*AclJwks) ProtoMessage()    {}
func (*AclJwks) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{1}
}
func (m *AclJwks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AclJwks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AclJwks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AclJwks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AclJwks.Merge(m, src)
}
func (m *AclJwks) XXX_Size() int {
	return m.Size()
}
func (m *AclJwks) XXX_DiscardUnknown() {
	xxx_messageInfo_AclJwks.DiscardUnknown(m)
}

var xxx_messageInfo_AclJwks proto.InternalMessageInfo

func (m *AclJwks) GetEncryptedJwks() map[string][]byte {
	if m != nil {
		return m.EncryptedJwks
	}
	return nil
}

// BucketItem is a pointer to a single piece of content in a bucket.
type BucketItem struct {
	// Name of the bucket item.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// DID of the item. If applicable // optional
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`
	// Timestamp of the last update. // optional
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Type of the resource
	Type ResourceIdentifier `protobuf:"varint,4,opt,name=type,proto3,enum=sonrio.sonr.bucket.ResourceIdentifier" json:"type,omitempty"`
	// References the schema and item within the bucket is associated with.
	// Bucket items do not need to use the same schema if the bucket is "generic" // optional
	SchemaDid string `protobuf:"bytes,5,opt,name=schema_did,json=schemaDid,proto3" json:"schema_did,omitempty"`
}

func (m *BucketItem) Reset()         { *m = BucketItem{} }
func (m *BucketItem) String() string { return proto.CompactTextString(m) }
func (*BucketItem) ProtoMessage()    {}
func (*BucketItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d51d2bd3c178d08, []int{2}
}
func (m *BucketItem) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BucketItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BucketItem.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BucketItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BucketItem.Merge(m, src)
}
func (m *BucketItem) XXX_Size() int {
	return m.Size()
}
func (m *BucketItem) XXX_DiscardUnknown() {
	xxx_messageInfo_BucketItem.DiscardUnknown(m)
}

var xxx_messageInfo_BucketItem proto.InternalMessageInfo

func (m *BucketItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BucketItem) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *BucketItem) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *BucketItem) GetType() ResourceIdentifier {
	if m != nil {
		return m.Type
	}
	return ResourceIdentifier_None
}

func (m *BucketItem) GetSchemaDid() string {
	if m != nil {
		return m.SchemaDid
	}
	return ""
}

func init() {
	proto.RegisterEnum("sonrio.sonr.bucket.BucketRole", BucketRole_name, BucketRole_value)
	proto.RegisterEnum("sonrio.sonr.bucket.BucketVisibility", BucketVisibility_name, BucketVisibility_value)
	proto.RegisterEnum("sonrio.sonr.bucket.ResourceIdentifier", ResourceIdentifier_name, ResourceIdentifier_value)
	proto.RegisterType((*WhereIs)(nil), "sonrio.sonr.bucket.WhereIs")
	proto.RegisterMapType((map[string]*AclJwks)(nil), "sonrio.sonr.bucket.WhereIs.ContentAclEntry")
	proto.RegisterType((*AclJwks)(nil), "sonrio.sonr.bucket.AclJwks")
	proto.RegisterMapType((map[string][]byte)(nil), "sonrio.sonr.bucket.AclJwks.EncryptedJwksEntry")
	proto.RegisterType((*BucketItem)(nil), "sonrio.sonr.bucket.BucketItem")
}

func init() { proto.RegisterFile("bucket/where_is.proto", fileDescriptor_9d51d2bd3c178d08) }

var fileDescriptor_9d51d2bd3c178d08 = []byte{
	// 623 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcd, 0x4e, 0xdb, 0x40,
	0x10, 0xce, 0xc6, 0x09, 0x4e, 0x26, 0x2d, 0x58, 0x2b, 0x2a, 0x59, 0xd0, 0x5a, 0x11, 0xaa, 0x50,
	0x44, 0x55, 0xa7, 0xd0, 0x0b, 0x42, 0x3d, 0x34, 0x24, 0xae, 0xe4, 0x0a, 0x85, 0x68, 0x21, 0xb4,
	0xe2, 0x12, 0x39, 0xce, 0xb4, 0x6c, 0xe3, 0xd8, 0x91, 0xbd, 0x81, 0xe6, 0x2d, 0xfa, 0x0c, 0xbd,
	0xf6, 0x45, 0x7a, 0xe4, 0xd8, 0x63, 0x05, 0x7d, 0x90, 0x6a, 0x37, 0x36, 0xa4, 0x90, 0x72, 0xda,
	0x99, 0xcf, 0xdf, 0xb7, 0xf3, 0xeb, 0x85, 0x27, 0xfd, 0x89, 0x3f, 0x44, 0x51, 0xbf, 0x38, 0xc3,
	0x18, 0x7b, 0x3c, 0xb1, 0xc7, 0x71, 0x24, 0x22, 0x4a, 0x93, 0x28, 0x8c, 0x79, 0x64, 0xcb, 0xc3,
	0x9e, 0x51, 0x36, 0xfe, 0x68, 0xa0, 0x7f, 0x90, 0x34, 0x37, 0xa1, 0x06, 0x68, 0x03, 0x3e, 0x30,
	0x49, 0x95, 0xd4, 0xca, 0x4c, 0x9a, 0xd4, 0x04, 0xdd, 0x8f, 0xd1, 0x13, 0x51, 0x6c, 0xe6, 0x15,
	0x9a, 0xb9, 0x74, 0x15, 0x8a, 0x81, 0xd7, 0xc7, 0xc0, 0xd4, 0x14, 0x3e, 0x73, 0x68, 0x0b, 0xe0,
	0x9c, 0x27, 0xbc, 0xcf, 0x03, 0x2e, 0xa6, 0x66, 0xa1, 0x4a, 0x6a, 0xcb, 0x3b, 0xcf, 0xed, 0xfb,
	0x61, 0xed, 0x7d, 0x75, 0x9c, 0xdc, 0x70, 0xd9, 0x9c, 0x8e, 0xee, 0x40, 0x21, 0x8e, 0x02, 0x34,
	0x8b, 0x4a, 0x6f, 0xfd, 0x5f, 0xcf, 0xa2, 0x00, 0x99, 0xe2, 0xd2, 0x75, 0x28, 0xf3, 0xa4, 0xe7,
	0xf9, 0x82, 0x9f, 0xa3, 0xb9, 0x54, 0x25, 0xb5, 0x12, 0x2b, 0xf1, 0xa4, 0xa1, 0x7c, 0xba, 0x0b,
	0xba, 0x1f, 0x85, 0x02, 0x43, 0x61, 0xea, 0x55, 0xad, 0x56, 0x79, 0xe8, 0x4e, 0x57, 0xe0, 0x88,
	0x65, 0x74, 0x7a, 0x00, 0x95, 0xd4, 0xec, 0x79, 0x7e, 0x60, 0x96, 0x94, 0xfa, 0xc5, 0x22, 0x75,
	0xda, 0x44, 0xbb, 0x39, 0xa3, 0x37, 0xfc, 0xc0, 0x09, 0x45, 0x3c, 0x65, 0xe0, 0xdf, 0x00, 0xf4,
	0x29, 0x94, 0x05, 0x1f, 0x61, 0x22, 0xbc, 0xd1, 0xd8, 0x2c, 0x57, 0x49, 0x4d, 0x63, 0xb7, 0xc0,
	0xda, 0x29, 0xac, 0xdc, 0x11, 0xcb, 0x89, 0x0c, 0x71, 0x9a, 0x4d, 0x64, 0x88, 0x53, 0xba, 0x0d,
	0xc5, 0x73, 0x2f, 0x98, 0xa0, 0x9a, 0x47, 0x65, 0x67, 0x7d, 0x51, 0x2a, 0x0d, 0x3f, 0x78, 0x7f,
	0x31, 0x4c, 0xd8, 0x8c, 0xb9, 0x97, 0xdf, 0x25, 0x1b, 0xdf, 0x09, 0xe8, 0x29, 0x4c, 0xbb, 0xb0,
	0x8c, 0xa1, 0x1f, 0x4f, 0xc7, 0x02, 0x07, 0xbd, 0x2f, 0x17, 0xc3, 0xc4, 0x24, 0xaa, 0x2c, 0xfb,
	0x81, 0xbb, 0x6c, 0x27, 0x53, 0x48, 0x6f, 0x56, 0xd9, 0x63, 0x9c, 0xc7, 0xd6, 0xde, 0x02, 0xbd,
	0x4f, 0x5a, 0x50, 0xc1, 0xea, 0x7c, 0x05, 0x8f, 0xe6, 0x93, 0xfc, 0x41, 0x00, 0x6e, 0x87, 0x40,
	0x29, 0x14, 0x42, 0x6f, 0x84, 0xa9, 0x56, 0xd9, 0xf2, 0xba, 0x49, 0xcc, 0xd3, 0x65, 0x94, 0xe6,
	0xbf, 0x3d, 0xd5, 0xee, 0xf4, 0x94, 0xee, 0x41, 0x41, 0x4c, 0xc7, 0x98, 0xae, 0xe2, 0xe6, 0xa2,
	0x0a, 0x19, 0x26, 0xd1, 0x24, 0xf6, 0xd1, 0x1d, 0x60, 0x28, 0xf8, 0x27, 0x8e, 0x31, 0x53, 0x1a,
	0xfa, 0x0c, 0x20, 0xf1, 0xcf, 0x70, 0xe4, 0xf5, 0xe4, 0x5f, 0x51, 0x54, 0x21, 0xcb, 0x33, 0xa4,
	0xc5, 0x07, 0x5b, 0xdb, 0x59, 0xb2, 0x72, 0x0b, 0x69, 0x09, 0x0a, 0xed, 0xc3, 0xb6, 0x63, 0xe4,
	0xe8, 0x0a, 0x54, 0x1a, 0x9d, 0xce, 0x81, 0xdb, 0x6c, 0x1c, 0xbb, 0x87, 0x6d, 0x83, 0xc8, 0x4f,
	0xdd, 0x23, 0x87, 0x19, 0xf9, 0xad, 0x37, 0x60, 0xdc, 0x5d, 0x7c, 0x49, 0xef, 0xb6, 0x8f, 0x3a,
	0x4e, 0xd3, 0x7d, 0xe7, 0x3a, 0x2d, 0x23, 0x47, 0x01, 0x96, 0x3a, 0xdd, 0xfd, 0x03, 0xb7, 0x69,
	0x10, 0x5a, 0x01, 0xbd, 0xc3, 0xdc, 0x93, 0xc6, 0xb1, 0x63, 0xe4, 0xb7, 0x5e, 0x01, 0xbd, 0x9f,
	0xab, 0x0a, 0x1c, 0x85, 0x68, 0xe4, 0xa8, 0x0e, 0x5a, 0xcb, 0x6d, 0x19, 0x44, 0x1a, 0x4d, 0xb7,
	0x65, 0xe4, 0xf7, 0x3f, 0xfe, 0xbc, 0xb2, 0xc8, 0xe5, 0x95, 0x45, 0x7e, 0x5f, 0x59, 0xe4, 0xdb,
	0xb5, 0x95, 0xbb, 0xbc, 0xb6, 0x72, 0xbf, 0xae, 0xad, 0x1c, 0xac, 0x66, 0x4d, 0x90, 0x95, 0x26,
	0x69, 0x2b, 0x3a, 0xe4, 0x74, 0xf3, 0x33, 0x17, 0x67, 0x93, 0xbe, 0xed, 0x47, 0xa3, 0xba, 0xfc,
	0xfe, 0x92, 0x47, 0xea, 0xac, 0x7f, 0xad, 0xa7, 0x6f, 0x8a, 0x12, 0xf4, 0x97, 0xd4, 0x8b, 0xf2,
	0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x38, 0x13, 0x41, 0x6a, 0x04, 0x00, 0x00,
}

func (m *WhereIs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhereIs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhereIs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != 0 {
		i = encodeVarintWhereIs(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x48
	}
	if len(m.ContentAcl) > 0 {
		for k := range m.ContentAcl {
			v := m.ContentAcl[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintWhereIs(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWhereIs(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWhereIs(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Content) > 0 {
		for iNdEx := len(m.Content) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Content[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintWhereIs(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.Role != 0 {
		i = encodeVarintWhereIs(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x28
	}
	if m.Visibility != 0 {
		i = encodeVarintWhereIs(dAtA, i, uint64(m.Visibility))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AclJwks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AclJwks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AclJwks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptedJwks) > 0 {
		for k := range m.EncryptedJwks {
			v := m.EncryptedJwks[k]
			baseI := i
			if len(v) > 0 {
				i -= len(v)
				copy(dAtA[i:], v)
				i = encodeVarintWhereIs(dAtA, i, uint64(len(v)))
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWhereIs(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWhereIs(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *BucketItem) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BucketItem) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BucketItem) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SchemaDid) > 0 {
		i -= len(m.SchemaDid)
		copy(dAtA[i:], m.SchemaDid)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.SchemaDid)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Type != 0 {
		i = encodeVarintWhereIs(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x20
	}
	if m.Timestamp != 0 {
		i = encodeVarintWhereIs(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWhereIs(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhereIs(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhereIs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhereIs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	if m.Visibility != 0 {
		n += 1 + sovWhereIs(uint64(m.Visibility))
	}
	if m.Role != 0 {
		n += 1 + sovWhereIs(uint64(m.Role))
	}
	if m.IsActive {
		n += 2
	}
	if len(m.Content) > 0 {
		for _, e := range m.Content {
			l = e.Size()
			n += 1 + l + sovWhereIs(uint64(l))
		}
	}
	if len(m.ContentAcl) > 0 {
		for k, v := range m.ContentAcl {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovWhereIs(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovWhereIs(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovWhereIs(uint64(mapEntrySize))
		}
	}
	if m.Timestamp != 0 {
		n += 1 + sovWhereIs(uint64(m.Timestamp))
	}
	return n
}

func (m *AclJwks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.EncryptedJwks) > 0 {
		for k, v := range m.EncryptedJwks {
			_ = k
			_ = v
			l = 0
			if len(v) > 0 {
				l = 1 + len(v) + sovWhereIs(uint64(len(v)))
			}
			mapEntrySize := 1 + len(k) + sovWhereIs(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovWhereIs(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *BucketItem) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovWhereIs(uint64(m.Timestamp))
	}
	if m.Type != 0 {
		n += 1 + sovWhereIs(uint64(m.Type))
	}
	l = len(m.SchemaDid)
	if l > 0 {
		n += 1 + l + sovWhereIs(uint64(l))
	}
	return n
}

func sovWhereIs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhereIs(x uint64) (n int) {
	return sovWhereIs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhereIs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhereIs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WhereIs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhereIs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Visibility", wireType)
			}
			m.Visibility = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Visibility |= BucketVisibility(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= BucketRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content, &BucketItem{})
			if err := m.Content[len(m.Content)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentAcl", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ContentAcl == nil {
				m.ContentAcl = make(map[string]*AclJwks)
			}
			var mapkey string
			var mapvalue *AclJwks
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWhereIs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhereIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWhereIs
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWhereIs
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhereIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthWhereIs
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthWhereIs
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &AclJwks{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWhereIs(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthWhereIs
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ContentAcl[mapkey] = mapvalue
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWhereIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhereIs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AclJwks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhereIs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AclJwks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AclJwks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedJwks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EncryptedJwks == nil {
				m.EncryptedJwks = make(map[string][]byte)
			}
			var mapkey string
			mapvalue := []byte{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWhereIs
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhereIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWhereIs
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWhereIs
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapbyteLen uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhereIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapbyteLen |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intMapbyteLen := int(mapbyteLen)
					if intMapbyteLen < 0 {
						return ErrInvalidLengthWhereIs
					}
					postbytesIndex := iNdEx + intMapbyteLen
					if postbytesIndex < 0 {
						return ErrInvalidLengthWhereIs
					}
					if postbytesIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = make([]byte, mapbyteLen)
					copy(mapvalue, dAtA[iNdEx:postbytesIndex])
					iNdEx = postbytesIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWhereIs(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthWhereIs
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.EncryptedJwks[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhereIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhereIs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BucketItem) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhereIs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BucketItem: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BucketItem: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= ResourceIdentifier(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthWhereIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhereIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SchemaDid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhereIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhereIs
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWhereIs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhereIs
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWhereIs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWhereIs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhereIs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhereIs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhereIs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhereIs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhereIs = fmt.Errorf("proto: unexpected end of group")
)
