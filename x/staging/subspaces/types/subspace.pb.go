// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/subspaces/v1beta1/subspace.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
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

// SubspaceType contains all the possible subspace types
type SubspaceType int32

const (
	// SubspaceTypeUnspecified identifies an unspecified type of subspace (used in
	// errors)
	SubspaceTypeUnspecified SubspaceType = 0
	// SubspaceTypeOpen identifies that users can interact inside the subspace
	// without the need to being registered in it
	SubspaceTypeOpen SubspaceType = 1
	// SubspaceTypeClosed identifies that users can't interact inside the subspace
	// without being registered in it
	SubspaceTypeClosed SubspaceType = 2
)

var SubspaceType_name = map[int32]string{
	0: "SUBSPACE_TYPE_UNSPECIFIED",
	1: "SUBSPACE_TYPE_OPEN",
	2: "SUBSPACE_TYPE_CLOSED",
}

var SubspaceType_value = map[string]int32{
	"SUBSPACE_TYPE_UNSPECIFIED": 0,
	"SUBSPACE_TYPE_OPEN":        1,
	"SUBSPACE_TYPE_CLOSED":      2,
}

func (x SubspaceType) String() string {
	return proto.EnumName(SubspaceType_name, int32(x))
}

func (SubspaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e657cf67bd23372d, []int{0}
}

