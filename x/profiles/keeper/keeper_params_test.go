package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	params := types.NewParams(
		types.NewNicknameParams(sdk.NewInt(3), sdk.NewInt(1000)),
		types.NewDTagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(1000)),
		sdk.NewInt(1000),
	)
	suite.k.SetParams(suite.ctx, params)

	actualParams := suite.k.GetParams(suite.ctx)
	suite.Require().Equal(params, actualParams)
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		expParams types.Params
	}{
		{
			name: "valid params do not error",
			store: func(ctx sdk.Context) {
				params := types.NewParams(
					types.NewNicknameParams(sdk.NewInt(3), sdk.NewInt(1000)),
					types.NewDTagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(1000)),
					sdk.NewInt(1000),
				)
				suite.k.SetParams(ctx, params)
			},
			shouldErr: false,
			expParams: types.NewParams(
				types.NewNicknameParams(sdk.NewInt(3), sdk.NewInt(1000)),
				types.NewDTagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(1000)),
				sdk.NewInt(1000),
			),
		},
		{
			name:      "invalid params panics",
			shouldErr: true,
			expParams: types.Params{},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			if tc.shouldErr {
				suite.Require().Panics(func() { suite.Require().Equal(tc.expParams, suite.k.GetParams(ctx)) })
			} else {
				suite.Require().NotPanics(func() { suite.Require().Equal(tc.expParams, suite.k.GetParams(ctx)) })
			}
		})
	}
}
