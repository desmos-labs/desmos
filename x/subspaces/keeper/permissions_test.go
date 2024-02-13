package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetUserPermissions() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		sectionID   uint32
		user        string
		permissions types.Permissions
		check       func(ctx sdk.Context)
	}{
		{
			name:        "permissions are set properly for user",
			subspaceID:  1,
			sectionID:   1,
			user:        "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			check: func(ctx sdk.Context) {
				permission := suite.k.GetUserPermissions(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().Equal(types.NewPermissions(types.PermissionEditSubspace), permission)
			},
		},
		{
			name: "existing permissions are overridden",
			store: func(ctx sdk.Context) {
				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
					types.NewPermissions(types.PermissionManageGroups),
				)
			},
			subspaceID:  1,
			sectionID:   0,
			user:        "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permissions: types.NewPermissions(types.PermissionDeleteSubspace),
			check: func(ctx sdk.Context) {
				permission := suite.k.GetUserPermissions(ctx, 1, 0, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().Equal(types.NewPermissions(types.PermissionDeleteSubspace), permission)
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

			suite.k.SetUserPermissions(ctx, tc.subspaceID, tc.sectionID, tc.user, tc.permissions)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasPermission() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		user       string
		permission types.Permission
		expResult  bool
	}{
		{
			name:       "subspace not found returns false",
			subspaceID: 1,
			sectionID:  0,
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
					nil,
				))
			},
			subspaceID: 1,
			sectionID:  0,
			user:       "cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
			permission: types.PermissionEverything,
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
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace),
				))
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
					types.NewPermissions(types.PermissionManageGroups),
				)
			},
			subspaceID: 1,
			sectionID:  0,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionManageGroups,
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
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
			},
			subspaceID: 1,
			sectionID:  0,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionEditSubspace,
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

			result := suite.k.HasPermission(ctx, tc.subspaceID, tc.sectionID, tc.user, tc.permission)
			suite.Require().Equal(tc.expResult, result)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserPermissions() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		subspaceID     uint64
		sectionID      uint32
		user           string
		expPermissions types.Permissions
	}{
		{
			name:           "not found user returns PermissionNothing",
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			expPermissions: nil,
		},
		{
			name: "found user returns the correct permission",
			store: func(ctx sdk.Context) {
				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
					types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace),
				)
			},
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			expPermissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace),
		},
		{
			name: "found user inside parent section returns correct permission",
			store: func(ctx sdk.Context) {
				// Store the section tree as follows
				//     root
				//    /   \
				//    A    B
				//    |
				//    C
				suite.k.SaveSection(ctx, types.DefaultSection(1))
				suite.k.SaveSection(ctx, types.NewSection(1, 1, 0, "A", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 2, 0, "B", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 3, 1, "C", ""))

				// Set the permission inside root
				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					types.NewPermissions(types.PermissionManageGroups),
				)
			},
			subspaceID:     1,
			sectionID:      3,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.NewPermissions(types.PermissionManageGroups),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			permission := suite.k.GetUserPermissions(ctx, tc.subspaceID, tc.sectionID, tc.user)
			suite.Require().Equal(tc.expPermissions, permission)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetGroupsInheritedPermissions() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		subspaceID     uint64
		sectionID      uint32
		user           string
		expPermissions types.Permissions
	}{
		{
			name:           "user in no group returns PermissionNothing",
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: nil,
		},
		{
			name: "user inside one group returns that group's permission",
			store: func(ctx sdk.Context) {
				suite.k.SaveSection(ctx, types.DefaultSection(1))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
			},
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.NewPermissions(types.PermissionEditSubspace),
		},
		{
			name: "user inside multiple groups returns the combination of the various permissions",
			store: func(ctx sdk.Context) {
				suite.k.SaveSection(ctx, types.DefaultSection(1))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					2,
					"Permission group",
					"This is a permissions group",
					types.NewPermissions(types.PermissionDeleteSubspace, types.PermissionSetPermissions),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.k.AddUserToGroup(ctx, 1, 2, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
			},
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace, types.PermissionSetPermissions),
		},
		{
			name: "user inside group of ancestor section returns correct permissions",
			store: func(ctx sdk.Context) {
				// Store the section tree as follows
				//  G1 ->   root
				//         /   \
				//  G2 ->  A    B
				//         |
				//         C
				suite.k.SaveSection(ctx, types.DefaultSection(1))
				suite.k.SaveSection(ctx, types.NewSection(1, 1, 0, "A", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 2, 0, "B", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 3, 1, "C", ""))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"G1",
					"",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					2,
					"G2",
					"",
					types.NewPermissions(types.PermissionDeleteSubspace, types.PermissionSetPermissions),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.k.AddUserToGroup(ctx, 1, 2, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
			},
			subspaceID:     1,
			sectionID:      3,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.CombinePermissions(types.PermissionDeleteSubspace, types.PermissionSetPermissions, types.PermissionEditSubspace),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			permissions := suite.k.GetGroupsInheritedPermissions(ctx, tc.subspaceID, tc.sectionID, tc.user)
			suite.Require().Equal(tc.expPermissions, permissions)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUsersWithRootPermissions() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		permissions types.Permissions
		expUsers    []string
	}{
		{
			name:        "subspace not found returns empty slice",
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			expUsers:    nil,
		},
		{
			name: "no users found returns empty slice",
			store: func(ctx sdk.Context) {

			},
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionEditSubspace),
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
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd")
			},
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			expUsers: []string{
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn", // Owner is always included
				"cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd",
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
					nil,
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd",
					types.NewPermissions(types.PermissionEditSubspace),
				)
			},
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			expUsers: []string{
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn", // Owner is always included
				"cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd",
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
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace, types.PermissionSetPermissions),
				))
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1xw69y2z3yf00rgfnly99628gn5c0x7fryyfv5e")

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					2,
					"Another test group",
					"This is a second test group",
					types.NewPermissions(types.PermissionSetPermissions),
				))

				suite.k.AddUserToGroup(ctx, 1, 2, "cosmos1e32dfqu7k9e5wj85cjtalqdd2zs6z7adgswnrn")
				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd",
					types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace),
				)
				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1f3e5dhpg3afanddld0kp6lkayz2qvuetf6hmv3",
					types.NewPermissions(types.PermissionSetPermissions),
				)
			},
			subspaceID:  1,
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			expUsers: []string{
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn", // Owner is always included
				"cosmos1xw69y2z3yf00rgfnly99628gn5c0x7fryyfv5e",
				"cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd",
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

			users := suite.k.GetUsersWithRootPermissions(ctx, tc.subspaceID, tc.permissions)
			suite.Require().Equal(tc.expUsers, users)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveUserPermissions() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "permissions are deleted for non existing user",
			subspaceID: 1,
			sectionID:  0,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			check: func(ctx sdk.Context) {
				permissions := suite.k.GetUserPermissions(ctx, 1, 0, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().Empty(permissions)
			},
		},
		{
			name: "permissionss are deleted for existing user",
			store: func(ctx sdk.Context) {
				suite.k.SetUserPermissions(ctx,
					1,
					1,
					"cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
					types.NewPermissions(types.PermissionManageGroups),
				)
			},
			subspaceID: 1,
			sectionID:  1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			check: func(ctx sdk.Context) {
				permissions := suite.k.GetUserPermissions(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().Empty(permissions)
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

			suite.k.RemoveUserPermissions(ctx, tc.subspaceID, tc.sectionID, tc.user)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
