package types

import (
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgCreateReport{}
	_ sdk.Msg = &MsgDeleteReport{}
	_ sdk.Msg = &MsgSupportStandardReason{}
	_ sdk.Msg = &MsgAddReason{}
	_ sdk.Msg = &MsgRemoveReason{}
)

// NewMsgCreateReport returns a new MsgCreateReport instance
func NewMsgCreateReport(
	subspaceID uint64,
	reasonID uint32,
	message string,
	target ReportTarget,
	reporter string,
) *MsgCreateReport {
	targetAny, err := codectypes.NewAnyWithValue(target)
	if err != nil {
		panic("failed to pack target to any type")
	}

	return &MsgCreateReport{
		SubspaceID: subspaceID,
		ReasonID:   reasonID,
		Message:    message,
		Reporter:   reporter,
		Target:     targetAny,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (msg *MsgCreateReport) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var target ReportTarget
	return unpacker.UnpackAny(msg.Target, &target)
}

// Route should return the name of the module
func (msg MsgCreateReport) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateReport) Type() string {
	return ActionCreateReport
}

// ValidateBasic runs stateless checks on the message
func (msg MsgCreateReport) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.ReasonID == 0 {
		return fmt.Errorf("invalid reason id: %d", msg.ReasonID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Reporter)
	if err != nil {
		return fmt.Errorf("invalid reporter address: %s", err)
	}

	return msg.Target.GetCachedValue().(ReportTarget).Validate()
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateReport) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgCreateReport) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Reporter)
	return []sdk.AccAddress{sender}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgDeleteReport returns a new MsgDeleteReport instance
func NewMsgDeleteReport(subspaceID uint64, reportID uint64, signer string) *MsgDeleteReport {
	return &MsgDeleteReport{
		SubspaceID: subspaceID,
		ReportID:   reportID,
		Signer:     signer,
	}
}

// Route should return the name of the module
func (msg MsgDeleteReport) Route() string { return RouterKey }

// Type should return the action
func (msg MsgDeleteReport) Type() string {
	return ActionDeleteReport
}

// ValidateBasic runs stateless checks on the message
func (msg MsgDeleteReport) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.ReportID == 0 {
		return fmt.Errorf("invalid report id: %d", msg.ReportID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return fmt.Errorf("invalid signer address: %s", err)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgDeleteReport) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgDeleteReport) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{sender}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgSupportStandardReason returns a new MsgSupportStandardReason instance
func NewMsgSupportStandardReason(subspaceID uint64, standardReasonID uint32, signer string) *MsgSupportStandardReason {
	return &MsgSupportStandardReason{
		SubspaceID:       subspaceID,
		StandardReasonID: standardReasonID,
		Signer:           signer,
	}
}

// Route should return the name of the module
func (msg MsgSupportStandardReason) Route() string { return RouterKey }

// Type should return the action
func (msg MsgSupportStandardReason) Type() string {
	return ActionSupportStandardReason
}

// ValidateBasic runs stateless checks on the message
func (msg MsgSupportStandardReason) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.StandardReasonID == 0 {
		return fmt.Errorf("invalid standard reason id: %d", msg.StandardReasonID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return fmt.Errorf("invalid signer address: %s", err)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgSupportStandardReason) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgSupportStandardReason) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{sender}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgAddReason returns a new MsgAddReason instance
func NewMsgAddReason(subspaceID uint64, title string, description string, signer string) *MsgAddReason {
	return &MsgAddReason{
		SubspaceID:  subspaceID,
		Title:       title,
		Description: description,
		Signer:      signer,
	}
}

// Route should return the name of the module
func (msg MsgAddReason) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAddReason) Type() string {
	return ActionAddReason
}

// ValidateBasic runs stateless checks on the message
func (msg MsgAddReason) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}

	if strings.TrimSpace(msg.Title) == "" {
		return fmt.Errorf("invalid reason title: %s", msg.Title)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return fmt.Errorf("invalid signer address: %s", err)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAddReason) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgAddReason) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{sender}
}

// --------------------------------------------------------------------------------------------------------------------

// NewMsgRemoveReason returns a new MsgRemoveReason instance
func NewMsgRemoveReason(subspaceID uint64, reasonID uint32, signer string) *MsgRemoveReason {
	return &MsgRemoveReason{
		SubspaceID: subspaceID,
		ReasonID:   reasonID,
		Signer:     signer,
	}
}

// Route should return the name of the module
func (msg MsgRemoveReason) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRemoveReason) Type() string {
	return ActionRemoveReason
}

// ValidateBasic runs stateless checks on the message
func (msg MsgRemoveReason) ValidateBasic() error {
	if msg.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", msg.SubspaceID)
	}

	if msg.ReasonID == 0 {
		return fmt.Errorf("invalid reason id: %d", msg.ReasonID)
	}

	_, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return fmt.Errorf("invalid signer address: %s", err)
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRemoveReason) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgRemoveReason) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Signer)
	return []sdk.AccAddress{sender}
}
