package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v4/x/fees/keeper"
	"github.com/desmos-labs/desmos/v4/x/fees/types"
)

func (suite *KeeperTestSuite) TestMsgServer_UpdateParams() {
	testCases := []struct {
		name        string
		msg         *types.MsgUpdateParams
		shouldErr   bool
		expResponse *types.MsgUpdateParams
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "invalid authority return error",
			msg: types.NewMsgUpdateParams(
				types.NewParams([]types.MinFee{
					types.NewMinFee(
						"test.v1beta1.MsgTest",
						sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
					)},
				),
				"invalid",
			),
			shouldErr: true,
		},
		{
			name: "set params properly",
			msg: types.NewMsgUpdateParams(
				types.NewParams([]types.MinFee{
					types.NewMinFee(
						"test.v1beta1.MsgTest",
						sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
					)},
				),
				"invalid",
			),
			shouldErr: true,
			check: func(ctx sdk.Context) {
				params := suite.keeper.GetParams(ctx)
				suite.Require().Equal(types.NewParams([]types.MinFee{
					types.NewMinFee(
						"test.v1beta1.MsgTest",
						sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000))),
					)},
				), params)
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			// Reset any event that might have been emitted during the setup
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the message
			service := keeper.NewMsgServerImpl(suite.keeper)
			res, err := service.UpdateParams(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
