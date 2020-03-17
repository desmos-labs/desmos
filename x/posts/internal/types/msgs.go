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
	Medias         PostMedias        `json:"medias,omitempty"`
	PollData       *PollData         `json:"poll_data,omitempty"`
}

// NewMsgCreatePost is a constructor function for MsgCreatePost
func NewMsgCreatePost(message string, parentID PostID, allowsComments bool, subspace string,
	optionalData map[string]string, owner sdk.AccAddress, creationDate time.Time,
	medias PostMedias, pollData *PollData) MsgCreatePost {
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

// Route should return the name of the module
func (msg MsgCreatePost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreatePost) Type() string { return ActionCreatePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreatePost) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Message)) == 0 && len(msg.Medias) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message or medias are required and cannot be both blank or empty")
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

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreatePost) MarshalJSON() ([]byte, error) {
	type temp MsgCreatePost
	return json.Marshal(temp(msg))
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
