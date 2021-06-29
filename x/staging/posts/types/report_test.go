package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func TestReport_AreReasonsValid(t *testing.T) {
	useCases := []struct {
		name          string
		report        types.Report
		paramsReports []string
		expBool       bool
	}{
		{
			name: "not contained reports reasons return false",
			report: types.NewReport(
				"",
				[]string{"skam"},
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			paramsReports: []string{"scam", "nudity", "violence"},
			expBool:       false,
		},
		{
			name: "contained reports reasons returns true",
			report: types.NewReport(
				"",
				[]string{"scam", "spam"},
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			paramsReports: []string{"scam", "nudity", "spam"},
			expBool:       true,
		},
	}

	for _, uc := range useCases {
		t.Run(uc.name, func(t *testing.T) {
			res := uc.report.AreReasonsValid(uc.paramsReports)
			require.Equal(t, uc.expBool, res)
		})
	}
}

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
				[]string{"scam"},
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("invalid post id: "),
		},
		{
			name: "empty report type returns error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{""},
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("report reason cannot be empty"),
		},
		{
			name: "empty report message returns error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"scam"},
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("report message cannot be empty"),
		},
		{
			name: "invalid report creator returns error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"scam"},
				"message",
				"",
			),
			expErr: fmt.Errorf("invalid user address: "),
		},
		{
			name: "valid reports returns no error",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"scam"},
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
		reports     []types.Report
		toAppend    types.Report
		expAppended bool
		expReports  []types.Report
	}{
		{
			name:        "report is appended to empty list",
			reports:     []types.Report{},
			toAppend:    types.NewReport("id", []string{"scam"}, "message", "user"),
			expAppended: true,
			expReports: []types.Report{
				types.NewReport("id", []string{"scam"}, "message", "user"),
			},
		},
		{
			name: "not present report is appended properly",
			reports: []types.Report{
				types.NewReport("id", []string{"scam"}, "message", "user"),
			},
			toAppend:    types.NewReport("id", []string{"scam"}, "message_2", "user"),
			expAppended: true,
			expReports: []types.Report{
				types.NewReport("id", []string{"scam"}, "message", "user"),
				types.NewReport("id", []string{"scam"}, "message_2", "user"),
			},
		},
		{
			name: "present report is not appended",
			reports: []types.Report{
				types.NewReport("id", []string{"scam"}, "message", "user"),
			},
			toAppend:    types.NewReport("id", []string{"scam"}, "message", "user"),
			expAppended: false,
			expReports: []types.Report{
				types.NewReport("id", []string{"scam"}, "message", "user"),
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

func TestReportsMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	reports := []types.Report{
		types.NewReport("id", []string{"scam"}, "message", "user"),
		types.NewReport("id", []string{"scam"}, "message_2", "user"),
	}

	bz := types.MustMarshalReports(reports, cdc)
	unmarshalled := types.MustUnmarshalReports(bz, cdc)
	require.Equal(t, reports, unmarshalled)
}
