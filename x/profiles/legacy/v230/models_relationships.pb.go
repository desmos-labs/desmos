// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/legacy/v230/models_relationships.proto

package v230

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

// Relationship is the struct of a relationship.
// It represent the concept of "follow" of traditional social networks.
type Relationship struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty" yaml:"recipient"`
	Subspace  string `protobuf:"bytes,3,opt,name=subspace,proto3" json:"subspace,omitempty" yaml:"subspace"`
}

func (m *Relationship) Reset()         { *m = Relationship{} }
func (m *Relationship) String() string { return proto.CompactTextString(m) }
func (*Relationship) ProtoMessage()    {}
func (*Relationship) Descriptor() ([]byte, []int) {
	return fileDescriptor_4929e4628d8fdc04, []int{0}
}
func (m *Relationship) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Relationship) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Relationship.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Relationship) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Relationship.Merge(m, src)
}
func (m *Relationship) XXX_Size() int {
	return m.Size()
}
func (m *Relationship) XXX_DiscardUnknown() {
	xxx_messageInfo_Relationship.DiscardUnknown(m)
}

var xxx_messageInfo_Relationship proto.InternalMessageInfo

func (m *Relationship) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Relationship) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Relationship) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

// UserBlock represents the fact that the Blocker has blocked the given Blocked
// user.
type UserBlock struct {
	// Blocker represents the address of the user blocking another one
	Blocker string `protobuf:"bytes,1,opt,name=blocker,proto3" json:"blocker,omitempty" yaml:"blocker"`
	// Blocked represents the address of the blocked user
	Blocked string `protobuf:"bytes,2,opt,name=blocked,proto3" json:"blocked,omitempty" yaml:"blocked"`
	// Reason represents the optional reason the user has been blocked for.
	Reason string `protobuf:"bytes,3,opt,name=reason,proto3" json:"reason,omitempty" yaml:"reason"`
	// Subspace contains the ID of the subspace inside which the user should be
	// blocked
	Subspace string `protobuf:"bytes,4,opt,name=subspace,proto3" json:"subspace,omitempty" yaml:"subspace"`
}

func (m *UserBlock) Reset()         { *m = UserBlock{} }
func (m *UserBlock) String() string { return proto.CompactTextString(m) }
func (*UserBlock) ProtoMessage()    {}
func (*UserBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_4929e4628d8fdc04, []int{1}
}
func (m *UserBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserBlock.Merge(m, src)
}
func (m *UserBlock) XXX_Size() int {
	return m.Size()
}
func (m *UserBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_UserBlock.DiscardUnknown(m)
}

var xxx_messageInfo_UserBlock proto.InternalMessageInfo

func (m *UserBlock) GetBlocker() string {
	if m != nil {
		return m.Blocker
	}
	return ""
}

func (m *UserBlock) GetBlocked() string {
	if m != nil {
		return m.Blocked
	}
	return ""
}

func (m *UserBlock) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *UserBlock) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

func init() {
	proto.RegisterType((*Relationship)(nil), "desmos.profiles.legacy.v230.Relationship")
	proto.RegisterType((*UserBlock)(nil), "desmos.profiles.legacy.v230.UserBlock")
}

func init() {
	proto.RegisterFile("desmos/profiles/legacy/v230/models_relationships.proto", fileDescriptor_4929e4628d8fdc04)
}

