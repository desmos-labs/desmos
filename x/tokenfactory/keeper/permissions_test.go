package keeper_test

import (
	"github.com/golang/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"
	"github.com/desmos-labs/desmos/v7/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestKeeper_ValidateSubspaceTokenPermission() {

	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		subspace  subspacestypes.Subspace
		sender    string
		denom     string
		shouldErr bool
	}{
		{
			name: "no permissions returns error",
			setup: func() {
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
			subspace: subspacestypes.Subspace{
				ID:       1,
				Treasury: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			sender:    "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			denom:     "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
			shouldErr: true,
		},
		{
			name: "denom does not exist returns error",
			setup: func() {
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
					Return(banktypes.Metadata{}, false)
			},
			subspace: subspacestypes.Subspace{
				ID:       1,
				Treasury: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			sender:    "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			denom:     "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
			shouldErr: true,
		},
		{
			name: "denom admin does not match subspace treasury returns error",
			setup: func() {
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
			},
			subspace: subspacestypes.Subspace{
				ID:       1,
				Treasury: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			sender:    "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			denom:     "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
			shouldErr: true,
		},
		{
			name: "subspace treasury does not match admin returns error",
			setup: func() {
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
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: ""})
			},
			subspace: subspacestypes.Subspace{
				ID:       1,
				Treasury: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			sender:    "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			denom:     "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
			shouldErr: true,
		},
		{
			name: "valid request returns no error",
			setup: func() {
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
			},
			store: func(ctx sdk.Context) {
				suite.k.SetAuthorityMetadata(ctx,
					"factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
					types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"})
			},
			subspace: subspacestypes.Subspace{
				ID:       1,
				Treasury: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			sender: "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			denom:  "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken",
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

			err := suite.k.ValidateManageTokenPermission(ctx, tc.subspace, tc.sender, tc.denom)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
