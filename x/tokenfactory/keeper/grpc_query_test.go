package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Params() {
	testCases := []struct {
		name      string
		setup     func()
		request   *types.QueryParamsRequest
		shouldErr bool
		expParams types.Params
	}{
		{
			name: "params are returned properly",
			setup: func() {
				suite.tfk.EXPECT().GetParams(gomock.Any()).
					Return(types.DefaultParams())
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
			if tc.setup != nil {
				tc.setup()
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
		setup     func()
		request   *types.QuerySubspaceDenomsRequest
		shouldErr bool
		expDenoms []string
	}{
		{
			name: "params are returned properly",
			setup: func() {
				suite.tfk.EXPECT().GetDenomsFromCreator(gomock.Any(), subspacestypes.GetTreasuryAddress(1).String()).
					Return([]string{"minttoken", "bitcoin"})
			},
			request:   types.NewQuerySubspaceDenomsRequest(1),
			shouldErr: false,
			expDenoms: []string{"minttoken", "bitcoin"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
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
