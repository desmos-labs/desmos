package types

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

// MsgCreatePost defines a CreatePost message
type MsgCreatePost struct {
	ParentID       PostID            `json:"parent_id"`
	Message        string            `json:"message"`
	AllowsComments bool              `json:"allows_comments"`
	Subspace       string            `json:"subspace"`
	OptionalData   map[string]string `json:"optional_data,omitempty"`
	Creator        sdk.AccAddress    `json:"creator"`
	CreationDate   time.Time         `json:"creation_date"`
	Medias         PostMedias        `json:"post_medias,omitempty"`
	PollData       *PollData         `json:"poll_data,omitempty"`
}

// NewMsgCreatePost is a constructor function for MsgSetName
func NewMsgCreatePost(message string, parentID PostID, allowsComments bool, subspace string,
	optionalData map[string]string, owner sdk.AccAddress, creationDate time.Time, medias PostMedias, pollData *PollData) MsgCreatePost {
	return MsgCreatePost{
		Message:        message,
		ParentID:       parentID,
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Creator:        owner,
		CreationDate:   creationDate,
		Medias:         medias,
		PollData:       pollData,
	}
}

// MarshalJSON implements the custom marshaling as Amino does not support
// the JSON signature omitempty
func (msg MsgCreatePost) MarshalJSON() ([]byte, error) {
	type msgCreatePost MsgCreatePost
	return json.Marshal(msgCreatePost(msg))
}

// Route should return the name of the module
func (msg MsgCreatePost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreatePost) Type() string { return ActionCreatePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreatePost) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Message)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message cannot be empty nor blank")
	}

	if len(msg.Message) > MaxPostMessageLength {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Post message cannot exceed %d characters", MaxPostMessageLength))
	}

	if !SubspaceRegEx.MatchString(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post subspace must be a valid sha-256 hash")
	}

	if len(msg.OptionalData) > MaxOptionalDataFieldsNumber {
		msg := fmt.Sprintf("Post optional data cannot be longer than %d fields", MaxOptionalDataFieldsNumber)
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg)
	}

	for key, value := range msg.OptionalData {
		if len(value) > MaxOptionalDataFieldValueLength {
			msg := fmt.Sprintf("Post optional data value lengths cannot be longer than %d. %s exceeds the limit",
				MaxOptionalDataFieldValueLength, key)
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg)
		}
	}

	if msg.CreationDate.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post creation date")
	}

	if msg.CreationDate.After(time.Now().UTC()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Creation date cannot be in the future")
	}

	if msg.Medias != nil {
		if err := msg.Medias.Validate(); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	if msg.PollData != nil {

		if !msg.PollData.Open {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Poll Post cannot be created closed")
		}
		if err := msg.PollData.Validate(); err != nil {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreatePost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreatePost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgEditPost
// ----------------------

// MsgEditPost defines the EditPostMessage message
type MsgEditPost struct {
	PostID   PostID         `json:"post_id"`
	Message  string         `json:"message"`
	Editor   sdk.AccAddress `json:"editor"`
	EditDate time.Time      `json:"edit_date"`
}

// NewMsgEditPost is the constructor function for MsgEditPost
func NewMsgEditPost(id PostID, message string, owner sdk.AccAddress, editDate time.Time) MsgEditPost {
	return MsgEditPost{
		PostID:   id,
		Message:  message,
		Editor:   owner,
		EditDate: editDate,
	}
}

// Route should return the name of the module
func (msg MsgEditPost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditPost) Type() string { return ActionEditPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditPost) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id")
	}

	if msg.Editor.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid editor address: %s", msg.Editor))
	}

	if len(strings.TrimSpace(msg.Message)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message cannot be empty nor blank")
	}

	if msg.EditDate.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid edit date")
	}

	if msg.EditDate.After(time.Now().UTC()) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Edit date cannot be in the future")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Editor}
}

// ----------------------
// --- MsgAnswerPollPost
// ----------------------

// MsgAnswerPollPost defines the AnswerPollPost message
type MsgAnswerPollPost struct {
	PostID      PostID         `json:"post_id"`
	UserAnswers []uint         `json:"provided_answers"`
	Answerer    sdk.AccAddress `json:"answerer"`
}

// NewMsgAnswerPollPost is the constructor function for MsgAnswerPollPost
func NewMsgAnswerPollPost(id PostID, providedAnswers []uint, answerer sdk.AccAddress) MsgAnswerPollPost {
	return MsgAnswerPollPost{
		PostID:      id,
		UserAnswers: providedAnswers,
		Answerer:    answerer,
	}
}

// Route should return the name of the module
func (msg MsgAnswerPollPost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAnswerPollPost) Type() string { return ActionAnswerPollPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgAnswerPollPost) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id")
	}

	if msg.Answerer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid answerer address: %s", msg.Answerer))
	}

	if len(msg.UserAnswers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Provided answers must contains at least one answer")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAnswerPollPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAnswerPollPost) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{msg.Answerer} }

// ----------------------
// --- MsgAddPostReaction
// ----------------------

// MsgAddPostReaction defines the message to be used to add a reaction to a post
type MsgAddPostReaction struct {
	PostID PostID         `json:"post_id"` // Id of the post to react to
	Value  string         `json:"value"`   // Reaction of the reaction
	User   sdk.AccAddress `json:"user"`    // Address of the user reacting to the post
}

// NewMsgAddPostReaction is a constructor function for MsgAddPostReaction
func NewMsgAddPostReaction(postID PostID, value string, user sdk.AccAddress) MsgAddPostReaction {
	return MsgAddPostReaction{
		PostID: postID,
		User:   user,
		Value:  value,
	}
}

// Route should return the name of the module
func (msg MsgAddPostReaction) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAddPostReaction) Type() string { return ActionAddPostReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgAddPostReaction) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id")
	}

	if msg.User.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid user address: %s", msg.User))
	}

	if len(strings.TrimSpace(msg.Value)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Reaction value cannot be empty nor blank"))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddPostReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAddPostReaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.User}
}

// ----------------------
// --- MsgRemovePostReaction
// ----------------------

// MsgRemovePostReaction defines the message to be used when wanting to remove
// an existing reaction from a specific user having a specific value
type MsgRemovePostReaction struct {
	PostID   PostID         `json:"post_id"`  // Id of the post to unlike
	User     sdk.AccAddress `json:"user"`     // Address of the user that has previously liked the post
	Reaction string         `json:"reaction"` // Reaction of the reaction to be removed
}

// MsgUnlikePostPost is the constructor of MsgRemovePostReaction
func NewMsgRemovePostReaction(postID PostID, user sdk.AccAddress, reaction string) MsgRemovePostReaction {
	return MsgRemovePostReaction{
		PostID:   postID,
		User:     user,
		Reaction: reaction,
	}
}

// Route should return the name of the module
func (msg MsgRemovePostReaction) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRemovePostReaction) Type() string { return ActionRemovePostReaction }

// ValidateBasic runs stateless checks on the message
func (msg MsgRemovePostReaction) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Invalid post id")
	}

	if msg.User.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid user address: %s", msg.User))
	}

	if len(strings.TrimSpace(msg.Reaction)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Reaction value cannot be empty nor blank")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRemovePostReaction) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRemovePostReaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.User}
}