// Subspace contains all the data of a Desmos subspace
type Subspace struct {
	// unique SHA-256 string that identifies the subspace
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"subspace_id" yaml:"subspace_id"`
	// human readable name of the subspace
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	// the address of the user that owns the subspace
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// the address of the subspace creator
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	// the creation time of the subspace
	CreationTime time.Time `protobuf:"bytes,5,opt,name=creation_time,json=creationTime,proto3,stdtime" json:"creation_time" yaml:"creation_time"`
	// the type of the subspace that indicates if it need registration or not
	Type SubspaceType `protobuf:"varint,6,opt,name=type,proto3,enum=desmos.subspaces.v1beta1.SubspaceType" json:"type" yaml:"type"`
}

func (m *Subspace) Reset()         { *m = Subspace{} }
func (m *Subspace) String() string { return proto.CompactTextString(m) }
func (*Subspace) ProtoMessage()    {}
func (*Subspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_e657cf67bd23372d, []int{0}
}
func (m *Subspace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subspace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subspace.Merge(m, src)
}
func (m *Subspace) XXX_Size() int {
	return m.Size()
}
func (m *Subspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Subspace.DiscardUnknown(m)
}

var xxx_messageInfo_Subspace proto.InternalMessageInfo

func (m *Subspace) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Subspace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subspace) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Subspace) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Subspace) GetCreationTime() time.Time {
	if m != nil {
		return m.CreationTime
	}
	return time.Time{}
}

func (m *Subspace) GetType() SubspaceType {
	if m != nil {
		return m.Type
	}
	return SubspaceTypeUnspecified
}

func init() {
	proto.RegisterEnum("desmos.subspaces.v1beta1.SubspaceType", SubspaceType_name, SubspaceType_value)
	proto.RegisterType((*Subspace)(nil), "desmos.subspaces.v1beta1.Subspace")
}

func init() {
	proto.RegisterFile("desmos/subspaces/v1beta1/subspace.proto", fileDescriptor_e657cf67bd23372d)
}

var fileDescriptor_e657cf67bd23372d = []byte{
	// 531 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x41, 0x8f, 0xd2, 0x4e,
	0x18, 0xc6, 0x3b, 0x2c, 0xbb, 0xff, 0xfd, 0xcf, 0xe2, 0x4a, 0x26, 0xc4, 0xad, 0x35, 0xe9, 0x34,
	0x35, 0x59, 0x89, 0xae, 0xad, 0xe0, 0x0d, 0x4f, 0x02, 0x35, 0x21, 0x31, 0x0b, 0x69, 0xe1, 0xa0,
	0x17, 0xd2, 0xd2, 0xd9, 0xda, 0x84, 0x76, 0x1a, 0x5a, 0x56, 0xf9, 0x06, 0x86, 0xd3, 0x5e, 0x4c,
	0xbc, 0x90, 0x6c, 0xe2, 0x17, 0xf0, 0x63, 0xec, 0x71, 0x8f, 0x9e, 0xaa, 0x81, 0x8b, 0xe1, 0xc8,
	0x27, 0x30, 0x9d, 0xd2, 0xb5, 0x1c, 0xbc, 0xbd, 0xf3, 0xbc, 0xbf, 0x67, 0x26, 0xef, 0xf3, 0x66,
	0xe0, 0x13, 0x9b, 0x84, 0x1e, 0x0d, 0xd5, 0x70, 0x6a, 0x85, 0x81, 0x39, 0x22, 0xa1, 0x7a, 0x59,
	0xb3, 0x48, 0x64, 0xd6, 0xee, 0x14, 0x25, 0x98, 0xd0, 0x88, 0x22, 0x3e, 0x05, 0x95, 0x3b, 0x50,
	0xd9, 0x82, 0x42, 0xc5, 0xa1, 0x0e, 0x65, 0x90, 0x9a, 0x54, 0x29, 0x2f, 0x60, 0x87, 0x52, 0x67,
	0x4c, 0x54, 0x76, 0xb2, 0xa6, 0x17, 0x6a, 0xe4, 0x7a, 0x24, 0x8c, 0x4c, 0x2f, 0x48, 0x01, 0xf9,
	0xcb, 0x1e, 0x3c, 0x34, 0xb6, 0x97, 0xa1, 0x57, 0xb0, 0xe0, 0xda, 0x3c, 0x90, 0x40, 0xf5, 0xff,
	0xe6, 0xb3, 0x65, 0x8c, 0x0b, 0x9d, 0xf6, 0x3a, 0xc6, 0x47, 0xd9, 0x63, 0x43, 0xd7, 0xde, 0xc4,
	0x18, 0xcd, 0x4c, 0x6f, 0xdc, 0x90, 0x73, 0xa2, 0xac, 0x17, 0x5c, 0x1b, 0x3d, 0x86, 0x45, 0xdf,
	0xf4, 0x08, 0x5f, 0x60, 0xf6, 0xfb, 0x9b, 0x18, 0x1f, 0xa5, 0x64, 0xa2, 0xca, 0x3a, 0x6b, 0xa2,
	0x53, 0xb8, 0x4f, 0x3f, 0xfa, 0x64, 0xc2, 0xef, 0x31, 0xaa, 0xbc, 0x89, 0x71, 0x29, 0xa5, 0x98,
	0x2c, 0xeb, 0x69, 0x1b, 0x9d, 0xc1, 0xff, 0x46, 0x13, 0x62, 0x46, 0x74, 0xc2, 0x17, 0x19, 0x89,
	0x36, 0x31, 0x3e, 0x4e, 0xc9, 0x6d, 0x43, 0xd6, 0x33, 0x04, 0x4d, 0xe0, 0x3d, 0x56, 0xba, 0xd4,
	0x1f, 0x26, 0x03, 0xf2, 0xfb, 0x12, 0xa8, 0x1e, 0xd5, 0x05, 0x25, 0x9d, 0x5e, 0xc9, 0xa6, 0x57,
	0xfa, 0xd9, 0xf4, 0xcd, 0xda, 0x4d, 0x8c, 0xb9, 0x75, 0x8c, 0x77, 0x8d, 0x9b, 0x18, 0x57, 0x72,
	0x8f, 0x64, 0xb2, 0x7c, 0xf5, 0x13, 0x03, 0xbd, 0x94, 0x69, 0xc9, 0x2d, 0xc8, 0x80, 0xc5, 0x68,
	0x16, 0x10, 0xfe, 0x40, 0x02, 0xd5, 0xe3, 0xfa, 0xa9, 0xf2, 0xaf, 0xc5, 0x28, 0x59, 0xba, 0xfd,
	0x59, 0x40, 0x9a, 0x27, 0xeb, 0x18, 0x33, 0xdf, 0xdf, 0x78, 0x92, 0x93, 0xac, 0x33, 0xb1, 0x71,
	0xf8, 0xf5, 0x1a, 0x83, 0xdf, 0xd7, 0x18, 0x3c, 0xfd, 0x0e, 0x60, 0x29, 0xef, 0x44, 0x0d, 0xf8,
	0xd0, 0x18, 0x34, 0x8d, 0xde, 0xeb, 0x96, 0x36, 0xec, 0xbf, 0xeb, 0x69, 0xc3, 0xc1, 0xb9, 0xd1,
	0xd3, 0x5a, 0x9d, 0x37, 0x1d, 0xad, 0x5d, 0xe6, 0x84, 0x47, 0xf3, 0x85, 0x74, 0x92, 0x37, 0x0c,
	0xfc, 0x30, 0x20, 0x23, 0xf7, 0xc2, 0x25, 0x36, 0x3a, 0x83, 0x68, 0xd7, 0xdb, 0xed, 0x69, 0xe7,
	0x65, 0x20, 0x54, 0xe6, 0x0b, 0xa9, 0x9c, 0x37, 0x75, 0x03, 0xe2, 0xa3, 0x17, 0xb0, 0xb2, 0x4b,
	0xb7, 0xde, 0x76, 0x0d, 0xad, 0x5d, 0x2e, 0x08, 0x0f, 0xe6, 0x0b, 0x09, 0xe5, 0xf9, 0xd6, 0x98,
	0x86, 0xc4, 0x16, 0x8a, 0x9f, 0xbf, 0x89, 0x5c, 0xb3, 0x7f, 0xb3, 0x14, 0xc1, 0xed, 0x52, 0x04,
	0xbf, 0x96, 0x22, 0xb8, 0x5a, 0x89, 0xdc, 0xed, 0x4a, 0xe4, 0x7e, 0xac, 0x44, 0xee, 0x7d, 0xc3,
	0x71, 0xa3, 0x0f, 0x53, 0x4b, 0x19, 0x51, 0x4f, 0x4d, 0x73, 0x7a, 0x3e, 0x36, 0xad, 0x70, 0x5b,
	0xab, 0x97, 0x75, 0xf5, 0x93, 0x1a, 0x46, 0xa6, 0xe3, 0xfa, 0x4e, 0xee, 0x0b, 0x24, 0x89, 0x84,
	0xd6, 0x01, 0x5b, 0xde, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x4f, 0x2f, 0x33, 0x23,
	0x03, 0x00, 0x00,
}

func (this *Subspace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Subspace)
	if !ok {
		that2, ok := that.(Subspace)
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
	if this.ID != that1.ID {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if !this.CreationTime.Equal(that1.CreationTime) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (m *Subspace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subspace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Subspace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintSubspace(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSubspace(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubspace(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubspace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Subspace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationTime)
	n += 1 + l + sovSubspace(uint64(l))
	if m.Type != 0 {
		n += 1 + sovSubspace(uint64(m.Type))
	}
	return n
}

func sovSubspace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubspace(x uint64) (n int) {
	return sovSubspace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Subspace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubspace
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
			return fmt.Errorf("proto: Subspace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subspace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SubspaceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSubspace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubspace
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
func skipSubspace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubspace
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
					return 0, ErrIntOverflowSubspace
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
					return 0, ErrIntOverflowSubspace
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
				return 0, ErrInvalidLengthSubspace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubspace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubspace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubspace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubspace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubspace = fmt.Errorf("proto: unexpected end of group")
)
