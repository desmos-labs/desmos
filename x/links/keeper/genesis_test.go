package keeper_test

// func (suite *KeeperTestSuite) TestExportGenesis() {
// 	tests := []struct {
// 		name       string
// 		state      *types.GenesisState
// 		expGenesis types.GenesisState
// 	}{
// 		{
// 			name:       "empty data is exported correctly",
// 			state:      types.NewGenesisState(types.PortID, nil),
// 			expGenesis: *types.NewGenesisState(types.PortID, nil),
// 		},
// 		{
// 			name: "data is exported correctly",
// 			state: types.NewGenesisState(
// 				types.PortID,
// 				[]types.Link{
// 					suite.testData.link,
// 				},
// 			),
// 			expGenesis: *types.NewGenesisState(
// 				types.PortID,
// 				[]types.Link{
// 					suite.testData.link,
// 				},
// 			),
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		suite.Run(test.name, func() {
// 			suite.SetupTest()
// 			suite.k.SetPort(suite.ctx, types.PortID)
// 			for _, link := range test.state.Links {
// 				suite.k.StoreLink(suite.ctx, link)
// 			}

// 			genesis := suite.k.ExportGenesis(suite.ctx)
// 			suite.Require().Equal(test.expGenesis, genesis)
// 		})
// 	}
// }

// func (suite *KeeperTestSuite) TestInitGenesis() {
// 	tests := []struct {
// 		name     string
// 		genesis  *types.GenesisState
// 		expPanic bool
// 		expState types.GenesisState
// 	}{
// 		{
// 			name:     "empty genesis is initialized properly",
// 			genesis:  types.NewGenesisState(types.PortID, nil),
// 			expPanic: false,
// 			expState: types.GenesisState{
// 				PortId: types.PortID,
// 				Links:  nil,
// 			},
// 		},
// 		{
// 			name: "proper genesis is initialized properly",
// 			genesis: types.NewGenesisState(
// 				types.PortID, []types.Link{
// 					suite.testData.link,
// 				}),
// 			expPanic: false,
// 			expState: *types.NewGenesisState(
// 				types.PortID, []types.Link{
// 					suite.testData.link,
// 				}),
// 		},
// 	}

// 	for _, test := range tests {
// 		test := test
// 		suite.Run(test.name, func() {
// 			suite.SetupTest()

// 			if test.expPanic {
// 				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, *test.genesis) })
// 			} else {
// 				suite.k.InitGenesis(suite.ctx, *test.genesis)
// 				suite.Require().Equal(types.PortID, suite.k.GetPort(suite.ctx))
// 				links := suite.k.GetAllLinks(suite.ctx)
// 				suite.Require().Equal(test.expState.Links, links)
// 			}
// 		})
// 	}
// }
