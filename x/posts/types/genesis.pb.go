// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/posts/v1beta1/genesis.proto

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

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Posts               Posts                  `protobuf:"bytes,1,rep,name=posts,proto3,castrepeated=Posts" json:"posts"`
	UsersPollAnswers    []UserPollAnswersEntry `protobuf:"bytes,2,rep,name=users_poll_answers,json=usersPollAnswers,proto3" json:"users_poll_answers"`
	PostsReactions      []PostReactionsEntry   `protobuf:"bytes,3,rep,name=posts_reactions,json=postsReactions,proto3" json:"posts_reactions"`
	RegisteredReactions []RegisteredReaction   `protobuf:"bytes,4,rep,name=registered_reactions,json=registeredReactions,proto3" json:"registered_reactions"`
	Params              Params                 `protobuf:"bytes,5,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e358c996c23f0348, []int{0}
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

func (m *GenesisState) GetPosts() Posts {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *GenesisState) GetUsersPollAnswers() []UserPollAnswersEntry {
	if m != nil {
		return m.UsersPollAnswers
	}
	return nil
}

func (m *GenesisState) GetPostsReactions() []PostReactionsEntry {
	if m != nil {
		return m.PostsReactions
	}
	return nil
}

func (m *GenesisState) GetRegisteredReactions() []RegisteredReaction {
	if m != nil {
		return m.RegisteredReactions
	}
	return nil
}

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// UserPollAnswerEntry represents an entry containing all the answers to a poll
type UserPollAnswersEntry struct {
	PostId      string       `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	UserAnswers []UserAnswer `protobuf:"bytes,2,rep,name=user_answers,json=userAnswers,proto3" json:"user_answers"`
}

func (m *UserPollAnswersEntry) Reset()         { *m = UserPollAnswersEntry{} }
func (m *UserPollAnswersEntry) String() string { return proto.CompactTextString(m) }
func (*UserPollAnswersEntry) ProtoMessage()    {}
func (*UserPollAnswersEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e358c996c23f0348, []int{1}
}
func (m *UserPollAnswersEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserPollAnswersEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserPollAnswersEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserPollAnswersEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserPollAnswersEntry.Merge(m, src)
}
func (m *UserPollAnswersEntry) XXX_Size() int {
	return m.Size()
}
func (m *UserPollAnswersEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_UserPollAnswersEntry.DiscardUnknown(m)
}

var xxx_messageInfo_UserPollAnswersEntry proto.InternalMessageInfo

func (m *UserPollAnswersEntry) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *UserPollAnswersEntry) GetUserAnswers() []UserAnswer {
	if m != nil {
		return m.UserAnswers
	}
	return nil
}

