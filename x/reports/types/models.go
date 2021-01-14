package types

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
)

// NewReport returns a Report
func NewReport(postID string, reportType string, message string, user string) Report {
	return Report{
		PostId:  postID,
		Type:    reportType,
		Message: message,
		User:    user,
	}
}

// Validate implements validator
func (r Report) Validate() error {
	if !postsTypes.IsValidPostID(r.PostId) {
		return fmt.Errorf("invalid post id: %s", r.PostId)
	}

	if len(strings.TrimSpace(r.Type)) == 0 {
		return fmt.Errorf("report type cannot be empty")
	}

	if len(strings.TrimSpace(r.Message)) == 0 {
		return fmt.Errorf("report message cannot be empty")
	}

	if len(r.User) == 0 {
		return fmt.Errorf("invalid user address: %s", r.User)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// AppendIfMissing appends the given report to the provided reports slice if not already present.
// If appended, returns the new slice and true.
// If not appended, returns the original slice and false.
func AppendIfMissing(reports []Report, report Report) (newSlice []Report, appended bool) {
	for _, existing := range reports {
		if existing.Equal(report) {
			return reports, false
		}
	}
	return append(reports, report), true
}

// ___________________________________________________________________________________________________________________

// MustMarshalReports marshals the given reports into an array of bytes.
// Panics on error.
func MustMarshalReports(reports []Report, cdc codec.BinaryMarshaler) []byte {
	bz, err := cdc.MarshalBinaryBare(&Reports{Reports: reports})
	if err != nil {
		panic(err)
	}
	return bz
}

// MustUnmarshalReports tries unmarshalling the given bz to a list of reports.
// Panics on error.
func MustUnmarshalReports(bz []byte, cdc codec.BinaryMarshaler) []Report {
	var wrapped Reports
	err := cdc.UnmarshalBinaryBare(bz, &wrapped)
	if err != nil {
		panic(err)
	}
	return wrapped.Reports
}
