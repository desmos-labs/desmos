package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

// ParseReportID parses the given value as a report id, returning an error if it's invalid
func ParseReportID(value string) (uint64, error) {
	if value == "" {
		return 0, nil
	}

	reportID, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid report id: %s", err)
	}
	return reportID, nil
}

// NewReport returns a new Report instance
func NewReport(
	subspaceID uint64,
	id uint64,
	reasonID uint32,
	message string,
	reporter string,
	data ReportData,
	creationDate time.Time,
) Report {
	dataAny, err := codectypes.NewAnyWithValue(data)
	if err != nil {
		panic("failed to pack data to any type")
	}

	return Report{
		SubspaceID:   subspaceID,
		ID:           id,
		ReasonID:     reasonID,
		Message:      message,
		Data:         dataAny,
		Reporter:     reporter,
		CreationDate: creationDate,
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

	err = r.Data.GetCachedValue().(ReportData).Validate()
	if err != nil {
		return err
	}

	if r.CreationDate.IsZero() {
		return fmt.Errorf("invalid report creation date: %s", r.CreationDate)
	}

	return nil
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
func NewUserData(user string) *UserData {
	return &UserData{
		User: user,
	}
}

// isReportData implements ReportData
func (data *UserData) isReportData() {}

// Validate implements ReportData
func (data *UserData) Validate() error {
	// We don't check the validity against sdk.AccAddress because the reported address might be another chain account
	if strings.TrimSpace(data.User) == "" {
		return fmt.Errorf("invalid reported user: %s", data.User)
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

// ParseReasonID parses the given value as a reason id, returning an error if it's invalid
func ParseReasonID(value string) (uint32, error) {
	if value == "" {
		return 0, nil
	}

	reasonID, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid reason id: %s", err)
	}
	return uint32(reasonID), nil
}

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
