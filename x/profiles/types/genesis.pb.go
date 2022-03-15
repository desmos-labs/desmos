// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v2/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the profiles module's genesis state.
type GenesisState struct {
	DTagTransferRequests []DTagTransferRequest `protobuf:"bytes,1,rep,name=dtag_transfer_requests,json=dtagTransferRequests,proto3" json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
	Params               Params                `protobuf:"bytes,2,opt,name=params,proto3" json:"params" yaml:"params"`
	IBCPortID            string                `protobuf:"bytes,3,opt,name=ibc_port_id,json=ibcPortId,proto3" json:"ibc_port_id,omitempty" yaml:"ibc_port_id"`
	ChainLinks           []ChainLink           `protobuf:"bytes,4,rep,name=chain_links,json=chainLinks,proto3" json:"chain_links" yaml:"chain_links"`
	ApplicationLinks     []ApplicationLink     `protobuf:"bytes,5,rep,name=application_links,json=applicationLinks,proto3" json:"application_links" yaml:"application_links"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_be71125223ae0fd1, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.profiles.v2.GenesisState")
}

func init() { proto.RegisterFile("desmos/profiles/v2/genesis.proto", fileDescriptor_be71125223ae0fd1) }

var fileDescriptor_be71125223ae0fd1 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x6a, 0xd4, 0x40,
	0x1c, 0xc7, 0x13, 0xbb, 0x16, 0x36, 0xab, 0xa0, 0x61, 0x95, 0x10, 0x68, 0x12, 0x52, 0xb0, 0x2b,
	0xd2, 0x04, 0xd7, 0x5b, 0xc1, 0x83, 0xd9, 0x82, 0x2c, 0x7a, 0x28, 0xb1, 0xa7, 0x5e, 0xc2, 0x24,
	0x99, 0xa6, 0x43, 0x93, 0xcc, 0x38, 0x33, 0x2d, 0xf6, 0x0d, 0x3c, 0x0a, 0x5e, 0x3d, 0xf8, 0x38,
	0x3d, 0xf6, 0xe8, 0x29, 0x48, 0xf6, 0x0d, 0xf6, 0x09, 0x24, 0x99, 0x89, 0x8d, 0x36, 0xdb, 0xdb,
	0x30, 0xbf, 0xcf, 0xf7, 0x0f, 0x3f, 0x7e, 0x9a, 0x93, 0x42, 0x56, 0x60, 0xe6, 0x13, 0x8a, 0x4f,
	0x51, 0x0e, 0x99, 0x7f, 0x39, 0xf7, 0x33, 0x58, 0x42, 0x86, 0x98, 0x47, 0x28, 0xe6, 0x58, 0xd7,
	0x05, 0xe1, 0x75, 0x84, 0x77, 0x39, 0x37, 0xa7, 0x19, 0xce, 0x70, 0x3b, 0xf6, 0x9b, 0x97, 0x20,
	0xcd, 0x17, 0x03, 0x5e, 0x05, 0x4e, 0x61, 0xce, 0x22, 0x02, 0x28, 0x28, 0xa4, 0xa3, 0xb9, 0x77,
	0x0f, 0x27, 0xbe, 0x24, 0xb8, 0xbf, 0x19, 0x4c, 0x39, 0xc8, 0x22, 0x0a, 0x3f, 0x5f, 0x40, 0xc6,
	0x3b, 0xdf, 0x57, 0x9b, 0xf1, 0xe4, 0x0c, 0xa0, 0x32, 0xca, 0x51, 0x79, 0xde, 0xc1, 0x2f, 0x37,
	0xc3, 0x80, 0x90, 0x3e, 0xea, 0x7e, 0x1f, 0x69, 0x8f, 0xde, 0x8b, 0x9d, 0x7c, 0xe2, 0x80, 0x43,
	0xfd, 0x87, 0xaa, 0x3d, 0x6f, 0x0b, 0x70, 0x0a, 0x4a, 0x76, 0x0a, 0xe9, 0xdf, 0x26, 0x86, 0xea,
	0x6c, 0xcd, 0x26, 0xf3, 0x3d, 0xef, 0xee, 0xd2, 0xbc, 0xc3, 0x63, 0x90, 0x1d, 0x4b, 0x41, 0x28,
	0xf8, 0xe0, 0xed, 0x75, 0x65, 0x2b, 0x75, 0x65, 0x4f, 0x07, 0x86, 0x6c, 0x5d, 0xd9, 0x3b, 0x57,
	0xa0, 0xc8, 0x0f, 0xdc, 0xe1, 0x30, 0x37, 0x9c, 0x36, 0x83, 0xff, 0x65, 0xfa, 0x52, 0xdb, 0x16,
	0xfb, 0x36, 0x1e, 0x38, 0xea, 0x6c, 0x32, 0x37, 0x87, 0xda, 0x1c, 0xb5, 0x44, 0xf0, 0xac, 0x29,
	0xb0, 0xae, 0xec, 0xc7, 0x22, 0x48, 0xe8, 0xdc, 0x50, 0x1a, 0xe8, 0x0b, 0x6d, 0x82, 0xe2, 0x24,
	0x22, 0x98, 0xf2, 0x08, 0xa5, 0xc6, 0x96, 0xa3, 0xce, 0xc6, 0xc1, 0x6e, 0x5d, 0xd9, 0xe3, 0x65,
	0xb0, 0x38, 0xc2, 0x94, 0x2f, 0x0f, 0xd7, 0x95, 0xad, 0x0b, 0x71, 0x8f, 0x74, 0xc3, 0x31, 0x8a,
	0x93, 0x16, 0x48, 0xf5, 0x13, 0x6d, 0xd2, 0xdb, 0xbf, 0x31, 0x6a, 0x57, 0xb4, 0x33, 0x54, 0x6a,
	0xd1, 0x60, 0x1f, 0x51, 0x79, 0x1e, 0x98, 0xb2, 0x97, 0xb4, 0xee, 0xe9, 0xdd, 0x50, 0x4b, 0x3a,
	0x8c, 0xe9, 0x54, 0x7b, 0x0a, 0x08, 0xc9, 0x51, 0x02, 0x38, 0xc2, 0x5d, 0xc2, 0xc3, 0x36, 0x61,
	0x77, 0x28, 0xe1, 0xdd, 0x2d, 0xdc, 0xe6, 0x38, 0x32, 0xc7, 0x10, 0x39, 0x77, 0xbc, 0xdc, 0xf0,
	0x09, 0xf8, 0x57, 0xc2, 0x0e, 0x46, 0x5f, 0x7f, 0xda, 0x4a, 0xf0, 0xe1, 0xba, 0xb6, 0xd4, 0x9b,
	0xda, 0x52, 0x7f, 0xd7, 0x96, 0xfa, 0x6d, 0x65, 0x29, 0x37, 0x2b, 0x4b, 0xf9, 0xb5, 0xb2, 0x94,
	0x93, 0xd7, 0x19, 0xe2, 0x67, 0x17, 0xb1, 0x97, 0xe0, 0xc2, 0x17, 0x15, 0xf6, 0x73, 0x10, 0x33,
	0xf9, 0x6e, 0x0e, 0xed, 0xcb, 0xed, 0xd9, 0xf1, 0x2b, 0x02, 0x59, 0xbc, 0xdd, 0x5e, 0xda, 0x9b,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb1, 0xcb, 0x82, 0x94, 0x8f, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ApplicationLinks) > 0 {
		for iNdEx := len(m.ApplicationLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ApplicationLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ChainLinks) > 0 {
		for iNdEx := len(m.ChainLinks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainLinks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.IBCPortID) > 0 {
		i -= len(m.IBCPortID)
		copy(dAtA[i:], m.IBCPortID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IBCPortID)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.DTagTransferRequests) > 0 {
		for iNdEx := len(m.DTagTransferRequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DTagTransferRequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DTagTransferRequests) > 0 {
		for _, e := range m.DTagTransferRequests {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.IBCPortID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.ChainLinks) > 0 {
		for _, e := range m.ChainLinks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ApplicationLinks) > 0 {
		for _, e := range m.ApplicationLinks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTagTransferRequests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DTagTransferRequests = append(m.DTagTransferRequests, DTagTransferRequest{})
			if err := m.DTagTransferRequests[len(m.DTagTransferRequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IBCPortID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IBCPortID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainLinks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainLinks = append(m.ChainLinks, ChainLink{})
			if err := m.ChainLinks[len(m.ChainLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationLinks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ApplicationLinks = append(m.ApplicationLinks, ApplicationLink{})
			if err := m.ApplicationLinks[len(m.ApplicationLinks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
