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
func MustMarshalReport(cdc codec.BinaryCodec, report Report) []byte {
	return cdc.MustMarshal(&report)
}

// MustUnmarshalReport unmarshal the given byte array to a report using the provided BinaryMarshaler
func MustUnmarshalReport(cdc codec.BinaryCodec, bz []byte) Report {
	var report Report
	cdc.MustUnmarshal(bz, &report)
	return report
}
