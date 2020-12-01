package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/fees/keeper"
	"github.com/desmos-labs/desmos/x/fees/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) Test_queryParams() {
	tests := []struct {
		name      string
		path      []string
		expResult types.Params
	}{
		{
			name:      "Returning posts parameters correctly",
			path:      []string{types.QueryParams},
			expResult: types.DefaultParams(),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.keeper.SetParams(suite.ctx, types.DefaultParams())

			querier := keeper.NewQuerier(suite.keeper, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})
			suite.Require().Nil(err)

			expected := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
			suite.Equal(string(expected), string(result))
		})
	}
}
