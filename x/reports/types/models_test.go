package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/app"

	"github.com/desmos-labs/desmos/x/reports/types"

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
			name: "empty report type returns error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("report type cannot be empty"),
		},
		{
			name: "empty report message returns error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"scam",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("report message cannot be empty"),
		},
		{
			name: "invalid report creator returns error",
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
			require.Equal(t, test.expErr, test.report.Validate())
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestAppendIfMissing(t *testing.T) {
	tests := []struct {
		name        string
		reports     []types.Report
		toAppend    types.Report
		expAppended bool
		expReports  []types.Report
	}{
		{
			name:        "report is appended to empty list",
			reports:     []types.Report{},
			toAppend:    types.NewReport("id", "type", "message", "user"),
			expAppended: true,
			expReports: []types.Report{
				types.NewReport("id", "type", "message", "user"),
			},
		},
		{
			name: "not present report is appended properly",
			reports: []types.Report{
				types.NewReport("id", "type", "message", "user"),
			},
			toAppend:    types.NewReport("id", "type", "message_2", "user"),
			expAppended: true,
			expReports: []types.Report{
				types.NewReport("id", "type", "message", "user"),
				types.NewReport("id", "type", "message_2", "user"),
			},
		},
		{
			name: "present report is not appended",
			reports: []types.Report{
				types.NewReport("id", "type", "message", "user"),
			},
			toAppend:    types.NewReport("id", "type", "message", "user"),
			expAppended: false,
			expReports: []types.Report{
				types.NewReport("id", "type", "message", "user"),
			},
		},
	}

	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			actual, appended := types.AppendIfMissing(test.reports, test.toAppend)
			require.Equal(t, appended, test.expAppended)
			require.Equal(t, actual, test.expReports)
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestReportsMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	reports := []types.Report{
		types.NewReport("id", "type", "message", "user"),
		types.NewReport("id", "type", "message_2", "user"),
	}

	bz := types.MustMarshalReports(reports, cdc)
	unmarshalled := types.MustUnmarshalReports(bz, cdc)
	require.Equal(t, reports, unmarshalled)
}
