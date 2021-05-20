package keeper_test

import (
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"time"
)

func (suite *KeeperTestsuite) TestKeeper_ExportGenesis() {
	tests := []struct {
		name string
		data struct {
			subspaces []types.Subspace
		}
		expected *types.GenesisState
	}{
		{
			name: "Default expected state",
			data: struct {
				subspaces []types.Subspace
			}{
				subspaces: nil,
			},
			expected: &types.GenesisState{Subspaces: nil},
		},
		{
			name: "Genesis exported successfully",
			data: struct {
				subspaces []types.Subspace
			}{
				subspaces: []types.Subspace{
					{
						ID:           "123",
						Name:         "test",
						Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						Open:         true,
						CreationTime: time.Time{},
					},
				},
			},
			expected: types.NewGenesisState([]types.Subspace{
				types.NewSubspace(
					"123",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					true,
					time.Time{},
				)},
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, subspace := range test.data.subspaces {
				suite.k.SaveSubspace(suite.ctx, subspace)
			}

			exported := suite.k.ExportGenesis(suite.ctx)
			suite.Equal(test.expected, exported)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_InitGenesis() {
	tests := []struct {
		name     string
		genesis  *types.GenesisState
		expError bool
		expState struct {
			subspaces []types.Subspace
		}
	}{
		{
			name:     "Default genesis is initialized properly",
			genesis:  types.DefaultGenesisState(),
			expError: false,
			expState: struct {
				subspaces []types.Subspace
			}{
				subspaces: nil,
			},
		},
		{
			name: "Invalid subspace panics",
			genesis: &types.GenesisState{
				Subspaces: []types.Subspace{
					{
						ID:           "",
						Name:         "test",
						Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						CreationTime: time.Time{},
					}},
			},
			expError: true,
		},
		{
			name: "Already existing subspace panics",
			genesis: &types.GenesisState{
				Subspaces: []types.Subspace{
					{
						ID:           "456",
						Name:         "test",
						Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						CreationTime: time.Time{},
					},
				},
			},
			expError: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.genesis.Subspaces != nil && test.genesis.Subspaces[0].ID == "456" {
				suite.k.SaveSubspace(suite.ctx, test.genesis.Subspaces[0])
			}

			if test.expError {
				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, *test.genesis) })
			} else {
				suite.k.InitGenesis(suite.ctx, *test.genesis)
				suite.Require().Equal(test.expState.subspaces, suite.k.GetAllSubspaces(suite.ctx))
			}
		})
	}
}
