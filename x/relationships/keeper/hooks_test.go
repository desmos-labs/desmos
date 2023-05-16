package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/relationships/types"
)

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name: "relationships are deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1wvpwlsvl7annmuasnzgftxstv5apjh8ed9zzc8",
					"cosmos18kn8l96p3zjgdg5v4udyevcs0hvqsl4pj3vlfj",
					1,
				))
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1wvpwlsvl7annmuasnzgftxstv5apjh8ed9zzc8",
					"cosmos18kn8l96p3zjgdg5v4udyevcs0hvqsl4pj3vlfj",
					2,
				))
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasRelationship(ctx,
					"cosmos1wvpwlsvl7annmuasnzgftxstv5apjh8ed9zzc8",
					"cosmos18kn8l96p3zjgdg5v4udyevcs0hvqsl4pj3vlfj",
					1,
				))
				suite.Require().True(suite.k.HasRelationship(ctx,
					"cosmos1wvpwlsvl7annmuasnzgftxstv5apjh8ed9zzc8",
					"cosmos18kn8l96p3zjgdg5v4udyevcs0hvqsl4pj3vlfj",
					2,
				))
			},
		},
		{
			name: "user blocks are deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1wvpwlsvl7annmuasnzgftxstv5apjh8ed9zzc8",
					"cosmos18kn8l96p3zjgdg5v4udyevcs0hvqsl4pj3vlfj",
					"",
					1,
				))
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1wvpwlsvl7annmuasnzgftxstv5apjh8ed9zzc8",
					"cosmos18kn8l96p3zjgdg5v4udyevcs0hvqsl4pj3vlfj",
					"",
					2,
				))
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserBlocked(ctx,
					"cosmos1wvpwlsvl7annmuasnzgftxstv5apjh8ed9zzc8",
					"cosmos18kn8l96p3zjgdg5v4udyevcs0hvqsl4pj3vlfj",
					1,
				))
				suite.Require().True(suite.k.HasUserBlocked(ctx,
					"cosmos1wvpwlsvl7annmuasnzgftxstv5apjh8ed9zzc8",
					"cosmos18kn8l96p3zjgdg5v4udyevcs0hvqsl4pj3vlfj",
					2,
				))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := suite.k.Hooks()
			hooks.AfterSubspaceDeleted(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
