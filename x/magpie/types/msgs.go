package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RouterKey was defined in your key.go file
const RouterKey = ModuleName

// MsgCreatePost defines a CreatePost message
type MsgCreatePost struct {
	Message string         `json:"message"`
	Time    time.Time      `json:"time"`
	Owner   sdk.AccAddress `json:"owner"`
}

// NewMsgCreatePost is a constructor function for MsgSetName
func NewMsgCreatePost(message string, time time.Time, owner sdk.AccAddress) MsgCreatePost {
	return MsgCreatePost{
		Message: message,
		Time:    time,
		Owner:   owner,
	}
}

// Route should return the name of the module
func (msg MsgCreatePost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreatePost) Type() string { return "create_post" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreatePost) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Message) == 0 || msg.Time.IsZero() {
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
	return []sdk.AccAddress{msg.Owner}
}

// MsgEditPost defines the EditPost message
type MsgEditPost struct {
	ID      string         `json:"id"`
	Message string         `json:"message"`
	Time    time.Time      `json:"time"`
	Owner   sdk.AccAddress `json:"owner"`
}

// NewMsgEditPost is the constructor function for MsgEditPost
func NewMsgEditPost(id string, message string, time time.Time, owner sdk.AccAddress) MsgEditPost {
	return MsgEditPost{
		ID:      id,
		Message: message,
		Time:    time,
		Owner:   owner,
	}
}

// Route should return the name of the module
func (msg MsgEditPost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgEditPost) Type() string { return "edit_post" }

// ValidateBasic runs stateless checks on the message
func (msg MsgEditPost) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Message) == 0 || msg.Time.IsZero() || len(msg.ID) == 0 {
		return sdk.ErrUnknownRequest("Post id, message and/or time cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgEditPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgEditPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgLike defines the MsgLike message
type MsgLike struct {
	PostID string         `json:"post_id"`
	Time   time.Time      `json:"time"`
	Liker  sdk.AccAddress `json:"liker"`
}

// NewMsgLike is a constructor function for MsgLike
func NewMsgLike(postID string, time time.Time, liker sdk.AccAddress) MsgLike {
	return MsgLike{
		PostID: postID,
		Time:   time,
		Liker:  liker,
	}
}

// Route should return the name of the module
func (msg MsgLike) Route() string { return RouterKey }

// Type should return the action
func (msg MsgLike) Type() string { return "like" }

// ValidateBasic runs stateless checks on the message
func (msg MsgLike) ValidateBasic() sdk.Error {
	if msg.Liker.Empty() {
		return sdk.ErrInvalidAddress(msg.Liker.String())
	}
	if len(msg.PostID) == 0 || msg.Time.IsZero() {
		return sdk.ErrUnknownRequest("Post id, and/or time cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgLike) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgLike) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Liker}
}

// MsgUnlike defines the MsgUnlike message
type MsgUnlike struct {
	ID    string
	Time  time.Time
	Liker sdk.AccAddress
}

// NewMsgUnlike is the contructor of MsgUnlike
func NewMsgUnlike(id string, time time.Time, liker sdk.AccAddress) MsgUnlike {
	return MsgUnlike{
		ID:    id,
		Time:  time,
		Liker: liker,
	}
}

// Route should return the name of the module
func (msg MsgUnlike) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnlike) Type() string { return "unlike" }

// ValidateBasic runs stateless checks on the message
func (msg MsgUnlike) ValidateBasic() sdk.Error {
	if msg.Liker.Empty() {
		return sdk.ErrInvalidAddress(msg.Liker.String())
	}
	if len(msg.ID) == 0 || msg.Time.IsZero() {
		return sdk.ErrUnknownRequest("Like id, and/or time cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnlike) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgUnlike) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Liker}
}
