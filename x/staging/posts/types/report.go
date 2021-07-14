package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewReport returns a Report
func NewReport(postID string, reasons []string, message, user string) Report {
	return Report{
		PostID:  postID,
		Reasons: reasons,
		Message: message,
		User:    user,
	}
}

// AreReasonsValid checks if the report reasons are present inside the paramsReasons
func (r Report) AreReasonsValid(paramsReasons []string) bool {
	exists := make(map[string]bool, len(paramsReasons))
	for _, reason := range paramsReasons {
		exists[reason] = true
	}

	for _, rr := range r.Reasons {
		if !exists[rr] {
			return false
		}
	}
	return true
}

// Validate implements validator
func (r Report) Validate() error {
	if !IsValidPostID(r.PostID) {
		return fmt.Errorf("invalid post id: %s", r.PostID)
	}

	if len(r.Reasons) == 0 {
		return fmt.Errorf("report reasons cannot be empty")
	}

	for _, reason := range r.Reasons {
		if strings.TrimSpace(reason) == "" {
			return fmt.Errorf("report reason cannot be empty or blank")
		}
	}

	if len(strings.TrimSpace(r.Message)) == 0 {
		return fmt.Errorf("report message cannot be empty or blank")
	}

	if len(r.User) == 0 {
		return fmt.Errorf("invalid user address: %s", r.User)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// MustMarshalReport marshal the given report using the given BinaryMarshaler
func MustMarshalReport(cdc codec.BinaryMarshaler, report Report) []byte {
	return cdc.MustMarshalBinaryBare(&report)
}

// MustUnmarshalReport unmarshal the given byte array to a report using the provided BinaryMarshaler
func MustUnmarshalReport(cdc codec.BinaryMarshaler, bz []byte) Report {
	var report Report
	cdc.MustUnmarshalBinaryBare(bz, &report)
	return report
}
