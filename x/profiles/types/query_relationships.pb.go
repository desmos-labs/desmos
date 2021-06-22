// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/query_relationships.proto

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

// QueryUserRelationshipsRequest is the request type for the
// Query/UserRelationships RPC method.
type QueryUserRelationshipsRequest struct {
	// address of the user to query the relationships for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// the id of the subspace to query the relationships for
	SubspaceId string `protobuf:"bytes,2,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryUserRelationshipsRequest) Reset()         { *m = QueryUserRelationshipsRequest{} }
func (m *QueryUserRelationshipsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUserRelationshipsRequest) ProtoMessage()    {}
func (*QueryUserRelationshipsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0b8922e87a25523, []int{0}
}
func (m *QueryUserRelationshipsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserRelationshipsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserRelationshipsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserRelationshipsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserRelationshipsRequest.Merge(m, src)
}
func (m *QueryUserRelationshipsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserRelationshipsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserRelationshipsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserRelationshipsRequest proto.InternalMessageInfo

// QueryUserRelationshipsResponse is the response type for the
// Query/UserRelationships RPC method.
type QueryUserRelationshipsResponse struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// relationships represent the list of all the relationships for the queried
	// user
	Relationships []Relationship `protobuf:"bytes,2,rep,name=relationships,proto3" json:"relationships"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageResponse `protobuf:"bytes,3,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryUserRelationshipsResponse) Reset()         { *m = QueryUserRelationshipsResponse{} }
func (m *QueryUserRelationshipsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserRelationshipsResponse) ProtoMessage()    {}
func (*QueryUserRelationshipsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0b8922e87a25523, []int{1}
}
func (m *QueryUserRelationshipsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserRelationshipsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserRelationshipsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserRelationshipsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserRelationshipsResponse.Merge(m, src)
}
func (m *QueryUserRelationshipsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserRelationshipsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserRelationshipsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserRelationshipsResponse proto.InternalMessageInfo

func (m *QueryUserRelationshipsResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *QueryUserRelationshipsResponse) GetRelationships() []Relationship {
	if m != nil {
		return m.Relationships
	}
	return nil
}

func (m *QueryUserRelationshipsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryUserBlocksRequest is the request type for the Query/UserBlocks RPC
// endpoint
type QueryUserBlocksRequest struct {
	// address of the user to query the blocks for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *QueryUserBlocksRequest) Reset()         { *m = QueryUserBlocksRequest{} }
func (m *QueryUserBlocksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUserBlocksRequest) ProtoMessage()    {}
func (*QueryUserBlocksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0b8922e87a25523, []int{2}
}
func (m *QueryUserBlocksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserBlocksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserBlocksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserBlocksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserBlocksRequest.Merge(m, src)
}
func (m *QueryUserBlocksRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserBlocksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserBlocksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserBlocksRequest proto.InternalMessageInfo

// QueryUserBlocksResponse is the response type for the Query/UserBlocks RPC
// method.
type QueryUserBlocksResponse struct {
	// blocks represent the list of all the blocks for the queried user
	Blocks []UserBlock `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks"`
}

func (m *QueryUserBlocksResponse) Reset()         { *m = QueryUserBlocksResponse{} }
func (m *QueryUserBlocksResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserBlocksResponse) ProtoMessage()    {}
func (*QueryUserBlocksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0b8922e87a25523, []int{3}
}
func (m *QueryUserBlocksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserBlocksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserBlocksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserBlocksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserBlocksResponse.Merge(m, src)
}
func (m *QueryUserBlocksResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserBlocksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserBlocksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserBlocksResponse proto.InternalMessageInfo

func (m *QueryUserBlocksResponse) GetBlocks() []UserBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryUserRelationshipsRequest)(nil), "desmos.profiles.v1beta1.QueryUserRelationshipsRequest")
	proto.RegisterType((*QueryUserRelationshipsResponse)(nil), "desmos.profiles.v1beta1.QueryUserRelationshipsResponse")
	proto.RegisterType((*QueryUserBlocksRequest)(nil), "desmos.profiles.v1beta1.QueryUserBlocksRequest")
	proto.RegisterType((*QueryUserBlocksResponse)(nil), "desmos.profiles.v1beta1.QueryUserBlocksResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/query_relationships.proto", fileDescriptor_c0b8922e87a25523)
}

