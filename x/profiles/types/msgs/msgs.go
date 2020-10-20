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
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
}

// NewMsgRequestDTagTransfer is a constructor function for MsgRequestDtagTransfer
func NewMsgRequestDTagTransfer(receiver, sender sdk.AccAddress) MsgRequestDTagTransfer {
	return MsgRequestDTagTransfer{
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgRequestDTagTransfer) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgRequestDTagTransfer) Type() string { return models.ActionRequestDtag }

func (msg MsgRequestDTagTransfer) ValidateBasic() error {
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Sender.Equals(msg.Receiver) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRequestDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRequestDTagTransfer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// ----------------------
// --- MsgAcceptDTagTransferRequest
// ----------------------

// MsgAcceptDTagTransferRequest represent a DTag transfer acceptance message
type MsgAcceptDTagTransferRequest struct {
	NewDTag  string         `json:"new_d_tag" yaml:"new_dtag"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
}

// NewMsgAcceptDTagTransfer is a constructor for MsgAcceptDTagTransferRequest
func NewMsgAcceptDTagTransfer(newDTag string, receiver, sender sdk.AccAddress) MsgAcceptDTagTransferRequest {
	return MsgAcceptDTagTransferRequest{
		NewDTag:  newDTag,
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgAcceptDTagTransferRequest) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgAcceptDTagTransferRequest) Type() string { return models.ActionAcceptDtagTransfer }

func (msg MsgAcceptDTagTransferRequest) ValidateBasic() error {
	if len(strings.TrimSpace(msg.NewDTag)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new dTag can't be empty")
	}

	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid current owner address: %s", msg.Receiver))
	}

	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiving user address: %s", msg.Sender))
	}

	if msg.Sender.Equals(msg.Receiver) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAcceptDTagTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAcceptDTagTransferRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Receiver}
}

// MsgRefuseDTagTransferRequest represent a DTag request rejection
type MsgRefuseDTagTransferRequest struct {
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
}

// NewMsgRefuseDTagTransferRequest is a constructor for MsgRefuseDTagTransferRequest
func NewMsgRefuseDTagTransferRequest(receiver, sender sdk.AccAddress) MsgRefuseDTagTransferRequest {
	return MsgRefuseDTagTransferRequest{
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgRefuseDTagTransferRequest) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgRefuseDTagTransferRequest) Type() string { return models.RefuseDTagTransferRequest }

func (msg MsgRefuseDTagTransferRequest) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	if msg.Sender.Equals(msg.Receiver) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRefuseDTagTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRefuseDTagTransferRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// MsgCancelDTagTransferRequest represent a DTag request rejection
type MsgCancelDTagTransferRequest struct {
	Sender   sdk.AccAddress `json:"sender" yaml:"sender"`
	Receiver sdk.AccAddress `json:"receiver" yaml:"receiver"`
}

// NewMsgCancelDTagTransferRequest is a constructor for MsgCancelDTagTransferRequest
func NewMsgCancelDTagTransferRequest(sender, receiver sdk.AccAddress) MsgCancelDTagTransferRequest {
	return MsgCancelDTagTransferRequest{
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgCancelDTagTransferRequest) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgCancelDTagTransferRequest) Type() string { return models.CancelDTagTransferRequest }

func (msg MsgCancelDTagTransferRequest) ValidateBasic() error {
	if msg.Receiver.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Receiver.Equals(msg.Sender) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
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
