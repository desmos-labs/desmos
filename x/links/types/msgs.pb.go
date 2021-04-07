// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/links/v1beta1/msgs.proto

package types

import (
	context "context"
	fmt "fmt"
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

// MsgCreateIBCAccountConnection represents the message to be used to create a
// link.
type MsgCreateIBCAccountConnection struct {
	Port                 string `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty" yaml:"port"`
	ChannelId            string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	TimeoutTimestamp     uint64 `protobuf:"varint,3,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	SourceChainPrefix    string `protobuf:"bytes,4,opt,name=source_chain_prefix,json=sourceChainPrefix,proto3" json:"source_chain_prefix,omitempty" yaml:"source_chain_prefix"`
	SourceAddress        string `protobuf:"bytes,5,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty" yaml:"source_address"`
	SourcePubKey         string `protobuf:"bytes,6,opt,name=source_pub_key,json=sourcePubKey,proto3" json:"source_pub_key,omitempty" yaml:"source_pub_key"`
	DestinationAddress   string `protobuf:"bytes,7,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty" yaml:"destination_address"`
	SourceSignature      string `protobuf:"bytes,8,opt,name=source_signature,json=sourceSignature,proto3" json:"source_signature,omitempty" yaml:"source_signature"`
	DestinationSignature string `protobuf:"bytes,9,opt,name=destination_signature,json=destinationSignature,proto3" json:"destination_signature,omitempty" yaml:"destination_signature"`
}

func (m *MsgCreateIBCAccountConnection) Reset()         { *m = MsgCreateIBCAccountConnection{} }
func (m *MsgCreateIBCAccountConnection) String() string { return proto.CompactTextString(m) }
func (*MsgCreateIBCAccountConnection) ProtoMessage()    {}
func (*MsgCreateIBCAccountConnection) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aac42b2e94dabba, []int{0}
}
func (m *MsgCreateIBCAccountConnection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateIBCAccountConnection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateIBCAccountConnection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateIBCAccountConnection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateIBCAccountConnection.Merge(m, src)
}
func (m *MsgCreateIBCAccountConnection) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateIBCAccountConnection) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateIBCAccountConnection.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateIBCAccountConnection proto.InternalMessageInfo

// MsgCreateIBCAccountConnectionResponse defines the Msg/CreatePost response
// type.
type MsgCreateIBCAccountConnectionResponse struct {
}

func (m *MsgCreateIBCAccountConnectionResponse) Reset()         { *m = MsgCreateIBCAccountConnectionResponse{} }
func (m *MsgCreateIBCAccountConnectionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateIBCAccountConnectionResponse) ProtoMessage()    {}
func (*MsgCreateIBCAccountConnectionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0aac42b2e94dabba, []int{1}
}
func (m *MsgCreateIBCAccountConnectionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateIBCAccountConnectionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateIBCAccountConnectionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateIBCAccountConnectionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateIBCAccountConnectionResponse.Merge(m, src)
}
func (m *MsgCreateIBCAccountConnectionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateIBCAccountConnectionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateIBCAccountConnectionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateIBCAccountConnectionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateIBCAccountConnection)(nil), "desmos.links.v1beta1.MsgCreateIBCAccountConnection")
	proto.RegisterType((*MsgCreateIBCAccountConnectionResponse)(nil), "desmos.links.v1beta1.MsgCreateIBCAccountConnectionResponse")
}

func init() { proto.RegisterFile("desmos/links/v1beta1/msgs.proto", fileDescriptor_0aac42b2e94dabba) }

var fileDescriptor_0aac42b2e94dabba = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0x87, 0x1b, 0xd6, 0x8d, 0xd5, 0xfc, 0xd9, 0xea, 0xb5, 0x22, 0x94, 0x91, 0x54, 0x46, 0x88,
	0x49, 0x88, 0x44, 0x63, 0x9c, 0xc6, 0x01, 0xd6, 0x0a, 0xa4, 0x0a, 0x0d, 0xa6, 0x00, 0x17, 0x2e,
	0x95, 0x93, 0x98, 0x34, 0x5a, 0x63, 0x47, 0xb1, 0x83, 0xd6, 0x6f, 0xc0, 0x09, 0x21, 0x0e, 0x9c,
	0xf7, 0x71, 0x38, 0xee, 0xc8, 0x29, 0x42, 0xed, 0x85, 0x73, 0x3e, 0x01, 0x8a, 0x9d, 0x76, 0x1d,
	0x2b, 0x3d, 0x70, 0xaa, 0xfd, 0xbc, 0x8f, 0x7f, 0xef, 0x5b, 0xc5, 0x06, 0xa6, 0x4f, 0x78, 0xc4,
	0xb8, 0x3d, 0x0c, 0xe9, 0x31, 0xb7, 0x3f, 0xed, 0xba, 0x44, 0xe0, 0x5d, 0x3b, 0xe2, 0x01, 0xb7,
	0xe2, 0x84, 0x09, 0x06, 0x1b, 0x4a, 0xb0, 0xa4, 0x60, 0x95, 0x42, 0xab, 0x11, 0xb0, 0x80, 0x49,
	0xc1, 0x2e, 0x56, 0xca, 0x45, 0xdf, 0x56, 0xc1, 0xdd, 0x43, 0x1e, 0x74, 0x13, 0x82, 0x05, 0xe9,
	0x75, 0xba, 0x07, 0x9e, 0xc7, 0x52, 0x2a, 0xba, 0x8c, 0x52, 0xe2, 0x89, 0x90, 0x51, 0x78, 0x0f,
	0x54, 0x63, 0x96, 0x08, 0x5d, 0x6b, 0x6b, 0x3b, 0xb5, 0xce, 0x46, 0x9e, 0x99, 0xd7, 0x46, 0x38,
	0x1a, 0xee, 0xa3, 0x82, 0x22, 0x47, 0x16, 0xe1, 0x13, 0x00, 0xbc, 0x01, 0xa6, 0x94, 0x0c, 0xfb,
	0xa1, 0xaf, 0x5f, 0x91, 0x6a, 0x33, 0xcf, 0xcc, 0xba, 0x52, 0xcf, 0x6b, 0xc8, 0xa9, 0x95, 0x9b,
	0x9e, 0x0f, 0x7b, 0xa0, 0x2e, 0xc2, 0x88, 0xb0, 0x54, 0xf4, 0x8b, 0x5f, 0x2e, 0x70, 0x14, 0xeb,
	0x2b, 0x6d, 0x6d, 0xa7, 0xda, 0xd9, 0xce, 0x33, 0x53, 0x57, 0x87, 0x2f, 0x29, 0xc8, 0xd9, 0x2c,
	0xd9, 0xbb, 0x29, 0x82, 0xaf, 0xc1, 0x16, 0x67, 0x69, 0xe2, 0x91, 0xbe, 0x37, 0xc0, 0x21, 0xed,
	0xc7, 0x09, 0xf9, 0x18, 0x9e, 0xe8, 0x55, 0x39, 0x89, 0x91, 0x67, 0x66, 0x4b, 0x85, 0x2d, 0x90,
	0x90, 0x53, 0x57, 0xb4, 0x5b, 0xc0, 0x23, 0xc9, 0xe0, 0x73, 0x70, 0xb3, 0x54, 0xb1, 0xef, 0x27,
	0x84, 0x73, 0x7d, 0x55, 0x46, 0xdd, 0xce, 0x33, 0xb3, 0x79, 0x21, 0xaa, 0xac, 0x23, 0xe7, 0x86,
	0x02, 0x07, 0x6a, 0x0f, 0x9f, 0xcd, 0x12, 0xe2, 0xd4, 0xed, 0x1f, 0x93, 0x91, 0xbe, 0xf6, 0x8f,
	0x84, 0xb2, 0x8e, 0x9c, 0xeb, 0x0a, 0x1c, 0xa5, 0xee, 0x2b, 0x32, 0x82, 0x6f, 0xc0, 0x96, 0x4f,
	0xb8, 0x08, 0x29, 0x2e, 0xbe, 0xc3, 0x6c, 0x8e, 0xab, 0x7f, 0xff, 0xa5, 0x05, 0x12, 0x72, 0xe0,
	0x1c, 0x9d, 0x4e, 0xf4, 0x12, 0x6c, 0x96, 0x1d, 0x79, 0x18, 0x50, 0x2c, 0xd2, 0x84, 0xe8, 0xeb,
	0x32, 0xed, 0x4e, 0x9e, 0x99, 0xb7, 0x2e, 0xcc, 0x34, 0x33, 0x90, 0xb3, 0xa1, 0xd0, 0xdb, 0x29,
	0x81, 0xef, 0x41, 0x73, 0xbe, 0xe7, 0x79, 0x58, 0x4d, 0x86, 0xb5, 0xf3, 0xcc, 0xdc, 0xbe, 0x3c,
	0xda, 0x5c, 0x62, 0x63, 0x8e, 0xcf, 0x62, 0xf7, 0xd7, 0x3f, 0x9f, 0x9a, 0x95, 0xdf, 0xa7, 0x66,
	0x05, 0x3d, 0x00, 0xf7, 0x97, 0xde, 0x49, 0x87, 0xf0, 0x98, 0x51, 0x4e, 0x1e, 0x7f, 0xd7, 0xc0,
	0xca, 0x21, 0x0f, 0xe0, 0x17, 0x0d, 0xb4, 0x96, 0x5c, 0xe1, 0x3d, 0x6b, 0xd1, 0x8b, 0xb0, 0x96,
	0xf6, 0x68, 0x3d, 0xfd, 0x8f, 0x43, 0xd3, 0xc1, 0x3a, 0x2f, 0x7e, 0x8c, 0x0d, 0xed, 0x6c, 0x6c,
	0x68, 0xbf, 0xc6, 0x86, 0xf6, 0x75, 0x62, 0x54, 0xce, 0x26, 0x46, 0xe5, 0xe7, 0xc4, 0xa8, 0x7c,
	0x78, 0x18, 0x84, 0x62, 0x90, 0xba, 0x96, 0xc7, 0x22, 0x5b, 0x35, 0x78, 0x34, 0xc4, 0x2e, 0x2f,
	0xd7, 0xf6, 0x49, 0xf9, 0xac, 0xc5, 0x28, 0x26, 0xdc, 0x5d, 0x93, 0x8f, 0x74, 0xef, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x24, 0xfd, 0x6d, 0x9f, 0xf3, 0x03, 0x00, 0x00,
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
	// CreateIBCAccountConnection defines the method to create a post
	CreateIBCAccountConnection(ctx context.Context, in *MsgCreateIBCAccountConnection, opts ...grpc.CallOption) (*MsgCreateIBCAccountConnectionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateIBCAccountConnection(ctx context.Context, in *MsgCreateIBCAccountConnection, opts ...grpc.CallOption) (*MsgCreateIBCAccountConnectionResponse, error) {
	out := new(MsgCreateIBCAccountConnectionResponse)
	err := c.cc.Invoke(ctx, "/desmos.links.v1beta1.Msg/CreateIBCAccountConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateIBCAccountConnection defines the method to create a post
	CreateIBCAccountConnection(context.Context, *MsgCreateIBCAccountConnection) (*MsgCreateIBCAccountConnectionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateIBCAccountConnection(ctx context.Context, req *MsgCreateIBCAccountConnection) (*MsgCreateIBCAccountConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIBCAccountConnection not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateIBCAccountConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateIBCAccountConnection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateIBCAccountConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desmos.links.v1beta1.Msg/CreateIBCAccountConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateIBCAccountConnection(ctx, req.(*MsgCreateIBCAccountConnection))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "desmos.links.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIBCAccountConnection",
			Handler:    _Msg_CreateIBCAccountConnection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "desmos/links/v1beta1/msgs.proto",
}

func (m *MsgCreateIBCAccountConnection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateIBCAccountConnection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateIBCAccountConnection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DestinationSignature) > 0 {
		i -= len(m.DestinationSignature)
		copy(dAtA[i:], m.DestinationSignature)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.DestinationSignature)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.SourceSignature) > 0 {
		i -= len(m.SourceSignature)
		copy(dAtA[i:], m.SourceSignature)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourceSignature)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.SourcePubKey) > 0 {
		i -= len(m.SourcePubKey)
		copy(dAtA[i:], m.SourcePubKey)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourcePubKey)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SourceChainPrefix) > 0 {
		i -= len(m.SourceChainPrefix)
		copy(dAtA[i:], m.SourceChainPrefix)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.SourceChainPrefix)))
		i--
		dAtA[i] = 0x22
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintMsgs(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Port) > 0 {
		i -= len(m.Port)
		copy(dAtA[i:], m.Port)
		i = encodeVarintMsgs(dAtA, i, uint64(len(m.Port)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateIBCAccountConnectionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateIBCAccountConnectionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateIBCAccountConnectionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgs(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateIBCAccountConnection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Port)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovMsgs(uint64(m.TimeoutTimestamp))
	}
	l = len(m.SourceChainPrefix)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.SourcePubKey)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.SourceSignature)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	l = len(m.DestinationSignature)
	if l > 0 {
		n += 1 + l + sovMsgs(uint64(l))
	}
	return n
}

func (m *MsgCreateIBCAccountConnectionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgs(x uint64) (n int) {
	return sovMsgs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateIBCAccountConnection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgCreateIBCAccountConnection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateIBCAccountConnection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Port = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChainPrefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceChainPrefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourcePubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgs
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
				return ErrInvalidLengthMsgs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func (m *MsgCreateIBCAccountConnectionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgs
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
			return fmt.Errorf("proto: MsgCreateIBCAccountConnectionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateIBCAccountConnectionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgs
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
func skipMsgs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
					return 0, ErrIntOverflowMsgs
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
				return 0, ErrInvalidLengthMsgs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgs = fmt.Errorf("proto: unexpected end of group")
)
