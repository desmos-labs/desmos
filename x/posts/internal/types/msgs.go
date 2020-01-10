package types

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreatePost interface {
	// Return the message type.
	// Must be alphanumeric or empty.
	Route() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Type() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() sdk.Error

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []sdk.AccAddress
}

// ----------------------
// --- MsgCreateTextPost
// ----------------------

// MsgCreateTextPost defines a CreatePost message
type MsgCreateTextPost struct {
	ParentID       PostID            `json:"parent_id"`
	Message        string            `json:"message"`
	AllowsComments bool              `json:"allows_comments"`
	Subspace       string            `json:"subspace"`
	OptionalData   map[string]string `json:"optional_data,omitempty"`
	Creator        sdk.AccAddress    `json:"creator"`
}

// NewMsgCreatePost is a constructor function for MsgSetName
func NewMsgCreatePost(message string, parentID PostID, allowsComments bool, subspace string,
	optionalData map[string]string, owner sdk.AccAddress) MsgCreateTextPost {
	return MsgCreateTextPost{
		Message:        message,
		ParentID:       parentID,
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Creator:        owner,
	}
}

// MarshalJSON implements the custom marshaling as Amino does not support
// the JSON signature omitempty
func (msg MsgCreateTextPost) MarshalJSON() ([]byte, error) {
	type msgCreatePost MsgCreateTextPost
	return json.Marshal(msgCreatePost(msg))
}

// Route should return the name of the module
func (msg MsgCreateTextPost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateTextPost) Type() string { return ActionCreatePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateTextPost) ValidateBasic() sdk.Error {
	if msg.Creator.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}

	if len(strings.TrimSpace(msg.Message)) == 0 {
		return sdk.ErrUnknownRequest("TextPost message cannot be empty nor blank")
	}

	if len(strings.TrimSpace(msg.Subspace)) == 0 {
		return sdk.ErrUnknownRequest("TextPost subspace cannot be empty nor blank")
	}

	if len(msg.OptionalData) > MaxOptionalDataFieldsNumber {
		return sdk.ErrUnknownRequest("TextPost optional data cannot be longer than 10 fields")
	}

	for key, value := range msg.OptionalData {
		if len(value) > MaxOptionalDataFieldValueLength {
			msg := fmt.Sprintf("TextPost optional data value lengths cannot be longer than 200. %s exceeds the limit", key)
			return sdk.ErrUnknownRequest(msg)
		}
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateTextPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateTextPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgCreateMediaPost
// ----------------------

// MsgCreateMediaPost defines a CreateMediaPost message
type MsgCreateMediaPost struct {
	MsgCreatePost MsgCreateTextPost `json:"msg_create_post"`
	Medias        PostMedias        `json:"post_medias"`
}

// NewMsgCreateMediaPost is a constructor function for MsgCreateMediaPost
func NewMsgCreateMediaPost(message string, parentID PostID, allowsComments bool, subspace string,
	optionalData map[string]string, owner sdk.AccAddress, medias PostMedias) MsgCreateMediaPost {
	return MsgCreateMediaPost{
		MsgCreatePost: NewMsgCreatePost(message, parentID, allowsComments, subspace, optionalData, owner),
		Medias:        medias,
	}
}

// Route should return the name of the module
func (msg MsgCreateMediaPost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateMediaPost) Type() string { return ActionCreateMediaPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateMediaPost) ValidateBasic() sdk.Error {
	if err := msg.MsgCreatePost.ValidateBasic(); err != nil {
		return err
	}

	for _, media := range msg.Medias {
		if err := media.Validate(); err != nil {
			return sdk.ErrUnknownRequest(err.Error())
		}
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateMediaPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateMediaPost) GetSigners() []sdk.AccAddress {
	return msg.MsgCreatePost.GetSigners()
}

// ----------------------
// --- MsgEditPost
// ----------------------

// MsgEditPost defines the EditPostMessage message
type MsgEditPost struct {
	PostID  PostID         `json:"post_id"`
	Message string         `json:"message"`
	Editor  sdk.AccAddress `json:"editor"`
}

// NewMsgEditPost is the constructor function for MsgEditPost
func NewMsgEditPost(id PostID, message string, owner sdk.AccAddress) MsgEditPost {
	return MsgEditPost{
		PostID:  id,
		Message: message,
		Editor:  owner,
	}
}

// Route should return the name of the module
func (msg MsgEditPost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditPost) Type() string { return ActionEditPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditPost) ValidateBasic() sdk.Error {
	if !msg.PostID.Valid() {
		return sdk.ErrUnknownRequest("Invalid post id")
	}

	if msg.Editor.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Invalid editor address: %s", msg.Editor))
	}

	if len(msg.Message) == 0 {
		return sdk.ErrUnknownRequest("TextPost message cannot be empty")
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
func (msg MsgAddPostReaction) ValidateBasic() sdk.Error {
	if !msg.PostID.Valid() {
		return sdk.ErrUnknownRequest("Invalid post id")
	}

	if msg.User.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Invalid user address: %s", msg.User))
	}

	if len(strings.TrimSpace(msg.Value)) == 0 {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Reaction value cannot be empty nor blank"))
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
func (msg MsgRemovePostReaction) ValidateBasic() sdk.Error {
	if !msg.PostID.Valid() {
		return sdk.ErrUnknownRequest("Invalid post id")
	}

	if msg.User.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Invalid user address: %s", msg.User))
	}

	if len(strings.TrimSpace(msg.Reaction)) == 0 {
		return sdk.ErrUnknownRequest("Reaction value cannot be empty nor blank")
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
