package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ----------------------
// --- MsgCreatePost
// ----------------------

// MsgCreatePost defines a CreatePost message
type MsgCreatePost struct {
	ParentID          PostID         `json:"parent_id"`
	Message           string         `json:"message"`
	AllowsComments    bool           `json:"allows_comments"`
	ExternalReference string         `json:"external_reference"`
	Creator           sdk.AccAddress `json:"creator"`
}

// NewMsgCreatePost is a constructor function for MsgSetName
func NewMsgCreatePost(message string, parentID PostID, allowsComments bool, externalReference string, owner sdk.AccAddress) MsgCreatePost {
	return MsgCreatePost{
		Message:           message,
		ParentID:          parentID,
		AllowsComments:    allowsComments,
		ExternalReference: externalReference,
		Creator:           owner,
	}
}

// Route should return the name of the module
func (msg MsgCreatePost) Route() string { return QuerierRoute }

// Type should return the action
func (msg MsgCreatePost) Type() string { return ActionCreatePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreatePost) ValidateBasic() sdk.Error {
	if msg.Creator.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Invalid creator address: %s", msg.Creator))
	}
	if len(msg.Message) == 0 {
		return sdk.ErrUnknownRequest("Post message cannot be empty")
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
func (msg MsgEditPost) Route() string { return QuerierRoute }

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
		return sdk.ErrUnknownRequest("Post message cannot be empty")
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
// --- MsgLikePost
// ----------------------

// MsgLikePost defines the MsgLikePost message
type MsgLikePost struct {
	PostID PostID         `json:"post_id"` // Id of the post to like
	Liker  sdk.AccAddress `json:"liker"`   // Address of the user liking the post
}

// NewMsgLikePost is a constructor function for MsgLikePost
func NewMsgLikePost(postID PostID, liker sdk.AccAddress) MsgLikePost {
	return MsgLikePost{
		PostID: postID,
		Liker:  liker,
	}
}

// Route should return the name of the module
func (msg MsgLikePost) Route() string { return QuerierRoute }

// Type should return the action
func (msg MsgLikePost) Type() string { return ActionLikePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgLikePost) ValidateBasic() sdk.Error {
	if !msg.PostID.Valid() {
		return sdk.ErrUnknownRequest("Invalid post id")
	}
	if msg.Liker.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Invalid liker address: %s", msg.Liker))
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgLikePost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgLikePost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Liker}
}

// ----------------------
// --- MsgUnlikePost
// ----------------------

// MsgUnlikePost defines the MsgUnlikePost message
type MsgUnlikePost struct {
	PostID PostID         `json:"post_id"` // Id of the post to unlike
	Liker  sdk.AccAddress `json:"liker"`   // Address of the user that has previously liked the post
}

// MsgUnlikePostPost is the constructor of MsgUnlikePost
func NewMsgUnlikePost(postID PostID, liker sdk.AccAddress) MsgUnlikePost {
	return MsgUnlikePost{
		PostID: postID,
		Liker:  liker,
	}
}

// Route should return the name of the module
func (msg MsgUnlikePost) Route() string { return QuerierRoute }

// Type should return the action
func (msg MsgUnlikePost) Type() string { return ActionUnlikePost }

// ValidateBasic runs stateless checks on the message
func (msg MsgUnlikePost) ValidateBasic() sdk.Error {
	if !msg.PostID.Valid() {
		return sdk.ErrUnknownRequest("Invalid post id")
	}
	if msg.Liker.Empty() {
		return sdk.ErrInvalidAddress(fmt.Sprintf("Invalid liker address: %s", msg.Liker))
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnlikePost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUnlikePost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Liker}
}
