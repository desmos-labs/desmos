package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"

	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveGrant() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		grant types.Grant
		check func(ctx sdk.Context)
	}{
		{
			name: "user grant is saved properly",
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)

				// Check if grant is set properly
				suite.Require().Equal(types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				), grant)

				// Check if account is set properly
				addr := sdk.MustAccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(suite.ak.HasAccount(ctx, addr))
			},
		},
		{
			name: "existing user grant is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))},
				))
			},
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)

				// Check if grant is set properly
				suite.Require().Equal(types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				), grant)

				// Check if account is set properly
				addr := sdk.MustAccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(suite.ak.HasAccount(ctx, addr))
			},
		},
		{
			name: "group grant is saved properly",
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// Check if grant is set properly
				suite.Require().Equal(types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				), grant)
			},
		},
		{
			name: "existing group grant is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))},
				))
			},
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			),
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// Check if grant is set properly
				suite.Require().Equal(types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				), grant)
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

			suite.k.SaveGrant(ctx, tc.grant)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveAllowanceToExpirationQueue() {
	expiration := time.Date(2100, 7, 7, 0, 0, 0, 0, time.UTC)
	newExpiration := time.Date(2100, 7, 14, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		grant types.Grant
		check func(ctx sdk.Context)
	}{
		{
			name: "new grant without expiration is saved properly",
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{
					SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
				},
			),
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// Check no grants inside expiration queue
				suite.Require().Empty(suite.k.GetAllGrantsInExpirationQueue(ctx))
			},
		},
		{
			name: "new grant with expiration is saved properly properly",
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{
					SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
					Expiration: &newExpiration,
				},
			),
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// Check grant is added into grant queue
				suite.Require().True(
					ctx.KVStore(suite.storeKey).Has(
						types.AllowanceExpirationQueueKey(&newExpiration, types.GroupAllowanceKey(1, 1)),
					),
				)
			},
		},
		{
			name: "grant without expiration overrides the grant with expiration properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
						Expiration: &expiration,
					},
				))
			},
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{
					SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
				},
			),
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// Check no grants inside expiration queue
				suite.Require().Empty(suite.k.GetAllGrantsInExpirationQueue(ctx))
			},
		},
		{
			name: "grant with expiration overrides the grant with expiration properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
						Expiration: &expiration,
					},
				))
			},
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{
					SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
					Expiration: &newExpiration,
				},
			),
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// Check old expiration is removed properly
				suite.Require().False(
					ctx.KVStore(suite.storeKey).Has(
						types.AllowanceExpirationQueueKey(&expiration, types.GroupAllowanceKey(1, 1)),
					),
				)

				// Check new expiration is added properly
				suite.Require().True(
					ctx.KVStore(suite.storeKey).Has(
						types.AllowanceExpirationQueueKey(&newExpiration, types.GroupAllowanceKey(1, 1)),
					),
				)
			},
		},
		{
			name: "grant without expiration overrides the grant without expiration properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
					},
				))
			},
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{
					SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
				},
			),
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// Check no grants inside expiration queue
				suite.Require().Empty(suite.k.GetAllGrantsInExpirationQueue(ctx))
			},
		},
		{
			name: "grant with expiration overrides the grant without expiration properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
					},
				))
			},
			grant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{
					SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
					Expiration: &newExpiration,
				},
			),
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// Check new expiration is added properly
				suite.Require().True(
					ctx.KVStore(suite.storeKey).Has(
						types.AllowanceExpirationQueueKey(&newExpiration, types.GroupAllowanceKey(1, 1)),
					),
				)
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

			suite.k.SaveGrant(ctx, tc.grant)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_HasUserGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		grantee    string
		expFound   bool
	}{
		{
			name:       "grant does not exist returns false",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			expFound:   false,
		},
		{
			name: "existing grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			expFound:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasUserGrant(ctx, tc.subspaceID, tc.grantee)
			suite.Require().Equal(tc.expFound, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		grantee    string
		expFound   bool
		expGrant   types.Grant
	}{
		{
			name:       "non-existing grant returns properly",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			expFound:   false,
		},
		{
			name: "existing grant returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			expFound:   true,
			expGrant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))},
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

			grant, found := suite.k.GetUserGrant(ctx, tc.subspaceID, tc.grantee)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expGrant, grant)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteUserGrant() {
	expiration := time.Date(2100, 7, 7, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		grantee    string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing grant is deleted properly",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))

				// Check no grants inside expiration queue
				suite.Require().Empty(suite.k.GetAllGrantsInExpirationQueue(ctx))
			},
		},
		{
			name: "existing grant is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1))),
						Expiration: &expiration,
					},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))

				// Check no grants inside expiration queue
				suite.Require().Empty(suite.k.GetAllGrantsInExpirationQueue(ctx))
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

			suite.k.DeleteUserGrant(ctx, tc.subspaceID, tc.grantee)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_HasGroupGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		groupID    uint32
		expFound   bool
	}{
		{
			name:       "grant does not exist returns false",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			expFound:   false,
		},
		{
			name: "existing grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))},
				))
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			expFound:   true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasGroupGrant(ctx, tc.subspaceID, tc.groupID)
			suite.Require().Equal(tc.expFound, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetGroupGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		groupID    uint32
		expFound   bool
		expGrant   types.Grant
	}{
		{
			name:       "non-existing grant returns properly",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			expFound:   false,
		},
		{
			name: "existing grant returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))},
				))
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			expFound:   true,
			expGrant: types.NewGrant(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewGroupGrantee(1),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))},
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

			grant, found := suite.k.GetGroupGrant(ctx, tc.subspaceID, tc.groupID)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().Equal(tc.expGrant, grant)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteGroupGrant() {
	expiration := time.Date(2100, 7, 7, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing grant is deleted properly",
			subspaceID: 1,
			groupID:    1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, 1))

				// Check no grants inside expiration queue
				suite.Require().Empty(suite.k.GetAllGrantsInExpirationQueue(ctx))
			},
		},
		{
			name: "existing grant is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1))),
						Expiration: &expiration,
					},
				))
			},
			subspaceID: 1,
			groupID:    1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, 1))

				// Check no grants inside expiration queue
				suite.Require().Empty(suite.k.GetAllGrantsInExpirationQueue(ctx))
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

			suite.k.DeleteGroupGrant(ctx, tc.subspaceID, tc.groupID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_RemoveExpiredAllowances() {
	expiration := time.Date(2100, 7, 7, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		expiration time.Time
		check      func(ctx sdk.Context)
	}{
		{
			name: "non expired grant is kept properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1))),
						Expiration: &expiration,
					},
				))
			},
			expiration: time.Date(2100, 7, 6, 0, 0, 0, 0, time.UTC),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasGroupGrant(ctx, 1, 1))

				// Check new expiration is kept properly
				suite.Require().True(
					ctx.KVStore(suite.storeKey).Has(
						types.AllowanceExpirationQueueKey(&expiration, types.GroupAllowanceKey(1, 1)),
					),
				)
			},
		},
		{
			name: "expired grant is removed properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1))),
						Expiration: &expiration,
					},
				))
			},
			expiration: time.Date(2100, 7, 8, 0, 0, 0, 0, time.UTC),
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, 1))

				// Check new expiration is removed properly
				suite.Require().False(
					ctx.KVStore(suite.storeKey).Has(
						types.AllowanceExpirationQueueKey(&expiration, types.GroupAllowanceKey(1, 1)),
					),
				)
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

			suite.k.RemoveExpiredAllowances(ctx, tc.expiration)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestSuite) TestKeeper_UseUserGrantedFees() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		grantee    string
		fees       sdk.Coins
		check      func(ctx sdk.Context)
		expUsed    bool
	}{
		{
			name:       "not existing grant returns false",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expUsed:    false,
		},
		{
			name: "invalid user grant returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(101))),
			expUsed:    false,
		},
		{
			name: "valid user grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(10))),
			expUsed:    true,
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)
				suite.Require().Equal(types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(90)))},
				), grant)
			},
		},
		{
			name: "consuming grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expUsed:    true,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))
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

			granteeAddr, err := sdk.AccAddressFromBech32(tc.grantee)
			suite.Require().NoError(err)

			used := suite.k.UseUserGrantedFees(ctx, tc.subspaceID, granteeAddr, tc.fees, nil)
			suite.Require().Equal(tc.expUsed, used)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_UseGroupGrantedFees() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		grantee    string
		fees       sdk.Coins
		check      func(ctx sdk.Context)
		expUsed    bool
	}{
		{
			name:       "not existing grant returns false",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expUsed:    false,
		},
		{
			name: "user not in the granted group returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expUsed:    false,
		},
		{
			name: "invalid grant returns false",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(101))),
			expUsed:    false,
		},
		{
			name: "valid grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(10))),
			expUsed:    true,
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(90)))},
				), grant)
			},
		},
		{
			name: "consuming grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expUsed:    true,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, 1))
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

			granteeAddr, err := sdk.AccAddressFromBech32(tc.grantee)
			suite.Require().NoError(err)

			used := suite.k.UseGroupGrantedFees(ctx, tc.subspaceID, granteeAddr, tc.fees, nil)
			suite.Require().Equal(tc.expUsed, used)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_UseGrantedFees() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		grantee    string
		fees       sdk.Coins
		expUsed    bool
	}{
		{
			name:       "not existing grant returns false",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expUsed:    false,
		},
		{
			name: "valid user grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(10))),
			expUsed:    true,
		},
		{
			name: "valid group grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(10))),
			expUsed:    true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			granteeAddr, err := sdk.AccAddressFromBech32(tc.grantee)
			suite.Require().NoError(err)

			used := suite.k.UseGrantedFees(ctx, tc.subspaceID, granteeAddr, tc.fees, nil)
			suite.Require().Equal(tc.expUsed, used)
		})
	}
}