var fileDescriptor_c0b8922e87a25523 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0x8f, 0xd3, 0x40,
	0x10, 0xf5, 0xde, 0x9d, 0x4e, 0xb0, 0x11, 0x8d, 0x85, 0x38, 0x5f, 0x04, 0x4e, 0x14, 0x09, 0x88,
	0x90, 0xd8, 0x55, 0x82, 0x44, 0x41, 0x85, 0x52, 0xf0, 0xd1, 0x71, 0x96, 0x68, 0xa0, 0x88, 0x76,
	0xe3, 0x39, 0x9f, 0x85, 0xe3, 0xdd, 0xf3, 0xd8, 0x88, 0xfc, 0x03, 0x4a, 0x7e, 0xc2, 0x55, 0xfc,
	0x96, 0x2b, 0x53, 0x52, 0x21, 0x94, 0x34, 0xfc, 0x0c, 0xe4, 0xdd, 0x4d, 0xce, 0x11, 0x31, 0x74,
	0x33, 0x9e, 0x79, 0x6f, 0xde, 0x7b, 0x5e, 0x3a, 0x8a, 0x01, 0xe7, 0x0a, 0xb9, 0x2e, 0xd4, 0x79,
	0x9a, 0x01, 0xf2, 0xcf, 0x23, 0x09, 0xa5, 0x18, 0xf1, 0xcb, 0x0a, 0x8a, 0xc5, 0xb4, 0x80, 0x4c,
	0x94, 0xa9, 0xca, 0xf1, 0x22, 0xd5, 0xc8, 0x74, 0xa1, 0x4a, 0xe5, 0x9f, 0x58, 0x08, 0xdb, 0x40,
	0x98, 0x83, 0x74, 0xef, 0x26, 0x2a, 0x51, 0x66, 0x87, 0xd7, 0x95, 0x5d, 0xef, 0xde, 0x4f, 0x94,
	0x4a, 0x32, 0xe0, 0x42, 0xa7, 0x5c, 0xe4, 0xb9, 0x2a, 0x2d, 0xa1, 0x9b, 0x9e, 0xba, 0xa9, 0xe9,
	0x64, 0x75, 0xce, 0x45, 0xbe, 0x70, 0xa3, 0x71, 0x9b, 0xb4, 0xb9, 0x8a, 0x21, 0xc3, 0x7d, 0xda,
	0xba, 0xa7, 0x33, 0x55, 0x63, 0xa6, 0x56, 0x85, 0x6d, 0xdc, 0xe8, 0x89, 0xed, 0xb8, 0x14, 0x08,
	0xd6, 0xdd, 0x96, 0x50, 0x8b, 0x24, 0xcd, 0x0d, 0x97, 0xdd, 0x1d, 0x7c, 0x27, 0xf4, 0xc1, 0x59,
	0xbd, 0xf2, 0x1e, 0xa1, 0x88, 0x9a, 0x77, 0x22, 0xb8, 0xac, 0x00, 0x4b, 0xdf, 0xa7, 0x47, 0x15,
	0x42, 0x11, 0x90, 0x3e, 0x19, 0xde, 0x8e, 0x4c, 0xed, 0xf7, 0x68, 0x07, 0x2b, 0x89, 0x5a, 0xcc,
	0x60, 0x9a, 0xc6, 0xc1, 0x81, 0x19, 0xd1, 0xcd, 0xa7, 0xb7, 0xb1, 0xff, 0x8a, 0xd2, 0x9b, 0x53,
	0xc1, 0x61, 0x9f, 0x0c, 0x3b, 0xe3, 0x47, 0xcc, 0xa9, 0xac, 0x75, 0x31, 0xa3, 0x6b, 0x13, 0x28,
	0x7b, 0x27, 0x12, 0x70, 0x07, 0xa3, 0x06, 0xf2, 0xc5, 0xad, 0xaf, 0x57, 0x3d, 0xef, 0xf7, 0x55,
	0xcf, 0x1b, 0x2c, 0x09, 0x0d, 0xdb, 0x84, 0xa2, 0x56, 0x39, 0xc2, 0x5e, 0xa5, 0x67, 0xf4, 0xce,
	0x4e, 0x7a, 0xc1, 0x41, 0xff, 0x70, 0xd8, 0x19, 0x3f, 0x64, 0x2d, 0xbf, 0x96, 0x35, 0xa9, 0x27,
	0x47, 0xd7, 0x3f, 0x7b, 0x5e, 0xb4, 0xcb, 0xe0, 0xbf, 0xde, 0xe3, 0xed, 0xf1, 0x7f, 0xbd, 0x59,
	0x8d, 0x4d, 0x73, 0x83, 0xe7, 0xf4, 0xde, 0xd6, 0xd1, 0x24, 0x53, 0xb3, 0x4f, 0xff, 0xca, 0xbc,
	0x11, 0xc5, 0x47, 0x7a, 0xf2, 0x17, 0xce, 0x45, 0xf0, 0x92, 0x1e, 0x4b, 0xf3, 0x25, 0x20, 0xc6,
	0xe7, 0xa0, 0xd5, 0xe7, 0x16, 0xec, 0x4c, 0x3a, 0xdc, 0xe4, 0xcd, 0xf5, 0x2a, 0x24, 0xcb, 0x55,
	0x48, 0x7e, 0xad, 0x42, 0xf2, 0x6d, 0x1d, 0x7a, 0xcb, 0x75, 0xe8, 0xfd, 0x58, 0x87, 0xde, 0x07,
	0x96, 0xa4, 0xe5, 0x45, 0x25, 0xd9, 0x4c, 0xcd, 0xb9, 0x65, 0x7d, 0x9a, 0x09, 0x89, 0xae, 0xe6,
	0x5f, 0x6e, 0x9e, 0x6f, 0xb9, 0xd0, 0x80, 0xf2, 0xd8, 0xbc, 0xb0, 0x67, 0x7f, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xbf, 0xd3, 0xa8, 0xa7, 0x79, 0x03, 0x00, 0x00,
}

func (m *QueryUserRelationshipsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserRelationshipsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserRelationshipsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryUserRelationshipsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserRelationshipsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserRelationshipsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			dAtA[i] = 0x12
		}
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

func (m *QueryUserBlocksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserBlocksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserBlocksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryRelationships(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUserBlocksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserBlocksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserBlocksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *QueryUserRelationshipsRequest) Size() (n int) {
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

func (m *QueryUserRelationshipsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
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

func (m *QueryUserBlocksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryRelationships(uint64(l))
	}
	return n
}

func (m *QueryUserBlocksResponse) Size() (n int) {
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
	return n
}

func sovQueryRelationships(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryRelationships(x uint64) (n int) {
	return sovQueryRelationships(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryUserRelationshipsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUserRelationshipsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserRelationshipsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryUserRelationshipsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUserRelationshipsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserRelationshipsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryUserBlocksRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUserBlocksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserBlocksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryUserBlocksResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUserBlocksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserBlocksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
