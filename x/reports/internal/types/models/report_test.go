package models_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/stretchr/testify/require"
)

func TestReport_String(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	report := types.NewReport("scam", "it's a trap", creator)
	require.Equal(t, `{"type":"scam","message":"it's a trap","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`, report.String())
}

func TestReport_Validate(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name   string
		report types.Report
		expErr error
	}{
		{
			name:   "empty reports's type returns error",
			report: types.NewReport("", "message", creator),
			expErr: fmt.Errorf("report type cannot be empty"),
		},
		{
			name:   "empty reports's message returns error",
			report: types.NewReport("scam", "", creator),
			expErr: fmt.Errorf("reports's message cannot be empty"),
		},
		{
			name:   "invalid reports's creator returns error",
			report: types.NewReport("scam", "message", sdk.AccAddress{}),
			expErr: fmt.Errorf("invalid user address "),
		},
		{
			name:   "valid reports returns no error",
			report: types.NewReport("scam", "message", creator),
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

func TestReport_Equals(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	types.NewReport("type", "message", creator)

	tests := []struct {
		name     string
		report   types.Report
		otherRep types.Report
		expBool  bool
	}{
		{
			name:     "equals reports returns true",
			report:   types.NewReport("scam", "it's a trap", creator),
			otherRep: types.NewReport("scam", "it's a trap", creator),
			expBool:  true,
		},
		{
			name:     "non-equals reports returns false",
			report:   types.NewReport("scam", "it's a trap", creator),
			otherRep: types.NewReport("spam", "it's a trap", creator),
			expBool:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actualBool := test.report.Equals(test.otherRep)
			require.Equal(t, actualBool, test.expBool)
		})
	}
}

func TestReports_String(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	types.NewReport("type", "message", creator)

	reports := types.Reports{
		types.NewReport("scam", "message", creator),
		types.NewReport("violence", "message", creator),
	}

	require.Equal(t, "Type - Message - User\nscam - message - cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\nviolence - message - cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", reports.String())
}

func TestReports_Validate(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name    string
		reports types.Reports
		expErr  error
	}{
		{
			name:    "empty reports's type returns error",
			reports: types.Reports{types.NewReport("", "message", creator)},
			expErr:  fmt.Errorf("report type cannot be empty"),
		},
		{
			name:    "empty reports's message returns error",
			reports: types.Reports{types.NewReport("scam", "", creator)},
			expErr:  fmt.Errorf("reports's message cannot be empty"),
		},
		{
			name:    "invalid reports's creator returns error",
			reports: types.Reports{types.NewReport("scam", "message", sdk.AccAddress{})},
			expErr:  fmt.Errorf("invalid user address "),
		},
		{
			name:    "valid reports returns no error",
			reports: types.Reports{types.NewReport("scam", "message", creator)},
			expErr:  nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actualErr := test.reports.Validate()
			require.Equal(t, actualErr, test.expErr)
		})
	}
}

func TestReports_Equals(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	types.NewReport("type", "message", creator)

	tests := []struct {
		name      string
		reports   types.Reports
		otherReps types.Reports
		expBool   bool
	}{
		{
			name:      "equals reports returns true",
			reports:   types.Reports{types.NewReport("scam", "it's a trap", creator)},
			otherReps: types.Reports{types.NewReport("scam", "it's a trap", creator)},
			expBool:   true,
		},
		{
			name:      "non-equals reports returns false",
			reports:   types.Reports{types.NewReport("scam", "it's a trap", creator)},
			otherReps: types.Reports{types.NewReport("spam", "it's a trap", creator)},
			expBool:   false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actualBool := test.reports.Equals(test.otherReps)
			require.Equal(t, actualBool, test.expBool)
		})
	}
}
