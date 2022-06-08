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
	reasonsIDs []uint32,
	message string,
	target ReportTarget,
	reporter string,
	creationDate time.Time,
) Report {
	targetAny, err := codectypes.NewAnyWithValue(target)
	if err != nil {
		panic("failed to pack target to any type")
	}

	return Report{
		SubspaceID:   subspaceID,
		ID:           id,
		ReasonsIDs:   reasonsIDs,
		Message:      message,
		Target:       targetAny,
		Reporter:     reporter,
		CreationDate: creationDate,
	}
}

// Validate implements fmt.Validator
func (r Report) Validate() error {
	if r.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", r.SubspaceID)
	}

	if r.ID == 0 {
		return fmt.Errorf("invalid report id: %d", r.ID)
	}

	if len(r.ReasonsIDs) == 0 {
		return fmt.Errorf("reasons ids cannot be empty")
	}

	for _, reasonID := range r.ReasonsIDs {
		if reasonID == 0 {
			return fmt.Errorf("invalid reason id: %d", reasonID)
		}
	}

	_, err := sdk.AccAddressFromBech32(r.Reporter)
	if err != nil {
		return fmt.Errorf("invalid reporter address: %s", err)
	}

	err = r.Target.GetCachedValue().(ReportTarget).Validate()
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
	var target ReportTarget
	return unpacker.UnpackAny(r.Target, &target)
}

// --------------------------------------------------------------------------------------------------------------------

// ReportTarget represents a generic report target
type ReportTarget interface {
	proto.Message

	isReportData()
	Validate() error
}

// --------------------------------------------------------------------------------------------------------------------

var _ ReportTarget = &UserTarget{}

// NewUserTarget returns a new UserTarget instance
func NewUserTarget(user string) *UserTarget {
	return &UserTarget{
		User: user,
	}
}

// isReportData implements ReportTarget
func (t *UserTarget) isReportData() {}

// Validate implements ReportTarget
func (t *UserTarget) Validate() error {
	// We don't check the validity against sdk.AccAddress because the reported address might be another chain account
	if strings.TrimSpace(t.User) == "" {
		return fmt.Errorf("invalid reported user: %s", t.User)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

var _ ReportTarget = &PostTarget{}

// NewPostTarget returns a new PostTarget instance
func NewPostTarget(postID uint64) *PostTarget {
	return &PostTarget{
		PostID: postID,
	}
}

// isReportData implements ReportTarget
func (t *PostTarget) isReportData() {}

// Validate implements ReportTarget
func (t *PostTarget) Validate() error {
	if t.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", t.PostID)
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

// ParseReasonsIDs parses the given comma-separated values as a list of reasons ids.
func ParseReasonsIDs(value string) ([]uint32, error) {
	strValues := strings.Split(value, ",")
	reasons := make([]uint32, len(strValues))
	for i, str := range strValues {
		reason, err := ParseReasonID(str)
		if err != nil {
			return nil, err
		}
		reasons[i] = reason
	}
	return reasons, nil
}

// ContainsReason returns true iff the given reasons contain the provided reasonID
func ContainsReason(reasons []uint32, reasonID uint32) bool {
	for _, reason := range reasons {
		if reason == reasonID {
			return true
		}
	}
	return false
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
