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
		name      string
		store     func(ctx sdk.Context)
		tx        sdk.Tx
		shouldErr bool
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
			shouldErr: false,
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
			shouldErr: true,
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
			shouldErr: false,
		},
	}

	for _, tc := range testsCases {
		ctx, _ := suite.ctx.CacheContext()
		if tc.store != nil {
			tc.store(ctx)
		}

		anteHandler := sdk.ChainAnteDecorators(keeper.NewMinFeeDecorator(suite.keeper))
		_, err := anteHandler(ctx, tc.tx, true)
		if tc.shouldErr {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)
		}
	}

}
