package types

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var (
	_ sdk.Msg = &MsgCreateDenom{}
)

// NewMsgCreateDenom creates a new MsgCreateDenom instance
func NewMsgCreateDenom(subspaceID uint64, sender, subdenom string) *MsgCreateDenom {
	return &MsgCreateDenom{
		SubspaceID: subspaceID,
		Sender:     sender,
		Subdenom:   subdenom,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgCreateDenom) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "Invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address: %s", err)
	}

	_, err = GetTokenDenom(msg.Sender, msg.Subdenom)
	if err != nil {
		return errors.Wrap(ErrInvalidDenom, err.Error())
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgCreateDenom) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// Route implements legacytx.LegacyMsg
func (msg MsgCreateDenom) Route() string { return RouterKey }

// Type implements legacytx.LegacyMsg
func (msg MsgCreateDenom) Type() string { return ActionCreateDenom }

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg = &MsgMint{}
)

// NewMsgMint creates a new MsgMint instance
func NewMsgMint(subspaceID uint64, sender string, amount sdk.Coin) *MsgMint {
	return &MsgMint{
		SubspaceID: subspaceID,
		Sender:     sender,
		Amount:     amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgMint) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address: %s", err)
	}

	if !msg.Amount.IsValid() || msg.Amount.Amount.Equal(math.ZeroInt()) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// Route implements legacytx.LegacyMsg
func (msg MsgMint) Route() string { return RouterKey }

// Type implements legacytx.LegacyMsg
func (msg MsgMint) Type() string { return ActionMint }

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg = &MsgBurn{}
)

// NewMsgCreateDenom creates a new MsgBurn instance
func NewMsgBurn(subspaceID uint64, sender string, amount sdk.Coin) *MsgBurn {
	return &MsgBurn{
		SubspaceID: subspaceID,
		Sender:     sender,
		Amount:     amount,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgBurn) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address: %s", err)
	}

	if !msg.Amount.IsValid() || msg.Amount.Amount.Equal(math.ZeroInt()) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// Route implements legacytx.LegacyMsg
func (msg MsgBurn) Route() string { return RouterKey }

// Type implements legacytx.LegacyMsg
func (msg MsgBurn) Type() string { return ActionBurn }

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg = &MsgSetDenomMetadata{}
)

// NewMsgCreateDenom creates a new MsgSetDenomMetadata instance
func NewMsgSetDenomMetadata(subspaceID uint64, sender string, metadata banktypes.Metadata) *MsgSetDenomMetadata {
	return &MsgSetDenomMetadata{
		SubspaceID: subspaceID,
		Sender:     sender,
		Metadata:   metadata,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgSetDenomMetadata) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id: %d", msg.SubspaceID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address: %s", err)
	}

	err = msg.Metadata.Validate()
	if err != nil {
		return err
	}

	_, _, err = DeconstructDenom(msg.Metadata.Base)
	return err
}

// GetSigners implements sdk.Msg
func (msg MsgSetDenomMetadata) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// Route implements legacytx.LegacyMsg
func (msg MsgSetDenomMetadata) Route() string { return RouterKey }

// Type implements legacytx.LegacyMsg
func (msg MsgSetDenomMetadata) Type() string { return ActionSetDenomMetadata }

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg = &MsgUpdateParams{}
)

// NewMsgUpdateParams creates a new MsgUpdateParams instance
func NewMsgUpdateParams(params Params, authority string) *MsgUpdateParams {
	return &MsgUpdateParams{
		Params:    params,
		Authority: authority,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid authority address: %s", err)
	}

	return msg.Params.Validate()
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{sender}
}

// Route implements legacytx.LegacyMsg
func (msg MsgUpdateParams) Route() string { return RouterKey }

// Type implements legacytx.LegacyMsg
func (msg MsgUpdateParams) Type() string { return ActionUpdateParams }
