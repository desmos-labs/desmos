// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/msg_server.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

func init() {
	proto.RegisterFile("desmos/profiles/v3/msg_server.proto", fileDescriptor_869194438e2ecf0d)
}

var fileDescriptor_869194438e2ecf0d = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0x5b, 0x21, 0x21, 0x61, 0x84, 0x00, 0xc3, 0x65, 0x3e, 0xe4, 0xc0, 0xdb, 0x18, 0xb0,
	0x78, 0x6b, 0x2f, 0x5c, 0xcb, 0x76, 0x83, 0x49, 0xa8, 0x8c, 0x0b, 0x12, 0xaa, 0x5c, 0xf7, 0xa9,
	0x17, 0xcd, 0xb5, 0x4d, 0xec, 0x44, 0xec, 0x13, 0x70, 0xe5, 0x03, 0xf0, 0x81, 0x38, 0xee, 0xc8,
	0x11, 0xb5, 0x5f, 0x04, 0x25, 0x6e, 0x43, 0xfa, 0xe2, 0x2e, 0xbd, 0xd9, 0x79, 0x7e, 0xff, 0x17,
	0x5b, 0x8a, 0xd1, 0xd3, 0x11, 0xd8, 0x89, 0xb6, 0xd4, 0xa4, 0x7a, 0x9c, 0x48, 0xb0, 0x34, 0xef,
	0xd2, 0x89, 0x15, 0x03, 0x0b, 0x69, 0x0e, 0x69, 0x6c, 0x52, 0xed, 0x34, 0xc6, 0x1e, 0x8a, 0x17,
	0x50, 0x9c, 0x77, 0xc9, 0x63, 0xa1, 0x85, 0x2e, 0xc7, 0xb4, 0x58, 0x79, 0x92, 0xec, 0x09, 0xad,
	0x85, 0x04, 0x5a, 0xee, 0x86, 0xd9, 0x98, 0x32, 0x75, 0xb5, 0x18, 0x71, 0x5d, 0x98, 0x0c, 0xbc,
	0xc6, 0x6f, 0xe6, 0xa3, 0xfd, 0x4d, 0x25, 0xf4, 0x08, 0x64, 0x49, 0x17, 0x9f, 0xe6, 0xe0, 0x61,
	0x18, 0x1c, 0x39, 0x26, 0x06, 0x29, 0x7c, 0xcb, 0xc0, 0xba, 0x85, 0xef, 0xf3, 0xcd, 0x87, 0x5b,
	0x75, 0x7d, 0x1d, 0xc2, 0x36, 0x79, 0x1e, 0x84, 0x60, 0x7e, 0xc1, 0x12, 0x35, 0x90, 0x89, 0xba,
	0xdc, 0x7a, 0xac, 0x02, 0x65, 0xc6, 0xd4, 0xc1, 0xce, 0xaf, 0x3b, 0xe8, 0xd6, 0x99, 0x15, 0xf8,
	0x2b, 0xba, 0xfb, 0x89, 0xe5, 0xf0, 0xd1, 0xf3, 0xf8, 0x49, 0xbc, 0x7e, 0xef, 0xf1, 0x99, 0x15,
	0x35, 0x86, 0xbc, 0xba, 0x99, 0xe9, 0x83, 0x35, 0x5a, 0x59, 0xc0, 0x1c, 0xdd, 0x3b, 0x05, 0x09,
	0xae, 0x0a, 0x78, 0x16, 0x10, 0x2f, 0x51, 0xe4, 0x4d, 0x13, 0xaa, 0x0a, 0xc9, 0xd0, 0xa3, 0xbe,
	0xbf, 0xb1, 0xd3, 0x73, 0x26, 0xce, 0x53, 0xa6, 0xec, 0x18, 0x52, 0x1c, 0xea, 0xb9, 0x81, 0x25,
	0x9d, 0xe6, 0x6c, 0x15, 0xfb, 0xa3, 0x8d, 0xf6, 0x4e, 0x98, 0xe2, 0x20, 0x97, 0xc7, 0xa5, 0x02,
	0x1f, 0x05, 0x1c, 0x83, 0x0a, 0xf2, 0x76, 0x57, 0xc5, 0x52, 0x93, 0x1e, 0xe7, 0x60, 0xdc, 0x2e,
	0x4d, 0x82, 0x8a, 0x60, 0x93, 0xa0, 0x62, 0xa9, 0x49, 0x1f, 0xc6, 0x99, 0x85, 0x5d, 0x9a, 0x04,
	0x15, 0xc1, 0x26, 0x41, 0x45, 0xd5, 0x44, 0xa2, 0x07, 0x1f, 0x12, 0x75, 0x79, 0x52, 0xfc, 0x22,
	0x3d, 0xce, 0x75, 0xa6, 0x1c, 0xde, 0x0f, 0xb8, 0xad, 0x82, 0x84, 0x36, 0x04, 0xab, 0xb4, 0x14,
	0xe1, 0xcf, 0x4a, 0xae, 0xe6, 0x1d, 0x04, 0x6c, 0xd6, 0x51, 0x72, 0xdc, 0x18, 0xad, 0x32, 0x13,
	0x74, 0xbf, 0xe8, 0xd3, 0x33, 0x46, 0x26, 0x9c, 0xb9, 0x44, 0x2b, 0xfc, 0x62, 0x4b, 0xef, 0x1a,
	0x47, 0xe2, 0x66, 0x5c, 0x15, 0xa5, 0xd1, 0x43, 0x5f, 0xa4, 0x1e, 0xf6, 0x72, 0x6b, 0xe5, 0x7a,
	0xdc, 0x51, 0x53, 0x72, 0x11, 0xf8, 0xee, 0xfd, 0xef, 0x69, 0xd4, 0xbe, 0x9e, 0x46, 0xed, 0xbf,
	0xd3, 0xa8, 0xfd, 0x73, 0x16, 0xb5, 0xae, 0x67, 0x51, 0xeb, 0xcf, 0x2c, 0x6a, 0x7d, 0x39, 0x16,
	0x89, 0xbb, 0xc8, 0x86, 0x31, 0xd7, 0x13, 0xea, 0x5d, 0x0f, 0x25, 0x1b, 0xda, 0xf9, 0x9a, 0xe6,
	0x1d, 0xfa, 0xfd, 0xff, 0xeb, 0xe7, 0xae, 0x0c, 0xd8, 0xe1, 0xed, 0xf2, 0xc9, 0xeb, 0xfe, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0x36, 0xcc, 0xa9, 0x79, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SaveProfile defines the method to save a profile
	SaveProfile(ctx context.Context, in *MsgSaveProfile, opts ...grpc.CallOption) (*MsgSaveProfileResponse, error)
	// DeleteProfile defines the method to delete an existing profile
	DeleteProfile(ctx context.Context, in *MsgDeleteProfile, opts ...grpc.CallOption) (*MsgDeleteProfileResponse, error)
	// RequestDTagTransfer defines the method to request another user to transfer
	// their DTag to you
	RequestDTagTransfer(ctx context.Context, in *MsgRequestDTagTransfer, opts ...grpc.CallOption) (*MsgRequestDTagTransferResponse, error)
	// CancelDTagTransferRequest defines the method to cancel an outgoing DTag
	// transfer request
	CancelDTagTransferRequest(ctx context.Context, in *MsgCancelDTagTransferRequest, opts ...grpc.CallOption) (*MsgCancelDTagTransferRequestResponse, error)
	// AcceptDTagTransferRequest defines the method to accept an incoming DTag
	// transfer request
	AcceptDTagTransferRequest(ctx context.Context, in *MsgAcceptDTagTransferRequest, opts ...grpc.CallOption) (*MsgAcceptDTagTransferRequestResponse, error)
	// RefuseDTagTransferRequest defines the method to refuse an incoming DTag
	// transfer request
	RefuseDTagTransferRequest(ctx context.Context, in *MsgRefuseDTagTransferRequest, opts ...grpc.CallOption) (*MsgRefuseDTagTransferRequestResponse, error)
	// LinkChainAccount defines a method to link an external chain account to a
	// profile
	LinkChainAccount(ctx context.Context, in *MsgLinkChainAccount, opts ...grpc.CallOption) (*MsgLinkChainAccountResponse, error)
	// UnlinkChainAccount defines a method to unlink an external chain account
	// from a profile
	UnlinkChainAccount(ctx context.Context, in *MsgUnlinkChainAccount, opts ...grpc.CallOption) (*MsgUnlinkChainAccountResponse, error)
	// LinkApplication defines a method to create a centralized application
	// link
	LinkApplication(ctx context.Context, in *MsgLinkApplication, opts ...grpc.CallOption) (*MsgLinkApplicationResponse, error)
	// UnlinkApplication defines a method to remove a centralized application
	UnlinkApplication(ctx context.Context, in *MsgUnlinkApplication, opts ...grpc.CallOption) (*MsgUnlinkApplicationResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SaveProfile(ctx context.Context, in *MsgSaveProfile, opts ...grpc.CallOption) (*MsgSaveProfileResponse, error) {
	out := new(MsgSaveProfileResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/SaveProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteProfile(ctx context.Context, in *MsgDeleteProfile, opts ...grpc.CallOption) (*MsgDeleteProfileResponse, error) {
	out := new(MsgDeleteProfileResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/DeleteProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestDTagTransfer(ctx context.Context, in *MsgRequestDTagTransfer, opts ...grpc.CallOption) (*MsgRequestDTagTransferResponse, error) {
	out := new(MsgRequestDTagTransferResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/RequestDTagTransfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelDTagTransferRequest(ctx context.Context, in *MsgCancelDTagTransferRequest, opts ...grpc.CallOption) (*MsgCancelDTagTransferRequestResponse, error) {
	out := new(MsgCancelDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/CancelDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) AcceptDTagTransferRequest(ctx context.Context, in *MsgAcceptDTagTransferRequest, opts ...grpc.CallOption) (*MsgAcceptDTagTransferRequestResponse, error) {
	out := new(MsgAcceptDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/AcceptDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RefuseDTagTransferRequest(ctx context.Context, in *MsgRefuseDTagTransferRequest, opts ...grpc.CallOption) (*MsgRefuseDTagTransferRequestResponse, error) {
	out := new(MsgRefuseDTagTransferRequestResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/RefuseDTagTransferRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LinkChainAccount(ctx context.Context, in *MsgLinkChainAccount, opts ...grpc.CallOption) (*MsgLinkChainAccountResponse, error) {
	out := new(MsgLinkChainAccountResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/LinkChainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnlinkChainAccount(ctx context.Context, in *MsgUnlinkChainAccount, opts ...grpc.CallOption) (*MsgUnlinkChainAccountResponse, error) {
	out := new(MsgUnlinkChainAccountResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/UnlinkChainAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LinkApplication(ctx context.Context, in *MsgLinkApplication, opts ...grpc.CallOption) (*MsgLinkApplicationResponse, error) {
	out := new(MsgLinkApplicationResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/LinkApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnlinkApplication(ctx context.Context, in *MsgUnlinkApplication, opts ...grpc.CallOption) (*MsgUnlinkApplicationResponse, error) {
	out := new(MsgUnlinkApplicationResponse)
	err := c.cc.Invoke(ctx, "/desmos.profiles.v3.Msg/UnlinkApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SaveProfile defines the method to save a profile
	SaveProfile(context.Context, *MsgSaveProfile) (*MsgSaveProfileResponse, error)
	// DeleteProfile defines the method to delete an existing profile
	DeleteProfile(context.Context, *MsgDeleteProfile) (*MsgDeleteProfileResponse, error)
	// RequestDTagTransfer defines the method to request another user to transfer
	// their DTag to you
	RequestDTagTransfer(context.Context, *MsgRequestDTagTransfer) (*MsgRequestDTagTransferResponse, error)
	// CancelDTagTransferRequest defines the method to cancel an outgoing DTag
	// transfer request
	CancelDTagTransferRequest(context.Context, *MsgCancelDTagTransferRequest) (*MsgCancelDTagTransferRequestResponse, error)
	// AcceptDTagTransferRequest defines the method to accept an incoming DTag
	// transfer request
	AcceptDTagTransferRequest(context.Context, *MsgAcceptDTagTransferRequest) (*MsgAcceptDTagTransferRequestResponse, error)
	// RefuseDTagTransferRequest defines the method to refuse an incoming DTag
	// transfer request
	RefuseDTagTransferRequest(context.Context, *MsgRefuseDTagTransferRequest) (*MsgRefuseDTagTransferRequestResponse, error)
	// LinkChainAccount defines a method to link an external chain account to a
	// profile
	LinkChainAccount(context.Context, *MsgLinkChainAccount) (*MsgLinkChainAccountResponse, error)
	// UnlinkChainAccount defines a method to unlink an external chain account
	// from a profile
	UnlinkChainAccount(context.Context, *MsgUnlinkChainAccount) (*MsgUnlinkChainAccountResponse, error)
	// LinkApplication defines a method to create a centralized application
	// link
	LinkApplication(context.Context, *MsgLinkApplication) (*MsgLinkApplicationResponse, error)
	// UnlinkApplication defines a method to remove a centralized application
	UnlinkApplication(context.Context, *MsgUnlinkApplication) (*MsgUnlinkApplicationResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SaveProfile(ctx context.Context, req *MsgSaveProfile) (*MsgSaveProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveProfile not implemented")
}
func (*UnimplementedMsgServer) DeleteProfile(ctx context.Context, req *MsgDeleteProfile) (*MsgDeleteProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfile not implemented")
}
func (*UnimplementedMsgServer) RequestDTagTransfer(ctx context.Context, req *MsgRequestDTagTransfer) (*MsgRequestDTagTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestDTagTransfer not implemented")
}
func (*UnimplementedMsgServer) CancelDTagTransferRequest(ctx context.Context, req *MsgCancelDTagTransferRequest) (*MsgCancelDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) AcceptDTagTransferRequest(ctx context.Context, req *MsgAcceptDTagTransferRequest) (*MsgAcceptDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) RefuseDTagTransferRequest(ctx context.Context, req *MsgRefuseDTagTransferRequest) (*MsgRefuseDTagTransferRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefuseDTagTransferRequest not implemented")
}
func (*UnimplementedMsgServer) LinkChainAccount(ctx context.Context, req *MsgLinkChainAccount) (*MsgLinkChainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkChainAccount not implemented")
}
func (*UnimplementedMsgServer) UnlinkChainAccount(ctx context.Context, req *MsgUnlinkChainAccount) (*MsgUnlinkChainAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkChainAccount not implemented")
}
func (*UnimplementedMsgServer) LinkApplication(ctx context.Context, req *MsgLinkApplication) (*MsgLinkApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LinkApplication not implemented")
}
func (*UnimplementedMsgServer) UnlinkApplication(ctx context.Context, req *MsgUnlinkApplication) (*MsgUnlinkApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnlinkApplication not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SaveProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSaveProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SaveProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/SaveProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SaveProfile(ctx, req.(*MsgSaveProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteProfile)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/DeleteProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteProfile(ctx, req.(*MsgDeleteProfile))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestDTagTransfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestDTagTransfer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestDTagTransfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/RequestDTagTransfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestDTagTransfer(ctx, req.(*MsgRequestDTagTransfer))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/CancelDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelDTagTransferRequest(ctx, req.(*MsgCancelDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_AcceptDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAcceptDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AcceptDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/AcceptDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AcceptDTagTransferRequest(ctx, req.(*MsgAcceptDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RefuseDTagTransferRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRefuseDTagTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RefuseDTagTransferRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/RefuseDTagTransferRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RefuseDTagTransferRequest(ctx, req.(*MsgRefuseDTagTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LinkChainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLinkChainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LinkChainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/LinkChainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LinkChainAccount(ctx, req.(*MsgLinkChainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnlinkChainAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnlinkChainAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnlinkChainAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/UnlinkChainAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnlinkChainAccount(ctx, req.(*MsgUnlinkChainAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LinkApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLinkApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LinkApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/LinkApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LinkApplication(ctx, req.(*MsgLinkApplication))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnlinkApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnlinkApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnlinkApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.profiles.v3.Msg/UnlinkApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnlinkApplication(ctx, req.(*MsgUnlinkApplication))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.profiles.v3.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveProfile",
			Handler:    _Msg_SaveProfile_Handler,
		},
		{
			MethodName: "DeleteProfile",
			Handler:    _Msg_DeleteProfile_Handler,
		},
		{
			MethodName: "RequestDTagTransfer",
			Handler:    _Msg_RequestDTagTransfer_Handler,
		},
		{
			MethodName: "CancelDTagTransferRequest",
			Handler:    _Msg_CancelDTagTransferRequest_Handler,
		},
		{
			MethodName: "AcceptDTagTransferRequest",
			Handler:    _Msg_AcceptDTagTransferRequest_Handler,
		},
		{
			MethodName: "RefuseDTagTransferRequest",
			Handler:    _Msg_RefuseDTagTransferRequest_Handler,
		},
		{
			MethodName: "LinkChainAccount",
			Handler:    _Msg_LinkChainAccount_Handler,
		},
		{
			MethodName: "UnlinkChainAccount",
			Handler:    _Msg_UnlinkChainAccount_Handler,
		},
		{
			MethodName: "LinkApplication",
			Handler:    _Msg_LinkApplication_Handler,
		},
		{
			MethodName: "UnlinkApplication",
			Handler:    _Msg_UnlinkApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/profiles/v3/msg_server.proto",
}
