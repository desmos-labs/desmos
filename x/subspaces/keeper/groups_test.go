package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetNextGroupID() {
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
				groupID := types.GetGroupIDFromBytes(store.Get(types.NextGroupIDStoreKey(1)))
				suite.Require().Equal(uint32(0), groupID)
			},
		},
		{
			name:       "non-zero group id",
			subspaceID: 1,
			groupID:    5,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				groupID := types.GetGroupIDFromBytes(store.Get(types.NextGroupIDStoreKey(1)))
				suite.Require().Equal(uint32(5), groupID)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			suite.k.SetNextGroupID(ctx, tc.subspaceID, tc.groupID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasNextGroupID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		expResult  bool
	}{
		{
			name:       "not found next group id returns false",
			subspaceID: 1,
			expResult:  false,
		},
		{
			name: "found next group id returns true",
			store: func(ctx sdk.Context) {
				suite.k.SetNextGroupID(ctx, 1, 1)
			},
			subspaceID: 1,
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

			result := suite.k.HasNextGroupID(ctx, tc.subspaceID)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetNextGroupID() {
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
				store.Set(types.NextGroupIDStoreKey(1), types.GetGroupIDBytes(1))
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

			id, err := suite.k.GetNextGroupID(ctx, tc.subspaceID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expID, id)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DeleteNextGroupID() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		check      func(ctx sdk.Context)
	}{
		{
			name:       "non existing next group id is deleted properly",
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasNextGroupID(ctx, 1))
			},
		},
		{
			name: "existing next group id is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SetNextGroupID(ctx, 1, 1)
			},
			subspaceID: 1,
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasNextGroupID(ctx, 1))
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

			suite.k.DeleteNextGroupID(ctx, tc.subspaceID)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (suite *KeeperTestsuite) TestKeeper_SaveUserGroup() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		group types.UserGroup
		check func(ctx sdk.Context)
	}{
		{
			name: "non existing group is stored properly",
			group: types.NewUserGroup(
				1,
				0,
				1,
				"Test group",
				"This is a test group",
				types.NewPermissions(types.PermissionEditSubspace),
			),
			check: func(ctx sdk.Context) {
				group, found := suite.k.GetUserGroup(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				), group)
			},
		},
		{
			name: "existing group is updated properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			group: types.NewUserGroup(
				1,
				0,
				1,
				"Edited test group",
				"This is an edited test group",
				types.NewPermissions(types.PermissionEverything),
			),
			check: func(ctx sdk.Context) {
				group, found := suite.k.GetUserGroup(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserGroup(
					1,
					0,
					1,
					"Edited test group",
					"This is an edited test group",
					types.NewPermissions(types.PermissionEverything),
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
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
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
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			subspaceID: 1,
			groupID:    1,
			expFound:   true,
			expGroup: types.NewUserGroup(
				1,
				0,
				1,
				"Test group",
				"This is a test group",
				types.NewPermissions(types.PermissionEditSubspace),
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
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
			},
			subspaceID: 1,
			groupID:    1,
			check: func(ctx sdk.Context) {
				hasGroup := suite.k.HasUserGroup(ctx, 1, 1)
				suite.Require().False(hasGroup)

				members := suite.k.GetUserGroupMembers(ctx, 1, 1)
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
		check      func(ctx sdk.Context)
	}{
		{
			name: "user is added properly to group",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			subspaceID: 1,
			groupID:    1,
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			check: func(ctx sdk.Context) {
				isMember := suite.k.IsMemberOfGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().True(isMember)
				acc, _ := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().True(suite.ak.HasAccount(ctx, acc))
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

			suite.k.AddUserToGroup(ctx, tc.subspaceID, tc.groupID, tc.user)
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
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
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

			result := suite.k.IsMemberOfGroup(ctx, tc.subspaceID, tc.groupID, tc.user)
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
				isMember := suite.k.IsMemberOfGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().False(isMember)
			},
		},
		{
			name: "existing user is removed properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
			},
			subspaceID: 1,
			groupID:    1,
			user:       "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			check: func(ctx sdk.Context) {
				isMember := suite.k.IsMemberOfGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
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

			suite.k.RemoveUserFromGroup(ctx, tc.subspaceID, tc.groupID, tc.user)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
