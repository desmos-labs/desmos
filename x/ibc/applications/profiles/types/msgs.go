package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

const (
	TypeMsgCreateApplicationLink = "create-app-link"
)

// NewMsgCreateApplicationLink creates a new MsgCreateApplicationLink instance
// nolint:interfacer
func NewMsgCreateApplicationLink(
	application *ApplicationData, verification *VerificationData, sender sdk.AccAddress,
	sourcePort, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
) *MsgCreateApplicationLink {
	return &MsgCreateApplicationLink{
		SourcePort:       sourcePort,
		SourceChannel:    sourceChannel,
		Sender:           sender.String(),
		Application:      application,
		VerificationData: verification,
		TimeoutHeight:    timeoutHeight,
		TimeoutTimestamp: timeoutTimestamp,
	}
}

// Route implements sdk.Msg
func (MsgCreateApplicationLink) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgCreateApplicationLink) Type() string {
	return TypeMsgCreateApplicationLink
}

// ValidateBasic performs a basic check of the MsgCreateApplicationLink fields.
// NOTE: timeout height or timestamp values can be 0 to disable the timeout.
func (msg MsgCreateApplicationLink) ValidateBasic() error {
	if err := host.PortIdentifierValidator(msg.SourcePort); err != nil {
		return sdkerrors.Wrap(err, "invalid source port ID")
	}
	if err := host.ChannelIdentifierValidator(msg.SourceChannel); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}

	// NOTE: sender format must be validated as it is required by the GetSigners function.
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg.
func (msg MsgCreateApplicationLink) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgCreateApplicationLink) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddr}
}
