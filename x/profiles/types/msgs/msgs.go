package msgs

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types/models"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

// MsgSaveProfile defines a SaveProfile message
type MsgSaveProfile struct {
	Dtag       string         `json:"dtag" yaml:"dtag"`
	Moniker    *string        `json:"moniker,omitempty" yaml:"moniker,omitempty"`
	Bio        *string        `json:"bio,omitempty" yaml:"bio,omitempty"`
	ProfilePic *string        `json:"profile_picture,omitempty" yaml:"profile_pic,omitempty"`
	CoverPic   *string        `json:"cover_picture,omitempty" yaml:"cover_pic,omitempty"`
	Creator    sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewMsgSaveProfile is a constructor function for MsgSaveProfile
func NewMsgSaveProfile(dtag string, moniker, bio, profilePic, coverPic *string, creator sdk.AccAddress) MsgSaveProfile {
	return MsgSaveProfile{
		Dtag:       dtag,
		Moniker:    moniker,
		Bio:        bio,
		ProfilePic: profilePic,
		CoverPic:   coverPic,
		Creator:    creator,
	}
}

// Route should return the name of the module
func (msg MsgSaveProfile) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgSaveProfile) Type() string { return models.ActionSaveProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgSaveProfile) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", msg.Creator))
	}

	if strings.TrimSpace(msg.Dtag) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "profile dtag cannot be empty or blank")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSaveProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgSaveProfile) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgDeleteProfile
// ----------------------

// MsgDeleteProfile defines a DeleteProfile message
type MsgDeleteProfile struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
}

// NewMsgDeleteProfile is a constructor function for MsgDeleteProfile
func NewMsgDeleteProfile(creator sdk.AccAddress) MsgDeleteProfile {
	return MsgDeleteProfile{
		Creator: creator,
	}
}

// Route should return the name of the module
func (msg MsgDeleteProfile) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgDeleteProfile) Type() string { return models.ActionDeleteProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteProfile) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator address: %s", msg.Creator))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProfile) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// ----------------------
// --- MsgRequestDTagTransfer
// ----------------------

// MsgRequestDTagTransfer define a Dtag transfer message
type MsgRequestDTagTransfer struct {
	CurrentOwner  sdk.AccAddress `json:"current_owner" yaml:"current_owner"`
	ReceivingUser sdk.AccAddress `json:"receiving_user" yaml:"receiving_user"`
}

// NewMsgRequestDTagTransfer is a constructor function for MsgRequestDtagTransfer
func NewMsgRequestDTagTransfer(owner, receiver sdk.AccAddress) MsgRequestDTagTransfer {
	return MsgRequestDTagTransfer{
		CurrentOwner:  owner,
		ReceivingUser: receiver,
	}
}

// Route should return the name of the module
func (msg MsgRequestDTagTransfer) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgRequestDTagTransfer) Type() string { return models.ActionRequestDtag }

func (msg MsgRequestDTagTransfer) ValidateBasic() error {
	if msg.CurrentOwner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid current owner address: %s", msg.CurrentOwner))
	}

	if msg.ReceivingUser.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiving user address: %s", msg.ReceivingUser))
	}

	if msg.ReceivingUser.Equals(msg.CurrentOwner) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the receiving user and current owner must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRequestDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRequestDTagTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.ReceivingUser}
}

// ----------------------
// --- MsgAcceptDTagTransfer
// ----------------------

// MsgAcceptDTagTransfer represent a DTag transfer acceptance message
type MsgAcceptDTagTransfer struct {
	NewDTag       string         `json:"new_d_tag" yaml:"new_dtag"`
	CurrentOwner  sdk.AccAddress `json:"owner" yaml:"owner"`
	ReceivingUser sdk.AccAddress `json:"receiving_user" yaml:"receiving_user"`
}

// NewMsgAcceptDTagTransfer is a constructor for MsgAcceptDTagTransfer
func NewMsgAcceptDTagTransfer(newDTag string, owner, receivingUser sdk.AccAddress) MsgAcceptDTagTransfer {
	return MsgAcceptDTagTransfer{
		NewDTag:       newDTag,
		CurrentOwner:  owner,
		ReceivingUser: receivingUser,
	}
}

// Route should return the name of the module
func (msg MsgAcceptDTagTransfer) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgAcceptDTagTransfer) Type() string { return models.ActionAcceptDtagTransfer }

func (msg MsgAcceptDTagTransfer) ValidateBasic() error {
	if len(strings.TrimSpace(msg.NewDTag)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new dTag can't be empty")
	}

	if msg.CurrentOwner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid current owner address: %s", msg.CurrentOwner))
	}

	if msg.ReceivingUser.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiving user address: %s", msg.ReceivingUser))
	}

	if msg.ReceivingUser.Equals(msg.CurrentOwner) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the receiving user and current owner must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAcceptDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAcceptDTagTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.CurrentOwner}
}

// MsgRefuseDTagTransferRequest represent a DTag request rejection
type MsgRefuseDTagTransferRequest struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"`
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
}

// NewMsgRefuseDTagTransferRequest is a constructor for MsgRefuseDTagTransferRequest
func NewMsgRefuseDTagTransferRequest(sender, owner sdk.AccAddress) MsgRefuseDTagTransferRequest {
	return MsgRefuseDTagTransferRequest{
		Sender: sender,
		Owner:  owner,
	}
}

// Route should return the name of the module
func (msg MsgRefuseDTagTransferRequest) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgRefuseDTagTransferRequest) Type() string { return models.RefuseDTagTransferRequest }

func (msg MsgRefuseDTagTransferRequest) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid owner address: %s", msg.Owner))
	}

	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Owner.Equals(msg.Sender) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner and sender addresses must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRefuseDTagTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRefuseDTagTransferRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

// MsgCancelDTagTransferRequest represent a DTag request rejection
type MsgCancelDTagTransferRequest struct {
	Sender sdk.AccAddress `json:"sender" yaml:"sender"`
	Owner  sdk.AccAddress `json:"owner" yaml:"owner"`
}

// NewMsgCancelDTagTransferRequest is a constructor for MsgCancelDTagTransferRequest
func NewMsgCancelDTagTransferRequest(sender, owner sdk.AccAddress) MsgCancelDTagTransferRequest {
	return MsgCancelDTagTransferRequest{
		Sender: sender,
		Owner:  owner,
	}
}

// Route should return the name of the module
func (msg MsgCancelDTagTransferRequest) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgCancelDTagTransferRequest) Type() string { return models.CancelDTagTransferRequest }

func (msg MsgCancelDTagTransferRequest) ValidateBasic() error {
	if msg.Owner.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid owner address: %s", msg.Owner))
	}

	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Owner.Equals(msg.Sender) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner and sender addresses must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelDTagTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgCancelDTagTransferRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
