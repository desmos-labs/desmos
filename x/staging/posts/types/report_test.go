package types_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/app"
	"testing"

	"github.com/stretchr/testify/require"

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

func TestReport_Marshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	report := types.NewReport(
		"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
		[]string{"scam"},
		"this is a test",
		"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
	)
	marshaled := types.MustMarshalReport(cdc, report)
	unmarshaled := types.MustUnmarshalReport(cdc, marshaled)
	require.Equal(t, report, unmarshaled)
}