var fileDescriptor_4929e4628d8fdc04 = []byte{
	// 387 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbd, 0x4e, 0xeb, 0x30,
	0x18, 0x86, 0xeb, 0x73, 0x8e, 0x7a, 0xda, 0xe8, 0xfc, 0x40, 0xe8, 0x50, 0x8a, 0x94, 0x20, 0x4f,
	0x20, 0x41, 0x8c, 0x5a, 0x09, 0xa4, 0x8e, 0xbd, 0x03, 0x22, 0xb1, 0xb0, 0x54, 0x4e, 0xe2, 0xa6,
	0x11, 0x4e, 0xbe, 0xc8, 0x4e, 0x2b, 0x7a, 0x17, 0x8c, 0x8c, 0xdd, 0xb9, 0x11, 0x16, 0xa4, 0x8e,
	0x4c, 0x11, 0x6a, 0x17, 0xe6, 0x5c, 0x01, 0x4a, 0x9c, 0xb4, 0xfc, 0x74, 0x61, 0xfb, 0x9c, 0xf7,
	0x79, 0x92, 0xf7, 0x53, 0xac, 0x9d, 0x7b, 0x4c, 0x86, 0x20, 0x49, 0x2c, 0x60, 0x14, 0x70, 0x26,
	0x09, 0x67, 0x3e, 0x75, 0x67, 0x64, 0xda, 0xed, 0x9d, 0x91, 0x10, 0x3c, 0xc6, 0xe5, 0x50, 0x30,
	0x4e, 0x93, 0x00, 0x22, 0x39, 0x0e, 0x62, 0x69, 0xc5, 0x02, 0x12, 0xd0, 0x0f, 0x94, 0x67, 0x55,
	0x9e, 0xa5, 0x3c, 0x2b, 0xf7, 0x3a, 0x2d, 0x1f, 0x7c, 0x28, 0x38, 0x92, 0x4f, 0x4a, 0xe9, 0xec,
	0xfb, 0x00, 0x3e, 0x67, 0xa4, 0x38, 0x39, 0x93, 0x11, 0xa1, 0xd1, 0xac, 0x8c, 0xcc, 0xcf, 0x51,
	0x12, 0x84, 0x4c, 0x26, 0x34, 0x8c, 0x2b, 0xd7, 0x85, 0xfc, 0x73, 0x43, 0xf5, 0x52, 0x75, 0x50,
	0x11, 0x7e, 0x40, 0xda, 0x1f, 0xfb, 0x5d, 0x43, 0xfd, 0x44, 0xfb, 0xed, 0x0a, 0x46, 0x13, 0x10,
	0x6d, 0x74, 0x88, 0x8e, 0x9a, 0x03, 0x3d, 0x4b, 0xcd, 0x7f, 0x33, 0x1a, 0xf2, 0x3e, 0x2e, 0x03,
	0x6c, 0x57, 0x88, 0xde, 0xd5, 0x9a, 0x82, 0xb9, 0x41, 0x1c, 0xb0, 0x28, 0x69, 0xff, 0x28, 0xf8,
	0x56, 0x96, 0x9a, 0x3b, 0x8a, 0x5f, 0x47, 0xd8, 0xde, 0x60, 0x3a, 0xd1, 0x1a, 0x72, 0xe2, 0xc8,
	0x98, 0xba, 0xac, 0xfd, 0xb3, 0x50, 0xf6, 0xb2, 0xd4, 0xfc, 0xaf, 0x94, 0x2a, 0xc1, 0xf6, 0x1a,
	0xea, 0x37, 0xee, 0xe7, 0x26, 0x7a, 0x9d, 0x9b, 0x08, 0x3f, 0x21, 0xad, 0x79, 0x25, 0x99, 0x18,
	0x70, 0x70, 0x6f, 0xf2, 0xaa, 0x4e, 0x3e, 0xb0, 0x2d, 0x55, 0xcb, 0x00, 0xdb, 0x15, 0xb2, 0xa1,
	0xbd, 0xb2, 0xe8, 0x17, 0xda, 0x5b, 0xd3, 0x9e, 0x7e, 0xac, 0xd5, 0x05, 0xa3, 0x12, 0xa2, 0xb2,
	0xe2, 0x6e, 0x96, 0x9a, 0x7f, 0xab, 0xad, 0xf2, 0xe7, 0xd8, 0x2e, 0x81, 0x0f, 0xfb, 0xfc, 0xfa,
	0xd6, 0x3e, 0x83, 0xcb, 0xc7, 0xa5, 0x81, 0x16, 0x4b, 0x03, 0xbd, 0x2c, 0x0d, 0x74, 0xb7, 0x32,
	0x6a, 0x8b, 0x95, 0x51, 0x7b, 0x5e, 0x19, 0xb5, 0xeb, 0x0b, 0x3f, 0x48, 0xc6, 0x13, 0xc7, 0x72,
	0x21, 0x24, 0xea, 0xb2, 0x9c, 0x72, 0xea, 0xc8, 0x72, 0x26, 0xd3, 0x2e, 0xb9, 0xdd, 0x7a, 0xeb,
	0x9c, 0x7a, 0xf1, 0x5f, 0x7b, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x95, 0x70, 0xc1, 0x81, 0x9b,
	0x02, 0x00, 0x00,
}

