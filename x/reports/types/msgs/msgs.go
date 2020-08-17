package msgs

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	postserrors "github.com/desmos-labs/desmos/x/posts/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types/models"
)

// ----------------------
// --- MsgReportPost
// ----------------------

// MsgReportPost defines a ReportPost message
type MsgReportPost struct {
	PostID posts.PostID  `json:"post_id" yaml:"post_id"`
	Report models.Report `json:"report" yaml:"report"`
}

// NewMsgReportPost returns a MsgReportPost object
func NewMsgReportPost(id posts.PostID, repType, message string, user sdk.AccAddress) MsgReportPost {
	return MsgReportPost{
		PostID: id,
		Report: models.NewReport(repType, message, user),
	}
}

// Route should return the name of the module
func (msg MsgReportPost) Route() string { return models.RouterKey }

// Type should return the action
func (msg MsgReportPost) Type() string { return models.ActionReportPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgReportPost) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return sdkerrors.Wrap(postserrors.ErrInvalidPostID, msg.PostID.String())
	}

	if err := msg.Report.Validate(); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgReportPost) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgsCodec.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgReportPost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Report.User}
}
