// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta2/models_chain_links.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ChainLink contains the data representing either an inter- or cross- chain
// link
type ChainLink struct {
	// User defines the destination profile address to link
	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" yaml:"user"`
	// Address contains the data of the external chain address to be connected
	// with the Desmos profile
	Address *types.Any `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// Proof contains the ownership proof of the external chain address
	Proof Proof `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof" yaml:"proof"`
	// ChainConfig contains the configuration of the external chain
	ChainConfig ChainConfig `protobuf:"bytes,4,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config" yaml:"chain_config"`
	// CreationTime represents the time in which the link has been created
	CreationTime time.Time `protobuf:"bytes,5,opt,name=creation_time,json=creationTime,proto3,stdtime" json:"creation_time" yaml:"creation_time"`
}

func (m *ChainLink) Reset()         { *m = ChainLink{} }
func (m *ChainLink) String() string { return proto.CompactTextString(m) }
func (*ChainLink) ProtoMessage()    {}
func (*ChainLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f6d8e6d63f50f5, []int{0}
}
func (m *ChainLink) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainLink.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainLink.Merge(m, src)
}
func (m *ChainLink) XXX_Size() int {
	return m.Size()
}
func (m *ChainLink) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainLink.DiscardUnknown(m)
}

var xxx_messageInfo_ChainLink proto.InternalMessageInfo

// ChainConfig contains the data of the chain with which the link is made.
type ChainConfig struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f6d8e6d63f50f5, []int{1}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

// Proof contains all the data used to verify a signature when linking an
// account to a profile
type Proof struct {
	// PubKey represents the public key associated with the address for which to
	// prove the ownership
	PubKey *types.Any `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty" yaml:"pub_key"`
	// Signature represents the hex-encoded signature of the PlainText value
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty" yaml:"signature"`
	// PlainText represents the value signed in order to produce the Signature
	PlainText string `protobuf:"bytes,3,opt,name=plain_text,json=plainText,proto3" json:"plain_text,omitempty" yaml:"plain_text"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f6d8e6d63f50f5, []int{2}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

// Bech32Address represents a Bech32-encoded address
type Bech32Address struct {
	// Value represents the Bech-32 encoded address value
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
	// Prefix represents the HRP of the Bech32 address
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty" yaml:"prefix"`
}

func (m *Bech32Address) Reset()         { *m = Bech32Address{} }
func (m *Bech32Address) String() string { return proto.CompactTextString(m) }
func (*Bech32Address) ProtoMessage()    {}
func (*Bech32Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f6d8e6d63f50f5, []int{3}
}
func (m *Bech32Address) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bech32Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bech32Address.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bech32Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bech32Address.Merge(m, src)
}
func (m *Bech32Address) XXX_Size() int {
	return m.Size()
}
func (m *Bech32Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Bech32Address.DiscardUnknown(m)
}

var xxx_messageInfo_Bech32Address proto.InternalMessageInfo

// Base58Address represents a Base58-encoded address
type Base58Address struct {
	// Value contains the Base58-encoded address
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
}

func (m *Base58Address) Reset()         { *m = Base58Address{} }
func (m *Base58Address) String() string { return proto.CompactTextString(m) }
func (*Base58Address) ProtoMessage()    {}
func (*Base58Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6f6d8e6d63f50f5, []int{4}
}
func (m *Base58Address) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Base58Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Base58Address.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Base58Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Base58Address.Merge(m, src)
}
func (m *Base58Address) XXX_Size() int {
	return m.Size()
}
func (m *Base58Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Base58Address.DiscardUnknown(m)
}

var xxx_messageInfo_Base58Address proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChainLink)(nil), "desmos.profiles.v1beta2.ChainLink")
	proto.RegisterType((*ChainConfig)(nil), "desmos.profiles.v1beta2.ChainConfig")
	proto.RegisterType((*Proof)(nil), "desmos.profiles.v1beta2.Proof")
	proto.RegisterType((*Bech32Address)(nil), "desmos.profiles.v1beta2.Bech32Address")
	proto.RegisterType((*Base58Address)(nil), "desmos.profiles.v1beta2.Base58Address")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta2/models_chain_links.proto", fileDescriptor_d6f6d8e6d63f50f5)
}

