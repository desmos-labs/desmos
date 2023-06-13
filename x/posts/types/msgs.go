package types

import (
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

var (
	_ sdk.Msg = &MsgCreatePost{}
	_ sdk.Msg = &MsgEditPost{}
	_ sdk.Msg = &MsgAddPostAttachment{}
	_ sdk.Msg = &MsgRemovePostAttachment{}
	_ sdk.Msg = &MsgDeletePost{}
	_ sdk.Msg = &MsgAnswerPoll{}

	_ subspacestypes.SocialMsg = &MsgCreatePost{}
	_ subspacestypes.SocialMsg = &MsgEditPost{}
	_ subspacestypes.SocialMsg = &MsgAddPostAttachment{}
	_ subspacestypes.SocialMsg = &MsgRemovePostAttachment{}
	_ subspacestypes.SocialMsg = &MsgDeletePost{}
	_ subspacestypes.SocialMsg = &MsgAnswerPoll{}
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
	tags []string,
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
		Tags:            tags,
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
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
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid entities: %s", err)
		}
	}

	for _, tag := range msg.Tags {
		if strings.TrimSpace(tag) == "" {
			return fmt.Errorf("invalid tag: %s", tag)
		}
	}

	for _, attachment := range msg.Attachments {
		err = attachment.GetCachedValue().(AttachmentContent).Validate()
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid attachment: %s", err)
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

// IsSocialMsg implements subspacestypes.SocialMsg
func (msg MsgCreatePost) IsSocialMsg() {}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgEditPost returns a new MsgEditPost instance
func NewMsgEditPost(
	subspaceID uint64,
	postID uint64,
	text string,
	entities *Entities,
	tags []string,
	editor string,
) *MsgEditPost {
	return &MsgEditPost{
		SubspaceID: subspaceID,
		PostID:     postID,
		Text:       text,
		Entities:   entities,
		Tags:       tags,
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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.Entities != nil {
		err := msg.Entities.Validate()
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid entities: %s", err)
		}
	}

	for _, tag := range msg.Tags {
		if strings.TrimSpace(tag) == "" {
			return fmt.Errorf("invalid tag: %s", tag)
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

// IsSocialMsg implements subspacestypes.SocialMsg
func (msg MsgEditPost) IsSocialMsg() {}

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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.Content == nil {
		return errors.Wrap(sdkerrors.ErrInvalidRequest, "invalid attachment content")
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

// IsSocialMsg implements subspacestypes.SocialMsg
func (msg MsgAddPostAttachment) IsSocialMsg() {}

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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.AttachmentID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid attachment id: %d", msg.AttachmentID)
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

// IsSocialMsg implements subspacestypes.SocialMsg
func (msg MsgRemovePostAttachment) IsSocialMsg() {}

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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
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

// IsSocialMsg implements subspacestypes.SocialMsg
func (msg MsgDeletePost) IsSocialMsg() {}

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
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.PollID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid poll id: %d", msg.PollID)
	}

	if len(msg.AnswersIndexes) == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "at least one answer is required")
	}

	// Check duplicated answers
	answers := map[uint32]bool{}
	for _, answer := range msg.AnswersIndexes {
		if _, ok := answers[answer]; ok {
			return errors.Wrapf(sdkerrors.ErrInvalidRequest, "duplicated answer index: %d", answer)
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

// IsSocialMsg implements subspacestypes.SocialMsg
func (msg MsgAnswerPoll) IsSocialMsg() {}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg                  = &MsgMovePost{}
	_ legacytx.LegacyMsg       = &MsgMovePost{}
	_ subspacestypes.SocialMsg = &MsgMovePost{}
)

// NewMsgMovePost returns a new MsgMovePost instance
func NewMsgMovePost(
	subspaceID uint64,
	postID uint64,
	targetSubspaceID uint64,
	targetSectionID uint32,
	owner string,
) *MsgMovePost {
	return &MsgMovePost{
		SubspaceID:       subspaceID,
		PostID:           postID,
		TargetSubspaceID: targetSubspaceID,
		TargetSectionID:  targetSectionID,
		Owner:            owner,
	}
}

// Route implements sdk.Msg
func (msg MsgMovePost) Route() string { return RouterKey }

// Type implements sdk.Msg
func (msg MsgMovePost) Type() string { return ActionMovePost }

// ValidateBasic implements sdk.Msg
func (msg MsgMovePost) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.TargetSubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid target subspace id: %d", msg.TargetSubspaceID)
	}

	if msg.SubspaceID == msg.TargetSubspaceID {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "can not move to the current subspace with id %d", msg.TargetSectionID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return fmt.Errorf("invalid owner address: %s", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgMovePost) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgMovePost) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{addr}
}

// IsSocialMsg implements subspacestypes.SocialMsg
func (msg MsgMovePost) IsSocialMsg() {}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgUpdateParams{}
	_ legacytx.LegacyMsg = &MsgUpdateParams{}
)

func NewMsgUpdateParams(params Params, authority string) *MsgUpdateParams {
	return &MsgUpdateParams{
		Params:    params,
		Authority: authority,
	}
}

// Route implements legacytx.LegacyMsg
func (msg MsgUpdateParams) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg MsgUpdateParams) Type() string {
	return ActionUpdateParams
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return err
	}

	return msg.Params.Validate()
}

// GetSignBytes implements sdk.Msg
func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	authority := sdk.MustAccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{authority}
}

// GetSigners implements legacytx.LegacyMsg
func (msg MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}
