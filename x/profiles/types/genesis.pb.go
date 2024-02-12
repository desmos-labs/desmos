// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// GenesisState defines the profiles module's genesis state.
type GenesisState struct {
	DTagTransferRequests     []DTagTransferRequest         `protobuf:"bytes,1,rep,name=dtag_transfer_requests,json=dtagTransferRequests,proto3" json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
	ChainLinks               []ChainLink                   `protobuf:"bytes,2,rep,name=chain_links,json=chainLinks,proto3" json:"chain_links" yaml:"chain_links"`
	ApplicationLinks         []ApplicationLink             `protobuf:"bytes,3,rep,name=application_links,json=applicationLinks,proto3" json:"application_links" yaml:"application_links"`
	DefaultExternalAddresses []DefaultExternalAddressEntry `protobuf:"bytes,4,rep,name=default_external_addresses,json=defaultExternalAddresses,proto3" json:"default_external_addresses" yaml:"default_external_addresses"`
	IBCPortID                string                        `protobuf:"bytes,5,opt,name=ibc_port_id,json=ibcPortId,proto3" json:"ibc_port_id,omitempty" yaml:"ibc_port_id"`
	Params                   Params                        `protobuf:"bytes,6,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd22d098f73f0a1c, []int{0}
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

// DefaultExternalAddressEntry contains the data of a default extnernal address
type DefaultExternalAddressEntry struct {
	Owner     string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ChainName string `protobuf:"bytes,2,opt,name=chain_name,json=chainName,proto3" json:"chain_name,omitempty"`
	Target    string `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
}

func (m *DefaultExternalAddressEntry) Reset()         { *m = DefaultExternalAddressEntry{} }
func (m *DefaultExternalAddressEntry) String() string { return proto.CompactTextString(m) }
func (*DefaultExternalAddressEntry) ProtoMessage()    {}
func (*DefaultExternalAddressEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd22d098f73f0a1c, []int{1}
}
func (m *DefaultExternalAddressEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DefaultExternalAddressEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DefaultExternalAddressEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DefaultExternalAddressEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefaultExternalAddressEntry.Merge(m, src)
}
func (m *DefaultExternalAddressEntry) XXX_Size() int {
	return m.Size()
}
func (m *DefaultExternalAddressEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DefaultExternalAddressEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DefaultExternalAddressEntry proto.InternalMessageInfo

func (m *DefaultExternalAddressEntry) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DefaultExternalAddressEntry) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *DefaultExternalAddressEntry) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.profiles.v3.GenesisState")
	proto.RegisterType((*DefaultExternalAddressEntry)(nil), "desmos.profiles.v3.DefaultExternalAddressEntry")
}

func init() { proto.RegisterFile("desmos/profiles/v3/genesis.proto", fileDescriptor_bd22d098f73f0a1c) }

var fileDescriptor_bd22d098f73f0a1c = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xe3, 0xfe, 0x89, 0x94, 0x0b, 0x48, 0xf4, 0x14, 0x55, 0x26, 0xa8, 0x76, 0x48, 0x05,
	0x14, 0x50, 0x6d, 0xd1, 0x0e, 0x48, 0xdd, 0xea, 0xb6, 0x42, 0x15, 0xa8, 0xaa, 0xd2, 0x4e, 0x2c,
	0xd6, 0xc5, 0xbe, 0xba, 0xa7, 0xda, 0x3e, 0x73, 0x77, 0x49, 0x9b, 0x9d, 0x81, 0x91, 0x2f, 0x80,
	0x04, 0x1b, 0x23, 0x03, 0xdf, 0x81, 0x8e, 0x15, 0x13, 0x53, 0x84, 0x92, 0x81, 0xbd, 0x9f, 0x00,
	0xf9, 0xee, 0x4c, 0x02, 0x75, 0x58, 0xa2, 0xb3, 0xdf, 0xdf, 0xf3, 0x3e, 0xcf, 0xbd, 0xf1, 0x0b,
	0x5a, 0x21, 0xe6, 0x09, 0xe5, 0x6e, 0xc6, 0xe8, 0x09, 0x89, 0x31, 0x77, 0xfb, 0x9b, 0x6e, 0x84,
	0x53, 0xcc, 0x09, 0x77, 0x32, 0x46, 0x05, 0x85, 0x50, 0x11, 0x4e, 0x41, 0x38, 0xfd, 0xcd, 0xe6,
	0x12, 0x4a, 0x48, 0x4a, 0x5d, 0xf9, 0xab, 0xb0, 0x66, 0x23, 0xa2, 0x11, 0x95, 0x47, 0x37, 0x3f,
	0xe9, 0xb7, 0x77, 0x03, 0x9a, 0x8b, 0x7d, 0x55, 0x50, 0x0f, 0xba, 0xf4, 0xb0, 0xc4, 0x39, 0xa1,
	0x21, 0x8e, 0xb9, 0x9f, 0x21, 0x86, 0x92, 0x82, 0x5b, 0x9f, 0xcd, 0x85, 0x02, 0x45, 0x3e, 0xc3,
	0x6f, 0x7a, 0x98, 0x8b, 0x02, 0x7f, 0x3a, 0x1b, 0x0f, 0x4e, 0x11, 0x49, 0xfd, 0x98, 0xa4, 0x67,
	0x05, 0xfc, 0x78, 0x36, 0x8c, 0xb2, 0x6c, 0x1a, 0x6d, 0x7f, 0x5b, 0x04, 0xb7, 0x5e, 0xa8, 0xc1,
	0x1c, 0x09, 0x24, 0x30, 0xfc, 0x64, 0x80, 0x65, 0x19, 0x40, 0x30, 0x94, 0xf2, 0x13, 0xcc, 0xfe,
	0x24, 0x31, 0x8d, 0xd6, 0xfc, 0x5a, 0x7d, 0xe3, 0x91, 0x73, 0x73, 0x72, 0xce, 0xee, 0x31, 0x8a,
	0x8e, 0xb5, 0xa0, 0xa3, 0x78, 0xcf, 0xbb, 0x1c, 0xda, 0x95, 0xd1, 0xd0, 0x6e, 0x94, 0x14, 0xf9,
	0xf5, 0xd0, 0x5e, 0x19, 0xa0, 0x24, 0xde, 0x6a, 0x97, 0x9b, 0xb5, 0x3f, 0xff, 0xfa, 0xf2, 0xc4,
	0xe8, 0x34, 0xf2, 0xea, 0xbf, 0x5a, 0xe8, 0x83, 0xfa, 0xd4, 0xa5, 0xcd, 0x39, 0x99, 0x6b, 0xa5,
	0x2c, 0xd7, 0x4e, 0x8e, 0xbd, 0x22, 0xe9, 0x99, 0x67, 0xe7, 0x69, 0xae, 0x87, 0x36, 0x54, 0xae,
	0x53, 0x7a, 0x6d, 0x05, 0x82, 0x82, 0xe5, 0xf0, 0x1c, 0x2c, 0xa1, 0x2c, 0x8b, 0x49, 0x80, 0x04,
	0xa1, 0x85, 0xcd, 0xbc, 0xb4, 0x59, 0x2d, 0xb3, 0xd9, 0x9e, 0xc0, 0xd2, 0xec, 0x81, 0x36, 0x33,
	0x95, 0xd9, 0x8d, 0x5e, 0xda, 0xf2, 0x0e, 0xfa, 0x5b, 0xc7, 0xe1, 0x07, 0x03, 0x34, 0x43, 0x7c,
	0x82, 0x7a, 0xb1, 0xf0, 0xf1, 0x85, 0xc0, 0x2c, 0x45, 0xb1, 0x8f, 0xc2, 0x90, 0x61, 0xce, 0x31,
	0x37, 0x17, 0x64, 0x04, 0xb7, 0xf4, 0x1f, 0x50, 0xaa, 0x3d, 0x2d, 0xda, 0x56, 0x9a, 0xbd, 0x54,
	0xb0, 0x81, 0xe7, 0xe8, 0x38, 0xf7, 0xf5, 0xc4, 0x67, 0x1a, 0xe8, 0x5c, 0x66, 0x58, 0xda, 0x0c,
	0x73, 0xb8, 0x03, 0xea, 0xa4, 0x1b, 0xf8, 0x19, 0x65, 0xc2, 0x27, 0xa1, 0xb9, 0xd8, 0x32, 0xd6,
	0x6a, 0xde, 0xea, 0x68, 0x68, 0xd7, 0xf6, 0xbd, 0x9d, 0x43, 0xca, 0xc4, 0xfe, 0xee, 0x64, 0xc6,
	0x53, 0x64, 0xbb, 0x53, 0x23, 0xdd, 0x40, 0x02, 0x21, 0x3c, 0x00, 0x55, 0xb5, 0x0a, 0x66, 0xb5,
	0x65, 0xac, 0xd5, 0x37, 0x9a, 0x65, 0xf7, 0x39, 0x94, 0x84, 0xd7, 0xd4, 0xd1, 0x6f, 0xab, 0x96,
	0x4a, 0xa7, 0x63, 0xea, 0x2e, 0x5b, 0x0b, 0xef, 0x3e, 0xda, 0x95, 0xf6, 0x5b, 0x03, 0xdc, 0xfb,
	0xcf, 0x10, 0xa0, 0x03, 0x16, 0xe9, 0x79, 0x8a, 0x99, 0x69, 0xc8, 0xd0, 0xe6, 0xf7, 0xaf, 0xeb,
	0x0d, 0xbd, 0xb9, 0x9a, 0x3b, 0x12, 0x8c, 0xa4, 0x51, 0x47, 0x61, 0x70, 0x05, 0xa8, 0x2f, 0xc2,
	0x4f, 0x51, 0x82, 0xcd, 0xb9, 0x5c, 0xd4, 0xa9, 0xc9, 0x37, 0x07, 0x28, 0xc1, 0x70, 0x19, 0x54,
	0x05, 0x62, 0x11, 0x16, 0xe6, 0xbc, 0x2c, 0xe9, 0x27, 0xef, 0xe5, 0xe5, 0xc8, 0x32, 0xae, 0x46,
	0x96, 0xf1, 0x73, 0x64, 0x19, 0xef, 0xc7, 0x56, 0xe5, 0x6a, 0x6c, 0x55, 0x7e, 0x8c, 0xad, 0xca,
	0xeb, 0x67, 0x11, 0x11, 0xa7, 0xbd, 0xae, 0x13, 0xd0, 0xc4, 0x55, 0x17, 0x5e, 0x8f, 0x51, 0x97,
	0xeb, 0xb3, 0xdb, 0x7f, 0xee, 0x5e, 0x4c, 0x36, 0x56, 0x0c, 0x32, 0xcc, 0xbb, 0x55, 0xb9, 0xa4,
	0x9b, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x40, 0xdd, 0x90, 0xcf, 0x04, 0x00, 0x00,
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.IBCPortID) > 0 {
		i -= len(m.IBCPortID)
		copy(dAtA[i:], m.IBCPortID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.IBCPortID)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DefaultExternalAddresses) > 0 {
		for iNdEx := len(m.DefaultExternalAddresses) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DefaultExternalAddresses[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
			dAtA[i] = 0x1a
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
			dAtA[i] = 0x12
		}
	}
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

func (m *DefaultExternalAddressEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DefaultExternalAddressEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DefaultExternalAddressEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Target) > 0 {
		i -= len(m.Target)
		copy(dAtA[i:], m.Target)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Target)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
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
	if len(m.DefaultExternalAddresses) > 0 {
		for _, e := range m.DefaultExternalAddresses {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.IBCPortID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *DefaultExternalAddressEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Target)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
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
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DefaultExternalAddresses", wireType)
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
			m.DefaultExternalAddresses = append(m.DefaultExternalAddresses, DefaultExternalAddressEntry{})
			if err := m.DefaultExternalAddresses[len(m.DefaultExternalAddresses)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
		case 6:
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
func (m *DefaultExternalAddressEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DefaultExternalAddressEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DefaultExternalAddressEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
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
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Target", wireType)
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
			m.Target = string(dAtA[iNdEx:postIndex])
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
