package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	posts "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/desmos-labs/desmos/x/reports/types/models"
)

func (suite *KeeperTestSuite) TestInvariants() {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
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
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			// nolint: errcheck
			suite.keeper.SaveReport(suite.ctx, test.postID, test.report)

			res, stop := keeper.AllInvariants(suite.keeper)(suite.ctx)

			suite.Equal(test.expResponse, res)
			suite.Equal(test.expBool, stop)
		})
	}
}
