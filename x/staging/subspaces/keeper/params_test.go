package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetParams() {
	nameParams := types.NewNameParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(10))
	params := types.NewParams(nameParams)

	suite.k.SetParams(suite.ctx, params)

	actualParams := suite.k.GetParams(suite.ctx)
	suite.Equal(params, actualParams)
}

func (suite *KeeperTestsuite) TestKeeper_GetParams() {
	nameParams := types.NewNameParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(10))
	params := types.NewParams(nameParams)

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
				suite.Equal(*test.expParams, suite.k.GetParams(suite.ctx))
			}
		})
	}
}
