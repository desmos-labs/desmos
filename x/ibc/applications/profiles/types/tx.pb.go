// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/ibc/applications/profiles/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// MsgCreateApplicationLink defines a msg to connect a profile with a centralized
// social network account (eg. Twitter, GitHub, etc).
type MsgCreateApplicationLink struct {
	// The sender of the connection request
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// Data of the application to which connect
	Application *ApplicationData `protobuf:"bytes,2,opt,name=application,proto3" json:"application,omitempty"`
	// Data used to verify the connection
	VerificationData *VerificationData `protobuf:"bytes,3,opt,name=verification_data,json=verificationData,proto3" json:"verification_data,omitempty"`
	// The port on which the packet will be sent
	SourcePort string `protobuf:"bytes,4,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// The channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,5,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,6,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,7,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
}

func (m *MsgCreateApplicationLink) Reset()         { *m = MsgCreateApplicationLink{} }
func (m *MsgCreateApplicationLink) String() string { return proto.CompactTextString(m) }
func (*MsgCreateApplicationLink) ProtoMessage()    {}
func (*MsgCreateApplicationLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a3cff6b0faea82f, []int{0}
}
func (m *MsgCreateApplicationLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateApplicationLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateApplicationLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateApplicationLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateApplicationLink.Merge(m, src)
}
func (m *MsgCreateApplicationLink) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateApplicationLink) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateApplicationLink.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateApplicationLink proto.InternalMessageInfo

// MsgCreateApplicationLinkResponse defines the Msg/ConnectProfile response type.
type MsgCreateApplicationLinkResponse struct {
}

func (m *MsgCreateApplicationLinkResponse) Reset()         { *m = MsgCreateApplicationLinkResponse{} }
func (m *MsgCreateApplicationLinkResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateApplicationLinkResponse) ProtoMessage()    {}
func (*MsgCreateApplicationLinkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a3cff6b0faea82f, []int{1}
}
func (m *MsgCreateApplicationLinkResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateApplicationLinkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateApplicationLinkResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateApplicationLinkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateApplicationLinkResponse.Merge(m, src)
}
func (m *MsgCreateApplicationLinkResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateApplicationLinkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateApplicationLinkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateApplicationLinkResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateApplicationLink)(nil), "desmos.ibc.applications.profiles.v1.MsgCreateApplicationLink")
	proto.RegisterType((*MsgCreateApplicationLinkResponse)(nil), "desmos.ibc.applications.profiles.v1.MsgCreateApplicationLinkResponse")
}

func init() {
	proto.RegisterFile("desmos/ibc/applications/profiles/v1/tx.proto", fileDescriptor_5a3cff6b0faea82f)
}

var fileDescriptor_5a3cff6b0faea82f = []byte{
	// 522 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0x4e, 0xec, 0xba, 0xea, 0x2c, 0x2d, 0xed, 0x60, 0x4b, 0x5c, 0x34, 0x59, 0xe2, 0x65, 0x0f,
	0x3a, 0x63, 0xaa, 0x22, 0x14, 0x44, 0xdd, 0x2a, 0x28, 0x58, 0x90, 0x20, 0x05, 0xbd, 0xac, 0x93,
	0xd9, 0x69, 0x76, 0x30, 0xc9, 0x84, 0xcc, 0x6c, 0x68, 0xff, 0x81, 0x47, 0xff, 0x80, 0xb0, 0x17,
	0xff, 0x4b, 0x8f, 0x3d, 0x7a, 0x5a, 0x64, 0xf7, 0xe2, 0x79, 0x7f, 0x81, 0x24, 0x99, 0xb5, 0x59,
	0xb1, 0xb2, 0xd0, 0x53, 0xde, 0xfb, 0xde, 0xf7, 0xbe, 0xef, 0x3d, 0x5e, 0x06, 0xdc, 0x1b, 0x30,
	0x19, 0x0b, 0x89, 0x79, 0x40, 0x31, 0x49, 0xd3, 0x88, 0x53, 0xa2, 0xb8, 0x48, 0x24, 0x4e, 0x33,
	0x71, 0xc4, 0x23, 0x26, 0x71, 0xee, 0x61, 0x75, 0x8c, 0xd2, 0x4c, 0x28, 0x01, 0xef, 0x56, 0x6c,
	0xc4, 0x03, 0x8a, 0xea, 0x6c, 0xb4, 0x60, 0xa3, 0xdc, 0x6b, 0x3f, 0x58, 0x45, 0x32, 0x16, 0x03,
	0x16, 0xc9, 0x4a, 0xb6, 0x7d, 0x33, 0x14, 0xa1, 0x28, 0x43, 0x5c, 0x44, 0x1a, 0xb5, 0xa9, 0x28,
	0x75, 0x02, 0x22, 0x19, 0xce, 0xbd, 0x80, 0x29, 0xe2, 0x61, 0x2a, 0x78, 0xa2, 0xeb, 0x4e, 0x61,
	0x40, 0x45, 0xc6, 0x30, 0x8d, 0x38, 0x4b, 0x54, 0x21, 0x5b, 0x45, 0x15, 0xc1, 0x1d, 0x37, 0x80,
	0x75, 0x20, 0xc3, 0xfd, 0x8c, 0x11, 0xc5, 0x5e, 0x9c, 0x4f, 0xf2, 0x96, 0x27, 0x9f, 0xe1, 0x0e,
	0x68, 0x4a, 0x96, 0x0c, 0x58, 0x66, 0x99, 0x1d, 0xb3, 0x7b, 0xc3, 0xd7, 0x19, 0x3c, 0x04, 0xad,
	0xda, 0xd0, 0xd6, 0x95, 0x8e, 0xd9, 0x6d, 0xed, 0x3e, 0x42, 0x2b, 0x2c, 0x8e, 0x6a, 0x16, 0x2f,
	0x89, 0x22, 0x7e, 0x5d, 0x08, 0x06, 0x60, 0x2b, 0x67, 0x19, 0x3f, 0xd2, 0x79, 0x7f, 0x40, 0x14,
	0xb1, 0xd6, 0x4a, 0xf5, 0xc7, 0x2b, 0xa9, 0x1f, 0xd6, 0xba, 0x4b, 0xf9, 0xcd, 0xfc, 0x2f, 0x04,
	0x3e, 0x01, 0x2d, 0x29, 0x46, 0x19, 0x65, 0xfd, 0x54, 0x64, 0xca, 0x6a, 0x14, 0x8b, 0xf5, 0x76,
	0xe6, 0x13, 0x07, 0x9e, 0x90, 0x38, 0xda, 0x73, 0x6b, 0x45, 0xd7, 0x07, 0x55, 0xf6, 0x4e, 0x64,
	0x0a, 0x3e, 0x07, 0x1b, 0xba, 0x46, 0x87, 0x24, 0x49, 0x58, 0x64, 0x5d, 0x2d, 0x7b, 0x6f, 0xcd,
	0x27, 0xce, 0xf6, 0x52, 0xaf, 0xae, 0xbb, 0xfe, 0x7a, 0x05, 0xec, 0x57, 0x39, 0xfc, 0x04, 0x36,
	0x14, 0x8f, 0x99, 0x18, 0xa9, 0xfe, 0x90, 0xf1, 0x70, 0xa8, 0xac, 0x66, 0xb9, 0x5b, 0xbb, 0x5c,
	0xaa, 0xb8, 0x12, 0xd2, 0xb7, 0xc9, 0x3d, 0xf4, 0xba, 0x64, 0xf4, 0xee, 0x9c, 0x4e, 0x1c, 0xe3,
	0xdc, 0x61, 0xb9, 0xdf, 0xf5, 0xd7, 0x35, 0x50, 0xb1, 0xe1, 0x1b, 0xb0, 0xb5, 0x60, 0x14, 0x5f,
	0xa9, 0x48, 0x9c, 0x5a, 0xd7, 0x3a, 0x66, 0xb7, 0xd1, 0xbb, 0x3d, 0x9f, 0x38, 0xd6, 0xb2, 0xc8,
	0x1f, 0x8a, 0xeb, 0x6f, 0x6a, 0xec, 0xfd, 0x02, 0xda, 0xbb, 0xfe, 0x65, 0xec, 0x18, 0xbf, 0xc6,
	0x8e, 0xe1, 0xba, 0xa0, 0x73, 0xd1, 0x1f, 0xe2, 0x33, 0x99, 0x8a, 0x44, 0xb2, 0xdd, 0xef, 0x26,
	0x58, 0x3b, 0x90, 0x21, 0xfc, 0x66, 0x82, 0xed, 0x7f, 0xff, 0x4b, 0x4f, 0x57, 0x3a, 0xe0, 0x45,
	0x46, 0xed, 0x57, 0x97, 0x6a, 0x5f, 0xcc, 0xd9, 0xfb, 0x70, 0x3a, 0xb5, 0xcd, 0xb3, 0xa9, 0x6d,
	0xfe, 0x9c, 0xda, 0xe6, 0xd7, 0x99, 0x6d, 0x9c, 0xcd, 0x6c, 0xe3, 0xc7, 0xcc, 0x36, 0x3e, 0x3e,
	0x0b, 0xb9, 0x1a, 0x8e, 0x02, 0x44, 0x45, 0x8c, 0x2b, 0xab, 0xfb, 0x11, 0x09, 0xa4, 0x8e, 0xf1,
	0xf1, 0x7f, 0x9e, 0xaa, 0x3a, 0x49, 0x99, 0x0c, 0x9a, 0xe5, 0x83, 0x7a, 0xf8, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x86, 0xcb, 0x7d, 0xe4, 0x2e, 0x04, 0x00, 0x00,
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
	// ConnectProfile defines a rpc handler method for MsgConnectProfile.
	CreateApplicationLink(ctx context.Context, in *MsgCreateApplicationLink, opts ...grpc.CallOption) (*MsgCreateApplicationLinkResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateApplicationLink(ctx context.Context, in *MsgCreateApplicationLink, opts ...grpc.CallOption) (*MsgCreateApplicationLinkResponse, error) {
	out := new(MsgCreateApplicationLinkResponse)
	err := c.cc.Invoke(ctx, "/desmos.ibc.applications.profiles.v1.Msg/CreateApplicationLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// ConnectProfile defines a rpc handler method for MsgConnectProfile.
	CreateApplicationLink(context.Context, *MsgCreateApplicationLink) (*MsgCreateApplicationLinkResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateApplicationLink(ctx context.Context, req *MsgCreateApplicationLink) (*MsgCreateApplicationLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApplicationLink not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateApplicationLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateApplicationLink)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateApplicationLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.ibc.applications.profiles.v1.Msg/CreateApplicationLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateApplicationLink(ctx, req.(*MsgCreateApplicationLink))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.ibc.applications.profiles.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateApplicationLink",
			Handler:    _Msg_CreateApplicationLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/ibc/applications/profiles/v1/tx.proto",
}

func (m *MsgCreateApplicationLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateApplicationLink) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateApplicationLink) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0x22
	}
	if m.VerificationData != nil {
		{
			size, err := m.VerificationData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Application != nil {
		{
			size, err := m.Application.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateApplicationLinkResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateApplicationLinkResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateApplicationLinkResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateApplicationLink) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Application != nil {
		l = m.Application.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if m.VerificationData != nil {
		l = m.VerificationData.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *MsgCreateApplicationLinkResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateApplicationLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateApplicationLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateApplicationLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Application", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Application == nil {
				m.Application = &ApplicationData{}
			}
			if err := m.Application.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerificationData == nil {
				m.VerificationData = &VerificationData{}
			}
			if err := m.VerificationData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateApplicationLinkResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateApplicationLinkResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateApplicationLinkResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
