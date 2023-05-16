package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/x/reports/types"
)

func TestReport_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		report    types.Report
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			report: types.NewReport(
				0,
				1,
				[]uint32{1},
				"",
				types.NewPostTarget(1),
				"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid id returns error",
			report: types.NewReport(
				1,
				0,
				[]uint32{1},
				"",
				types.NewPostTarget(1),
				"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "empty reasons ids returns error",
			report: types.NewReport(
				1,
				1,
				nil,
				"",
				types.NewPostTarget(1),
				"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid reason id returns error",
			report: types.NewReport(
				1,
				1,
				[]uint32{0},
				"",
				types.NewPostTarget(1),
				"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid reporter returns error",
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"",
				types.NewPostTarget(1),
				"",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid target returns error",
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"",
				types.NewPostTarget(0),
				"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid time returns error",
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"",
				types.NewPostTarget(1),
				"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
				time.Time{},
			),
			shouldErr: true,
		},
		{
			name: "valid report returns no error",
			report: types.NewReport(
				1,
				1,
				[]uint32{1},
				"",
				types.NewPostTarget(1),
				"cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.report.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

// --------------------------------------------------------------------------------------------------------------------

func TestUserTarget_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		data      *types.UserTarget
		shouldErr bool
	}{
		{
			name:      "invalid user address returns error",
			data:      types.NewUserTarget(""),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			data:      types.NewUserTarget("cosmos1atdl3cpms89md5qa3rxtql0drtgftch2zgkr7v"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.data.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestPostTarget_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		data      *types.PostTarget
		shouldErr bool
	}{
		{
			name:      "invalid post id returns error",
			data:      types.NewPostTarget(0),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			data:      types.NewPostTarget(1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.data.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

// --------------------------------------------------------------------------------------------------------------------

func TestReason_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		reason    types.Reason
		shouldErr bool
	}{
		{
			name:      "invali subspace id returns error",
			reason:    types.NewReason(0, 1, "Spam", "This content is spam"),
			shouldErr: true,
		},
		{
			name:      "invalid id returns error",
			reason:    types.NewReason(1, 0, "Spam", "This content is spam"),
			shouldErr: true,
		},
		{
			name:      "invalid title returns error",
			reason:    types.NewReason(1, 1, "", "This content is spam"),
			shouldErr: true,
		},
		{
			name:      "valid reason returns no error",
			reason:    types.NewReason(1, 1, "Spam", "This content is spam"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.reason.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
