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
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())
			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)
				suite.Equal(string(expectedIndented), string(result))
			}

		})
	}
}
