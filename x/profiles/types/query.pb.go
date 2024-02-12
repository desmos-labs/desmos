// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0x4f, 0x6b, 0x13, 0x4f,
	0x18, 0xc7, 0xbb, 0x3f, 0xf8, 0x55, 0x18, 0x04, 0xe5, 0x01, 0x41, 0x43, 0xdc, 0x62, 0x35, 0x7f,
	0x48, 0xb3, 0xbb, 0x4d, 0x83, 0xd5, 0x88, 0x1e, 0xda, 0xc6, 0x43, 0x51, 0xb0, 0x4a, 0x4f, 0x5e,
	0xc2, 0x24, 0x99, 0x6c, 0x17, 0x37, 0x33, 0xdb, 0x9d, 0x4d, 0x6c, 0x28, 0xb9, 0x78, 0xf3, 0x52,
	0x04, 0xef, 0x9e, 0x7c, 0x09, 0xbe, 0x07, 0xbd, 0x88, 0x05, 0x2f, 0x1e, 0x25, 0xf1, 0x85, 0x48,
	0x66, 0x67, 0x12, 0x8d, 0x3b, 0x49, 0x9a, 0xdb, 0x92, 0xf9, 0x7e, 0x9f, 0xe7, 0xf3, 0x9d, 0xe7,
	0x61, 0x82, 0xcc, 0x26, 0xe1, 0x6d, 0xc6, 0x9d, 0x20, 0x64, 0x2d, 0xcf, 0x27, 0xdc, 0xe9, 0x96,
	0x9d, 0xe3, 0x0e, 0x09, 0x7b, 0x76, 0x10, 0xb2, 0x88, 0x01, 0xc4, 0xe7, 0xb6, 0x3a, 0xb7, 0xbb,
	0xe5, 0x54, 0xda, 0x65, 0xcc, 0xf5, 0x89, 0x83, 0x03, 0xcf, 0xc1, 0x94, 0xb2, 0x08, 0x47, 0x1e,
	0xa3, 0x3c, 0x76, 0xa4, 0xb2, 0xba, 0x8a, 0x35, 0xf9, 0x8b, 0xd4, 0x15, 0xb5, 0xba, 0x66, 0x84,
	0xdd, 0x5a, 0x48, 0x8e, 0x3b, 0x84, 0x47, 0xaa, 0x6a, 0x46, 0x5f, 0x15, 0x87, 0xb8, 0xad, 0x64,
	0x05, 0xad, 0xac, 0x71, 0x84, 0x3d, 0x5a, 0xf3, 0x3d, 0xfa, 0x4a, 0x69, 0xf3, 0x5a, 0x2d, 0x0e,
	0x82, 0x3f, 0x95, 0x5b, 0x67, 0x97, 0xd1, 0xff, 0xcf, 0x47, 0x27, 0xf0, 0xd6, 0x40, 0x97, 0x0e,
	0x62, 0x3d, 0xe4, 0xec, 0x7f, 0xef, 0xc6, 0x16, 0x32, 0xa9, 0x78, 0x11, 0x47, 0x48, 0xe5, 0xe7,
	0x0b, 0x79, 0xc0, 0x28, 0x27, 0xeb, 0x1b, 0x6f, 0xbe, 0xff, 0x7a, 0xff, 0x5f, 0x06, 0x6e, 0x3b,
	0x09, 0x6c, 0xe3, 0xef, 0xd3, 0x0e, 0x27, 0x61, 0x1f, 0xbe, 0x19, 0x28, 0xbd, 0x4f, 0x1b, 0xac,
	0xed, 0x51, 0xb7, 0x7a, 0x88, 0xdd, 0xc3, 0x10, 0x53, 0xde, 0x22, 0xa1, 0x6c, 0xcb, 0xe1, 0xa1,
	0xb6, 0xef, 0x2c, 0x9b, 0xa2, 0x7e, 0xb4, 0xa4, 0x5b, 0x46, 0xd9, 0x12, 0x51, 0x8a, 0x50, 0x48,
	0x8a, 0x32, 0x9a, 0xb0, 0x15, 0x49, 0xab, 0xa5, 0x46, 0x0d, 0x67, 0x06, 0x42, 0x7b, 0xa3, 0x39,
	0x3d, 0x1d, 0x5d, 0x3e, 0x14, 0xb4, 0x04, 0x13, 0x91, 0xa2, 0xdd, 0x58, 0x48, 0x2b, 0xd9, 0x72,
	0x82, 0xed, 0x16, 0xac, 0x25, 0xb1, 0x89, 0x45, 0xb1, 0xc4, 0xf8, 0xe1, 0xa3, 0x81, 0xae, 0x8c,
	0xfd, 0xcf, 0x5e, 0x53, 0x12, 0x72, 0x70, 0xe6, 0x77, 0x8a, 0x95, 0x0a, 0x6d, 0x73, 0x71, 0x83,
	0xe4, 0xb3, 0x05, 0x5f, 0x1e, 0xb2, 0x73, 0xf8, 0x1c, 0x16, 0x23, 0x7d, 0x36, 0xd0, 0xf5, 0x2a,
	0x69, 0xe1, 0x8e, 0x1f, 0x3d, 0x3e, 0x89, 0x48, 0x48, 0xb1, 0xbf, 0xd3, 0x6c, 0x86, 0x84, 0x73,
	0xc2, 0xe1, 0xbe, 0xb6, 0xbd, 0xce, 0xa2, 0xc0, 0x2b, 0x4b, 0x38, 0x65, 0x82, 0x6d, 0x91, 0x60,
	0x13, 0xec, 0xc4, 0xe9, 0xc7, 0x6e, 0x8b, 0x48, 0xbb, 0x85, 0xc7, 0xb0, 0x1f, 0x0c, 0x74, 0x75,
	0x27, 0x08, 0x7c, 0xaf, 0x21, 0xde, 0x94, 0x78, 0x0f, 0xf4, 0x17, 0x38, 0x2d, 0x55, 0xe4, 0xa5,
	0x0b, 0x38, 0x24, 0x71, 0x46, 0x10, 0xaf, 0xc1, 0xcd, 0x24, 0x62, 0x1c, 0x04, 0x72, 0x23, 0xbe,
	0x1a, 0xe8, 0xc6, 0x54, 0x8d, 0xdd, 0xde, 0x9e, 0xef, 0x11, 0x1a, 0xed, 0x57, 0xa1, 0xb2, 0x68,
	0xdf, 0x89, 0x47, 0x21, 0x3f, 0x58, 0xc6, 0x2a, 0xd9, 0x2b, 0x82, 0xbd, 0x0c, 0xa5, 0x99, 0xec,
	0x4e, 0x43, 0xf8, 0xb8, 0x73, 0x1a, 0x7f, 0xd4, 0xbc, 0x66, 0x1f, 0x3e, 0x19, 0xe8, 0xda, 0x54,
	0x03, 0xb9, 0xe7, 0x77, 0x17, 0x05, 0xfa, 0x7b, 0xdb, 0xb7, 0x2f, 0x6a, 0x93, 0x19, 0x8a, 0x22,
	0x43, 0x16, 0xee, 0xcc, 0xce, 0x20, 0x37, 0xbe, 0x8f, 0x56, 0x0f, 0xc4, 0xbb, 0x0f, 0x59, 0xfd,
	0xe3, 0x2a, 0x04, 0x8a, 0x2b, 0x37, 0x57, 0x27, 0x41, 0xd6, 0x05, 0x48, 0x1a, 0x52, 0x89, 0x6f,
	0xb0, 0xd0, 0xee, 0x3e, 0xf9, 0x32, 0x30, 0x8d, 0xf3, 0x81, 0x69, 0xfc, 0x1c, 0x98, 0xc6, 0xbb,
	0xa1, 0xb9, 0x72, 0x3e, 0x34, 0x57, 0x7e, 0x0c, 0xcd, 0x95, 0x97, 0x25, 0xd7, 0x8b, 0x8e, 0x3a,
	0x75, 0xbb, 0xc1, 0xda, 0xd2, 0x6f, 0xf9, 0xb8, 0xce, 0x55, 0xad, 0xee, 0x3d, 0xe7, 0x64, 0x52,
	0x30, 0xea, 0x05, 0x84, 0xd7, 0x57, 0xc5, 0x9f, 0x4c, 0xf9, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x24, 0x12, 0x6d, 0x26, 0x8b, 0x07, 0x00, 0x00,
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
	// DefaultExternalAddresses queries the default addresses associated to the
	// given user and (optionally) chain name
	DefaultExternalAddresses(ctx context.Context, in *QueryDefaultExternalAddressesRequest, opts ...grpc.CallOption) (*QueryDefaultExternalAddressesResponse, error)
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

