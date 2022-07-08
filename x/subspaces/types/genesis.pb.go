// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/subspaces/v2/genesis.proto

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

// GenesisState contains the data of the genesis state for the subspaces module
type GenesisState struct {
	InitialSubspaceID uint64                 `protobuf:"varint,1,opt,name=initial_subspace_id,json=initialSubspaceId,proto3" json:"initial_subspace_id,omitempty"`
	SubspacesData     []SubspaceData         `protobuf:"bytes,2,rep,name=subspaces_data,json=subspacesData,proto3" json:"subspaces_data"`
	Subspaces         []Subspace             `protobuf:"bytes,3,rep,name=subspaces,proto3" json:"subspaces"`
	Sections          []Section              `protobuf:"bytes,4,rep,name=sections,proto3" json:"sections"`
	UserPermissions   []UserPermission       `protobuf:"bytes,5,rep,name=user_permissions,json=userPermissions,proto3" json:"user_permissions"`
	UserGroups        []UserGroup            `protobuf:"bytes,6,rep,name=user_groups,json=userGroups,proto3" json:"user_groups"`
	UserGroupsMembers []UserGroupMemberEntry `protobuf:"bytes,7,rep,name=user_groups_members,json=userGroupsMembers,proto3" json:"user_groups_members"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e29defb77aaf744c, []int{0}
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

func (m *GenesisState) GetInitialSubspaceID() uint64 {
	if m != nil {
		return m.InitialSubspaceID
	}
	return 0
}

func (m *GenesisState) GetSubspacesData() []SubspaceData {
	if m != nil {
		return m.SubspacesData
	}
	return nil
}

func (m *GenesisState) GetSubspaces() []Subspace {
	if m != nil {
		return m.Subspaces
	}
	return nil
}

func (m *GenesisState) GetSections() []Section {
	if m != nil {
		return m.Sections
	}
	return nil
}

func (m *GenesisState) GetUserPermissions() []UserPermission {
	if m != nil {
		return m.UserPermissions
	}
	return nil
}

func (m *GenesisState) GetUserGroups() []UserGroup {
	if m != nil {
		return m.UserGroups
	}
	return nil
}

func (m *GenesisState) GetUserGroupsMembers() []UserGroupMemberEntry {
	if m != nil {
		return m.UserGroupsMembers
	}
	return nil
}

// SubspaceData contains the genesis data for a single subspace
type SubspaceData struct {
	SubspaceID    uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	NextGroupID   uint32 `protobuf:"varint,2,opt,name=next_group_id,json=nextGroupId,proto3" json:"next_group_id,omitempty"`
	NextSectionID uint32 `protobuf:"varint,3,opt,name=next_section_id,json=nextSectionId,proto3" json:"next_section_id,omitempty"`
}

func (m *SubspaceData) Reset()         { *m = SubspaceData{} }
func (m *SubspaceData) String() string { return proto.CompactTextString(m) }
func (*SubspaceData) ProtoMessage()    {}
func (*SubspaceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e29defb77aaf744c, []int{1}
}
func (m *SubspaceData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubspaceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubspaceData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubspaceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubspaceData.Merge(m, src)
}
func (m *SubspaceData) XXX_Size() int {
	return m.Size()
}
func (m *SubspaceData) XXX_DiscardUnknown() {
	xxx_messageInfo_SubspaceData.DiscardUnknown(m)
}

var xxx_messageInfo_SubspaceData proto.InternalMessageInfo

func (m *SubspaceData) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *SubspaceData) GetNextGroupID() uint32 {
	if m != nil {
		return m.NextGroupID
	}
	return 0
}

func (m *SubspaceData) GetNextSectionID() uint32 {
	if m != nil {
		return m.NextSectionID
	}
	return 0
}

// UserGroupMemberEntry contains the details of a user group member
type UserGroupMemberEntry struct {
	SubspaceID uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	GroupID    uint32 `protobuf:"varint,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	User       string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (m *UserGroupMemberEntry) Reset()         { *m = UserGroupMemberEntry{} }
func (m *UserGroupMemberEntry) String() string { return proto.CompactTextString(m) }
func (*UserGroupMemberEntry) ProtoMessage()    {}
func (*UserGroupMemberEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e29defb77aaf744c, []int{2}
}
func (m *UserGroupMemberEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserGroupMemberEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserGroupMemberEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserGroupMemberEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserGroupMemberEntry.Merge(m, src)
}
func (m *UserGroupMemberEntry) XXX_Size() int {
	return m.Size()
}
func (m *UserGroupMemberEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserGroupMemberEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserGroupMemberEntry proto.InternalMessageInfo

func (m *UserGroupMemberEntry) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *UserGroupMemberEntry) GetGroupID() uint32 {
	if m != nil {
		return m.GroupID
	}
	return 0
}

func (m *UserGroupMemberEntry) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.subspaces.v2.GenesisState")
	proto.RegisterType((*SubspaceData)(nil), "desmos.subspaces.v2.SubspaceData")
	proto.RegisterType((*UserGroupMemberEntry)(nil), "desmos.subspaces.v2.UserGroupMemberEntry")
}

