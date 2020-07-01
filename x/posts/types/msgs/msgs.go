package msgs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/types/models"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

// MsgCreatePost defines a CreatePost message
type MsgCreatePost struct {
	ParentID       models.PostID     `json:"parent_id" yaml:"parent_id"`
	Message        string            `json:"message" yaml:"message"`
	AllowsComments bool              `json:"allows_comments" yaml:"allows_comments"`
	Subspace       string            `json:"subspace" yaml:"subspace"`
	OptionalData   map[string]string `json:"optional_data,omitempty" yaml:"optional_data,omitempty"`
	Creator        sdk.AccAddress    `json:"creator" yaml:"creator"`
	CreationDate   time.Time         `json:"creation_date" yaml:"creation_date"`
	Medias         models.PostMedias `json:"medias,omitempty" yaml:"medias,omitempty"`
	PollData       *models.PollData  `json:"poll_data,omitempty" yaml:"poll_data,omitempty"`
}

// NewMsgCreatePost is a constructor function for MsgCreatePost
func NewMsgCreatePost(message string, parentID models.PostID, allowsComments bool, subspace string,
	optionalData map[string]string, owner sdk.AccAddress, creationDate time.Time,
	medias models.PostMedias, pollData *models.PollData) MsgCreatePost {
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
func (msg MsgCreatePost) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgCreatePost) Type() string { return models.ActionCreatePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreatePost) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Message)) == 0 && len(msg.Medias) == 0 && msg.PollData == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post message, medias or poll are required and cannot be all blank or empty")
	}

	if !models.Sha256RegEx.MatchString(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post subspace must be a valid sha-256 hash")
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
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
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
	PostID   models.PostID  `json:"post_id" yaml:"post_id"`
	Message  string         `json:"message" yaml:"message"`
	Editor   sdk.AccAddress `json:"editor" yaml:"editor"`
	EditDate time.Time      `json:"edit_date" yaml:"edit_date"`
}

// NewMsgEditPost is the constructor function for MsgEditPost
func NewMsgEditPost(id models.PostID, message string, owner sdk.AccAddress, editDate time.Time) MsgEditPost {
	return MsgEditPost{
		PostID:   id,
		Message:  message,
		Editor:   owner,
		EditDate: editDate,
	}
}

// Route should return the name of the module
func (msg MsgEditPost) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgEditPost) Type() string { return models.ActionEditPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditPost) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Invalid post id: %s", msg.PostID))
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
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Editor}
}
