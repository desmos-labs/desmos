package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	nsParams := types.NewMonikerParams(sdk.NewInt(3), sdk.NewInt(1000))
	monikerParams := types.NewDtagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(1000))

	params := types.NewParams(nsParams, monikerParams, sdk.NewInt(1000))
	suite.k.SetParams(suite.ctx, params)

	actualParams := suite.k.GetParams(suite.ctx)
	suite.Require().Equal(params, actualParams)
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	nsParams := types.NewMonikerParams(sdk.NewInt(3), sdk.NewInt(1000))
	monikerParams := types.NewDtagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(1000))
	params := types.NewParams(nsParams, monikerParams, sdk.NewInt(1000))

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
				suite.k.SetParams(suite.ctx, *test.params)
			}

			if test.expParams != nil {
				suite.Require().Equal(*test.expParams, suite.k.GetParams(suite.ctx))
			}
		})
	}
}
