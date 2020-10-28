package keeper_test

import (
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	tests := []struct {
		name        string
		report      types.Report
		expBool     bool
		expResponse string
	}{
		{
			name: "Invariants not violated",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expBool:     true,
			expResponse: "Every invariant condition is fulfilled correctly",
		},
		{
			name: "ValidReportIDs invariant violated",
			report: types.NewReport(
				"123",
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expBool:     true,
			expResponse: "stored: invalid stored' IDs invariant\nThe following list contains invalid postIDs:\n 123\n\n",
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			err := suite.keeper.SaveReport(suite.ctx, test.report)
			suite.Require().NoError(err)

			res, stop := keeper.AllInvariants(suite.keeper)(suite.ctx)

			suite.Require().Equal(test.expResponse, res)
			suite.Require().Equal(test.expBool, stop)
		})
	}
}
