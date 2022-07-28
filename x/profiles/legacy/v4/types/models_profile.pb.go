// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v1beta1/models_profile.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Profile represents a generic first on Desmos, containing the information of a
// single user
type Profile struct {
	// Account represents the base Cosmos account associated with this profile
	Account *types.Any `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// DTag represents the unique tag of this profile
	DTag string `protobuf:"bytes,2,opt,name=dtag,proto3" json:"dtag,omitempty" yaml:"dtag"`
	// Nickname contains the custom human readable name of the profile
	Nickname string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty" yaml:"nickname"`
	// Bio contains the biography of the profile
	Bio string `protobuf:"bytes,4,opt,name=bio,proto3" json:"bio,omitempty" yaml:"bio"`
	// Pictures contains the data about the pictures associated with he profile
	Pictures Pictures `protobuf:"bytes,5,opt,name=pictures,proto3" json:"pictures" yaml:"pictures"`
	// CreationTime represents the time in which the profile has been created
	CreationDate time.Time `protobuf:"bytes,6,opt,name=creation_date,json=creationDate,proto3,stdtime" json:"creation_date" yaml:"creation_date"`
}

func (m *Profile) Reset()      { *m = Profile{} }
func (*Profile) ProtoMessage() {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a19232e029005b86, []int{0}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return m.Size()
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

// Pictures contains the data of a user profile's related pictures
type Pictures struct {
	// Profile contains the URL to the profile picture
	Profile string `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty" yaml:"profile"`
	// Cover contains the URL to the cover picture
	Cover string `protobuf:"bytes,2,opt,name=cover,proto3" json:"cover,omitempty" yaml:"cover"`
}

func (m *Pictures) Reset()         { *m = Pictures{} }
func (m *Pictures) String() string { return proto.CompactTextString(m) }
func (*Pictures) ProtoMessage()    {}
func (*Pictures) Descriptor() ([]byte, []int) {
	return fileDescriptor_a19232e029005b86, []int{1}
}
func (m *Pictures) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pictures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pictures.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pictures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pictures.Merge(m, src)
}
func (m *Pictures) XXX_Size() int {
	return m.Size()
}
func (m *Pictures) XXX_DiscardUnknown() {
	xxx_messageInfo_Pictures.DiscardUnknown(m)
}

var xxx_messageInfo_Pictures proto.InternalMessageInfo

func (m *Pictures) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *Pictures) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func init() {
	proto.RegisterType((*Profile)(nil), "desmos.profiles.v1beta1.Profile")
	proto.RegisterType((*Pictures)(nil), "desmos.profiles.v1beta1.Pictures")
}

func init() {
	proto.RegisterFile("desmos/profiles/v1beta1/models_profile.proto", fileDescriptor_a19232e029005b86)
}