func (c *queryClient) ChainLinks(ctx context.Context, in *QueryChainLinksRequest, opts ...grpc.CallOption) (*QueryChainLinksResponse, error) {
	out := new(QueryChainLinksResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/ChainLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ChainLinkOwners(ctx context.Context, in *QueryChainLinkOwnersRequest, opts ...grpc.CallOption) (*QueryChainLinkOwnersResponse, error) {
	out := new(QueryChainLinkOwnersResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/ChainLinkOwners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DefaultExternalAddresses(ctx context.Context, in *QueryDefaultExternalAddressesRequest, opts ...grpc.CallOption) (*QueryDefaultExternalAddressesResponse, error) {
	out := new(QueryDefaultExternalAddressesResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/DefaultExternalAddresses", in, out, opts...)
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

func (c *queryClient) ApplicationLinkOwners(ctx context.Context, in *QueryApplicationLinkOwnersRequest, opts ...grpc.CallOption) (*QueryApplicationLinkOwnersResponse, error) {
	out := new(QueryApplicationLinkOwnersResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Query/ApplicationLinkOwners", in, out, opts...)
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
	// DefaultExternalAddresses queries the default addresses associated to the
	// given user and (optionally) chain name
	DefaultExternalAddresses(context.Context, *QueryDefaultExternalAddressesRequest) (*QueryDefaultExternalAddressesResponse, error)
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
func (*UnimplementedQueryServer) DefaultExternalAddresses(ctx context.Context, req *QueryDefaultExternalAddressesRequest) (*QueryDefaultExternalAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DefaultExternalAddresses not implemented")
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
		FullMethod: "/desmos.profiles.v3.Query/ChainLinkOwners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ChainLinkOwners(ctx, req.(*QueryChainLinkOwnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DefaultExternalAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDefaultExternalAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DefaultExternalAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Query/DefaultExternalAddresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DefaultExternalAddresses(ctx, req.(*QueryDefaultExternalAddressesRequest))
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
		FullMethod: "/desmos.profiles.v3.Query/ApplicationLinkOwners",
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
		FullMethod: "/desmos.profiles.v3.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
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
			MethodName: "ChainLinks",
			Handler:    _Query_ChainLinks_Handler,
		},
		{
			MethodName: "ChainLinkOwners",
			Handler:    _Query_ChainLinkOwners_Handler,
		},
		{
			MethodName: "DefaultExternalAddresses",
			Handler:    _Query_DefaultExternalAddresses_Handler,
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
	Metadata: "desmos/profiles/v3/query.proto",
}
