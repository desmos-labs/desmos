// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/subspaces/v3/msgs_feegrant.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgGrantAllowance adds grants for the grantee to spend up allowance of fees
// from the treasury inside the given subspace
type MsgGrantAllowance struct {
	// Id of the subspace inside which where the allowance should be granted
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty" yaml:"subspace_id"`
	// Address of the user granting the allowance
	Granter string `protobuf:"bytes,2,opt,name=granter,proto3" json:"granter,omitempty" yaml:"granter"`
	// Target being granted the allowance
	Grantee *types.Any `protobuf:"bytes,3,opt,name=grantee,proto3" json:"grantee,omitempty" yaml:"grantee"`
	// Allowance can be any allowance type that implements AllowanceI
	Allowance *types.Any `protobuf:"bytes,4,opt,name=allowance,proto3" json:"allowance,omitempty" yaml:"allowance"`
}

func (m *MsgGrantAllowance) Reset()         { *m = MsgGrantAllowance{} }
func (m *MsgGrantAllowance) String() string { return proto.CompactTextString(m) }
func (*MsgGrantAllowance) ProtoMessage()    {}
func (*MsgGrantAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9368a6a8079b960, []int{0}
}
func (m *MsgGrantAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantAllowance.Merge(m, src)
}
func (m *MsgGrantAllowance) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantAllowance proto.InternalMessageInfo

func (m *MsgGrantAllowance) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *MsgGrantAllowance) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *MsgGrantAllowance) GetGrantee() *types.Any {
	if m != nil {
		return m.Grantee
	}
	return nil
}

func (m *MsgGrantAllowance) GetAllowance() *types.Any {
	if m != nil {
		return m.Allowance
	}
	return nil
}

// MsgGrantAllowanceResponse defines the Msg/GrantAllowanceResponse response
// type.
type MsgGrantAllowanceResponse struct {
}

func (m *MsgGrantAllowanceResponse) Reset()         { *m = MsgGrantAllowanceResponse{} }
func (m *MsgGrantAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGrantAllowanceResponse) ProtoMessage()    {}
func (*MsgGrantAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9368a6a8079b960, []int{1}
}
func (m *MsgGrantAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantAllowanceResponse.Merge(m, src)
}
func (m *MsgGrantAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantAllowanceResponse proto.InternalMessageInfo

// MsgRevokeAllowance removes any existing allowance to the grantee inside the
// subspace
type MsgRevokeAllowance struct {
	// If of the subspace inside which the allowance to be deleted is
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty" yaml:"subspace_id"`
	// Address of the user that created the allowance
	Granter string `protobuf:"bytes,2,opt,name=granter,proto3" json:"granter,omitempty" yaml:"granter"`
	// Target being revoked the allowance
	Grantee *types.Any `protobuf:"bytes,3,opt,name=grantee,proto3" json:"grantee,omitempty" yaml:"grantee"`
}

func (m *MsgRevokeAllowance) Reset()         { *m = MsgRevokeAllowance{} }
func (m *MsgRevokeAllowance) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeAllowance) ProtoMessage()    {}
func (*MsgRevokeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9368a6a8079b960, []int{2}
}
func (m *MsgRevokeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeAllowance.Merge(m, src)
}
func (m *MsgRevokeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeAllowance proto.InternalMessageInfo

func (m *MsgRevokeAllowance) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *MsgRevokeAllowance) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *MsgRevokeAllowance) GetGrantee() *types.Any {
	if m != nil {
		return m.Grantee
	}
	return nil
}

// MsgRevokeAllowanceResponse defines the Msg/RevokeAllowanceResponse
// response type.
type MsgRevokeAllowanceResponse struct {
}

func (m *MsgRevokeAllowanceResponse) Reset()         { *m = MsgRevokeAllowanceResponse{} }
func (m *MsgRevokeAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeAllowanceResponse) ProtoMessage()    {}
func (*MsgRevokeAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9368a6a8079b960, []int{3}
}
func (m *MsgRevokeAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeAllowanceResponse.Merge(m, src)
}
func (m *MsgRevokeAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeAllowanceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgGrantAllowance)(nil), "desmos.subspaces.v3.MsgGrantAllowance")
	proto.RegisterType((*MsgGrantAllowanceResponse)(nil), "desmos.subspaces.v3.MsgGrantAllowanceResponse")
	proto.RegisterType((*MsgRevokeAllowance)(nil), "desmos.subspaces.v3.MsgRevokeAllowance")
	proto.RegisterType((*MsgRevokeAllowanceResponse)(nil), "desmos.subspaces.v3.MsgRevokeAllowanceResponse")
}

