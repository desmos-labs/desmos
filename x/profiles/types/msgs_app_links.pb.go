// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/msgs_app_links.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// MsgLinkApplication defines a msg to connect a profile with a
// centralized application account (eg. Twitter, GitHub, etc).
type MsgLinkApplication struct {
	// The sender of the connection request
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	// LinkData contains the data related to the application to which connect
	LinkData Data `protobuf:"bytes,2,opt,name=link_data,json=linkData,proto3" json:"link_data" yaml:"link_data"`
	// Hex encoded call data that will be sent to the data source in order to
	// verify the link
	CallData string `protobuf:"bytes,3,opt,name=call_data,json=callData,proto3" json:"call_data,omitempty" yaml:"call_data"`
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

func (m *MsgLinkApplication) Reset()         { *m = MsgLinkApplication{} }
func (m *MsgLinkApplication) String() string { return proto.CompactTextString(m) }
func (*MsgLinkApplication) ProtoMessage()    {}
func (*MsgLinkApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_29dfbdba444598ee, []int{0}
}
func (m *MsgLinkApplication) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkApplication.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkApplication.Merge(m, src)
}
func (m *MsgLinkApplication) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkApplication.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkApplication proto.InternalMessageInfo

func (m *MsgLinkApplication) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgLinkApplication) GetLinkData() Data {
	if m != nil {
		return m.LinkData
	}
	return Data{}
}

func (m *MsgLinkApplication) GetCallData() string {
	if m != nil {
		return m.CallData
	}
	return ""
}

func (m *MsgLinkApplication) GetSourcePort() string {
	if m != nil {
		return m.SourcePort
	}
	return ""
}

func (m *MsgLinkApplication) GetSourceChannel() string {
	if m != nil {
		return m.SourceChannel
	}
	return ""
}

func (m *MsgLinkApplication) GetTimeoutHeight() types.Height {
	if m != nil {
		return m.TimeoutHeight
	}
	return types.Height{}
}

func (m *MsgLinkApplication) GetTimeoutTimestamp() uint64 {
	if m != nil {
		return m.TimeoutTimestamp
	}
	return 0
}

// MsgLinkApplicationResponse defines the Msg/LinkApplication
// response type.
type MsgLinkApplicationResponse struct {
}

func (m *MsgLinkApplicationResponse) Reset()         { *m = MsgLinkApplicationResponse{} }
func (m *MsgLinkApplicationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgLinkApplicationResponse) ProtoMessage()    {}
func (*MsgLinkApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_29dfbdba444598ee, []int{1}
}
func (m *MsgLinkApplicationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgLinkApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgLinkApplicationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgLinkApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgLinkApplicationResponse.Merge(m, src)
}
func (m *MsgLinkApplicationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgLinkApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgLinkApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgLinkApplicationResponse proto.InternalMessageInfo

// MsgUnlinkApplication defines a msg to delete an application link from a user
// profile
type MsgUnlinkApplication struct {
	// Application represents the name of the application to unlink
	Application string `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty" yaml:"application"`
	// Username represents the username inside the application to unlink
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" yaml:"username"`
	// Signer represents the Desmos account to which the application should be
	// unlinked
	Signer string `protobuf:"bytes,3,opt,name=signer,proto3" json:"signer,omitempty" yaml:"signer"`
}

func (m *MsgUnlinkApplication) Reset()         { *m = MsgUnlinkApplication{} }
func (m *MsgUnlinkApplication) String() string { return proto.CompactTextString(m) }
func (*MsgUnlinkApplication) ProtoMessage()    {}
func (*MsgUnlinkApplication) Descriptor() ([]byte, []int) {
	return fileDescriptor_29dfbdba444598ee, []int{2}
}
func (m *MsgUnlinkApplication) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlinkApplication) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlinkApplication.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlinkApplication) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlinkApplication.Merge(m, src)
}
func (m *MsgUnlinkApplication) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlinkApplication) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlinkApplication.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlinkApplication proto.InternalMessageInfo

