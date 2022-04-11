// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/supply/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
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

// QueryTotalSupplyRequest is the request type for Query/TotalSupply RPC method
type QueryTotalSupplyRequest struct {
	// coin denom to query the circulating supply for
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// divider_exponent is a factor used to power the divider used to convert the supply to the desired representation
	DividerExponent uint64 `protobuf:"varint,2,opt,name=divider_exponent,json=dividerExponent,proto3" json:"divider_exponent,omitempty"`
}

func (m *QueryTotalSupplyRequest) Reset()         { *m = QueryTotalSupplyRequest{} }
func (m *QueryTotalSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSupplyRequest) ProtoMessage()    {}
func (*QueryTotalSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{0}
}
func (m *QueryTotalSupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSupplyRequest.Merge(m, src)
}
func (m *QueryTotalSupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSupplyRequest proto.InternalMessageInfo

// QueryTotalSupplyResponse is the response type for the Query/TotalSupply RPC method
type QueryTotalSupplyResponse struct {
	TotalSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=total_supply,json=totalSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_supply"`
}

func (m *QueryTotalSupplyResponse) Reset()         { *m = QueryTotalSupplyResponse{} }
func (m *QueryTotalSupplyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSupplyResponse) ProtoMessage()    {}
func (*QueryTotalSupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{1}
}
func (m *QueryTotalSupplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSupplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSupplyResponse.Merge(m, src)
}
func (m *QueryTotalSupplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSupplyResponse proto.InternalMessageInfo

// QueryCirculatingSupplyRequest is the request type for the Query/CirculatingSupply RPC method
type QueryCirculatingSupplyRequest struct {
	// coin denom to query the circulating supply for
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// divider_exponent is a factor used to power the divider used to convert the supply to the desired representation
	DividerExponent uint64 `protobuf:"varint,2,opt,name=divider_exponent,json=dividerExponent,proto3" json:"divider_exponent,omitempty"`
}

func (m *QueryCirculatingSupplyRequest) Reset()         { *m = QueryCirculatingSupplyRequest{} }
func (m *QueryCirculatingSupplyRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCirculatingSupplyRequest) ProtoMessage()    {}
func (*QueryCirculatingSupplyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{2}
}
func (m *QueryCirculatingSupplyRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCirculatingSupplyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCirculatingSupplyRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCirculatingSupplyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCirculatingSupplyRequest.Merge(m, src)
}
func (m *QueryCirculatingSupplyRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCirculatingSupplyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCirculatingSupplyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCirculatingSupplyRequest proto.InternalMessageInfo

// QueryCirculatingSupplyRequest is the response type for the Query/CirculatingSupply RPC method
type QueryCirculatingSupplyResponse struct {
	CirculatingSupply github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=circulating_supply,json=circulatingSupply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"circulating_supply"`
}

func (m *QueryCirculatingSupplyResponse) Reset()         { *m = QueryCirculatingSupplyResponse{} }
func (m *QueryCirculatingSupplyResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCirculatingSupplyResponse) ProtoMessage()    {}
func (*QueryCirculatingSupplyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_821941a4adac8710, []int{3}
}
func (m *QueryCirculatingSupplyResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCirculatingSupplyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCirculatingSupplyResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCirculatingSupplyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCirculatingSupplyResponse.Merge(m, src)
}
func (m *QueryCirculatingSupplyResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCirculatingSupplyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCirculatingSupplyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCirculatingSupplyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryTotalSupplyRequest)(nil), "desmos.supply.v1.QueryTotalSupplyRequest")
	proto.RegisterType((*QueryTotalSupplyResponse)(nil), "desmos.supply.v1.QueryTotalSupplyResponse")
	proto.RegisterType((*QueryCirculatingSupplyRequest)(nil), "desmos.supply.v1.QueryCirculatingSupplyRequest")
	proto.RegisterType((*QueryCirculatingSupplyResponse)(nil), "desmos.supply.v1.QueryCirculatingSupplyResponse")
}

func init() { proto.RegisterFile("desmos/supply/v1/query.proto", fileDescriptor_821941a4adac8710) }

var fileDescriptor_821941a4adac8710 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xde, 0x09, 0x56, 0x74, 0x2a, 0xd8, 0x0e, 0x05, 0xe3, 0x52, 0x37, 0x25, 0x87, 0x92, 0x8a,
	0x99, 0x31, 0xd6, 0x93, 0xc7, 0x8a, 0x87, 0x1e, 0x1b, 0x3d, 0x09, 0x12, 0x66, 0x77, 0xc7, 0x75,
	0x74, 0x33, 0x33, 0xdd, 0x99, 0xdd, 0x36, 0x88, 0x17, 0x4f, 0x7a, 0x13, 0xf4, 0x07, 0xf4, 0x07,
	0xf8, 0x43, 0x7a, 0x11, 0x0a, 0x5e, 0xc4, 0x43, 0x91, 0xc4, 0x83, 0x3f, 0x43, 0x32, 0x33, 0x31,
	0x21, 0x21, 0x50, 0x90, 0x9e, 0x76, 0xdf, 0xfb, 0xe6, 0xbd, 0xef, 0xfb, 0xde, 0x9b, 0x81, 0x9b,
	0x29, 0xd3, 0x7d, 0xa9, 0x89, 0x2e, 0x95, 0xca, 0x07, 0xa4, 0xea, 0x90, 0xc3, 0x92, 0x15, 0x03,
	0xac, 0x0a, 0x69, 0x24, 0x5a, 0x73, 0x28, 0x76, 0x28, 0xae, 0x3a, 0xe1, 0x46, 0x26, 0x33, 0x69,
	0x41, 0x32, 0xfe, 0x73, 0xe7, 0xc2, 0xcd, 0x4c, 0xca, 0x2c, 0x67, 0x84, 0x2a, 0x4e, 0xa8, 0x10,
	0xd2, 0x50, 0xc3, 0xa5, 0xd0, 0x1e, 0xbd, 0xed, 0x51, 0x1b, 0xc5, 0xe5, 0x4b, 0x42, 0x85, 0x27,
	0x08, 0xa3, 0x79, 0xe8, 0xa8, 0xa0, 0x4a, 0xb1, 0xe2, 0x5f, 0x69, 0x22, 0xc7, 0x02, 0x7a, 0x8e,
	0xd1, 0x05, 0x93, 0x52, 0x17, 0x91, 0x98, 0x6a, 0x46, 0xaa, 0x4e, 0xcc, 0x0c, 0xed, 0x90, 0x44,
	0x72, 0xe1, 0xf0, 0x66, 0x0a, 0x6f, 0x1d, 0x8c, 0xad, 0x3c, 0x93, 0x86, 0xe6, 0x4f, 0xad, 0x81,
	0x2e, 0x3b, 0x2c, 0x99, 0x36, 0x68, 0x03, 0xae, 0xa4, 0x4c, 0xc8, 0x7e, 0x1d, 0x6c, 0x81, 0xd6,
	0xf5, 0xae, 0x0b, 0xd0, 0x0e, 0x5c, 0x4b, 0x79, 0xc5, 0x53, 0x56, 0xf4, 0xd8, 0xb1, 0x92, 0x82,
	0x09, 0x53, 0xaf, 0x6d, 0x81, 0xd6, 0x95, 0xee, 0x4d, 0x9f, 0x7f, 0xe2, 0xd3, 0x8f, 0xae, 0x7d,
	0x38, 0x69, 0x04, 0x7f, 0x4e, 0x1a, 0x41, 0xf3, 0x08, 0xd6, 0x17, 0x59, 0xb4, 0x92, 0x42, 0x33,
	0x74, 0x00, 0x6f, 0x98, 0x71, 0xba, 0xe7, 0xc6, 0xe7, 0xd8, 0xf6, 0xf0, 0xe9, 0x79, 0x23, 0xf8,
	0x79, 0xde, 0xd8, 0xce, 0xb8, 0x79, 0x55, 0xc6, 0x38, 0x91, 0x7d, 0x6f, 0xcc, 0x7f, 0xda, 0x3a,
	0x7d, 0x43, 0xcc, 0x40, 0x31, 0x8d, 0xf7, 0x85, 0xe9, 0xae, 0x9a, 0x69, 0xeb, 0x19, 0xe2, 0xd7,
	0xf0, 0x8e, 0x25, 0x7e, 0xcc, 0x8b, 0xa4, 0xcc, 0xa9, 0xe1, 0x22, 0xbb, 0x34, 0x93, 0x1f, 0x01,
	0x8c, 0x96, 0x91, 0x79, 0xaf, 0x2f, 0x20, 0x4a, 0xa6, 0xe0, 0xff, 0x39, 0x5e, 0x4f, 0xe6, 0x69,
	0xa6, 0x5a, 0x1e, 0x7c, 0xab, 0xc1, 0x15, 0xab, 0x05, 0x7d, 0x01, 0x70, 0x75, 0x66, 0xec, 0x68,
	0x07, 0xcf, 0xdf, 0x56, 0xbc, 0xe4, 0x02, 0x84, 0x77, 0x2f, 0x72, 0xd4, 0x39, 0x6b, 0xe2, 0xf7,
	0xdf, 0x7f, 0x7f, 0xae, 0xb5, 0xd0, 0x36, 0x59, 0x78, 0x2a, 0x76, 0x33, 0x6d, 0x1f, 0xbf, 0xb5,
	0x03, 0x7e, 0x87, 0xbe, 0x02, 0xb8, 0xbe, 0x30, 0x27, 0x44, 0x96, 0x30, 0x2e, 0x5b, 0x5f, 0x78,
	0xff, 0xe2, 0x05, 0x5e, 0xe8, 0x43, 0x2b, 0x14, 0xa3, 0x7b, 0x8b, 0x42, 0x67, 0x06, 0x3a, 0x27,
	0x77, 0x6f, 0xff, 0x74, 0x18, 0x81, 0xb3, 0x61, 0x04, 0x7e, 0x0d, 0x23, 0xf0, 0x69, 0x14, 0x05,
	0x67, 0xa3, 0x28, 0xf8, 0x31, 0x8a, 0x82, 0xe7, 0x64, 0x66, 0x5d, 0xae, 0x63, 0x3b, 0xa7, 0xb1,
	0x9e, 0x74, 0xaf, 0x76, 0xc9, 0xf1, 0x84, 0xc2, 0xee, 0x2e, 0xbe, 0x6a, 0x1f, 0xde, 0xee, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xc7, 0x88, 0x80, 0x54, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// TotalSupply queries the total supply of the given denom
	TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error)
	// CirculatingSupply queries the amount of tokens circulating in the market of the given denom
	CirculatingSupply(ctx context.Context, in *QueryCirculatingSupplyRequest, opts ...grpc.CallOption) (*QueryCirculatingSupplyResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TotalSupply(ctx context.Context, in *QueryTotalSupplyRequest, opts ...grpc.CallOption) (*QueryTotalSupplyResponse, error) {
	out := new(QueryTotalSupplyResponse)
	err := c.cc.Invoke(ctx, "/desmos.supply.v1.Query/TotalSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CirculatingSupply(ctx context.Context, in *QueryCirculatingSupplyRequest, opts ...grpc.CallOption) (*QueryCirculatingSupplyResponse, error) {
	out := new(QueryCirculatingSupplyResponse)
	err := c.cc.Invoke(ctx, "/desmos.supply.v1.Query/CirculatingSupply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// TotalSupply queries the total supply of the given denom
	TotalSupply(context.Context, *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error)
	// CirculatingSupply queries the amount of tokens circulating in the market of the given denom
	CirculatingSupply(context.Context, *QueryCirculatingSupplyRequest) (*QueryCirculatingSupplyResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) TotalSupply(ctx context.Context, req *QueryTotalSupplyRequest) (*QueryTotalSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSupply not implemented")
}
func (*UnimplementedQueryServer) CirculatingSupply(ctx context.Context, req *QueryCirculatingSupplyRequest) (*QueryCirculatingSupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CirculatingSupply not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_TotalSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.supply.v1.Query/TotalSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSupply(ctx, req.(*QueryTotalSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CirculatingSupply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCirculatingSupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CirculatingSupply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.supply.v1.Query/CirculatingSupply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CirculatingSupply(ctx, req.(*QueryCirculatingSupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.supply.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TotalSupply",
			Handler:    _Query_TotalSupply_Handler,
		},
		{
			MethodName: "CirculatingSupply",
			Handler:    _Query_CirculatingSupply_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/supply/v1/query.proto",
}

func (m *QueryTotalSupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DividerExponent != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.DividerExponent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryTotalSupplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSupplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSupplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalSupply.Size()
		i -= size
		if _, err := m.TotalSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryCirculatingSupplyRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCirculatingSupplyRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCirculatingSupplyRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DividerExponent != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.DividerExponent))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryCirculatingSupplyResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCirculatingSupplyResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCirculatingSupplyResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.CirculatingSupply.Size()
		i -= size
		if _, err := m.CirculatingSupply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryTotalSupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.DividerExponent != 0 {
		n += 1 + sovQuery(uint64(m.DividerExponent))
	}
	return n
}

func (m *QueryTotalSupplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TotalSupply.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryCirculatingSupplyRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.DividerExponent != 0 {
		n += 1 + sovQuery(uint64(m.DividerExponent))
	}
	return n
}

func (m *QueryCirculatingSupplyResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CirculatingSupply.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryTotalSupplyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTotalSupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DividerExponent", wireType)
			}
			m.DividerExponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DividerExponent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTotalSupplyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTotalSupplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSupplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCirculatingSupplyRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCirculatingSupplyRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCirculatingSupplyRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DividerExponent", wireType)
			}
			m.DividerExponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DividerExponent |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCirculatingSupplyResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCirculatingSupplyResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCirculatingSupplyResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CirculatingSupply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CirculatingSupply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
