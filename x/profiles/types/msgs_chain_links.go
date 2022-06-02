package types

import (
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// nolint:interfacer
func NewMsgLinkChainAccount(
	chainAddress AddressData,
	proof Proof,
	chainConfig ChainConfig,
	signer string,
) *MsgLinkChainAccount {
	addressAny, err := codectypes.NewAnyWithValue(chainAddress)
	if err != nil {
		panic("failed to pack public key to any type")
	}
	return &MsgLinkChainAccount{
		ChainAddress: addressAny,
		Proof:        proof,
		ChainConfig:  chainConfig,
		Signer:       signer,
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
	if msg.ChainAddress == nil {
		return fmt.Errorf("source address cannot be nil")
	}
	if err := msg.Proof.Validate(); err != nil {
		return err
	}
	if err := msg.ChainConfig.Validate(); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Signer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid destination address: %s", msg.Signer)
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
	if err := unpacker.UnpackAny(msg.ChainAddress, &address); err != nil {
		return err
	}

	if err := msg.Proof.UnpackInterfaces(unpacker); err != nil {
		return err
	}
	return nil
}

// GetSigners defines whose signature is required
func (msg MsgLinkChainAccount) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
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

// ___________________________________________________________________________________________________________________

func NewMsgSetDefaultExternalAddress(chainName, address, signer string) *MsgSetDefaultExternalAddress {
	return &MsgSetDefaultExternalAddress{
		ChainName:       chainName,
		ExternalAddress: address,
		Signer:          signer,
	}
}

// Route should return the name of the module
func (msg MsgSetDefaultExternalAddress) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSetDefaultExternalAddress) Type() string {
	return ActionSetDefaultExternalAddress
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSetDefaultExternalAddress) ValidateBasic() error {
	if strings.TrimSpace(msg.ChainName) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "chain name cannot be empty or blank")
	}

	if strings.TrimSpace(msg.ExternalAddress) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid external address")
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid signer")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSetDefaultExternalAddress) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSetDefaultExternalAddress) GetSigners() []sdk.AccAddress {
	signer, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{signer}
}
