package keeper_test

import "github.com/desmos-labs/desmos/x/links/types"

func (suite *KeeperTestSuite) TestExportGenesis() {
	tests := []struct {
		name       string
		state      *types.GenesisState
		expGenesis types.GenesisState
	}{
		{
			name:       "empty data is exported correctly",
			state:      types.NewGenesisState(types.PortID, nil),
			expGenesis: *types.NewGenesisState(types.PortID, nil),
		},
		{
			name: "data is exported correctly",
			state: types.NewGenesisState(
				types.PortID,
				[]types.Link{
					types.NewLink("cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn", "cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn"),
				},
			),
			expGenesis: *types.NewGenesisState(
				types.PortID,
				[]types.Link{
					types.NewLink("cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn", "cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn"),
				},
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetPort(suite.ctx, types.PortID)
			for _, link := range test.state.Links {
				suite.k.StoreLink(suite.ctx, link)
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
		expState types.GenesisState
	}{
		{
			name:    "empty genesis is initialized properly",
			genesis: types.NewGenesisState(types.PortID, nil),
			expState: types.GenesisState{
				PortId: types.PortID,
				Links:  nil,
			},
		},
		{
			name: "proper genesis is initialized properly",
			genesis: types.NewGenesisState(types.PortID, []types.Link{
				types.NewLink("cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn", "cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn"),
			}),
			expState: types.GenesisState{
				PortId: types.PortID,
				Links: []types.Link{
					types.NewLink("cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn", "cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn"),
				},
			},
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

				links := suite.k.GetAllLinks(suite.ctx)
				suite.Require().Equal(test.expState.Links, links)
			}
		})
	}
}
