package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	types2 "github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	tests := []struct {
		name      string
		params    types2.Params
		expError  bool
		expParams types2.Params
	}{
		{
			name:     "Storing empty params returns error",
			params:   types2.Params{},
			expError: true,
		},
		{
			name:      "Default params are stored properly",
			params:    types2.DefaultParams(),
			expError:  false,
			expParams: types2.DefaultParams(),
		},
		{
			name:      "Non default params are stored properly",
			params:    types2.NewParams(sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1)),
			expError:  false,
			expParams: types2.NewParams(sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1)),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.expError {
				suite.Require().Panics(func() { suite.k.SetParams(suite.ctx, test.params) })
			} else {
				suite.k.SetParams(suite.ctx, test.params)
				suite.Require().Equal(test.expParams, suite.k.GetParams(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	params := types2.DefaultParams()
	suite.k.SetParams(suite.ctx, params)

	actualParams := suite.k.GetParams(suite.ctx)

	suite.Require().Equal(params, actualParams)

	tests := []struct {
		name      string
		params    *types2.Params
		expParams *types2.Params
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
