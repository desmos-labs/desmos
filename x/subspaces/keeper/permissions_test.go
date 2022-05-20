package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetPermissions() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		user       string
		permission types.Permissions
		check      func(ctx sdk.Context)
	}{
		{
			name:       "permission is set properly for user",
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.NewPermissions(types.PermissionEditSubspace),
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				permission := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Equal(types.NewPermissions(types.PermissionEditSubspace), permission)
			},
		},
		{
			name: "existing permission is overridden",
			store: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.NewPermissions(types.PermissionManageGroups))
			},
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.NewPermissions(types.PermissionWrite),
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				permission := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Equal(types.NewPermissions(types.PermissionWrite), permission)
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
					types.CombinePermissions(types.PermissionWrite, types.PermissionEditSubspace),
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
					types.CombinePermissions(types.PermissionWrite, types.PermissionEditSubspace),
				))

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, userAddr, types.NewPermissions(types.PermissionManageGroups))
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
		expPermissions types.Permissions
	}{
		{
			name:           "not found user returns PermissionNothing",
			subspaceID:     1,
			user:           "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			expPermissions: nil,
		},
		{
			name: "found user returns the correct permission",
			store: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.CombinePermissions(types.PermissionWrite, types.PermissionManageGroups))
			},
			subspaceID:     1,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.CombinePermissions(types.PermissionWrite, types.PermissionManageGroups),
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
		expPermissions types.Permissions
	}{
		{
			name:           "user in no group returns PermissionNothing",
			subspaceID:     1,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: nil,
		},
		{
			name: "user inside one group returns that group's permission",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionWrite),
				))

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)
			},
			subspaceID:     1,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.NewPermissions(types.PermissionWrite),
		},
		{
			name: "user inside multiple groups returns the combination of the various permissions",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionWrite),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					2,
					"Permission group",
					"This is a permissions group",
					types.CombinePermissions(types.PermissionSetPermissions, types.PermissionEditSubspace),
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
			expPermissions: types.CombinePermissions(types.PermissionWrite, types.PermissionSetPermissions, types.PermissionEditSubspace),
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

func (suite *KeeperTestsuite) TestKeeper_GetUsersWithPermissions() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		permissions types.Permissions
		shouldErr   bool
		expUsers    []string
	}{
		{
			name:        "subspace not found returns empty slice",
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionWrite),
			shouldErr:   false,
			expUsers:    nil,
		},
		{
			name: "no users found returns empty slice",
			store: func(ctx sdk.Context) {

			},
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionWrite),
			shouldErr:   false,
			expUsers:    nil,
		},
		{
			name: "users with permissions inherited from groups are returned properly",
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
					types.NewPermissions(types.PermissionWrite),
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, sdkAddr)
				suite.Require().NoError(err)
			},
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionWrite),
			shouldErr:   false,
			expUsers: []string{
				"cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn", // Owner is always included
			},
		},
		{
			name: "users with individually set permissions are returned properly",
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

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd")
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.NewPermissions(types.PermissionWrite))
			},
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionWrite),
			shouldErr:   false,
			expUsers: []string{
				"cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn", // Owner is always included
			},
		},
		{
			name: "multiple users are returned properly",
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
					types.CombinePermissions(types.PermissionWrite, types.PermissionSetPermissions),
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1xw69y2z3yf00rgfnly99628gn5c0x7fryyfv5e")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, sdkAddr)
				suite.Require().NoError(err)

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					2,
					"Another test group",
					"This is a second test group",
					types.NewPermissions(types.PermissionSetPermissions),
				))

				sdkAddr, err = sdk.AccAddressFromBech32("cosmos1e32dfqu7k9e5wj85cjtalqdd2zs6z7adgswnrn")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 2, sdkAddr)
				suite.Require().NoError(err)

				sdkAddr, err = sdk.AccAddressFromBech32("cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd")
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.CombinePermissions(types.PermissionWrite, types.PermissionEditSubspace))

				sdkAddr, err = sdk.AccAddressFromBech32("cosmos1f3e5dhpg3afanddld0kp6lkayz2qvuetf6hmv3")
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.NewPermissions(types.PermissionEditSubspace))
			},
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionWrite),
			shouldErr:   false,
			expUsers: []string{
				"cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd",
				"cosmos1xw69y2z3yf00rgfnly99628gn5c0x7fryyfv5e",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn", // Owner is always included
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

			users, err := suite.k.GetUsersWithPermission(ctx, tc.subspaceID, tc.permissions)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Len(users, len(tc.expUsers))

				for _, user := range users {
					suite.Require().Contains(tc.expUsers, user.String())
				}
			}
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
				suite.Require().Empty(permission)
			},
		},
		{
			name: "permission is deleted for existing user",
			store: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.NewPermissions(types.PermissionManageGroups))
			},
			subspaceID: 1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				permission := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Empty(permission)
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
