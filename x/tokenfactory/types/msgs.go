package types

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	tokenfactorytypes "github.com/osmosis-labs/osmosis/v15/x/tokenfactory/types"
)

var (
	_ sdk.Msg            = &MsgCreateDenom{}
	_ legacytx.LegacyMsg = &MsgCreateDenom{}
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

	_, err = tokenfactorytypes.GetTokenDenom(msg.Sender, msg.Subdenom)
	if err != nil {
		return errors.Wrap(tokenfactorytypes.ErrInvalidDenom, err.Error())
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

// GetSignBytes implements legacytx.LegacyMsg
func (msg MsgCreateDenom) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgMint{}
	_ legacytx.LegacyMsg = &MsgMint{}
)

// NewMsgMint creates a new MsgMint instance
func NewMsgMint(subspaceID uint64, sender string, amount sdk.Coin, mintToAddress string) *MsgMint {
	return &MsgMint{
		SubspaceID:    subspaceID,
		Sender:        sender,
		Amount:        amount,
		MintToAddress: mintToAddress,
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

	if !msg.Amount.IsValid() || msg.Amount.Amount.Equal(sdk.ZeroInt()) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	_, err = sdk.AccAddressFromBech32(msg.MintToAddress)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid mint to address: %s", err)
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgMint) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// implements legacytx.LegacyMsg
func (msg MsgMint) Route() string { return RouterKey }

// implements legacytx.LegacyMsg
func (msg MsgMint) Type() string { return ActionMint }

// implements legacytx.LegacyMsg
func (msg MsgMint) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgBurn{}
	_ legacytx.LegacyMsg = &MsgBurn{}
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

	if !msg.Amount.IsValid() || msg.Amount.Amount.Equal(sdk.ZeroInt()) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgBurn) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// implements legacytx.LegacyMsg
func (msg MsgBurn) Route() string { return RouterKey }

// implements legacytx.LegacyMsg
func (msg MsgBurn) Type() string { return ActionBurn }

// implements legacytx.LegacyMsg
func (msg MsgBurn) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgSetDenomMetadata{}
	_ legacytx.LegacyMsg = &MsgSetDenomMetadata{}
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

	_, _, err = tokenfactorytypes.DeconstructDenom(msg.Metadata.Base)
	return err
}

// GetSigners implements sdk.Msg
func (msg MsgSetDenomMetadata) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// implements legacytx.LegacyMsg
func (msg MsgSetDenomMetadata) Route() string { return RouterKey }

// implements legacytx.LegacyMsg
func (msg MsgSetDenomMetadata) Type() string { return ActionSetDenomMetadata }

// implements legacytx.LegacyMsg
func (msg MsgSetDenomMetadata) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ sdk.Msg            = &MsgUpdateParams{}
	_ legacytx.LegacyMsg = &MsgUpdateParams{}
)

// NewMsgCreateDenom creates a new MsgSetDenomMetadata instance

// ValidateBasic implements sdk.Msg
func (msg MsgUpdateParams) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid authority address: %s", err)
	}

	return ToOsmosisTokenFactoryParams(msg.Params).Validate()
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateParams) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Authority)
	return []sdk.AccAddress{sender}
}

// implements legacytx.LegacyMsg
func (msg MsgUpdateParams) Route() string { return RouterKey }

// implements legacytx.LegacyMsg
func (msg MsgUpdateParams) Type() string { return ActionUpdateParams }

// implements legacytx.LegacyMsg
func (msg MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}
