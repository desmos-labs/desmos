package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_ExportGenesis() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "subspaces are exported correctly",
			store: func(ctx sdk.Context) {
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
				nil,
				nil,
			),
		},
		{
			name: "permissions are exported correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetPermissions(ctx, 1, "group", types.PermissionWrite)

				suite.k.SaveSubspace(ctx, types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SetPermissions(ctx, 2, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm", types.PermissionSetPermissions)
			},
			expGenesis: types.NewGenesisState(
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
				nil,
				[]types.ACLEntry{
					types.NewACLEntry(1, "group", types.PermissionWrite),
					types.NewACLEntry(2, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm", types.PermissionSetPermissions),
				},
			),
		},
		{
			name: "user groups are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, 1, "group", types.PermissionWrite)

				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 1, "group", userAddr)

				suite.k.SaveSubspace(ctx, types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveUserGroup(ctx, 2, "another-group", types.PermissionManageGroups)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 2, "another-group", userAddr)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 2, "another-group", userAddr)
			},
			expGenesis: types.NewGenesisState(
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
				[]types.UserGroup{
					types.NewUserGroup(1, "group", []string{
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					}),
					types.NewUserGroup(2, "another-group", []string{
						"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					}),
				},
				[]types.ACLEntry{
					types.NewACLEntry(1, "group", types.PermissionWrite),
					types.NewACLEntry(2, "another-group", types.PermissionManageGroups),
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
				ACL: []types.ACLEntry{
					types.NewACLEntry(1, "group", types.PermissionWrite),
					types.NewACLEntry(2, "another-group", types.PermissionManageGroups),
				},
				UserGroups: []types.UserGroup{
					types.NewUserGroup(1, "group", []string{
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					}),
					types.NewUserGroup(2, "another-group", []string{
						"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					}),
				},
			},
			check: func(ctx sdk.Context) {
				subspaces := suite.k.GetAllSubspaces(ctx)
				suite.Require().Len(subspaces, 2)

				var firstSubspaceGroups = 0
				suite.k.IterateSubspaceGroups(ctx, 1, func(index int64, groupName string) (stop bool) {
					firstSubspaceGroups += 1
					return false
				})
				suite.Require().Equal(firstSubspaceGroups, 1)

				groupPermissions := suite.k.GetPermissions(ctx, 1, "group")
				suite.Require().Equal(types.PermissionWrite, groupPermissions)

				var groupMembers = 0
				suite.k.IterateGroupMembers(ctx, 1, "group", func(index int64, member sdk.AccAddress) (stop bool) {
					groupMembers += 1
					return false
				})
				suite.Require().Equal(1, groupMembers)

				var secondSubspaceGroups = 0
				suite.k.IterateSubspaceGroups(ctx, 2, func(index int64, groupName string) (stop bool) {
					secondSubspaceGroups += 1
					return false
				})
				suite.Require().Equal(secondSubspaceGroups, 1)

				var anotherGroupMembers = 0
				suite.k.IterateGroupMembers(ctx, 2, "another-group", func(index int64, member sdk.AccAddress) (stop bool) {
					anotherGroupMembers += 1
					return false
				})
				suite.Require().Equal(2, anotherGroupMembers)

				anotherGroupPermissions := suite.k.GetPermissions(ctx, 2, "another-group")
				suite.Require().Equal(types.PermissionManageGroups, anotherGroupPermissions)
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
