package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// NewMsgLinkApplication creates a new MsgLinkApplication instance
// nolint:interfacer
func NewMsgLinkApplication(
	linkData Data, callData OracleRequest_CallData, sender sdk.AccAddress,
	sourcePort, sourceChannel string, timeoutHeight clienttypes.Height, timeoutTimestamp uint64,
) *MsgLinkApplication {
	return &MsgLinkApplication{
		Sender:           sender.String(),
		LinkData:         linkData,
		CallData:         callData,
		SourcePort:       sourcePort,
		SourceChannel:    sourceChannel,
		TimeoutHeight:    timeoutHeight,
		TimeoutTimestamp: timeoutTimestamp,
	}
}

// Route implements sdk.Msg
func (MsgLinkApplication) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgLinkApplication) Type() string {
	return ActionLinkApplication
}

// ValidateBasic performs a basic check of the MsgLinkApplication fields.
// NOTE: timeout height or timestamp values can be 0 to disable the timeout.
func (msg MsgLinkApplication) ValidateBasic() error {
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
func (msg MsgLinkApplication) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgLinkApplication) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddr}
}

// -------------------------------------------------------------------------------------------------------------------

// NewMsgUnlinkApplication creates a new MsgUnlinkApplication instance
// nolint:interfacer
func NewMsgUnlinkApplication(application, username string, signer sdk.AccAddress) *MsgUnlinkApplication {
	return &MsgUnlinkApplication{
		Application: application,
		Username:    username,
		Signer:      signer.String(),
	}
}

// Route implements sdk.Msg
func (MsgUnlinkApplication) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgUnlinkApplication) Type() string {
	return ActionUnlinkApplication
}

// ValidateBasic performs a basic check of the MsgUnlinkApplication fields.
// NOTE: timeout height or timestamp values can be 0 to disable the timeout.
func (msg MsgUnlinkApplication) ValidateBasic() error {
	if len(strings.TrimSpace(msg.Application)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "application cannot be empty or blank")
	}
	if len(strings.TrimSpace(msg.Username)) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "username cannot be empty or blank")
	}

	// NOTE: sender format must be validated as it is required by the GetSigners function.
	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg.
func (msg MsgUnlinkApplication) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners implements sdk.Msg
func (msg MsgUnlinkApplication) GetSigners() []sdk.AccAddress {
	valAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddr}
}
