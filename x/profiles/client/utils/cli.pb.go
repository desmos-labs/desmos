// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta2/client/cli.proto

package utils

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	types1 "github.com/desmos-labs/desmos/x/profiles/types"
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

// ChainLinkJSON contains the data required to create a ChainLink using the CLI
// command
type ChainLinkJSON struct {
	// Address contains the data of the external chain address to be connected
	// with the Desmos profile
	Address *types.Any `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// Proof contains the ownership proof of the external chain address
	Proof types1.Proof `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof" yaml:"proof"`
	// ChainConfig contains the configuration of the external chain
	ChainConfig types1.ChainConfig `protobuf:"bytes,3,opt,name=chain_config,json=chainConfig,proto3" json:"chain_config" yaml:"chain_config"`
}

func (m *ChainLinkJSON) Reset()         { *m = ChainLinkJSON{} }
func (m *ChainLinkJSON) String() string { return proto.CompactTextString(m) }
func (*ChainLinkJSON) ProtoMessage()    {}
func (*ChainLinkJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_12f99432bbe19bc1, []int{0}
}
func (m *ChainLinkJSON) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainLinkJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainLinkJSON.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainLinkJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainLinkJSON.Merge(m, src)
}
func (m *ChainLinkJSON) XXX_Size() int {
	return m.Size()
}
func (m *ChainLinkJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainLinkJSON.DiscardUnknown(m)
}

var xxx_messageInfo_ChainLinkJSON proto.InternalMessageInfo

func (m *ChainLinkJSON) GetAddress() *types.Any {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ChainLinkJSON) GetProof() types1.Proof {
	if m != nil {
		return m.Proof
	}
	return types1.Proof{}
}

func (m *ChainLinkJSON) GetChainConfig() types1.ChainConfig {
	if m != nil {
		return m.ChainConfig
	}
	return types1.ChainConfig{}
}

func init() {
	proto.RegisterType((*ChainLinkJSON)(nil), "desmos.profiles.v1beta2.client.ChainLinkJSON")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta2/client/cli.proto", fileDescriptor_12f99432bbe19bc1)
}

