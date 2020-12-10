package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/fees/types"
	posts "github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	params := types.DefaultParams()
	suite.keeper.SetParams(suite.ctx, params)

	actualParams := suite.keeper.GetParams(suite.ctx)

	suite.Equal(params, actualParams)
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	params := types.DefaultParams()
	suite.keeper.SetParams(suite.ctx, params)

	actualParams := suite.keeper.GetParams(suite.ctx)

	suite.Equal(params, actualParams)

	tests := []struct {
		name      string
		params    *types.Params
		expParams *types.Params
	}{
		{
			name:      "Returning previously set params",
			params:    &params,
			expParams: &params,
		},
		{
			name:      "Returning nothing",
			params:    nil,
			expParams: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			if test.params != nil {
				suite.keeper.SetParams(suite.ctx, *test.params)
			}

			if test.expParams != nil {
				suite.Equal(*test.expParams, suite.keeper.GetParams(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_CheckFees() {
	tests := []struct {
		name      string
		params    types.Params
		givenFees sdk.Coins
		msgs      []sdk.Msg
		expError  bool
	}{
		{
			name: "Not enough fees returns error",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			givenFees: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 150)),
			msgs: []sdk.Msg{
				posts.NewMsgCreatePost(
					"My new post",
					"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					nil,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					nil,
					nil,
				),
			},
			expError: true,
		},
		{
			name: "Enough fees works properly",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			givenFees: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)),
			msgs: []sdk.Msg{
				posts.NewMsgCreatePost(
					"My new post",
					"dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					nil,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					nil,
					nil,
				),
			},
			expError: false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.keeper.SetParams(suite.ctx, test.params)
			err := suite.keeper.CheckFees(suite.ctx, test.givenFees, test.msgs)
			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
