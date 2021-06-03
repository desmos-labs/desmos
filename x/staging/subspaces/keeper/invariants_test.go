package keeper_test

import (
	"github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"time"
)

func (suite *KeeperTestsuite) TestInvariants() {
	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.NoError(err)

	tests := []struct {
		name      string
		subspaces []types.Subspace
		expStop   bool
	}{
		{
			name: "All invariants are not violated",
			subspaces: []types.Subspace{
				{
					ID:           "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					Name:         "test",
					Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					CreationTime: date,
					Type:         types.Open,
				},
			},
			expStop: true,
		},
		{
			name: "Valid subspace invariant violated",
			subspaces: []types.Subspace{
				{
					ID:           "",
					Name:         "test",
					Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					CreationTime: date,
				},
			},
			expStop: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, sub := range test.subspaces {
				_ = suite.k.SaveSubspace(suite.ctx, sub, sub.Owner)
			}

			_, stop := keeper.AllInvariants(suite.k)(suite.ctx)
			suite.Require().Equal(test.expStop, stop)
		})
	}
}
