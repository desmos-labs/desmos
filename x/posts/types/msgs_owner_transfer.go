package types

import (
	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

var (
	_ sdk.Msg            = &MsgRequestPostOwnerTransfer{}
	_ legacytx.LegacyMsg = &MsgRequestPostOwnerTransfer{}
)

// MsgRequestPostOwnerTransfer returns a new MsgRequestPostOwnerTransfer instance
func NewMsgRequestPostOwnerTransfer(subspaceID uint64, postID uint64, receiver string, sender string) *MsgRequestPostOwnerTransfer {
	return &MsgRequestPostOwnerTransfer{
		SubspaceID: subspaceID,
		PostID:     postID,
		Receiver:   receiver,
		Sender:     sender,
	}
}

// Route implements legacytx.LegacyMsg
func (msg MsgRequestPostOwnerTransfer) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg MsgRequestPostOwnerTransfer) Type() string {
	return ActionRequestPostOwnerTransfer
}

// ValidateBasic implements sdk.Msg
func (msg MsgRequestPostOwnerTransfer) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return err
	}

	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgRequestPostOwnerTransfer) GetSigners() []sdk.AccAddress {
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// GetSigners implements legacytx.LegacyMsg
func (msg MsgRequestPostOwnerTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgCancelPostOwnerTransfer{}
	_ legacytx.LegacyMsg = &MsgCancelPostOwnerTransfer{}
)

// MsgCancelPostOwnerTransfer returns a new MsgCancelPostOwnerTransfer instance
func NewMsgCancelPostOwnerTransfer(subspaceID uint64, postID uint64, sender string) *MsgCancelPostOwnerTransfer {
	return &MsgCancelPostOwnerTransfer{
		SubspaceID: subspaceID,
		PostID:     postID,
		Sender:     sender,
	}
}

// Route implements legacytx.LegacyMsg
func (msg MsgCancelPostOwnerTransfer) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg MsgCancelPostOwnerTransfer) Type() string {
	return ActionCancelPostOwnerTransfer
}

// ValidateBasic implements sdk.Msg
func (msg MsgCancelPostOwnerTransfer) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgCancelPostOwnerTransfer) GetSigners() []sdk.AccAddress {
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// GetSigners implements legacytx.LegacyMsg
func (msg MsgCancelPostOwnerTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgAcceptPostOwnerTransfer{}
	_ legacytx.LegacyMsg = &MsgAcceptPostOwnerTransfer{}
)

// MsgAcceptPostOwnerTransfer returns a new MsgAcceptPostOwnerTransfer instance
func NewMsgAcceptPostOwnerTransfer(subspaceID uint64, postID uint64, receiver string) *MsgAcceptPostOwnerTransfer {
	return &MsgAcceptPostOwnerTransfer{
		SubspaceID: subspaceID,
		PostID:     postID,
		Receiver:   receiver,
	}
}

// Route implements legacytx.LegacyMsg
func (msg MsgAcceptPostOwnerTransfer) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg MsgAcceptPostOwnerTransfer) Type() string {
	return ActionAcceptPostOwnerTransfer
}

// ValidateBasic implements sdk.Msg
func (msg MsgAcceptPostOwnerTransfer) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return err
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgAcceptPostOwnerTransfer) GetSigners() []sdk.AccAddress {
	receiver := sdk.MustAccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{receiver}
}

// GetSigners implements legacytx.LegacyMsg
func (msg MsgAcceptPostOwnerTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgRefusePostOwnerTransfer{}
	_ legacytx.LegacyMsg = &MsgRefusePostOwnerTransfer{}
)

// MsgRefusePostOwnerTransfer returns a new MsgRefusePostOwnerTransfer instance
func NewMsgRefusePostOwnerTransfer(subspaceID uint64, postID uint64, receiver string) *MsgRefusePostOwnerTransfer {
	return &MsgRefusePostOwnerTransfer{
		SubspaceID: subspaceID,
		PostID:     postID,
		Receiver:   receiver,
	}
}

// Route implements legacytx.LegacyMsg
func (msg MsgRefusePostOwnerTransfer) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg MsgRefusePostOwnerTransfer) Type() string {
	return ActionRefusePostOwnerTransfer
}

// ValidateBasic implements sdk.Msg
func (msg MsgRefusePostOwnerTransfer) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return err
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg MsgRefusePostOwnerTransfer) GetSigners() []sdk.AccAddress {
	receiver := sdk.MustAccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{receiver}
}

// GetSigners implements legacytx.LegacyMsg
func (msg MsgRefusePostOwnerTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(&msg))
}
