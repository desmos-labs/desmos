package keeper_test

import "github.com/desmos-labs/desmos/x/staging/reports/types"

func (suite *KeeperTestSuite) TestExportGenesis() {
	tests := []struct {
		name  string
		state struct {
			reports []types.Report
		}
		expGenesis *types.GenesisState
	}{
		{
			name: "empty data is exported correctly",
			state: struct {
				reports []types.Report
			}{
				reports: []types.Report{},
			},
			expGenesis: types.NewGenesisState(nil),
		},
		{
			name: "data is exported correctly",
			state: struct {
				reports []types.Report
			}{
				reports: []types.Report{
					types.NewReport("post_id", "types", "message", "user"),
					types.NewReport("post_id", "types", "message_2", "user"),
				},
			},
			expGenesis: types.NewGenesisState(
				[]types.Report{
					types.NewReport("post_id", "types", "message", "user"),
					types.NewReport("post_id", "types", "message_2", "user"),
				},
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, report := range test.state.reports {
				suite.k.SaveReport(suite.ctx, report)
			}

			genesis := suite.k.ExportGenesis(suite.ctx)
			suite.Require().Equal(test.expGenesis, genesis)
		})
	}
}

func (suite *KeeperTestSuite) TestInitGenesis() {
	tests := []struct {
		name     string
		genesis  *types.GenesisState
		expPanic bool
		expState struct {
			reports []types.Report
		}
	}{
		{
			name:    "empty genesis is initialized properly",
			genesis: types.NewGenesisState(nil),
			expState: struct {
				reports []types.Report
			}{
				reports: nil,
			},
		},
		{
			name: "proper genesis is initialized properly",
			genesis: types.NewGenesisState([]types.Report{
				types.NewReport("post_id", "types", "message", "user"),
				types.NewReport("post_id", "types", "message_2", "user"),
			}),
			expState: struct {
				reports []types.Report
			}{
				reports: []types.Report{
					types.NewReport("post_id", "types", "message", "user"),
					types.NewReport("post_id", "types", "message_2", "user"),
				},
			},
		},
		{
			name: "double reports panics",
			genesis: types.NewGenesisState([]types.Report{
				types.NewReport("post_id", "type", "message", "user"),
				types.NewReport("post_id", "type", "message", "user"),
			}),
			expPanic: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.expPanic {
				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, *test.genesis) })
			} else {
				suite.k.InitGenesis(suite.ctx, *test.genesis)

				reports := suite.k.GetAllReports(suite.ctx)
				suite.Require().Equal(test.expState.reports, reports)
			}
		})
	}
}
