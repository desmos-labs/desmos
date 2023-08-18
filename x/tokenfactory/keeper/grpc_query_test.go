package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
	"github.com/desmos-labs/desmos/v6/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Params() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryParamsRequest
		shouldErr bool
		expParams types.Params
	}{
		{
			name: "params are returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			request:   types.NewQueryParamsRequest(),
			shouldErr: false,
			expParams: types.DefaultParams(),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.Params(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expParams, res.Params)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_SubspaceDenoms() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QuerySubspaceDenomsRequest
		shouldErr bool
		expDenoms []string
	}{
		{
			name: "denoms are returned properly",
			store: func(ctx sdk.Context) {
				suite.k.AddDenomFromCreator(ctx, subspacestypes.GetTreasuryAddress(1).String(), "bitcoin")
				suite.k.AddDenomFromCreator(ctx, subspacestypes.GetTreasuryAddress(1).String(), "minttoken")
			},
			request:   types.NewQuerySubspaceDenomsRequest(1),
			shouldErr: false,
			expDenoms: []string{"bitcoin", "minttoken"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.SubspaceDenoms(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expDenoms, res.Denoms)
			}
		})
	}
}
