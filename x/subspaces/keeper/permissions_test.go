package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetPermissions() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		user       string
		permission types.Permission
		check      func(ctx sdk.Context)
	}{
		{
			name:       "permission is set properly for user",
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionChangeInfo,
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				permission := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Equal(types.PermissionChangeInfo, permission)
			},
		},
		{
			name: "existing permission is overridden",
			store: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionManageGroups)
			},
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionWrite,
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				permission := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Equal(types.PermissionWrite, permission)
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

			sdkAddr, err := sdk.AccAddressFromBech32(tc.user)
			suite.Require().NoError(err)

			suite.k.SetUserPermissions(ctx, tc.subspaceID, sdkAddr, tc.permission)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_HasPermission() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		user       string
		permission types.Permission
		expResult  bool
	}{
		{
			name:       "subspace not found returns false",
			subspaceID: 1,
			user:       "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			expResult:  false,
		},
		{
			name: "subspace owner returns always true",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			subspaceID: 1,
			user:       "cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
			permission: types.PermissionEverything,
			expResult:  true,
		},
		{
			name: "user with inherited permission returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite|types.PermissionChangeInfo,
				))

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)
			},
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionWrite,
			expResult:  true,
		},
		{
			name: "user with custom permission returns true",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite|types.PermissionChangeInfo,
				))

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, userAddr, types.PermissionManageGroups)
			},
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionManageGroups,
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

			sdkAddr, err := sdk.AccAddressFromBech32(tc.user)
			suite.Require().NoError(err)

			result := suite.k.HasPermission(ctx, tc.subspaceID, sdkAddr, tc.permission)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetUserPermissions() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		subspaceID     uint64
		user           string
		expPermissions types.Permission
	}{
		{
			name:           "not found user returns PermissionNothing",
			subspaceID:     1,
			user:           "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			expPermissions: types.PermissionNothing,
		},
		{
			name: "found user returns the correct permission",
			store: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionWrite|types.PermissionManageGroups)
			},
			subspaceID:     1,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionWrite | types.PermissionManageGroups,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			sdkAddr, err := sdk.AccAddressFromBech32(tc.user)
			suite.Require().NoError(err)

			permission := suite.k.GetUserPermissions(ctx, tc.subspaceID, sdkAddr)
			suite.Require().Equal(tc.expPermissions, permission)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetGroupsInheritedPermissions() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		subspaceID     uint64
		user           string
		expPermissions types.Permission
	}{
		{
			name:           "user in no group returns PermissionNothing",
			subspaceID:     1,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionNothing,
		},
		{
			name: "user inside one group returns that group's permission",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)
			},
			subspaceID:     1,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionWrite,
		},
		{
			name: "user inside multiple groups returns the combination of the various permissions",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					2,
					"Permission group",
					"This is a permissions group",
					types.PermissionSetPermissions|types.PermissionChangeInfo,
				))

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 2, userAddr)
				suite.Require().NoError(err)
			},
			subspaceID:     1,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionWrite | types.PermissionChangeInfo | types.PermissionSetPermissions,
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

			permissions := suite.k.GetGroupsInheritedPermissions(ctx, tc.subspaceID, userAddr)
			suite.Require().Equal(tc.expPermissions, permissions)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RemoveUserPermissions() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "permission is deleted for non existing user",
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				permission := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Equal(types.PermissionNothing, permission)
			},
		},
		{
			name: "permission is deleted for existing user",
			store: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionManageGroups)
			},
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				permission := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Equal(types.PermissionNothing, permission)
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

			sdkAddr, err := sdk.AccAddressFromBech32(tc.user)
			suite.Require().NoError(err)

			suite.k.RemoveUserPermissions(ctx, tc.subspaceID, sdkAddr)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
