// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/posts/v1beta1/reactions.proto

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

// RegisteredReaction represents a registered reaction that can be referenced
// by its shortCode inside post reactions
type RegisteredReaction struct {
	ShortCode string `protobuf:"bytes,1,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty" yaml:"short_code"`
	Value     string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
	Subspace  string `protobuf:"bytes,3,opt,name=subspace,proto3" json:"subspace,omitempty" yaml:"subspace"`
	Creator   string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *RegisteredReaction) Reset()         { *m = RegisteredReaction{} }
func (m *RegisteredReaction) String() string { return proto.CompactTextString(m) }
func (*RegisteredReaction) ProtoMessage()    {}
func (*RegisteredReaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_279d343e0105421c, []int{0}
}
func (m *RegisteredReaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteredReaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteredReaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteredReaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredReaction.Merge(m, src)
}
func (m *RegisteredReaction) XXX_Size() int {
	return m.Size()
}
func (m *RegisteredReaction) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredReaction.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredReaction proto.InternalMessageInfo

func (m *RegisteredReaction) GetShortCode() string {
	if m != nil {
		return m.ShortCode
	}
	return ""
}

func (m *RegisteredReaction) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RegisteredReaction) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

func (m *RegisteredReaction) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// PostReaction is a struct of a user reaction to a post
type PostReaction struct {
	PostID    string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty" yaml:"post_id"`
	ShortCode string `protobuf:"bytes,2,opt,name=short_code,json=shortCode,proto3" json:"short_code,omitempty" yaml:"short_code"`
	Value     string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
	Owner     string `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
}

func (m *PostReaction) Reset()         { *m = PostReaction{} }
func (m *PostReaction) String() string { return proto.CompactTextString(m) }
func (*PostReaction) ProtoMessage()    {}
func (*PostReaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_279d343e0105421c, []int{1}
}
func (m *PostReaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostReaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostReaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostReaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReaction.Merge(m, src)
}
func (m *PostReaction) XXX_Size() int {
	return m.Size()
}
func (m *PostReaction) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReaction.DiscardUnknown(m)
}

var xxx_messageInfo_PostReaction proto.InternalMessageInfo

func (m *PostReaction) GetPostID() string {
	if m != nil {
		return m.PostID
	}
	return ""
}

func (m *PostReaction) GetShortCode() string {
	if m != nil {
		return m.ShortCode
	}
	return ""
}

func (m *PostReaction) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *PostReaction) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisteredReaction)(nil), "desmos.posts.v1beta1.RegisteredReaction")
	proto.RegisterType((*PostReaction)(nil), "desmos.posts.v1beta1.PostReaction")
}

func init() {
	proto.RegisterFile("desmos/posts/v1beta1/reactions.proto", fileDescriptor_279d343e0105421c)
}

var fileDescriptor_279d343e0105421c = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3d, 0x6b, 0xeb, 0x30,
	0x18, 0x85, 0xa3, 0xe4, 0xe6, 0x4b, 0x84, 0xfb, 0xe1, 0x9b, 0x0b, 0xe6, 0x72, 0xb1, 0x2f, 0xa2,
	0x84, 0x0e, 0xad, 0x45, 0xfa, 0x35, 0x64, 0x4c, 0xbb, 0x64, 0x0b, 0x1a, 0xbb, 0x04, 0xd9, 0x16,
	0x8e, 0x21, 0x89, 0x8c, 0xa5, 0xa4, 0xcd, 0xbf, 0xe8, 0xd8, 0x31, 0x3f, 0xa7, 0x63, 0x96, 0x42,
	0xa1, 0x60, 0x8a, 0xb3, 0x74, 0xf6, 0x2f, 0x28, 0x96, 0x95, 0xa4, 0x74, 0x28, 0x74, 0x7b, 0xa5,
	0xf3, 0xbc, 0x2f, 0x9c, 0xc3, 0x81, 0x07, 0x3e, 0x13, 0x53, 0x2e, 0x70, 0xc4, 0x85, 0x14, 0x78,
	0xd1, 0x75, 0x99, 0xa4, 0x5d, 0x1c, 0x33, 0xea, 0xc9, 0x90, 0xcf, 0x84, 0x13, 0xc5, 0x5c, 0x72,
	0xa3, 0x5d, 0x50, 0x8e, 0xa2, 0x1c, 0x4d, 0xfd, 0x6d, 0x07, 0x3c, 0xe0, 0x0a, 0xc0, 0xf9, 0x54,
	0xb0, 0xe8, 0x19, 0x40, 0x83, 0xb0, 0x20, 0x14, 0x92, 0xc5, 0xcc, 0x27, 0xfa, 0x92, 0x71, 0x06,
	0xa1, 0x18, 0xf3, 0x58, 0x8e, 0x3c, 0xee, 0x33, 0x13, 0xfc, 0x07, 0x87, 0xcd, 0xfe, 0x9f, 0x2c,
	0xb1, 0x7f, 0x2d, 0xe9, 0x74, 0xd2, 0x43, 0x7b, 0x0d, 0x91, 0xa6, 0x7a, 0x5c, 0x72, 0x9f, 0x19,
	0x1d, 0x58, 0x5d, 0xd0, 0xc9, 0x9c, 0x99, 0x65, 0xb5, 0xf0, 0x33, 0x4b, 0xec, 0x56, 0xb1, 0xa0,
	0xbe, 0x11, 0x29, 0x64, 0x03, 0xc3, 0x86, 0x98, 0xbb, 0x22, 0xa2, 0x1e, 0x33, 0x2b, 0x0a, 0xfd,
	0x9d, 0x25, 0xf6, 0x0f, 0x7d, 0x5b, 0x2b, 0x88, 0xec, 0x20, 0xe3, 0x08, 0xd6, 0xbd, 0x98, 0x51,
	0xc9, 0x63, 0xf3, 0x9b, 0xe2, 0x8d, 0x2c, 0xb1, 0xbf, 0x17, 0xbc, 0x16, 0x10, 0xd9, 0x22, 0xbd,
	0xc6, 0xfd, 0xca, 0x06, 0xaf, 0x2b, 0x1b, 0xa0, 0x47, 0x00, 0x5b, 0x43, 0x2e, 0xe4, 0xce, 0xd7,
	0x39, 0xac, 0xe7, 0xa9, 0x8c, 0x42, 0x5f, 0x9b, 0xfa, 0x97, 0x26, 0x76, 0x2d, 0x47, 0x06, 0x57,
	0xfb, 0x93, 0x1a, 0x41, 0xa4, 0x96, 0x4f, 0x03, 0xff, 0x43, 0x1c, 0xe5, 0xaf, 0xc6, 0x51, 0xf9,
	0x3c, 0x8e, 0x0e, 0xac, 0xf2, 0x9b, 0x19, 0xdb, 0x7a, 0x7b, 0xc7, 0xa9, 0x6f, 0x44, 0x0a, 0x79,
	0xef, 0xab, 0x3f, 0x7c, 0x48, 0x2d, 0xb0, 0x4e, 0x2d, 0xf0, 0x92, 0x5a, 0xe0, 0x6e, 0x63, 0x95,
	0xd6, 0x1b, 0xab, 0xf4, 0xb4, 0xb1, 0x4a, 0xd7, 0x17, 0x41, 0x28, 0xc7, 0x73, 0xd7, 0xf1, 0xf8,
	0x14, 0x17, 0x35, 0x38, 0x9e, 0x50, 0x57, 0xe8, 0x19, 0x2f, 0x4e, 0xf0, 0x2d, 0x16, 0x92, 0x06,
	0xe1, 0x2c, 0xd0, 0x2d, 0x92, 0xcb, 0x88, 0x09, 0xb7, 0xa6, 0xea, 0x70, 0xfa, 0x16, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0xa1, 0x87, 0xff, 0x62, 0x02, 0x00, 0x00,
}

func (this *RegisteredReaction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RegisteredReaction)
	if !ok {
		that2, ok := that.(RegisteredReaction)
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
	if this.ShortCode != that1.ShortCode {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Subspace != that1.Subspace {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	return true
}
func (this *PostReaction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PostReaction)
	if !ok {
		that2, ok := that.(PostReaction)
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
	if this.PostID != that1.PostID {
		return false
	}
	if this.ShortCode != that1.ShortCode {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	return true
}
func (m *RegisteredReaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteredReaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisteredReaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintReactions(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintReactions(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintReactions(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ShortCode) > 0 {
		i -= len(m.ShortCode)
		copy(dAtA[i:], m.ShortCode)
		i = encodeVarintReactions(dAtA, i, uint64(len(m.ShortCode)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostReaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostReaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostReaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintReactions(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintReactions(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ShortCode) > 0 {
		i -= len(m.ShortCode)
		copy(dAtA[i:], m.ShortCode)
		i = encodeVarintReactions(dAtA, i, uint64(len(m.ShortCode)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PostID) > 0 {
		i -= len(m.PostID)
		copy(dAtA[i:], m.PostID)
		i = encodeVarintReactions(dAtA, i, uint64(len(m.PostID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReactions(dAtA []byte, offset int, v uint64) int {
	offset -= sovReactions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisteredReaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ShortCode)
	if l > 0 {
		n += 1 + l + sovReactions(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovReactions(uint64(l))
	}
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovReactions(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovReactions(uint64(l))
	}
	return n
}

func (m *PostReaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostID)
	if l > 0 {
		n += 1 + l + sovReactions(uint64(l))
	}
	l = len(m.ShortCode)
	if l > 0 {
		n += 1 + l + sovReactions(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovReactions(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovReactions(uint64(l))
	}
	return n
}

func sovReactions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReactions(x uint64) (n int) {
	return sovReactions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisteredReaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReactions
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
			return fmt.Errorf("proto: RegisteredReaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteredReaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReactions
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
				return ErrInvalidLengthReactions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReactions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReactions
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
				return ErrInvalidLengthReactions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReactions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReactions
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
				return ErrInvalidLengthReactions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReactions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReactions
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
				return ErrInvalidLengthReactions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReactions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReactions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReactions
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
func (m *PostReaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReactions
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
			return fmt.Errorf("proto: PostReaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostReaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReactions
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
				return ErrInvalidLengthReactions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReactions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PostID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ShortCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReactions
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
				return ErrInvalidLengthReactions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReactions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ShortCode = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReactions
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
				return ErrInvalidLengthReactions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReactions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReactions
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
				return ErrInvalidLengthReactions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReactions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReactions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReactions
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
func skipReactions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReactions
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
					return 0, ErrIntOverflowReactions
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
					return 0, ErrIntOverflowReactions
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
				return 0, ErrInvalidLengthReactions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReactions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReactions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReactions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReactions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReactions = fmt.Errorf("proto: unexpected end of group")
)
