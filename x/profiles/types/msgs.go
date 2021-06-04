package types

import (
	"fmt"
	"strings"

	"github.com/desmos-labs/desmos/x/commons"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgSaveProfile
// ----------------------

// NewMsgSaveProfile returns a new MsgSaveProfile instance
func NewMsgSaveProfile(dTag string, nickname, bio, profilePic, coverPic string, creator string) *MsgSaveProfile {
	return &MsgSaveProfile{
		DTag:           dTag,
		Nickname:       nickname,
		Bio:            bio,
		ProfilePicture: profilePic,
		CoverPicture:   coverPic,
		Creator:        creator,
	}
}

// Route should return the name of the module
func (msg MsgSaveProfile) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSaveProfile) Type() string { return ActionSaveProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgSaveProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator: %s", msg.Creator))
	}

	if strings.TrimSpace(msg.DTag) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "profile dtag cannot be empty or blank")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSaveProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSaveProfile) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgDeleteProfile is a constructor function for MsgDeleteProfile
func NewMsgDeleteProfile(creator string) *MsgDeleteProfile {
	return &MsgDeleteProfile{
		Creator: creator,
	}
}

// Route should return the name of the module
func (msg MsgDeleteProfile) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteProfile) Type() string { return ActionDeleteProfile }

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteProfile) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid creator: %s", msg.Creator))
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteProfile) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteProfile) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgRequestDTagTransfer is a constructor function for MsgRequestDTagTransfer
func NewMsgRequestDTagTransfer(sender, receiver string) *MsgRequestDTagTransfer {
	return &MsgRequestDTagTransfer{
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgRequestDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRequestDTagTransfer) Type() string { return ActionRequestDTag }

func (msg MsgRequestDTagTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRequestDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRequestDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgCancelDTagTransferRequest is a constructor for MsgCancelDTagTransfer
func NewMsgCancelDTagTransferRequest(sender, receiver string) *MsgCancelDTagTransfer {
	return &MsgCancelDTagTransfer{
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgCancelDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCancelDTagTransfer) Type() string { return ActionCancelDTagTransferRequest }

func (msg MsgCancelDTagTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	if msg.Receiver == msg.Sender {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCancelDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCancelDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgAcceptDTagTransfer is a constructor for MsgAcceptDTagTransfer
func NewMsgAcceptDTagTransfer(newDTag string, sender, receiver string) *MsgAcceptDTagTransfer {
	return &MsgAcceptDTagTransfer{
		NewDTag:  newDTag,
		Sender:   sender,
		Receiver: receiver,
	}
}

// Route should return the name of the module
func (msg MsgAcceptDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAcceptDTagTransfer) Type() string { return ActionAcceptDTagTransfer }

func (msg MsgAcceptDTagTransfer) ValidateBasic() error {
	if strings.TrimSpace(msg.NewDTag) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "new DTag can't be empty")
	}

	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address: %s", msg.Sender)
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address: %s", msg.Receiver)
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAcceptDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgAcceptDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

// NewMsgRefuseDTagTransferRequest is a constructor for MsgRefuseDTagTransfer
func NewMsgRefuseDTagTransferRequest(sender, receiver string) *MsgRefuseDTagTransfer {
	return &MsgRefuseDTagTransfer{
		Receiver: receiver,
		Sender:   sender,
	}
}

// Route should return the name of the module
func (msg MsgRefuseDTagTransfer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRefuseDTagTransfer) Type() string { return ActionRefuseDTagTransferRequest }

func (msg MsgRefuseDTagTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid sender address: %s", msg.Sender))
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid receiver address: %s", msg.Receiver))
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the sender and receiver must be different")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRefuseDTagTransfer) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRefuseDTagTransfer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.Receiver)
	return []sdk.AccAddress{addr}
}

// ___________________________________________________________________________________________________________________

func NewMsgCreateRelationship(creator, recipient, subspace string) *MsgCreateRelationship {
	return &MsgCreateRelationship{
		Sender:   creator,
		Receiver: recipient,
		Subspace: subspace,
	}
}

// Route should return the name of the module
func (msg MsgCreateRelationship) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateRelationship) Type() string {
	return ActionCreateRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateRelationship) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid receiver address")
	}

	if msg.Sender == msg.Receiver {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "sender and receiver must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateRelationship) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}

// ___________________________________________________________________________________________________________________

func NewMsgDeleteRelationship(user, counterparty, subspace string) *MsgDeleteRelationship {
	return &MsgDeleteRelationship{
		User:         user,
		Counterparty: counterparty,
		Subspace:     subspace,
	}
}

// Route should return the name of the module
func (msg MsgDeleteRelationship) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteRelationship) Type() string {
	return ActionDeleteRelationship
}

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteRelationship) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid user address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Counterparty)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid counterparty address")
	}

	if msg.User == msg.Counterparty {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user and counterparty must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a sha-256")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteRelationship) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteRelationship) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{sender}
}

// ___________________________________________________________________________________________________________________

func NewMsgBlockUser(blocker, blocked, reason, subspace string) *MsgBlockUser {
	return &MsgBlockUser{
		Blocker:  blocker,
		Blocked:  blocked,
		Reason:   reason,
		Subspace: subspace,
	}
}

// Route should return the name of the module
func (msg MsgBlockUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgBlockUser) Type() string {
	return ActionBlockUser
}

// ValidateBasic runs stateless checks on the message
func (msg MsgBlockUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Blocker)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker address")
	}

	_, err = sdk.AccAddressFromBech32(msg.Blocked)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked address")
	}

	if msg.Blocker == msg.Blocked {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "blocker and blocked must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgBlockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgBlockUser) GetSigners() []sdk.AccAddress {
	blocker, _ := sdk.AccAddressFromBech32(msg.Blocker)
	return []sdk.AccAddress{blocker}
}

// ___________________________________________________________________________________________________________________

func NewMsgUnblockUser(blocker, blocked, subspace string) *MsgUnblockUser {
	return &MsgUnblockUser{
		Blocker:  blocker,
		Blocked:  blocked,
		Subspace: subspace,
	}
}

// Route should return the name of the module
func (msg MsgUnblockUser) Route() string { return RouterKey }

// Type should return the action
func (msg MsgUnblockUser) Type() string {
	return ActionUnblockUser
}

// ValidateBasic runs stateless checks on the message
func (msg MsgUnblockUser) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Blocker)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocker")
	}

	_, err = sdk.AccAddressFromBech32(msg.Blocked)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid blocked")
	}

	if msg.Blocker == msg.Blocked {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "blocker and blocked must be different")
	}

	if !commons.IsValidSubspace(msg.Subspace) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "subspace must be a valid sha-256 hash")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgUnblockUser) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgUnblockUser) GetSigners() []sdk.AccAddress {
	blocker, _ := sdk.AccAddressFromBech32(msg.Blocker)
	return []sdk.AccAddress{blocker}
}

// ___________________________________________________________________________________________________________________

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
