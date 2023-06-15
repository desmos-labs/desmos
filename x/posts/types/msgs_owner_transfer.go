package types

import (
	"cosmossdk.io/errors"
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
func (msg *MsgRequestPostOwnerTransfer) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg *MsgRequestPostOwnerTransfer) Type() string {
	return ActionRequestPostOwnerTransfer
}

// ValidateBasic implements sdk.Msg
func (msg *MsgRequestPostOwnerTransfer) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.PostID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %d", msg.PostID)
	}

	if msg.Sender == msg.Receiver {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "receiver cannot be the same as sender")
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
func (msg *MsgRequestPostOwnerTransfer) GetSigners() []sdk.AccAddress {
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// GetSigners implements legacytx.LegacyMsg
func (msg *MsgRequestPostOwnerTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgCancelPostOwnerTransferRequest{}
	_ legacytx.LegacyMsg = &MsgCancelPostOwnerTransferRequest{}
)

// MsgCancelPostOwnerTransferRequest returns a new MsgCancelPostOwnerTransferRequest instance
func NewMsgCancelPostOwnerTransferRequest(subspaceID uint64, postID uint64, sender string) *MsgCancelPostOwnerTransferRequest {
	return &MsgCancelPostOwnerTransferRequest{
		SubspaceID: subspaceID,
		PostID:     postID,
		Sender:     sender,
	}
}

// Route implements legacytx.LegacyMsg
func (msg *MsgCancelPostOwnerTransferRequest) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg *MsgCancelPostOwnerTransferRequest) Type() string {
	return ActionCancelPostOwnerTransfer
}

// ValidateBasic implements sdk.Msg
func (msg *MsgCancelPostOwnerTransferRequest) ValidateBasic() error {
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
func (msg *MsgCancelPostOwnerTransferRequest) GetSigners() []sdk.AccAddress {
	sender := sdk.MustAccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// GetSigners implements legacytx.LegacyMsg
func (msg *MsgCancelPostOwnerTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgAcceptPostOwnerTransferRequest{}
	_ legacytx.LegacyMsg = &MsgAcceptPostOwnerTransferRequest{}
)

// MsgAcceptPostOwnerTransferRequest returns a new MsgAcceptPostOwnerTransferRequest instance
func NewMsgAcceptPostOwnerTransferRequest(subspaceID uint64, postID uint64, receiver string) *MsgAcceptPostOwnerTransferRequest {
	return &MsgAcceptPostOwnerTransferRequest{
		SubspaceID: subspaceID,
		PostID:     postID,
		Receiver:   receiver,
	}
}

// Route implements legacytx.LegacyMsg
func (msg *MsgAcceptPostOwnerTransferRequest) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg *MsgAcceptPostOwnerTransferRequest) Type() string {
	return ActionAcceptPostOwnerTransfer
}

// ValidateBasic implements sdk.Msg
func (msg *MsgAcceptPostOwnerTransferRequest) ValidateBasic() error {
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
func (msg *MsgAcceptPostOwnerTransferRequest) GetSigners() []sdk.AccAddress {
	receiver := sdk.MustAccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{receiver}
}

// GetSigners implements legacytx.LegacyMsg
func (msg *MsgAcceptPostOwnerTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgRefusePostOwnerTransferRequest{}
	_ legacytx.LegacyMsg = &MsgRefusePostOwnerTransferRequest{}
)

// MsgRefusePostOwnerTransferRequest returns a new MsgRefusePostOwnerTransferRequest instance
func NewMsgRefusePostOwnerTransferRequest(subspaceID uint64, postID uint64, receiver string) *MsgRefusePostOwnerTransferRequest {
	return &MsgRefusePostOwnerTransferRequest{
		SubspaceID: subspaceID,
		PostID:     postID,
		Receiver:   receiver,
	}
}

// Route implements legacytx.LegacyMsg
func (msg *MsgRefusePostOwnerTransferRequest) Route() string {
	return RouterKey
}

// Type implements legacytx.LegacyMsg
func (msg *MsgRefusePostOwnerTransferRequest) Type() string {
	return ActionRefusePostOwnerTransfer
}

// ValidateBasic implements sdk.Msg
func (msg *MsgRefusePostOwnerTransferRequest) ValidateBasic() error {
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
func (msg *MsgRefusePostOwnerTransferRequest) GetSigners() []sdk.AccAddress {
	receiver := sdk.MustAccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{receiver}
}

// GetSigners implements legacytx.LegacyMsg
func (msg *MsgRefusePostOwnerTransferRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCodec.MustMarshalJSON(msg))
}
