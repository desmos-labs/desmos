package types

import (
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

// NewReport returns a new Report instance
func NewReport(subspaceID uint64, id uint64, reasonID uint32, message string, reporter string, data ReportData) Report {
	dataAny, err := codectypes.NewAnyWithValue(data)
	if err != nil {
		panic("failed to pack data to any type")
	}

	return Report{
		SubspaceID: subspaceID,
		ID:         id,
		ReasonID:   reasonID,
		Message:    message,
		Reporter:   reporter,
		Data:       dataAny,
	}
}

// Validate implements fmt.Validator
func (r Report) Validate() error {
	if r.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", r.Size())
	}

	if r.ID == 0 {
		return fmt.Errorf("invalid report id: %d", r.ID)
	}

	if r.ReasonID == 0 {
		return fmt.Errorf("invalid reason id: %d", r.ReasonID)
	}

	_, err := sdk.AccAddressFromBech32(r.Reporter)
	if err != nil {
		return fmt.Errorf("invalid reporter address: %s", err)
	}

	return r.Data.GetCachedValue().(ReportData).Validate()
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *Report) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var data ReportData
	return unpacker.UnpackAny(r.Data, &data)
}

// --------------------------------------------------------------------------------------------------------------------

// ReportData represents a generic report data
type ReportData interface {
	proto.Message

	isReportData()
	Validate() error
}

// --------------------------------------------------------------------------------------------------------------------

var _ ReportData = &UserData{}

// NewUserData returns a new UserData instance
func NewUserData(user sdk.AccAddress) *UserData {
	return &UserData{
		User: user.String(),
	}
}

// isReportData implements ReportData
func (data *UserData) isReportData() {}

// Validate implements ReportData
func (data *UserData) Validate() error {
	_, err := sdk.AccAddressFromBech32(data.User)
	if err != nil {
		return fmt.Errorf("invalid reported user address: %s", err)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

var _ ReportData = &PostData{}

// NewPostData returns a new PostData instance
func NewPostData(postID uint64) *PostData {
	return &PostData{
		PostID: postID,
	}
}

// isReportData implements ReportData
func (data *PostData) isReportData() {}

// Validate implements ReportData
func (data *PostData) Validate() error {
	if data.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", data.PostID)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewReason returns a new Reason instance
func NewReason(subspaceID uint64, id uint32, title string, description string) Reason {
	return Reason{
		SubspaceID:  subspaceID,
		ID:          id,
		Title:       title,
		Description: description,
	}
}

// Validate implements fmt.Validator
func (r Reason) Validate() error {
	if r.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", r.SubspaceID)
	}

	if r.ID == 0 {
		return fmt.Errorf("invalid reason id: %d", r.ID)
	}

	if strings.TrimSpace(r.Title) == "" {
		return fmt.Errorf("invalid reason title: %s", r.Title)
	}

	return nil
}
