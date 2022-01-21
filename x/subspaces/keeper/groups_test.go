package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_HasUserGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupName  string
		expResult  bool
	}{
		{
			name:       "not found group returns false",
			subspaceID: 1,
			groupName:  "group",
			expResult:  false,
		},
		{
			name: "found group returns true",
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
			},
			subspaceID: 1,
			groupName:  "group",
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			result := suite.k.HasUserGroup(ctx, tc.subspaceID, tc.groupName)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_SaveUserGroup() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		groupName   string
		permissions types.Permission
		check       func(ctx sdk.Context)
	}{
		{
			name:        "non existing group is stored properly",
			subspaceID:  1,
			groupName:   "group",
			permissions: types.PermissionWrite,
			check: func(ctx sdk.Context) {
				hasGroup := suite.k.HasUserGroup(ctx, 1, "group")
				suite.Require().True(hasGroup)

				permissions := suite.k.GetPermissions(ctx, 1, "group")
				suite.Require().Equal(types.PermissionWrite, permissions)
			},
		},
		{
			name: "existing group permissions are updated properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, 1, "group", types.PermissionWrite)
			},
			subspaceID:  1,
			groupName:   "group",
			permissions: types.PermissionManageGroups,
			check: func(ctx sdk.Context) {
				hasGroup := suite.k.HasUserGroup(ctx, 1, "group")
				suite.Require().True(hasGroup)

				permissions := suite.k.GetPermissions(ctx, 1, "group")
				suite.Require().Equal(types.PermissionManageGroups, permissions)
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

			suite.k.SaveUserGroup(ctx, tc.subspaceID, tc.groupName, tc.permissions)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteUserGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupName  string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing group is deleted properly",
			subspaceID: 1,
			groupName:  "group",
			check: func(ctx sdk.Context) {
				hasGroup := suite.k.HasUserGroup(ctx, 1, "group")
				suite.Require().False(hasGroup)
			},
		},
		{
			name: "existing group is deleted properly, members are deleted and permissions are cleared",
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
				suite.Require().NoError(err)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 1, "group", userAddr)
				suite.Require().NoError(err)
			},
			subspaceID: 1,
			groupName:  "group",
			check: func(ctx sdk.Context) {
				hasGroup := suite.k.HasUserGroup(ctx, 1, "group")
				suite.Require().False(hasGroup)

				var members []sdk.AccAddress
				suite.k.IterateGroupMembers(ctx, 1, "group", func(index int64, member sdk.AccAddress) (stop bool) {
					members = append(members, member)
					return false
				})
				suite.Require().Empty(members)

				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.PermissionStoreKey(1, "group")))
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

			suite.k.DeleteUserGroup(ctx, tc.subspaceID, tc.groupName)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_IsMemberOfGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupName  string
		user       string
		expResult  bool
	}{
		{
			name:       "not being part of group returns false",
			subspaceID: 1,
			groupName:  "group",
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			expResult:  false,
		},
		{
			name: "being part of group returns true",
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
				suite.Require().NoError(err)
			},
			subspaceID: 1,
			groupName:  "group",
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			expResult:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			userAddr, err := sdk.AccAddressFromBech32(tc.user)
			suite.Require().NoError(err)

			result := suite.k.IsMemberOfGroup(ctx, tc.subspaceID, tc.groupName, userAddr)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_AddUserToGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupName  string
		user       string
		shouldErr  bool
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing group returns error",
			subspaceID: 1,
			groupName:  "group",
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			shouldErr:  true,
		},
		{
			name: "user is added properly to group",
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
			},
			subspaceID: 1,
			groupName:  "group",
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			shouldErr:  false,
			check: func(ctx sdk.Context) {
				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				isMember := suite.k.IsMemberOfGroup(ctx, 1, "group", userAddr)
				suite.Require().True(isMember)
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

			userAddr, err := sdk.AccAddressFromBech32(tc.user)
			suite.Require().NoError(err)

			err = suite.k.AddUserToGroup(ctx, tc.subspaceID, tc.groupName, userAddr)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RemoveUserFromGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupName  string
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing user is removed properly",
			subspaceID: 1,
			groupName:  "group",
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			check: func(ctx sdk.Context) {
				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				isMember := suite.k.IsMemberOfGroup(ctx, 1, "group", userAddr)
				suite.Require().False(isMember)
			},
		},
		{
			name: "existing user is removed properly",
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
				suite.Require().NoError(err)
			},
			subspaceID: 1,
			groupName:  "group",
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			check: func(ctx sdk.Context) {
				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				isMember := suite.k.IsMemberOfGroup(ctx, 1, "group", userAddr)
				suite.Require().False(isMember)
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

			userAddr, err := sdk.AccAddressFromBech32(tc.user)
			suite.Require().NoError(err)

			suite.k.RemoveUserFromGroup(ctx, tc.subspaceID, tc.groupName, userAddr)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
