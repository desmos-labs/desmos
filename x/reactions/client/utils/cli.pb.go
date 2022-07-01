// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/reactions/v1/client/cli.proto

package utils

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/desmos-labs/desmos/v4/x/reactions/types"
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

// SetReactionsParamsJSON contains the data that can be specified when setting a
// subspace reactions params using the CLI command
type SetReactionsParamsJSON struct {
	// Params related to RegisteredReactionValue reactions
	RegisteredReactionParams types.RegisteredReactionValueParams `protobuf:"bytes,1,opt,name=registered_reaction_params,json=registeredReactionParams,proto3" json:"registered_reaction_params"`
	// Params related to FreeTextValue reactions
	FreeTextParams types.FreeTextValueParams `protobuf:"bytes,2,opt,name=free_text_params,json=freeTextParams,proto3" json:"free_text_params"`
}

func (m *SetReactionsParamsJSON) Reset()         { *m = SetReactionsParamsJSON{} }
func (m *SetReactionsParamsJSON) String() string { return proto.CompactTextString(m) }
func (*SetReactionsParamsJSON) ProtoMessage()    {}
func (*SetReactionsParamsJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_aecb29e08714fd3f, []int{0}
}
func (m *SetReactionsParamsJSON) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetReactionsParamsJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetReactionsParamsJSON.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetReactionsParamsJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetReactionsParamsJSON.Merge(m, src)
}
func (m *SetReactionsParamsJSON) XXX_Size() int {
	return m.Size()
}
func (m *SetReactionsParamsJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_SetReactionsParamsJSON.DiscardUnknown(m)
}

var xxx_messageInfo_SetReactionsParamsJSON proto.InternalMessageInfo

func (m *SetReactionsParamsJSON) GetRegisteredReactionParams() types.RegisteredReactionValueParams {
	if m != nil {
		return m.RegisteredReactionParams
	}
	return types.RegisteredReactionValueParams{}
}

func (m *SetReactionsParamsJSON) GetFreeTextParams() types.FreeTextValueParams {
	if m != nil {
		return m.FreeTextParams
	}
	return types.FreeTextValueParams{}
}

func init() {
	proto.RegisterType((*SetReactionsParamsJSON)(nil), "desmos.reactions.v1.client.SetReactionsParamsJSON")
}

func init() {
	proto.RegisterFile("desmos/reactions/v1/client/cli.proto", fileDescriptor_aecb29e08714fd3f)
}

