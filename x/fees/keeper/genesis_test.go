package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "params are exported properly",
			store: func(ctx sdk.Context) {
				suite.keeper.SetParams(ctx, types.NewParams(
					[]types.MinFee{
						types.NewMinFee(
							sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
							sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
						),
					},
				))
			},
			expGenesis: types.NewGenesisState(
				types.NewParams(
					[]types.MinFee{
						types.NewMinFee(
							sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
							sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
						),
					},
				),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			genesis := suite.keeper.ExportGenesis(ctx)
			suite.Require().Equal(tc.expGenesis, genesis)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_InitGenesis() {
	testCases := []struct {
		name  string
		data  types.GenesisState
		check func(ctx sdk.Context)
	}{
		{
			name: "params are initialized properly",
			data: types.GenesisState{
				Params: types.NewParams(
					[]types.MinFee{
						types.NewMinFee(
							sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
							sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
						),
					},
				),
			},
			check: func(ctx sdk.Context) {
				params := types.NewParams(
					[]types.MinFee{
						types.NewMinFee(
							sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
							sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
						),
					},
				)
				suite.Require().Equal(params, suite.keeper.GetParams(ctx))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			suite.keeper.InitGenesis(ctx, tc.data)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
