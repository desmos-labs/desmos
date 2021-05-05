package types

import (
	"encoding/hex"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

func NewMsgCreateIBCAccountConnection(
	port string,
	channelID string,
	timeoutTimestamp uint64,
	sourceAddress string,
	sourcePubKey string,
	destinationAddress string,
	sourceSignature string,
	destinationSignature string,
) *MsgCreateIBCAccountConnection {
	return &MsgCreateIBCAccountConnection{
		Port:                 port,
		ChannelId:            channelID,
		TimeoutTimestamp:     timeoutTimestamp,
		SourceAddress:        sourceAddress,
		SourcePubKey:         sourcePubKey,
		DestinationAddress:   destinationAddress,
		SourceSignature:      sourceSignature,
		DestinationSignature: destinationSignature,
	}
}

// Route should return the name of the module
func (msg MsgCreateIBCAccountConnection) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateIBCAccountConnection) Type() string { return ActionIBCAccountConnection }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateIBCAccountConnection) ValidateBasic() error {
	if err := host.PortIdentifierValidator(msg.Port); err != nil {
		return sdkerrors.Wrap(err, "invalid source port ID")
	}
	if err := host.ChannelIdentifierValidator(msg.ChannelId); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}
	srcAccAddr, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source address (%s)", err)
	}
	srcPubKeyBz, err := hex.DecodeString(msg.SourcePubKey)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid source pubkey")
	}
	srcSig, err := hex.DecodeString(msg.SourceSignature)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid source signature")
	}
	_, err = hex.DecodeString(msg.DestinationSignature)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid destination signature")
	}
	srcPubKey := &secp256k1.PubKey{Key: srcPubKeyBz}
	if !srcAccAddr.Equals(sdk.AccAddress(srcPubKey.Address().Bytes())) {
		return sdkerrors.Wrap(err, "source pubkey and source address are mismatched")
	}
	link := NewLink(msg.SourceAddress, msg.DestinationAddress)
	linkBz, _ := link.Marshal()
	if !VerifySignature(linkBz, srcSig, srcPubKey) {
		return sdkerrors.Wrap(err, "failed to verify source signature")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateIBCAccountConnection) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCreateIBCAccountConnection) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.SourceAddress)
	return []sdk.AccAddress{sender}
}

// MarshalJSON implements the json.Mashaler interface.
// This is done due to the fact that Amino does not respect omitempty clauses
func (msg MsgCreateIBCAccountConnection) MarshalJSON() ([]byte, error) {
	type temp MsgCreateIBCAccountConnection
	return json.Marshal(temp(msg))
}

// ___________________________________________________________________________________________________________________

func NewMsgCreateIBCAccountLink(
	port string,
	channelID string,
	timeoutTimestamp uint64,
	sourceAddress string,
	sourcePubKey string,
	signature string,
) *MsgCreateIBCAccountLink {
	return &MsgCreateIBCAccountLink{
		Port:             port,
		ChannelId:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		SourceAddress:    sourceAddress,
		SourcePubKey:     sourcePubKey,
		Signature:        signature,
	}
}

// Route should return the name of the module
func (msg MsgCreateIBCAccountLink) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateIBCAccountLink) Type() string { return ActionIBCAccountLink }

// ValidateBasic runs stateless checks on the message
func (msg *MsgCreateIBCAccountLink) ValidateBasic() error {
	if err := host.PortIdentifierValidator(msg.Port); err != nil {
		return sdkerrors.Wrap(err, "invalid source port ID")
	}
	if err := host.ChannelIdentifierValidator(msg.ChannelId); err != nil {
		return sdkerrors.Wrap(err, "invalid source channel ID")
	}
	accAddr, err := sdk.AccAddressFromBech32(msg.SourceAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid source address (%s)", err)
	}
	PubKeyBz, err := hex.DecodeString(msg.SourcePubKey)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid pubkey")
	}
	_, err = hex.DecodeString(msg.Signature)
	if err != nil {
		return sdkerrors.Wrap(err, "invalid signature")
	}
	srcPubKey := &secp256k1.PubKey{Key: PubKeyBz}
	if !accAddr.Equals(sdk.AccAddress(srcPubKey.Address().Bytes())) {
		return sdkerrors.Wrap(err, "pubkey and address are mismatched")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgCreateIBCAccountLink) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg *MsgCreateIBCAccountLink) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.SourceAddress)
	return []sdk.AccAddress{sender}
}
