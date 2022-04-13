package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Params() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryParamsRequest
		expParams types.Params
	}{
		{
			name: "valid request returns proper data",
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
			req: types.NewQueryParamsRequest(),
			expParams: types.NewParams(
				[]types.MinFee{
					types.NewMinFee(
						sdk.MsgTypeURL(&profilestypes.MsgSaveProfile{}),
						sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))),
					),
				},
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

			res, err := suite.keeper.Params(sdk.WrapSDKContext(ctx), tc.req)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expParams, res.Params)
		})
	}
}