var fileDescriptor_a19232e029005b86 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xb1, 0x6e, 0xd3, 0x40,
	0x1c, 0xc6, 0x6d, 0x92, 0x36, 0xee, 0x35, 0x14, 0x74, 0x44, 0x8a, 0xc9, 0xe0, 0x0b, 0x37, 0xa0,
	0x4a, 0xb4, 0x3e, 0x15, 0x98, 0x02, 0x4b, 0xad, 0x2e, 0x6c, 0x95, 0x55, 0x18, 0x58, 0xa2, 0xb3,
	0x73, 0x35, 0x06, 0xdb, 0x17, 0xd9, 0x97, 0x08, 0xbf, 0x01, 0x03, 0x43, 0xc7, 0x8e, 0x79, 0x08,
	0x1e, 0xa2, 0x62, 0xea, 0xc8, 0x64, 0x50, 0xb2, 0x30, 0xfb, 0x09, 0x90, 0x7d, 0x77, 0xad, 0x00,
	0x75, 0xbb, 0xfb, 0x7f, 0xbf, 0xff, 0x77, 0x9f, 0x3e, 0x1b, 0x1c, 0xcc, 0x58, 0x91, 0xf2, 0x82,
	0xcc, 0x73, 0x7e, 0x1e, 0x27, 0xac, 0x20, 0xcb, 0xa3, 0x80, 0x09, 0x7a, 0x44, 0x52, 0x3e, 0x63,
	0x49, 0x31, 0x55, 0x73, 0x77, 0x9e, 0x73, 0xc1, 0xe1, 0x50, 0xd2, 0xae, 0xa6, 0x5d, 0x45, 0x8f,
	0x06, 0x11, 0x8f, 0x78, 0xcb, 0x90, 0xe6, 0x24, 0xf1, 0xd1, 0xe3, 0x88, 0xf3, 0x28, 0x61, 0xa4,
	0xbd, 0x05, 0x8b, 0x73, 0x42, 0xb3, 0x52, 0x49, 0xe8, 0x5f, 0x49, 0xc4, 0x29, 0x2b, 0x04, 0x4d,
	0xe7, 0x7a, 0x37, 0xe4, 0xcd, 0x53, 0x53, 0x69, 0x2a, 0x2f, 0x52, 0xc2, 0x5f, 0x3b, 0xa0, 0x77,
	0x2a, 0x13, 0xc0, 0xd7, 0xa0, 0x47, 0xc3, 0x90, 0x2f, 0x32, 0x61, 0x9b, 0x63, 0x73, 0x7f, 0xf7,
	0xf9, 0xc0, 0x95, 0xce, 0xae, 0x76, 0x76, 0x8f, 0xb3, 0xd2, 0xeb, 0x7f, 0xff, 0x76, 0x68, 0x1d,
	0x4b, 0xf0, 0x8d, 0xaf, 0x57, 0xe0, 0x33, 0xd0, 0x9d, 0x09, 0x1a, 0xd9, 0xf7, 0xc6, 0xe6, 0xfe,
	0x8e, 0x37, 0x5c, 0x57, 0xa8, 0x7b, 0x72, 0x46, 0xa3, 0xba, 0x42, 0xbb, 0x25, 0x4d, 0x93, 0x09,
	0x6e, 0x54, 0xec, 0xb7, 0x10, 0x24, 0xc0, 0xca, 0xe2, 0xf0, 0x53, 0x46, 0x53, 0x66, 0x77, 0xda,
	0x85, 0x47, 0x75, 0x85, 0x1e, 0x48, 0x50, 0x2b, 0xd8, 0xbf, 0x81, 0xe0, 0x18, 0x74, 0x82, 0x98,
	0xdb, 0xdd, 0x96, 0xdd, 0xab, 0x2b, 0x04, 0x24, 0x1b, 0xc4, 0x1c, 0xfb, 0x8d, 0x04, 0xdf, 0x01,
	0x6b, 0x1e, 0x87, 0x62, 0x91, 0xb3, 0xc2, 0xde, 0x6a, 0xe3, 0x3f, 0x71, 0xef, 0xa8, 0xd8, 0x3d,
	0x55, 0xa0, 0x37, 0xbc, 0xaa, 0x90, 0x71, 0xfb, 0xb2, 0x36, 0xc0, 0xfe, 0x8d, 0x17, 0xa4, 0xe0,
	0x7e, 0x98, 0x33, 0x2a, 0x62, 0x9e, 0x4d, 0x67, 0x54, 0x30, 0x7b, 0xbb, 0x35, 0x1f, 0xfd, 0xd7,
	0xcd, 0x99, 0x6e, 0xdd, 0x1b, 0x2b, 0xd7, 0x81, 0x74, 0xfd, 0x6b, 0x1d, 0x5f, 0xfc, 0x44, 0xa6,
	0xdf, 0xd7, 0xb3, 0x13, 0x2a, 0xd8, 0xc4, 0xfa, 0xb2, 0x42, 0xc6, 0xe5, 0x0a, 0x19, 0xf8, 0x23,
	0xb0, 0x74, 0x36, 0x78, 0x00, 0x7a, 0x2a, 0x78, 0xfb, 0x39, 0x76, 0x3c, 0x58, 0x57, 0x68, 0x4f,
	0x05, 0x95, 0x02, 0xf6, 0x35, 0x02, 0x9f, 0x82, 0xad, 0x90, 0x2f, 0x59, 0xae, 0xfa, 0x7f, 0x58,
	0x57, 0xa8, 0xaf, 0x9e, 0x6f, 0xc6, 0xd8, 0x97, 0xf2, 0xc4, 0xba, 0x5c, 0x21, 0xf3, 0xf7, 0x0a,
	0x99, 0xde, 0xdb, 0xab, 0xb5, 0x63, 0x5e, 0xaf, 0x1d, 0xf3, 0xd7, 0xda, 0x31, 0x2f, 0x36, 0x8e,
	0x71, 0xbd, 0x71, 0x8c, 0x1f, 0x1b, 0xc7, 0x78, 0xff, 0x2a, 0x8a, 0xc5, 0x87, 0x45, 0xe0, 0x86,
	0x3c, 0x25, 0xb2, 0xc2, 0xc3, 0x84, 0x06, 0x85, 0x3a, 0x93, 0xe5, 0x4b, 0xf2, 0xf9, 0xf6, 0x27,
	0x4f, 0x58, 0x44, 0xc3, 0xb2, 0x19, 0x8a, 0x72, 0xce, 0x8a, 0x60, 0xbb, 0x2d, 0xe4, 0xc5, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x96, 0x94, 0x7e, 0x0e, 0x03, 0x00, 0x00,
}

