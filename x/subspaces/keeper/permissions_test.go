package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SetUserPermissions() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		user       string
		permission types.Permission
		check      func(ctx sdk.Context)
	}{
		{
			name:       "permission is set properly for user",
			subspaceID: 1,
			sectionID:  1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionChangeInfo,
			check: func(ctx sdk.Context) {
				permission := suite.k.GetUserPermissions(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().Equal(types.PermissionChangeInfo, permission)
			},
		},
		{
			name: "existing permission is overridden",
			store: func(ctx sdk.Context) {
				suite.k.SetUserPermissions(ctx, 1, 0, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn", types.PermissionManageGroups)
			},
			subspaceID: 1,
			sectionID:  0,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionWrite,
			check: func(ctx sdk.Context) {
				permission := suite.k.GetUserPermissions(ctx, 1, 0, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
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

			suite.k.SetUserPermissions(ctx, tc.subspaceID, tc.sectionID, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn", tc.permission)

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
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite|types.PermissionChangeInfo,
				))
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.k.SetUserPermissions(ctx, 1, 0, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn", types.PermissionManageGroups)
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
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite|types.PermissionChangeInfo,
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
			},
			subspaceID: 1,
			sectionID:  0,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionWrite,
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

func (suite *KeeperTestsuite) TestKeeper_GetUserPermissions() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		subspaceID     uint64
		sectionID      uint32
		user           string
		expPermissions types.Permission
	}{
		{
			name:           "not found user returns PermissionNothing",
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			expPermissions: types.PermissionNothing,
		},
		{
			name: "found user returns the correct permission",
			store: func(ctx sdk.Context) {
				suite.k.SetUserPermissions(ctx, 1, 0, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn", types.PermissionWrite|types.PermissionManageGroups)
			},
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionWrite | types.PermissionManageGroups,
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
				suite.k.SetUserPermissions(ctx, 1, 0, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn", types.PermissionManageGroups)
			},
			subspaceID:     1,
			sectionID:      3,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionManageGroups,
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

func (suite *KeeperTestsuite) TestKeeper_GetGroupsInheritedPermissions() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		subspaceID     uint64
		sectionID      uint32
		user           string
		expPermissions types.Permission
	}{
		{
			name:           "user in no group returns PermissionNothing",
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionNothing,
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
					types.PermissionWrite,
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
			},
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionWrite,
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
					types.PermissionWrite,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					2,
					"Permission group",
					"This is a permissions group",
					types.PermissionSetPermissions|types.PermissionChangeInfo,
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.k.AddUserToGroup(ctx, 1, 2, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
			},
			subspaceID:     1,
			sectionID:      0,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionWrite | types.PermissionChangeInfo | types.PermissionSetPermissions,
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

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 1, "G1", "", types.PermissionWrite))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 1, 2, "G2", "", types.PermissionSetPermissions|types.PermissionChangeInfo))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.k.AddUserToGroup(ctx, 1, 2, "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
			},
			subspaceID:     1,
			sectionID:      3,
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

			permissions := suite.k.GetGroupsInheritedPermissions(ctx, tc.subspaceID, tc.sectionID, tc.user)
			suite.Require().Equal(tc.expPermissions, permissions)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetUsersWithPermissions() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		permissions types.Permission
		shouldErr   bool
		expUsers    []string
	}{
		{
			name:        "subspace not found returns empty slice",
			subspaceID:  1,
			permissions: types.PermissionWrite,
			shouldErr:   false,
			expUsers:    nil,
		},
		{
			name: "no users found returns empty slice",
			store: func(ctx sdk.Context) {

			},
			subspaceID:  1,
			permissions: types.PermissionWrite,
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
					0,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd")
			},
			subspaceID:  1,
			permissions: types.PermissionWrite,
			shouldErr:   false,
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
				))

				suite.k.SetUserPermissions(ctx, 1, 0, "cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd", types.PermissionWrite)
			},
			subspaceID:  1,
			permissions: types.PermissionWrite,
			shouldErr:   false,
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
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite|types.PermissionSetPermissions,
				))
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1xw69y2z3yf00rgfnly99628gn5c0x7fryyfv5e")

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					2,
					"Another test group",
					"This is a second test group",
					types.PermissionSetPermissions,
				))

				suite.k.AddUserToGroup(ctx, 1, 2, "cosmos1e32dfqu7k9e5wj85cjtalqdd2zs6z7adgswnrn")
				suite.k.SetUserPermissions(ctx, 1, 0, "cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd", types.PermissionWrite|types.PermissionChangeInfo)
				suite.k.SetUserPermissions(ctx, 1, 0, "cosmos1f3e5dhpg3afanddld0kp6lkayz2qvuetf6hmv3", types.PermissionChangeInfo)
			},
			subspaceID:  1,
			permissions: types.PermissionWrite,
			shouldErr:   false,
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

			users, err := suite.k.GetUsersWithPermission(ctx, tc.subspaceID, tc.permissions)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expUsers, users)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RemoveUserPermissions() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name:       "permission is deleted for non existing user",
			subspaceID: 1,
			sectionID:  0,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			check: func(ctx sdk.Context) {
				permission := suite.k.GetUserPermissions(ctx, 1, 0, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().Equal(types.PermissionNothing, permission)
			},
		},
		{
			name: "permission is deleted for existing user",
			store: func(ctx sdk.Context) {
				suite.k.SetUserPermissions(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn", types.PermissionManageGroups)
			},
			subspaceID: 1,
			sectionID:  1,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			check: func(ctx sdk.Context) {
				permission := suite.k.GetUserPermissions(ctx, 1, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
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

			suite.k.RemoveUserPermissions(ctx, tc.subspaceID, tc.sectionID, tc.user)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
