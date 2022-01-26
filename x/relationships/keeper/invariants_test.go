package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v2/x/relationships/types"
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

				block := types.NewUserBlock("blocker", "blocked", "reason", 0)
				store.Set(
					types.UserBlockStoreKey(block.Blocker, block.SubspaceID, block.Blocked),
					suite.cdc.MustMarshal(&block),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid user blocks",
				fmt.Sprintf("%s%s",
					"The following list contains invalid user blocks:\n",
					"[Blocker]: blocker, [Blocked]: blocked, [SubspaceID]: 0\n",
				),
			),
		},
		{
			name: "ValidRelationshipsInvariant broken",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				relationship := types.NewRelationship("creator", "recipient", 0)
				store.Set(
					types.RelationshipsStoreKey(relationship.Creator, relationship.Counterparty, relationship.SubspaceID),
					suite.cdc.MustMarshal(&relationship),
				)
			},
			expBroken: true,
			expResponse: sdk.FormatInvariant(types.ModuleName, "invalid relationships",
				fmt.Sprintf("%s%s",
					"The following list contains invalid relationships:\n",
					"[Creator]: creator, [Counterparty]: recipient, [SubspaceID]: 0\n",
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
