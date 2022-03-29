package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/fees/types"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
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
				types.NewMinFee("desmos.profiles.v2.MsgSaveProfile", sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000)))),
			}),
			givenFees: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 150)),
			msgs: []sdk.Msg{
				profilestypes.NewMsgSaveProfile("", "", "", "", "", ""),
			},
			expError: true,
		},
		{
			name: "Enough fees works properly",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("desmos.profiles.v2.MsgSaveProfile", sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000)))),
			}),
			givenFees: sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000)),
			msgs: []sdk.Msg{
				profilestypes.NewMsgSaveProfile("", "", "", "", "", ""),
			},
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
