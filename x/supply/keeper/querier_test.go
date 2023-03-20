package keeper_test

import (
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/supply/keeper"
	"github.com/desmos-labs/desmos/v4/x/supply/types"
)

func (suite *KeeperTestSuite) TestQuerier_QueryTotalSupply() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		path      []string
		request   abci.RequestQuery
		expSupply sdk.Int
	}{
		{
			name: "total supply is returned correctly",
			store: func(ctx sdk.Context) {
				suite.setupSupply(ctx,
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(200))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(300))),
				)
			},
			path: []string{types.QueryTotalSupply},
			request: abci.RequestQuery{
				Data: suite.cdc.MustMarshal(
					types.NewQueryTotalRequest(sdk.DefaultBondDenom, 3),
				),
			},
			expSupply: sdk.NewInt(1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			querier := keeper.NewQuerier(suite.k)
			res, err := querier(ctx, tc.path, tc.request)
			suite.Require().NoError(err)

			var supply sdk.Int
			err = supply.Unmarshal(res)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expSupply, supply)
		})
	}
}

func (suite *KeeperTestSuite) TestQuerier_QueryCirculatingSupply() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		path      []string
		request   abci.RequestQuery
		expSupply sdk.Int
	}{
		{
			name: "circulating supply is returned correctly",
			store: func(ctx sdk.Context) {
				suite.setupSupply(ctx,
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(200))),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(300))),
				)
			},
			path: []string{types.QueryCirculatingSupply, "1000"},
			request: abci.RequestQuery{
				Data: suite.cdc.MustMarshal(
					types.NewQueryCirculatingRequest(sdk.DefaultBondDenom, 0),
				),
			},
			expSupply: sdk.NewInt(500),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			querier := keeper.NewQuerier(suite.k)
			res, err := querier(ctx, tc.path, tc.request)
			suite.Require().NoError(err)

			var supply sdk.Int
			err = supply.Unmarshal(res)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expSupply, supply)
		})
	}
}
