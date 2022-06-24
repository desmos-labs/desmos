// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/posts/v1/genesis.proto

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

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	SubspacesData []SubspaceDataEntry `protobuf:"bytes,1,rep,name=subspaces_data,json=subspacesData,proto3" json:"subspaces_data"`
	PostsData     []PostDataEntry     `protobuf:"bytes,2,rep,name=posts_data,json=postsData,proto3" json:"posts_data"`
	Posts         []Post              `protobuf:"bytes,3,rep,name=posts,proto3" json:"posts"`
	Attachments   []Attachment        `protobuf:"bytes,4,rep,name=attachments,proto3" json:"attachments"`
	ActivePolls   []ActivePollData    `protobuf:"bytes,5,rep,name=active_polls,json=activePolls,proto3" json:"active_polls"`
	UserAnswers   []UserAnswer        `protobuf:"bytes,6,rep,name=user_answers,json=userAnswers,proto3" json:"user_answers"`
	Params        Params              `protobuf:"bytes,7,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b7acf2775f2913, []int{0}
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

func (m *GenesisState) GetPostsData() []PostDataEntry {
	if m != nil {
		return m.PostsData
	}
	return nil
}

func (m *GenesisState) GetPosts() []Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *GenesisState) GetAttachments() []Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *GenesisState) GetActivePolls() []ActivePollData {
	if m != nil {
		return m.ActivePolls
	}
	return nil
}

func (m *GenesisState) GetUserAnswers() []UserAnswer {
	if m != nil {
		return m.UserAnswers
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// SubspaceDataEntry contains the data for a given subspace
type SubspaceDataEntry struct {
	SubspaceID    uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	InitialPostID uint64 `protobuf:"varint,2,opt,name=initial_post_id,json=initialPostId,proto3" json:"initial_post_id,omitempty"`
}

func (m *SubspaceDataEntry) Reset()         { *m = SubspaceDataEntry{} }
func (m *SubspaceDataEntry) String() string { return proto.CompactTextString(m) }
func (*SubspaceDataEntry) ProtoMessage()    {}
func (*SubspaceDataEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b7acf2775f2913, []int{1}
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

func (m *SubspaceDataEntry) GetInitialPostID() uint64 {
	if m != nil {
		return m.InitialPostID
	}
	return 0
}

// PostDataEntry contains the data of a given post
type PostDataEntry struct {
	SubspaceID          uint64 `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	PostID              uint64 `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	InitialAttachmentID uint32 `protobuf:"varint,3,opt,name=initial_attachment_id,json=initialAttachmentId,proto3" json:"initial_attachment_id,omitempty"`
}

func (m *PostDataEntry) Reset()         { *m = PostDataEntry{} }
func (m *PostDataEntry) String() string { return proto.CompactTextString(m) }
func (*PostDataEntry) ProtoMessage()    {}
func (*PostDataEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b7acf2775f2913, []int{2}
}
func (m *PostDataEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostDataEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostDataEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostDataEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostDataEntry.Merge(m, src)
}
func (m *PostDataEntry) XXX_Size() int {
	return m.Size()
}
func (m *PostDataEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PostDataEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PostDataEntry proto.InternalMessageInfo

func (m *PostDataEntry) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *PostDataEntry) GetPostID() uint64 {
	if m != nil {
		return m.PostID
	}
	return 0
}

func (m *PostDataEntry) GetInitialAttachmentID() uint32 {
	if m != nil {
		return m.InitialAttachmentID
	}
	return 0
}

// ActivePollData contains the data of an active poll
type ActivePollData struct {
	SubspaceID uint64    `protobuf:"varint,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty"`
	PostID     uint64    `protobuf:"varint,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	PollID     uint32    `protobuf:"varint,3,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
	EndDate    time.Time `protobuf:"bytes,4,opt,name=end_date,json=endDate,proto3,stdtime" json:"end_date"`
}

func (m *ActivePollData) Reset()         { *m = ActivePollData{} }
func (m *ActivePollData) String() string { return proto.CompactTextString(m) }
func (*ActivePollData) ProtoMessage()    {}
func (*ActivePollData) Descriptor() ([]byte, []int) {
	return fileDescriptor_26b7acf2775f2913, []int{3}
}
func (m *ActivePollData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivePollData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivePollData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivePollData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivePollData.Merge(m, src)
}
func (m *ActivePollData) XXX_Size() int {
	return m.Size()
}
func (m *ActivePollData) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivePollData.DiscardUnknown(m)
}

var xxx_messageInfo_ActivePollData proto.InternalMessageInfo

func (m *ActivePollData) GetSubspaceID() uint64 {
	if m != nil {
		return m.SubspaceID
	}
	return 0
}

