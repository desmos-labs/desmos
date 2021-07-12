// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/models_params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params contains the parameters for the profiles module
type Params struct {
	NicknameParams NicknameParams                         `protobuf:"bytes,1,opt,name=nickname_params,json=nicknameParams,proto3" json:"nickname_params" yaml:"nickname_params"`
	DTagParams     DTagParams                             `protobuf:"bytes,2,opt,name=dtag_params,json=dtagParams,proto3" json:"dtag_params" yaml:"dtag_params"`
	MaxBioLength   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_bio_length,json=maxBioLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_bio_length" yaml:"max_bio_length"`
	Oracle         OracleParams                           `protobuf:"bytes,4,opt,name=oracle,proto3" json:"oracle" yaml:"oracle"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_a621950d5c07fbad, []int{0}
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
	MinNicknameLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=min_nickname_length,json=minNicknameLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_nickname_length" yaml:"min_nickname_length"`
	MaxNicknameLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=max_nickname_length,json=maxNicknameLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_nickname_length" yaml:"max_nickname_length"`
}

func (m *NicknameParams) Reset()         { *m = NicknameParams{} }
func (m *NicknameParams) String() string { return proto.CompactTextString(m) }
func (*NicknameParams) ProtoMessage()    {}
func (*NicknameParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a621950d5c07fbad, []int{1}
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
	RegEx         string                                 `protobuf:"bytes,1,opt,name=reg_ex,json=regEx,proto3" json:"reg_ex,omitempty" yaml:"reg_ex"`
	MinDTagLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_dtag_length,json=minDtagLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_dtag_length" yaml:"min_dtag_length"`
	MaxDTagLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_dtag_length,json=maxDtagLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_dtag_length" yaml:"max_dtag_length"`
}