func (m *MsgUnlinkApplication) GetApplication() string {
	if m != nil {
		return m.Application
	}
	return ""
}

func (m *MsgUnlinkApplication) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MsgUnlinkApplication) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

// MsgUnlinkApplicationResponse defines the Msg/UnlinkApplication response
// type.
type MsgUnlinkApplicationResponse struct {
}

func (m *MsgUnlinkApplicationResponse) Reset()         { *m = MsgUnlinkApplicationResponse{} }
func (m *MsgUnlinkApplicationResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnlinkApplicationResponse) ProtoMessage()    {}
func (*MsgUnlinkApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_29dfbdba444598ee, []int{3}
}
func (m *MsgUnlinkApplicationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlinkApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlinkApplicationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlinkApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlinkApplicationResponse.Merge(m, src)
}
func (m *MsgUnlinkApplicationResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlinkApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlinkApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlinkApplicationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgLinkApplication)(nil), "desmos.profiles.v3.MsgLinkApplication")
	proto.RegisterType((*MsgLinkApplicationResponse)(nil), "desmos.profiles.v3.MsgLinkApplicationResponse")
	proto.RegisterType((*MsgUnlinkApplication)(nil), "desmos.profiles.v3.MsgUnlinkApplication")
	proto.RegisterType((*MsgUnlinkApplicationResponse)(nil), "desmos.profiles.v3.MsgUnlinkApplicationResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v3/msgs_app_links.proto", fileDescriptor_29dfbdba444598ee)
}

