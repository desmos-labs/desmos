package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewMsgReportPost returns a MsgReportPost object
func NewMsgReportPost(id string, reportType, message string, user string) *MsgReportPost {
	return &MsgReportPost{
		PostId:     id,
		ReportType: reportType,
		Message:    message,
		User:       user,
	}
}

// Route should return the name of the module
func (msg MsgReportPost) Route() string { return RouterKey }

// Type should return the action
func (msg MsgReportPost) Type() string { return ActionReportPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgReportPost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.User)
	}

	report := NewReport(msg.PostId, msg.ReportType, msg.Message, msg.User)
	err = report.Validate()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgReportPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners defines whose signature is required
func (msg MsgReportPost) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(msg.User)
	return []sdk.AccAddress{addr}
}
