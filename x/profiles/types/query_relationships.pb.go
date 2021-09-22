// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta2/query_relationships.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryRelationshipsRequest is the request type for the
// Query/Relationships RPC method.
type QueryRelationshipsRequest struct {
	// address of the user to query the relationships for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// subspace to query the relationships for
	SubspaceId string `protobuf:"bytes,2,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRelationshipsRequest) Reset()         { *m = QueryRelationshipsRequest{} }
func (m *QueryRelationshipsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRelationshipsRequest) ProtoMessage()    {}
func (*QueryRelationshipsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40abc27f3ab2eb88, []int{0}
}
func (m *QueryRelationshipsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRelationshipsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRelationshipsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRelationshipsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRelationshipsRequest.Merge(m, src)
}
func (m *QueryRelationshipsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRelationshipsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRelationshipsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRelationshipsRequest proto.InternalMessageInfo

// QueryRelationshipsResponse is the response type for the
// Query/Relationships RPC method.
type QueryRelationshipsResponse struct {
	Relationships []Relationship `protobuf:"bytes,1,rep,name=relationships,proto3" json:"relationships"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryRelationshipsResponse) Reset()         { *m = QueryRelationshipsResponse{} }
func (m *QueryRelationshipsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRelationshipsResponse) ProtoMessage()    {}
func (*QueryRelationshipsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40abc27f3ab2eb88, []int{1}
}
func (m *QueryRelationshipsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRelationshipsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRelationshipsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRelationshipsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRelationshipsResponse.Merge(m, src)
}
func (m *QueryRelationshipsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRelationshipsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRelationshipsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRelationshipsResponse proto.InternalMessageInfo

func (m *QueryRelationshipsResponse) GetRelationships() []Relationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

func (m *QueryRelationshipsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryBlocksRequest is the request type for the Query/Blocks RPC
// endpoint
type QueryBlocksRequest struct {
	// address of the user to query the blocks for
	User       string             `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	SubspaceId string             `protobuf:"bytes,2,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	Pagination *query.PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryBlocksRequest) Reset()         { *m = QueryBlocksRequest{} }
func (m *QueryBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryBlocksRequest) ProtoMessage()    {}
func (*QueryBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_40abc27f3ab2eb88, []int{2}
}
func (m *QueryBlocksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlocksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlocksRequest.Merge(m, src)
}
func (m *QueryBlocksRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlocksRequest proto.InternalMessageInfo

// QueryBlocksResponse is the response type for the Query/Blocks RPC
// method.
type QueryBlocksResponse struct {
	Blocks     []UserBlock         `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryBlocksResponse) Reset()         { *m = QueryBlocksResponse{} }
func (m *QueryBlocksResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBlocksResponse) ProtoMessage()    {}
func (*QueryBlocksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_40abc27f3ab2eb88, []int{3}
}
func (m *QueryBlocksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBlocksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBlocksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBlocksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBlocksResponse.Merge(m, src)
}
func (m *QueryBlocksResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBlocksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBlocksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBlocksResponse proto.InternalMessageInfo

func (m *QueryBlocksResponse) GetBlocks() []UserBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *QueryBlocksResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRelationshipsRequest)(nil), "desmos.profiles.v1beta2.QueryRelationshipsRequest")
	proto.RegisterType((*QueryRelationshipsResponse)(nil), "desmos.profiles.v1beta2.QueryRelationshipsResponse")
	proto.RegisterType((*QueryBlocksRequest)(nil), "desmos.profiles.v1beta2.QueryBlocksRequest")
	proto.RegisterType((*QueryBlocksResponse)(nil), "desmos.profiles.v1beta2.QueryBlocksResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta2/query_relationships.proto", fileDescriptor_40abc27f3ab2eb88)
}

var fileDescriptor_40abc27f3ab2eb88 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0xc6, 0xbd, 0x77, 0xa7, 0x13, 0x6c, 0x44, 0xb3, 0x20, 0x91, 0x44, 0xc8, 0x89, 0x22, 0x01,
	0x11, 0x12, 0xbb, 0x4a, 0xe8, 0xa8, 0x50, 0x0a, 0xfe, 0x74, 0x9c, 0x25, 0x1a, 0x9a, 0x68, 0xd7,
	0x9e, 0xf3, 0x59, 0x38, 0xde, 0x3d, 0x8f, 0x8d, 0xc8, 0x1b, 0x50, 0xd2, 0xd3, 0x44, 0x3c, 0x05,
	0x8f, 0x70, 0xe5, 0x95, 0x54, 0x08, 0x25, 0x0d, 0x8f, 0x81, 0xbc, 0xbb, 0xa7, 0xb3, 0xa5, 0x8b,
	0x68, 0x28, 0xe8, 0x66, 0x76, 0xe6, 0x9b, 0xf9, 0xcd, 0x97, 0x98, 0xce, 0x12, 0xc0, 0x95, 0x46,
	0x61, 0x4a, 0x7d, 0x9a, 0xe5, 0x80, 0xe2, 0xe3, 0x4c, 0x41, 0x25, 0xe7, 0xe2, 0xbc, 0x86, 0x72,
	0xbd, 0x2c, 0x21, 0x97, 0x55, 0xa6, 0x0b, 0x3c, 0xcb, 0x0c, 0x72, 0x53, 0xea, 0x4a, 0xb3, 0xfb,
	0x4e, 0xc2, 0xaf, 0x24, 0xdc, 0x4b, 0x86, 0xf7, 0x52, 0x9d, 0x6a, 0xdb, 0x23, 0x9a, 0xc8, 0xb5,
	0x0f, 0x1f, 0xa4, 0x5a, 0xa7, 0x39, 0x08, 0x69, 0x32, 0x21, 0x8b, 0x42, 0x57, 0x6e, 0xa0, 0xaf,
	0x0e, 0x7c, 0xd5, 0x66, 0xaa, 0x3e, 0x15, 0xb2, 0x58, 0xfb, 0xd2, 0x7c, 0x1f, 0xda, 0x4a, 0x27,
	0x90, 0xe3, 0x4d, 0x6c, 0xc3, 0x41, 0xac, 0x1b, 0xcd, 0xd2, 0x51, 0xb8, 0xc4, 0x97, 0x9e, 0xb8,
	0x4c, 0x28, 0x89, 0xe0, 0xae, 0xf3, 0x03, 0x67, 0xc2, 0xc8, 0x34, 0x2b, 0xec, 0x2c, 0xd7, 0x3b,
	0xf9, 0x46, 0xe8, 0xe0, 0xa4, 0x69, 0x89, 0xda, 0x3b, 0x22, 0x38, 0xaf, 0x01, 0x2b, 0xc6, 0xe8,
	0x51, 0x8d, 0x50, 0xf6, 0xc9, 0x98, 0x4c, 0x6f, 0x47, 0x36, 0x66, 0x23, 0xda, 0xc3, 0x5a, 0xa1,
	0x91, 0x31, 0x2c, 0xb3, 0xa4, 0x7f, 0x60, 0x4b, 0xf4, 0xea, 0xe9, 0x4d, 0xc2, 0x5e, 0x52, 0x7a,
	0xbd, 0xa6, 0x7f, 0x38, 0x26, 0xd3, 0xde, 0xfc, 0x11, 0xf7, 0x84, 0x0d, 0x13, 0xb7, 0x4c, 0xde,
	0xcc, 0x19, 0x7f, 0x2b, 0x53, 0xf0, 0x0b, 0xa3, 0x96, 0xf2, 0xf9, 0xad, 0xcf, 0x9b, 0x51, 0xf0,
	0x7b, 0x33, 0x0a, 0x26, 0xdf, 0x09, 0x1d, 0xde, 0x04, 0x89, 0x46, 0x17, 0x08, 0xec, 0x84, 0xde,
	0xe9, 0x38, 0xd4, 0x27, 0xe3, 0xc3, 0x69, 0x6f, 0xfe, 0x90, 0xef, 0xf9, 0xf9, 0x78, 0x7b, 0xcc,
	0xe2, 0xe8, 0xe2, 0xe7, 0x28, 0x88, 0xba, 0x13, 0xd8, 0xab, 0xce, 0x0d, 0x07, 0xf6, 0x86, 0xc7,
	0x7f, 0xbd, 0xc1, 0xf1, 0xb4, 0x8f, 0x98, 0x7c, 0x25, 0x94, 0x59, 0xf4, 0x45, 0xae, 0xe3, 0x0f,
	0xff, 0x9b, 0xb1, 0x1b, 0x42, 0xef, 0x76, 0xe8, 0xbc, 0xa3, 0x2f, 0xe8, 0xb1, 0xb2, 0x2f, 0xde,
	0xca, 0xc9, 0x5e, 0x2b, 0xdf, 0x21, 0x94, 0x56, 0xec, 0x7d, 0xf4, 0xba, 0x7f, 0x66, 0xe0, 0xe2,
	0xf5, 0xc5, 0x36, 0x24, 0x97, 0xdb, 0x90, 0xfc, 0xda, 0x86, 0xe4, 0xcb, 0x2e, 0x0c, 0x2e, 0x77,
	0x61, 0xf0, 0x63, 0x17, 0x06, 0xef, 0x79, 0x9a, 0x55, 0x67, 0xb5, 0xe2, 0xb1, 0x5e, 0x09, 0x87,
	0xf7, 0x34, 0x97, 0x0a, 0x7d, 0x2c, 0x3e, 0x5d, 0x7f, 0x4e, 0xd5, 0xda, 0x00, 0xaa, 0x63, 0xfb,
	0x8f, 0x7f, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x57, 0x46, 0x69, 0xc3, 0x09, 0x04, 0x00, 0x00,
}

func (m *QueryRelationshipsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRelationshipsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRelationshipsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRelationships(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SubspaceId) > 0 {
		i -= len(m.SubspaceId)
		copy(dAtA[i:], m.SubspaceId)
		i = encodeVarintQueryRelationships(dAtA, i, uint64(len(m.SubspaceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryRelationships(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRelationshipsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRelationshipsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRelationshipsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRelationships(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Relationships) > 0 {
		for iNdEx := len(m.Relationships) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Relationships[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryRelationships(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryBlocksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlocksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlocksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRelationships(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SubspaceId) > 0 {
		i -= len(m.SubspaceId)
		copy(dAtA[i:], m.SubspaceId)
		i = encodeVarintQueryRelationships(dAtA, i, uint64(len(m.SubspaceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryRelationships(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBlocksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBlocksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBlocksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQueryRelationships(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Blocks) > 0 {
		for iNdEx := len(m.Blocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Blocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryRelationships(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryRelationships(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryRelationships(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRelationshipsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	l = len(m.SubspaceId)
	if l > 0 {
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	return n
}

func (m *QueryRelationshipsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Relationships) > 0 {
		for _, e := range m.Relationships {
			l = e.Size()
			n += 1 + l + sovQueryRelationships(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	return n
}

func (m *QueryBlocksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	l = len(m.SubspaceId)
	if l > 0 {
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	return n
}

func (m *QueryBlocksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Blocks) > 0 {
		for _, e := range m.Blocks {
			l = e.Size()
			n += 1 + l + sovQueryRelationships(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	return n
}

func sovQueryRelationships(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryRelationships(x uint64) (n int) {
	return sovQueryRelationships(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRelationshipsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRelationships
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
			return fmt.Errorf("proto: QueryRelationshipsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRelationshipsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubspaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRelationships(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRelationships
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
func (m *QueryRelationshipsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRelationships
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
			return fmt.Errorf("proto: QueryRelationshipsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRelationshipsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relationships", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relationships = append(m.Relationships, Relationship{})
			if err := m.Relationships[len(m.Relationships)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRelationships(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRelationships
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
func (m *QueryBlocksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRelationships
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
			return fmt.Errorf("proto: QueryBlocksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlocksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubspaceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRelationships(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRelationships
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
func (m *QueryBlocksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryRelationships
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
			return fmt.Errorf("proto: QueryBlocksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBlocksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocks = append(m.Blocks, UserBlock{})
			if err := m.Blocks[len(m.Blocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryRelationships
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
				return ErrInvalidLengthQueryRelationships
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryRelationships(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryRelationships
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
func skipQueryRelationships(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryRelationships
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
					return 0, ErrIntOverflowQueryRelationships
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
					return 0, ErrIntOverflowQueryRelationships
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
				return 0, ErrInvalidLengthQueryRelationships
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryRelationships
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryRelationships
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryRelationships        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryRelationships          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryRelationships = fmt.Errorf("proto: unexpected end of group")
)
