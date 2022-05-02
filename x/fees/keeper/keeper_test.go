package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetParams() {
	testCases := []struct {
		name   string
		store  func(ctx sdk.Context)
		params types.Params
	}{
		{
			name: "params are stored properly",
			params: types.NewParams([]types.MinFee{
				types.NewMinFee(
					sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
				),
			}),
		},
		{
			name: "params are overridden properly",
			store: func(ctx sdk.Context) {
				suite.keeper.SetParams(ctx, types.NewParams([]types.MinFee{
					types.NewMinFee(
						sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
						sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
					),
				}))
			},
			params: types.NewParams([]types.MinFee{
				types.NewMinFee(
					sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(20000))),
				),
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.keeper.SetParams(ctx, tc.params)
			suite.Require().Equal(tc.params, suite.keeper.GetParams(ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetParams() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expParams types.Params
	}{
		{
			name: "params are returned properly",
			store: func(ctx sdk.Context) {
				suite.keeper.SetParams(ctx, types.NewParams([]types.MinFee{
					types.NewMinFee(
						sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
						sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
					),
				}))
			},
			expParams: types.NewParams([]types.MinFee{
				types.NewMinFee(
					sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
					sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
				),
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			suite.Require().Equal(tc.expParams, suite.keeper.GetParams(ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_CheckFees() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		fees      sdk.Coins
		msgs      []sdk.Msg
		shouldErr bool
	}{
		{
			name: "not enough fees returns error - single message",
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
			msgs:      []sdk.Msg{&profilestypes.MsgSaveProfile{}},
			fees:      sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5)),
			shouldErr: true,
		},
		{
			name: "not enough fees returns error - multiple messages",
			store: func(ctx sdk.Context) {
				suite.keeper.SetParams(ctx, types.NewParams(
					[]types.MinFee{
						types.NewMinFee(
							sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
							sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
						),
						types.NewMinFee(
							sdk.MsgTypeURL(&profilestypes.MsgCancelDTagTransferRequest{}),
							sdk.NewCoins(sdk.NewCoin("photino", sdk.NewInt(1))),
						),
					},
				))
			},
			msgs: []sdk.Msg{
				&profilestypes.MsgSaveProfile{},
				&profilestypes.MsgCancelDTagTransferRequest{},
			},
			fees: sdk.NewCoins(
				sdk.NewInt64Coin(sdk.DefaultBondDenom, 10),
			),
			shouldErr: true,
		},
		{
			name: "enough fees work properly - single message",
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
			msgs:      []sdk.Msg{&profilestypes.MsgSaveProfile{}},
			fees:      sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10)),
			shouldErr: false,
		},
		{
			name: "enough fees work properly - multiple messages",
			store: func(ctx sdk.Context) {
				suite.keeper.SetParams(ctx, types.NewParams(
					[]types.MinFee{
						types.NewMinFee(
							sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
							sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
						),
						types.NewMinFee(
							sdk.MsgTypeURL(&profilestypes.MsgCancelDTagTransferRequest{}),
							sdk.NewCoins(sdk.NewCoin("photino", sdk.NewInt(1))),
						),
					},
				))
			},
			msgs: []sdk.Msg{
				&profilestypes.MsgSaveProfile{},
				&profilestypes.MsgSaveProfile{},
				&profilestypes.MsgCancelDTagTransferRequest{},
			},
			fees: sdk.NewCoins(
				sdk.NewInt64Coin(sdk.DefaultBondDenom, 20),
				sdk.NewInt64Coin("photino", 1),
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.keeper.CheckFees(ctx, tc.msgs, tc.fees)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
