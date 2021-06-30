// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/posts/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params contains the parameters for the posts module
type Params struct {
	MaxPostMessageLength                    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=max_post_message_length,json=maxPostMessageLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_post_message_length" yaml:"max_post_message_length"`
	MaxAdditionalAttributesFieldsNumber     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=max_additional_attributes_fields_number,json=maxAdditionalAttributesFieldsNumber,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_additional_attributes_fields_number" yaml:"max_additional_attributes_fields_number"`
	MaxAdditionalAttributesFieldValueLength github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=max_additional_attributes_field_value_length,json=maxAdditionalAttributesFieldValueLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_additional_attributes_field_value_length" yaml:"max_additional_attributes_field_value_length"`
	MaxAdditionalAttributesFieldKeyLength   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=max_additional_attributes_field_key_length,json=maxAdditionalAttributesFieldKeyLength,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_additional_attributes_field_key_length" yaml:"max_additional_attributes_field_key_length"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_785f2b1a10bf8af9, []int{0}
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

func init() {
	proto.RegisterType((*Params)(nil), "desmos.posts.v1beta1.Params")
}

func init() { proto.RegisterFile("desmos/posts/v1beta1/params.proto", fileDescriptor_785f2b1a10bf8af9) }

var fileDescriptor_785f2b1a10bf8af9 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xbb, 0x4e, 0xf3, 0x30,
	0x14, 0xc7, 0xe3, 0xef, 0xd2, 0x21, 0xfa, 0xa6, 0xa8, 0xd2, 0x57, 0x31, 0xa4, 0x10, 0x04, 0x45,
	0x40, 0x63, 0x55, 0xdd, 0xd8, 0x5a, 0x09, 0x24, 0xc4, 0x45, 0x55, 0x06, 0x06, 0x16, 0xe3, 0x34,
	0x26, 0x8d, 0x1a, 0xc7, 0x51, 0xec, 0x54, 0xc9, 0x13, 0xb0, 0xf2, 0x34, 0x8c, 0x4c, 0x0c, 0x1d,
	0x2b, 0xb1, 0x20, 0x86, 0x0a, 0xb5, 0x6f, 0xc0, 0x13, 0xa0, 0x38, 0xa1, 0x85, 0x81, 0x5e, 0x26,
	0x5b, 0xd6, 0xef, 0x9c, 0xf3, 0xfb, 0x5b, 0x3a, 0xea, 0x96, 0x43, 0x38, 0x65, 0x1c, 0x86, 0x8c,
	0x0b, 0x0e, 0x07, 0x0d, 0x9b, 0x08, 0xdc, 0x80, 0x21, 0x8e, 0x30, 0xe5, 0x66, 0x18, 0x31, 0xc1,
	0xb4, 0x72, 0x8e, 0x98, 0x12, 0x31, 0x0b, 0x64, 0xa3, 0xec, 0x32, 0x97, 0x49, 0x00, 0x66, 0xb7,
	0x9c, 0x35, 0x9e, 0xff, 0xaa, 0xa5, 0x8e, 0x2c, 0xd6, 0xee, 0x80, 0xfa, 0x9f, 0xe2, 0x04, 0x65,
	0x65, 0x88, 0x12, 0xce, 0xb1, 0x4b, 0x90, 0x4f, 0x02, 0x57, 0xf4, 0x2a, 0x60, 0x13, 0xec, 0xfd,
	0x6b, 0x77, 0x86, 0xe3, 0xaa, 0xf2, 0x3a, 0xae, 0xee, 0xba, 0x9e, 0xe8, 0xc5, 0xb6, 0xd9, 0x65,
	0x14, 0x76, 0x99, 0xd4, 0xc9, 0x8f, 0x3a, 0x77, 0xfa, 0x50, 0xa4, 0x21, 0xe1, 0xe6, 0x69, 0x20,
	0xde, 0xc7, 0x55, 0x3d, 0xc5, 0xd4, 0x3f, 0x32, 0x7e, 0x68, 0x6b, 0x58, 0x65, 0x8a, 0x93, 0x0e,
	0xe3, 0xe2, 0x22, 0x7f, 0x3f, 0x97, 0xcf, 0xda, 0x03, 0x50, 0x6b, 0x59, 0x09, 0x76, 0x1c, 0x4f,
	0x78, 0x2c, 0xc0, 0x3e, 0xc2, 0x42, 0x44, 0x9e, 0x1d, 0x0b, 0xc2, 0xd1, 0xad, 0x47, 0x7c, 0x87,
	0xa3, 0x20, 0xa6, 0x36, 0x89, 0x2a, 0xbf, 0xa4, 0xd9, 0xcd, 0xda, 0x66, 0xe6, 0xdc, 0x6c, 0x85,
	0x31, 0x86, 0xb5, 0x4d, 0x71, 0xd2, 0x9a, 0x81, 0xad, 0x19, 0x77, 0x22, 0xb1, 0x4b, 0x49, 0x69,
	0x4f, 0x40, 0x3d, 0x5c, 0xd2, 0x11, 0x0d, 0xb0, 0x1f, 0xcf, 0xfe, 0xf5, 0xb7, 0xb4, 0x27, 0x6b,
	0xdb, 0x37, 0x57, 0xb2, 0xff, 0x36, 0xcb, 0xb0, 0x6a, 0x8b, 0x22, 0x5c, 0x65, 0x68, 0xf1, 0xff,
	0x8f, 0x40, 0xdd, 0x5f, 0xd6, 0xba, 0x4f, 0xd2, 0xcf, 0x10, 0x7f, 0x64, 0x88, 0xee, 0xda, 0x21,
	0x1a, 0xab, 0x85, 0x98, 0x4f, 0x32, 0xac, 0x9d, 0x45, 0x11, 0xce, 0x48, 0x9a, 0x07, 0x68, 0x1f,
	0x0f, 0x27, 0x3a, 0x18, 0x4d, 0x74, 0xf0, 0x36, 0xd1, 0xc1, 0xfd, 0x54, 0x57, 0x46, 0x53, 0x5d,
	0x79, 0x99, 0xea, 0xca, 0xf5, 0xc1, 0x17, 0xbb, 0x7c, 0x4d, 0xea, 0x3e, 0xb6, 0x79, 0x71, 0x87,
	0x49, 0xb1, 0x57, 0x52, 0xd3, 0x2e, 0xc9, 0x1d, 0x69, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x50,
	0x0f, 0xfc, 0xf3, 0x74, 0x03, 0x00, 0x00,
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
		size := m.MaxAdditionalAttributesFieldKeyLength.Size()
		i -= size
		if _, err := m.MaxAdditionalAttributesFieldKeyLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.MaxAdditionalAttributesFieldValueLength.Size()
		i -= size
		if _, err := m.MaxAdditionalAttributesFieldValueLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MaxAdditionalAttributesFieldsNumber.Size()
		i -= size
		if _, err := m.MaxAdditionalAttributesFieldsNumber.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.MaxPostMessageLength.Size()
		i -= size
		if _, err := m.MaxPostMessageLength.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.MaxPostMessageLength.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxAdditionalAttributesFieldsNumber.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxAdditionalAttributesFieldValueLength.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MaxAdditionalAttributesFieldKeyLength.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPostMessageLength", wireType)
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
			if err := m.MaxPostMessageLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAdditionalAttributesFieldsNumber", wireType)
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
			if err := m.MaxAdditionalAttributesFieldsNumber.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAdditionalAttributesFieldValueLength", wireType)
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
			if err := m.MaxAdditionalAttributesFieldValueLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAdditionalAttributesFieldKeyLength", wireType)
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
			if err := m.MaxAdditionalAttributesFieldKeyLength.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