// PostReactionEntry represents an entry containing all the reactions to a post
type PostReactionsEntry struct {
	PostId    string         `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Reactions []PostReaction `protobuf:"bytes,2,rep,name=reactions,proto3" json:"reactions"`
}

func (m *PostReactionsEntry) Reset()         { *m = PostReactionsEntry{} }
func (m *PostReactionsEntry) String() string { return proto.CompactTextString(m) }
func (*PostReactionsEntry) ProtoMessage()    {}
func (*PostReactionsEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e358c996c23f0348, []int{2}
}
func (m *PostReactionsEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PostReactionsEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PostReactionsEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PostReactionsEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostReactionsEntry.Merge(m, src)
}
func (m *PostReactionsEntry) XXX_Size() int {
	return m.Size()
}
func (m *PostReactionsEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_PostReactionsEntry.DiscardUnknown(m)
}

var xxx_messageInfo_PostReactionsEntry proto.InternalMessageInfo

func (m *PostReactionsEntry) GetPostId() string {
	if m != nil {
		return m.PostId
	}
	return ""
}

func (m *PostReactionsEntry) GetReactions() []PostReaction {
	if m != nil {
		return m.Reactions
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "desmos.posts.v1beta1.GenesisState")
	proto.RegisterType((*UserPollAnswersEntry)(nil), "desmos.posts.v1beta1.UserPollAnswersEntry")
	proto.RegisterType((*PostReactionsEntry)(nil), "desmos.posts.v1beta1.PostReactionsEntry")
}

func init() {
	proto.RegisterFile("desmos/posts/v1beta1/genesis.proto", fileDescriptor_e358c996c23f0348)
}

var fileDescriptor_e358c996c23f0348 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0x87, 0x13, 0xbb, 0xad, 0xec, 0x74, 0xfd, 0xc3, 0x18, 0x30, 0x14, 0xc9, 0xc6, 0xe0, 0x21,
	0x28, 0x26, 0xec, 0x7a, 0xf3, 0x22, 0x16, 0x56, 0xd9, 0xdb, 0x12, 0x11, 0xc1, 0x83, 0x61, 0xb2,
	0x19, 0x62, 0x20, 0xcd, 0x94, 0x79, 0x27, 0x6a, 0xfd, 0x14, 0x7e, 0x0b, 0xc1, 0x4f, 0xd2, 0x63,
	0x8f, 0x9e, 0x54, 0xda, 0x2f, 0x22, 0x99, 0x99, 0xa4, 0xc5, 0x26, 0xf5, 0x96, 0xc9, 0x3c, 0xbf,
	0xe7, 0x9d, 0x79, 0xdf, 0x41, 0x5e, 0x4a, 0x61, 0xc6, 0x20, 0x9c, 0x33, 0x10, 0x10, 0x7e, 0x3a,
	0x4b, 0xa8, 0x20, 0x67, 0x61, 0x46, 0x4b, 0x0a, 0x39, 0x04, 0x73, 0xce, 0x04, 0xc3, 0x96, 0x62,
	0x02, 0xc9, 0x04, 0x9a, 0x99, 0x58, 0x19, 0xcb, 0x98, 0x04, 0xc2, 0xfa, 0x4b, 0xb1, 0x13, 0xb7,
	0xd3, 0xa7, 0x92, 0x87, 0x89, 0xa2, 0x68, 0x88, 0x47, 0x9d, 0x04, 0xa7, 0xe4, 0x5a, 0xe4, 0xac,
	0x6c, 0xa8, 0x87, 0xdd, 0x1e, 0xc2, 0xc9, 0x4c, 0x23, 0xde, 0xf7, 0x01, 0x3a, 0x79, 0xad, 0xae,
	0xf2, 0x46, 0x10, 0x41, 0xf1, 0x0b, 0x34, 0x94, 0xb8, 0x6d, 0xba, 0x03, 0x7f, 0x7c, 0x3e, 0x09,
	0xba, 0x6e, 0x16, 0x5c, 0x31, 0x10, 0xd3, 0x5b, 0xcb, 0x5f, 0xa7, 0xc6, 0x8f, 0xdf, 0xa7, 0xc3,
	0x7a, 0x05, 0x91, 0xca, 0xe1, 0x0f, 0x08, 0x57, 0x40, 0x39, 0xc4, 0xf5, 0x79, 0x63, 0x52, 0xc2,
	0x67, 0xca, 0xc1, 0xbe, 0x21, 0x6d, 0x8f, 0xbb, 0x6d, 0x6f, 0x81, 0xf2, 0x2b, 0x56, 0x14, 0x2f,
	0x15, 0x7c, 0x51, 0x0a, 0xbe, 0x98, 0x1e, 0xd5, 0xf6, 0xe8, 0xae, 0x74, 0xed, 0x6c, 0xe2, 0x77,
	0xe8, 0x8e, 0x4c, 0xc7, 0xed, 0x6d, 0xed, 0x81, 0x94, 0xfb, 0xfd, 0x47, 0x8d, 0x1a, 0x74, 0x57,
	0x7d, 0x5b, 0x72, 0xed, 0x16, 0x26, 0xc8, 0xe2, 0x34, 0xcb, 0x41, 0x50, 0x4e, 0xd3, 0x1d, 0xfb,
	0xd1, 0x21, 0x7b, 0xd4, 0x26, 0x1a, 0x91, 0xb6, 0xdf, 0xe3, 0x7b, 0x3b, 0x80, 0x9f, 0xa3, 0x91,
	0xea, 0xbe, 0x3d, 0x74, 0x4d, 0x7f, 0x7c, 0xfe, 0xa0, 0xe7, 0xc8, 0x92, 0xd1, 0x22, 0x9d, 0xf0,
	0xbe, 0x22, 0xab, 0xab, 0x4f, 0xf8, 0x3e, 0xba, 0x59, 0xa7, 0xe3, 0x3c, 0xb5, 0x4d, 0xd7, 0xf4,
	0x8f, 0xa3, 0x51, 0xbd, 0xbc, 0x4c, 0xf1, 0x25, 0x3a, 0xa9, 0x9b, 0xf7, 0xcf, 0x08, 0xdc, 0xfe,
	0x11, 0x28, 0xad, 0x2e, 0x3b, 0xae, 0xda, 0x3f, 0xe0, 0x55, 0x08, 0xef, 0xb7, 0xb1, 0xbf, 0xf2,
	0x2b, 0x74, 0xbc, 0x6d, 0x9f, 0x2a, 0xeb, 0xfd, 0x7f, 0x38, 0xba, 0xf0, 0x36, 0x3a, 0xbd, 0x58,
	0xae, 0x1d, 0x73, 0xb5, 0x76, 0xcc, 0x3f, 0x6b, 0xc7, 0xfc, 0xb6, 0x71, 0x8c, 0xd5, 0xc6, 0x31,
	0x7e, 0x6e, 0x1c, 0xe3, 0xfd, 0x93, 0x2c, 0x17, 0x1f, 0xab, 0x24, 0xb8, 0x66, 0xb3, 0x50, 0x89,
	0x9f, 0x16, 0x24, 0x01, 0xfd, 0x1d, 0x7e, 0xd1, 0x4f, 0x5e, 0x2c, 0xe6, 0x14, 0x92, 0x91, 0x7c,
	0xea, 0xcf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x78, 0x5b, 0x30, 0x58, 0xc9, 0x03, 0x00, 0x00,
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
	dAtA[i] = 0x2a
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
			dAtA[i] = 0x22
		}
	}
	if len(m.PostsReactions) > 0 {
		for iNdEx := len(m.PostsReactions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PostsReactions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.UsersPollAnswers) > 0 {
		for iNdEx := len(m.UsersPollAnswers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UsersPollAnswers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UserPollAnswersEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserPollAnswersEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserPollAnswersEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
			dAtA[i] = 0x12
		}
	}
	if len(m.PostId) > 0 {
		i -= len(m.PostId)
		copy(dAtA[i:], m.PostId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PostId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PostReactionsEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PostReactionsEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PostReactionsEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
			dAtA[i] = 0x12
		}
	}
	if len(m.PostId) > 0 {
		i -= len(m.PostId)
		copy(dAtA[i:], m.PostId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PostId)))
		i--
		dAtA[i] = 0xa
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
	if len(m.Posts) > 0 {
		for _, e := range m.Posts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UsersPollAnswers) > 0 {
		for _, e := range m.UsersPollAnswers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PostsReactions) > 0 {
		for _, e := range m.PostsReactions {
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *UserPollAnswersEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.UserAnswers) > 0 {
		for _, e := range m.UserAnswers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *PostReactionsEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PostId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Reactions) > 0 {
		for _, e := range m.Reactions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsersPollAnswers", wireType)
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
			m.UsersPollAnswers = append(m.UsersPollAnswers, UserPollAnswersEntry{})
			if err := m.UsersPollAnswers[len(m.UsersPollAnswers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostsReactions", wireType)
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
			m.PostsReactions = append(m.PostsReactions, PostReactionsEntry{})
			if err := m.PostsReactions[len(m.PostsReactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
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
		case 5:
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
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
func (m *UserPollAnswersEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: UserPollAnswersEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserPollAnswersEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
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
			m.PostId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
func (m *PostReactionsEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PostReactionsEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PostReactionsEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostId", wireType)
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
			m.PostId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
			m.Reactions = append(m.Reactions, PostReaction{})
			if err := m.Reactions[len(m.Reactions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
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
