// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/client/cli.proto

package utils

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/desmos-labs/desmos/v4/x/profiles/types"
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
	return fileDescriptor_8dd0252be1e0ade7, []int{0}
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
	proto.RegisterType((*ChainLinkJSON)(nil), "desmos.profiles.v3.client.ChainLinkJSON")
}

func init() {
	proto.RegisterFile("desmos/profiles/v3/client/cli.proto", fileDescriptor_8dd0252be1e0ade7)
}

var fileDescriptor_8dd0252be1e0ade7 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xbf, 0x4e, 0xfb, 0x30,
	0x10, 0xc7, 0x93, 0xfe, 0xf4, 0x03, 0x94, 0x16, 0x86, 0xd0, 0xa1, 0x2d, 0x52, 0x82, 0xc2, 0x82,
	0x84, 0x6a, 0x4b, 0xb4, 0x03, 0xea, 0xd6, 0x14, 0x16, 0x84, 0x00, 0x85, 0x8d, 0x25, 0xca, 0xff,
	0x5a, 0x75, 0xec, 0x2a, 0x4e, 0x2b, 0xfa, 0x16, 0x8c, 0x8c, 0xdd, 0x78, 0x01, 0x1e, 0xa2, 0x62,
	0xea, 0xc8, 0x54, 0xa1, 0x76, 0x61, 0xee, 0x13, 0xa0, 0xd8, 0x89, 0xca, 0x90, 0x29, 0x77, 0xb9,
	0xcf, 0xf7, 0x7b, 0x77, 0xb6, 0x95, 0x33, 0x3f, 0x60, 0x31, 0x65, 0x70, 0x9c, 0xd0, 0x10, 0xe1,
	0x80, 0xc1, 0x69, 0x07, 0x7a, 0x18, 0x05, 0x24, 0xcd, 0x3e, 0x60, 0x9c, 0xd0, 0x94, 0xaa, 0x4d,
	0x01, 0x81, 0x02, 0x02, 0xd3, 0x0e, 0x10, 0x50, 0xab, 0x1e, 0xd1, 0x88, 0x72, 0x0a, 0x66, 0x91,
	0x10, 0xb4, 0x9a, 0x11, 0xa5, 0x11, 0x0e, 0x20, 0xcf, 0xdc, 0x49, 0x08, 0x1d, 0x32, 0x2b, 0x4a,
	0x1e, 0xcd, 0xbc, 0x6c, 0xa1, 0x11, 0x49, 0x5e, 0xba, 0x28, 0x99, 0x25, 0xa6, 0x7e, 0x80, 0x99,
	0xed, 0x0d, 0x1d, 0x44, 0x6c, 0x8c, 0xc8, 0x28, 0x87, 0x8d, 0xf7, 0x8a, 0x72, 0x38, 0xc8, 0xfe,
	0xde, 0x21, 0x32, 0xba, 0x7d, 0x7a, 0xb8, 0x57, 0x7d, 0x65, 0xdf, 0xf1, 0xfd, 0x24, 0x60, 0xac,
	0x21, 0x9f, 0xca, 0xe7, 0xd5, 0xcb, 0x3a, 0x10, 0x63, 0x80, 0x62, 0x0c, 0xd0, 0x27, 0x33, 0xb3,
	0xbb, 0x5d, 0xe9, 0x47, 0x33, 0x27, 0xc6, 0x3d, 0x23, 0xc7, 0x8d, 0xcf, 0x8f, 0xb6, 0x56, 0xb2,
	0x62, 0x5f, 0x94, 0xaf, 0x9d, 0xd4, 0xb1, 0x0a, 0x6b, 0xf5, 0x46, 0xf9, 0x3f, 0x4e, 0x28, 0x0d,
	0x1b, 0x15, 0xde, 0xa3, 0x09, 0x4a, 0x84, 0x8f, 0x19, 0x60, 0xd6, 0x17, 0x2b, 0x5d, 0xda, 0xae,
	0xf4, 0x9a, 0x68, 0xc6, 0x55, 0x86, 0x25, 0xd4, 0xaa, 0xad, 0xd4, 0xc4, 0x4e, 0x1e, 0x25, 0x21,
	0x8a, 0x1a, 0xff, 0xb8, 0x9b, 0x5e, 0xe6, 0xc6, 0xb7, 0x1c, 0x70, 0xcc, 0x3c, 0xc9, 0x3d, 0x8f,
	0x85, 0xe7, 0x5f, 0x0b, 0xc3, 0xaa, 0x7a, 0x3b, 0xb2, 0x77, 0xf0, 0x36, 0xd7, 0xe5, 0x9f, 0xb9,
	0x2e, 0x9b, 0xd6, 0x62, 0xad, 0xc9, 0xcb, 0xb5, 0x26, 0x7f, 0xaf, 0x35, 0xf9, 0x75, 0xa3, 0x49,
	0xcb, 0x8d, 0x26, 0x7d, 0x6d, 0x34, 0xe9, 0xf9, 0x2a, 0x42, 0xe9, 0x70, 0xe2, 0x02, 0x8f, 0xc6,
	0x50, 0x34, 0x6e, 0x63, 0xc7, 0x65, 0x79, 0x0c, 0xa7, 0x5d, 0xf8, 0xb2, 0xbb, 0x8c, 0xfc, 0x55,
	0x4c, 0x52, 0x84, 0x99, 0xbb, 0xc7, 0x8f, 0xb4, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x83, 0x63,
	0x21, 0x18, 0x3f, 0x02, 0x00, 0x00,
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
