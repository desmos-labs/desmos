// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v2/query.proto

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

func init() { proto.RegisterFile("desmos/profiles/v2/query.proto", fileDescriptor_fba4df0d7bde4d7c) }

var fileDescriptor_fba4df0d7bde4d7c = []byte{
	// 613 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x3d, 0x6f, 0xd3, 0x40,
	0x18, 0xc7, 0x6b, 0x24, 0x8a, 0x74, 0x0b, 0xe8, 0x04, 0x43, 0xa2, 0xe0, 0x8a, 0x42, 0x5e, 0x94,
	0x26, 0x3e, 0x92, 0x02, 0x52, 0x11, 0x0c, 0xb4, 0x5d, 0x2a, 0x90, 0x28, 0xa8, 0x13, 0x4b, 0x74,
	0x76, 0x2e, 0xee, 0x09, 0xe7, 0xee, 0xea, 0x73, 0x02, 0x51, 0x95, 0x85, 0x8d, 0x05, 0x21, 0xb1,
	0x33, 0xf1, 0x11, 0xf8, 0x08, 0x0c, 0x2c, 0x88, 0x4a, 0x2c, 0x8c, 0x28, 0xe1, 0x83, 0x20, 0xdf,
	0x4b, 0xa3, 0x86, 0x38, 0x49, 0xbb, 0xd9, 0xbe, 0xdf, 0xff, 0x9e, 0xdf, 0x73, 0x7e, 0x74, 0xc0,
	0x6d, 0x13, 0xd9, 0xe5, 0x12, 0x89, 0x98, 0x77, 0x68, 0x44, 0x24, 0xea, 0x37, 0xd1, 0x51, 0x8f,
	0xc4, 0x03, 0x4f, 0xc4, 0x3c, 0xe1, 0x10, 0xea, 0x75, 0xcf, 0xae, 0x7b, 0xfd, 0x66, 0xfe, 0x7a,
	0xc8, 0x43, 0xae, 0x96, 0x51, 0xfa, 0xa4, 0xc9, 0x7c, 0x21, 0xe4, 0x3c, 0x8c, 0x08, 0xc2, 0x82,
	0x22, 0xcc, 0x18, 0x4f, 0x70, 0x42, 0x39, 0x93, 0x66, 0x35, 0x67, 0x56, 0xd5, 0x9b, 0xdf, 0xeb,
	0x20, 0xcc, 0x4c, 0x89, 0x7c, 0x29, 0x4b, 0xa1, 0x65, 0xbe, 0x18, 0xae, 0x96, 0xc9, 0xb5, 0x13,
	0x1c, 0xb6, 0x62, 0x72, 0xd4, 0x23, 0x32, 0xb1, 0x05, 0x8b, 0xd9, 0xbb, 0xe2, 0x18, 0x77, 0x2d,
	0x56, 0xcd, 0xc4, 0x82, 0x43, 0x4c, 0x59, 0x2b, 0xa2, 0xec, 0xb5, 0x65, 0x2b, 0x99, 0x2c, 0x16,
	0xe2, 0x0c, 0x99, 0x0b, 0x78, 0x4a, 0xb6, 0xf4, 0x21, 0xe9, 0x17, 0x5b, 0x50, 0xbf, 0x21, 0x1f,
	0x4b, 0xa2, 0xd3, 0xa8, 0xdf, 0xf0, 0x49, 0x82, 0x1b, 0x48, 0xe0, 0x90, 0x32, 0x75, 0x6a, 0x9a,
	0x6d, 0x7e, 0x03, 0xe0, 0xf2, 0x8b, 0x14, 0x81, 0xef, 0x1d, 0x70, 0x65, 0x5f, 0x97, 0x85, 0x65,
	0xef, 0xff, 0x7f, 0xe2, 0x29, 0xcc, 0x10, 0x2f, 0xf5, 0x49, 0xe4, 0x2b, 0x8b, 0x41, 0x29, 0x38,
	0x93, 0x64, 0x7d, 0xe3, 0xdd, 0xaf, 0xbf, 0x9f, 0x2e, 0x15, 0xe1, 0x6d, 0x34, 0xa3, 0xc5, 0xd3,
	0xe7, 0xe3, 0x9e, 0x24, 0xf1, 0x10, 0xfe, 0x74, 0x40, 0x61, 0x8f, 0x05, 0xbc, 0x4b, 0x59, 0xb8,
	0x7b, 0x80, 0xc3, 0x83, 0x18, 0x33, 0xd9, 0x21, 0xb1, 0x29, 0x2b, 0xe1, 0xa3, 0xcc, 0xba, 0xf3,
	0x62, 0xd6, 0xfa, 0xf1, 0x05, 0xd3, 0xa6, 0x95, 0xa6, 0x6a, 0xa5, 0x06, 0xab, 0xb3, 0x5a, 0x49,
	0x07, 0xa5, 0x9e, 0x98, 0x68, 0xdd, 0x4e, 0x0c, 0xfc, 0xe0, 0x00, 0xb0, 0x93, 0xfe, 0xee, 0x67,
	0xe9, 0x3f, 0x84, 0xd5, 0x4c, 0x83, 0x09, 0x64, 0x6d, 0x37, 0x96, 0x62, 0x8d, 0x5b, 0x59, 0xb9,
	0xdd, 0x82, 0x6b, 0xb3, 0xdc, 0xd4, 0xbc, 0xd5, 0xd5, 0x14, 0xc1, 0x2f, 0x0e, 0xb8, 0x7a, 0x9a,
	0x7f, 0xfe, 0x86, 0x91, 0x58, 0x42, 0xb4, 0xb8, 0x92, 0x26, 0xad, 0xda, 0xdd, 0xe5, 0x03, 0xc6,
	0xcf, 0x53, 0x7e, 0x15, 0x58, 0x5a, 0xe0, 0x87, 0xb8, 0x56, 0xfa, 0xec, 0x80, 0x6b, 0x4f, 0x84,
	0x88, 0x68, 0xa0, 0xa6, 0x56, 0x9f, 0x5e, 0x76, 0xd9, 0x69, 0xd4, 0x8a, 0x36, 0xce, 0x91, 0x30,
	0xa6, 0x45, 0x65, 0xba, 0x06, 0x6f, 0xce, 0x32, 0xc5, 0x42, 0x98, 0x73, 0xfc, 0xe1, 0x80, 0xdc,
	0xd4, 0x1e, 0xdb, 0x83, 0x9d, 0x88, 0x12, 0x96, 0xec, 0xed, 0xc2, 0xad, 0x65, 0xeb, 0x4e, 0x32,
	0x56, 0xf9, 0xe1, 0x45, 0xa2, 0xc6, 0x7d, 0x4b, 0xb9, 0x6f, 0xc2, 0xc6, 0x5c, 0x77, 0x14, 0xa8,
	0x9c, 0x44, 0xc7, 0xfa, 0xa1, 0x45, 0xdb, 0x43, 0xf8, 0xd5, 0x01, 0x37, 0xa6, 0x0a, 0x98, 0xe9,
	0xb8, 0xbf, 0xac, 0xd0, 0xd9, 0x19, 0x79, 0x70, 0xde, 0x98, 0xe9, 0xa1, 0xa6, 0x7a, 0x28, 0xc1,
	0x3b, 0xf3, 0x7b, 0x30, 0x73, 0x32, 0x04, 0xab, 0xfb, 0xea, 0xd2, 0x85, 0xa5, 0xec, 0x2b, 0x49,
	0x01, 0xd6, 0xab, 0xbc, 0x90, 0x33, 0x22, 0xeb, 0x4a, 0xa4, 0x00, 0xf3, 0x33, 0x6f, 0x2e, 0xc5,
	0x6e, 0x3f, 0xfd, 0x3e, 0x72, 0x9d, 0x93, 0x91, 0xeb, 0xfc, 0x19, 0xb9, 0xce, 0xc7, 0xb1, 0xbb,
	0x72, 0x32, 0x76, 0x57, 0x7e, 0x8f, 0xdd, 0x95, 0x57, 0x8d, 0x90, 0x26, 0x87, 0x3d, 0xdf, 0x0b,
	0x78, 0xd7, 0xe4, 0xeb, 0x11, 0xf6, 0xa5, 0xdd, 0xab, 0x7f, 0x0f, 0xbd, 0x9d, 0x6c, 0x98, 0x0c,
	0x04, 0x91, 0xfe, 0xaa, 0xba, 0x9a, 0x37, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x59, 0x36,
	0x67, 0x39, 0x07, 0x00, 0x00,
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
	// ChainLinks queries the chain links associated to the given user, if
	// provided. Otherwise it queries all the chain links stored.
	ChainLinks(ctx context.Context, in *QueryChainLinksRequest, opts ...grpc.CallOption) (*QueryChainLinksResponse, error)
	// ChainLinkOwners queries for the owners of chain links, optionally searching
	// for a specific chain name and external address
	ChainLinkOwners(ctx context.Context, in *QueryChainLinkOwnersRequest, opts ...grpc.CallOption) (*QueryChainLinkOwnersResponse, error)
	// ApplicationLinks queries the applications links associated to the given
	// user, if provided. Otherwise, it queries all the application links stored.
	ApplicationLinks(ctx context.Context, in *QueryApplicationLinksRequest, opts ...grpc.CallOption) (*QueryApplicationLinksResponse, error)
	// ApplicationLinkByClientID queries a single application link for a given
	// client id.
	ApplicationLinkByClientID(ctx context.Context, in *QueryApplicationLinkByClientIDRequest, opts ...grpc.CallOption) (*QueryApplicationLinkByClientIDResponse, error)
	// ApplicationLinkOwners queries for the owners of applications links,
	// optionally searching for a specific application and username.
	ApplicationLinkOwners(ctx context.Context, in *QueryApplicationLinkOwnersRequest, opts ...grpc.CallOption) (*QueryApplicationLinkOwnersResponse, error)
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
	err := c.cc.Invoke(ctx, "/desmos.profiles.v2.Query/Profile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IncomingDTagTransferRequests(ctx context.Context, in *QueryIncomingDTagTransferRequestsRequest, opts ...grpc.CallOption) (*QueryIncomingDTagTransferRequestsResponse, error) {
	out := new(QueryIncomingDTagTransferRequestsResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v2.Query/IncomingDTagTransferRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainLinks(ctx context.Context, in *QueryChainLinksRequest, opts ...grpc.CallOption) (*QueryChainLinksResponse, error) {
	out := new(QueryChainLinksResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v2.Query/ChainLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainLinkOwners(ctx context.Context, in *QueryChainLinkOwnersRequest, opts ...grpc.CallOption) (*QueryChainLinkOwnersResponse, error) {
	out := new(QueryChainLinkOwnersResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v2.Query/ChainLinkOwners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ApplicationLinks(ctx context.Context, in *QueryApplicationLinksRequest, opts ...grpc.CallOption) (*QueryApplicationLinksResponse, error) {
	out := new(QueryApplicationLinksResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v2.Query/ApplicationLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ApplicationLinkByClientID(ctx context.Context, in *QueryApplicationLinkByClientIDRequest, opts ...grpc.CallOption) (*QueryApplicationLinkByClientIDResponse, error) {
	out := new(QueryApplicationLinkByClientIDResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v2.Query/ApplicationLinkByClientID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ApplicationLinkOwners(ctx context.Context, in *QueryApplicationLinkOwnersRequest, opts ...grpc.CallOption) (*QueryApplicationLinkOwnersResponse, error) {
	out := new(QueryApplicationLinkOwnersResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v2.Query/ApplicationLinkOwners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v2.Query/Params", in, out, opts...)
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
	// ChainLinks queries the chain links associated to the given user, if
	// provided. Otherwise it queries all the chain links stored.
	ChainLinks(context.Context, *QueryChainLinksRequest) (*QueryChainLinksResponse, error)
	// ChainLinkOwners queries for the owners of chain links, optionally searching
	// for a specific chain name and external address
	ChainLinkOwners(context.Context, *QueryChainLinkOwnersRequest) (*QueryChainLinkOwnersResponse, error)
	// ApplicationLinks queries the applications links associated to the given
	// user, if provided. Otherwise, it queries all the application links stored.
	ApplicationLinks(context.Context, *QueryApplicationLinksRequest) (*QueryApplicationLinksResponse, error)
	// ApplicationLinkByClientID queries a single application link for a given
	// client id.
	ApplicationLinkByClientID(context.Context, *QueryApplicationLinkByClientIDRequest) (*QueryApplicationLinkByClientIDResponse, error)
	// ApplicationLinkOwners queries for the owners of applications links,
	// optionally searching for a specific application and username.
	ApplicationLinkOwners(context.Context, *QueryApplicationLinkOwnersRequest) (*QueryApplicationLinkOwnersResponse, error)
	// Params queries the profiles module params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
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
func (*UnimplementedQueryServer) ChainLinks(ctx context.Context, req *QueryChainLinksRequest) (*QueryChainLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainLinks not implemented")
}
func (*UnimplementedQueryServer) ChainLinkOwners(ctx context.Context, req *QueryChainLinkOwnersRequest) (*QueryChainLinkOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainLinkOwners not implemented")
}
func (*UnimplementedQueryServer) ApplicationLinks(ctx context.Context, req *QueryApplicationLinksRequest) (*QueryApplicationLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationLinks not implemented")
}
func (*UnimplementedQueryServer) ApplicationLinkByClientID(ctx context.Context, req *QueryApplicationLinkByClientIDRequest) (*QueryApplicationLinkByClientIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationLinkByClientID not implemented")
}
func (*UnimplementedQueryServer) ApplicationLinkOwners(ctx context.Context, req *QueryApplicationLinkOwnersRequest) (*QueryApplicationLinkOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationLinkOwners not implemented")
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
		FullMethod: "/desmos.profiles.v2.Query/Profile",
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
		FullMethod: "/desmos.profiles.v2.Query/IncomingDTagTransferRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IncomingDTagTransferRequests(ctx, req.(*QueryIncomingDTagTransferRequestsRequest))
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
		FullMethod: "/desmos.profiles.v2.Query/ChainLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainLinks(ctx, req.(*QueryChainLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ChainLinkOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryChainLinkOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ChainLinkOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v2.Query/ChainLinkOwners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainLinkOwners(ctx, req.(*QueryChainLinkOwnersRequest))
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
		FullMethod: "/desmos.profiles.v2.Query/ApplicationLinks",
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
		FullMethod: "/desmos.profiles.v2.Query/ApplicationLinkByClientID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ApplicationLinkByClientID(ctx, req.(*QueryApplicationLinkByClientIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ApplicationLinkOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryApplicationLinkOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ApplicationLinkOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v2.Query/ApplicationLinkOwners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ApplicationLinkOwners(ctx, req.(*QueryApplicationLinkOwnersRequest))
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
		FullMethod: "/desmos.profiles.v2.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.profiles.v2.Query",
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
			MethodName: "ChainLinks",
			Handler:    _Query_ChainLinks_Handler,
		},
		{
			MethodName: "ChainLinkOwners",
			Handler:    _Query_ChainLinkOwners_Handler,
		},
		{
			MethodName: "ApplicationLinks",
			Handler:    _Query_ApplicationLinks_Handler,
		},
		{
			MethodName: "ApplicationLinkByClientID",
			Handler:    _Query_ApplicationLinkByClientID_Handler,
		},
		{
			MethodName: "ApplicationLinkOwners",
			Handler:    _Query_ApplicationLinkOwners_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/profiles/v2/query.proto",
}