func (this *Pictures) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Pictures)
	if !ok {
		that2, ok := that.(Pictures)
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
	if this.Profile != that1.Profile {
		return false
	}
	if this.Cover != that1.Cover {
		return false
	}
	return true
}
func (m *Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Profile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreationDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationDate):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintModelsProfile(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x32
	{
		size, err := m.Pictures.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintModelsProfile(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Bio) > 0 {
		i -= len(m.Bio)
		copy(dAtA[i:], m.Bio)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.Bio)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Nickname) > 0 {
		i -= len(m.Nickname)
		copy(dAtA[i:], m.Nickname)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.Nickname)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DTag) > 0 {
		i -= len(m.DTag)
		copy(dAtA[i:], m.DTag)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.DTag)))
		i--
		dAtA[i] = 0x12
	}
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintModelsProfile(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pictures) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pictures) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pictures) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cover) > 0 {
		i -= len(m.Cover)
		copy(dAtA[i:], m.Cover)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.Cover)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Profile) > 0 {
		i -= len(m.Profile)
		copy(dAtA[i:], m.Profile)
		i = encodeVarintModelsProfile(dAtA, i, uint64(len(m.Profile)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintModelsProfile(dAtA []byte, offset int, v uint64) int {
	offset -= sovModelsProfile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = len(m.DTag)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = len(m.Bio)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = m.Pictures.Size()
	n += 1 + l + sovModelsProfile(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationDate)
	n += 1 + l + sovModelsProfile(uint64(l))
	return n
}

func (m *Pictures) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Profile)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	l = len(m.Cover)
	if l > 0 {
		n += 1 + l + sovModelsProfile(uint64(l))
	}
	return n
}

func sovModelsProfile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModelsProfile(x uint64) (n int) {
	return sovModelsProfile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Profile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsProfile
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
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &types.Any{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DTag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DTag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pictures", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Pictures.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreationDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsProfile
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
func (m *Pictures) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModelsProfile
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
			return fmt.Errorf("proto: Pictures: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pictures: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Profile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Profile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cover", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModelsProfile
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
				return ErrInvalidLengthModelsProfile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModelsProfile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cover = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModelsProfile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModelsProfile
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
func skipModelsProfile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModelsProfile
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
					return 0, ErrIntOverflowModelsProfile
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
					return 0, ErrIntOverflowModelsProfile
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
				return 0, ErrInvalidLengthModelsProfile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModelsProfile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModelsProfile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModelsProfile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModelsProfile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModelsProfile = fmt.Errorf("proto: unexpected end of group")
)
