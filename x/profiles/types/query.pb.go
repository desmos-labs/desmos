// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

func init() { proto.RegisterFile("desmos/profiles/v3/query.proto", fileDescriptor_bcbdebc2a1cf2f2b) }

var fileDescriptor_bcbdebc2a1cf2f2b = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xc7, 0x63, 0x24, 0x5a, 0xc9, 0x13, 0x3a, 0xb1, 0xc4, 0x0a, 0xae, 0x28, 0x4a, 0x53, 0xa5,
	0xad, 0x0f, 0x27, 0x53, 0x11, 0x0c, 0xb4, 0x5d, 0x2a, 0x18, 0x0a, 0xea, 0xc4, 0x62, 0x9d, 0x9d,
	0x8b, 0x7b, 0xc2, 0xb9, 0xbb, 0xfa, 0xce, 0x11, 0x51, 0x95, 0x85, 0x8d, 0x05, 0x21, 0xb1, 0xf3,
	0x79, 0x58, 0x10, 0x95, 0x58, 0x18, 0x51, 0x82, 0xf8, 0x1c, 0xc8, 0xf7, 0xd2, 0x88, 0x62, 0xf7,
	0x25, 0xdb, 0x5d, 0xee, 0xf7, 0xdc, 0xff, 0x77, 0x4f, 0x1e, 0xd9, 0xf5, 0x07, 0x58, 0x8c, 0x98,
	0x80, 0x3c, 0x67, 0x43, 0x92, 0x61, 0x01, 0xc7, 0x7d, 0x78, 0x5a, 0xe0, 0x7c, 0x12, 0xf0, 0x9c,
	0x49, 0x06, 0x80, 0x3e, 0x0f, 0xec, 0x79, 0x30, 0xee, 0x7b, 0xf7, 0x53, 0x96, 0x32, 0x75, 0x0c,
	0xcb, 0x95, 0x26, 0xbd, 0x56, 0xca, 0x58, 0x9a, 0x61, 0x88, 0x38, 0x81, 0x88, 0x52, 0x26, 0x91,
	0x24, 0x8c, 0x0a, 0x73, 0xda, 0x34, 0xa7, 0x6a, 0x17, 0x17, 0x43, 0x88, 0xa8, 0x89, 0xf0, 0x36,
	0xea, 0x14, 0x22, 0xf3, 0x8b, 0xe1, 0xb6, 0x6b, 0xb9, 0x81, 0x44, 0x69, 0x94, 0xe3, 0xd3, 0x02,
	0x0b, 0x69, 0x03, 0xdb, 0xf5, 0xb7, 0xa2, 0x1c, 0x8d, 0x2c, 0xd6, 0xad, 0xc5, 0x92, 0x13, 0x44,
	0x68, 0x94, 0x11, 0xfa, 0xd6, 0xb2, 0x9b, 0xb5, 0x2c, 0xe2, 0xfc, 0x1f, 0xb2, 0x99, 0xb0, 0x92,
	0x8c, 0x74, 0x93, 0xf4, 0xc6, 0x06, 0xea, 0x1d, 0x8c, 0x91, 0xc0, 0xba, 0x1a, 0x8e, 0xc3, 0x18,
	0x4b, 0x14, 0x42, 0x8e, 0x52, 0x42, 0x55, 0xd7, 0x34, 0xdb, 0xfb, 0xb3, 0xea, 0xde, 0x7d, 0x55,
	0x22, 0xe0, 0x83, 0xe3, 0xae, 0x1e, 0xe9, 0x58, 0xd0, 0x09, 0xfe, 0xff, 0x4f, 0x02, 0x85, 0x19,
	0xe2, 0xb5, 0xee, 0x84, 0xb7, 0x79, 0x3d, 0x28, 0x38, 0xa3, 0x02, 0xaf, 0x6f, 0xbd, 0xff, 0xf1,
	0xfb, 0xf3, 0x9d, 0x36, 0x78, 0x04, 0x2b, 0x9e, 0x78, 0xb1, 0x3e, 0x2b, 0x04, 0xce, 0xa7, 0xe0,
	0xbb, 0xe3, 0xb6, 0x0e, 0x69, 0xc2, 0x46, 0x84, 0xa6, 0x07, 0xc7, 0x28, 0x3d, 0xce, 0x11, 0x15,
	0x43, 0x9c, 0x9b, 0x58, 0x01, 0x9e, 0xd6, 0xe6, 0x5e, 0x55, 0x66, 0xad, 0x9f, 0x2d, 0x59, 0x6d,
	0x9e, 0xd2, 0x53, 0x4f, 0xd9, 0x06, 0xdd, 0xaa, 0xa7, 0xa8, 0x41, 0x91, 0xa6, 0xf4, 0x62, 0x62,
	0xc0, 0xd4, 0x5d, 0x39, 0x52, 0x43, 0x01, 0x36, 0xea, 0x5b, 0xa6, 0x00, 0x2b, 0xd9, 0xb9, 0x96,
	0x33, 0x3a, 0xeb, 0x4a, 0xa7, 0x05, 0xbc, 0xca, 0xce, 0xea, 0xd0, 0x8f, 0x8e, 0xeb, 0xee, 0x97,
	0xd3, 0xf6, 0xb2, 0x1c, 0x21, 0xd0, 0xad, 0xbd, 0x7b, 0x01, 0x59, 0x8f, 0xad, 0x1b, 0xb1, 0xc6,
	0xa5, 0xa3, 0x5c, 0x1e, 0x82, 0xb5, 0x2a, 0x17, 0x35, 0xee, 0x3b, 0x6a, 0x88, 0xc1, 0x17, 0xc7,
	0xbd, 0xf7, 0x9c, 0xf3, 0x8c, 0x24, 0x6a, 0x1a, 0xb5, 0xd6, 0xe3, 0xda, 0xa8, 0xcb, 0xa8, 0x95,
	0x0b, 0x6f, 0x51, 0x61, 0x14, 0xdb, 0x4a, 0x71, 0x0d, 0x3c, 0xa8, 0x52, 0x44, 0x9c, 0x1b, 0xc1,
	0x6f, 0x8e, 0xdb, 0xbc, 0x74, 0xc7, 0xde, 0x64, 0x3f, 0x23, 0x98, 0xca, 0xc3, 0x03, 0xb0, 0x7b,
	0xd3, 0xdc, 0x45, 0x8d, 0x55, 0x7e, 0xb2, 0x4c, 0xa9, 0x71, 0xdf, 0x55, 0xee, 0x7d, 0x10, 0x5e,
	0xe9, 0x0e, 0x13, 0x55, 0x27, 0xe0, 0x99, 0x5e, 0x44, 0x64, 0x30, 0xdd, 0x7b, 0xf1, 0x75, 0xe6,
	0x3b, 0xe7, 0x33, 0xdf, 0xf9, 0x35, 0xf3, 0x9d, 0x4f, 0x73, 0xbf, 0x71, 0x3e, 0xf7, 0x1b, 0x3f,
	0xe7, 0x7e, 0xe3, 0x4d, 0x98, 0x12, 0x79, 0x52, 0xc4, 0x41, 0xc2, 0x46, 0xe6, 0xda, 0x9d, 0x0c,
	0xc5, 0xc2, 0x46, 0x8c, 0x7b, 0xf0, 0xdd, 0x22, 0x47, 0x4e, 0x38, 0x16, 0xf1, 0x8a, 0xfa, 0x78,
	0xf4, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x50, 0xc7, 0x0c, 0xa0, 0xdb, 0x05, 0x00, 0x00,
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
	// Profile queries the profile of a specific user given their DTag or address.
	// If the queried user does not have a profile, the returned response will
	// contain a null profile.
	Profile(ctx context.Context, in *QueryProfileRequest, opts ...grpc.CallOption) (*QueryProfileResponse, error)
	// IncomingDTagTransferRequests queries all the DTag transfers requests that
	// have been made towards the user with the given address
	IncomingDTagTransferRequests(ctx context.Context, in *QueryIncomingDTagTransferRequestsRequest, opts ...grpc.CallOption) (*QueryIncomingDTagTransferRequestsResponse, error)
	// Params queries the profiles module params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// ChainLinks queries the chain links associated to the given user, if
	// provided. Otherwise it queries all the chain links stored.
	ChainLinks(ctx context.Context, in *QueryChainLinksRequest, opts ...grpc.CallOption) (*QueryChainLinksResponse, error)
	// ApplicationLinks queries the applications links associated to the given
	// user, if provided. Otherwise, it queries all the application links stored.
	ApplicationLinks(ctx context.Context, in *QueryApplicationLinksRequest, opts ...grpc.CallOption) (*QueryApplicationLinksResponse, error)
	// ApplicationLinkByClientID queries a single application link for a given
	// client id.
	ApplicationLinkByClientID(ctx context.Context, in *QueryApplicationLinkByClientIDRequest, opts ...grpc.CallOption) (*QueryApplicationLinkByClientIDResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Profile(ctx context.Context, in *QueryProfileRequest, opts ...grpc.CallOption) (*QueryProfileResponse, error) {
	out := new(QueryProfileResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IncomingDTagTransferRequests(ctx context.Context, in *QueryIncomingDTagTransferRequestsRequest, opts ...grpc.CallOption) (*QueryIncomingDTagTransferRequestsResponse, error) {
	out := new(QueryIncomingDTagTransferRequestsResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/IncomingDTagTransferRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainLinks(ctx context.Context, in *QueryChainLinksRequest, opts ...grpc.CallOption) (*QueryChainLinksResponse, error) {
	out := new(QueryChainLinksResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/ChainLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ApplicationLinks(ctx context.Context, in *QueryApplicationLinksRequest, opts ...grpc.CallOption) (*QueryApplicationLinksResponse, error) {
	out := new(QueryApplicationLinksResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/ApplicationLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ApplicationLinkByClientID(ctx context.Context, in *QueryApplicationLinkByClientIDRequest, opts ...grpc.CallOption) (*QueryApplicationLinkByClientIDResponse, error) {
	out := new(QueryApplicationLinkByClientIDResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/ApplicationLinkByClientID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Profile queries the profile of a specific user given their DTag or address.
	// If the queried user does not have a profile, the returned response will
	// contain a null profile.
	Profile(context.Context, *QueryProfileRequest) (*QueryProfileResponse, error)
	// IncomingDTagTransferRequests queries all the DTag transfers requests that
	// have been made towards the user with the given address
	IncomingDTagTransferRequests(context.Context, *QueryIncomingDTagTransferRequestsRequest) (*QueryIncomingDTagTransferRequestsResponse, error)
	// Params queries the profiles module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// ChainLinks queries the chain links associated to the given user, if
	// provided. Otherwise it queries all the chain links stored.
	ChainLinks(context.Context, *QueryChainLinksRequest) (*QueryChainLinksResponse, error)
	// ApplicationLinks queries the applications links associated to the given
	// user, if provided. Otherwise, it queries all the application links stored.
	ApplicationLinks(context.Context, *QueryApplicationLinksRequest) (*QueryApplicationLinksResponse, error)
	// ApplicationLinkByClientID queries a single application link for a given
	// client id.
	ApplicationLinkByClientID(context.Context, *QueryApplicationLinkByClientIDRequest) (*QueryApplicationLinkByClientIDResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Profile(ctx context.Context, req *QueryProfileRequest) (*QueryProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Profile not implemented")
}
func (*UnimplementedQueryServer) IncomingDTagTransferRequests(ctx context.Context, req *QueryIncomingDTagTransferRequestsRequest) (*QueryIncomingDTagTransferRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncomingDTagTransferRequests not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) ChainLinks(ctx context.Context, req *QueryChainLinksRequest) (*QueryChainLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainLinks not implemented")
}
func (*UnimplementedQueryServer) ApplicationLinks(ctx context.Context, req *QueryApplicationLinksRequest) (*QueryApplicationLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationLinks not implemented")
}
func (*UnimplementedQueryServer) ApplicationLinkByClientID(ctx context.Context, req *QueryApplicationLinkByClientIDRequest) (*QueryApplicationLinkByClientIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationLinkByClientID not implemented")
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
		FullMethod: "/desmos.profiles.v3.Query/Profile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Profile(ctx, req.(*QueryProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IncomingDTagTransferRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIncomingDTagTransferRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IncomingDTagTransferRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Query/IncomingDTagTransferRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IncomingDTagTransferRequests(ctx, req.(*QueryIncomingDTagTransferRequestsRequest))
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
		FullMethod: "/desmos.profiles.v3.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChainLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChainLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChainLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Query/ChainLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainLinks(ctx, req.(*QueryChainLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ApplicationLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryApplicationLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ApplicationLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Query/ApplicationLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ApplicationLinks(ctx, req.(*QueryApplicationLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ApplicationLinkByClientID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryApplicationLinkByClientIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ApplicationLinkByClientID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Query/ApplicationLinkByClientID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ApplicationLinkByClientID(ctx, req.(*QueryApplicationLinkByClientIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.profiles.v3.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Profile",
			Handler:    _Query_Profile_Handler,
		},
		{
			MethodName: "IncomingDTagTransferRequests",
			Handler:    _Query_IncomingDTagTransferRequests_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "ChainLinks",
			Handler:    _Query_ChainLinks_Handler,
		},
		{
			MethodName: "ApplicationLinks",
			Handler:    _Query_ApplicationLinks_Handler,
		},
		{
			MethodName: "ApplicationLinkByClientID",
			Handler:    _Query_ApplicationLinkByClientID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/profiles/v3/query.proto",
}
