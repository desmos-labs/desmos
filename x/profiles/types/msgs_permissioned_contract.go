package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgSavePermissionedContract
// ----------------------

// NewMsgSavePermissionedContractReference is a constructor function for MsgSavePermissionedContractReference
func NewMsgSavePermissionedContractReference(contractAddress, admin string, message json.RawMessage) *MsgSavePermissionedContractReference {
	return &MsgSavePermissionedContractReference{
		Address: contractAddress,
		Admin:   admin,
		Message: message,
	}
}

// Route should return the name of the module
func (msg MsgSavePermissionedContractReference) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSavePermissionedContractReference) Type() string { return ActionRequestDTag }

func (msg MsgSavePermissionedContractReference) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid contract address: %s", msg.Admin))
	}

	_, err = sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid admin address: %s", msg.Admin))
	}

	if msg.Address == msg.Admin {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the contract address and admin address must be different")
	}

	if !json.Valid(msg.Message) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the json message is not valid")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSavePermissionedContractReference) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSavePermissionedContractReference) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}
