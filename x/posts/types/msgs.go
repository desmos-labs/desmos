package types

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &MsgCreatePost{}
)

// NewMsgCreatePost returns a new MsgCreatePost instance
func NewMsgCreatePost(
	subspaceID uint64,
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
			panic("failed to pack content to any type")
		}
		attachmentsAnis[i] = attachmentAny
	}

	return &MsgCreatePost{
		SubspaceID:      subspaceID,
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
func (a *MsgCreatePost) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, attachment := range a.Attachments {
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
