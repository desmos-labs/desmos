// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/query_app_links.proto

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

// QueryUserApplicationLinkRequest represents the request used when querying an
// application link using an application name and username for a given user
type QueryUserApplicationLinkRequest struct {
	// User contains the Desmos profile address associated for which the link
	// should be searched for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Application represents the application name associated with the link
	Application string `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
	// Username represents the username inside the application associated with the
	// link
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
}

func (m *QueryUserApplicationLinkRequest) Reset()         { *m = QueryUserApplicationLinkRequest{} }
func (m *QueryUserApplicationLinkRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUserApplicationLinkRequest) ProtoMessage()    {}
func (*QueryUserApplicationLinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d2a2b45fd238cb, []int{0}
}
func (m *QueryUserApplicationLinkRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserApplicationLinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserApplicationLinkRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserApplicationLinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserApplicationLinkRequest.Merge(m, src)
}
func (m *QueryUserApplicationLinkRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserApplicationLinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserApplicationLinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserApplicationLinkRequest proto.InternalMessageInfo

func (m *QueryUserApplicationLinkRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *QueryUserApplicationLinkRequest) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *QueryUserApplicationLinkRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

// QueryUserApplicationLinkResponse represents the response to the query
// allowing to get an application link for a specific user, searching via the
// application name and username
type QueryUserApplicationLinkResponse struct {
	Link ApplicationLink `protobuf:"bytes,1,opt,name=link,proto3" json:"link"`
}

func (m *QueryUserApplicationLinkResponse) Reset()         { *m = QueryUserApplicationLinkResponse{} }
func (m *QueryUserApplicationLinkResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserApplicationLinkResponse) ProtoMessage()    {}
func (*QueryUserApplicationLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d2a2b45fd238cb, []int{1}
}
func (m *QueryUserApplicationLinkResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserApplicationLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserApplicationLinkResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserApplicationLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserApplicationLinkResponse.Merge(m, src)
}
func (m *QueryUserApplicationLinkResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserApplicationLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserApplicationLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserApplicationLinkResponse proto.InternalMessageInfo

func (m *QueryUserApplicationLinkResponse) GetLink() ApplicationLink {
	if m != nil {
		return m.Link
	}
	return ApplicationLink{}
}

// QueryUserApplicationLinksRequest represents the request used when querying
// the application links of a specific user
type QueryUserApplicationLinksRequest struct {
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	// Pagination defines an optional pagination for the request
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryUserApplicationLinksRequest) Reset()         { *m = QueryUserApplicationLinksRequest{} }
func (m *QueryUserApplicationLinksRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUserApplicationLinksRequest) ProtoMessage()    {}
func (*QueryUserApplicationLinksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d2a2b45fd238cb, []int{2}
}
func (m *QueryUserApplicationLinksRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserApplicationLinksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserApplicationLinksRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserApplicationLinksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserApplicationLinksRequest.Merge(m, src)
}
func (m *QueryUserApplicationLinksRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserApplicationLinksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserApplicationLinksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserApplicationLinksRequest proto.InternalMessageInfo

func (m *QueryUserApplicationLinksRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *QueryUserApplicationLinksRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryUserApplicationLinksResponse represents the response to the query used
// to get the application links for a specific user
type QueryUserApplicationLinksResponse struct {
	Links []ApplicationLink `protobuf:"bytes,1,rep,name=links,proto3" json:"links"`
	// Pagination defines the pagination response
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryUserApplicationLinksResponse) Reset()         { *m = QueryUserApplicationLinksResponse{} }
func (m *QueryUserApplicationLinksResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUserApplicationLinksResponse) ProtoMessage()    {}
func (*QueryUserApplicationLinksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d2a2b45fd238cb, []int{3}
}
func (m *QueryUserApplicationLinksResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUserApplicationLinksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUserApplicationLinksResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUserApplicationLinksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUserApplicationLinksResponse.Merge(m, src)
}
func (m *QueryUserApplicationLinksResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUserApplicationLinksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUserApplicationLinksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUserApplicationLinksResponse proto.InternalMessageInfo

func (m *QueryUserApplicationLinksResponse) GetLinks() []ApplicationLink {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *QueryUserApplicationLinksResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryUserApplicationLinkRequest)(nil), "desmos.profiles.v1beta1.QueryUserApplicationLinkRequest")
	proto.RegisterType((*QueryUserApplicationLinkResponse)(nil), "desmos.profiles.v1beta1.QueryUserApplicationLinkResponse")
	proto.RegisterType((*QueryUserApplicationLinksRequest)(nil), "desmos.profiles.v1beta1.QueryUserApplicationLinksRequest")
	proto.RegisterType((*QueryUserApplicationLinksResponse)(nil), "desmos.profiles.v1beta1.QueryUserApplicationLinksResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/query_app_links.proto", fileDescriptor_18d2a2b45fd238cb)
}

var fileDescriptor_18d2a2b45fd238cb = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6e, 0xda, 0x30,
	0x18, 0xc7, 0x93, 0xc1, 0xa6, 0xcd, 0xdc, 0xa2, 0x49, 0x0b, 0xd1, 0x14, 0xb2, 0x1c, 0x36, 0x34,
	0x09, 0x5b, 0xb0, 0x27, 0x18, 0x9a, 0xb6, 0x1d, 0x7a, 0x68, 0x23, 0xf5, 0xd2, 0x0b, 0x72, 0xc0,
	0xa4, 0x16, 0x89, 0x6d, 0xe2, 0xa4, 0x2a, 0x97, 0x3e, 0x43, 0x9f, 0xa5, 0x4f, 0xc1, 0x91, 0x63,
	0x4f, 0x55, 0x05, 0x2f, 0x52, 0xc5, 0x76, 0x01, 0xa1, 0x86, 0x4a, 0xbd, 0xf9, 0xcb, 0xff, 0xff,
	0xf9, 0xfb, 0xfd, 0xbf, 0x18, 0xf4, 0x26, 0x44, 0x66, 0x5c, 0x22, 0x91, 0xf3, 0x29, 0x4d, 0x89,
	0x44, 0x57, 0xfd, 0x98, 0x14, 0xb8, 0x8f, 0xe6, 0x25, 0xc9, 0x17, 0x23, 0x2c, 0xc4, 0x28, 0xa5,
	0x6c, 0x26, 0xa1, 0xc8, 0x79, 0xc1, 0x9d, 0x2f, 0xda, 0x0e, 0x9f, 0xed, 0xd0, 0xd8, 0xbd, 0xcf,
	0x09, 0x4f, 0xb8, 0xf2, 0xa0, 0xea, 0xa4, 0xed, 0xde, 0xd7, 0x84, 0xf3, 0x24, 0x25, 0x08, 0x0b,
	0x8a, 0x30, 0x63, 0xbc, 0xc0, 0x05, 0xe5, 0xcc, 0x5c, 0xe6, 0xb5, 0x8d, 0xaa, 0xaa, 0xb8, 0x9c,
	0x22, 0xcc, 0x16, 0x46, 0x82, 0x75, 0x58, 0x19, 0x9f, 0x90, 0x54, 0x1e, 0x72, 0x79, 0xed, 0x31,
	0xaf, 0xfc, 0x23, 0x4d, 0xa0, 0x0b, 0x23, 0xfd, 0xd4, 0x15, 0x8a, 0xb1, 0x24, 0x3a, 0xd5, 0xf6,
	0x32, 0x81, 0x13, 0xca, 0x14, 0x92, 0xf6, 0x86, 0x12, 0x74, 0xce, 0x2a, 0xc7, 0xb9, 0x24, 0xf9,
	0x6f, 0x21, 0x52, 0x3a, 0x56, 0xea, 0x09, 0x65, 0xb3, 0x88, 0xcc, 0x4b, 0x22, 0x0b, 0xc7, 0x01,
	0xcd, 0x52, 0x92, 0xdc, 0xb5, 0x03, 0xbb, 0xfb, 0x29, 0x52, 0x67, 0x27, 0x00, 0x2d, 0xbc, 0x73,
	0xbb, 0xef, 0x94, 0xb4, 0xff, 0xc9, 0xf1, 0xc0, 0xc7, 0xca, 0xc9, 0x70, 0x46, 0xdc, 0x86, 0x92,
	0xb7, 0x75, 0x38, 0x05, 0x41, 0xfd, 0x50, 0x29, 0x38, 0x93, 0xc4, 0x19, 0x82, 0x66, 0x15, 0x57,
	0x4d, 0x6d, 0x0d, 0xba, 0xb0, 0xe6, 0x37, 0xc0, 0x83, 0xfe, 0x61, 0x73, 0xf9, 0xd0, 0xb1, 0x22,
	0xd5, 0x1b, 0xde, 0xd4, 0xcf, 0x91, 0xc7, 0xd2, 0xfd, 0x05, 0x60, 0xb7, 0x28, 0x15, 0xae, 0x35,
	0xf8, 0x0e, 0xcd, 0x8e, 0xab, 0xad, 0x42, 0xb5, 0xd5, 0x2d, 0xc3, 0x29, 0x4e, 0x88, 0xb9, 0x2f,
	0xda, 0xeb, 0x0c, 0xef, 0x6c, 0xf0, 0xed, 0x08, 0x80, 0x49, 0xfa, 0x07, 0xbc, 0x57, 0x3f, 0xd6,
	0xb5, 0x83, 0xc6, 0x1b, 0xa2, 0xea, 0x66, 0xe7, 0xdf, 0x0b, 0xcc, 0x3f, 0x5e, 0x65, 0xd6, 0x08,
	0xfb, 0xd0, 0xc3, 0xff, 0xcb, 0xb5, 0x6f, 0xaf, 0xd6, 0xbe, 0xfd, 0xb8, 0xf6, 0xed, 0xdb, 0x8d,
	0x6f, 0xad, 0x36, 0xbe, 0x75, 0xbf, 0xf1, 0xad, 0x0b, 0x98, 0xd0, 0xe2, 0xb2, 0x8c, 0xe1, 0x98,
	0x67, 0x48, 0x33, 0xf6, 0x52, 0x1c, 0x4b, 0x73, 0x46, 0xd7, 0xbb, 0xb7, 0x5b, 0x2c, 0x04, 0x91,
	0xf1, 0x07, 0xf5, 0xc4, 0x7e, 0x3d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x04, 0xe8, 0xcc, 0x72,
	0x03, 0x00, 0x00,
}

func (m *QueryUserApplicationLinkRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserApplicationLinkRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserApplicationLinkRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Application) > 0 {
		i -= len(m.Application)
		copy(dAtA[i:], m.Application)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.Application)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUserApplicationLinkResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserApplicationLinkResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserApplicationLinkResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryUserApplicationLinksRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserApplicationLinksRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserApplicationLinksRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQueryAppLinks(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUserApplicationLinksResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUserApplicationLinksResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUserApplicationLinksResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
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
				i = encodeVarintQueryAppLinks(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQueryAppLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovQueryAppLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryUserApplicationLinkRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	l = len(m.Application)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	return n
}

func (m *QueryUserApplicationLinkResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Link.Size()
	n += 1 + l + sovQueryAppLinks(uint64(l))
	return n
}

func (m *QueryUserApplicationLinksRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	return n
}

func (m *QueryUserApplicationLinksResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Links) > 0 {
		for _, e := range m.Links {
			l = e.Size()
			n += 1 + l + sovQueryAppLinks(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQueryAppLinks(uint64(l))
	}
	return n
}

func sovQueryAppLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQueryAppLinks(x uint64) (n int) {
	return sovQueryAppLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryUserApplicationLinkRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryUserApplicationLinkRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserApplicationLinkRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Application", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Application = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func (m *QueryUserApplicationLinkResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryUserApplicationLinkResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserApplicationLinkResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Link", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
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
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func (m *QueryUserApplicationLinksRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryUserApplicationLinksRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserApplicationLinksRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
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
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
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
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func (m *QueryUserApplicationLinksResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQueryAppLinks
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
			return fmt.Errorf("proto: QueryUserApplicationLinksResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUserApplicationLinksResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Links", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Links = append(m.Links, ApplicationLink{})
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
					return ErrIntOverflowQueryAppLinks
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
				return ErrInvalidLengthQueryAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQueryAppLinks
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
			skippy, err := skipQueryAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQueryAppLinks
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
func skipQueryAppLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQueryAppLinks
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
					return 0, ErrIntOverflowQueryAppLinks
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
					return 0, ErrIntOverflowQueryAppLinks
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
				return 0, ErrInvalidLengthQueryAppLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQueryAppLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQueryAppLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQueryAppLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQueryAppLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQueryAppLinks = fmt.Errorf("proto: unexpected end of group")
)
