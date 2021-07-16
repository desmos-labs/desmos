// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/models_packets.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// LinkChainAccountPacketData defines the object that should be sent inside a
// MsgSendPacket when wanting to link an external chain to a Desmos profile
// using IBC
type LinkChainAccountPacketData struct {
	// SourceAddress contains the details of the external chain address
	SourceAddress *types.Any `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty" yaml:"source_address"`
	// SourceProof represents the proof of ownership of the source address
	SourceProof Proof `protobuf:"bytes,2,opt,name=source_proof,json=sourceProof,proto3" json:"source_proof" yaml:"source_proof"`
	// SourceChainConfig contains the details of the source chain
	SourceChainConfig ChainConfig `protobuf:"bytes,3,opt,name=source_chain_config,json=sourceChainConfig,proto3" json:"source_chain_config" yaml:"source_chain_config"`
	// DestinationAddress represents the Desmos address of the profile that should
	// be linked with the external account
	DestinationAddress string `protobuf:"bytes,4,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty" yaml:"destination_address"`
	// DestinationProof contains the proof of ownership of the DestinationAddress
	DestinationProof Proof `protobuf:"bytes,5,opt,name=destination_proof,json=destinationProof,proto3" json:"destination_proof" yaml:"destination_proof"`
}

func (m *LinkChainAccountPacketData) Reset()         { *m = LinkChainAccountPacketData{} }
func (m *LinkChainAccountPacketData) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountPacketData) ProtoMessage()    {}
func (*LinkChainAccountPacketData) Descriptor() ([]byte, []int) {
	return fileDescriptor_923faf54c46abe52, []int{0}
}
func (m *LinkChainAccountPacketData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountPacketData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountPacketData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountPacketData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountPacketData.Merge(m, src)
}
func (m *LinkChainAccountPacketData) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountPacketData) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountPacketData.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountPacketData proto.InternalMessageInfo

