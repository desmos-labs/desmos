package types_test

import (
	"fmt"
	"testing"

	types2 "github.com/desmos-labs/desmos/x/posts/types"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"
)

func TestReport_Validate(t *testing.T) {
	tests := []struct {
		name   string
		report types2.Report
		expErr error
	}{
		{
			name: "invalid post id returns error",
			report: types2.NewReport(
				"",
				"scam",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("invalid post id: "),
		},
		{
			name: "empty report type returns error",
			report: types2.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("report type cannot be empty"),
		},
		{
			name: "empty report message returns error",
			report: types2.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("report message cannot be empty"),
		},
		{
			name: "invalid report creator returns error",
			report: types2.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"message",
				"",
			),
			expErr: fmt.Errorf("invalid user address: "),
		},
		{
			name: "valid reports returns no error",
			report: types2.NewReport(
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
			require.Equal(t, test.expErr, test.report.Validate())
		})
	}
}

func TestAppendIfMissing(t *testing.T) {
	tests := []struct {
		name        string
		reports     []types2.Report
		toAppend    types2.Report
		expAppended bool
		expReports  []types2.Report
	}{
		{
			name:        "report is appended to empty list",
			reports:     []types2.Report{},
			toAppend:    types2.NewReport("id", "type", "message", "user"),
			expAppended: true,
			expReports: []types2.Report{
				types2.NewReport("id", "type", "message", "user"),
			},
		},
		{
			name: "not present report is appended properly",
			reports: []types2.Report{
				types2.NewReport("id", "type", "message", "user"),
			},
			toAppend:    types2.NewReport("id", "type", "message_2", "user"),
			expAppended: true,
			expReports: []types2.Report{
				types2.NewReport("id", "type", "message", "user"),
				types2.NewReport("id", "type", "message_2", "user"),
			},
		},
		{
			name: "present report is not appended",
			reports: []types2.Report{
				types2.NewReport("id", "type", "message", "user"),
			},
			toAppend:    types2.NewReport("id", "type", "message", "user"),
			expAppended: false,
			expReports: []types2.Report{
				types2.NewReport("id", "type", "message", "user"),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			actual, appended := types2.AppendIfMissing(test.reports, test.toAppend)
			require.Equal(t, appended, test.expAppended)
			require.Equal(t, actual, test.expReports)
		})
	}
}

func TestReportsMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	reports := []types2.Report{
		types2.NewReport("id", "type", "message", "user"),
		types2.NewReport("id", "type", "message_2", "user"),
	}

	bz := types2.MustMarshalReports(reports, cdc)
	unmarshalled := types2.MustUnmarshalReports(bz, cdc)
	require.Equal(t, reports, unmarshalled)
}
