package cli

import (
	"fmt"

	"github.com/spf13/pflag"

	"github.com/desmos-labs/desmos/v7/x/reports/types"
)

const (
	FlagUser     = "user"
	FlagPostID   = "post-id"
	FlagReporter = "reporter"
	FlagMessage  = "message"
)

// ReadReportTarget reads the given flags and returns the report target
func ReadReportTarget(flagSet *pflag.FlagSet) (types.ReportTarget, error) {
	userTarget, err := flagSet.GetString(FlagUser)
	if err != nil {
		return nil, err
	}

	postTarget, err := flagSet.GetUint64(FlagPostID)
	if err != nil {
		return nil, err
	}

	switch {
	case userTarget != "" && postTarget != 0:
		return nil, fmt.Errorf("only one of --%s or --%s must be used", FlagUser, FlagPostID)

	case userTarget != "":
		return types.NewUserTarget(userTarget), nil

	case postTarget != 0:
		return types.NewPostTarget(postTarget), nil
	}

	return nil, nil
}
