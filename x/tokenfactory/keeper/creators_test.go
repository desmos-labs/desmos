package keeper_test

import sdk "github.com/cosmos/cosmos-sdk/types"

func (suite *KeeperTestSuite) TestKeeper_AddDenomFromCreator() {
	testCases := []struct {
		name    string
		creator string
		denom   string
		check   func(ctx sdk.Context)
	}{
		{
			name:    "add denom creator properly",
			creator: "creator",
			denom:   "denom",
			check: func(ctx sdk.Context) {
				store := suite.k.GetCreatorPrefixStore(ctx, "creator")
				suite.Require().Equal(
					[]byte("denom"),
					store.Get([]byte("denom")),
				)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			suite.k.AddDenomFromCreator(ctx, tc.creator, tc.denom)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDenomsFromCreator() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		creator   string
		denom     string
		expResult []string
	}{
		{
			name: "get denoms from creator properly",
			store: func(ctx sdk.Context) {
				suite.k.AddDenomFromCreator(ctx, "creator", "bitcoin")
				suite.k.AddDenomFromCreator(ctx, "creator", "litecoin")
			},
			creator:   "creator",
			expResult: []string{"bitcoin", "litecoin"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.Require().Equal(
				tc.expResult,
				suite.k.GetDenomsFromCreator(ctx, tc.creator),
			)
		})
	}
}