var fileDescriptor_29dfbdba444598ee = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcd, 0x8e, 0xd3, 0x3c,
	0x14, 0x6d, 0xbe, 0x99, 0x6f, 0x68, 0x5d, 0x75, 0x98, 0x09, 0x05, 0x85, 0xaa, 0x24, 0x95, 0x37,
	0x74, 0x16, 0x24, 0x2a, 0x45, 0x02, 0xb1, 0x82, 0xc0, 0x02, 0x04, 0x23, 0x50, 0x04, 0x1b, 0x36,
	0xc5, 0x4d, 0x4d, 0x6a, 0x8d, 0x63, 0x47, 0xb1, 0x5b, 0x31, 0x6f, 0xc1, 0x83, 0xf0, 0x20, 0xb3,
	0x9c, 0x25, 0xab, 0x08, 0xb5, 0x4b, 0x76, 0x79, 0x02, 0x14, 0xdb, 0xfd, 0xa3, 0xb3, 0xca, 0xf1,
	0xb9, 0xe7, 0x38, 0xd7, 0xd7, 0xc7, 0xe0, 0xe1, 0x04, 0x8b, 0x94, 0x8b, 0x20, 0xcb, 0xf9, 0x37,
	0x42, 0xb1, 0x08, 0xe6, 0xc3, 0x20, 0x15, 0x89, 0x18, 0xa1, 0x2c, 0x1b, 0x51, 0xc2, 0x2e, 0x84,
	0x9f, 0xe5, 0x5c, 0x72, 0xdb, 0xd6, 0x42, 0x7f, 0x25, 0xf4, 0xe7, 0xc3, 0xce, 0xd9, 0x4d, 0x66,
	0x3e, 0xc1, 0x74, 0xcf, 0xde, 0x69, 0x27, 0x3c, 0xe1, 0x0a, 0x06, 0x15, 0x32, 0xac, 0x47, 0xc6,
	0x71, 0x10, 0xf3, 0x1c, 0x07, 0x31, 0x25, 0x98, 0xc9, 0x60, 0x3e, 0x30, 0x48, 0x0b, 0xe0, 0x9f,
	0x03, 0x60, 0x9f, 0x8b, 0xe4, 0x3d, 0x61, 0x17, 0x2f, 0xb3, 0x8c, 0x92, 0x18, 0x49, 0xc2, 0x99,
	0x7d, 0x06, 0x8e, 0x04, 0x66, 0x13, 0x9c, 0x3b, 0x56, 0xcf, 0xea, 0x37, 0xc2, 0xd3, 0xb2, 0xf0,
	0x5a, 0x97, 0x28, 0xa5, 0xcf, 0xa1, 0xe6, 0x61, 0x64, 0x04, 0xf6, 0x07, 0xd0, 0xa8, 0xfa, 0x18,
	0x4d, 0x90, 0x44, 0xce, 0x7f, 0x3d, 0xab, 0xdf, 0x7c, 0xec, 0xf8, 0xfb, 0x67, 0xf1, 0x5f, 0x23,
	0x89, 0x42, 0xe7, 0xaa, 0xf0, 0x6a, 0x65, 0xe1, 0x9d, 0xe8, 0xbd, 0xd6, 0x46, 0x18, 0xd5, 0x2b,
	0x5c, 0x69, 0xec, 0x01, 0x68, 0xc4, 0x88, 0x52, 0xbd, 0xe1, 0x81, 0xfa, 0x7d, 0x7b, 0x63, 0x59,
	0x97, 0x60, 0x54, 0xaf, 0xb0, 0xb2, 0x3c, 0x05, 0x4d, 0xc1, 0x67, 0x79, 0x8c, 0x47, 0x19, 0xcf,
	0xa5, 0x73, 0xa8, 0x4c, 0xf7, 0xca, 0xc2, 0xb3, 0x4d, 0xcf, 0x9b, 0x22, 0x8c, 0x80, 0x5e, 0x7d,
	0xe4, 0xb9, 0xb4, 0x5f, 0x80, 0x63, 0x53, 0x8b, 0xa7, 0x88, 0x31, 0x4c, 0x9d, 0xff, 0x95, 0xf7,
	0x7e, 0x59, 0x78, 0x77, 0x77, 0xbc, 0xa6, 0x0e, 0xa3, 0x96, 0x26, 0x5e, 0xe9, 0xb5, 0xfd, 0x15,
	0x1c, 0x4b, 0x92, 0x62, 0x3e, 0x93, 0xa3, 0x29, 0x26, 0xc9, 0x54, 0x3a, 0x47, 0x6a, 0x06, 0x1d,
	0x9f, 0x8c, 0x63, 0xbf, 0x1a, 0xbd, 0x6f, 0x06, 0x3e, 0x1f, 0xf8, 0x6f, 0x94, 0x22, 0x7c, 0x60,
	0xa6, 0x60, 0xfe, 0xb0, 0xeb, 0x87, 0x51, 0xcb, 0x10, 0x5a, 0x6d, 0xbf, 0x05, 0xa7, 0x2b, 0x45,
	0xf5, 0x15, 0x12, 0xa5, 0x99, 0x73, 0xab, 0x67, 0xf5, 0x0f, 0xc3, 0x6e, 0x59, 0x78, 0xce, 0xee,
	0x26, 0x6b, 0x09, 0x8c, 0x4e, 0x0c, 0xf7, 0x69, 0x4d, 0x75, 0x41, 0x67, 0xff, 0xb2, 0x23, 0x2c,
	0x32, 0xce, 0x04, 0x86, 0x3f, 0x2d, 0xd0, 0x3e, 0x17, 0xc9, 0x67, 0x46, 0xff, 0x49, 0xc3, 0x33,
	0xd0, 0x44, 0x9b, 0xa5, 0x89, 0xc4, 0xd6, 0x78, 0xb7, 0x8a, 0x30, 0xda, 0x96, 0xda, 0x01, 0xa8,
	0xcf, 0x04, 0xce, 0x19, 0x4a, 0xb1, 0xca, 0x46, 0x23, 0xbc, 0x53, 0x16, 0xde, 0x6d, 0x6d, 0x5b,
	0x55, 0x60, 0xb4, 0x16, 0xa9, 0xe0, 0x91, 0x84, 0xe1, 0xdc, 0xdc, 0xfc, 0x76, 0xf0, 0x14, 0x5f,
	0x05, 0x4f, 0x03, 0x17, 0x74, 0x6f, 0xea, 0x76, 0x75, 0x9c, 0xf0, 0xdd, 0xd5, 0xc2, 0xb5, 0xae,
	0x17, 0xae, 0xf5, 0x7b, 0xe1, 0x5a, 0x3f, 0x96, 0x6e, 0xed, 0x7a, 0xe9, 0xd6, 0x7e, 0x2d, 0xdd,
	0xda, 0x97, 0x41, 0x42, 0xe4, 0x74, 0x36, 0xf6, 0x63, 0x9e, 0x06, 0x3a, 0xa9, 0x8f, 0x28, 0x1a,
	0x0b, 0x83, 0x83, 0xf9, 0x93, 0xe0, 0xfb, 0xe6, 0xc9, 0xc9, 0xcb, 0x0c, 0x8b, 0xf1, 0x91, 0x7a,
	0x2e, 0xc3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x29, 0x0f, 0xc1, 0x99, 0xcf, 0x03, 0x00, 0x00,
}