var fileDescriptor_d6f6d8e6d63f50f5 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0x8f, 0x69, 0xd3, 0x2a, 0x97, 0x04, 0xda, 0x23, 0x88, 0xd0, 0x4a, 0xbe, 0x72, 0x20, 0x54,
	0x86, 0xda, 0x90, 0x82, 0x84, 0x3a, 0x51, 0x97, 0x01, 0x01, 0x03, 0x58, 0x9d, 0x58, 0xa2, 0xb3,
	0x73, 0x71, 0xad, 0xda, 0x3e, 0xcb, 0x77, 0xae, 0x92, 0x89, 0x95, 0xb1, 0x23, 0x63, 0x27, 0x3e,
	0x01, 0x9f, 0x80, 0xa9, 0x62, 0xaa, 0x98, 0x98, 0x0c, 0x6a, 0x17, 0xe6, 0x7c, 0x02, 0x74, 0x77,
	0x76, 0x13, 0x5a, 0x15, 0x89, 0xed, 0xf9, 0xfd, 0xfe, 0xbc, 0x77, 0xef, 0x3d, 0x19, 0x3c, 0x1a,
	0x50, 0x1e, 0x33, 0x6e, 0xa7, 0x19, 0x1b, 0x86, 0x11, 0xe5, 0xf6, 0xc1, 0x63, 0x8f, 0x0a, 0xd2,
	0xb3, 0x63, 0x36, 0xa0, 0x11, 0xef, 0xfb, 0x7b, 0x24, 0x4c, 0xfa, 0x51, 0x98, 0xec, 0x73, 0x2b,
	0xcd, 0x98, 0x60, 0xf0, 0xb6, 0x56, 0x58, 0x95, 0xc2, 0x2a, 0x15, 0x2b, 0x9d, 0x80, 0x05, 0x4c,
	0x71, 0x6c, 0x19, 0x69, 0xfa, 0xca, 0x9d, 0x80, 0xb1, 0x20, 0xa2, 0xb6, 0xfa, 0xf2, 0xf2, 0xa1,
	0x4d, 0x92, 0x71, 0x09, 0xa1, 0x8b, 0x90, 0x08, 0x63, 0xca, 0x05, 0x89, 0xd3, 0x4a, 0xeb, 0x33,
	0x59, 0xaa, 0xaf, 0x4d, 0xf5, 0x87, 0x86, 0xf0, 0xe7, 0x39, 0xd0, 0xd8, 0x91, 0xbd, 0xbd, 0x09,
	0x93, 0x7d, 0x78, 0x0f, 0xcc, 0xe7, 0x9c, 0x66, 0x5d, 0x63, 0xcd, 0x58, 0x6f, 0x38, 0x37, 0x26,
	0x05, 0x6a, 0x8e, 0x49, 0x1c, 0x6d, 0x61, 0x99, 0xc5, 0xae, 0x02, 0xe1, 0x3b, 0xb0, 0x48, 0x06,
	0x83, 0x8c, 0x72, 0xde, 0xbd, 0xb6, 0x66, 0xac, 0x37, 0x7b, 0x1d, 0x4b, 0x37, 0x60, 0x55, 0x0d,
	0x58, 0xdb, 0xc9, 0xd8, 0xb9, 0x3b, 0x29, 0xd0, 0x75, 0xad, 0x2e, 0xe9, 0xf8, 0xdb, 0x97, 0x8d,
	0xe6, 0xb6, 0x8e, 0x5f, 0x10, 0x41, 0xdc, 0xca, 0x07, 0xbe, 0x02, 0xf5, 0x34, 0x63, 0x6c, 0xd8,
	0x9d, 0x53, 0x86, 0xa6, 0x75, 0xc5, 0x6c, 0xac, 0xb7, 0x92, 0xe5, 0x74, 0x8e, 0x0b, 0x54, 0x9b,
	0x14, 0xa8, 0xa5, 0xed, 0x95, 0x14, 0xbb, 0xda, 0x02, 0x0e, 0x40, 0x4b, 0x0f, 0xdb, 0x67, 0xc9,
	0x30, 0x0c, 0xba, 0xf3, 0xca, 0xf2, 0xfe, 0x95, 0x96, 0xea, 0xf5, 0x3b, 0x8a, 0xeb, 0xac, 0x96,
	0xc6, 0x37, 0xb5, 0xf1, 0xac, 0x0f, 0x76, 0x9b, 0xfe, 0x94, 0x09, 0x09, 0x68, 0xfb, 0x19, 0x25,
	0x22, 0x64, 0x49, 0x5f, 0x8e, 0xbb, 0x5b, 0x57, 0x65, 0x56, 0x2e, 0x8d, 0x62, 0xb7, 0xda, 0x85,
	0xb3, 0x56, 0x9a, 0x77, 0x4a, 0xf3, 0x59, 0x39, 0x3e, 0xfc, 0x89, 0x0c, 0xb7, 0x55, 0xe5, 0xa4,
	0x68, 0xab, 0xf5, 0xf1, 0x08, 0xd5, 0x3e, 0x1d, 0x21, 0xe3, 0xf7, 0x11, 0x32, 0xf0, 0x73, 0xd0,
	0x9c, 0xe9, 0x54, 0x6e, 0x2a, 0x21, 0x31, 0xbd, 0xbc, 0x29, 0x99, 0xc5, 0xae, 0x02, 0x2f, 0x38,
	0x7c, 0x35, 0x40, 0x5d, 0xcd, 0x0f, 0x6e, 0x83, 0xc5, 0x34, 0xf7, 0xfa, 0xfb, 0x74, 0xac, 0xf4,
	0x57, 0x6d, 0x10, 0x4e, 0x37, 0x58, 0xd2, 0xb1, 0xbb, 0x90, 0xe6, 0xde, 0x6b, 0x3a, 0x86, 0x3d,
	0xd0, 0xe0, 0x61, 0x90, 0x10, 0x91, 0x67, 0x54, 0x9d, 0x41, 0xc3, 0xe9, 0x4c, 0x0a, 0xb4, 0xa4,
	0xe9, 0xe7, 0x10, 0x76, 0xa7, 0x34, 0xf8, 0x04, 0x80, 0x34, 0x92, 0x13, 0x15, 0x74, 0x24, 0xd4,
	0xaa, 0x1b, 0xce, 0xad, 0x49, 0x81, 0x96, 0xcb, 0x1a, 0xe7, 0x18, 0x76, 0x1b, 0xea, 0x63, 0x97,
	0x8e, 0xc4, 0x85, 0x47, 0x7c, 0x00, 0x6d, 0x87, 0xfa, 0x7b, 0x9b, 0xbd, 0xf2, 0x8e, 0xe0, 0x03,
	0x50, 0x3f, 0x20, 0x51, 0x5e, 0x4d, 0x62, 0x69, 0x7a, 0x16, 0x2a, 0x8d, 0x5d, 0x0d, 0xc3, 0x87,
	0x60, 0x21, 0xcd, 0xe8, 0x30, 0x1c, 0x95, 0xdd, 0x2e, 0x4f, 0x0a, 0xd4, 0xae, 0xee, 0x47, 0xe6,
	0xe5, 0xdb, 0x54, 0xb0, 0xb5, 0x3a, 0x5b, 0xf1, 0xfb, 0xdf, 0x37, 0x8b, 0x77, 0x41, 0xdb, 0x21,
	0x9c, 0x3e, 0x7d, 0xf6, 0x9f, 0x0d, 0xfc, 0xd3, 0xd5, 0x79, 0x79, 0x7c, 0x6a, 0x1a, 0x27, 0xa7,
	0xa6, 0xf1, 0xeb, 0xd4, 0x34, 0x0e, 0xcf, 0xcc, 0xda, 0xc9, 0x99, 0x59, 0xfb, 0x71, 0x66, 0xd6,
	0xde, 0x5b, 0x41, 0x28, 0xf6, 0x72, 0xcf, 0xf2, 0x59, 0x6c, 0xeb, 0x13, 0xde, 0x88, 0x88, 0xc7,
	0xcb, 0xd8, 0x1e, 0x4d, 0xff, 0x38, 0x62, 0x9c, 0x52, 0xee, 0x2d, 0xa8, 0x15, 0x6e, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x00, 0x05, 0x01, 0x61, 0x91, 0x04, 0x00, 0x00,
}

