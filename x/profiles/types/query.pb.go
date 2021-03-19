// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryProfileRequest is the request type for the Query/Profile RPC method.
type QueryProfileRequest struct {
	// Address or DTag of the user to query the profile for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *QueryProfileRequest) Reset()         { *m = QueryProfileRequest{} }
func (m *QueryProfileRequest) String() string { return proto.CompactTextString(m) }
func (*QueryProfileRequest) ProtoMessage()    {}
func (*QueryProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0074f57a59f38d, []int{0}
}
func (m *QueryProfileRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProfileRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfileRequest.Merge(m, src)
}
func (m *QueryProfileRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfileRequest proto.InternalMessageInfo

// QueryProfileResponse is the response type for the Query/Profile RPC method.
type QueryProfileResponse struct {
	Profile *types.Any `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (m *QueryProfileResponse) Reset()         { *m = QueryProfileResponse{} }
func (m *QueryProfileResponse) String() string { return proto.CompactTextString(m) }
func (*QueryProfileResponse) ProtoMessage()    {}
func (*QueryProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0074f57a59f38d, []int{1}
}
func (m *QueryProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryProfileResponse.Merge(m, src)
}
func (m *QueryProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryProfileResponse proto.InternalMessageInfo

func (m *QueryProfileResponse) GetProfile() *types.Any {
	if m != nil {
		return m.Profile
	}
	return nil
}

// QueryDTagTransfersRequest is the request type for the Query/DTagTransfers RPC
// endpoint
type QueryDTagTransfersRequest struct {
	// Address or DTag of the user to query the transfer requests for
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *QueryDTagTransfersRequest) Reset()         { *m = QueryDTagTransfersRequest{} }
func (m *QueryDTagTransfersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDTagTransfersRequest) ProtoMessage()    {}
func (*QueryDTagTransfersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0074f57a59f38d, []int{2}
}
func (m *QueryDTagTransfersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDTagTransfersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDTagTransfersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDTagTransfersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDTagTransfersRequest.Merge(m, src)
}
func (m *QueryDTagTransfersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDTagTransfersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDTagTransfersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDTagTransfersRequest proto.InternalMessageInfo

// QueryDTagTransfersResponse is the response type for the Query/DTagTransfers
// RPC method.
type QueryDTagTransfersResponse struct {
	// relationships represent the list of all the blocks for the queried user
	Requests []DTagTransferRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests"`
}

func (m *QueryDTagTransfersResponse) Reset()         { *m = QueryDTagTransfersResponse{} }
func (m *QueryDTagTransfersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDTagTransfersResponse) ProtoMessage()    {}
func (*QueryDTagTransfersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0074f57a59f38d, []int{3}
}
func (m *QueryDTagTransfersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDTagTransfersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDTagTransfersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDTagTransfersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDTagTransfersResponse.Merge(m, src)
}
func (m *QueryDTagTransfersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDTagTransfersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDTagTransfersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDTagTransfersResponse proto.InternalMessageInfo

func (m *QueryDTagTransfersResponse) GetRequests() []DTagTransferRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

// QueryParamsRequest is the request type for the Query/Params RPC endpoint
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0074f57a59f38d, []int{4}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e0074f57a59f38d, []int{5}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryProfileRequest)(nil), "desmos.profiles.v1beta1.QueryProfileRequest")
	proto.RegisterType((*QueryProfileResponse)(nil), "desmos.profiles.v1beta1.QueryProfileResponse")
	proto.RegisterType((*QueryDTagTransfersRequest)(nil), "desmos.profiles.v1beta1.QueryDTagTransfersRequest")
	proto.RegisterType((*QueryDTagTransfersResponse)(nil), "desmos.profiles.v1beta1.QueryDTagTransfersResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "desmos.profiles.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "desmos.profiles.v1beta1.QueryParamsResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/query.proto", fileDescriptor_5e0074f57a59f38d)
}