func (m *MsgLinkApplication) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkApplication) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkApplication) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x38
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CallData) > 0 {
		i -= len(m.CallData)
		copy(dAtA[i:], m.CallData)
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(len(m.CallData)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.LinkData.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgLinkApplicationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgLinkApplicationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgLinkApplicationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUnlinkApplication) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlinkApplication) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlinkApplication) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Application) > 0 {
		i -= len(m.Application)
		copy(dAtA[i:], m.Application)
		i = encodeVarintMsgsAppLinks(dAtA, i, uint64(len(m.Application)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnlinkApplicationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlinkApplicationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlinkApplicationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsAppLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsAppLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgLinkApplication) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMsgsAppLinks(uint64(l))
	}
	l = m.LinkData.Size()
	n += 1 + l + sovMsgsAppLinks(uint64(l))
	l = len(m.CallData)
	if l > 0 {
		n += 1 + l + sovMsgsAppLinks(uint64(l))
	}
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovMsgsAppLinks(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovMsgsAppLinks(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovMsgsAppLinks(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovMsgsAppLinks(uint64(m.TimeoutTimestamp))
	}
	return n
}

func (m *MsgLinkApplicationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUnlinkApplication) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Application)
	if l > 0 {
		n += 1 + l + sovMsgsAppLinks(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovMsgsAppLinks(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovMsgsAppLinks(uint64(l))
	}
	return n
}

func (m *MsgUnlinkApplicationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsAppLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsAppLinks(x uint64) (n int) {
	return sovMsgsAppLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgLinkApplication) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsAppLinks
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
			return fmt.Errorf("proto: MsgLinkApplication: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkApplication: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LinkData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CallData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CallData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
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
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
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
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
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
					return ErrIntOverflowMsgsAppLinks
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
			skippy, err := skipMsgsAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsAppLinks
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
func (m *MsgLinkApplicationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsAppLinks
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
			return fmt.Errorf("proto: MsgLinkApplicationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgLinkApplicationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsAppLinks
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
func (m *MsgUnlinkApplication) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsAppLinks
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
			return fmt.Errorf("proto: MsgUnlinkApplication: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlinkApplication: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Application", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Application = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsAppLinks
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
				return ErrInvalidLengthMsgsAppLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsAppLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsAppLinks
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
func (m *MsgUnlinkApplicationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsAppLinks
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
			return fmt.Errorf("proto: MsgUnlinkApplicationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlinkApplicationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsAppLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsAppLinks
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
func skipMsgsAppLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsAppLinks
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
					return 0, ErrIntOverflowMsgsAppLinks
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
					return 0, ErrIntOverflowMsgsAppLinks
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
				return 0, ErrInvalidLengthMsgsAppLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsAppLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsAppLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsAppLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsAppLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsAppLinks = fmt.Errorf("proto: unexpected end of group")
)