func (this *ChainLink) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChainLink)
	if !ok {
		that2, ok := that.(ChainLink)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.User != that1.User {
		return false
	}
	if !this.Address.Equal(that1.Address) {
		return false
	}
	if !this.Proof.Equal(&that1.Proof) {
		return false
	}
	if !this.ChainConfig.Equal(&that1.ChainConfig) {
		return false
	}
	if !this.CreationTime.Equal(that1.CreationTime) {
		return false
	}
	return true
}
func (this *ChainConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChainConfig)
	if !ok {
		that2, ok := that.(ChainConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	return true
}
func (this *Proof) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Proof)
	if !ok {
		that2, ok := that.(Proof)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.PubKey.Equal(that1.PubKey) {
		return false
	}
	if this.Signature != that1.Signature {
		return false
	}
	if this.PlainText != that1.PlainText {
		return false
	}
	return true
}
func (this *Bech32Address) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Bech32Address)
	if !ok {
		that2, ok := that.(Bech32Address)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Prefix != that1.Prefix {
		return false
	}
	return true
}
func (this *Base58Address) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Base58Address)
	if !ok {
		that2, ok := that.(Base58Address)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	return true
}
func (m *ChainLink) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainLink) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainLink) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintModelsChainLinks(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	{
		size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Address != nil {
		{
			size, err := m.Address.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsChainLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PlainText) > 0 {
		i -= len(m.PlainText)
		copy(dAtA[i:], m.PlainText)
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(len(m.PlainText)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsChainLinks(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Bech32Address) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bech32Address) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bech32Address) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Prefix) > 0 {
		i -= len(m.Prefix)
		copy(dAtA[i:], m.Prefix)
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(len(m.Prefix)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Base58Address) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Base58Address) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Base58Address) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintModelsChainLinks(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsChainLinks(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsChainLinks(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainLink) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	l = m.Proof.Size()
	n += 1 + l + sovModelsChainLinks(uint64(l))
	l = m.ChainConfig.Size()
	n += 1 + l + sovModelsChainLinks(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationTime)
	n += 1 + l + sovModelsChainLinks(uint64(l))
	return n
}

func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	return n
}

func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	l = len(m.PlainText)
	if l > 0 {
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	return n
}

func (m *Bech32Address) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	return n
}

