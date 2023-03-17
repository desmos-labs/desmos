// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v3/msgs_profile.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

// MsgSaveProfile represents a message to save a profile.
type MsgSaveProfile struct {
	// DTag of the profile. If it shouldn't be changed, [do-no-modify] can be used
	// instead.
	DTag string `protobuf:"bytes,1,opt,name=dtag,proto3" json:"dtag,omitempty" yaml:"dtag"`
	// Nickname of the profile. If it shouldn't be changed, [do-no-modify] can be
	// used instead.
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty" yaml:"nickname"`
	// Bio of the profile. If it shouldn't be changed, [do-no-modify] can be used
	// instead.
	Bio string `protobuf:"bytes,3,opt,name=bio,proto3" json:"bio,omitempty" yaml:"bio"`
	// URL to the profile picture. If it shouldn't be changed, [do-no-modify] can
	// be used instead.
	ProfilePicture string `protobuf:"bytes,4,opt,name=profile_picture,json=profilePicture,proto3" json:"profile_picture,omitempty" yaml:"profile_picture"`
	// URL to the profile cover. If it shouldn't be changed, [do-no-modify] can be
	// used instead.
	CoverPicture string `protobuf:"bytes,5,opt,name=cover_picture,json=coverPicture,proto3" json:"cover_picture,omitempty" yaml:"cover_picture"`
	// Address of the user associated to the profile
	Creator string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *MsgSaveProfile) Reset()         { *m = MsgSaveProfile{} }
func (m *MsgSaveProfile) String() string { return proto.CompactTextString(m) }
func (*MsgSaveProfile) ProtoMessage()    {}
func (*MsgSaveProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea75cf4ca5bb3a3, []int{0}
}
func (m *MsgSaveProfile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveProfile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveProfile.Merge(m, src)
}
func (m *MsgSaveProfile) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveProfile.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveProfile proto.InternalMessageInfo

func (m *MsgSaveProfile) GetDTag() string {
	if m != nil {
		return m.DTag
	}
	return ""
}

func (m *MsgSaveProfile) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *MsgSaveProfile) GetBio() string {
	if m != nil {
		return m.Bio
	}
	return ""
}

func (m *MsgSaveProfile) GetProfilePicture() string {
	if m != nil {
		return m.ProfilePicture
	}
	return ""
}

func (m *MsgSaveProfile) GetCoverPicture() string {
	if m != nil {
		return m.CoverPicture
	}
	return ""
}

func (m *MsgSaveProfile) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// MsgSaveProfileResponse defines the Msg/SaveProfile response type.
type MsgSaveProfileResponse struct {
}

func (m *MsgSaveProfileResponse) Reset()         { *m = MsgSaveProfileResponse{} }
func (m *MsgSaveProfileResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSaveProfileResponse) ProtoMessage()    {}
func (*MsgSaveProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea75cf4ca5bb3a3, []int{1}
}
func (m *MsgSaveProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveProfileResponse.Merge(m, src)
}
func (m *MsgSaveProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveProfileResponse proto.InternalMessageInfo

// MsgDeleteProfile represents the message used to delete an existing profile.
type MsgDeleteProfile struct {
	// Address associated to the profile to be deleted
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
}

func (m *MsgDeleteProfile) Reset()         { *m = MsgDeleteProfile{} }
func (m *MsgDeleteProfile) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProfile) ProtoMessage()    {}
func (*MsgDeleteProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea75cf4ca5bb3a3, []int{2}
}
func (m *MsgDeleteProfile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProfile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProfile.Merge(m, src)
}
func (m *MsgDeleteProfile) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProfile.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProfile proto.InternalMessageInfo

func (m *MsgDeleteProfile) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

// MsgDeleteProfileResponse defines the Msg/DeleteProfile response type.
type MsgDeleteProfileResponse struct {
}

