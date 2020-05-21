package models_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReport_String(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	report := types.NewReport("scam", "it's a trap", creator)
	require.Equal(t, `{"type":"scam","message":"it's a trap","user":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`, report.String())
}

func TestReport_ValidReportType(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name    string
		reports types.Reports
		expBool bool
	}{
		{
			name: "Valid reports type",
			reports: types.Reports{
				types.NewReport("scam", "it's a trap", creator),
				types.NewReport("nudity", "it's a trap", creator),
				types.NewReport("violence", "it's a trap", creator),
				types.NewReport("intimidation", "it's a trap", creator),
				types.NewReport("suicide or self-harm", "it's a trap", creator),
				types.NewReport("fake news", "it's a trap", creator),
				types.NewReport("spam", "it's a trap", creator),
				types.NewReport("unauthorized sale", "it's a trap", creator),
				types.NewReport("hatred incitement", "it's a trap", creator),
				types.NewReport("promotion of drug use", "it's a trap", creator),
				types.NewReport("non-consensual intimate images", "it's a trap", creator),
				types.NewReport("pornography", "it's a trap", creator),
				types.NewReport("children abuse", "it's a trap", creator),
				types.NewReport("animals abuse", "it's a trap", creator),
				types.NewReport("bullying", "it's a trap", creator),
			},
			expBool: true,
		},
		{
			name: "Invalid reports type",
			reports: models.Reports{
				types.NewReport("type", "mess", creator),
			},
			expBool: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			for _, rep := range test.reports {
				require.Equal(t, test.expBool, rep.ValidReportType())
			}
		})
	}
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
			expErr: fmt.Errorf("invalid reports type, please referes to our official reports's type list to check the valid ones"),
		},
		{
			name:   "invalid reports's type returns error",
			report: types.NewReport("type", "message", creator),
			expErr: fmt.Errorf("invalid reports type, please referes to our official reports's type list to check the valid ones"),
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
			expErr:  fmt.Errorf("invalid reports type, please referes to our official reports's type list to check the valid ones"),
		},
		{
			name:    "invalid reports's type returns error",
			reports: types.Reports{types.NewReport("type", "message", creator)},
			expErr:  fmt.Errorf("invalid reports type, please referes to our official reports's type list to check the valid ones"),
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

func TestReports_AppendIfMissing(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	types.NewReport("type", "message", creator)

	tests := []struct {
		name       string
		reports    types.Reports
		report     types.Report
		expReports types.Reports
		expBool    bool
	}{
		{
			name:    "new reports has been appended correctly",
			reports: types.Reports{types.NewReport("scam", "it's a trap", creator)},
			report:  types.NewReport("spam", "it's a trap", creator),
			expReports: types.Reports{
				types.NewReport("scam", "it's a trap", creator),
				types.NewReport("spam", "it's a trap", creator),
			},
			expBool: true,
		},
		{
			name:       "already existent reports has not been appended",
			reports:    types.Reports{types.NewReport("scam", "it's a trap", creator)},
			report:     types.NewReport("scam", "it's a trap", creator),
			expReports: types.Reports{types.NewReport("scam", "it's a trap", creator)},
			expBool:    false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actualReps, actualBool := test.reports.AppendIfMissing(test.report)
			require.Equal(t, actualReps, test.expReports)
			require.Equal(t, actualBool, test.expBool)
		})
	}
}
