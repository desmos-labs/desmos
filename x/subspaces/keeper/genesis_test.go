package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_ExportGenesis() {
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
			expGenesis: types.NewGenesisState(1, nil, nil, nil, nil, nil, nil),
		},
		{
			name: "subspaces are exported correctly",
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
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				))
			},
			expGenesis: types.NewGenesisState(
				3,
				[]types.GenesisSubspace{
					types.NewGenesisSubspace(
						types.NewSubspace(
							1,
							"Test subspace",
							"This is a test subspace",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						),
						1,
					),
					types.NewGenesisSubspace(
						types.NewSubspace(
							2,
							"Another test subspace",
							"This is another test subspace",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
						),
						1,
					),
				},
				nil,
				[]types.UserGroup{
					types.DefaultUserGroup(1),
					types.DefaultUserGroup(2),
				},
				nil,
			),
		},
		{
			name: "permissions are exported correctly",
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

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 2, sdkAddr, types.PermissionSetPermissions)
			},
			expGenesis: types.NewGenesisState(
				3,
				[]types.GenesisSubspace{
					types.NewGenesisSubspace(
						types.NewSubspace(
							2,
							"Another test subspace",
							"This is another test subspace",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
						),
						1,
					),
				},
				[]types.ACLEntry{
					types.NewACLEntry(2, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm", types.PermissionSetPermissions),
				},
				[]types.UserGroup{
					types.DefaultUserGroup(2),
				},
				nil,
			),
		},
		{
			name: "user groups are exported properly",
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
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)

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
					1,
					"Another test group",
					"This is another test group",
					types.PermissionWrite,
				))

				userAddr, err = sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 2, 1, userAddr)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 2, 1, userAddr)
			},
			expGenesis: types.NewGenesisState(
				3,
				[]types.GenesisSubspace{
					types.NewGenesisSubspace(
						types.NewSubspace(
							1,
							"Test subspace",
							"This is a test subspace",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						),
						2,
					),
					types.NewGenesisSubspace(
						types.NewSubspace(
							2,
							"Another test subspace",
							"This is another test subspace",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
						),
						2,
					),
				},
				nil,
				[]types.UserGroup{
					types.DefaultUserGroup(1),
					types.NewUserGroup(
						1,
						1,
						"Test group",
						"This is a test group",
						types.PermissionWrite,
					),
					types.DefaultUserGroup(2),
					types.NewUserGroup(
						2,
						1,
						"Another test group",
						"This is another test group",
						types.PermissionWrite,
					),
				},
				[]types.UserGroupMembersEntry{
					types.NewUserGroupMembersEntry(1, 1, []string{
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					}),
					types.NewUserGroupMembersEntry(2, 1, []string{
						"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					}),
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

			genesis := suite.k.ExportGenesis(ctx)
			suite.Require().Equal(tc.expGenesis, genesis)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_InitGenesis() {
	testCases := []struct {
		name    string
		genesis types.GenesisState
		check   func(ctx sdk.Context)
	}{
		{
			name: "all data is imported properly",
			genesis: types.GenesisState{
				InitialSubspaceID: 3,
				Subspaces: []types.GenesisSubspace{
					types.NewGenesisSubspace(
						types.NewSubspace(
							1,
							"Test subspace",
							"This is a test subspace",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						),
						2,
					),
					types.NewGenesisSubspace(
						types.NewSubspace(
							2,
							"Another test subspace",
							"This is another test subspace",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
							time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
						),
						14,
					),
				},
				ACL: []types.ACLEntry{
					types.NewACLEntry(2, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm", types.PermissionSetPermissions),
				},
				UserGroups: []types.UserGroup{
					types.NewUserGroup(
						1,
						1,
						"Test group",
						"This is a test group",
						types.PermissionWrite,
					),
					types.NewUserGroup(
						2,
						1,
						"Another test group",
						"This is another test group",
						types.PermissionWrite,
					),
					types.NewUserGroup(
						2,
						13,
						"High id test group",
						"This is another test group",
						types.PermissionWrite,
					),
				},
				UserGroupsMembers: []types.UserGroupMembersEntry{
					types.NewUserGroupMembersEntry(1, 1, []string{
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					}),
					types.NewUserGroupMembersEntry(2, 1, []string{
						"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					}),
				},
			},
			check: func(ctx sdk.Context) {
				subspaceID, err := suite.k.GetSubspaceID(ctx)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(3), subspaceID)

				subspaces := suite.k.GetAllSubspaces(ctx)
				suite.Require().Len(subspaces, 2)

				// Check the fist subspace data
				firstSubspaceGroupID, err := suite.k.GetNextGroupID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), firstSubspaceGroupID)

				expectedFirstSubspaceGroups := []types.UserGroup{
					types.DefaultUserGroup(1),
					types.NewUserGroup(
						1,
						1,
						"Test group",
						"This is a test group",
						types.PermissionWrite,
					),
				}
				storedFirstSubspacesGroups := suite.k.GetSubspaceUserGroups(ctx, 1)
				suite.Require().Equal(expectedFirstSubspaceGroups, storedFirstSubspacesGroups)

				groupMembers := suite.k.GetUserGroupMembers(ctx, 1, 1)
				suite.Require().Len(groupMembers, 1)

				// Check the second subspace data
				secondSubspaceGroupID, err := suite.k.GetNextGroupID(ctx, 2)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(14), secondSubspaceGroupID)

				expectedSecondSubspaceGroups := []types.UserGroup{
					types.DefaultUserGroup(2),
					types.NewUserGroup(
						2,
						1,
						"Another test group",
						"This is another test group",
						types.PermissionWrite,
					),
					types.NewUserGroup(
						2,
						13,
						"High id test group",
						"This is another test group",
						types.PermissionWrite,
					),
				}
				storedSecondSubspaceGroups := suite.k.GetSubspaceUserGroups(ctx, 2)
				suite.Require().Equal(expectedSecondSubspaceGroups, storedSecondSubspaceGroups)

				anotherGroupMembers := suite.k.GetUserGroupMembers(ctx, 2, 1)
				suite.Require().Len(anotherGroupMembers, 2)

				// Check user permissions
				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				storedUserPermissions := suite.k.GetUserPermissions(ctx, 2, userAddr)
				suite.Require().Equal(types.PermissionSetPermissions, storedUserPermissions)
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