func init() { proto.RegisterFile("desmos/subspaces/v2/genesis.proto", fileDescriptor_e29defb77aaf744c) }

var fileDescriptor_e29defb77aaf744c = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x73, 0x4d, 0x68, 0xd2, 0x73, 0xd3, 0x90, 0x4b, 0x91, 0xac, 0x0a, 0xec, 0xb4, 0x48,
	0x28, 0x0c, 0xd8, 0x52, 0xca, 0x02, 0x03, 0x12, 0x51, 0xa2, 0x2a, 0x12, 0x54, 0x28, 0x85, 0x85,
	0xc5, 0x72, 0xe2, 0x93, 0xb1, 0x14, 0xff, 0x91, 0xdf, 0x73, 0x94, 0x7e, 0x08, 0x24, 0x46, 0xc6,
	0x7e, 0x04, 0x56, 0xbe, 0x41, 0xc7, 0x8e, 0x4c, 0x16, 0x72, 0x16, 0x3e, 0x06, 0xf2, 0xf9, 0xec,
	0x98, 0x62, 0x2a, 0xb1, 0x9d, 0xdf, 0xf7, 0x79, 0x7e, 0xf7, 0x9c, 0xdf, 0x3b, 0x7c, 0x6c, 0x51,
	0x70, 0x7d, 0xd0, 0x21, 0x9a, 0x43, 0x60, 0x2e, 0x28, 0xe8, 0xab, 0xa1, 0x6e, 0x53, 0x8f, 0x82,
	0x03, 0x5a, 0x10, 0xfa, 0xcc, 0x27, 0xbd, 0x4c, 0xa2, 0x15, 0x12, 0x6d, 0x35, 0x3c, 0x3a, 0xb4,
	0x7d, 0xdb, 0xe7, 0x7d, 0x3d, 0x5d, 0x65, 0xd2, 0xa3, 0x7e, 0x15, 0xcd, 0xf5, 0x2d, 0xba, 0x14,
	0xb0, 0x93, 0x6f, 0x0d, 0xbc, 0x7f, 0x96, 0xe1, 0x2f, 0x98, 0xc9, 0x28, 0x99, 0xe0, 0x9e, 0xe3,
	0x39, 0xcc, 0x31, 0x97, 0x46, 0xee, 0x32, 0x1c, 0x4b, 0x46, 0x7d, 0x34, 0x68, 0x8c, 0x1e, 0x24,
	0xb1, 0xda, 0x9d, 0x66, 0xed, 0x0b, 0xd1, 0x9d, 0x8e, 0x67, 0x5d, 0xe7, 0x56, 0xc9, 0x22, 0xe7,
	0xf8, 0xa0, 0xd8, 0xd4, 0xb0, 0x4c, 0x66, 0xca, 0x3b, 0xfd, 0xfa, 0x40, 0x1a, 0x1e, 0x6b, 0x15,
	0xe9, 0xb5, 0xdc, 0x38, 0x36, 0x99, 0x39, 0x6a, 0x5c, 0xc7, 0x6a, 0x6d, 0xd6, 0x2e, 0x04, 0x69,
	0x91, 0xbc, 0xc6, 0x7b, 0x45, 0x41, 0xae, 0x73, 0xd4, 0xa3, 0x3b, 0x51, 0x02, 0xb3, 0x75, 0x91,
	0x57, 0xb8, 0x05, 0x74, 0xc1, 0x1c, 0xdf, 0x03, 0xb9, 0xc1, 0x09, 0x0f, 0xab, 0x09, 0x99, 0x48,
	0x00, 0x0a, 0x0f, 0x79, 0x8f, 0xef, 0x47, 0x40, 0x43, 0x23, 0xa0, 0xa1, 0xeb, 0x00, 0x70, 0xce,
	0x3d, 0xce, 0x79, 0x5c, 0xc9, 0xf9, 0x00, 0x34, 0x7c, 0x57, 0x68, 0x05, 0xae, 0x13, 0xfd, 0x51,
	0x05, 0x32, 0xc1, 0x12, 0xa7, 0xda, 0xa1, 0x1f, 0x05, 0x20, 0xef, 0x72, 0xa0, 0xf2, 0x4f, 0xe0,
	0x59, 0x2a, 0x13, 0x2c, 0x1c, 0xe5, 0x05, 0x20, 0x06, 0xee, 0x95, 0x30, 0x86, 0x4b, 0xdd, 0x39,
	0x0d, 0x41, 0x6e, 0x72, 0xdc, 0xd3, 0xbb, 0x71, 0x6f, 0xb9, 0x78, 0xe2, 0xb1, 0xf0, 0x52, 0x90,
	0xbb, 0x5b, 0x72, 0xd6, 0x84, 0x97, 0xad, 0xaf, 0x57, 0x2a, 0xfa, 0x75, 0xa5, 0xa2, 0x93, 0xef,
	0x08, 0xef, 0x97, 0x07, 0x46, 0x74, 0x2c, 0xfd, 0x7d, 0x55, 0x0e, 0x92, 0x58, 0xc5, 0xa5, 0x3b,
	0x82, 0x61, 0x7b, 0x39, 0x4e, 0x71, 0xdb, 0xa3, 0x6b, 0x96, 0x85, 0x4d, 0x2d, 0x3b, 0x7d, 0x34,
	0x68, 0x8f, 0x3a, 0x49, 0xac, 0x4a, 0xe7, 0x74, 0xcd, 0xf8, 0xce, 0xd3, 0xf1, 0x4c, 0xf2, 0x8a,
	0x0f, 0x8b, 0xbc, 0xc0, 0x1d, 0x6e, 0x12, 0xf3, 0x48, 0x6d, 0x75, 0x6e, 0xeb, 0x26, 0xb1, 0xda,
	0x4e, 0x6d, 0x62, 0x70, 0xd3, 0xf1, 0x8c, 0xe3, 0xf3, 0x4f, 0xab, 0x94, 0xfd, 0x33, 0xc2, 0x87,
	0x55, 0xe7, 0xfe, 0xff, 0x33, 0x3c, 0xc1, 0xad, 0x5b, 0xf1, 0xa5, 0x24, 0x56, 0x9b, 0x79, 0xf4,
	0xa6, 0x2d, 0x62, 0x13, 0xdc, 0x48, 0x7f, 0x26, 0xcf, 0xba, 0x37, 0xe3, 0xeb, 0x6d, 0x9e, 0xd1,
	0x9b, 0xeb, 0x44, 0x41, 0x37, 0x89, 0x82, 0x7e, 0x26, 0x0a, 0xfa, 0xb2, 0x51, 0x6a, 0x37, 0x1b,
	0xa5, 0xf6, 0x63, 0xa3, 0xd4, 0x3e, 0x0e, 0x6d, 0x87, 0x7d, 0x8a, 0xe6, 0xda, 0xc2, 0x77, 0xf5,
	0x6c, 0x7a, 0xcf, 0x96, 0xe6, 0x1c, 0xc4, 0x5a, 0x5f, 0x3d, 0xd7, 0xd7, 0xa5, 0x67, 0xcd, 0x2e,
	0x03, 0x0a, 0xf3, 0x5d, 0xfe, 0xa6, 0x4f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x21, 0xa7, 0xab,
	0x7a, 0x45, 0x04, 0x00, 0x00,
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
	if this.InitialSubspaceID != that1.InitialSubspaceID {
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
	if len(this.Subspaces) != len(that1.Subspaces) {
		return false
	}
	for i := range this.Subspaces {
		if !this.Subspaces[i].Equal(&that1.Subspaces[i]) {
			return false
		}
	}
	if len(this.Sections) != len(that1.Sections) {
		return false
	}
	for i := range this.Sections {
		if !this.Sections[i].Equal(&that1.Sections[i]) {
			return false
		}
	}
	if len(this.UserPermissions) != len(that1.UserPermissions) {
		return false
	}
	for i := range this.UserPermissions {
		if !this.UserPermissions[i].Equal(&that1.UserPermissions[i]) {
			return false
		}
	}
	if len(this.UserGroups) != len(that1.UserGroups) {
		return false
	}
	for i := range this.UserGroups {
		if !this.UserGroups[i].Equal(&that1.UserGroups[i]) {
			return false
		}
	}
	if len(this.UserGroupsMembers) != len(that1.UserGroupsMembers) {
		return false
	}
	for i := range this.UserGroupsMembers {
		if !this.UserGroupsMembers[i].Equal(&that1.UserGroupsMembers[i]) {
			return false
		}
	}
	return true
}
func (this *SubspaceData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SubspaceData)
	if !ok {
		that2, ok := that.(SubspaceData)
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
	if this.NextGroupID != that1.NextGroupID {
		return false
	}
	if this.NextSectionID != that1.NextSectionID {
		return false
	}
	return true
}
func (this *UserGroupMemberEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UserGroupMemberEntry)
	if !ok {
		that2, ok := that.(UserGroupMemberEntry)
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
	if this.GroupID != that1.GroupID {
		return false
	}
	if this.User != that1.User {
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
	if len(m.UserGroupsMembers) > 0 {
		for iNdEx := len(m.UserGroupsMembers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserGroupsMembers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.UserGroups) > 0 {
		for iNdEx := len(m.UserGroups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserGroups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.UserPermissions) > 0 {
		for iNdEx := len(m.UserPermissions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserPermissions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Sections) > 0 {
		for iNdEx := len(m.Sections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Sections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Subspaces) > 0 {
		for iNdEx := len(m.Subspaces) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Subspaces[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
			dAtA[i] = 0x12
		}
	}
	if m.InitialSubspaceID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.InitialSubspaceID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SubspaceData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubspaceData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubspaceData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextSectionID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextSectionID))
		i--
		dAtA[i] = 0x18
	}
	if m.NextGroupID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextGroupID))
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

func (m *UserGroupMemberEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserGroupMemberEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserGroupMemberEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x1a
	}
	if m.GroupID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GroupID))
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
	if m.InitialSubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.InitialSubspaceID))
	}
	if len(m.SubspacesData) > 0 {
		for _, e := range m.SubspacesData {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Subspaces) > 0 {
		for _, e := range m.Subspaces {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Sections) > 0 {
		for _, e := range m.Sections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserPermissions) > 0 {
		for _, e := range m.UserPermissions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserGroups) > 0 {
		for _, e := range m.UserGroups {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserGroupsMembers) > 0 {
		for _, e := range m.UserGroupsMembers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *SubspaceData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.NextGroupID != 0 {
		n += 1 + sovGenesis(uint64(m.NextGroupID))
	}
	if m.NextSectionID != 0 {
		n += 1 + sovGenesis(uint64(m.NextSectionID))
	}
	return n
}

func (m *UserGroupMemberEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.GroupID != 0 {
		n += 1 + sovGenesis(uint64(m.GroupID))
	}
	l = len(m.User)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialSubspaceID", wireType)
			}
			m.InitialSubspaceID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialSubspaceID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
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
			m.SubspacesData = append(m.SubspacesData, SubspaceData{})
			if err := m.SubspacesData[len(m.SubspacesData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspaces", wireType)
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
			m.Subspaces = append(m.Subspaces, Subspace{})
			if err := m.Subspaces[len(m.Subspaces)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sections", wireType)
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
			m.Sections = append(m.Sections, Section{})
			if err := m.Sections[len(m.Sections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserPermissions", wireType)
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
			m.UserPermissions = append(m.UserPermissions, UserPermission{})
			if err := m.UserPermissions[len(m.UserPermissions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserGroups", wireType)
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
			m.UserGroups = append(m.UserGroups, UserGroup{})
			if err := m.UserGroups[len(m.UserGroups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserGroupsMembers", wireType)
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
			m.UserGroupsMembers = append(m.UserGroupsMembers, UserGroupMemberEntry{})
			if err := m.UserGroupsMembers[len(m.UserGroupsMembers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *SubspaceData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SubspaceData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubspaceData: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field NextGroupID", wireType)
			}
			m.NextGroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextGroupID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextSectionID", wireType)
			}
			m.NextSectionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextSectionID |= uint32(b&0x7F) << shift
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
func (m *UserGroupMemberEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UserGroupMemberEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserGroupMemberEntry: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field GroupID", wireType)
			}
			m.GroupID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
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
			m.User = string(dAtA[iNdEx:postIndex])
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
