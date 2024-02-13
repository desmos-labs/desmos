package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/reactions/types"
)

func (suite *KeeperTestSuite) TestKeeper_HasReacted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		user       string
		subspaceID uint64
		postID     uint64
		value      types.ReactionValue
		expResult  bool
	}{
		{
			name: "different subspace returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			subspaceID: 2,
			postID:     1,
			user:       "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			value:      types.NewRegisteredReactionValue(1),
			expResult:  false,
		},
		{
			name: "different post id returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			subspaceID: 1,
			postID:     2,
			user:       "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			value:      types.NewRegisteredReactionValue(1),
			expResult:  false,
		},
		{
			name: "different user returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			subspaceID: 1,
			postID:     1,
			user:       "cosmos1qds55pz0scvw8l6fm44jxq57wwuk4hcun6rqhd",
			value:      types.NewRegisteredReactionValue(1),
			expResult:  false,
		},
		{
			name: "different value type returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			subspaceID: 1,
			postID:     1,
			user:       "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			value:      types.NewFreeTextValue("Wow"),
			expResult:  false,
		},
		{
			name: "different value inner value returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			subspaceID: 1,
			postID:     1,
			user:       "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			value:      types.NewRegisteredReactionValue(2),
			expResult:  false,
		},
		{
			name: "correct data returns true -- registered reaction value",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			subspaceID: 1,
			postID:     1,
			user:       "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			value:      types.NewRegisteredReactionValue(1),
			expResult:  true,
		},
		{
			name: "correct data returns true -- free text value",
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewFreeTextValue("test"),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			subspaceID: 1,
			postID:     1,
			user:       "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			value:      types.NewFreeTextValue("test"),
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasReacted(ctx, tc.subspaceID, tc.postID, tc.user, tc.value)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}