func init() {
	proto.RegisterFile("desmos/subspaces/v3/msgs_feegrant.proto", fileDescriptor_c9368a6a8079b960)
}

var fileDescriptor_c9368a6a8079b960 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x93, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x9b, 0xba, 0x28, 0x3b, 0x0b, 0xa2, 0xb1, 0x87, 0x6e, 0x57, 0x92, 0x32, 0x08, 0x16,
	0xa1, 0x33, 0x74, 0x8b, 0x97, 0xbd, 0x35, 0xac, 0x2c, 0x05, 0xf7, 0x92, 0xbd, 0x79, 0x29, 0x93,
	0xe4, 0xdb, 0xb1, 0x98, 0x64, 0x42, 0x26, 0x8d, 0xe6, 0xee, 0x03, 0xf8, 0x30, 0xfb, 0x10, 0xe2,
	0x69, 0xf1, 0xe4, 0x29, 0x48, 0xfa, 0x06, 0xf5, 0x05, 0x24, 0x99, 0x4c, 0x5c, 0xac, 0xf8, 0x02,
	0x7b, 0x9b, 0x2f, 0xff, 0xef, 0xfb, 0x7e, 0xf3, 0xff, 0x0f, 0x41, 0x2f, 0x03, 0x90, 0x91, 0x90,
	0x54, 0x6e, 0x3c, 0x99, 0x30, 0x1f, 0x24, 0xcd, 0xe7, 0x34, 0x92, 0x5c, 0xae, 0xae, 0x01, 0x78,
	0xca, 0xe2, 0x8c, 0x24, 0xa9, 0xc8, 0x84, 0xf9, 0x4c, 0x35, 0x92, 0xae, 0x91, 0xe4, 0xf3, 0xd1,
	0x80, 0x0b, 0x2e, 0x1a, 0x9d, 0xd6, 0x27, 0xd5, 0x3a, 0x3a, 0xe6, 0x42, 0xf0, 0x10, 0x68, 0x53,
	0x79, 0x9b, 0x6b, 0xca, 0xe2, 0x42, 0x4b, 0xbe, 0xa8, 0xb7, 0xac, 0xd4, 0x8c, 0x2a, 0x94, 0x84,
	0x7f, 0xf5, 0xd1, 0xd3, 0x4b, 0xc9, 0x2f, 0x6a, 0xe6, 0x22, 0x0c, 0xc5, 0x47, 0x16, 0xfb, 0x60,
	0xbe, 0x41, 0x47, 0x9a, 0xb8, 0x5a, 0x07, 0x43, 0x63, 0x6c, 0x4c, 0x0e, 0x9c, 0x17, 0x55, 0x69,
	0xa3, 0xab, 0xf6, 0xf3, 0xf2, 0x7c, 0x57, 0xda, 0x66, 0xc1, 0xa2, 0xf0, 0x0c, 0xdf, 0x69, 0xc5,
	0x2e, 0xd2, 0xd5, 0x32, 0x30, 0xcf, 0xd1, 0xa3, 0xc6, 0x0c, 0xa4, 0xc3, 0xfe, 0xd8, 0x98, 0x1c,
	0x3a, 0xaf, 0x76, 0xa5, 0xfd, 0x58, 0x0d, 0xb5, 0x02, 0xfe, 0x7e, 0x33, 0x1d, 0xb4, 0x37, 0x5a,
	0x04, 0x41, 0x0a, 0x52, 0x5e, 0x65, 0xe9, 0x3a, 0xe6, 0xae, 0x1e, 0x35, 0x99, 0xde, 0x02, 0xc3,
	0x07, 0x63, 0x63, 0x72, 0x74, 0x3a, 0x20, 0xca, 0x2a, 0xd1, 0x56, 0xc9, 0x22, 0x2e, 0x9c, 0xd9,
	0xdf, 0xbb, 0x01, 0x7f, 0xbb, 0x99, 0x9e, 0xfc, 0x23, 0x40, 0x72, 0xa1, 0x74, 0x8d, 0x00, 0x33,
	0x46, 0x87, 0x4c, 0x9b, 0x1f, 0x1e, 0xfc, 0x07, 0x72, 0xb6, 0x2b, 0xed, 0x27, 0x0a, 0xd2, 0x0d,
	0xd4, 0x18, 0xdc, 0x5a, 0xe8, 0x9e, 0x2f, 0x9f, 0x79, 0x90, 0xb1, 0x19, 0xe9, 0x32, 0x5d, 0xba,
	0x7f, 0x10, 0xf8, 0x04, 0x1d, 0xef, 0x85, 0xee, 0x82, 0x4c, 0x44, 0x2c, 0x01, 0x7f, 0xee, 0x23,
	0xf3, 0x52, 0x72, 0x17, 0x72, 0xf1, 0x01, 0xee, 0xed, 0x9b, 0xe0, 0xe7, 0x68, 0xb4, 0x9f, 0x82,
	0x0e, 0xc9, 0x79, 0xfb, 0xb5, 0xb2, 0x8c, 0xdb, 0xca, 0x32, 0x7e, 0x56, 0x96, 0xf1, 0x65, 0x6b,
	0xf5, 0x6e, 0xb7, 0x56, 0xef, 0xc7, 0xd6, 0xea, 0xbd, 0x3b, 0xe5, 0xeb, 0xec, 0xfd, 0xc6, 0x23,
	0xbe, 0x88, 0xa8, 0x02, 0x4d, 0x43, 0xe6, 0xc9, 0xf6, 0x4c, 0xf3, 0xd7, 0xf4, 0xd3, 0x9d, 0xff,
	0x2e, 0x2b, 0x12, 0x90, 0xde, 0xc3, 0xe6, 0xd6, 0xf3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3a,
	0x1a, 0x7c, 0xb7, 0x98, 0x03, 0x00, 0x00,
}

