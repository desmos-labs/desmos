package types

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreatePost{}
	_ sdk.Msg = &MsgEditPost{}
	_ sdk.Msg = &MsgAddPostAttachment{}
	_ sdk.Msg = &MsgRemovePostAttachment{}
	_ sdk.Msg = &MsgDeletePost{}
	_ sdk.Msg = &MsgAnswerPoll{}
)

// NewMsgCreatePost returns a new MsgCreatePost instance
func NewMsgCreatePost(
	subspaceID uint64,
	sectionID uint32,
	externalID string,
	text string,
	conversationID uint64,
	replySettings ReplySetting,
	entities *Entities,
	attachments []AttachmentContent,
	referencedPosts []PostReference,
	author string,
) *MsgCreatePost {
	attachmentsAnis := make([]*codectypes.Any, len(attachments))
	for i, attachment := range attachments {
		attachmentAny, err := codectypes.NewAnyWithValue(attachment)
		if err != nil {
			panic("failed to pack attachment content to any type")
		}
		attachmentsAnis[i] = attachmentAny
	}

	return &MsgCreatePost{
		SubspaceID:      subspaceID,
		SectionID:       sectionID,
		ExternalID:      externalID,
		Text:            text,
		Entities:        entities,
		Attachments:     attachmentsAnis,
		Author:          author,
		ConversationID:  conversationID,
		ReplySettings:   replySettings,
		ReferencedPosts: referencedPosts,
	}
}

// Route implements sdk.Msg
func (msg MsgCreatePost) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgCreatePost) Type() string { return ActionCreatePost }

// ValidateBasic implements sdk.Msg
func (msg MsgCreatePost) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Author)
	if err != nil {
		return fmt.Errorf("invalid author address: %s", err)
	}

	if msg.ReplySettings == REPLY_SETTING_UNSPECIFIED {
		return fmt.Errorf("invalid reply setting: %s", msg.ReplySettings)
	}

	if msg.Entities != nil {
		err := msg.Entities.Validate()
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid entities: %s", err)
		}
	}

	for _, attachment := range msg.Attachments {
		err = attachment.GetCachedValue().(AttachmentContent).Validate()
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid attachment: %s", err)
		}
	}

	for _, reference := range msg.ReferencedPosts {
		err = reference.Validate()
		if err != nil {
			return fmt.Errorf("invalid post reference: %s", err)
		}
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgCreatePost) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgCreatePost) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Author)
	return []sdk.AccAddress{addr}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg *MsgCreatePost) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, attachment := range msg.Attachments {
		var content AttachmentContent
		err := unpacker.UnpackAny(attachment, &content)
		if err != nil {
			return err
		}
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgEditPost returns a new MsgEditPost instance
func NewMsgEditPost(
	subspaceID uint64,
	postID uint64,
	text string,
	entities *Entities,
	editor string,
) *MsgEditPost {
	return &MsgEditPost{
		SubspaceID: subspaceID,
		PostID:     postID,
		Text:       text,
		Entities:   entities,
		Editor:     editor,
	}
}

// Route implements sdk.Msg
func (msg MsgEditPost) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgEditPost) Type() string { return ActionEditPost }

// ValidateBasic implements sdk.Msg
func (msg MsgEditPost) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.Entities != nil {
		err := msg.Entities.Validate()
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid entities: %s", err)
		}
	}

	_, err := sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return fmt.Errorf("invalid editor address: %s", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgEditPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgEditPost) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgAddPostAttachment returns a new MsgAddPostAttachment instance
func NewMsgAddPostAttachment(
	subspaceID uint64,
	postID uint64,
	content AttachmentContent,
	editor string,
) *MsgAddPostAttachment {
	contentAny, err := codectypes.NewAnyWithValue(content)
	if err != nil {
		panic(fmt.Errorf("failed to pack attachment content to any"))
	}

	return &MsgAddPostAttachment{
		SubspaceID: subspaceID,
		PostID:     postID,
		Content:    contentAny,
		Editor:     editor,
	}
}

// Route implements sdk.Msg
func (msg MsgAddPostAttachment) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgAddPostAttachment) Type() string { return ActionAddPostAttachment }

// ValidateBasic implements sdk.Msg
func (msg MsgAddPostAttachment) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.Content == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid attachment content")
	}

	_, err := sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return fmt.Errorf("invalid editor address: %s", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgAddPostAttachment) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAddPostAttachment) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
	return []sdk.AccAddress{addr}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg *MsgAddPostAttachment) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var content AttachmentContent
	return unpacker.UnpackAny(msg.Content, &content)
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgRemovePostAttachment returns a new MsgRemovePostAttachment instance
func NewMsgRemovePostAttachment(subspaceID uint64, postID uint64, attachmentID uint32, editor string) *MsgRemovePostAttachment {
	return &MsgRemovePostAttachment{
		SubspaceID:   subspaceID,
		PostID:       postID,
		AttachmentID: attachmentID,
		Editor:       editor,
	}
}

// Route implements sdk.Msg
func (msg MsgRemovePostAttachment) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgRemovePostAttachment) Type() string { return ActionRemovePostAttachment }

// ValidateBasic implements sdk.Msg
func (msg MsgRemovePostAttachment) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.AttachmentID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid attachment id: %d", msg.AttachmentID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return fmt.Errorf("invalid editor address: %s", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgRemovePostAttachment) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgRemovePostAttachment) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgDeletePost returns a new MsgDeletePost instance
func NewMsgDeletePost(subspaceID uint64, postID uint64, signer string) *MsgDeletePost {
	return &MsgDeletePost{
		SubspaceID: subspaceID,
		PostID:     postID,
		Signer:     signer,
	}
}

// Route implements sdk.Msg
func (msg MsgDeletePost) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgDeletePost) Type() string { return ActionDeletePost }

// ValidateBasic implements sdk.Msg
func (msg MsgDeletePost) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return fmt.Errorf("invalid signer address: %s", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgDeletePost) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgDeletePost) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgAnswerPoll returns a new MsgAnswerPoll instance
func NewMsgAnswerPoll(
	subspaceID uint64,
	postID uint64,
	pollID uint32,
	answersIndexes []uint32,
	signer string,
) *MsgAnswerPoll {
	return &MsgAnswerPoll{
		SubspaceID:     subspaceID,
		PostID:         postID,
		PollID:         pollID,
		AnswersIndexes: answersIndexes,
		Signer:         signer,
	}
}

// Route implements sdk.Msg
func (msg MsgAnswerPoll) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgAnswerPoll) Type() string { return ActionAnswerPoll }

// ValidateBasic implements sdk.Msg
func (msg MsgAnswerPoll) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.PollID == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid poll id: %d", msg.PollID)
	}

	if len(msg.AnswersIndexes) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "at least one answer is required")
	}

	// Check duplicated answers
	answers := map[uint32]bool{}
	for _, answer := range msg.AnswersIndexes {
		if _, ok := answers[answer]; ok {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "duplicated answer index: %d", answer)
		}
		answers[answer] = true
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return fmt.Errorf("invalid signer address: %s", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgAnswerPoll) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgAnswerPoll) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{addr}
}