func (m *ActivePollData) GetPostID() uint64 {
	if m != nil {
		return m.PostID
	}
	return 0
}

func (m *ActivePollData) GetPollID() uint32 {
	if m != nil {
		return m.PollID
	}
	return 0
}

func (m *ActivePollData) GetEndDate() time.Time {
	if m != nil {
		return m.EndDate
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.posts.v1.GenesisState")
	proto.RegisterType((*SubspaceDataEntry)(nil), "desmos.posts.v1.SubspaceDataEntry")
	proto.RegisterType((*PostDataEntry)(nil), "desmos.posts.v1.PostDataEntry")
	proto.RegisterType((*ActivePollData)(nil), "desmos.posts.v1.ActivePollData")
}

func init() { proto.RegisterFile("desmos/posts/v1/genesis.proto", fileDescriptor_26b7acf2775f2913) }

var fileDescriptor_26b7acf2775f2913 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xbf, 0x6e, 0xd3, 0x40,
	0x18, 0xcf, 0x35, 0x69, 0x12, 0x2e, 0x7f, 0xaa, 0xba, 0x54, 0xb5, 0x02, 0xd8, 0x51, 0x58, 0xb2,
	0x60, 0x2b, 0x05, 0x06, 0x58, 0x50, 0x53, 0x23, 0x6a, 0x31, 0x50, 0xa5, 0xb0, 0xb0, 0x44, 0x97,
	0xf8, 0x70, 0x2d, 0x9d, 0x7d, 0x96, 0xef, 0x12, 0xe8, 0x23, 0xb0, 0x75, 0x64, 0xec, 0x83, 0xf0,
	0x00, 0x1d, 0x3b, 0x21, 0xa6, 0x80, 0x9c, 0x85, 0x85, 0x77, 0x40, 0x3e, 0x9f, 0xf3, 0xcf, 0x30,
	0x30, 0x74, 0xbb, 0xbb, 0xef, 0xf7, 0x2f, 0xf9, 0x3e, 0x7f, 0xf0, 0x81, 0x83, 0x99, 0x4f, 0x99,
	0x19, 0x52, 0xc6, 0x99, 0x39, 0xed, 0x99, 0x2e, 0x0e, 0x30, 0xf3, 0x98, 0x11, 0x46, 0x94, 0x53,
	0x65, 0x27, 0x2d, 0x1b, 0xa2, 0x6c, 0x4c, 0x7b, 0xad, 0xbb, 0x2e, 0x75, 0xa9, 0xa8, 0x99, 0xc9,
	0x29, 0x85, 0xb5, 0x74, 0x97, 0x52, 0x97, 0x60, 0x53, 0xdc, 0x46, 0x93, 0x0f, 0x26, 0xf7, 0x7c,
	0xcc, 0x38, 0xf2, 0x43, 0x09, 0xb8, 0xbf, 0x69, 0xe3, 0x53, 0x07, 0x13, 0xe9, 0xd2, 0xf9, 0x5d,
	0x84, 0xf5, 0x57, 0xa9, 0xef, 0x19, 0x47, 0x1c, 0x2b, 0x6f, 0x60, 0x93, 0x4d, 0x46, 0x2c, 0x44,
	0x63, 0xcc, 0x86, 0x0e, 0xe2, 0x48, 0x05, 0xed, 0x62, 0xb7, 0x76, 0xd8, 0x31, 0x36, 0xf2, 0x18,
	0x67, 0x12, 0x66, 0x21, 0x8e, 0x5e, 0x06, 0x3c, 0xba, 0xe8, 0x97, 0xae, 0x67, 0x7a, 0x61, 0xd0,
	0x58, 0xf0, 0x93, 0x8a, 0x72, 0x0c, 0xa1, 0xa0, 0xa4, 0x62, 0x5b, 0x42, 0x4c, 0xcb, 0x89, 0x9d,
	0x52, 0xc6, 0x37, 0x85, 0xee, 0x88, 0xaa, 0x10, 0xe9, 0xc1, 0x6d, 0x71, 0x51, 0x8b, 0x82, 0xbf,
	0xff, 0x57, 0xbe, 0xa4, 0xa5, 0x48, 0xe5, 0x18, 0xd6, 0x10, 0xe7, 0x68, 0x7c, 0xee, 0xe3, 0x80,
	0x33, 0xb5, 0x24, 0x88, 0xf7, 0x72, 0xc4, 0xa3, 0x05, 0x46, 0xd2, 0x57, 0x59, 0xca, 0x09, 0xac,
	0xa3, 0x31, 0xf7, 0xa6, 0x78, 0x18, 0x52, 0x42, 0x98, 0xba, 0x2d, 0x54, 0xf4, 0xbc, 0x8a, 0x00,
	0x9d, 0x52, 0x42, 0x92, 0xb8, 0x0b, 0xa5, 0xc5, 0x2b, 0x53, 0x2c, 0x58, 0x9f, 0x30, 0x1c, 0x0d,
	0x51, 0xc0, 0x3e, 0xe2, 0x88, 0xa9, 0xe5, 0x7f, 0xe4, 0x79, 0xc7, 0x70, 0x74, 0x24, 0x30, 0x99,
	0xca, 0x64, 0xf1, 0xc2, 0x94, 0xa7, 0xb0, 0x1c, 0xa2, 0x08, 0xf9, 0x4c, 0xad, 0xb4, 0x41, 0xb7,
	0x76, 0x78, 0x90, 0xff, 0x23, 0x44, 0x59, 0x72, 0x25, 0xf8, 0x79, 0xf5, 0xcb, 0x95, 0x0e, 0x7e,
	0x5d, 0xe9, 0xa0, 0xf3, 0x19, 0xc0, 0xdd, 0x5c, 0xe3, 0x14, 0x13, 0xd6, 0xb2, 0xa6, 0x0d, 0x3d,
	0x47, 0x05, 0x6d, 0xd0, 0x2d, 0xf5, 0x9b, 0xf1, 0x4c, 0x87, 0x19, 0xd6, 0xb6, 0x06, 0x30, 0x83,
	0xd8, 0x8e, 0xf2, 0x0c, 0xee, 0x78, 0x81, 0xc7, 0x3d, 0x44, 0x86, 0x89, 0x73, 0x42, 0xda, 0x12,
	0xa4, 0xdd, 0x78, 0xa6, 0x37, 0xec, 0xb4, 0x94, 0xf4, 0xc4, 0xb6, 0x06, 0x0d, 0x6f, 0xe5, 0xea,
	0xac, 0x64, 0xf9, 0x0a, 0x60, 0x63, 0xad, 0xef, 0xff, 0x9f, 0xe3, 0x21, 0xac, 0xac, 0xfb, 0xc3,
	0x78, 0xa6, 0x97, 0xa5, 0x71, 0x39, 0x14, 0x8e, 0xca, 0x6b, 0xb8, 0x9f, 0x85, 0x5d, 0xf6, 0x36,
	0xa1, 0x14, 0xdb, 0xa0, 0xdb, 0xe8, 0x1f, 0xc4, 0x33, 0x7d, 0x4f, 0x46, 0x5e, 0x4e, 0x83, 0x6d,
	0x0d, 0xf6, 0xbc, 0xdc, 0xe3, 0x6a, 0xfc, 0x6f, 0x00, 0x36, 0xd7, 0xfb, 0x7e, 0x4b, 0xf9, 0x05,
	0x88, 0x90, 0x65, 0x62, 0x09, 0x22, 0x24, 0x05, 0x11, 0x62, 0x3b, 0xca, 0x0b, 0x58, 0xc5, 0x81,
	0x93, 0x7c, 0x64, 0x58, 0x2d, 0x89, 0xd9, 0x68, 0x19, 0xe9, 0x6a, 0x30, 0xb2, 0xd5, 0x60, 0xbc,
	0xcd, 0x56, 0x43, 0xbf, 0x9a, 0x8c, 0xc7, 0xe5, 0x0f, 0x1d, 0x0c, 0x2a, 0x38, 0x70, 0x2c, 0xc4,
	0xf1, 0xf2, 0x87, 0xf5, 0x4f, 0xae, 0x63, 0x0d, 0xdc, 0xc4, 0x1a, 0xf8, 0x19, 0x6b, 0xe0, 0x72,
	0xae, 0x15, 0x6e, 0xe6, 0x5a, 0xe1, 0xfb, 0x5c, 0x2b, 0xbc, 0x37, 0x5c, 0x8f, 0x9f, 0x4f, 0x46,
	0xc6, 0x98, 0xfa, 0x66, 0x3a, 0x78, 0x8f, 0x08, 0x1a, 0x31, 0x79, 0x36, 0xa7, 0x4f, 0xcc, 0x4f,
	0x72, 0xcf, 0xf0, 0x8b, 0x10, 0xb3, 0x51, 0x59, 0x58, 0x3f, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff,
	0xe0, 0x5d, 0x71, 0x70, 0xeb, 0x04, 0x00, 0x00,
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
	if len(this.PostsData) != len(that1.PostsData) {
		return false
	}
	for i := range this.PostsData {
		if !this.PostsData[i].Equal(&that1.PostsData[i]) {
			return false
		}
	}
	if len(this.Posts) != len(that1.Posts) {
		return false
	}
	for i := range this.Posts {
		if !this.Posts[i].Equal(&that1.Posts[i]) {
			return false
		}
	}
	if len(this.Attachments) != len(that1.Attachments) {
		return false
	}
	for i := range this.Attachments {
		if !this.Attachments[i].Equal(&that1.Attachments[i]) {
			return false
		}
	}
	if len(this.ActivePolls) != len(that1.ActivePolls) {
		return false
	}
	for i := range this.ActivePolls {
		if !this.ActivePolls[i].Equal(&that1.ActivePolls[i]) {
			return false
		}
	}
	if len(this.UserAnswers) != len(that1.UserAnswers) {
		return false
	}
	for i := range this.UserAnswers {
		if !this.UserAnswers[i].Equal(&that1.UserAnswers[i]) {
			return false
		}
	}
	if !this.Params.Equal(&that1.Params) {
		return false
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
	if this.InitialPostID != that1.InitialPostID {
		return false
	}
	return true
}
func (this *PostDataEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PostDataEntry)
	if !ok {
		that2, ok := that.(PostDataEntry)
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
	if this.PostID != that1.PostID {
		return false
	}
	if this.InitialAttachmentID != that1.InitialAttachmentID {
		return false
	}
	return true
}
func (this *ActivePollData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ActivePollData)
	if !ok {
		that2, ok := that.(ActivePollData)
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
	if this.PostID != that1.PostID {
		return false
	}
	if this.PollID != that1.PollID {
		return false
	}
	if !this.EndDate.Equal(that1.EndDate) {
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.UserAnswers) > 0 {
		for iNdEx := len(m.UserAnswers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserAnswers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ActivePolls) > 0 {
		for iNdEx := len(m.ActivePolls) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ActivePolls[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Attachments) > 0 {
		for iNdEx := len(m.Attachments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attachments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Posts) > 0 {
		for iNdEx := len(m.Posts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Posts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.PostsData) > 0 {
		for iNdEx := len(m.PostsData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PostsData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.InitialPostID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.InitialPostID))
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

func (m *PostDataEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostDataEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostDataEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.InitialAttachmentID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.InitialAttachmentID))
		i--
		dAtA[i] = 0x18
	}
	if m.PostID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PostID))
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

func (m *ActivePollData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivePollData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivePollData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndDate):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if m.PollID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PollID))
		i--
		dAtA[i] = 0x18
	}
	if m.PostID != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.PostID))
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
	if len(m.PostsData) > 0 {
		for _, e := range m.PostsData {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Posts) > 0 {
		for _, e := range m.Posts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Attachments) > 0 {
		for _, e := range m.Attachments {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ActivePolls) > 0 {
		for _, e := range m.ActivePolls {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserAnswers) > 0 {
		for _, e := range m.UserAnswers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
	if m.InitialPostID != 0 {
		n += 1 + sovGenesis(uint64(m.InitialPostID))
	}
	return n
}

func (m *PostDataEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.PostID != 0 {
		n += 1 + sovGenesis(uint64(m.PostID))
	}
	if m.InitialAttachmentID != 0 {
		n += 1 + sovGenesis(uint64(m.InitialAttachmentID))
	}
	return n
}

func (m *ActivePollData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SubspaceID != 0 {
		n += 1 + sovGenesis(uint64(m.SubspaceID))
	}
	if m.PostID != 0 {
		n += 1 + sovGenesis(uint64(m.PostID))
	}
	if m.PollID != 0 {
		n += 1 + sovGenesis(uint64(m.PollID))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndDate)
	n += 1 + l + sovGenesis(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field PostsData", wireType)
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
			m.PostsData = append(m.PostsData, PostDataEntry{})
			if err := m.PostsData[len(m.PostsData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Posts", wireType)
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
			m.Posts = append(m.Posts, Post{})
			if err := m.Posts[len(m.Posts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attachments", wireType)
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
			m.Attachments = append(m.Attachments, Attachment{})
			if err := m.Attachments[len(m.Attachments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivePolls", wireType)
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
			m.ActivePolls = append(m.ActivePolls, ActivePollData{})
			if err := m.ActivePolls[len(m.ActivePolls)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserAnswers", wireType)
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
			m.UserAnswers = append(m.UserAnswers, UserAnswer{})
			if err := m.UserAnswers[len(m.UserAnswers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
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
				return fmt.Errorf("proto: wrong wireType = %d for field InitialPostID", wireType)
			}
			m.InitialPostID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialPostID |= uint64(b&0x7F) << shift
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
func (m *PostDataEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PostDataEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostDataEntry: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			m.PostID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialAttachmentID", wireType)
			}
			m.InitialAttachmentID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialAttachmentID |= uint32(b&0x7F) << shift
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
func (m *ActivePollData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ActivePollData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivePollData: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field PostID", wireType)
			}
			m.PostID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PostID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollID", wireType)
			}
			m.PollID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PollID |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndDate, dAtA[iNdEx:postIndex]); err != nil {
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
