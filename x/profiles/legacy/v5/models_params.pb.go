// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v2/models_params.proto

package v5

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Params contains the parameters for the profiles module
type Params struct {
	Nickname NicknameParams `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname" yaml:"nickname"`
	DTag     DTagParams     `protobuf:"bytes,2,opt,name=dtag,proto3" json:"dtag" yaml:"dtag"`
	Bio      BioParams      `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio" yaml:"bio"`
	Oracle   OracleParams   `protobuf:"bytes,4,opt,name=oracle,proto3" json:"oracle" yaml:"oracle"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b093f5fee9c9f7d, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// NicknameParams defines the parameters related to the profiles nicknames
type NicknameParams struct {
	MinLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=min_length,json=minLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_length" yaml:"min_length"`
	MaxLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=max_length,json=maxLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_length" yaml:"max_length"`
}

func (m *NicknameParams) Reset()         { *m = NicknameParams{} }
func (m *NicknameParams) String() string { return proto.CompactTextString(m) }
func (*NicknameParams) ProtoMessage()    {}
func (*NicknameParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b093f5fee9c9f7d, []int{1}
}
func (m *NicknameParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NicknameParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NicknameParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NicknameParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NicknameParams.Merge(m, src)
}
func (m *NicknameParams) XXX_Size() int {
	return m.Size()
}
func (m *NicknameParams) XXX_DiscardUnknown() {
	xxx_messageInfo_NicknameParams.DiscardUnknown(m)
}

var xxx_messageInfo_NicknameParams proto.InternalMessageInfo

// DTagParams defines the parameters related to profile DTags
type DTagParams struct {
	RegEx     string                                 `protobuf:"bytes,1,opt,name=reg_ex,json=regEx,proto3" json:"reg_ex,omitempty" yaml:"reg_ex"`
	MinLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_length,json=minLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_length" yaml:"min_length"`
	MaxLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_length,json=maxLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_length" yaml:"max_length"`
}

func (m *DTagParams) Reset()         { *m = DTagParams{} }
func (m *DTagParams) String() string { return proto.CompactTextString(m) }
func (*DTagParams) ProtoMessage()    {}
func (*DTagParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b093f5fee9c9f7d, []int{2}
}
func (m *DTagParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DTagParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DTagParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DTagParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DTagParams.Merge(m, src)
}
func (m *DTagParams) XXX_Size() int {
	return m.Size()
}
func (m *DTagParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DTagParams.DiscardUnknown(m)
}

var xxx_messageInfo_DTagParams proto.InternalMessageInfo

// BioParams defines the parameters related to profile biography
type BioParams struct {
	MaxLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_length,json=maxLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_length" yaml:"max_length"`
}

func (m *BioParams) Reset()         { *m = BioParams{} }
func (m *BioParams) String() string { return proto.CompactTextString(m) }
func (*BioParams) ProtoMessage()    {}
func (*BioParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b093f5fee9c9f7d, []int{3}
}
func (m *BioParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BioParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BioParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BioParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BioParams.Merge(m, src)
}
func (m *BioParams) XXX_Size() int {
	return m.Size()
}
func (m *BioParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BioParams.DiscardUnknown(m)
}

var xxx_messageInfo_BioParams proto.InternalMessageInfo

// OracleParams defines the parameters related to the oracle
// that will be used to verify the ownership of a centralized
// application account by a Desmos profile
type OracleParams struct {
	// ScriptID represents the ID of the oracle script to be called to verify the
	// data
	ScriptID uint64 `protobuf:"varint,1,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty" yaml:"script_id"`
	// AskCount represents the number of oracles to which ask to verify the data
	AskCount uint64 `protobuf:"varint,2,opt,name=ask_count,json=askCount,proto3" json:"ask_count,omitempty" yaml:"ask_count"`
	// MinCount represents the minimum count of oracles that should complete the
	// verification successfully
	MinCount uint64 `protobuf:"varint,3,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty" yaml:"min_count"`
	// PrepareGas represents the amount of gas to be used during the preparation
	// stage of the oracle script
	PrepareGas uint64 `protobuf:"varint,4,opt,name=prepare_gas,json=prepareGas,proto3" json:"prepare_gas,omitempty" yaml:"prepare_gas"`
	// ExecuteGas represents the amount of gas to be used during the execution of
	// the oracle script
	ExecuteGas uint64 `protobuf:"varint,5,opt,name=execute_gas,json=executeGas,proto3" json:"execute_gas,omitempty" yaml:"execute_gas"`
	// FeeAmount represents the amount of fees to be payed in order to execute the
	// oracle script
	FeeAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=fee_amount,json=feeAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_amount" yaml:"fee_amount"`
}

