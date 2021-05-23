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
	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.NoError(err)

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
			name: "Valid genesis initialized correctly",
			genesis: &types.GenesisState{
				Subspaces: []types.Subspace{
					{
						ID:           "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						Name:         "test",
						Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						CreationTime: date,
					},
				},
			},
			expState: struct {
				subspaces []types.Subspace
			}{
				subspaces: []types.Subspace{
					{
						ID:           "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						Name:         "test",
						Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						CreationTime: date,
					},
				},
			},
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.expError {
				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, *test.genesis) })
			} else {
				suite.k.InitGenesis(suite.ctx, *test.genesis)
				suite.Require().Equal(test.expState.subspaces, suite.k.GetAllSubspaces(suite.ctx))
			}
		})
	}
}
