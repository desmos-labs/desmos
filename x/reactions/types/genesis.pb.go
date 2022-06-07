// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/reactions/v1/genesis.proto

package types

import (
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

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	SubspacesData       []SubspaceDataEntry       `protobuf:"bytes,1,rep,name=subspaces_data,json=subspacesData,proto3" json:"subspaces_data"`
	RegisteredReactions []RegisteredReaction      `protobuf:"bytes,2,rep,name=registered_reactions,json=registeredReactions,proto3" json:"registered_reactions"`
	Reactions           []Reaction                `protobuf:"bytes,3,rep,name=reactions,proto3" json:"reactions"`
	SubspacesParams     []SubspaceReactionsParams `protobuf:"bytes,4,rep,name=subspaces_params,json=subspacesParams,proto3" json:"subspaces_params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b04d309ae6ac6f, []int{0}
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

func (m *GenesisState) GetSubspacesData() []SubspaceDataEntry {
	if m != nil {
		return m.SubspacesData
	}
	return nil
}

func (m *GenesisState) GetRegisteredReactions() []RegisteredReaction {
	if m != nil {
		return m.RegisteredReactions
	}
	return nil
}

func (m *GenesisState) GetReactions() []Reaction {
	if m != nil {
		return m.Reactions
	}
	return nil
}

func (m *GenesisState) GetSubspacesParams() []SubspaceReactionsParams {
	if m != nil {
		return m.SubspacesParams
	}
	return nil
}

// SubspaceDataEntry contains the data related to a single subspace
type SubspaceDataEntry struct {
	// Id of the subspace to which the data relates
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	// Id of the next registered reaction inside the subspace
	RegisteredReactionID uint32 `protobuf:"varint,2,opt,name=registered_reaction_id,json=registeredReactionId,proto3" json:"registered_reaction_id,omitempty"`
	// Id of the next reaction inside the subspace
	ReactionID uint64 `protobuf:"varint,3,opt,name=reaction_id,json=reactionId,proto3" json:"reaction_id,omitempty"`
}

func (m *SubspaceDataEntry) Reset()         { *m = SubspaceDataEntry{} }
func (m *SubspaceDataEntry) String() string { return proto.CompactTextString(m) }
func (*SubspaceDataEntry) ProtoMessage()    {}
func (*SubspaceDataEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_81b04d309ae6ac6f, []int{1}
}
func (m *SubspaceDataEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubspaceDataEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubspaceDataEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubspaceDataEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubspaceDataEntry.Merge(m, src)
}
func (m *SubspaceDataEntry) XXX_Size() int {
	return m.Size()
}
func (m *SubspaceDataEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SubspaceDataEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SubspaceDataEntry proto.InternalMessageInfo

func (m *SubspaceDataEntry) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *SubspaceDataEntry) GetRegisteredReactionID() uint32 {
	if m != nil {
		return m.RegisteredReactionID
	}
	return 0
}

func (m *SubspaceDataEntry) GetReactionID() uint64 {
	if m != nil {
		return m.ReactionID
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.reactions.v1.GenesisState")
	proto.RegisterType((*SubspaceDataEntry)(nil), "desmos.reactions.v1.SubspaceDataEntry")
}

func init() { proto.RegisterFile("desmos/reactions/v1/genesis.proto", fileDescriptor_81b04d309ae6ac6f) }

var fileDescriptor_81b04d309ae6ac6f = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8f, 0x93, 0x40,
	0x1c, 0xc5, 0x99, 0xb6, 0x31, 0x3a, 0xeb, 0xae, 0xca, 0x12, 0x83, 0x9b, 0x08, 0xb8, 0x07, 0xed,
	0x41, 0x99, 0xec, 0xee, 0xcd, 0x9b, 0xa4, 0xc6, 0x90, 0x18, 0x63, 0xe8, 0xcd, 0xc4, 0xe0, 0x00,
	0x23, 0x92, 0x14, 0x86, 0xcc, 0x4c, 0x1b, 0xfb, 0x2d, 0x3c, 0x7a, 0xdc, 0x8f, 0xd3, 0x9b, 0x3d,
	0x7a, 0x6a, 0x0c, 0xbd, 0x78, 0xf5, 0x1b, 0x18, 0x60, 0x80, 0x2a, 0xe8, 0x8d, 0x99, 0xf7, 0xde,
	0x8f, 0xff, 0xff, 0x65, 0xe0, 0xa3, 0x88, 0xf0, 0x94, 0x72, 0xc4, 0x08, 0x0e, 0x45, 0x42, 0x33,
	0x8e, 0x56, 0x17, 0x28, 0x26, 0x19, 0xe1, 0x09, 0xb7, 0x73, 0x46, 0x05, 0x55, 0x4f, 0x6b, 0x8b,
	0xdd, 0x5a, 0xec, 0xd5, 0xc5, 0x99, 0x16, 0xd3, 0x98, 0x56, 0x3a, 0x2a, 0xbf, 0x6a, 0xeb, 0xd9,
	0x83, 0x98, 0xd2, 0x78, 0x41, 0x50, 0x75, 0x0a, 0x96, 0x1f, 0x11, 0xce, 0xd6, 0x52, 0x32, 0xff,
	0x96, 0x44, 0x92, 0x12, 0x2e, 0x70, 0x9a, 0x37, 0xd9, 0x90, 0x96, 0xbf, 0xf1, 0x6b, 0x68, 0x7d,
	0x90, 0x92, 0x35, 0x34, 0x64, 0x4a, 0x23, 0xb2, 0x90, 0x8e, 0xf3, 0x5f, 0x23, 0x78, 0xfb, 0x55,
	0x3d, 0xf5, 0x5c, 0x60, 0x41, 0xd4, 0x39, 0x3c, 0xe1, 0xcb, 0x80, 0xe7, 0x38, 0x24, 0xdc, 0x8f,
	0xb0, 0xc0, 0x3a, 0xb0, 0xc6, 0xd3, 0xa3, 0xcb, 0xc7, 0xf6, 0xc0, 0x36, 0xf6, 0x5c, 0x5a, 0x67,
	0x58, 0xe0, 0x97, 0x99, 0x60, 0x6b, 0x67, 0xb2, 0xd9, 0x99, 0x8a, 0x77, 0xdc, 0x32, 0x4a, 0x45,
	0xfd, 0x00, 0x35, 0x46, 0xe2, 0x84, 0x0b, 0xc2, 0x48, 0xe4, 0xb7, 0x04, 0x7d, 0x54, 0xa1, 0x9f,
	0x0c, 0xa2, 0xbd, 0x36, 0xe0, 0xc9, 0x6b, 0xc9, 0x3e, 0x65, 0x3d, 0x85, 0xab, 0x2f, 0xe0, 0xad,
	0x0e, 0x3b, 0xae, 0xb0, 0x0f, 0xff, 0x81, 0xfd, 0x03, 0xd6, 0xa5, 0xd4, 0xf7, 0xf0, 0x6e, 0xb7,
	0x79, 0x8e, 0x19, 0x4e, 0xb9, 0x3e, 0xa9, 0x48, 0x4f, 0xff, 0xbb, 0x7b, 0x3b, 0xc4, 0xdb, 0x2a,
	0x23, 0xc1, 0x77, 0x5a, 0x56, 0x7d, 0xfd, 0xfc, 0xe6, 0xd7, 0x6b, 0x13, 0xfc, 0xbc, 0x36, 0xc1,
	0xf9, 0x37, 0x00, 0xef, 0xf5, 0x8a, 0x53, 0x11, 0x3c, 0x6a, 0x22, 0x7e, 0x12, 0xe9, 0xc0, 0x02,
	0xd3, 0x89, 0x73, 0x52, 0xec, 0x4c, 0xd8, 0x78, 0xdd, 0x99, 0x07, 0x1b, 0x8b, 0x1b, 0xa9, 0x6f,
	0xe0, 0xfd, 0x81, 0x52, 0xcb, 0xec, 0xc8, 0x02, 0xd3, 0x63, 0x47, 0x2f, 0x76, 0xa6, 0xd6, 0x6f,
	0xd1, 0x9d, 0x79, 0x5a, 0xbf, 0x41, 0x37, 0x2a, 0x07, 0x38, 0x84, 0x8c, 0xbb, 0x01, 0x0e, 0xa2,
	0x90, 0xb5, 0x81, 0x6e, 0x23, 0xe7, 0xf5, 0xa6, 0x30, 0xc0, 0xb6, 0x30, 0xc0, 0x8f, 0xc2, 0x00,
	0x5f, 0xf6, 0x86, 0xb2, 0xdd, 0x1b, 0xca, 0xf7, 0xbd, 0xa1, 0xbc, 0xbb, 0x8c, 0x13, 0xf1, 0x69,
	0x19, 0xd8, 0x21, 0x4d, 0x51, 0x5d, 0xe2, 0xb3, 0x05, 0x0e, 0xb8, 0xfc, 0x46, 0xab, 0x2b, 0xf4,
	0xf9, 0xe0, 0x75, 0x8a, 0x75, 0x4e, 0x78, 0x70, 0xa3, 0x7a, 0x9a, 0x57, 0xbf, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xad, 0x58, 0xed, 0x82, 0x63, 0x03, 0x00, 0x00,
}

func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
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
	if len(this.SubspacesData) != len(that1.SubspacesData) {
		return false
	}
	for i := range this.SubspacesData {
		if !this.SubspacesData[i].Equal(&that1.SubspacesData[i]) {
			return false
		}
	}
	if len(this.RegisteredReactions) != len(that1.RegisteredReactions) {
		return false
	}
	for i := range this.RegisteredReactions {
		if !this.RegisteredReactions[i].Equal(&that1.RegisteredReactions[i]) {
			return false
		}
	}
	if len(this.Reactions) != len(that1.Reactions) {
		return false
	}
	for i := range this.Reactions {
		if !this.Reactions[i].Equal(&that1.Reactions[i]) {
			return false
		}
	}
	if len(this.SubspacesParams) != len(that1.SubspacesParams) {
		return false
	}
	for i := range this.SubspacesParams {
		if !this.SubspacesParams[i].Equal(&that1.SubspacesParams[i]) {
			return false
		}
	}
	return true
}
func (this *SubspaceDataEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubspaceDataEntry)
	if !ok {
		that2, ok := that.(SubspaceDataEntry)
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
	if this.SubspaceID != that1.SubspaceID {
		return false
	}
	if this.RegisteredReactionID != that1.RegisteredReactionID {
		return false
	}
	if this.ReactionID != that1.ReactionID {
		return false
	}
	return true
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
	if len(m.SubspacesParams) > 0 {
		for iNdEx := len(m.SubspacesParams) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubspacesParams[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Reactions) > 0 {
		for iNdEx := len(m.Reactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Reactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.RegisteredReactions) > 0 {
		for iNdEx := len(m.RegisteredReactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RegisteredReactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.SubspacesData) > 0 {
		for iNdEx := len(m.SubspacesData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SubspacesData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *SubspaceDataEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubspaceDataEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubspaceDataEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReactionID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ReactionID))
		i--
		dAtA[i] = 0x18
	}
	if m.RegisteredReactionID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.RegisteredReactionID))
		i--
		dAtA[i] = 0x10
	}
	if m.SubspaceID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SubspaceID))
		i--
		dAtA[i] = 0x8
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
	if len(m.SubspacesData) > 0 {
		for _, e := range m.SubspacesData {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RegisteredReactions) > 0 {
		for _, e := range m.RegisteredReactions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Reactions) > 0 {
		for _, e := range m.Reactions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SubspacesParams) > 0 {
		for _, e := range m.SubspacesParams {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *SubspaceDataEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.RegisteredReactionID != 0 {
		n += 1 + sovGenesis(uint64(m.RegisteredReactionID))
	}
	if m.ReactionID != 0 {
		n += 1 + sovGenesis(uint64(m.ReactionID))
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
				return fmt.Errorf("proto: wrong wireType = %d for field SubspacesData", wireType)
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
			m.SubspacesData = append(m.SubspacesData, SubspaceDataEntry{})
			if err := m.SubspacesData[len(m.SubspacesData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredReactions", wireType)
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
			m.RegisteredReactions = append(m.RegisteredReactions, RegisteredReaction{})
			if err := m.RegisteredReactions[len(m.RegisteredReactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reactions", wireType)
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
			m.Reactions = append(m.Reactions, Reaction{})
			if err := m.Reactions[len(m.Reactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspacesParams", wireType)
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
			m.SubspacesParams = append(m.SubspacesParams, SubspaceReactionsParams{})
			if err := m.SubspacesParams[len(m.SubspacesParams)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *SubspaceDataEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SubspaceDataEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubspaceDataEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			m.SubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegisteredReactionID", wireType)
			}
			m.RegisteredReactionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegisteredReactionID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReactionID", wireType)
			}
			m.ReactionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReactionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