func (m *Base58Address) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovModelsChainLinks(uint64(l))
	}
	return n
}

func sovModelsChainLinks(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsChainLinks(x uint64) (n int) {
	return sovModelsChainLinks(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainLink) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsChainLinks
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
			return fmt.Errorf("proto: ChainLink: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainLink: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Address == nil {
				m.Address = &types.Any{}
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsChainLinks
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
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsChainLinks
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsChainLinks
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
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsChainLinks
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &types.Any{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlainText", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlainText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsChainLinks
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
func (m *Bech32Address) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsChainLinks
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
			return fmt.Errorf("proto: Bech32Address: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bech32Address: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsChainLinks
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
func (m *Base58Address) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsChainLinks
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
			return fmt.Errorf("proto: Base58Address: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Base58Address: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsChainLinks
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
				return ErrInvalidLengthModelsChainLinks
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsChainLinks
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsChainLinks(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsChainLinks
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
func skipModelsChainLinks(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsChainLinks
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
					return 0, ErrIntOverflowModelsChainLinks
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
					return 0, ErrIntOverflowModelsChainLinks
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
				return 0, ErrInvalidLengthModelsChainLinks
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsChainLinks
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsChainLinks
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsChainLinks        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsChainLinks          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsChainLinks = fmt.Errorf("proto: unexpected end of group")
)
