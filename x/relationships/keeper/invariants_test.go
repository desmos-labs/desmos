package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		expResponse string
		expBroken   bool
	}{
		{
			name:        "empty state does not break invariants",
			expResponse: "Every invariant condition is fulfilled correctly",
			expBroken:   false,
		},
		{
			name: "ValidUserBlocksInvariant broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(
					types.UserBlockStoreKey("user", "user", 1),
					types.MustMarshalUserBlock(suite.cdc, types.NewUserBlock("user", "user", "reason", 1)),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid user blocks",
				fmt.Sprintf("%s%s",
					"The following list contains invalid user blocks:\n",
					"[Blocker]: user, [Blocked]: user, [SubspaceID]: 1\n",
				),
			),
		},
		{
			name: "ValidRelationshipsInvariant broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(
					types.RelationshipsStoreKey("user", "user", 1),
					types.MustMarshalRelationship(suite.cdc, types.NewRelationship("user", "user", 1)),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid relationships",
				fmt.Sprintf("%s%s",
					"The following list contains invalid relationships:\n",
					"[Creator]: user, [Counterparty]: user, [SubspaceID]: 1\n",
				),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, broken := keeper.AllInvariants(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
			suite.Require().Equal(tc.expResponse, res)
		})
	}
}
