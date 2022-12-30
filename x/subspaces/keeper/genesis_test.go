package keeper_test

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	allowanceAny, err := codectypes.NewAnyWithValue(&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
	suite.Require().NoError(err)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "subspace id is exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 1)
			},
			expGenesis: types.NewGenesisState(1, nil, nil, nil, nil, nil, nil, nil, nil),
		},
		{
			name: "subspaces and their data are exported correctly",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 3)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextSectionID(ctx, 1, 2)
				suite.k.SetNextGroupID(ctx, 1, 3)

				suite.k.SaveSubspace(ctx, types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextSectionID(ctx, 2, 10)
				suite.k.SetNextGroupID(ctx, 2, 11)
			},
			expGenesis: types.NewGenesisState(
				3,
				[]types.SubspaceData{
					types.NewSubspaceData(1, 2, 3),
					types.NewSubspaceData(2, 10, 11),
				},
				[]types.Subspace{
					types.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
					types.NewSubspace(
						2,
						"Another test subspace",
						"This is another test subspace",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
					),
				},
				[]types.Section{
					types.DefaultSection(1),
					types.DefaultSection(2),
				},
				nil,
				[]types.UserGroup{
					types.DefaultUserGroup(1),
					types.DefaultUserGroup(2),
				},
				nil,
				nil,
				nil,
			),
		},
		{
			name: "sections are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 1)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				))

				suite.k.SaveSection(ctx, types.NewSection(
					2,
					3,
					1,
					"Another test section",
					"This is another test section",
				))
			},
			expGenesis: types.NewGenesisState(1, nil, nil, []types.Section{
				types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				),
				types.NewSection(
					2,
					3,
					1,
					"Another test section",
					"This is another test section",
				),
			}, nil, nil, nil, nil, nil),
		},
		{
			name: "user permissions are exported correctly",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 3)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SetUserPermissions(ctx,
					2,
					0,
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					types.NewPermissions(types.PermissionSetPermissions),
				)
			},
			expGenesis: types.NewGenesisState(
				3,
				[]types.SubspaceData{
					types.NewSubspaceData(2, 1, 1),
				},
				[]types.Subspace{
					types.NewSubspace(
						2,
						"Another test subspace",
						"This is another test subspace",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
					),
				},
				[]types.Section{
					types.DefaultSection(2),
				},
				[]types.UserPermission{
					types.NewUserPermission(
						2,
						0,
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						types.NewPermissions(types.PermissionSetPermissions),
					),
				},
				[]types.UserGroup{
					types.DefaultUserGroup(2),
				},
				nil,
				nil,
				nil,
			),
		},
		{
			name: "user groups and members are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 3)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextGroupID(ctx, 1, 2)
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")

				suite.k.SaveSubspace(ctx, types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextGroupID(ctx, 2, 2)
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					2,
					0,
					1,
					"Another test group",
					"This is another test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 2, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.AddUserToGroup(ctx, 2, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
			},
			expGenesis: types.NewGenesisState(
				3,
				[]types.SubspaceData{
					types.NewSubspaceData(1, 1, 2),
					types.NewSubspaceData(2, 1, 2),
				},
				[]types.Subspace{
					types.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
					types.NewSubspace(
						2,
						"Another test subspace",
						"This is another test subspace",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
					),
				},
				[]types.Section{
					types.DefaultSection(1),
					types.DefaultSection(2),
				},
				nil,
				[]types.UserGroup{
					types.DefaultUserGroup(1),
					types.NewUserGroup(
						1,
						0,
						1,
						"Test group",
						"This is a test group",
						types.NewPermissions(types.PermissionEditSubspace),
					),
					types.DefaultUserGroup(2),
					types.NewUserGroup(
						2,
						0,
						1,
						"Another test group",
						"This is another test group",
						types.NewPermissions(types.PermissionEditSubspace),
					),
				},
				[]types.UserGroupMemberEntry{
					types.NewUserGroupMemberEntry(1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"),
					types.NewUserGroupMemberEntry(2, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"),
					types.NewUserGroupMemberEntry(2, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				},
				nil,
				nil,
			),
		},
		{
			name: "user grants are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 3)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGrant(ctx, types.UserGrant{
					SubspaceID: 1,
					Granter:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					Allowance:  allowanceAny,
				})
				suite.k.SaveUserGrant(ctx, types.UserGrant{
					SubspaceID: 1,
					Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					Allowance:  allowanceAny,
				})
			},
			expGenesis: types.NewGenesisState(
				3,
				[]types.SubspaceData{
					types.NewSubspaceData(1, 1, 1),
				},
				[]types.Subspace{
					types.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
				[]types.Section{
					types.DefaultSection(1),
				},
				nil,
				[]types.UserGroup{
					types.DefaultUserGroup(1),
				},
				nil,
				[]types.UserGrant{{
					SubspaceID: 1,
					Granter:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					Allowance:  allowanceAny,
				}, {
					SubspaceID: 1,
					Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					Allowance:  allowanceAny,
				}},
				nil,
			),
		},
		{
			name: "group grants are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetSubspaceID(ctx, 3)
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetNextGroupID(ctx, 1, 2)
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")

				suite.k.SaveGroupGrant(ctx, types.GroupGrant{
					SubspaceID: 1,
					Granter:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					GroupID:    1,
					Allowance:  allowanceAny,
				})
				suite.k.SaveGroupGrant(ctx, types.GroupGrant{
					SubspaceID: 1,
					Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					GroupID:    1,
					Allowance:  allowanceAny,
				})
			},
			expGenesis: types.NewGenesisState(
				3,
				[]types.SubspaceData{
					types.NewSubspaceData(1, 1, 2),
				},
				[]types.Subspace{
					types.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
				[]types.Section{
					types.DefaultSection(1),
				},
				nil,
				[]types.UserGroup{
					types.DefaultUserGroup(1),
					types.NewUserGroup(
						1,
						0,
						1,
						"Test group",
						"This is a test group",
						types.NewPermissions(types.PermissionEditSubspace),
					),
				},
				[]types.UserGroupMemberEntry{
					types.NewUserGroupMemberEntry(1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"),
					types.NewUserGroupMemberEntry(1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				},
				nil,
				[]types.GroupGrant{{
					SubspaceID: 1,
					Granter:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					GroupID:    1,
					Allowance:  allowanceAny,
				}, {
					SubspaceID: 1,
					Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					GroupID:    1,
					Allowance:  allowanceAny,
				}},
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

			genesis := suite.k.ExportGenesis(ctx)
			suite.Require().Equal(tc.expGenesis, genesis)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_InitGenesis() {
	allowanceAny, err := codectypes.NewAnyWithValue(&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))})
	suite.Require().NoError(err)

	testCases := []struct {
		name    string
		genesis types.GenesisState
		check   func(ctx sdk.Context)
	}{
		{
			name: "initial subspace id is imported properly",
			genesis: types.GenesisState{
				InitialSubspaceID: 2,
			},
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetSubspaceID(ctx)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(2), stored)
			},
		},
		{
			name: "subspaces data are imported properly",
			genesis: types.GenesisState{
				SubspacesData: []types.SubspaceData{
					types.NewSubspaceData(1, 10, 20),
				},
			},
			check: func(ctx sdk.Context) {
				nextSectionID, err := suite.k.GetNextSectionID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(10), nextSectionID)

				nextGroupID, err := suite.k.GetNextGroupID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(20), nextGroupID)
			},
		},
		{
			name: "subspaces are imported properly",
			genesis: types.GenesisState{
				Subspaces: []types.Subspace{
					types.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), stored)
			},
		},
		{
			name: "sections are imported properly",
			genesis: types.GenesisState{
				Sections: []types.Section{
					types.NewSection(
						1,
						2,
						0,
						"Test section",
						"This is a test section",
					),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetSection(ctx, 1, 2)
				suite.Require().True(found)
				suite.Equal(types.NewSection(
					1,
					2,
					0,
					"Test section",
					"This is a test section",
				), stored)
			},
		},
		{
			name: "user groups are imported properly",
			genesis: types.GenesisState{
				UserGroups: []types.UserGroup{
					types.NewUserGroup(
						1,
						0,
						1,
						"Test group",
						"This is a test group",
						types.NewPermissions(types.PermissionEditSubspace),
					),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetUserGroup(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				), stored)
			},
		},
		{
			name: "user group members are imported properly",
			genesis: types.GenesisState{
				UserGroupsMembers: []types.UserGroupMemberEntry{
					types.NewUserGroupMemberEntry(2, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"),
					types.NewUserGroupMemberEntry(2, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				},
			},
			check: func(ctx sdk.Context) {
				groupMembers := suite.k.GetUserGroupMembers(ctx, 2, 1)
				suite.Require().Len(groupMembers, 2)
			},
		},
		{
			name: "user permissions are imported properly",
			genesis: types.GenesisState{
				UserPermissions: []types.UserPermission{
					types.NewUserPermission(
						2,
						0,
						"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
						types.NewPermissions(types.PermissionSetPermissions),
					),
				},
			},
			check: func(ctx sdk.Context) {
				storedUserPermissions := suite.k.GetUserPermissions(ctx, 2, 0, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e")
				suite.Require().Equal(types.NewPermissions(types.PermissionSetPermissions), storedUserPermissions)
			},
		},
		{
			name: "user grants are imported properly",
			genesis: types.GenesisState{
				UserGrants: []types.UserGrant{{
					SubspaceID: 1,
					Granter:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					Allowance:  allowanceAny,
				}},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetUserGrant(ctx, 1, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().True(found)
				suite.Require().Equal(types.UserGrant{
					SubspaceID: 1,
					Granter:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					Grantee:    "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					Allowance:  allowanceAny,
				}, stored)
			},
		},
		{
			name: "group grants are imported properly",
			genesis: types.GenesisState{
				GroupGrants: []types.GroupGrant{{
					SubspaceID: 1,
					Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					GroupID:    1,
					Allowance:  allowanceAny,
				}},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetGroupGrant(ctx, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53", 1)
				suite.Require().True(found)
				suite.Require().Equal(types.GroupGrant{
					SubspaceID: 1,
					Granter:    "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					GroupID:    1,
					Allowance:  allowanceAny,
				}, stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			suite.k.InitGenesis(ctx, tc.genesis)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
