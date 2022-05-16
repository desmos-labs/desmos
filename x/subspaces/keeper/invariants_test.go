package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestsuite) TestInvariants() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "valid subspace invariant violated",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: true,
		},
		{
			name: "valid user groups invariant violated - invalid data",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					0,
					"This is a test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			expBroken: true,
		},
		{
			name: "valid user groups invariant violated - missing associated subspace",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"This is a test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			expBroken: true,
		},
		{
			name: "no invariant is violated",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			expBroken: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if test.store != nil {
				test.store(ctx)
			}

			_, broken := keeper.AllInvariants(suite.k)(ctx)
			suite.Require().Equal(test.expBroken, broken)
		})
	}
}
