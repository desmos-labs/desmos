package msgs

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models/common"
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
func (msg MsgReportPost) Route() string { return common.RouterKey }

// Type should return the action
func (msg MsgReportPost) Type() string { return common.ActionReportPost }

// ValidateBasic runs stateless checks on the message
func (msg MsgReportPost) ValidateBasic() error {
	if !msg.PostID.Valid() {
		return fmt.Errorf("invalid postID: %s", msg.PostID)
	}

	if err := msg.Report.Validate(); err != nil {
		return err
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