func (m *MsgDeleteProfileResponse) Reset()         { *m = MsgDeleteProfileResponse{} }
func (m *MsgDeleteProfileResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteProfileResponse) ProtoMessage()    {}
func (*MsgDeleteProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5ea75cf4ca5bb3a3, []int{3}
}
func (m *MsgDeleteProfileResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteProfileResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteProfileResponse.Merge(m, src)
}
func (m *MsgDeleteProfileResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteProfileResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSaveProfile)(nil), "desmos.profiles.v3.MsgSaveProfile")
	proto.RegisterType((*MsgSaveProfileResponse)(nil), "desmos.profiles.v3.MsgSaveProfileResponse")
	proto.RegisterType((*MsgDeleteProfile)(nil), "desmos.profiles.v3.MsgDeleteProfile")
	proto.RegisterType((*MsgDeleteProfileResponse)(nil), "desmos.profiles.v3.MsgDeleteProfileResponse")
}

func init() {
	proto.RegisterFile("desmos/profiles/v3/msgs_profile.proto", fileDescriptor_5ea75cf4ca5bb3a3)
}

var fileDescriptor_5ea75cf4ca5bb3a3 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x37, 0xed, 0x5a, 0x75, 0xd4, 0xad, 0x8c, 0xa5, 0x8e, 0x39, 0x24, 0x65, 0x40, 0x14,
	0xb4, 0x3b, 0xc8, 0x7a, 0x12, 0x04, 0xa9, 0xbd, 0x49, 0xa1, 0x44, 0x4f, 0x5e, 0x96, 0x49, 0xf6,
	0x75, 0x0c, 0x26, 0xfb, 0x62, 0x66, 0x36, 0xb8, 0xdf, 0xc2, 0xcf, 0xe4, 0xc9, 0x63, 0x8f, 0x9e,
	0x82, 0xec, 0x7e, 0x83, 0x7c, 0x02, 0xc9, 0x4c, 0xb2, 0xda, 0x45, 0xe9, 0xed, 0xbd, 0xf7, 0xff,
	0xfd, 0xff, 0x3c, 0x1e, 0x8f, 0x3c, 0x9e, 0x81, 0xce, 0x51, 0x8b, 0xa2, 0xc4, 0x8b, 0x34, 0x03,
	0x2d, 0xaa, 0x89, 0xc8, 0xb5, 0xd2, 0xd3, 0x6e, 0x30, 0x2e, 0x4a, 0x34, 0x48, 0xa9, 0xc3, 0xc6,
	0x3d, 0x36, 0xae, 0x26, 0xfe, 0x81, 0x42, 0x85, 0x56, 0x16, 0x6d, 0xe5, 0x48, 0xff, 0x91, 0x42,
	0x54, 0x19, 0x08, 0xdb, 0xc5, 0x8b, 0x0b, 0x21, 0xe7, 0xcb, 0x5e, 0x4a, 0xb0, 0x0d, 0x99, 0x3a,
	0x8f, 0x6b, 0x3a, 0xe9, 0xc9, 0xbf, 0xd6, 0xc0, 0x19, 0x64, 0x5b, 0x8b, 0xf8, 0xc7, 0xff, 0x07,
	0x67, 0x46, 0xaa, 0x69, 0x09, 0x5f, 0x16, 0xa0, 0x4d, 0x97, 0xcb, 0xbf, 0xef, 0x90, 0xd1, 0x99,
	0x56, 0xef, 0x65, 0x05, 0xe7, 0xce, 0x41, 0x9f, 0x91, 0x61, 0x4b, 0x32, 0xef, 0xc8, 0x7b, 0x7a,
	0xfb, 0xe4, 0xe1, 0xaa, 0x0e, 0x87, 0xa7, 0x1f, 0xa4, 0x6a, 0xea, 0xf0, 0xce, 0x52, 0xe6, 0xd9,
	0x2b, 0xde, 0xaa, 0x3c, 0xb2, 0x10, 0x15, 0xe4, 0xd6, 0x3c, 0x4d, 0x3e, 0xcf, 0x65, 0x0e, 0x6c,
	0xc7, 0x1a, 0x1e, 0x34, 0x75, 0xb8, 0xef, 0xc0, 0x5e, 0xe1, 0xd1, 0x06, 0xa2, 0x47, 0x64, 0x37,
	0x4e, 0x91, 0xed, 0x5a, 0x76, 0xd4, 0xd4, 0x21, 0x71, 0x6c, 0x9c, 0x22, 0x8f, 0x5a, 0x89, 0xbe,
	0x25, 0xfb, 0xdd, 0xf2, 0xd3, 0x22, 0x4d, 0xcc, 0xa2, 0x04, 0x36, 0xb4, 0xb4, 0xdf, 0xd4, 0xe1,
	0xa1, 0xa3, 0xb7, 0x00, 0x1e, 0x8d, 0xba, 0xc9, 0xb9, 0x1b, 0xd0, 0xd7, 0xe4, 0x5e, 0x82, 0x15,
	0x94, 0x9b, 0x88, 0x1b, 0x36, 0x82, 0x35, 0x75, 0x78, 0xe0, 0x22, 0xae, 0xc8, 0x3c, 0xba, 0x6b,
	0xfb, 0xde, 0xfe, 0x9c, 0xdc, 0x4c, 0x4a, 0x90, 0x06, 0x4b, 0xb6, 0x67, 0x8d, 0xb4, 0xa9, 0xc3,
	0x51, 0x67, 0x74, 0x02, 0x8f, 0x7a, 0x84, 0x33, 0x72, 0x78, 0xf5, 0x86, 0x11, 0xe8, 0x02, 0xe7,
	0x1a, 0xf8, 0x1b, 0x72, 0xff, 0x4c, 0xab, 0x53, 0xc8, 0xc0, 0x6c, 0xee, 0xfb, 0x57, 0xb6, 0x77,
	0x7d, 0xb6, 0x4f, 0xd8, 0x76, 0x42, 0x9f, 0x7e, 0xf2, 0xee, 0xc7, 0x2a, 0xf0, 0x2e, 0x57, 0x81,
	0xf7, 0x6b, 0x15, 0x78, 0xdf, 0xd6, 0xc1, 0xe0, 0x72, 0x1d, 0x0c, 0x7e, 0xae, 0x83, 0xc1, 0xc7,
	0x17, 0x2a, 0x35, 0x9f, 0x16, 0xf1, 0x38, 0xc1, 0x5c, 0xb8, 0x87, 0x38, 0xce, 0x64, 0xac, 0xbb,
	0x5a, 0x54, 0x2f, 0xc5, 0xd7, 0x3f, 0x1f, 0x62, 0x96, 0x05, 0xe8, 0x78, 0xcf, 0x3e, 0xc4, 0xe4,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x25, 0x9b, 0x00, 0xf1, 0x02, 0x00, 0x00,
}

