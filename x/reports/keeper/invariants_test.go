package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/desmos-labs/desmos/x/reports/types/models"
	"github.com/stretchr/testify/require"
)

func TestInvariants(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	postID := posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	report := models.NewReport("type", "message", creator)

	tests := []struct {
		name        string
		postID      posts.PostID
		report      types.Report
		expBool     bool
		expResponse string
	}{
		{
			name:        "Invariants not violated",
			postID:      postID,
			report:      report,
			expBool:     true,
			expResponse: "Every invariant condition is fulfilled correctly",
		},
		{
			name:        "ValidReportIDs invariant violated",
			postID:      "123",
			report:      report,
			expBool:     true,
			expResponse: "reports: invalid reports' IDs invariant\nThe following list contains invalid postIDs:\n 123\n\n",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k, _ := SetupTestInput()
			// nolint: errcheck
			k.SaveReport(ctx, test.postID, test.report)

			res, stop := keeper.AllInvariants(k)(ctx)

			require.Equal(t, test.expResponse, res)
			require.Equal(t, test.expBool, stop)
		})
	}
}
