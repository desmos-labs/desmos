package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/commons"
	"strings"
)

// NewMsgCreateSubspace is a constructor function for MsgCreateSubspace
func NewMsgCreateSubspace(subspace string, creator string) *MsgCreateSubspace {
	return &MsgCreateSubspace{
		Name:    subspace,
		Creator: creator,
	}
}

// Route should return the name of the module
func (msg MsgCreateSubspace) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateSubspace) Type() string { return ActionCreateSubspace }

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateSubspace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator")
	}

	if len(strings.TrimSpace(msg.Name)) == 0 {
		return sdkerrors.Wrap(ErrInvalidSubspaceName, "subspace name should not be empty")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateSubspace) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgCreateSubspace) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreateSubspace) MarshalJSON() ([]byte, error) {
	type temp MsgCreateSubspace
	return json.Marshal(temp(msg))
}

// NewMsgAddAdmin is a constructor function for MsgAddAdmin
func NewMsgAddAdmin(id string, newAdmin, admin string) *MsgAddAdmin {
	return &MsgAddAdmin{
		SubspaceId: id,
		NewAdmin:   newAdmin,
		Admin:      admin,
	}
}

// Route should return the name of the module
func (msg MsgAddAdmin) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAddAdmin) Type() string { return ActionAddAdmin }

// ValidateBasic runs stateless checks on the message
func (msg MsgAddAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Admin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid admin address")
	}

	_, err = sdk.AccAddressFromBech32(msg.NewAdmin)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new admin address")
	}

	if !commons.IsValidSubspace(msg.SubspaceId) {
		return sdkerrors.Wrap(ErrInvalidSubspaceId, "subspace id must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddAdmin) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines the required signature
func (msg MsgAddAdmin) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{addr}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgAddAdmin) MarshalJSON() ([]byte, error) {
	type temp MsgAddAdmin
	return json.Marshal(temp(msg))
}