func (m *OracleParams) Reset()         { *m = OracleParams{} }
func (m *OracleParams) String() string { return proto.CompactTextString(m) }
func (*OracleParams) ProtoMessage()    {}
func (*OracleParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b093f5fee9c9f7d, []int{4}
}
func (m *OracleParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OracleParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OracleParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OracleParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OracleParams.Merge(m, src)
}
func (m *OracleParams) XXX_Size() int {
	return m.Size()
}
func (m *OracleParams) XXX_DiscardUnknown() {
	xxx_messageInfo_OracleParams.DiscardUnknown(m)
}

var xxx_messageInfo_OracleParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "desmos.profiles.v2.Params")
	proto.RegisterType((*NicknameParams)(nil), "desmos.profiles.v2.NicknameParams")
	proto.RegisterType((*DTagParams)(nil), "desmos.profiles.v2.DTagParams")
	proto.RegisterType((*BioParams)(nil), "desmos.profiles.v2.BioParams")
	proto.RegisterType((*OracleParams)(nil), "desmos.profiles.v2.OracleParams")
}

func init() {
	proto.RegisterFile("desmos/profiles/v2/models_params.proto", fileDescriptor_1b093f5fee9c9f7d)
}

var fileDescriptor_1b093f5fee9c9f7d = []byte{
	// 718 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x73, 0x6b, 0x94, 0x4c, 0xca, 0xa5, 0x56, 0x81, 0x50, 0x84, 0x1d, 0xcd, 0xa2, 0xca,
	0xa6, 0xb6, 0x1a, 0xa8, 0x90, 0x2a, 0xb1, 0xc0, 0x69, 0x85, 0x2a, 0x01, 0xa9, 0x0c, 0x12, 0x12,
	0x9b, 0x68, 0xec, 0x4c, 0xdc, 0x51, 0x6c, 0x8f, 0xe5, 0x71, 0xa2, 0x74, 0xc5, 0x96, 0x25, 0x4f,
	0x80, 0x58, 0x23, 0x1e, 0xa4, 0x2b, 0xd4, 0x25, 0x62, 0x61, 0x50, 0xfa, 0x06, 0x79, 0x02, 0x34,
	0x97, 0x5c, 0x5a, 0x2a, 0x51, 0x40, 0xac, 0x32, 0x9e, 0x73, 0xfe, 0xef, 0x9c, 0x39, 0xbf, 0x3d,
	0x01, 0x9b, 0x3d, 0xcc, 0x42, 0xca, 0xac, 0x38, 0xa1, 0x7d, 0x12, 0x60, 0x66, 0x8d, 0x5a, 0x56,
	0x48, 0x7b, 0x38, 0x60, 0xdd, 0x18, 0x25, 0x28, 0x64, 0x66, 0x9c, 0xd0, 0x94, 0x6a, 0x9a, 0xcc,
	0x33, 0x67, 0x79, 0xe6, 0xa8, 0xb5, 0xb1, 0xee, 0x53, 0x9f, 0x8a, 0xb0, 0xc5, 0x57, 0x32, 0x73,
	0x43, 0xf7, 0xa8, 0x20, 0xba, 0x88, 0x61, 0x6b, 0xb4, 0xed, 0xe2, 0x14, 0x6d, 0x5b, 0x1e, 0x25,
	0xd1, 0x2c, 0xee, 0x53, 0xea, 0x07, 0xd8, 0x12, 0x4f, 0xee, 0xb0, 0x6f, 0xf5, 0x86, 0x09, 0x4a,
	0x09, 0x55, 0x71, 0xf8, 0xa5, 0x00, 0xca, 0x87, 0xa2, 0xb4, 0xf6, 0x1a, 0x54, 0x22, 0xe2, 0x0d,
	0x22, 0x14, 0xe2, 0x7a, 0xbe, 0x91, 0x6f, 0xd6, 0x5a, 0xd0, 0xfc, 0xb5, 0x0f, 0xf3, 0x85, 0xca,
	0x91, 0x2a, 0xfb, 0xce, 0x49, 0x66, 0xe4, 0xa6, 0x99, 0x71, 0xe3, 0x18, 0x85, 0xc1, 0x2e, 0x9c,
	0x11, 0xa0, 0x33, 0x87, 0x69, 0x1d, 0x50, 0xea, 0xa5, 0xc8, 0xaf, 0x17, 0x04, 0x54, 0xbf, 0x0c,
	0xba, 0xf7, 0x0a, 0xf9, 0x0a, 0x78, 0x8f, 0x03, 0x27, 0x99, 0x51, 0xe2, 0x7b, 0xd3, 0xcc, 0xa8,
	0x49, 0x30, 0x27, 0x40, 0x47, 0x80, 0xb4, 0x36, 0x28, 0xba, 0x84, 0xd6, 0x8b, 0x82, 0x77, 0xff,
	0x32, 0x9e, 0x4d, 0xa8, 0xc2, 0x69, 0xaa, 0x3f, 0x20, 0x31, 0x2e, 0xa1, 0xd0, 0xe1, 0x6a, 0xad,
	0x03, 0xca, 0x34, 0x41, 0x5e, 0x80, 0xeb, 0x25, 0xc1, 0x69, 0x5c, 0xc6, 0xe9, 0x88, 0x0c, 0x85,
	0xba, 0xa5, 0x50, 0xd7, 0x24, 0x4a, 0xaa, 0xa1, 0xa3, 0x30, 0xbb, 0xa5, 0x77, 0x1f, 0x8d, 0x1c,
	0xcc, 0xf2, 0xe0, 0xfa, 0xf9, 0x11, 0x69, 0x2e, 0x00, 0x21, 0x89, 0xba, 0x01, 0x8e, 0xfc, 0xf4,
	0x48, 0x8c, 0x76, 0xd5, 0x6e, 0x73, 0xd6, 0xb7, 0xcc, 0xd8, 0xf4, 0x49, 0x7a, 0x34, 0x74, 0x4d,
	0x8f, 0x86, 0x96, 0xb2, 0x52, 0xfe, 0x6c, 0xb1, 0xde, 0xc0, 0x4a, 0x8f, 0x63, 0xcc, 0xcc, 0x83,
	0x28, 0x9d, 0x66, 0xc6, 0x9a, 0xac, 0xba, 0x20, 0x41, 0xa7, 0x1a, 0x92, 0xe8, 0x99, 0x58, 0x8b,
	0x1a, 0x68, 0x3c, 0xab, 0x51, 0xf8, 0xc7, 0x1a, 0x73, 0x12, 0xaf, 0x81, 0xc6, 0xb2, 0x86, 0x3a,
	0xe0, 0x87, 0x02, 0x00, 0x0b, 0xbb, 0xb4, 0x26, 0x28, 0x27, 0xd8, 0xef, 0xe2, 0xb1, 0x38, 0x58,
	0xd5, 0x5e, 0x5b, 0x0c, 0x48, 0xee, 0x43, 0x67, 0x25, 0xc1, 0xfe, 0xfe, 0x58, 0xa3, 0xe7, 0xc6,
	0x20, 0x5b, 0x3c, 0xfc, 0xb3, 0x16, 0x27, 0x99, 0x51, 0x7d, 0x3e, 0x3b, 0xf3, 0x6f, 0x67, 0x42,
	0xcf, 0xcd, 0xa4, 0xf8, 0xd7, 0x05, 0x67, 0x03, 0xb8, 0xe2, 0x80, 0x86, 0xa0, 0x3a, 0x7f, 0xfd,
	0x2e, 0xf8, 0x52, 0xfc, 0x8f, 0xbe, 0x7c, 0x2e, 0x82, 0xd5, 0xe5, 0xd7, 0x55, 0x7b, 0x0c, 0xaa,
	0xcc, 0x4b, 0x48, 0x9c, 0x76, 0x49, 0x4f, 0x98, 0x53, 0xb2, 0x1b, 0x93, 0xcc, 0xa8, 0xbc, 0x14,
	0x9b, 0x07, 0x7b, 0xd3, 0xcc, 0xb8, 0x29, 0xb9, 0xf3, 0x34, 0xe8, 0x54, 0xe4, 0xfa, 0xa0, 0xa7,
	0x6d, 0x83, 0x2a, 0x62, 0x83, 0xae, 0x47, 0x87, 0x51, 0x2a, 0xdc, 0x2a, 0xd9, 0xeb, 0x0b, 0xc9,
	0x3c, 0x04, 0x9d, 0x0a, 0x62, 0x83, 0x36, 0x5f, 0x72, 0x09, 0xb7, 0x42, 0x4a, 0x8a, 0x17, 0x25,
	0xf3, 0x10, 0x74, 0x2a, 0x21, 0x89, 0xa4, 0xe4, 0x11, 0xa8, 0xc5, 0x09, 0x8e, 0x51, 0x82, 0xbb,
	0x3e, 0x62, 0xe2, 0x53, 0x2c, 0xd9, 0xb7, 0xa7, 0x99, 0xa1, 0x49, 0xd1, 0x52, 0x10, 0x3a, 0x40,
	0x3d, 0x3d, 0x45, 0x8c, 0x0b, 0xf1, 0x18, 0x7b, 0xc3, 0x54, 0x0a, 0x57, 0x2e, 0x0a, 0x97, 0x82,
	0xd0, 0x01, 0xea, 0x89, 0x0b, 0xdf, 0x02, 0xd0, 0xc7, 0xb8, 0x8b, 0x42, 0xd1, 0x65, 0xb9, 0x51,
	0x6c, 0xd6, 0x5a, 0x77, 0x4d, 0x39, 0x78, 0x93, 0x5f, 0xa3, 0xa6, 0xba, 0x46, 0xcd, 0x36, 0x25,
	0x91, 0xbd, 0xaf, 0x3e, 0x7a, 0x65, 0xc1, 0x42, 0x0a, 0x3f, 0x7d, 0x37, 0x9a, 0x57, 0x70, 0x90,
	0x53, 0x98, 0x53, 0xed, 0x63, 0xfc, 0x44, 0xe8, 0xa4, 0x5d, 0x76, 0xe7, 0x64, 0xa2, 0xe7, 0x4f,
	0x27, 0x7a, 0xfe, 0xc7, 0x44, 0xcf, 0xbf, 0x3f, 0xd3, 0x73, 0xa7, 0x67, 0x7a, 0xee, 0xeb, 0x99,
	0x9e, 0x7b, 0xb3, 0xb3, 0x04, 0x95, 0x57, 0xd2, 0x56, 0x80, 0x5c, 0xa6, 0xd6, 0xd6, 0xe8, 0xa1,
	0x35, 0x5e, 0xfc, 0x81, 0x04, 0xd8, 0x47, 0xde, 0xb1, 0x35, 0xda, 0x71, 0xcb, 0xe2, 0x42, 0x7f,
	0xf0, 0x33, 0x00, 0x00, 0xff, 0xff, 0x95, 0x96, 0x45, 0x5c, 0x64, 0x06, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Oracle.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.Bio.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DTag.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Nickname.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *NicknameParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NicknameParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NicknameParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxLength.Size()
		i -= size
		if _, err := m.MaxLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinLength.Size()
		i -= size
		if _, err := m.MinLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DTagParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DTagParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DTagParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxLength.Size()
		i -= size
		if _, err := m.MaxLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinLength.Size()
		i -= size
		if _, err := m.MinLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RegEx) > 0 {
		i -= len(m.RegEx)
		copy(dAtA[i:], m.RegEx)
		i = encodeVarintModelsParams(dAtA, i, uint64(len(m.RegEx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BioParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BioParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BioParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxLength.Size()
		i -= size
		if _, err := m.MaxLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	return len(dAtA) - i, nil
}

func (m *OracleParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OracleParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OracleParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FeeAmount) > 0 {
		for iNdEx := len(m.FeeAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintModelsParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.ExecuteGas != 0 {
		i = encodeVarintModelsParams(dAtA, i, uint64(m.ExecuteGas))
		i--
		dAtA[i] = 0x28
	}
	if m.PrepareGas != 0 {
		i = encodeVarintModelsParams(dAtA, i, uint64(m.PrepareGas))
		i--
		dAtA[i] = 0x20
	}
	if m.MinCount != 0 {
		i = encodeVarintModelsParams(dAtA, i, uint64(m.MinCount))
		i--
		dAtA[i] = 0x18
	}
	if m.AskCount != 0 {
		i = encodeVarintModelsParams(dAtA, i, uint64(m.AskCount))
		i--
		dAtA[i] = 0x10
	}
	if m.ScriptID != 0 {
		i = encodeVarintModelsParams(dAtA, i, uint64(m.ScriptID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Nickname.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.DTag.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.Bio.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.Oracle.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	return n
}

func (m *NicknameParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinLength.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.MaxLength.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	return n
}

func (m *DTagParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RegEx)
	if l > 0 {
		n += 1 + l + sovModelsParams(uint64(l))
	}
	l = m.MinLength.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.MaxLength.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	return n
}

func (m *BioParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxLength.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	return n
}

func (m *OracleParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ScriptID != 0 {
		n += 1 + sovModelsParams(uint64(m.ScriptID))
	}
	if m.AskCount != 0 {
		n += 1 + sovModelsParams(uint64(m.AskCount))
	}
	if m.MinCount != 0 {
		n += 1 + sovModelsParams(uint64(m.MinCount))
	}
	if m.PrepareGas != 0 {
		n += 1 + sovModelsParams(uint64(m.PrepareGas))
	}
	if m.ExecuteGas != 0 {
		n += 1 + sovModelsParams(uint64(m.ExecuteGas))
	}
	if len(m.FeeAmount) > 0 {
		for _, e := range m.FeeAmount {
			l = e.Size()
			n += 1 + l + sovModelsParams(uint64(l))
		}
	}
	return n
}

func sovModelsParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsParams(x uint64) (n int) {
	return sovModelsParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
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
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Nickname.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTag", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
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
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DTag.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bio", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
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
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
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
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Oracle.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsParams
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
func (m *NicknameParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsParams
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
			return fmt.Errorf("proto: NicknameParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NicknameParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsParams
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
func (m *DTagParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsParams
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
			return fmt.Errorf("proto: DTagParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DTagParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegEx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
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
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegEx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsParams
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
func (m *BioParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsParams
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
			return fmt.Errorf("proto: BioParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BioParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsParams
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
func (m *OracleParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsParams
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
			return fmt.Errorf("proto: OracleParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OracleParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ScriptID", wireType)
			}
			m.ScriptID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ScriptID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AskCount", wireType)
			}
			m.AskCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AskCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinCount", wireType)
			}
			m.MinCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrepareGas", wireType)
			}
			m.PrepareGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrepareGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteGas", wireType)
			}
			m.ExecuteGas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExecuteGas |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsParams
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
				return ErrInvalidLengthModelsParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeAmount = append(m.FeeAmount, types.Coin{})
			if err := m.FeeAmount[len(m.FeeAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsParams
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
func skipModelsParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsParams
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
					return 0, ErrIntOverflowModelsParams
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
					return 0, ErrIntOverflowModelsParams
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
				return 0, ErrInvalidLengthModelsParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsParams = fmt.Errorf("proto: unexpected end of group")
)
