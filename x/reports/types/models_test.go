package types_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/reports/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReport_Validate(t *testing.T) {
	tests := []struct {
		name   string
		report types.Report
		expErr error
	}{
		{
			name: "invalid post id returns error",
			report: types.NewReport(
				"",
				"scam",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("invalid post id: "),
		},
		{
			name: "empty reports's type returns error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("report type cannot be empty"),
		},
		{
			name: "empty reports's message returns error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("report message cannot be empty"),
		},
		{
			name: "invalid reports's creator returns error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"message",
				"",
			),
			expErr: fmt.Errorf("invalid user address: "),
		},
		{
			name: "valid reports returns no error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actualErr := test.report.Validate()
			require.Equal(t, actualErr, test.expErr)
		})
	}
}
