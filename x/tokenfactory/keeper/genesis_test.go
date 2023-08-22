package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v6/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "genesis denoms are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.AddDenomFromCreator(ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/bitcoin")

				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/bitcoin",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})

				suite.k.AddDenomFromCreator(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/litecoin")

				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/litecoin",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			expGenesis: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{Denom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/bitcoin", AuthorityMetadata: types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}},
					{Denom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/litecoin", AuthorityMetadata: types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}},
				},
			},
		},
		{
			name: "params are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			expGenesis: &types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{},
				Params:        types.DefaultParams(),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			genesis := suite.k.ExportGenesis(ctx)
			suite.Require().Equal(tc.expGenesis, genesis)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_InitGenesis() {
	testCases := []struct {
		name  string
		setup func()
		data  types.GenesisState
		check func(ctx sdk.Context)
	}{
		{
			name: "denoms are imported properly",
			setup: func() {
				suite.ak.EXPECT().GetModuleAccount(gomock.Any(), types.ModuleName)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/bitcoin").
					Return(banktypes.Metadata{}, true)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/litecoin").
					Return(banktypes.Metadata{}, true)
			},
			data: types.GenesisState{
				FactoryDenoms: []types.GenesisDenom{
					{Denom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/bitcoin", AuthorityMetadata: types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}},
					{Denom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/litecoin", AuthorityMetadata: types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}},
				},
			},
			check: func(ctx sdk.Context) {
				suite.Require().Equal(
					[]string{"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/bitcoin", "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/litecoin"},
					suite.k.GetDenomsFromCreator(ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				)

				suite.Require().Equal(
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
					suite.k.GetAuthorityMetadata(ctx, "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/bitcoin"),
				)

				suite.Require().Equal(
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
					suite.k.GetAuthorityMetadata(ctx, "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/litecoin"),
				)
			},
		},
		{
			name: "params are imported properly",
			setup: func() {
				suite.ak.EXPECT().GetModuleAccount(gomock.Any(), types.ModuleName)
			},
			data: types.GenesisState{
				Params: types.DefaultParams(),
			},
			check: func(ctx sdk.Context) {
				suite.Require().Equal(types.DefaultParams(), suite.k.GetParams(ctx))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}

			suite.k.InitGenesis(ctx, tc.data)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
