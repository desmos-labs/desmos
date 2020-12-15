package keeper_test

import (
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	tests := []struct {
		name       string
		report     types.Report
		shouldStop bool
	}{
		{
			name: "Invariants not violated",
			report: types.NewReport(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldStop: true,
		},
		{
			name: "ValidReportIDs invariant violated",
			report: types.NewReport(
				"123",
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldStop: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.keeper.SaveReport(suite.ctx, test.report)

			_, stop := keeper.AllInvariants(suite.keeper)(suite.ctx)
			suite.Require().Equal(test.shouldStop, stop)
		})
	}
}
