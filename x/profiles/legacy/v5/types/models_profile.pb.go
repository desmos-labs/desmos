// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/profiles/v2/models_profile.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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
	return fileDescriptor_089dd63594c4b06b, []int{0}
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
	return fileDescriptor_089dd63594c4b06b, []int{1}
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
	proto.RegisterType((*Profile)(nil), "desmos.profiles.v2.Profile")
	proto.RegisterType((*Pictures)(nil), "desmos.profiles.v2.Pictures")
}

func init() {
	proto.RegisterFile("desmos/profiles/v2/models_profile.proto", fileDescriptor_089dd63594c4b06b)
}

var fileDescriptor_089dd63594c4b06b = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x13, 0xd6, 0xad, 0x9d, 0x37, 0x06, 0x32, 0x95, 0x16, 0xaa, 0x29, 0xae, 0x7c, 0x80,
	0x49, 0x30, 0x5b, 0x2b, 0x42, 0x48, 0xe5, 0xb4, 0x68, 0x17, 0x6e, 0x25, 0x1a, 0x17, 0x2e, 0x95,
	0x93, 0x7a, 0x59, 0x20, 0x89, 0xab, 0xc6, 0x8d, 0xe8, 0x1b, 0x70, 0xdc, 0x71, 0x12, 0x97, 0x3e,
	0x04, 0x0f, 0x31, 0x71, 0xda, 0x91, 0x53, 0x40, 0xed, 0x85, 0x73, 0x9e, 0x00, 0x25, 0xb6, 0x37,
	0xc1, 0x6e, 0xf6, 0xf7, 0xff, 0x7d, 0xff, 0xef, 0xcb, 0xdf, 0x01, 0xcf, 0x27, 0x3c, 0x4f, 0x45,
	0x4e, 0xa7, 0x33, 0x71, 0x1e, 0x27, 0x3c, 0xa7, 0xc5, 0x80, 0xa6, 0x62, 0xc2, 0x93, 0x7c, 0xac,
	0x4b, 0x64, 0x3a, 0x13, 0x52, 0x40, 0xa8, 0x40, 0x62, 0x40, 0x52, 0x0c, 0x7a, 0xdd, 0x48, 0x44,
	0xa2, 0x91, 0x69, 0x7d, 0x52, 0x64, 0xef, 0x69, 0x24, 0x44, 0x94, 0x70, 0xda, 0xdc, 0x82, 0xf9,
	0x39, 0x65, 0xd9, 0x42, 0x4b, 0xe8, 0x7f, 0x49, 0xc6, 0x29, 0xcf, 0x25, 0x4b, 0xa7, 0xa6, 0x37,
	0x14, 0xf5, 0x94, 0xb1, 0x32, 0x55, 0x17, 0x25, 0xe1, 0x6f, 0x1b, 0xa0, 0x3d, 0x52, 0xc3, 0xe1,
	0x08, 0xb4, 0x59, 0x18, 0x8a, 0x79, 0x26, 0x1d, 0xbb, 0x6f, 0x1f, 0xee, 0x0c, 0xba, 0x44, 0x39,
	0x13, 0xe3, 0x4c, 0x4e, 0xb2, 0x85, 0xd7, 0xff, 0xf1, 0xfd, 0xe8, 0x40, 0x9b, 0xb0, 0xb9, 0xbc,
	0x20, 0xc5, 0x71, 0xc0, 0x25, 0x3b, 0x26, 0x27, 0xaa, 0xf9, 0x9d, 0x6f, 0x6c, 0xe0, 0x0b, 0xd0,
	0x9a, 0x48, 0x16, 0x39, 0x0f, 0xfa, 0xf6, 0xe1, 0xb6, 0xb7, 0xbf, 0x2a, 0x51, 0xeb, 0xf4, 0x8c,
	0x45, 0x55, 0x89, 0x76, 0x16, 0x2c, 0x4d, 0x86, 0xb8, 0x56, 0xb1, 0xdf, 0x40, 0x90, 0x82, 0x4e,
	0x16, 0x87, 0x9f, 0x33, 0x96, 0x72, 0x67, 0xa3, 0x69, 0x78, 0x52, 0x95, 0xe8, 0x91, 0x02, 0x8d,
	0x82, 0xfd, 0x5b, 0x08, 0xf6, 0xc1, 0x46, 0x10, 0x0b, 0xa7, 0xd5, 0xb0, 0x7b, 0x55, 0x89, 0x80,
	0x62, 0x83, 0x58, 0x60, 0xbf, 0x96, 0xe0, 0x7b, 0xd0, 0x99, 0xc6, 0xa1, 0x9c, 0xcf, 0x78, 0xee,
	0x6c, 0x36, 0x9f, 0x74, 0x40, 0xee, 0x27, 0x4e, 0x46, 0x9a, 0xf1, 0xf6, 0xaf, 0x4b, 0x64, 0xdd,
	0x0d, 0x35, 0xbd, 0xd8, 0xbf, 0xb5, 0x81, 0x0c, 0x3c, 0x0c, 0x67, 0x9c, 0xc9, 0x58, 0x64, 0xe3,
	0x09, 0x93, 0xdc, 0xd9, 0x6a, 0x7c, 0x7b, 0xf7, 0xa2, 0x3a, 0x33, 0x8f, 0xe0, 0xf5, 0xb5, 0x6b,
	0x57, 0xb9, 0xfe, 0xd3, 0x8e, 0x2f, 0x7f, 0x21, 0xdb, 0xdf, 0x35, 0xb5, 0x53, 0x26, 0xf9, 0xb0,
	0xf3, 0x75, 0x89, 0xac, 0xab, 0x25, 0xb2, 0xf0, 0x27, 0xd0, 0x31, 0xbb, 0xc1, 0x97, 0xa0, 0xad,
	0x77, 0x6e, 0x5e, 0x67, 0xdb, 0x83, 0x55, 0x89, 0xf6, 0xf4, 0xa2, 0x4a, 0xc0, 0xbe, 0x41, 0xe0,
	0x33, 0xb0, 0x19, 0x8a, 0x82, 0xcf, 0x74, 0xf4, 0x8f, 0xab, 0x12, 0xed, 0xea, 0xf1, 0x75, 0x19,
	0xfb, 0x4a, 0x1e, 0x76, 0xae, 0x96, 0xc8, 0xfe, 0xb3, 0x44, 0xb6, 0xf7, 0xe1, 0x7a, 0xe5, 0xda,
	0x37, 0x2b, 0xd7, 0xfe, 0xbd, 0x72, 0xed, 0xcb, 0xb5, 0x6b, 0xdd, 0xac, 0x5d, 0xeb, 0xe7, 0xda,
	0xb5, 0x3e, 0xbe, 0x8d, 0x62, 0x79, 0x31, 0x0f, 0x48, 0x28, 0x52, 0xaa, 0xd2, 0x3b, 0x4a, 0x58,
	0x90, 0xeb, 0x33, 0x2d, 0xde, 0xd0, 0x2f, 0x77, 0x7f, 0x7a, 0xc2, 0x23, 0x16, 0x2e, 0x68, 0xf1,
	0x9a, 0xca, 0xc5, 0x94, 0xe7, 0xc1, 0x56, 0x13, 0xc8, 0xab, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xcf, 0xff, 0x46, 0xf0, 0x13, 0x03, 0x00, 0x00,
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
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CreationDate, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreationDate):])
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
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreationDate)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CreationDate, dAtA[iNdEx:postIndex]); err != nil {
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
