package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ----------------------
// --- MsgAnswerPoll
// ----------------------

// MsgAnswerPoll defines the AnswerPoll message
type MsgAnswerPoll struct {
	PostID      PostID         `json:"post_id" yaml:"post_id"`
	UserAnswers []AnswerID     `json:"answers" yaml:"answers"`
	Answerer    sdk.AccAddress `json:"answerer" yaml:"answerer"`
}

// NewMsgAnswerPoll is the constructor function for MsgAnswerPoll
func NewMsgAnswerPoll(id PostID, providedAnswers []AnswerID, answerer sdk.AccAddress) MsgAnswerPoll {
	return MsgAnswerPoll{
		PostID:      id,
		UserAnswers: providedAnswers,
		Answerer:    answerer,
	}
}

// Route should return the name of the module
func (msg MsgAnswerPoll) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAnswerPoll) Type() string { return ActionAnswerPoll }

// ValidateBasic runs stateless checks on the message
func (msg MsgAnswerPoll) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Invalid post id: %s", msg.PostID))
	}

	if msg.Answerer.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("Invalid answerer address: %s", msg.Answerer))
	}

	if len(msg.UserAnswers) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Provided answers must contains at least one answer")
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgAnswerPoll) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAnswerPoll) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{msg.Answerer} }
