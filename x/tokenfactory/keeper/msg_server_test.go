package keeper_test

import (
	"fmt"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"

	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v6/x/tokenfactory/keeper"
	"github.com/desmos-labs/desmos/v6/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestMsgServer_CreateDenom() {
	testCases := []struct {
		name        string
		setup       func()
		msg         *types.MsgCreateDenom
		shouldErr   bool
		expResponse *types.MsgCreateDenomResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "subspace does not exist returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.Subspace{}, false)
			},
			msg: types.NewMsgCreateDenom(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"uminttoken",
			),
			shouldErr: true,
		},
		{
			name: "no permissions returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(false)
			},
			msg: types.NewMsgCreateDenom(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"uminttoken",
			),
			shouldErr: true,
		},
		{
			name: "create denom failed returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(true)
			},
			msg: types.NewMsgCreateDenom(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"uminttoken",
			),
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					HasSupply(gomock.Any(), "uminttoken").
					Return(false)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, false).
					Times(2)

				suite.bk.EXPECT().
					SetDenomMetaData(gomock.Any(), banktypes.Metadata{
						DenomUnits: []*banktypes.DenomUnit{{
							Denom:    "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
							Exponent: 0,
						}},
						Base: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					})
			},
			msg: types.NewMsgCreateDenom(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"uminttoken",
			),
			expResponse: &types.MsgCreateDenomResponse{
				NewTokenDenom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateDenom{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
				),
				sdk.NewEvent(
					types.EventTypeCreateDenom,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeCreator, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
					sdk.NewAttribute(types.AttributeNewTokenDenom, "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken"),
				),
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
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.CreateDenom(ctx, tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_Mint() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgMint
		shouldErr   bool
		expResponse *types.MsgMintResponse
		expEvents   sdk.Events
	}{
		{
			name: "subspace does not exist returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.Subspace{}, false)
			},
			msg: types.NewMsgMint(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			shouldErr: true,
		},
		{
			name: "no permissions returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(false)
			},
			msg: types.NewMsgMint(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			shouldErr: true,
		},
		{
			name: "mint failed returns error - bank mint coins failed",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, true)

				suite.bk.EXPECT().
					MintCoins(
						gomock.Any(),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(fmt.Errorf("error"))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			msg: types.NewMsgMint(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			shouldErr: true,
		},
		{
			name: "mint failed returns error - module send coins to treasury failed",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, true)

				suite.bk.EXPECT().
					MintCoins(
						gomock.Any(),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(nil)

				suite.bk.EXPECT().
					SendCoinsFromModuleToAccount(
						gomock.Any(),
						types.ModuleName,
						sdk.MustAccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(fmt.Errorf("error"))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			msg: types.NewMsgMint(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, true)

				suite.bk.EXPECT().
					MintCoins(
						gomock.Any(),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(nil)

				suite.bk.EXPECT().
					SendCoinsFromModuleToAccount(
						gomock.Any(),
						types.ModuleName,
						sdk.MustAccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(nil)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			msg: types.NewMsgMint(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			expResponse: &types.MsgMintResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgMint{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
				),
				sdk.NewEvent(
					types.EventTypeMint,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeMintToAddress, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeAmount, sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)).String()),
				),
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.Mint(ctx, tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_Burn() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgBurn
		shouldErr   bool
		expResponse *types.MsgBurnResponse
		expEvents   sdk.Events
	}{
		{
			name: "subspace does not exist returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.Subspace{}, false)
			},
			msg: types.NewMsgBurn(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			shouldErr: true,
		},
		{
			name: "no permissions returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(false)
			},
			msg: types.NewMsgBurn(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			shouldErr: true,
		},
		{
			name: "burn failed returns error - send coins to module failed",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, true)

				suite.bk.EXPECT().
					SendCoinsFromAccountToModule(
						gomock.Any(),
						sdk.MustAccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(fmt.Errorf("error"))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			msg: types.NewMsgBurn(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			shouldErr: true,
		},
		{
			name: "burn failed returns error - bank burn failed",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, true)

				suite.bk.EXPECT().
					SendCoinsFromAccountToModule(
						gomock.Any(),
						sdk.MustAccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(nil)

				suite.bk.EXPECT().
					BurnCoins(
						gomock.Any(),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(fmt.Errorf("error"))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			msg: types.NewMsgBurn(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, true)

				suite.bk.EXPECT().
					SendCoinsFromAccountToModule(
						gomock.Any(),
						sdk.MustAccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(nil)

				suite.bk.EXPECT().
					BurnCoins(
						gomock.Any(),
						types.ModuleName,
						sdk.NewCoins(sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100))),
					).
					Return(nil)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			msg: types.NewMsgBurn(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)),
			),
			expResponse: &types.MsgBurnResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgBurn{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
				),
				sdk.NewEvent(
					types.EventTypeBurn,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeBurnFromAddress, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeAmount, sdk.NewCoin("factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", math.NewInt(100)).String()),
				),
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.Burn(ctx, tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_SetDenomMetadata() {
	metadata := banktypes.Metadata{
		Name:        "Mint Token",
		Symbol:      "MTK",
		Description: "The custom token of the test subspace.",
		DenomUnits: []*banktypes.DenomUnit{
			{Denom: "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken", Exponent: uint32(0), Aliases: nil},
			{Denom: "minttoken", Exponent: uint32(6), Aliases: []string{"minttoken"}},
		},
		Base:    "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
		Display: "minttoken",
	}

	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgSetDenomMetadata
		shouldErr   bool
		expResponse *types.MsgSetDenomMetadataResponse
		expEvents   sdk.Events
	}{
		{
			name: "subspace does not exist returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.Subspace{}, false)
			},
			msg: types.NewMsgSetDenomMetadata(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				metadata,
			),
			shouldErr: true,
		},
		{
			name: "no permissions returns error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(false)
			},
			msg: types.NewMsgSetDenomMetadata(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				metadata,
			),
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			setup: func() {
				suite.sk.EXPECT().
					GetSubspace(gomock.Any(), uint64(1)).
					Return(subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(1),
						uint32(subspacestypes.RootSectionID),
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						types.PermissionManageSubspaceTokens,
					).
					Return(true)

				suite.bk.EXPECT().
					GetDenomMetaData(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(banktypes.Metadata{}, true)

				suite.bk.EXPECT().
					SetDenomMetaData(gomock.Any(), metadata)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			msg: types.NewMsgSetDenomMetadata(
				1,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				metadata,
			),
			expResponse: &types.MsgSetDenomMetadataResponse{},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgSetDenomMetadata{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
				),
				sdk.NewEvent(
					types.EventTypeSetDenomMetadata,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeDenom, metadata.Base),
					sdk.NewAttribute(types.AttributeDenomMetadata, metadata.String()),
				),
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.SetDenomMetadata(ctx, tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_UpdateParams() {
	testCases := []struct {
		name      string
		setup     func()
		msg       *types.MsgUpdateParams
		shouldErr bool
		check     func(ctx sdk.Context)
		expEvents sdk.Events
	}{
		{
			name: "invalid authority return error",
			msg: types.NewMsgUpdateParams(
				types.DefaultParams(),
				"invalid",
			),
			shouldErr: true,
		},
		{
			name: "set params properly",
			msg: types.NewMsgUpdateParams(
				types.NewParams(sdk.NewCoins(sdk.NewCoin("test", math.NewInt(100)))),
				authtypes.NewModuleAddress("gov").String(),
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				suite.Require().Equal(
					types.NewParams(sdk.NewCoins(sdk.NewCoin("test", math.NewInt(100)))),
					suite.k.GetParams(ctx),
				)
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgUpdateParams{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"),
				),
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

			// Reset any event that might have been emitted during the setup
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.UpdateParams(ctx, tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
