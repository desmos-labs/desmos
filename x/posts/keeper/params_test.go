package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/posts/types"
	"github.com/stretchr/testify/require"
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
