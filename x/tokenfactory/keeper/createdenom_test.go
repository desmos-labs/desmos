package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestKeeper_CreateDenom() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		shouldErr bool
		creator   string
		subdenom  string
		expDenom  string
		check     func(ctx sdk.Context)
	}{
		{
			name: "subdenom with existing supply returns error",
			setup: func() {
				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(true)
			},
			subdenom:  "uminttoken",
			creator:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			shouldErr: true,
		},
		{
			name: "invalid token denoms returns error",
			setup: func() {
				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(false)
			},
			subdenom:  "uminttoken",
			creator:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/",
			shouldErr: true,
		},
		{
			name: "existing denom returns error",
			setup: func() {
				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(false)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, true)
			},
			subdenom:  "uminttoken",
			creator:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			shouldErr: true,
		},
		{
			name: "burn creation fees returns error - send coins to module failed",
			setup: func() {
				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(false)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, false)

				suite.bk.EXPECT().
					SendCoinsFromAccountToModule(
						gomock.Any(),
						sdk.MustAccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
					).
					Return(fmt.Errorf("send coin to module error"))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))))
			},
			subdenom:  "uminttoken",
			creator:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			shouldErr: true,
		},
		{
			name: "burn creation fees returns error - bank burn failed",
			setup: func() {
				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(false)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, false)

				suite.bk.EXPECT().
					SendCoinsFromAccountToModule(
						gomock.Any(),
						sdk.MustAccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
					).
					Return(nil)

				suite.bk.EXPECT().
					BurnCoins(gomock.Any(), types.ModuleName, sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))).
					Return(fmt.Errorf("bank burn coins error"))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))))
			},
			subdenom:  "uminttoken",
			creator:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			shouldErr: true,
		},
		{
			name: "create denom properly",
			setup: func() {
				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(false)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, false).Times(2)

				suite.bk.EXPECT().
					SendCoinsFromAccountToModule(
						gomock.Any(),
						sdk.MustAccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100))),
					).
					Return(nil)

				suite.bk.EXPECT().
					BurnCoins(gomock.Any(), types.ModuleName, sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))).
					Return(nil)

				suite.bk.EXPECT().
					SetDenomMetaData(gomock.Any(), banktypes.Metadata{
						DenomUnits: []*banktypes.DenomUnit{{
							Denom:    "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
							Exponent: 0,
						}},
						Base: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					})
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))))
			},
			check: func(ctx sdk.Context) {
				suite.Require().Equal(
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
					suite.k.GetAuthorityMetadata(ctx, "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken"),
				)
				suite.Require().Equal(
					[]string{"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken"},
					suite.k.GetDenomsFromCreator(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				)
			},
			subdenom: "uminttoken",
			creator:  "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expDenom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
		},
		{
			name: "create denom properly - no creation fees",
			setup: func() {
				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(false)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, false).Times(2)

				suite.bk.EXPECT().
					SetDenomMetaData(gomock.Any(), banktypes.Metadata{
						DenomUnits: []*banktypes.DenomUnit{{
							Denom:    "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
							Exponent: 0,
						}},
						Base: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					})
			},
			check: func(ctx sdk.Context) {
				suite.Require().Equal(
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
					suite.k.GetAuthorityMetadata(ctx, "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken"),
				)
				suite.Require().Equal(
					[]string{"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken"},
					suite.k.GetDenomsFromCreator(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				)

			},
			subdenom: "uminttoken",
			creator:  "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expDenom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			denom, err := suite.k.CreateDenom(ctx, tc.creator, tc.subdenom)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expDenom, denom)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
