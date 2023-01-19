package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveGrant() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		granter    string
		grantee    types.Grantee
		allowance  feegrant.FeeAllowanceI
		check      func(ctx sdk.Context)
	}{
		{
			name:       "user grant is saved properly",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			allowance:  &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)

				// check if grant is set properly
				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
					grant)

				// check if account is set properly
				addr := sdk.MustAccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(suite.ak.HasAccount(ctx, addr))
			},
		},
		{
			name: "existing user grant is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}))
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			allowance:  &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)

				// check if grant is set properly
				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
					grant)

				// check if account is set properly
				addr := sdk.MustAccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(suite.ak.HasAccount(ctx, addr))
			},
		},
		{
			name:       "group grant is saved properly",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    types.NewGroupGrantee(1),
			allowance:  &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// check if grant is set properly
				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1), &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}),
					grant)
			},
		},
		{
			name: "existing group grant is overridden properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}))
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			grantee:    types.NewGroupGrantee(1),
			allowance:  &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)

				// check if grant is set properly
				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}), grant)
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
			suite.k.SaveGrant(ctx, types.NewGrant(tc.subspaceID,
				tc.granter,
				tc.grantee,
				tc.allowance))

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
		expResult  bool
	}{
		{
			name:       "grant does not exist returns false",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			expResult:  false,
		},
		{
			name: "existing grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
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
			suite.Require().Equal(tc.expResult, suite.k.HasUserGrant(ctx, tc.subspaceID, tc.grantee))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserGrant() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		grantee     string
		shouldFound bool
		expResult   types.Grant
	}{
		{
			name:        "non-existing grant returns error",
			subspaceID:  1,
			grantee:     "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			shouldFound: false,
		},
		{
			name: "existing grant returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}))
			},
			subspaceID:  1,
			grantee:     "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			shouldFound: true,
			expResult: types.NewGrant(1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}),
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

			suite.Require().Equal(tc.shouldFound, found)
			suite.Require().Equal(tc.expResult, grant)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteUserGrant() {
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
			},
		},
		{
			name: "existing grant is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
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
		expResult  bool
	}{
		{
			name:       "grant does not exist returns false",
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
			expResult:  false,
		},
		{
			name: "existing grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1), &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}))
			},
			subspaceID: 1,
			granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:    1,
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
			suite.Require().Equal(tc.expResult, suite.k.HasGroupGrant(ctx, tc.subspaceID, tc.groupID))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetGroupGrant() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		granter     string
		groupID     uint32
		shouldFound bool
		expResult   types.Grant
	}{
		{
			name:        "non-existing grant returns error",
			subspaceID:  1,
			granter:     "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:     1,
			shouldFound: false,
		},
		{
			name: "existing grant returns no error",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}))
			},
			subspaceID:  1,
			granter:     "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			groupID:     1,
			shouldFound: true,
			expResult: types.NewGrant(
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

			suite.Require().Equal(tc.shouldFound, found)
			suite.Require().Equal(tc.expResult, grant)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteGroupGrant() {
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
			},
		},
		{
			name: "existing grant is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1)))}))
			},
			subspaceID: 1,
			groupID:    1,
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

			suite.k.DeleteGroupGrant(ctx, tc.subspaceID, tc.groupID)

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
		expResult  bool
	}{
		{
			name:       "no any grant exists returns false",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expResult:  false,
		},
		{
			name: "invalid user grant returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(101))),
			expResult:  false,
		},
		{
			name: "valid user grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(10))),
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)
				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(90)))}),
					grant)
			},
			expResult: true,
		},
		{
			name: "use up user grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasUserGrant(ctx, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))
			},
			expResult: true,
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

			suite.Require().Equal(tc.expResult, suite.k.UseUserGrantedFees(ctx, tc.subspaceID, granteeAddr, tc.fees, nil))

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
		expResult  bool
	}{
		{
			name:       "no any grant exists returns false",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expResult:  false,
		},
		{
			name: "user not in the granted group returns false",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expResult:  false,
		},
		{
			name: "invalid grant returns false",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(101))),
			expResult:  false,
		},
		{
			name: "valid grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(10))),
			check: func(ctx sdk.Context) {
				grant, found := suite.k.GetGroupGrant(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(90)))}),
					grant)
			},
			expResult: true,
		},
		{
			name: "use up grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1), &feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasGroupGrant(ctx, 1, 1))
			},
			expResult: true,
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

			suite.Require().Equal(tc.expResult, suite.k.UseGroupGrantedFees(ctx, tc.subspaceID, granteeAddr, tc.fees, nil))

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
		expResult  bool
	}{
		{
			name:       "no any grant exists returns false",
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
			expResult:  false,
		},
		{
			name: "valid user grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(10))),
			expResult:  true,
		},
		{
			name: "valid group grant returns true",
			store: func(ctx sdk.Context) {
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))}))
			},
			subspaceID: 1,
			grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			fees:       sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(10))),
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

			granteeAddr, err := sdk.AccAddressFromBech32(tc.grantee)
			suite.Require().NoError(err)

			suite.Require().Equal(tc.expResult, suite.k.UseGrantedFees(ctx, tc.subspaceID, granteeAddr, tc.fees, nil))
		})
	}
}
