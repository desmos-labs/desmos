package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"

	"github.com/desmos-labs/desmos/v3/x/fees/keeper"
	"github.com/desmos-labs/desmos/v3/x/fees/types"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (suite *KeeperTestSuite) TestFeeDecorator() {
	testsCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		tx         sdk.Tx
		shouldErr  bool
		zeroHeight bool
		simulate   bool
	}{
		{
			name: "not existing min fees returns no error",
			store: func(ctx sdk.Context) {
				suite.keeper.SetParams(ctx, types.NewParams(nil))
			},
			tx: legacytx.NewStdTx(
				[]sdk.Msg{profilestypes.NewMsgSaveProfile("", "", "", "", "", "")},
				legacytx.NewStdFee(10, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 0))),
				nil,
				"",
			),
			shouldErr:  false,
			zeroHeight: false,
			simulate:   true,
		},
		{
			name: "not enough fees returns error",
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
			tx: legacytx.NewStdTx(
				[]sdk.Msg{profilestypes.NewMsgSaveProfile("", "", "", "", "", "")},
				legacytx.NewStdFee(10, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5))),
				nil,
				"",
			),
			shouldErr:  true,
			zeroHeight: false,
			simulate:   true,
		},
		{
			name: "block height 0 skips minimum fees checks",
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
			tx: legacytx.NewStdTx(
				[]sdk.Msg{profilestypes.NewMsgSaveProfile("", "", "", "", "", "")},
				legacytx.NewStdFee(10, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 5))),
				nil,
				"",
			),
			shouldErr:  false,
			simulate:   false,
			zeroHeight: true,
		},
		{
			name: "enough fees returns no error",
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
			tx: legacytx.NewStdTx(
				[]sdk.Msg{profilestypes.NewMsgSaveProfile("", "", "", "", "", "")},
				legacytx.NewStdFee(10, sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 10))),
				nil,
				"",
			),
			shouldErr:  false,
			zeroHeight: false,
			simulate:   true,
		},
	}

	for _, tc := range testsCases {
		ctx, _ := suite.ctx.CacheContext()
		if tc.store != nil {
			if tc.zeroHeight {
				ctx = suite.ctx.WithBlockHeight(0)
			}
			tc.store(ctx)
		}

		anteHandler := sdk.ChainAnteDecorators(keeper.NewMinFeeDecorator(suite.keeper))
		_, err := anteHandler(ctx, tc.tx, tc.simulate)
		if tc.shouldErr {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)
		}
	}

}