func (m *DTagParams) Reset()         { *m = DTagParams{} }
func (m *DTagParams) String() string { return proto.CompactTextString(m) }
func (*DTagParams) ProtoMessage()    {}
func (*DTagParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a621950d5c07fbad, []int{2}
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

// OracleParams defines the parameters related to the oracle
// that will be used to verify the ownership of a centralized
// application account by a Desmos profile
type OracleParams struct {
	// ScriptID represents the ID of the oracle script to be called to verify the data
	ScriptID int64 `protobuf:"varint,1,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty" yaml:"script_id"`
	// AskCount represents the number of oracles to which ask to verify the data
	AskCount uint64 `protobuf:"varint,2,opt,name=ask_count,json=askCount,proto3" json:"ask_count,omitempty" yaml:"ask_count"`
	// MinCount represents the minimum count of oracles that should complete the verification successfully
	MinCount uint64 `protobuf:"varint,3,opt,name=min_count,json=minCount,proto3" json:"min_count,omitempty" yaml:"min_count"`
	// PrepareGas represents the amount of gas to be used during the preparation stage of the oracle script
	PrepareGas uint64 `protobuf:"varint,4,opt,name=prepare_gas,json=prepareGas,proto3" json:"prepare_gas,omitempty" yaml:"prepare_gas"`
	// ExecuteGas represents the amount of gas to be used during the execution of the oracle script
	ExecuteGas uint64 `protobuf:"varint,5,opt,name=execute_gas,json=executeGas,proto3" json:"execute_gas,omitempty" yaml:"execute_gas"`
	// FeePayer represents the key of the account that is going to pay for oracle fees if needed
	FeePayer string `protobuf:"bytes,6,opt,name=fee_payer,json=feePayer,proto3" json:"fee_payer,omitempty" yaml:"fee_payer"`
	// FeeAmount represents the amount of fees to be payed in order to execute the oracle script
	FeeAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,7,rep,name=fee_amount,json=feeAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fee_amount" yaml:"fee_amount"`
}

func (m *OracleParams) Reset()         { *m = OracleParams{} }
func (m *OracleParams) String() string { return proto.CompactTextString(m) }
func (*OracleParams) ProtoMessage()    {}
func (*OracleParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_a621950d5c07fbad, []int{3}
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
	proto.RegisterType((*Params)(nil), "desmos.profiles.v1beta1.Params")
	proto.RegisterType((*NicknameParams)(nil), "desmos.profiles.v1beta1.NicknameParams")
	proto.RegisterType((*DTagParams)(nil), "desmos.profiles.v1beta1.DTagParams")
	proto.RegisterType((*OracleParams)(nil), "desmos.profiles.v1beta1.OracleParams")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/models_params.proto", fileDescriptor_a621950d5c07fbad)
}

var fileDescriptor_a621950d5c07fbad = []byte{
	// 755 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x93, 0x34, 0x24, 0x93, 0x3e, 0xa8, 0x69, 0x4b, 0xe8, 0xc2, 0x8e, 0x06, 0x01, 0x91,
	0x50, 0x1d, 0xa5, 0x2c, 0x90, 0x2a, 0xb1, 0xc0, 0x6d, 0x55, 0x2a, 0x15, 0xa8, 0x4c, 0x37, 0xb0,
	0xb1, 0x26, 0xce, 0xd4, 0xb5, 0x12, 0x7b, 0x2c, 0x8f, 0x8b, 0x5c, 0x09, 0x09, 0xb1, 0x63, 0x89,
	0xf8, 0x02, 0xd6, 0xec, 0xf9, 0x03, 0x16, 0x5d, 0x76, 0x89, 0x58, 0xb8, 0x28, 0xfd, 0x83, 0x7c,
	0x01, 0x9a, 0x47, 0xe2, 0x24, 0xa5, 0x12, 0x85, 0x95, 0x67, 0xe6, 0xde, 0x7b, 0xce, 0xb9, 0xc7,
	0xe3, 0x6b, 0xf0, 0xb0, 0x83, 0xa9, 0x4f, 0x68, 0x33, 0x8c, 0xc8, 0xa1, 0xd7, 0xc3, 0xb4, 0xf9,
	0xb6, 0xd5, 0xc6, 0x31, 0x6a, 0x35, 0x7d, 0xd2, 0xc1, 0x3d, 0x6a, 0x87, 0x28, 0x42, 0x3e, 0x35,
	0xc2, 0x88, 0xc4, 0x44, 0xbd, 0x2d, 0x92, 0x8d, 0x61, 0xb2, 0x21, 0x93, 0x57, 0x97, 0x5c, 0xe2,
	0x12, 0x9e, 0xd3, 0x64, 0x2b, 0x91, 0xbe, 0xaa, 0x39, 0x84, 0x63, 0xb7, 0x11, 0xc5, 0x23, 0x5c,
	0x87, 0x78, 0x81, 0x88, 0xc3, 0x6f, 0x05, 0x50, 0xda, 0xe7, 0xf8, 0x6a, 0x08, 0x16, 0x02, 0xcf,
	0xe9, 0x06, 0xc8, 0xc7, 0x92, 0xb2, 0xa6, 0xd4, 0x95, 0x46, 0x75, 0xfd, 0x81, 0x71, 0x05, 0xa7,
	0xf1, 0x42, 0xe6, 0x0b, 0x04, 0x53, 0x3b, 0x4d, 0xf5, 0xdc, 0x20, 0xd5, 0x57, 0x4e, 0x90, 0xdf,
	0xdb, 0x80, 0x53, 0x68, 0xd0, 0x9a, 0x0f, 0x26, 0xf2, 0xd5, 0x00, 0x54, 0x3b, 0x31, 0x72, 0x87,
	0x6c, 0x79, 0xce, 0x76, 0xf7, 0x4a, 0xb6, 0xad, 0x03, 0xe4, 0x4a, 0xa6, 0x06, 0x63, 0xea, 0xa7,
	0x3a, 0xc8, 0xce, 0x06, 0xa9, 0xae, 0x0a, 0xde, 0x31, 0x4c, 0x68, 0x01, 0xb6, 0x93, 0x7c, 0x3e,
	0x98, 0xf7, 0x51, 0x62, 0xb7, 0x3d, 0x62, 0xf7, 0x70, 0xe0, 0xc6, 0x47, 0xb5, 0x42, 0x5d, 0x69,
	0xcc, 0x9a, 0x3b, 0x0c, 0xed, 0x67, 0xaa, 0xdf, 0x77, 0xbd, 0xf8, 0xe8, 0xb8, 0x6d, 0x38, 0xc4,
	0x6f, 0x4a, 0xdf, 0xc4, 0x63, 0x8d, 0x76, 0xba, 0xcd, 0xf8, 0x24, 0xc4, 0xd4, 0xd8, 0x0d, 0xe2,
	0x41, 0xaa, 0x2f, 0x0b, 0xa6, 0x49, 0x34, 0x68, 0xcd, 0xfa, 0x28, 0x31, 0x3d, 0xb2, 0xc7, 0xb7,
	0xea, 0x01, 0x28, 0x91, 0x08, 0x39, 0x3d, 0x5c, 0x2b, 0xf2, 0xce, 0xee, 0x5d, 0xd9, 0xd9, 0x4b,
	0x9e, 0x26, 0x7b, 0x5b, 0x96, 0x2e, 0xce, 0x09, 0x0e, 0x01, 0x01, 0x2d, 0x89, 0xb5, 0x51, 0xfc,
	0xf8, 0x45, 0xcf, 0xc1, 0xcf, 0x79, 0x30, 0x3f, 0xe9, 0xbe, 0xfa, 0x0e, 0xdc, 0xf2, 0xbd, 0xc0,
	0x1e, 0xb9, 0x2e, 0x5b, 0x54, 0x78, 0x8b, 0x7b, 0xd7, 0x6e, 0x71, 0x55, 0xb6, 0x78, 0x19, 0x12,
	0x5a, 0x8b, 0xbe, 0x17, 0x0c, 0xd9, 0x65, 0xb3, 0x8c, 0x1d, 0x25, 0x97, 0xd8, 0xf3, 0xff, 0xc9,
	0x7e, 0x19, 0x92, 0xb1, 0xa3, 0x64, 0x92, 0x5d, 0x9a, 0xf2, 0x3d, 0x0f, 0xc6, 0x2e, 0x84, 0xda,
	0x00, 0xa5, 0x08, 0xbb, 0x36, 0x4e, 0xb8, 0x07, 0x15, 0x73, 0x31, 0x33, 0x55, 0x9c, 0x43, 0x6b,
	0x26, 0xc2, 0xee, 0x76, 0xa2, 0x7e, 0x50, 0xc0, 0x02, 0x6b, 0x94, 0xdf, 0x9c, 0x09, 0xe5, 0xaf,
	0xaf, 0xa7, 0xbc, 0x9f, 0xea, 0x73, 0xcf, 0xbd, 0x80, 0x89, 0x10, 0xca, 0xb2, 0xaf, 0x61, 0x0a,
	0x1f, 0x5a, 0x73, 0xbe, 0x17, 0x6c, 0xc5, 0xc3, 0x44, 0xa1, 0x01, 0x25, 0x13, 0x1a, 0x0a, 0xff,
	0xac, 0x01, 0x25, 0x7f, 0xd4, 0x30, 0x89, 0xcf, 0x34, 0xa0, 0x24, 0xd3, 0x20, 0x6d, 0x3c, 0x2f,
	0x80, 0xd9, 0xf1, 0x1b, 0xa9, 0x3e, 0x01, 0x15, 0xea, 0x44, 0x5e, 0x18, 0xdb, 0x5e, 0x87, 0x7b,
	0x59, 0x30, 0xeb, 0xfd, 0x54, 0x2f, 0xbf, 0xe2, 0x87, 0xbb, 0x5b, 0x83, 0x54, 0xbf, 0x29, 0x08,
	0x46, 0x69, 0xd0, 0x2a, 0x8b, 0xf5, 0x6e, 0x47, 0x6d, 0x81, 0x0a, 0xa2, 0x5d, 0xdb, 0x21, 0xc7,
	0x41, 0xcc, 0x6d, 0x2d, 0x9a, 0x4b, 0x59, 0xc9, 0x28, 0x04, 0xad, 0x32, 0xa2, 0xdd, 0x4d, 0xb6,
	0x64, 0x25, 0xcc, 0x2f, 0x51, 0x52, 0x98, 0x2e, 0x19, 0x85, 0xa0, 0x55, 0xf6, 0xbd, 0x40, 0x94,
	0x3c, 0x06, 0xd5, 0x30, 0xc2, 0x21, 0x8a, 0xb0, 0xed, 0x22, 0xca, 0x3f, 0xb9, 0xa2, 0xb9, 0x92,
	0x4d, 0x85, 0xb1, 0x20, 0xb4, 0x80, 0xdc, 0xed, 0x20, 0xca, 0x0a, 0x71, 0x82, 0x9d, 0xe3, 0x58,
	0x14, 0xce, 0x4c, 0x17, 0x8e, 0x05, 0xa1, 0x05, 0xe4, 0x8e, 0x15, 0xb6, 0x40, 0xe5, 0x10, 0xb3,
	0xe9, 0x76, 0x82, 0xa3, 0x5a, 0x89, 0x5f, 0xb1, 0x31, 0x91, 0xa3, 0x10, 0xb4, 0xca, 0x87, 0x18,
	0xef, 0xb3, 0xa5, 0xfa, 0x1e, 0x00, 0x76, 0x8e, 0x7c, 0xde, 0xd8, 0x8d, 0x7a, 0xa1, 0x51, 0x5d,
	0xbf, 0x63, 0x88, 0xb7, 0x68, 0xb0, 0x19, 0x3d, 0x1a, 0x09, 0x9b, 0xc4, 0x0b, 0xcc, 0x6d, 0x39,
	0x0a, 0x16, 0x33, 0x48, 0x51, 0x0a, 0xbf, 0x9e, 0xeb, 0x8d, 0xbf, 0xb8, 0x0e, 0x0c, 0x85, 0x5a,
	0x4c, 0xe6, 0x53, 0x5e, 0x27, 0xde, 0xb0, 0xf9, 0xec, 0xb4, 0xaf, 0x29, 0x67, 0x7d, 0x4d, 0xf9,
	0xd5, 0xd7, 0x94, 0x4f, 0x17, 0x5a, 0xee, 0xec, 0x42, 0xcb, 0xfd, 0xb8, 0xd0, 0x72, 0x6f, 0x8c,
	0x31, 0x50, 0x31, 0xad, 0xd6, 0x7a, 0xa8, 0x4d, 0xe5, 0xba, 0x99, 0x64, 0x3f, 0x29, 0x4e, 0xd0,
	0x2e, 0xf1, 0xdf, 0xc8, 0xa3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x28, 0x78, 0x92, 0x71, 0xc4,
	0x06, 0x00, 0x00,
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
		size := m.MaxBioLength.Size()
		i -= size
		if _, err := m.MaxBioLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DTagParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.NicknameParams.MarshalToSizedBuffer(dAtA[:i])
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
		size := m.MaxNicknameLength.Size()
		i -= size
		if _, err := m.MaxNicknameLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinNicknameLength.Size()
		i -= size
		if _, err := m.MinNicknameLength.MarshalTo(dAtA[i:]); err != nil {
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
		size := m.MaxDTagLength.Size()
		i -= size
		if _, err := m.MaxDTagLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintModelsParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinDTagLength.Size()
		i -= size
		if _, err := m.MinDTagLength.MarshalTo(dAtA[i:]); err != nil {
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
			dAtA[i] = 0x3a
		}
	}
	if len(m.FeePayer) > 0 {
		i -= len(m.FeePayer)
		copy(dAtA[i:], m.FeePayer)
		i = encodeVarintModelsParams(dAtA, i, uint64(len(m.FeePayer)))
		i--
		dAtA[i] = 0x32
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
	l = m.NicknameParams.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.DTagParams.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.MaxBioLength.Size()
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
	l = m.MinNicknameLength.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.MaxNicknameLength.Size()
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
	l = m.MinDTagLength.Size()
	n += 1 + l + sovModelsParams(uint64(l))
	l = m.MaxDTagLength.Size()
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
	l = len(m.FeePayer)
	if l > 0 {
		n += 1 + l + sovModelsParams(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field NicknameParams", wireType)
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
			if err := m.NicknameParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTagParams", wireType)
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
			if err := m.DTagParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxBioLength", wireType)
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
			if err := m.MaxBioLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field MinNicknameLength", wireType)
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
			if err := m.MinNicknameLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxNicknameLength", wireType)
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
			if err := m.MaxNicknameLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				return fmt.Errorf("proto: wrong wireType = %d for field MinDTagLength", wireType)
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
			if err := m.MinDTagLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDTagLength", wireType)
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
			if err := m.MaxDTagLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
				m.ScriptID |= int64(b&0x7F) << shift
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
				return fmt.Errorf("proto: wrong wireType = %d for field FeePayer", wireType)
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
			m.FeePayer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
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
