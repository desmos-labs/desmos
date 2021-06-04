package types

import (
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

// nolint:interfacer
func NewMsgLinkChainAccount(
	sourceAddress AddressData,
	sourceProof Proof,
	sourceChainConfig ChainConfig,
	destinationAddress string,
	destinationProof Proof,
) *MsgLinkChainAccount {
	addressAny, err := codectypes.NewAnyWithValue(sourceAddress)
	if err != nil {
		panic("failed to pack public key to any type")
	}
	return &MsgLinkChainAccount{
		SourceAddress:      addressAny,
		SourceProof:        sourceProof,
		SourceChainConfig:  sourceChainConfig,
		DestinationAddress: destinationAddress,
		DestinationProof:   destinationProof,
	}
}

// Route should return the name of the module
func (msg MsgLinkChainAccount) Route() string { return RouterKey }

// Type should return the action
func (msg MsgLinkChainAccount) Type() string {
	return ActionLinkChainAccount
}

// ValidateBasic runs stateless checks on the message
func (msg MsgLinkChainAccount) ValidateBasic() error {
	if msg.SourceAddress == nil {
		return fmt.Errorf("source address cannot be nil")
	}
	if err := msg.SourceProof.Validate(); err != nil {
		return err
	}
	if err := msg.SourceChainConfig.Validate(); err != nil {
		return err
	}
	if _, err := sdk.AccAddressFromBech32(msg.DestinationAddress); err != nil {
		return fmt.Errorf("invalid destination address: %s", msg.DestinationAddress)
	}
	if err := msg.DestinationProof.Validate(); err != nil {
		return err
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgLinkChainAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg *MsgLinkChainAccount) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var address AddressData
	if err := unpacker.UnpackAny(msg.SourceAddress, &address); err != nil {
		return err
	}
	if err := msg.SourceProof.UnpackInterfaces(unpacker); err != nil {
		return err
	}
	if err := msg.DestinationProof.UnpackInterfaces(unpacker); err != nil {
		return err
	}

	return nil
}

// GetSigners defines whose signature is required
func (msg MsgLinkChainAccount) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.DestinationAddress)
	return []sdk.AccAddress{signer}
}

// ___________________________________________________________________________________________________________________

func NewMsgUnlinkChainAccount(owner, chainName, target string) *MsgUnlinkChainAccount {
	return &MsgUnlinkChainAccount{
		Owner:     owner,
		ChainName: chainName,
		Target:    target,
	}
}

// Route should return the name of the module
func (msg MsgUnlinkChainAccount) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnlinkChainAccount) Type() string {
	return ActionUnlinkChainAccount
}

// ValidateBasic runs stateless checks on the message
func (msg MsgUnlinkChainAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid owner")
	}

	if strings.TrimSpace(msg.ChainName) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "chain name cannot be empty or blank")
	}

	if strings.TrimSpace(msg.Target) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid target")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnlinkChainAccount) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgUnlinkChainAccount) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Owner)
	return []sdk.AccAddress{signer}
}
