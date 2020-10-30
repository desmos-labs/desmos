// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/params.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params contains the parameters for the relationships module
type Params struct {
	MonikerParams MonikerParams                          `protobuf:"bytes,1,opt,name=moniker_params,json=monikerParams,proto3" json:"moniker_params" yaml:"moniker_params"`
	DtagParams    DTagParams                             `protobuf:"bytes,2,opt,name=dtag_params,json=dtagParams,proto3" json:"dtag_params" yaml:"dtag_params"`
	MaxBioLength  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_bio_length,json=maxBioLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_bio_length" yaml:"max_bio_length"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_821862d20041ec2d, []int{0}
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

// MonikerParams defines the parameters related to the profiles monikers
type MonikerParams struct {
	MinMonikerLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=min_moniker_length,json=minMonikerLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_length" yaml:"min_length"`
	MaxMonikerLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=max_moniker_length,json=maxMonikerLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_length" yaml:"max_length"`
}

func (m *MonikerParams) Reset()      { *m = MonikerParams{} }
func (*MonikerParams) ProtoMessage() {}
func (*MonikerParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_821862d20041ec2d, []int{1}
}
func (m *MonikerParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MonikerParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MonikerParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MonikerParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonikerParams.Merge(m, src)
}
func (m *MonikerParams) XXX_Size() int {
	return m.Size()
}
func (m *MonikerParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MonikerParams.DiscardUnknown(m)
}

var xxx_messageInfo_MonikerParams proto.InternalMessageInfo

// DTagParams defines the parameters related to profile DTags
type DTagParams struct {
	RegEx         string                                 `protobuf:"bytes,1,opt,name=reg_ex,json=regEx,proto3" json:"reg_ex,omitempty" yaml:"reg_ex"`
	MinDtagLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_dtag_length,json=minDtagLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_length" yaml:"min_length"`
	MaxDtagLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_dtag_length,json=maxDtagLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_length" yaml:"max_length"`
}

func (m *DTagParams) Reset()      { *m = DTagParams{} }
func (*DTagParams) ProtoMessage() {}
func (*DTagParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_821862d20041ec2d, []int{2}
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

func init() {
	proto.RegisterType((*Params)(nil), "desmos.profiles.v1beta1.Params")
	proto.RegisterType((*MonikerParams)(nil), "desmos.profiles.v1beta1.MonikerParams")
	proto.RegisterType((*DTagParams)(nil), "desmos.profiles.v1beta1.DTagParams")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/params.proto", fileDescriptor_821862d20041ec2d)
}

var fileDescriptor_821862d20041ec2d = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x31, 0x6f, 0xd4, 0x30,
	0x18, 0x86, 0x93, 0x54, 0x9c, 0x84, 0xdb, 0x14, 0x1a, 0x81, 0xa8, 0x18, 0x9c, 0x2a, 0xa0, 0xea,
	0x18, 0x1a, 0xab, 0xb0, 0x31, 0x46, 0x65, 0x40, 0x02, 0x09, 0xa2, 0x2e, 0xb0, 0x44, 0x4e, 0xcf,
	0xb8, 0x56, 0xcf, 0xf1, 0x29, 0x31, 0xc8, 0xb7, 0x50, 0x7e, 0x02, 0xfc, 0x0f, 0x7e, 0xc8, 0x8d,
	0x1d, 0x2b, 0x86, 0x88, 0xde, 0x6d, 0x8c, 0x37, 0x32, 0xa1, 0xd8, 0xbe, 0xbb, 0xa4, 0xa8, 0x42,
	0x80, 0x3a, 0xc5, 0xfe, 0xf4, 0x7d, 0xef, 0xfb, 0xf8, 0xb5, 0x62, 0xf0, 0x70, 0x40, 0x2a, 0x2e,
	0x2a, 0x34, 0x2a, 0xc5, 0x3b, 0x36, 0x24, 0x15, 0xfa, 0xb0, 0x9f, 0x13, 0x89, 0xf7, 0xd1, 0x08,
	0x97, 0x98, 0x57, 0xf1, 0xa8, 0x14, 0x52, 0x04, 0xf7, 0x4c, 0x57, 0xbc, 0xe8, 0x8a, 0x6d, 0xd7,
	0xfd, 0x3b, 0x54, 0x50, 0xa1, 0x7b, 0x50, 0xb3, 0x32, 0xed, 0xd1, 0x4f, 0x0f, 0xf4, 0x5e, 0xe9,
	0xf9, 0xe0, 0x23, 0xd8, 0xe4, 0xa2, 0x60, 0x27, 0xa4, 0xcc, 0x8c, 0xe2, 0xb6, 0xbb, 0xe3, 0xf6,
	0xd7, 0x1f, 0xef, 0xc6, 0x57, 0x48, 0xc6, 0x2f, 0x4d, 0xbb, 0x99, 0x4f, 0xd0, 0xa4, 0x0e, 0x9d,
	0x1f, 0x75, 0x78, 0x49, 0x65, 0x5e, 0x87, 0x77, 0xc7, 0x98, 0x0f, 0x9f, 0x46, 0xdd, 0x7a, 0x94,
	0xfa, 0xbc, 0x3d, 0x1f, 0x08, 0xb0, 0x3e, 0x90, 0x98, 0x2e, 0xcc, 0x3d, 0x6d, 0xfe, 0xe0, 0x4a,
	0xf3, 0x83, 0x43, 0x4c, 0xad, 0xf3, 0x23, 0xeb, 0xdc, 0x9e, 0x9f, 0xd7, 0x61, 0x60, 0x6c, 0x5b,
	0xc5, 0x28, 0x05, 0xcd, 0xce, 0x1a, 0x9e, 0x82, 0x4d, 0x8e, 0x55, 0x96, 0x33, 0x91, 0x0d, 0x49,
	0x41, 0xe5, 0xf1, 0xf6, 0xda, 0x8e, 0xdb, 0xdf, 0x48, 0xde, 0x34, 0x72, 0xdf, 0xea, 0x70, 0x97,
	0x32, 0x79, 0xfc, 0x3e, 0x8f, 0x8f, 0x04, 0x47, 0x47, 0x42, 0x67, 0x6f, 0x3e, 0x7b, 0xd5, 0xe0,
	0x04, 0xc9, 0xf1, 0x88, 0x54, 0xf1, 0xf3, 0x42, 0xea, 0x23, 0x77, 0x74, 0x5a, 0x47, 0xee, 0xd4,
	0xa3, 0x74, 0x83, 0x63, 0x95, 0x30, 0xf1, 0xc2, 0x6c, 0xbf, 0x78, 0xc0, 0xef, 0x64, 0x18, 0x9c,
	0x82, 0x80, 0xb3, 0x22, 0x5b, 0x24, 0x65, 0xb1, 0x5c, 0x8d, 0xf5, 0xfa, 0xaf, 0xb1, 0x40, 0xa3,
	0xb5, 0x44, 0xda, 0xb2, 0x48, 0xcb, 0x5a, 0x94, 0xde, 0xe6, 0xac, 0xb0, 0x00, 0x06, 0x49, 0x03,
	0x60, 0x75, 0x19, 0xc0, 0xfb, 0x67, 0x00, 0xac, 0x7e, 0x07, 0x58, 0xd6, 0x1a, 0x00, 0xac, 0x3a,
	0x00, 0xd1, 0x57, 0x0f, 0x80, 0xd5, 0xd5, 0x06, 0x7d, 0xd0, 0x2b, 0x09, 0xcd, 0x88, 0xd2, 0x21,
	0xdc, 0x4c, 0xb6, 0xe6, 0x75, 0xe8, 0x1b, 0x1d, 0x53, 0x8f, 0xd2, 0x1b, 0x25, 0xa1, 0xcf, 0x54,
	0x30, 0x06, 0xb7, 0x9a, 0xa3, 0xe9, 0xdb, 0xfe, 0x6f, 0xec, 0x3f, 0xe4, 0xe6, 0x73, 0x56, 0x1c,
	0x48, 0x4c, 0x6d, 0x68, 0x8d, 0x35, 0x56, 0x1d, 0xeb, 0xb5, 0xeb, 0x4a, 0xcc, 0xe7, 0x58, 0xad,
	0xac, 0x93, 0xc3, 0xc9, 0x05, 0x74, 0xce, 0x2f, 0xa0, 0xf3, 0x69, 0x0a, 0x9d, 0xc9, 0x14, 0xba,
	0x67, 0x53, 0xe8, 0x7e, 0x9f, 0x42, 0xf7, 0xf3, 0x0c, 0x3a, 0x67, 0x33, 0xe8, 0x9c, 0xcf, 0xa0,
	0xf3, 0x36, 0x6e, 0xf9, 0x9b, 0xff, 0x69, 0x6f, 0x88, 0xf3, 0xca, 0xae, 0x91, 0x5a, 0xbd, 0x29,
	0x9a, 0x25, 0xef, 0xe9, 0xc7, 0xe1, 0xc9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x92, 0x9f,
	0x98, 0x73, 0x04, 0x00, 0x00,
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
		size := m.MaxBioLength.Size()
		i -= size
		if _, err := m.MaxBioLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.DtagParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.MonikerParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MonikerParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonikerParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MonikerParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxMonikerLength.Size()
		i -= size
		if _, err := m.MaxMonikerLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MinMonikerLength.Size()
		i -= size
		if _, err := m.MinMonikerLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
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
		size := m.MaxDtagLength.Size()
		i -= size
		if _, err := m.MaxDtagLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinDtagLength.Size()
		i -= size
		if _, err := m.MinDtagLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.RegEx) > 0 {
		i -= len(m.RegEx)
		copy(dAtA[i:], m.RegEx)
		i = encodeVarintParams(dAtA, i, uint64(len(m.RegEx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	l = m.MonikerParams.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.DtagParams.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxBioLength.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *MonikerParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MinMonikerLength.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxMonikerLength.Size()
	n += 1 + l + sovParams(uint64(l))
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
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.MinDtagLength.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxDtagLength.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field MonikerParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MonikerParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DtagParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DtagParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxBioLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *MonikerParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: MonikerParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonikerParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinMonikerLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinMonikerLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMonikerLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxMonikerLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthParams
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
				return ErrIntOverflowParams
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
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RegEx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDtagLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinDtagLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDtagLength", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxDtagLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthParams
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
