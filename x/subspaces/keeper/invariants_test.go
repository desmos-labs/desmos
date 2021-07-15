package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestInvariants() {
	tests := []struct {
		name    string
		store   func(ctx sdk.Context)
		expStop bool
	}{
		{
			name: "All invariants are not violated",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2050, 01, 01, 15, 15, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			expStop: true,
		},
		{
			name: "Valid subspace invariant violated",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2050, 01, 01, 15, 15, 00, 000, time.UTC),
				)
				_ = suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
			},
			expStop: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			_, stop := keeper.AllInvariants(suite.k)(suite.ctx)
			suite.Require().Equal(test.expStop, stop)
		})
	}
}