func (this *Relationship) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Relationship)
	if !ok {
		that2, ok := that.(Relationship)
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
	if this.Creator != that1.Creator {
		return false
	}
	if this.Recipient != that1.Recipient {
		return false
	}
	if this.Subspace != that1.Subspace {
		return false
	}
	return true
}
func (this *UserBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserBlock)
	if !ok {
		that2, ok := that.(UserBlock)
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
	if this.Blocker != that1.Blocker {
		return false
	}
	if this.Blocked != that1.Blocked {
		return false
	}
	if this.Reason != that1.Reason {
		return false
	}
	if this.Subspace != that1.Subspace {
		return false
	}
	return true
}
func (m *Relationship) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Relationship) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Relationship) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintModelsRelationships(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintModelsRelationships(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintModelsRelationships(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UserBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintModelsRelationships(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintModelsRelationships(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Blocked) > 0 {
		i -= len(m.Blocked)
		copy(dAtA[i:], m.Blocked)
		i = encodeVarintModelsRelationships(dAtA, i, uint64(len(m.Blocked)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Blocker) > 0 {
		i -= len(m.Blocker)
		copy(dAtA[i:], m.Blocker)
		i = encodeVarintModelsRelationships(dAtA, i, uint64(len(m.Blocker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsRelationships(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsRelationships(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Relationship) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovModelsRelationships(uint64(l))
	}
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovModelsRelationships(uint64(l))
	}
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovModelsRelationships(uint64(l))
	}
	return n
}

func (m *UserBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Blocker)
	if l > 0 {
		n += 1 + l + sovModelsRelationships(uint64(l))
	}
	l = len(m.Blocked)
	if l > 0 {
		n += 1 + l + sovModelsRelationships(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovModelsRelationships(uint64(l))
	}
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovModelsRelationships(uint64(l))
	}
	return n
}

func sovModelsRelationships(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsRelationships(x uint64) (n int) {
	return sovModelsRelationships(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Relationship) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsRelationships
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
			return fmt.Errorf("proto: Relationship: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Relationship: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsRelationships
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
				return ErrInvalidLengthModelsRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsRelationships
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
				return ErrInvalidLengthModelsRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsRelationships
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
				return ErrInvalidLengthModelsRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsRelationships(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsRelationships
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
func (m *UserBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsRelationships
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
			return fmt.Errorf("proto: UserBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsRelationships
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
				return ErrInvalidLengthModelsRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blocked", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsRelationships
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
				return ErrInvalidLengthModelsRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Blocked = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsRelationships
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
				return ErrInvalidLengthModelsRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsRelationships
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
				return ErrInvalidLengthModelsRelationships
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsRelationships
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsRelationships(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsRelationships
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
func skipModelsRelationships(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsRelationships
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
					return 0, ErrIntOverflowModelsRelationships
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
					return 0, ErrIntOverflowModelsRelationships
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
				return 0, ErrInvalidLengthModelsRelationships
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsRelationships
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsRelationships
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsRelationships        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsRelationships          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsRelationships = fmt.Errorf("proto: unexpected end of group")
)
