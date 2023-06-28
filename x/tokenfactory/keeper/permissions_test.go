package keeper_test

import (
	"fmt"

	"github.com/golang/mock/gomock"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

func (suite *KeeperTestSuite) TestKeeper_ValidateSubspaceTokenPermission() {

	testCases := []struct {
		name      string
		setup     func()
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
			name: "get denom authority failed returns error",
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

				suite.tfk.EXPECT().
					GetAuthorityMetadata(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(types.DenomAuthorityMetadata{}, fmt.Errorf("error"))
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

				suite.tfk.EXPECT().
					GetAuthorityMetadata(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(types.DenomAuthorityMetadata{Admin: "non-treasury-account"}, nil)
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

				suite.tfk.EXPECT().
					GetAuthorityMetadata(gomock.Any(), "factory/cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47/uminttoken").
					Return(types.DenomAuthorityMetadata{Admin: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"}, nil)
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

			err := suite.k.ValidateManageTokenPermission(ctx, tc.subspace, tc.sender, tc.denom)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