var fileDescriptor_5e0074f57a59f38d = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0xf7, 0xd1, 0xd2, 0x86, 0x2b, 0x2c, 0x47, 0x24, 0x5a, 0x0b, 0xd9, 0xc5, 0x20, 0x35, 0x40,
	0x73, 0x47, 0x93, 0x09, 0x04, 0x43, 0x23, 0x06, 0x58, 0x10, 0x44, 0x99, 0x58, 0xaa, 0x73, 0x72,
	0x31, 0x91, 0x1c, 0x9f, 0xeb, 0x3b, 0x23, 0x22, 0xc4, 0xc2, 0xc4, 0xc0, 0x80, 0xc4, 0xc2, 0xd8,
	0x95, 0x9d, 0x0f, 0x51, 0x31, 0x55, 0x62, 0x61, 0x01, 0xa1, 0x84, 0x81, 0x8f, 0x81, 0x72, 0xf7,
	0x1c, 0x5a, 0x29, 0x6e, 0xc3, 0x66, 0xfb, 0x7e, 0xff, 0xde, 0xfb, 0xf9, 0xf0, 0xf5, 0x9e, 0x50,
	0x43, 0xa9, 0x58, 0x9a, 0xc9, 0xfe, 0x20, 0x16, 0x8a, 0xbd, 0xdc, 0x09, 0x85, 0xe6, 0x3b, 0x6c,
	0x3f, 0x17, 0xd9, 0x88, 0xa6, 0x99, 0xd4, 0x92, 0x5c, 0xb1, 0x20, 0x5a, 0x80, 0x28, 0x80, 0xdc,
	0x6a, 0x24, 0x23, 0x69, 0x30, 0x6c, 0xfa, 0x64, 0xe1, 0xee, 0xd5, 0x48, 0xca, 0x28, 0x16, 0x8c,
	0xa7, 0x03, 0xc6, 0x93, 0x44, 0x6a, 0xae, 0x07, 0x32, 0x51, 0x70, 0xba, 0x01, 0xa7, 0xe6, 0x2d,
	0xcc, 0xfb, 0x8c, 0x27, 0xe0, 0xe3, 0xde, 0x28, 0x0b, 0x93, 0xf2, 0x8c, 0x0f, 0xd5, 0x59, 0xa8,
	0xa1, 0xec, 0x89, 0x78, 0x66, 0xd3, 0x95, 0x53, 0xd4, 0x9e, 0x4d, 0x67, 0x5f, 0xec, 0x51, 0xd0,
	0xc4, 0x97, 0x9f, 0x4d, 0xa7, 0x7b, 0x6a, 0x05, 0xda, 0x62, 0x3f, 0x17, 0x4a, 0x13, 0x82, 0x97,
	0x73, 0x25, 0xb2, 0x75, 0xb4, 0x89, 0x6a, 0x17, 0xda, 0xe6, 0xf9, 0x5e, 0xe5, 0xdd, 0x81, 0xef,
	0xfc, 0x39, 0xf0, 0x9d, 0xa0, 0x83, 0xab, 0x27, 0x49, 0x2a, 0x95, 0x89, 0x12, 0xe4, 0x3e, 0x5e,
	0x85, 0x20, 0x86, 0xb8, 0xd6, 0xa8, 0x52, 0x3b, 0x20, 0x2d, 0x06, 0xa4, 0xbb, 0xc9, 0xa8, 0x75,
	0xf1, 0xeb, 0x97, 0x7a, 0x65, 0xb7, 0xdb, 0x95, 0x79, 0xa2, 0x1f, 0xb7, 0x0b, 0x4a, 0x70, 0x17,
	0x6f, 0x18, 0xd5, 0x87, 0x1d, 0x1e, 0x75, 0x32, 0x9e, 0xa8, 0xbe, 0xc8, 0xd4, 0x62, 0x81, 0x62,
	0xec, 0xce, 0xa3, 0x42, 0xac, 0x27, 0xb8, 0x92, 0x59, 0x19, 0xb5, 0x8e, 0x36, 0x97, 0x6a, 0x6b,
	0x8d, 0x6d, 0x5a, 0xd2, 0x22, 0x3d, 0xae, 0x00, 0xde, 0xad, 0xe5, 0xc3, 0x9f, 0xbe, 0xd3, 0x9e,
	0x69, 0x04, 0x55, 0x4c, 0xec, 0xf8, 0xa6, 0x09, 0x40, 0x05, 0x9d, 0x62, 0x93, 0xf0, 0x15, 0xcc,
	0x1f, 0xe0, 0x15, 0xdb, 0x18, 0xac, 0xc4, 0x2f, 0xb5, 0xb6, 0x44, 0x70, 0x03, 0x52, 0xe3, 0xc7,
	0x12, 0x3e, 0x6f, 0x64, 0xc9, 0x27, 0x84, 0x57, 0x61, 0xe1, 0xa4, 0x3c, 0xff, 0x9c, 0x32, 0xdd,
	0xfa, 0x82, 0x68, 0x9b, 0x38, 0xb8, 0xf3, 0xf6, 0xdb, 0xef, 0x8f, 0xe7, 0x6e, 0x91, 0x1a, 0x2b,
	0xfd, 0x05, 0x8b, 0x0f, 0xaf, 0xa7, 0x3d, 0xbc, 0x21, 0x9f, 0x11, 0xbe, 0x74, 0x62, 0xf5, 0xa4,
	0x71, 0xba, 0xe5, 0xbc, 0x8a, 0xdd, 0xe6, 0x7f, 0x71, 0x20, 0x2c, 0x33, 0x61, 0x6f, 0x92, 0xad,
	0xd2, 0xb0, 0x3d, 0xcd, 0xa3, 0x3d, 0x3d, 0x4b, 0xf6, 0x1e, 0xe1, 0x15, 0xbb, 0x69, 0x72, 0xfb,
	0x8c, 0xbd, 0x1c, 0xaf, 0xd7, 0xdd, 0x5e, 0x0c, 0x0c, 0xb1, 0xb6, 0x4c, 0xac, 0x6b, 0xc4, 0x67,
	0xa7, 0x5f, 0xe3, 0xd6, 0xa3, 0xc3, 0xb1, 0x87, 0x8e, 0xc6, 0x1e, 0xfa, 0x35, 0xf6, 0xd0, 0x87,
	0x89, 0xe7, 0x1c, 0x4d, 0x3c, 0xe7, 0xfb, 0xc4, 0x73, 0x9e, 0xd3, 0x68, 0xa0, 0x5f, 0xe4, 0x21,
	0xed, 0xca, 0x21, 0x88, 0xd4, 0x63, 0x1e, 0xaa, 0x42, 0xf0, 0xd5, 0x3f, 0x49, 0x3d, 0x4a, 0x85,
	0x0a, 0x57, 0xcc, 0x1d, 0x6b, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x72, 0x92, 0xdc, 0x7f, 0xc6,
	0x04, 0x00, 0x00,
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
	// Profile queries the profile of a specific user
	Profile(ctx context.Context, in *QueryProfileRequest, opts ...grpc.CallOption) (*QueryProfileResponse, error)
	// DTagTransfers queries all the DTag transfers requests
	DTagTransfers(ctx context.Context, in *QueryDTagTransfersRequest, opts ...grpc.CallOption) (*QueryDTagTransfersResponse, error)
	// Params queries the profiles module params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Profile(ctx context.Context, in *QueryProfileRequest, opts ...grpc.CallOption) (*QueryProfileResponse, error) {
	out := new(QueryProfileResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Query/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DTagTransfers(ctx context.Context, in *QueryDTagTransfersRequest, opts ...grpc.CallOption) (*QueryDTagTransfersResponse, error) {
	out := new(QueryDTagTransfersResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Query/DTagTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Profile queries the profile of a specific user
	Profile(context.Context, *QueryProfileRequest) (*QueryProfileResponse, error)
	// DTagTransfers queries all the DTag transfers requests
	DTagTransfers(context.Context, *QueryDTagTransfersRequest) (*QueryDTagTransfersResponse, error)
	// Params queries the profiles module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Profile(ctx context.Context, req *QueryProfileRequest) (*QueryProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (*UnimplementedQueryServer) DTagTransfers(ctx context.Context, req *QueryDTagTransfersRequest) (*QueryDTagTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DTagTransfers not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Profile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Profile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Query/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Profile(ctx, req.(*QueryProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DTagTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDTagTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DTagTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Query/DTagTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DTagTransfers(ctx, req.(*QueryDTagTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.profiles.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Profile",
			Handler:    _Query_Profile_Handler,
		},
		{
			MethodName: "DTagTransfers",
			Handler:    _Query_DTagTransfers_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/profiles/v1beta1/query.proto",
}

func (m *QueryProfileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProfileRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProfileRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Profile != nil {
		{
			size, err := m.Profile.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDTagTransfersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDTagTransfersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDTagTransfersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDTagTransfersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDTagTransfersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDTagTransfersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
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
func (m *QueryProfileRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Profile != nil {
		l = m.Profile.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDTagTransfersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDTagTransfersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryProfileRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryProfileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProfileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
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
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryProfileResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Profile == nil {
				m.Profile = &types.Any{}
			}
			if err := m.Profile.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryDTagTransfersRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDTagTransfersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDTagTransfersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
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
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryDTagTransfersResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDTagTransfersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDTagTransfersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, DTagTransferRequest{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
