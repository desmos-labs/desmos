// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/posts/v2/client/cli.proto

package utils

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/desmos-labs/desmos/v5/x/posts/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// CreatePostJSON contains the data that can be specified when creating a Post
// using the CLi command
type CreatePostJSON struct {
	// (optional) External id for this post
	ExternalID string `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	// (optional) Text of the post
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// (optional) Entities connected to this post
	Entities *types.Entities `protobuf:"bytes,3,opt,name=entities,proto3" json:"entities,omitempty"`
	// Tags related to this post
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// Attachments of the post
	Attachments []*types1.Any `protobuf:"bytes,5,rep,name=attachments,proto3" json:"attachments,omitempty"`
	// (optional) Id of the original post of the conversation
	ConversationID uint64 `protobuf:"varint,6,opt,name=conversation_id,json=conversationId,proto3" json:"conversation_id,omitempty"`
	// Reply settings of this post
	ReplySettings types.ReplySetting `protobuf:"varint,7,opt,name=reply_settings,json=replySettings,proto3,enum=desmos.posts.v3.ReplySetting" json:"reply_settings,omitempty"`
	// A list this posts references (either as a reply, repost or quote)
	ReferencedPosts []types.PostReference `protobuf:"bytes,8,rep,name=referenced_posts,json=referencedPosts,proto3" json:"referenced_posts"`
}

func (m *CreatePostJSON) Reset()         { *m = CreatePostJSON{} }
func (m *CreatePostJSON) String() string { return proto.CompactTextString(m) }
func (*CreatePostJSON) ProtoMessage()    {}
func (*CreatePostJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5899b795c71461f, []int{0}
}
func (m *CreatePostJSON) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreatePostJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreatePostJSON.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreatePostJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePostJSON.Merge(m, src)
}
func (m *CreatePostJSON) XXX_Size() int {
	return m.Size()
}
func (m *CreatePostJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePostJSON.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePostJSON proto.InternalMessageInfo

func (m *CreatePostJSON) GetExternalID() string {
	if m != nil {
		return m.ExternalID
	}
	return ""
}

func (m *CreatePostJSON) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CreatePostJSON) GetEntities() *types.Entities {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *CreatePostJSON) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *CreatePostJSON) GetAttachments() []*types1.Any {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *CreatePostJSON) GetConversationID() uint64 {
	if m != nil {
		return m.ConversationID
	}
	return 0
}

func (m *CreatePostJSON) GetReplySettings() types.ReplySetting {
	if m != nil {
		return m.ReplySettings
	}
	return types.REPLY_SETTING_UNSPECIFIED
}

func (m *CreatePostJSON) GetReferencedPosts() []types.PostReference {
	if m != nil {
		return m.ReferencedPosts
	}
	return nil
}

// EditPostJSON contains the data that can be specified when editing a Post
// using the CLI command
type EditPostJSON struct {
	// New text of the post
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// New entities connected to this post
	Entities *types.Entities `protobuf:"bytes,2,opt,name=entities,proto3" json:"entities,omitempty"`
	// New tags associated to this post
	Tags []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (m *EditPostJSON) Reset()         { *m = EditPostJSON{} }
func (m *EditPostJSON) String() string { return proto.CompactTextString(m) }
func (*EditPostJSON) ProtoMessage()    {}
func (*EditPostJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5899b795c71461f, []int{1}
}
func (m *EditPostJSON) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EditPostJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EditPostJSON.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EditPostJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EditPostJSON.Merge(m, src)
}
func (m *EditPostJSON) XXX_Size() int {
	return m.Size()
}
func (m *EditPostJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_EditPostJSON.DiscardUnknown(m)
}

var xxx_messageInfo_EditPostJSON proto.InternalMessageInfo

func (m *EditPostJSON) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *EditPostJSON) GetEntities() *types.Entities {
	if m != nil {
		return m.Entities
	}
	return nil
}

func (m *EditPostJSON) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePostJSON)(nil), "desmos.posts.v3.client.CreatePostJSON")
	proto.RegisterType((*EditPostJSON)(nil), "desmos.posts.v3.client.EditPostJSON")
}

func init() { proto.RegisterFile("desmos/posts/v2/client/cli.proto", fileDescriptor_d5899b795c71461f) }

var fileDescriptor_d5899b795c71461f = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xe3, 0x3a, 0x94, 0xb2, 0x01, 0x17, 0xad, 0x2a, 0xe4, 0x56, 0xe0, 0x58, 0x3d, 0xf9,
	0x82, 0x57, 0x32, 0x94, 0x0b, 0x27, 0xd2, 0xe4, 0x10, 0x0e, 0x04, 0xb9, 0x37, 0x2e, 0xd1, 0xc6,
	0x9e, 0xba, 0x2b, 0xd9, 0xbb, 0x91, 0x77, 0x12, 0x25, 0x6f, 0xc1, 0x83, 0xf0, 0x20, 0x3d, 0xf6,
	0xc8, 0x29, 0x42, 0xce, 0x8b, 0x20, 0xff, 0x6b, 0xa3, 0x44, 0x1c, 0x38, 0x79, 0xc6, 0xdf, 0x6f,
	0x3e, 0xef, 0x7e, 0xeb, 0x25, 0x6e, 0x0c, 0x3a, 0x53, 0x9a, 0xcd, 0x95, 0x46, 0xcd, 0x96, 0x01,
	0x8b, 0x52, 0x01, 0x12, 0xcb, 0x87, 0x3f, 0xcf, 0x15, 0x2a, 0xfa, 0xa6, 0x26, 0xfc, 0x8a, 0xf0,
	0x97, 0x81, 0x5f, 0x13, 0x17, 0x67, 0x89, 0x4a, 0x54, 0x85, 0xb0, 0xb2, 0xaa, 0xe9, 0x8b, 0xf3,
	0x44, 0xa9, 0x24, 0x05, 0x56, 0x75, 0xb3, 0xc5, 0x2d, 0xe3, 0x72, 0xdd, 0x48, 0xfd, 0x7d, 0x09,
	0x45, 0x06, 0x1a, 0x79, 0x36, 0x6f, 0x67, 0x23, 0x55, 0x7e, 0x69, 0x5a, 0x9b, 0xd6, 0x4d, 0x23,
	0xbd, 0xdd, 0x5f, 0x66, 0xa6, 0x62, 0x48, 0x1b, 0xf5, 0xf2, 0x97, 0x49, 0xac, 0xeb, 0x1c, 0x38,
	0xc2, 0x77, 0xa5, 0xf1, 0xeb, 0xcd, 0xe4, 0x1b, 0x65, 0xa4, 0x07, 0x2b, 0x84, 0x5c, 0xf2, 0x74,
	0x2a, 0x62, 0xdb, 0x70, 0x0d, 0xef, 0xc5, 0xc0, 0x2a, 0x36, 0x7d, 0x32, 0x6a, 0x5e, 0x8f, 0x87,
	0x21, 0x69, 0x91, 0x71, 0x4c, 0x29, 0xe9, 0x22, 0xac, 0xd0, 0x3e, 0x2a, 0xc9, 0xb0, 0xaa, 0xe9,
	0x15, 0x39, 0x01, 0x89, 0x02, 0x05, 0x68, 0xdb, 0x74, 0x0d, 0xaf, 0x17, 0x9c, 0xfb, 0xfb, 0x69,
	0x8c, 0x1a, 0x20, 0x7c, 0x44, 0x2b, 0x2b, 0x9e, 0x68, 0xbb, 0xeb, 0x9a, 0x95, 0x15, 0x4f, 0x34,
	0xfd, 0x44, 0x7a, 0x1c, 0x91, 0x47, 0x77, 0x19, 0x48, 0xd4, 0xf6, 0x33, 0xd7, 0xf4, 0x7a, 0xc1,
	0x99, 0x5f, 0x47, 0xe2, 0xb7, 0x91, 0xf8, 0x5f, 0xe4, 0x3a, 0xdc, 0x05, 0xe9, 0x67, 0x72, 0x1a,
	0x29, 0xb9, 0x84, 0x5c, 0x73, 0x14, 0x4a, 0x96, 0x7b, 0x39, 0x76, 0x0d, 0xaf, 0x3b, 0xa0, 0xc5,
	0xa6, 0x6f, 0x5d, 0xef, 0x48, 0xe3, 0x61, 0x68, 0xed, 0xa2, 0xe3, 0x98, 0x0e, 0x89, 0x95, 0xc3,
	0x3c, 0x5d, 0x4f, 0x35, 0x20, 0x0a, 0x99, 0x68, 0xfb, 0xb9, 0x6b, 0x78, 0x56, 0xf0, 0xee, 0x60,
	0x17, 0x61, 0x89, 0xdd, 0xd4, 0x54, 0xf8, 0x2a, 0xdf, 0xe9, 0x34, 0x9d, 0x90, 0xd7, 0x39, 0xdc,
	0x42, 0x0e, 0x32, 0x82, 0x78, 0x5a, 0x8d, 0xd8, 0x27, 0xd5, 0xfa, 0x9d, 0x03, 0x9f, 0x32, 0xff,
	0xb0, 0x85, 0x07, 0xdd, 0xfb, 0x4d, 0xbf, 0x13, 0x9e, 0x3e, 0x4d, 0x97, 0xb2, 0xbe, 0xcc, 0xc8,
	0xcb, 0x51, 0x2c, 0xf0, 0xf1, 0xac, 0xda, 0xe8, 0x8d, 0x7f, 0x44, 0x7f, 0xf4, 0xff, 0xd1, 0x9b,
	0x4f, 0xd1, 0x0f, 0x26, 0xf7, 0x85, 0x63, 0x3c, 0x14, 0x8e, 0xf1, 0xa7, 0x70, 0x8c, 0x9f, 0x5b,
	0xa7, 0xf3, 0xb0, 0x75, 0x3a, 0xbf, 0xb7, 0x4e, 0xe7, 0xc7, 0x55, 0x22, 0xf0, 0x6e, 0x31, 0xf3,
	0x23, 0x95, 0xb1, 0xda, 0xfc, 0x7d, 0xca, 0x67, 0xba, 0xa9, 0xd9, 0xf2, 0x23, 0x5b, 0x35, 0x7f,
	0x5c, 0x73, 0x2b, 0x16, 0x28, 0x52, 0x3d, 0x3b, 0xae, 0x8e, 0xeb, 0xc3, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x0c, 0xb8, 0x9b, 0xa0, 0x3c, 0x03, 0x00, 0x00,
}

func (m *CreatePostJSON) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreatePostJSON) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreatePostJSON) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReferencedPosts) > 0 {
		for iNdEx := len(m.ReferencedPosts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ReferencedPosts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCli(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.ReplySettings != 0 {
		i = encodeVarintCli(dAtA, i, uint64(m.ReplySettings))
		i--
		dAtA[i] = 0x38
	}
	if m.ConversationID != 0 {
		i = encodeVarintCli(dAtA, i, uint64(m.ConversationID))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Attachments) > 0 {
		for iNdEx := len(m.Attachments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attachments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCli(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintCli(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if m.Entities != nil {
		{
			size, err := m.Entities.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCli(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintCli(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ExternalID) > 0 {
		i -= len(m.ExternalID)
		copy(dAtA[i:], m.ExternalID)
		i = encodeVarintCli(dAtA, i, uint64(len(m.ExternalID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EditPostJSON) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EditPostJSON) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EditPostJSON) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintCli(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Entities != nil {
		{
			size, err := m.Entities.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCli(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Text) > 0 {
		i -= len(m.Text)
		copy(dAtA[i:], m.Text)
		i = encodeVarintCli(dAtA, i, uint64(len(m.Text)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCli(dAtA []byte, offset int, v uint64) int {
	offset -= sovCli(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CreatePostJSON) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ExternalID)
	if l > 0 {
		n += 1 + l + sovCli(uint64(l))
	}
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovCli(uint64(l))
	}
	if m.Entities != nil {
		l = m.Entities.Size()
		n += 1 + l + sovCli(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovCli(uint64(l))
		}
	}
	if len(m.Attachments) > 0 {
		for _, e := range m.Attachments {
			l = e.Size()
			n += 1 + l + sovCli(uint64(l))
		}
	}
	if m.ConversationID != 0 {
		n += 1 + sovCli(uint64(m.ConversationID))
	}
	if m.ReplySettings != 0 {
		n += 1 + sovCli(uint64(m.ReplySettings))
	}
	if len(m.ReferencedPosts) > 0 {
		for _, e := range m.ReferencedPosts {
			l = e.Size()
			n += 1 + l + sovCli(uint64(l))
		}
	}
	return n
}

func (m *EditPostJSON) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Text)
	if l > 0 {
		n += 1 + l + sovCli(uint64(l))
	}
	if m.Entities != nil {
		l = m.Entities.Size()
		n += 1 + l + sovCli(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovCli(uint64(l))
		}
	}
	return n
}

func sovCli(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCli(x uint64) (n int) {
	return sovCli(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreatePostJSON) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCli
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
			return fmt.Errorf("proto: CreatePostJSON: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreatePostJSON: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entities == nil {
				m.Entities = &types.Entities{}
			}
			if err := m.Entities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attachments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attachments = append(m.Attachments, &types1.Any{})
			if err := m.Attachments[len(m.Attachments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConversationID", wireType)
			}
			m.ConversationID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ConversationID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplySettings", wireType)
			}
			m.ReplySettings = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplySettings |= types.ReplySetting(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReferencedPosts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReferencedPosts = append(m.ReferencedPosts, types.PostReference{})
			if err := m.ReferencedPosts[len(m.ReferencedPosts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCli(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCli
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
func (m *EditPostJSON) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCli
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
			return fmt.Errorf("proto: EditPostJSON: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EditPostJSON: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Text = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entities == nil {
				m.Entities = &types.Entities{}
			}
			if err := m.Entities.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCli
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
				return ErrInvalidLengthCli
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCli
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCli(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCli
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
func skipCli(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCli
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
					return 0, ErrIntOverflowCli
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
					return 0, ErrIntOverflowCli
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
				return 0, ErrInvalidLengthCli
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCli
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCli
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCli        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCli          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCli = fmt.Errorf("proto: unexpected end of group")
)
