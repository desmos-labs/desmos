package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	tests := []struct {
		name      string
		params    types.Params
		expError  bool
		expParams types.Params
	}{
		{
			name:     "Storing empty params returns error",
			params:   types.Params{},
			expError: true,
		},
		{
			name:      "Default params are stored properly",
			params:    types.DefaultParams(),
			expError:  false,
			expParams: types.DefaultParams(),
		},
		{
			name:      "Non default params are stored properly",
			params:    types.NewParams(sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1)),
			expError:  false,
			expParams: types.NewParams(sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1)),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.expError {
				suite.Require().Panics(func() { suite.keeper.SetParams(suite.ctx, test.params) })
			} else {
				suite.keeper.SetParams(suite.ctx, test.params)
				suite.Require().Equal(test.expParams, suite.keeper.GetParams(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	params := types.DefaultParams()
	suite.keeper.SetParams(suite.ctx, params)

	actualParams := suite.keeper.GetParams(suite.ctx)

	suite.Require().Equal(params, actualParams)

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
				suite.Require().Equal(*test.expParams, suite.keeper.GetParams(suite.ctx))
			}
		})
	}
}