// LinkChainAccountPacketAck defines a struct for the packet acknowledgment
type LinkChainAccountPacketAck struct {
	// SourceAddress contains the external address that has been linked properly
	// with the profile
	SourceAddress string `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
}

func (m *LinkChainAccountPacketAck) Reset()         { *m = LinkChainAccountPacketAck{} }
func (m *LinkChainAccountPacketAck) String() string { return proto.CompactTextString(m) }
func (*LinkChainAccountPacketAck) ProtoMessage()    {}
func (*LinkChainAccountPacketAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_923faf54c46abe52, []int{1}
}
func (m *LinkChainAccountPacketAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LinkChainAccountPacketAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LinkChainAccountPacketAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LinkChainAccountPacketAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinkChainAccountPacketAck.Merge(m, src)
}
func (m *LinkChainAccountPacketAck) XXX_Size() int {
	return m.Size()
}
func (m *LinkChainAccountPacketAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LinkChainAccountPacketAck.DiscardUnknown(m)
}

var xxx_messageInfo_LinkChainAccountPacketAck proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LinkChainAccountPacketData)(nil), "desmos.profiles.v1beta1.LinkChainAccountPacketData")
	proto.RegisterType((*LinkChainAccountPacketAck)(nil), "desmos.profiles.v1beta1.LinkChainAccountPacketAck")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/models_packets.proto", fileDescriptor_923faf54c46abe52)
}

var fileDescriptor_923faf54c46abe52 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0xb6, 0x21, 0x20, 0xd5, 0x01, 0x44, 0x9d, 0x22, 0xd2, 0x20, 0x39, 0x91, 0x05, 0x52, 0x16,
	0x74, 0x86, 0xc2, 0xae, 0xbb, 0xb8, 0x2c, 0xba, 0x40, 0xa2, 0xca, 0x92, 0x05, 0xd1, 0x78, 0x3c,
	0x71, 0x47, 0xb1, 0xe7, 0x59, 0x9e, 0x09, 0x6a, 0x6e, 0xc0, 0x92, 0x23, 0x20, 0x71, 0x05, 0x0e,
	0x51, 0xb1, 0xea, 0x92, 0x55, 0x84, 0x92, 0x1b, 0xf4, 0x04, 0xc8, 0x33, 0x63, 0xea, 0x40, 0x2a,
	0xba, 0x9b, 0xf7, 0xde, 0xf7, 0xf3, 0xfc, 0xf9, 0x79, 0x2f, 0x13, 0x26, 0x73, 0x90, 0xb8, 0x28,
	0x61, 0xca, 0x33, 0x26, 0xf1, 0xa7, 0xc3, 0x98, 0x29, 0x72, 0x88, 0x73, 0x48, 0x58, 0x26, 0x27,
	0x05, 0xa1, 0x33, 0xa6, 0x24, 0x2a, 0x4a, 0x50, 0xe0, 0x3f, 0x35, 0x68, 0x54, 0xa3, 0x91, 0x45,
	0xf7, 0xf6, 0x52, 0x48, 0x41, 0x63, 0x70, 0xf5, 0x32, 0xf0, 0xde, 0x7e, 0x0a, 0x90, 0x66, 0x0c,
	0xeb, 0x2a, 0x9e, 0x4f, 0x31, 0x11, 0x8b, 0x7a, 0x44, 0xa1, 0x52, 0x9a, 0x18, 0x8e, 0x29, 0xec,
	0xe8, 0xd5, 0x7f, 0x56, 0xa2, 0x67, 0x84, 0x8b, 0x49, 0xc6, 0xc5, 0xcc, 0x32, 0xc2, 0x6f, 0x2d,
	0xaf, 0xf7, 0x8e, 0x8b, 0xd9, 0x71, 0x35, 0x19, 0x51, 0x0a, 0x73, 0xa1, 0x4e, 0xf5, 0xe2, 0x6f,
	0x89, 0x22, 0x3e, 0xf3, 0x1e, 0x49, 0x98, 0x97, 0x94, 0x4d, 0x48, 0x92, 0x94, 0x4c, 0xca, 0xae,
	0x3b, 0x70, 0x87, 0xed, 0xd7, 0x7b, 0xc8, 0xec, 0x87, 0xea, 0xfd, 0xd0, 0x48, 0x2c, 0xa2, 0xe1,
	0xd5, 0xb2, 0xff, 0x64, 0x41, 0xf2, 0xec, 0x28, 0xdc, 0x64, 0x85, 0x3f, 0xbe, 0x1f, 0xb4, 0x47,
	0xe6, 0x5d, 0xe9, 0x8e, 0x1f, 0x9a, 0xb9, 0x6d, 0xf9, 0x1f, 0xbd, 0x07, 0x96, 0x50, 0x94, 0x00,
	0xd3, 0xee, 0x1d, 0x6d, 0x12, 0xa0, 0x1b, 0x32, 0x43, 0xa7, 0x15, 0x2a, 0x7a, 0x76, 0xb1, 0xec,
	0x3b, 0x57, 0xcb, 0x7e, 0x67, 0xc3, 0x52, 0x2b, 0x84, 0xe3, 0xb6, 0x29, 0x35, 0xd2, 0x3f, 0xf7,
	0x3a, 0x76, 0x6a, 0x12, 0xa0, 0x20, 0xa6, 0x3c, 0xed, 0xde, 0xd5, 0x36, 0xcf, 0x6f, 0xb4, 0xd1,
	0xa1, 0x1c, 0x6b, 0x6c, 0x14, 0x5a, 0xb3, 0xde, 0x86, 0x59, 0x53, 0x2e, 0x1c, 0xef, 0x9a, 0x6e,
	0x83, 0xe6, 0xbf, 0xf7, 0x3a, 0x09, 0x93, 0x8a, 0x0b, 0xa2, 0x38, 0x88, 0x3f, 0x29, 0xb6, 0x06,
	0xee, 0x70, 0x27, 0x0a, 0xae, 0xf5, 0xb6, 0x80, 0xc2, 0xb1, 0xdf, 0xe8, 0xd6, 0x51, 0xe5, 0xde,
	0x6e, 0x13, 0x6b, 0xf2, 0xba, 0x77, 0xab, 0xbc, 0x06, 0xf6, 0x13, 0xba, 0xff, 0x5a, 0xda, 0xd0,
	0x1e, 0x37, 0x7a, 0x9a, 0x73, 0xd4, 0xfa, 0xfc, 0xb5, 0xef, 0x84, 0x27, 0xde, 0xfe, 0xf6, 0x23,
	0x19, 0xd1, 0x99, 0xff, 0x62, 0xeb, 0x8d, 0xec, 0xfc, 0xf5, 0x8f, 0x8d, 0x52, 0x74, 0x72, 0xb1,
	0x0a, 0xdc, 0xcb, 0x55, 0xe0, 0xfe, 0x5a, 0x05, 0xee, 0x97, 0x75, 0xe0, 0x5c, 0xae, 0x03, 0xe7,
	0xe7, 0x3a, 0x70, 0x3e, 0xa0, 0x94, 0xab, 0xb3, 0x79, 0x8c, 0x28, 0xe4, 0xd8, 0x7c, 0xc7, 0x41,
	0x46, 0x62, 0x69, 0xdf, 0xf8, 0xfc, 0xfa, 0xa8, 0xd5, 0xa2, 0x60, 0x32, 0xbe, 0xaf, 0x4f, 0xef,
	0xcd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x7e, 0x0e, 0x6b, 0x87, 0x03, 0x00, 0x00,
}

func (m *LinkChainAccountPacketData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountPacketData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountPacketData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DestinationProof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.DestinationAddress) > 0 {
		i -= len(m.DestinationAddress)
		copy(dAtA[i:], m.DestinationAddress)
		i = encodeVarintModelsPackets(dAtA, i, uint64(len(m.DestinationAddress)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.SourceChainConfig.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.SourceProof.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsPackets(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.SourceAddress != nil {
		{
			size, err := m.SourceAddress.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsPackets(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LinkChainAccountPacketAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LinkChainAccountPacketAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LinkChainAccountPacketAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SourceAddress) > 0 {
		i -= len(m.SourceAddress)
		copy(dAtA[i:], m.SourceAddress)
		i = encodeVarintModelsPackets(dAtA, i, uint64(len(m.SourceAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsPackets(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsPackets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LinkChainAccountPacketData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceAddress != nil {
		l = m.SourceAddress.Size()
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	l = m.SourceProof.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	l = m.SourceChainConfig.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	l = len(m.DestinationAddress)
	if l > 0 {
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	l = m.DestinationProof.Size()
	n += 1 + l + sovModelsPackets(uint64(l))
	return n
}

func (m *LinkChainAccountPacketAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SourceAddress)
	if l > 0 {
		n += 1 + l + sovModelsPackets(uint64(l))
	}
	return n
}

func sovModelsPackets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsPackets(x uint64) (n int) {
	return sovModelsPackets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LinkChainAccountPacketData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsPackets
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
			return fmt.Errorf("proto: LinkChainAccountPacketData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountPacketData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SourceAddress == nil {
				m.SourceAddress = &types.Any{}
			}
			if err := m.SourceAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SourceProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SourceChainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationProof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DestinationProof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsPackets
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
func (m *LinkChainAccountPacketAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsPackets
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
			return fmt.Errorf("proto: LinkChainAccountPacketAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LinkChainAccountPacketAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsPackets
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
				return ErrInvalidLengthModelsPackets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsPackets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsPackets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsPackets
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
func skipModelsPackets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsPackets
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
					return 0, ErrIntOverflowModelsPackets
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
					return 0, ErrIntOverflowModelsPackets
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
				return 0, ErrInvalidLengthModelsPackets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsPackets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsPackets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsPackets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsPackets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsPackets = fmt.Errorf("proto: unexpected end of group")
)