var fileDescriptor_12f99432bbe19bc1 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x8f, 0xda, 0x30,
	0x18, 0x87, 0x13, 0xaa, 0xfe, 0x51, 0xa0, 0x1d, 0x52, 0x06, 0x4a, 0x25, 0xa7, 0x8d, 0x3a, 0xb0,
	0x60, 0xb7, 0x54, 0x5d, 0xd8, 0x08, 0x9d, 0x50, 0x55, 0x5a, 0xba, 0x75, 0x41, 0x4e, 0xe2, 0x04,
	0x0b, 0x27, 0x8e, 0x62, 0x53, 0x95, 0x6f, 0xd1, 0xb1, 0x53, 0xc5, 0x87, 0xe8, 0x87, 0x40, 0x37,
	0x31, 0xde, 0x84, 0x4e, 0xb0, 0xdc, 0xcc, 0x27, 0x38, 0xc5, 0x76, 0xc4, 0xe9, 0x24, 0xa6, 0xbc,
	0xce, 0xfb, 0xfc, 0x1e, 0xbf, 0x76, 0xe2, 0xf4, 0x62, 0x22, 0x32, 0x2e, 0x50, 0x51, 0xf2, 0x84,
	0x32, 0x22, 0xd0, 0xaf, 0x0f, 0x21, 0x91, 0x78, 0x80, 0x22, 0x46, 0x49, 0x2e, 0xab, 0x07, 0x2c,
	0x4a, 0x2e, 0xb9, 0x0b, 0x34, 0x09, 0x6b, 0x12, 0x1a, 0x12, 0x6a, 0xb2, 0xdb, 0x4e, 0x79, 0xca,
	0x15, 0x8a, 0xaa, 0x4a, 0xa7, 0xba, 0xaf, 0x52, 0xce, 0x53, 0x46, 0x90, 0x5a, 0x85, 0xab, 0x04,
	0xe1, 0x7c, 0x6d, 0x5a, 0xde, 0xc3, 0x96, 0xa4, 0x19, 0x11, 0x12, 0x67, 0x45, 0x9d, 0x8d, 0x78,
	0xb5, 0xe3, 0x5c, 0x4b, 0xf5, 0xc2, 0xb4, 0xde, 0x5f, 0x1a, 0x3b, 0xe3, 0x31, 0x61, 0x62, 0x1e,
	0x2d, 0x30, 0xcd, 0xe7, 0x8c, 0xe6, 0x4b, 0x93, 0xf0, 0xff, 0x35, 0x9c, 0xe7, 0xe3, 0xea, 0xed,
	0x17, 0x9a, 0x2f, 0x27, 0x3f, 0xa6, 0x5f, 0xdd, 0xef, 0xce, 0x53, 0x1c, 0xc7, 0x25, 0x11, 0xa2,
	0x63, 0xbf, 0xb1, 0x7b, 0xcd, 0x41, 0x1b, 0xea, 0x89, 0x60, 0x3d, 0x11, 0x1c, 0xe5, 0xeb, 0xe0,
	0xed, 0x69, 0xef, 0xbd, 0x58, 0xe3, 0x8c, 0x0d, 0x7d, 0x83, 0xfb, 0x57, 0xff, 0xfb, 0xcd, 0x91,
	0xae, 0x3f, 0x63, 0x89, 0x67, 0xb5, 0xc7, 0x9d, 0x38, 0x8f, 0x8b, 0x92, 0xf3, 0xa4, 0xd3, 0x50,
	0x42, 0x00, 0x2f, 0xdd, 0xd9, 0xb7, 0x8a, 0x0a, 0xda, 0xdb, 0xbd, 0x67, 0x9d, 0xf6, 0x5e, 0x4b,
	0xeb, 0x55, 0xd4, 0x9f, 0x69, 0x85, 0x1b, 0x3b, 0x2d, 0x7d, 0x8a, 0x88, 0xe7, 0x09, 0x4d, 0x3b,
	0x8f, 0x94, 0xf2, 0xdd, 0x45, 0xa5, 0x3a, 0xdc, 0x58, 0xb1, 0xc1, 0x6b, 0x23, 0x7e, 0xa9, 0xc5,
	0xf7, 0x3d, 0xfe, 0xac, 0x19, 0x9d, 0xc9, 0xe1, 0xb3, 0xbf, 0x1b, 0xcf, 0xbe, 0xdd, 0x78, 0x76,
	0x30, 0xdd, 0x1e, 0x80, 0xbd, 0x3b, 0x00, 0xfb, 0xe6, 0x00, 0xec, 0x3f, 0x47, 0x60, 0xed, 0x8e,
	0xc0, 0xba, 0x3e, 0x02, 0xeb, 0xe7, 0xa7, 0x94, 0xca, 0xc5, 0x2a, 0x84, 0x11, 0xcf, 0x90, 0xde,
	0xbd, 0xcf, 0x70, 0x28, 0x4c, 0x8d, 0x7e, 0x9f, 0xbf, 0x82, 0xf9, 0x69, 0x56, 0x92, 0x32, 0x11,
	0x3e, 0x51, 0xd7, 0xf8, 0xf1, 0x2e, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xa7, 0x9e, 0x6e, 0x63, 0x02,
	0x00, 0x00,
}

func (this *ChainLinkJSON) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ChainLinkJSON)
	if !ok {
		that2, ok := that.(ChainLinkJSON)
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
	if !this.Address.Equal(that1.Address) {
		return false
	}
	if !this.Proof.Equal(&that1.Proof) {
		return false
	}
	if !this.ChainConfig.Equal(&that1.ChainConfig) {
		return false
	}
	return true
}
func (m *ChainLinkJSON) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainLinkJSON) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainLinkJSON) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCli(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCli(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Address != nil {
		{
			size, err := m.Address.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCli(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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
func (m *ChainLinkJSON) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Address != nil {
		l = m.Address.Size()
		n += 1 + l + sovCli(uint64(l))
	}
	l = m.Proof.Size()
	n += 1 + l + sovCli(uint64(l))
	l = m.ChainConfig.Size()
	n += 1 + l + sovCli(uint64(l))
	return n
}

func sovCli(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCli(x uint64) (n int) {
	return sovCli(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainLinkJSON) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ChainLinkJSON: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainLinkJSON: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			if m.Address == nil {
				m.Address = &types.Any{}
			}
			if err := m.Address.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
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
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainConfig", wireType)
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
			if err := m.ChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
