// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/query_chain_links.proto

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

// QueryUserChainLinkRequest represents the request that should be used in order
// to retrieve the link associated with the provided user, for the given chain
// and having the given target address
type QueryUserChainLinkRequest struct {
	// User represents the Desmos address of the user to which search the link for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// ChainName contains the name of the chain to which search the link for
	ChainName string `protobuf:"bytes,2,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	// Target must contain the external address to which query the link for
	Target string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
}

func (m *QueryUserChainLinkRequest) Reset()         { *m = QueryUserChainLinkRequest{} }
func (m *QueryUserChainLinkRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUserChainLinkRequest) ProtoMessage()    {}
func (*QueryUserChainLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44c4be38a628772, []int{0}
}
func (m *QueryUserChainLinkRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserChainLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserChainLinkRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserChainLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserChainLinkRequest.Merge(m, src)
}
func (m *QueryUserChainLinkRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserChainLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserChainLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserChainLinkRequest proto.InternalMessageInfo

func (m *QueryUserChainLinkRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *QueryUserChainLinkRequest) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *QueryUserChainLinkRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

// QueryUserChainLinkResponse contains the data that is returned when querying a
// specific chain link
type QueryUserChainLinkResponse struct {
	Link ChainLink `protobuf:"bytes,1,opt,name=link,proto3" json:"link"`
}

func (m *QueryUserChainLinkResponse) Reset()         { *m = QueryUserChainLinkResponse{} }
func (m *QueryUserChainLinkResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserChainLinkResponse) ProtoMessage()    {}
func (*QueryUserChainLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44c4be38a628772, []int{1}
}
func (m *QueryUserChainLinkResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserChainLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserChainLinkResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserChainLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserChainLinkResponse.Merge(m, src)
}
func (m *QueryUserChainLinkResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserChainLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserChainLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserChainLinkResponse proto.InternalMessageInfo

func (m *QueryUserChainLinkResponse) GetLink() ChainLink {
	if m != nil {
		return m.Link
	}
	return ChainLink{}
}

// QueryChainLinksRequest is the request type for the
// Query/ChainLinks RPC endpoint
type QueryChainLinksRequest struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Pagination defines an optional pagination for the request
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryChainLinksRequest) Reset()         { *m = QueryChainLinksRequest{} }
func (m *QueryChainLinksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryChainLinksRequest) ProtoMessage()    {}
func (*QueryChainLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44c4be38a628772, []int{2}
}
func (m *QueryChainLinksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChainLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChainLinksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChainLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChainLinksRequest.Merge(m, src)
}
func (m *QueryChainLinksRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryChainLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChainLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChainLinksRequest proto.InternalMessageInfo

func (m *QueryChainLinksRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *QueryChainLinksRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryChainLinksResponse is the response type for the
// Query/ChainLinks RPC method.
type QueryChainLinksResponse struct {
	Links []ChainLink `protobuf:"bytes,1,rep,name=links,proto3" json:"links"`
	// Pagination defines the pagination response
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryChainLinksResponse) Reset()         { *m = QueryChainLinksResponse{} }
func (m *QueryChainLinksResponse) String() string { return proto.CompactTextString(m) }
func (*QueryChainLinksResponse) ProtoMessage()    {}
func (*QueryChainLinksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c44c4be38a628772, []int{3}
}
func (m *QueryChainLinksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryChainLinksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryChainLinksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryChainLinksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryChainLinksResponse.Merge(m, src)
}
func (m *QueryChainLinksResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryChainLinksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryChainLinksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryChainLinksResponse proto.InternalMessageInfo

func (m *QueryChainLinksResponse) GetLinks() []ChainLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *QueryChainLinksResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryUserChainLinkRequest)(nil), "desmos.profiles.v1beta1.QueryUserChainLinkRequest")
	proto.RegisterType((*QueryUserChainLinkResponse)(nil), "desmos.profiles.v1beta1.QueryUserChainLinkResponse")
	proto.RegisterType((*QueryChainLinksRequest)(nil), "desmos.profiles.v1beta1.QueryChainLinksRequest")
	proto.RegisterType((*QueryChainLinksResponse)(nil), "desmos.profiles.v1beta1.QueryChainLinksResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/query_chain_links.proto", fileDescriptor_c44c4be38a628772)
}

var fileDescriptor_c44c4be38a628772 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x4d, 0xd8, 0xb2, 0xd2, 0x7a, 0x6f, 0x11, 0xda, 0x4d, 0x23, 0x08, 0xab, 0x1c, 0x00, 0x21,
	0x61, 0xd3, 0x72, 0x45, 0x1c, 0x16, 0x09, 0x0e, 0x20, 0x04, 0x91, 0xb8, 0xec, 0xa5, 0x72, 0xda,
	0xa9, 0x6b, 0x35, 0xb1, 0xd3, 0xd8, 0xa9, 0xe8, 0x5f, 0xf0, 0x0d, 0x7c, 0x4d, 0x8f, 0x3d, 0x72,
	0x42, 0xa8, 0xfd, 0x11, 0x14, 0xdb, 0x4d, 0x2b, 0xb5, 0x05, 0x71, 0xf3, 0xe4, 0xcd, 0x9b, 0xf7,
	0xde, 0xc4, 0x46, 0x64, 0x04, 0xaa, 0x90, 0x8a, 0x94, 0x95, 0x1c, 0xf3, 0x1c, 0x14, 0x99, 0xf7,
	0x32, 0xd0, 0xb4, 0x47, 0x66, 0x35, 0x54, 0x8b, 0xc1, 0x70, 0x42, 0xb9, 0x18, 0xe4, 0x5c, 0x4c,
	0x15, 0x2e, 0x2b, 0xa9, 0x65, 0x70, 0x6d, 0x09, 0x78, 0x4b, 0xc0, 0x8e, 0x10, 0x3d, 0x60, 0x92,
	0x49, 0xd3, 0x43, 0x9a, 0x93, 0x6d, 0x8f, 0x1e, 0x32, 0x29, 0x59, 0x0e, 0x84, 0x96, 0x9c, 0x50,
	0x21, 0xa4, 0xa6, 0x9a, 0x4b, 0xe1, 0x86, 0x45, 0x5d, 0x87, 0x9a, 0x2a, 0xab, 0xc7, 0x84, 0x8a,
	0x85, 0x83, 0x5e, 0x9e, 0x32, 0x56, 0xc8, 0x11, 0xe4, 0xea, 0xd0, 0x59, 0xd4, 0x1d, 0xca, 0x86,
	0x31, 0xb0, 0x1e, 0x6c, 0xe1, 0xa0, 0xe7, 0xb6, 0x22, 0x19, 0x55, 0x60, 0x93, 0xb5, 0xe3, 0x4a,
	0xca, 0xb8, 0x30, 0xa6, 0x6c, 0x6f, 0x32, 0x46, 0xdd, 0x2f, 0x4d, 0xc7, 0x57, 0x05, 0xd5, 0xdb,
	0x46, 0xe4, 0x23, 0x17, 0xd3, 0x14, 0x66, 0x35, 0x28, 0x1d, 0x04, 0xa8, 0x53, 0x2b, 0xa8, 0x42,
	0xff, 0xc6, 0x7f, 0x76, 0x91, 0x9a, 0x73, 0xf0, 0x08, 0x21, 0x6b, 0x46, 0xd0, 0x02, 0xc2, 0x7b,
	0x06, 0xb9, 0x30, 0x5f, 0x3e, 0xd1, 0x02, 0x82, 0x2b, 0x74, 0xae, 0x69, 0xc5, 0x40, 0x87, 0x67,
	0x06, 0x72, 0x55, 0x72, 0x87, 0xa2, 0x63, 0x3a, 0xaa, 0x94, 0x42, 0x41, 0xf0, 0x1a, 0x75, 0x9a,
	0x6c, 0x46, 0xe8, 0xb2, 0x9f, 0xe0, 0x13, 0x5b, 0xc7, 0x2d, 0xf3, 0xb6, 0xb3, 0xfc, 0xf5, 0xd8,
	0x4b, 0x0d, 0x2b, 0xd1, 0xe8, 0xca, 0xcc, 0x6e, 0x51, 0xf5, 0xb7, 0x00, 0xef, 0x10, 0xda, 0x6d,
	0xc1, 0x04, 0xb8, 0xec, 0x3f, 0xc1, 0x6e, 0x81, 0xcd, 0xca, 0xb0, 0x59, 0x59, 0xab, 0xf9, 0x99,
	0x32, 0x70, 0xf3, 0xd2, 0x3d, 0x66, 0xf2, 0xc3, 0x47, 0xd7, 0x07, 0xb2, 0x2e, 0xcf, 0x1b, 0x74,
	0xdf, 0xfc, 0xab, 0xd0, 0xbf, 0x39, 0xfb, 0xaf, 0x40, 0x96, 0x16, 0xbc, 0x3f, 0xe2, 0xf1, 0xe9,
	0x3f, 0x3d, 0x5a, 0xf1, 0x7d, 0x93, 0xb7, 0x1f, 0x96, 0xeb, 0xd8, 0x5f, 0xad, 0x63, 0xff, 0xf7,
	0x3a, 0xf6, 0xbf, 0x6f, 0x62, 0x6f, 0xb5, 0x89, 0xbd, 0x9f, 0x9b, 0xd8, 0xbb, 0xeb, 0x31, 0xae,
	0x27, 0x75, 0x86, 0x87, 0xb2, 0x70, 0xaf, 0xe2, 0x45, 0x4e, 0x33, 0xb5, 0x7d, 0x21, 0xf3, 0x3e,
	0xf9, 0xb6, 0xbb, 0x8d, 0x7a, 0x51, 0x82, 0xca, 0xce, 0xcd, 0x95, 0x79, 0xf5, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x2b, 0xd8, 0xd3, 0x0f, 0x46, 0x03, 0x00, 0x00,
}

func (m *QueryUserChainLinkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserChainLinkRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserChainLinkRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintQueryChainLinks(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintQueryChainLinks(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryChainLinks(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUserChainLinkResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserChainLinkResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserChainLinkResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Link.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQueryChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryChainLinksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChainLinksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChainLinksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQueryChainLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryChainLinks(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryChainLinksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryChainLinksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryChainLinksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQueryChainLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Links) > 0 {
		for iNdEx := len(m.Links) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Links[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQueryChainLinks(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryChainLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryChainLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryUserChainLinkRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryChainLinks(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovQueryChainLinks(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovQueryChainLinks(uint64(l))
	}
	return n
}

func (m *QueryUserChainLinkResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Link.Size()
	n += 1 + l + sovQueryChainLinks(uint64(l))
	return n
}

func (m *QueryChainLinksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryChainLinks(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryChainLinks(uint64(l))
	}
	return n
}

func (m *QueryChainLinksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovQueryChainLinks(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryChainLinks(uint64(l))
	}
	return n
}

func sovQueryChainLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryChainLinks(x uint64) (n int) {
	return sovQueryChainLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryUserChainLinkRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryChainLinks
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
			return fmt.Errorf("proto: QueryUserChainLinkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserChainLinkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryChainLinks
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
				return ErrInvalidLengthQueryChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryChainLinks
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
				return ErrInvalidLengthQueryChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryChainLinks
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
				return ErrInvalidLengthQueryChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Target = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryChainLinks
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
func (m *QueryUserChainLinkResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryChainLinks
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
			return fmt.Errorf("proto: QueryUserChainLinkResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserChainLinkResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryChainLinks
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
				return ErrInvalidLengthQueryChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Link.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryChainLinks
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
func (m *QueryChainLinksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryChainLinks
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
			return fmt.Errorf("proto: QueryChainLinksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChainLinksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryChainLinks
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
				return ErrInvalidLengthQueryChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryChainLinks
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
				return ErrInvalidLengthQueryChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryChainLinks
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
			skippy, err := skipQueryChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryChainLinks
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
func (m *QueryChainLinksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryChainLinks
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
			return fmt.Errorf("proto: QueryChainLinksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryChainLinksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryChainLinks
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
				return ErrInvalidLengthQueryChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, ChainLink{})
			if err := m.Links[len(m.Links)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowQueryChainLinks
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
				return ErrInvalidLengthQueryChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryChainLinks
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
			skippy, err := skipQueryChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryChainLinks
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
func skipQueryChainLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryChainLinks
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
					return 0, ErrIntOverflowQueryChainLinks
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
					return 0, ErrIntOverflowQueryChainLinks
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
				return 0, ErrInvalidLengthQueryChainLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryChainLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryChainLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryChainLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryChainLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryChainLinks = fmt.Errorf("proto: unexpected end of group")
)
