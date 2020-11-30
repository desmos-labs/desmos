package types

import (
	"encoding/json"
	"strings"

	emoji "github.com/desmos-labs/Go-Emoji-Utils"

	commonerrors "github.com/desmos-labs/desmos/x/commons/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewMsgCreatePost is a constructor function for MsgCreatePost
func NewMsgCreatePost(
	message string, parentID string, allowsComments bool, subspace string,
	optionalData OptionalData, owner string, attachments Attachments, pollData *PollData,
) *MsgCreatePost {
	return &MsgCreatePost{
		Message:        message,
		ParentID:       parentID,
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Creator:        owner,
		Attachments:    attachments,
		PollData:       pollData,
	}
}

// Route should return the name of the module
func (msg MsgCreatePost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreatePost) Type() string { return ActionCreatePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreatePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator")
	}

	if msg.ParentID != "" && !IsValidPostID(msg.ParentID) {
		return sdkerrors.Wrap(ErrInvalidPostID, msg.ParentID)
	}

	if len(strings.TrimSpace(msg.Message)) == 0 && len(msg.Attachments) == 0 && msg.PollData == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			"post message, attachments or poll are required and cannot be all blank or empty")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(ErrInvalidSubspace, "post subspace must be a valid sha-256 hash")
	}

	for _, attachment := range msg.Attachments {
		err := attachment.Validate()
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	if msg.PollData != nil {
		if err := msg.PollData.Validate(); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreatePost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreatePost) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreatePost) MarshalJSON() ([]byte, error) {
	type temp MsgCreatePost
	return json.Marshal(temp(msg))
}

// ___________________________________________________________________________________________________________________

// NewMsgEditPost is the constructor function for MsgEditPost
func NewMsgEditPost(
	id string, message string, attachments Attachments, pollData *PollData, owner string,
) *MsgEditPost {
	return &MsgEditPost{
		PostID:      id,
		Message:     message,
		Attachments: attachments,
		PollData:    pollData,
		Editor:      owner,
	}
}

// Route should return the name of the module
func (msg MsgEditPost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditPost) Type() string { return ActionEditPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditPost) ValidateBasic() error {
	if !IsValidPostID(msg.PostID) {
		return sdkerrors.Wrap(ErrInvalidPostID, msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid editor")
	}

	if len(strings.TrimSpace(msg.Message)) == 0 && len(msg.Attachments) == 0 && msg.PollData == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			"post message, attachments or poll are required and cannot be all blank or empty")
	}

	for _, attachment := range msg.Attachments {
		err := attachment.Validate()
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	if msg.PollData != nil {
		err := msg.PollData.Validate()
		if err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditPost) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Editor)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgAddPostReaction is a constructor function for MsgAddPostReaction
func NewMsgAddPostReaction(postID string, value string, user string) *MsgAddPostReaction {
	return &MsgAddPostReaction{
		PostID:   postID,
		User:     user,
		Reaction: value,
	}
}

// Route should return the name of the module
func (msg MsgAddPostReaction) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAddPostReaction) Type() string { return ActionAddPostReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgAddPostReaction) ValidateBasic() error {
	if !IsValidPostID(msg.PostID) {
		return sdkerrors.Wrap(ErrInvalidPostID, msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user")
	}

	_, err = emoji.LookupEmoji(msg.Reaction)
	if !IsValidReactionCode(msg.Reaction) && err != nil {
		return sdkerrors.Wrap(ErrInvalidReactionCode, msg.Reaction)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddPostReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgAddPostReaction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgRemovePostReaction is the constructor of MsgRemovePostReaction
func NewMsgRemovePostReaction(postID string, user string, value string) *MsgRemovePostReaction {
	return &MsgRemovePostReaction{
		PostID:   postID,
		User:     user,
		Reaction: value,
	}
}

// Route should return the name of the module
func (msg MsgRemovePostReaction) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRemovePostReaction) Type() string { return ActionRemovePostReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgRemovePostReaction) ValidateBasic() error {
	if !IsValidPostID(msg.PostID) {
		return sdkerrors.Wrap(ErrInvalidPostID, msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user")
	}

	_, err = emoji.LookupEmoji(msg.Reaction)
	if !IsValidReactionCode(msg.Reaction) && err != nil {
		return sdkerrors.Wrap(ErrInvalidReactionCode, msg.Reaction)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRemovePostReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRemovePostReaction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgAnswerPoll is the constructor function for MsgAnswerPoll
func NewMsgAnswerPoll(id string, providedAnswers []string, answerer string) *MsgAnswerPoll {
	return &MsgAnswerPoll{
		PostID:      id,
		UserAnswers: providedAnswers,
		Answerer:    answerer,
	}
}

// Route should return the name of the module
func (msg MsgAnswerPoll) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAnswerPoll) Type() string { return ActionAnswerPoll }

// ValidateBasic runs stateless checks on the message
func (msg MsgAnswerPoll) ValidateBasic() error {
	if !IsValidPostID(msg.PostID) {
		return sdkerrors.Wrap(ErrInvalidPostID, msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Answerer)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid answerer")
	}

	if len(msg.UserAnswers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "provided answer must contains at least one answer")
	}

	for _, answer := range msg.UserAnswers {
		if strings.TrimSpace(answer) == "" {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid answer")
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAnswerPoll) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgAnswerPoll) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Answerer)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgRegisterReaction is a constructor function for MsgRegisterReaction
func NewMsgRegisterReaction(creator string, shortCode, value, subspace string) *MsgRegisterReaction {
	return &MsgRegisterReaction{
		ShortCode: shortCode,
		Value:     value,
		Subspace:  subspace,
		Creator:   creator,
	}
}

// Route should return the name of the module
func (msg MsgRegisterReaction) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRegisterReaction) Type() string { return ActionRegisterReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgRegisterReaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid creator")
	}

	if !IsValidReactionCode(msg.ShortCode) {
		return sdkerrors.Wrap(ErrInvalidReactionCode, msg.ShortCode)
	}

	if !commons.IsURIValid(msg.Value) {
		return sdkerrors.Wrap(commonerrors.ErrInvalidURI, "reaction value should be a valid uri")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(ErrInvalidSubspace, "reaction subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRegisterReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRegisterReaction) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}
