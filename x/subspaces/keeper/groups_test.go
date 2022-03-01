package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetGroupID() {
	testCases := []struct {
		name       string
		subspaceID uint64
		groupID    uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "zero group id",
			subspaceID: 1,
			groupID:    0,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				groupID := types.GetGroupIDFromBytes(store.Get(types.GroupIDStoreKey(1)))
				suite.Require().Equal(uint32(0), groupID)
			},
		},
		{
			name:       "non-zero group id",
			subspaceID: 1,
			groupID:    5,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				groupID := types.GetGroupIDFromBytes(store.Get(types.GroupIDStoreKey(1)))
				suite.Require().Equal(uint32(5), groupID)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			suite.k.SetGroupID(ctx, tc.subspaceID, tc.groupID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetGroupID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		shouldErr  bool
		expID      uint32
	}{
		{
			name:       "group id not set",
			subspaceID: 1,
			shouldErr:  true,
		},
		{
			name: "group id set",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.GroupIDStoreKey(1), types.GetGroupIDBytes(1))
			},
			subspaceID: 1,
			shouldErr:  false,
			expID:      1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			id, err := suite.k.GetGroupID(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expID, id)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SaveUserGroup() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		group       types.UserGroup
		permissions types.Permission
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing group is stored properly",
			group: types.NewUserGroup(
				1,
				1,
				"Test group",
				"This is a test group",
				types.PermissionWrite,
			),
			permissions: types.PermissionWrite,
			check: func(ctx sdk.Context) {
				group, found := suite.k.GetUserGroup(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				), group)
			},
		},
		{
			name: "existing group is updated properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			group: types.NewUserGroup(
				1,
				1,
				"Edited test group",
				"This is an edited test group",
				types.PermissionChangeInfo,
			),
			permissions: types.PermissionManageGroups,
			check: func(ctx sdk.Context) {
				group, found := suite.k.GetUserGroup(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserGroup(
					1,
					1,
					"Edited test group",
					"This is an edited test group",
					types.PermissionChangeInfo,
				), group)
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

			suite.k.SaveUserGroup(ctx, tc.group)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasUserGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		expResult  bool
	}{
		{
			name:       "not found group returns false",
			subspaceID: 1,
			groupID:    1,
			expResult:  false,
		},
		{
			name: "found group returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			subspaceID: 1,
			groupID:    1,
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

			result := suite.k.HasUserGroup(ctx, tc.subspaceID, tc.groupID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetUserGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		expFound   bool
		expGroup   types.UserGroup
	}{
		{
			name:       "not found group returns false",
			subspaceID: 1,
			groupID:    1,
			expFound:   false,
		},
		{
			name: "found group returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			subspaceID: 1,
			groupID:    1,
			expFound:   true,
			expGroup: types.NewUserGroup(
				1,
				1,
				"Test group",
				"This is a test group",
				types.PermissionWrite,
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

			group, found := suite.k.GetUserGroup(ctx, tc.subspaceID, tc.groupID)
			suite.Require().Equal(tc.expFound, found)
			if tc.expFound {
				suite.Require().Equal(tc.expGroup, group)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteUserGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing group is deleted properly",
			subspaceID: 1,
			groupID:    1,
			check: func(ctx sdk.Context) {
				hasGroup := suite.k.HasUserGroup(ctx, 1, 1)
				suite.Require().False(hasGroup)
			},
		},
		{
			name: "existing group is deleted properly and members are cleared",
			store: func(ctx sdk.Context) {
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
				suite.Require().NoError(err)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)
			},
			subspaceID: 1,
			groupID:    1,
			check: func(ctx sdk.Context) {
				hasGroup := suite.k.HasUserGroup(ctx, 1, 1)
				suite.Require().False(hasGroup)

				members := suite.k.GetGroupMembers(ctx, 1, 1)
				suite.Require().Empty(members)
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

			suite.k.DeleteUserGroup(ctx, tc.subspaceID, tc.groupID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_AddUserToGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		user       string
		shouldErr  bool
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing group returns error",
			subspaceID: 1,
			groupID:    1,
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			shouldErr:  true,
		},
		{
			name: "user is added properly to group",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			subspaceID: 1,
			groupID:    1,
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			shouldErr:  false,
			check: func(ctx sdk.Context) {
				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				isMember := suite.k.IsMemberOfGroup(ctx, 1, 1, userAddr)
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

			err = suite.k.AddUserToGroup(ctx, tc.subspaceID, tc.groupID, userAddr)
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

func (suite *KeeperTestsuite) TestKeeper_IsMemberOfGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		user       string
		expResult  bool
	}{
		{
			name:       "group 0 always returns true",
			subspaceID: 1,
			groupID:    0,
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			expResult:  true,
		},
		{
			name:       "not being part of group returns false",
			subspaceID: 1,
			groupID:    1,
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			expResult:  false,
		},
		{
			name: "being part of group returns true",
			store: func(ctx sdk.Context) {

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
				suite.Require().NoError(err)
			},
			subspaceID: 1,
			groupID:    1,
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

			result := suite.k.IsMemberOfGroup(ctx, tc.subspaceID, tc.groupID, userAddr)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RemoveUserFromGroup() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing user is removed properly",
			subspaceID: 1,
			groupID:    1,
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			check: func(ctx sdk.Context) {
				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				isMember := suite.k.IsMemberOfGroup(ctx, 1, 1, userAddr)
				suite.Require().False(isMember)
			},
		},
		{
			name: "existing user is removed properly",
			store: func(ctx sdk.Context) {
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
				suite.Require().NoError(err)
			},
			subspaceID: 1,
			groupID:    1,
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			check: func(ctx sdk.Context) {
				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				isMember := suite.k.IsMemberOfGroup(ctx, 1, 1, userAddr)
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

			suite.k.RemoveUserFromGroup(ctx, tc.subspaceID, tc.groupID, userAddr)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
