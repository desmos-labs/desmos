// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/keeper/types.proto

package keeper

import (
	fmt "fmt"
	types "github.com/desmos-labs/desmos/x/profiles/types"
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

// DtagRequests contains the DTag transfer requests made towards a user
type WrappedDTagTransferRequests struct {
	Requests []types.DTagTransferRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests"`
}

func (m *WrappedDTagTransferRequests) Reset()         { *m = WrappedDTagTransferRequests{} }
func (m *WrappedDTagTransferRequests) String() string { return proto.CompactTextString(m) }
func (*WrappedDTagTransferRequests) ProtoMessage()    {}
func (*WrappedDTagTransferRequests) Descriptor() ([]byte, []int) {
	return fileDescriptor_732b728428b8b385, []int{0}
}
func (m *WrappedDTagTransferRequests) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedDTagTransferRequests) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedDTagTransferRequests.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedDTagTransferRequests) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedDTagTransferRequests.Merge(m, src)
}
func (m *WrappedDTagTransferRequests) XXX_Size() int {
	return m.Size()
}
func (m *WrappedDTagTransferRequests) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedDTagTransferRequests.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedDTagTransferRequests proto.InternalMessageInfo

func (m *WrappedDTagTransferRequests) GetRequests() []types.DTagTransferRequest {
	if m != nil {
		return m.Requests
	}
	return nil
}

func init() {
	proto.RegisterType((*WrappedDTagTransferRequests)(nil), "desmos.profiles.v1beta1.keeper.WrappedDTagTransferRequests")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/keeper/types.proto", fileDescriptor_732b728428b8b385)
}

var fileDescriptor_732b728428b8b385 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4a, 0x49, 0x2d, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0x2d, 0xd6, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0xcf, 0x4e, 0x4d, 0x2d, 0x48, 0x2d, 0xd2, 0x2f, 0xa9, 0x2c, 0x48, 0x2d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x83, 0xa8, 0xd5, 0x83, 0xa9, 0xd5, 0x83, 0xaa,
	0xd5, 0x83, 0xa8, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd5, 0x07, 0xb1, 0x20, 0xba,
	0xa4, 0x54, 0x70, 0xd9, 0x90, 0x9b, 0x9f, 0x92, 0x9a, 0x03, 0x35, 0x5b, 0xa9, 0x9c, 0x4b, 0x3a,
	0xbc, 0x28, 0xb1, 0xa0, 0x20, 0x35, 0xc5, 0x25, 0x24, 0x31, 0x3d, 0xa4, 0x28, 0x31, 0xaf, 0x38,
	0x2d, 0xb5, 0x28, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0xa4, 0x58, 0xc8, 0x8f, 0x8b, 0xa3, 0x08,
	0xca, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xd2, 0xd1, 0xc3, 0xe5, 0x1a, 0x2c, 0x06, 0x38,
	0xb1, 0x9c, 0xb8, 0x27, 0xcf, 0x10, 0x04, 0x37, 0xc3, 0x8a, 0x63, 0xc6, 0x02, 0x79, 0xc6, 0x17,
	0x0b, 0xe4, 0x19, 0x9d, 0x3c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23,
	0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a,
	0x3f, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x62, 0x97, 0x6e, 0x4e,
	0x62, 0x52, 0x31, 0x94, 0xad, 0x5f, 0x81, 0xf0, 0x11, 0xc4, 0xff, 0x49, 0x6c, 0x60, 0xaf, 0x18,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x05, 0x21, 0x73, 0x8c, 0x54, 0x01, 0x00, 0x00,
}

func (this *WrappedDTagTransferRequests) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WrappedDTagTransferRequests)
	if !ok {
		that2, ok := that.(WrappedDTagTransferRequests)
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
	if len(this.Requests) != len(that1.Requests) {
		return false
	}
	for i := range this.Requests {
		if !this.Requests[i].Equal(&that1.Requests[i]) {
			return false
		}
	}
	return true
}
func (m *WrappedDTagTransferRequests) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedDTagTransferRequests) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedDTagTransferRequests) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WrappedDTagTransferRequests) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WrappedDTagTransferRequests) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: WrappedDTagTransferRequests: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedDTagTransferRequests: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, types.DTagTransferRequest{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
