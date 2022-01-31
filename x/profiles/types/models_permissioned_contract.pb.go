// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/models_permissioned_contract.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// PermissionedContract represents a reference to a permissioned contract that runs automatically on chain
type PermissionedContract struct {
	Address  string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	Admin    string   `protobuf:"bytes,2,opt,name=admin,proto3" json:"admin,omitempty" yaml:"admin"`
	Messages [][]byte `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty" yaml:"messages"`
}

func (m *PermissionedContract) Reset()         { *m = PermissionedContract{} }
func (m *PermissionedContract) String() string { return proto.CompactTextString(m) }
func (*PermissionedContract) ProtoMessage()    {}
func (*PermissionedContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_650ff3f38f46759a, []int{0}
}
func (m *PermissionedContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PermissionedContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PermissionedContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PermissionedContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PermissionedContract.Merge(m, src)
}
func (m *PermissionedContract) XXX_Size() int {
	return m.Size()
}
func (m *PermissionedContract) XXX_DiscardUnknown() {
	xxx_messageInfo_PermissionedContract.DiscardUnknown(m)
}

var xxx_messageInfo_PermissionedContract proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PermissionedContract)(nil), "desmos.profiles.v1beta1.PermissionedContract")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/models_permissioned_contract.proto", fileDescriptor_650ff3f38f46759a)
}

var fileDescriptor_650ff3f38f46759a = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x31, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x63, 0x2a, 0xa0, 0x44, 0x15, 0xa0, 0x50, 0x89, 0xd2, 0x21, 0xa9, 0x32, 0xa0, 0x0e,
	0x10, 0x53, 0xd8, 0x3a, 0x96, 0x91, 0x05, 0x75, 0x64, 0x89, 0x9c, 0xc4, 0x4d, 0x2d, 0x62, 0x3b,
	0xca, 0x73, 0x2b, 0x7a, 0x03, 0x46, 0x8e, 0xd0, 0x8d, 0xab, 0x30, 0x76, 0x64, 0xaa, 0x50, 0xbb,
	0x30, 0xf7, 0x04, 0xa8, 0xb1, 0x03, 0x88, 0x85, 0xed, 0x3d, 0x7d, 0xdf, 0x6f, 0xeb, 0x7f, 0x76,
	0x3f, 0xa1, 0xc0, 0x25, 0xe0, 0xbc, 0x90, 0x23, 0x96, 0x51, 0xc0, 0xd3, 0x5e, 0x44, 0x15, 0xe9,
	0x61, 0x2e, 0x13, 0x9a, 0x41, 0x98, 0xd3, 0x82, 0x33, 0x00, 0x26, 0x05, 0x4d, 0xc2, 0x58, 0x0a,
	0x55, 0x90, 0x58, 0x05, 0x79, 0x21, 0x95, 0x74, 0x4e, 0x75, 0x36, 0xa8, 0xb2, 0x81, 0xc9, 0xb6,
	0x9b, 0xa9, 0x4c, 0x65, 0xe9, 0xe0, 0xed, 0xa4, 0xf5, 0xf6, 0x59, 0x2a, 0x65, 0x9a, 0x51, 0x5c,
	0x6e, 0xd1, 0x64, 0x84, 0x89, 0x98, 0x19, 0xe4, 0xfd, 0x45, 0x8a, 0x71, 0x0a, 0x8a, 0xf0, 0xbc,
	0xca, 0xc6, 0x72, 0xfb, 0x55, 0xa8, 0x1f, 0xd5, 0x8b, 0x41, 0x57, 0xff, 0x34, 0x88, 0xc7, 0x84,
	0x89, 0x30, 0x63, 0xe2, 0xd1, 0x24, 0xfc, 0x57, 0x64, 0x37, 0xef, 0x7f, 0xf5, 0xba, 0x35, 0xb5,
	0x9c, 0x0b, 0x7b, 0x9f, 0x24, 0x49, 0x41, 0x01, 0x5a, 0xa8, 0x83, 0xba, 0x07, 0x03, 0x67, 0xb3,
	0xf4, 0x0e, 0x67, 0x84, 0x67, 0x7d, 0xdf, 0x00, 0x7f, 0x58, 0x29, 0xce, 0xb9, 0xbd, 0x4b, 0x12,
	0xce, 0x44, 0x6b, 0xa7, 0x74, 0x8f, 0x37, 0x4b, 0xaf, 0x51, 0xb9, 0x9c, 0x09, 0x7f, 0xa8, 0xb1,
	0x83, 0xed, 0x3a, 0xa7, 0x00, 0x24, 0xa5, 0xd0, 0xaa, 0x75, 0x6a, 0xdd, 0xc6, 0xe0, 0x64, 0xb3,
	0xf4, 0x8e, 0xb4, 0x5a, 0x11, 0x7f, 0xf8, 0x2d, 0xf5, 0xeb, 0xcf, 0x73, 0xcf, 0xfa, 0x9c, 0x7b,
	0x68, 0x70, 0xf7, 0xb6, 0x72, 0xd1, 0x62, 0xe5, 0xa2, 0x8f, 0x95, 0x8b, 0x5e, 0xd6, 0xae, 0xb5,
	0x58, 0xbb, 0xd6, 0xfb, 0xda, 0xb5, 0x1e, 0x7a, 0x29, 0x53, 0xe3, 0x49, 0x14, 0xc4, 0x92, 0x63,
	0x7d, 0x80, 0xcb, 0x8c, 0x44, 0x60, 0x66, 0x3c, 0xbd, 0xc6, 0x4f, 0x3f, 0x17, 0x51, 0xb3, 0x9c,
	0x42, 0xb4, 0x57, 0xb6, 0xbf, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x43, 0x32, 0x18, 0x2a, 0xf3,
	0x01, 0x00, 0x00,
}