func (m *MsgGrantAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgsFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Grantee != nil {
		{
			size, err := m.Grantee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgsFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintMsgsFeegrant(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x12
	}
	if m.SubspaceID != 0 {
		i = encodeVarintMsgsFeegrant(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgGrantAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRevokeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Grantee != nil {
		{
			size, err := m.Grantee.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMsgsFeegrant(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintMsgsFeegrant(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0x12
	}
	if m.SubspaceID != 0 {
		i = encodeVarintMsgsFeegrant(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsFeegrant(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsFeegrant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgGrantAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovMsgsFeegrant(uint64(m.SubspaceID))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovMsgsFeegrant(uint64(l))
	}
	if m.Grantee != nil {
		l = m.Grantee.Size()
		n += 1 + l + sovMsgsFeegrant(uint64(l))
	}
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovMsgsFeegrant(uint64(l))
	}
	return n
}

func (m *MsgGrantAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRevokeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovMsgsFeegrant(uint64(m.SubspaceID))
	}
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovMsgsFeegrant(uint64(l))
	}
	if m.Grantee != nil {
		l = m.Grantee.Size()
		n += 1 + l + sovMsgsFeegrant(uint64(l))
	}
	return n
}

func (m *MsgRevokeAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsFeegrant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsFeegrant(x uint64) (n int) {
	return sovMsgsFeegrant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgGrantAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsFeegrant
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
			return fmt.Errorf("proto: MsgGrantAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsFeegrant
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
				return ErrInvalidLengthMsgsFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsFeegrant
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
				return ErrInvalidLengthMsgsFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Grantee == nil {
				m.Grantee = &types.Any{}
			}
			if err := m.Grantee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsFeegrant
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
				return ErrInvalidLengthMsgsFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Allowance == nil {
				m.Allowance = &types.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsFeegrant
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
func (m *MsgGrantAllowanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsFeegrant
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
			return fmt.Errorf("proto: MsgGrantAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsFeegrant
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
func (m *MsgRevokeAllowance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsFeegrant
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
			return fmt.Errorf("proto: MsgRevokeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsFeegrant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsFeegrant
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
				return ErrInvalidLengthMsgsFeegrant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsFeegrant
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
				return ErrInvalidLengthMsgsFeegrant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsFeegrant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Grantee == nil {
				m.Grantee = &types.Any{}
			}
			if err := m.Grantee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsFeegrant
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
func (m *MsgRevokeAllowanceResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsFeegrant
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
			return fmt.Errorf("proto: MsgRevokeAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsFeegrant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsFeegrant
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
func skipMsgsFeegrant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsFeegrant
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
					return 0, ErrIntOverflowMsgsFeegrant
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
					return 0, ErrIntOverflowMsgsFeegrant
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
				return 0, ErrInvalidLengthMsgsFeegrant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsFeegrant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsFeegrant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsFeegrant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsFeegrant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsFeegrant = fmt.Errorf("proto: unexpected end of group")
)