func (m *MsgSaveProfile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveProfile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveProfile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CoverPicture) > 0 {
		i -= len(m.CoverPicture)
		copy(dAtA[i:], m.CoverPicture)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.CoverPicture)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ProfilePicture) > 0 {
		i -= len(m.ProfilePicture)
		copy(dAtA[i:], m.ProfilePicture)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.ProfilePicture)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bio) > 0 {
		i -= len(m.Bio)
		copy(dAtA[i:], m.Bio)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Bio)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DTag) > 0 {
		i -= len(m.DTag)
		copy(dAtA[i:], m.DTag)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.DTag)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSaveProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProfile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProfile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProfile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMsgsProfile(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteProfileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteProfileResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteProfileResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintMsgsProfile(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsgsProfile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSaveProfile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DTag)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Bio)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.ProfilePicture)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.CoverPicture)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	return n
}

func (m *MsgSaveProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteProfile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMsgsProfile(uint64(l))
	}
	return n
}

func (m *MsgDeleteProfileResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMsgsProfile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsgsProfile(x uint64) (n int) {
	return sovMsgsProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSaveProfile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgSaveProfile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveProfile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DTag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProfilePicture", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProfilePicture = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoverPicture", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CoverPicture = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgSaveProfileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgSaveProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgDeleteProfile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgDeleteProfile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProfile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsgsProfile
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
				return ErrInvalidLengthMsgsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMsgsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func (m *MsgDeleteProfileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsgsProfile
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
			return fmt.Errorf("proto: MsgDeleteProfileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteProfileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMsgsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsgsProfile
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
func skipMsgsProfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsgsProfile
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
					return 0, ErrIntOverflowMsgsProfile
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
					return 0, ErrIntOverflowMsgsProfile
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
				return 0, ErrInvalidLengthMsgsProfile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsgsProfile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsgsProfile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsgsProfile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsgsProfile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsgsProfile = fmt.Errorf("proto: unexpected end of group")
)