func (this *PermissionedContract) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PermissionedContract)
	if !ok {
		that2, ok := that.(PermissionedContract)
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
	if this.Address != that1.Address {
		return false
	}
	if this.Admin != that1.Admin {
		return false
	}
	if len(this.Messages) != len(that1.Messages) {
		return false
	}
	for i := range this.Messages {
		if !bytes.Equal(this.Messages[i], that1.Messages[i]) {
			return false
		}
	}
	return true
}
func (m *PermissionedContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PermissionedContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PermissionedContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Messages[iNdEx])
			copy(dAtA[i:], m.Messages[iNdEx])
			i = encodeVarintModelsPermissionedContract(dAtA, i, uint64(len(m.Messages[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintModelsPermissionedContract(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintModelsPermissionedContract(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsPermissionedContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsPermissionedContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PermissionedContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovModelsPermissionedContract(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovModelsPermissionedContract(uint64(l))
	}
	if len(m.Messages) > 0 {
		for _, b := range m.Messages {
			l = len(b)
			n += 1 + l + sovModelsPermissionedContract(uint64(l))
		}
	}
	return n
}

func sovModelsPermissionedContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsPermissionedContract(x uint64) (n int) {
	return sovModelsPermissionedContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PermissionedContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsPermissionedContract
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
			return fmt.Errorf("proto: PermissionedContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PermissionedContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPermissionedContract
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
				return ErrInvalidLengthModelsPermissionedContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPermissionedContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPermissionedContract
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
				return ErrInvalidLengthModelsPermissionedContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPermissionedContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPermissionedContract
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
				return ErrInvalidLengthModelsPermissionedContract
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPermissionedContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, make([]byte, postIndex-iNdEx))
			copy(m.Messages[len(m.Messages)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsPermissionedContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsPermissionedContract
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
func skipModelsPermissionedContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsPermissionedContract
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
					return 0, ErrIntOverflowModelsPermissionedContract
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
					return 0, ErrIntOverflowModelsPermissionedContract
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
				return 0, ErrInvalidLengthModelsPermissionedContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsPermissionedContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsPermissionedContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsPermissionedContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsPermissionedContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsPermissionedContract = fmt.Errorf("proto: unexpected end of group")
)