var fileDescriptor_aecb29e08714fd3f = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xfb, 0x30,
	0x10, 0xc6, 0x93, 0xbf, 0xfe, 0x62, 0x08, 0x12, 0x42, 0x15, 0x42, 0x25, 0x43, 0x5a, 0x21, 0x86,
	0x2e, 0xd8, 0x6a, 0x61, 0x61, 0x65, 0x60, 0x60, 0x00, 0xd4, 0x22, 0x84, 0x58, 0x22, 0x27, 0xbd,
	0x1a, 0x4b, 0x76, 0x5d, 0xd9, 0x97, 0xa8, 0xbc, 0x05, 0x8f, 0xd5, 0xb1, 0x23, 0x13, 0x42, 0xed,
	0xc6, 0x53, 0xa0, 0xc6, 0x4e, 0x85, 0x68, 0xa6, 0xdc, 0xe5, 0xfb, 0xdd, 0xf7, 0xd9, 0xe7, 0xe8,
	0x6c, 0x0c, 0x56, 0x69, 0x4b, 0x0d, 0xb0, 0x1c, 0x85, 0x9e, 0x5a, 0x5a, 0xf6, 0x69, 0x2e, 0x05,
	0x4c, 0x71, 0xf3, 0x21, 0x33, 0xa3, 0x51, 0xb7, 0x62, 0x47, 0x91, 0x2d, 0x45, 0xca, 0x3e, 0x71,
	0x54, 0x7c, 0xc4, 0x35, 0xd7, 0x15, 0x46, 0x37, 0x95, 0x9b, 0x88, 0x4f, 0xb8, 0xd6, 0x5c, 0x02,
	0xad, 0xba, 0xac, 0x98, 0x50, 0x36, 0x7d, 0xf3, 0x52, 0xe7, 0xaf, 0x84, 0x42, 0x81, 0x45, 0xa6,
	0x66, 0xf5, 0x6c, 0xae, 0x37, 0x69, 0xa9, 0x33, 0x75, 0x8d, 0x97, 0xba, 0x4d, 0xc7, 0x55, 0x7a,
	0x0c, 0xd2, 0x13, 0xa7, 0xdf, 0x61, 0x74, 0x3c, 0x02, 0x1c, 0xd6, 0xc0, 0x03, 0x33, 0x4c, 0xd9,
	0xdb, 0xd1, 0xfd, 0x5d, 0xab, 0x8c, 0x62, 0x03, 0x5c, 0x58, 0x04, 0x03, 0xe3, 0xb4, 0xb6, 0x48,
	0x67, 0x15, 0xd1, 0x0e, 0xbb, 0x61, 0x6f, 0x7f, 0x30, 0x20, 0x4d, 0x57, 0x1d, 0x6e, 0xc7, 0x6a,
	0xdf, 0x27, 0x26, 0x0b, 0x70, 0xde, 0xd7, 0xff, 0x17, 0x9f, 0x9d, 0x60, 0xd8, 0x36, 0x3b, 0x90,
	0xd3, 0x5b, 0xcf, 0xd1, 0xe1, 0xc4, 0x00, 0xa4, 0x08, 0x73, 0xac, 0xd3, 0xfe, 0x55, 0x69, 0xbd,
	0xc6, 0xb4, 0x1b, 0x03, 0xf0, 0x08, 0x73, 0xdc, 0xcd, 0x38, 0x98, 0x78, 0xc9, 0xff, 0x1d, 0x2d,
	0x56, 0x49, 0xb8, 0x5c, 0x25, 0xe1, 0xd7, 0x2a, 0x09, 0xdf, 0xd7, 0x49, 0xb0, 0x5c, 0x27, 0xc1,
	0xc7, 0x3a, 0x09, 0x5e, 0xae, 0xb8, 0xc0, 0xd7, 0x22, 0x23, 0xb9, 0x56, 0xd4, 0x65, 0x9c, 0x4b,
	0x96, 0x59, 0x5f, 0xd3, 0xf2, 0x92, 0xce, 0x7f, 0x2d, 0xd1, 0x3f, 0x78, 0x81, 0x42, 0xda, 0x6c,
	0xaf, 0x5a, 0xe4, 0xc5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0xea, 0x64, 0xd8, 0x1b, 0x02,
	0x00, 0x00,
}

func (m *SetReactionsParamsJSON) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetReactionsParamsJSON) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetReactionsParamsJSON) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FreeTextParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCli(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.RegisteredReactionParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCli(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintCli(dAtA []byte, offset int, v uint64) int {
	offset -= sovCli(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SetReactionsParamsJSON) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.RegisteredReactionParams.Size()
	n += 1 + l + sovCli(uint64(l))
	l = m.FreeTextParams.Size()
	n += 1 + l + sovCli(uint64(l))
	return n
}

func sovCli(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCli(x uint64) (n int) {
	return sovCli(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SetReactionsParamsJSON) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCli
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
			return fmt.Errorf("proto: SetReactionsParamsJSON: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetReactionsParamsJSON: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredReactionParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RegisteredReactionParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FreeTextParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FreeTextParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCli(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCli
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
func skipCli(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCli
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
					return 0, ErrIntOverflowCli
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
					return 0, ErrIntOverflowCli
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
				return 0, ErrInvalidLengthCli
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCli
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCli
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCli        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCli          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCli = fmt.Errorf("proto: unexpected end of group")
)
