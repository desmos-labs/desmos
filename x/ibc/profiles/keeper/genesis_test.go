package keeper_test

import "github.com/desmos-labs/desmos/x/ibc/profiles/types"

func (suite *KeeperTestSuite) TestExportGenesis() {
	tests := []struct {
		name       string
		state      types.GenesisState
		expGenesis types.GenesisState
	}{
		{
			name: "Data is exported correctly",
			state: *types.NewGenesisState(
				types.PortID,
			),
			expGenesis: *types.NewGenesisState(
				types.PortID,
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.k.SetPort(suite.ctx, types.PortID)

			genesis := suite.k.ExportGenesis(suite.ctx)
			suite.Require().Equal(test.expGenesis, *genesis)
		})
	}
}

func (suite *KeeperTestSuite) TestInitGenesis() {
	tests := []struct {
		name     string
		genesis  types.GenesisState
		expPanic bool
		expState types.GenesisState
	}{
		{
			name: "proper genesis is initialized properly",
			genesis: *types.NewGenesisState(
				types.PortID),
			expPanic: false,
			expState: *types.NewGenesisState(
				types.PortID),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.expPanic {
				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, test.genesis) })
			} else {
				suite.k.InitGenesis(suite.ctx, test.genesis)
				suite.Require().Equal(types.PortID, suite.k.GetPort(suite.ctx))
			}
		})
	}
}
