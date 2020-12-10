package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/fees/keeper"
	"github.com/desmos-labs/desmos/x/fees/types"
)

func (suite *KeeperTestSuite) Test_queryParams() {
	tests := []struct {
		name      string
		path      []string
		params    types.Params
		expResult types.Params
	}{
		{
			name:      "Returning fees default parameters correctly",
			params:    types.DefaultParams(),
			path:      []string{types.QueryParams},
			expResult: types.DefaultParams(),
		},
		{
			name: "Returning fees parameters correctly",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
			path: []string{types.QueryParams},
			expResult: types.NewParams([]types.MinFee{
				types.NewMinFee("create_post", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
			}),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.keeper.SetParams(suite.ctx, test.params)

			querier := keeper.NewQuerier(suite.keeper, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})
			suite.Require().Nil(err)

			expected := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
			suite.Equal(string(expected), string(result))
		})
	}
}
