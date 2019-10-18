package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgCreatePost defines a CreatePost message
type MsgCreatePost struct {
	ParentID      string         `json:"parent_id"`
	Message       string         `json:"message"`
	Created       time.Time      `json:"created"`
	Owner         sdk.AccAddress `json:"owner"`
	Namespace     string         `json:"namespace"`
	ExternalOwner string         `json:"external_owner"`
}

// NewMsgCreatePost is a constructor function for MsgSetName
func NewMsgCreatePost(message string, parentID string, time time.Time, owner sdk.AccAddress, namespace string, externalOwner string) MsgCreatePost {
	return MsgCreatePost{
		Message:       message,
		ParentID:      parentID,
		Created:       time,
		Owner:         owner,
		Namespace:     namespace,
		ExternalOwner: externalOwner,
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
	if len(msg.Message) == 0 {
		return sdk.ErrUnknownRequest("Post message cannot be empty")
	}
	if msg.Created.IsZero() {
		return sdk.ErrUnknownRequest("The created time cannot be empty")
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
	PostID        string         `json:"post_id"`
	Created       time.Time      `json:"created"`
	Liker         sdk.AccAddress `json:"liker"`
	Namespace     string         `json:"namespace"`
	ExternalOwner string         `json:"external_owner"`
}

// NewMsgLike is a constructor function for MsgLike
func NewMsgLike(postID string, created time.Time, liker sdk.AccAddress, namespace string, externalOwner string) MsgLike {
	return MsgLike{
		PostID:        postID,
		Created:       created,
		Liker:         liker,
		Namespace:     namespace,
		ExternalOwner: externalOwner,
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
	if len(msg.PostID) == 0 || msg.Created.IsZero() {
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

// MsgCreateSession defines the MsgCreateSession message
type MsgCreateSession struct {
	Owner         sdk.AccAddress `json:"owner"`
	Created       time.Time      `json:"created"`
	Namespace     string         `json:"namespace"`
	ExternalOwner string         `json:"external_owner"`
	Pubkey        string         `json:"pubkey"`
	Signature     string         `json:"signature"`
}

// NewMsgCreateSession is the contructor of MsgCreateSession
func NewMsgCreateSession(created time.Time, owner sdk.AccAddress, namespace string, externalOwner string, pubkey string, signature string) MsgCreateSession {
	return MsgCreateSession{
		Created:       created,
		Owner:         owner,
		Namespace:     namespace,
		ExternalOwner: externalOwner,
		Pubkey:        pubkey,
		Signature:     signature,
	}
}

// Route should return the name of the module
func (msg MsgCreateSession) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateSession) Type() string { return "create_session" }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateSession) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrUnknownRequest("Message owner cannot be empty.")
	}

	if msg.Created.IsZero() {
		return sdk.ErrUnknownRequest("Session created time cannot be empty")
	}

	if len(msg.Namespace) == 0 {
		return sdk.ErrUnknownRequest("Session namespace cannot be empty")
	}

	if len(msg.Pubkey) == 0 {
		return sdk.ErrUnknownRequest("Signer pubkey cannot be empty")
	}

	// the external signer address doesn't have to be exists on Desmos
	if msg.ExternalOwner == "" {
		return sdk.ErrUnknownRequest("Session external owner cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateSession) GetSignBytes() []byte {
	// fmt.Printf("%x", ModuleCdc.MustMarshalJSON(msg))
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateSession) GetSigners() []sdk.AccAddress {
	// addresses := []sdk.AccAddress{}
	// // address, err := sdk.AccAddressFromBech32(msg.ExternalOwner.String())

	// address, err := mputils.GetAccAddressFromExternal(msg.ExternalOwner, msg.Namespace)

	// if err != nil{
	// 	return nil
	// }

	// addresses = append(addresses, address)

	// return addresses
	return []sdk.AccAddress{msg.Owner}
}
