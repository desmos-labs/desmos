package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/internal/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	min := sdk.NewInt(3)
	max := sdk.NewInt(1000)
	nsParams := types.NewMonikerParams(min, max)
	monikerParams := types.NewDtagParams("^[A-Za-z0-9_]+$", min, max)

	params := types.NewParams(nsParams, monikerParams, max)

	suite.keeper.SetParams(suite.ctx, params)

	actualParams := suite.keeper.GetParams(suite.ctx)

	suite.Equal(params, actualParams)
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	min := sdk.NewInt(3)
	max := sdk.NewInt(1000)
	nsParams := types.NewMonikerParams(min, max)
	monikerParams := types.NewDtagParams("^[A-Za-z0-9_]+$", min, max)
	params := types.NewParams(nsParams, monikerParams, max)

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
