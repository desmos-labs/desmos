package keeper_test

import (
	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	params := types.DefaultParams()
	suite.k.SetParams(suite.ctx, params)

	actualParams := suite.k.GetParams(suite.ctx)

	suite.Require().Equal(params, actualParams)
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	params := types.DefaultParams()
	suite.k.SetParams(suite.ctx, params)

	actualParams := suite.k.GetParams(suite.ctx)

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
				suite.k.SetParams(suite.ctx, *test.params)
			}

			if test.expParams != nil {
				suite.Require().Equal(*test.expParams, suite.k.GetParams(suite.ctx))
			}
		})
	}
}
