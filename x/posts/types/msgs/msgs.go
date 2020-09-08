package msgs

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	postserrors "github.com/desmos-labs/desmos/x/posts/types/errors"

	"github.com/desmos-labs/desmos/x/commons"
	"github.com/desmos-labs/desmos/x/posts/types/models"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

// MsgCreatePost defines a CreatePost message
type MsgCreatePost struct {
	ParentID       models.PostID      `json:"parent_id" yaml:"parent_id"`
	Message        string             `json:"message" yaml:"message"`
	AllowsComments bool               `json:"allows_comments" yaml:"allows_comments"`
	Subspace       string             `json:"subspace" yaml:"subspace"`
	OptionalData   map[string]string  `json:"optional_data,omitempty" yaml:"optional_data,omitempty"`
	Creator        sdk.AccAddress     `json:"creator" yaml:"creator"`
	Attachments    models.Attachments `json:"attachments,omitempty" yaml:"attachments,omitempty"`
	PollData       *models.PollData   `json:"poll_data,omitempty" yaml:"poll_data,omitempty"`
}

// NewMsgCreatePost is a constructor function for MsgCreatePost
func NewMsgCreatePost(message string, parentID models.PostID, allowsComments bool, subspace string,
	optionalData map[string]string, owner sdk.AccAddress, attachments models.Attachments, pollData *models.PollData) MsgCreatePost {
	return MsgCreatePost{
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
func (msg MsgCreatePost) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgCreatePost) Type() string { return models.ActionCreatePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreatePost) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Message)) == 0 && len(msg.Attachments) == 0 && msg.PollData == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			"post message, attachments or poll are required and cannot be all blank or empty")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(postserrors.ErrInvalidSubspace, "post subspace must be a valid sha-256 hash")
	}

	if msg.Attachments != nil {
		if err := msg.Attachments.Validate(); err != nil {
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
	PostID      models.PostID      `json:"post_id" yaml:"post_id"`
	Message     string             `json:"message" yaml:"message"`
	Attachments models.Attachments `json:"attachments,omitempty" yaml:"attachments,omitempty"`
	PollData    *models.PollData   `json:"poll_data,omitempty" yaml:"poll_data,omitempty"`
	Editor      sdk.AccAddress     `json:"editor" yaml:"editor"`
}

// NewMsgEditPost is the constructor function for MsgEditPost
func NewMsgEditPost(id models.PostID, message string,
	attachments models.Attachments, pollData *models.PollData, owner sdk.AccAddress) MsgEditPost {
	return MsgEditPost{
		PostID:      id,
		Message:     message,
		Attachments: attachments,
		PollData:    pollData,
		Editor:      owner,
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

	if len(strings.TrimSpace(msg.Message)) == 0 && len(msg.Attachments) == 0 && msg.PollData == nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			"post message, attachments or poll are required and cannot be all blank or empty")
	}

	if msg.Attachments != nil {
		if err := msg.Attachments.Validate(); err != nil {
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
func (msg MsgEditPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Editor}
}
