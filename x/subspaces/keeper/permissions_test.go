package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestGetPermissions() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		subspaceID     uint64
		target         string
		expPermissions types.Permission
	}{
		{
			name:           "not found target returns PermissionNothing",
			subspaceID:     1,
			target:         "",
			expPermissions: types.PermissionNothing,
		},
		{
			name: "found user returns the correct permission",
			store: func(ctx sdk.Context) {
				suite.k.SetPermissions(
					ctx,
					1,
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					types.PermissionWrite|types.PermissionAddLink,
				)
			},
			subspaceID:     1,
			target:         "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionWrite | types.PermissionAddLink,
		},
		{
			name: "found group returns the correct permission",
			store: func(ctx sdk.Context) {
				suite.k.SetPermissions(
					ctx, 1, "my_group", types.PermissionChangeInfo|types.PermissionSetPermissions)
			},
			subspaceID:     1,
			target:         "my_group",
			expPermissions: types.PermissionChangeInfo | types.PermissionSetPermissions,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			permission := suite.k.GetPermissions(ctx, tc.subspaceID, tc.target)
			suite.Require().Equal(tc.expPermissions, permission)
		})
	}
}

func (suite *KeeperTestsuite) TestGetGroupsInheritedPermissions() {
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
				suite.k.SaveUserGroup(ctx, 1, "write-group", types.PermissionWrite)

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.Require().NoError(err)
				err = suite.k.AddUserToGroup(ctx, 1, "write-group", userAddr)
				suite.Require().NoError(err)
			},
			subspaceID:     1,
			user:           "cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
			expPermissions: types.PermissionWrite,
		},
		{
			name: "user inside multiple groups returns the combination of the various permissions",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserGroup(ctx, 1,
					"write-group", types.PermissionWrite)
				suite.k.SaveUserGroup(ctx, 1,
					"permission-group", types.PermissionSetPermissions|types.PermissionChangeInfo)

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, "write-group", userAddr)
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, "permission-group", userAddr)
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

func (suite *KeeperTestsuite) TestHasPermission() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		target     string
		permission types.Permission
		expError   bool
		expResult  bool
	}{
		{
			name:       "subspace not found returns error",
			subspaceID: 1,
			expError:   true,
		},
		{
			name: "group with permission returns true",
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
				suite.k.SetPermissions(ctx, 1, "group", types.PermissionWrite)
			},
			subspaceID: 1,
			target:     "group",
			permission: types.PermissionWrite,
			expError:   false,
			expResult:  true,
		},
		{
			name: "group without permission returns false",
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
			target:     "group",
			permission: types.PermissionWrite,
			expError:   false,
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
			target:     "cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
			permission: types.PermissionEverything,
			expError:   false,
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

				suite.k.SaveUserGroup(ctx, 1, "group", types.PermissionWrite|types.PermissionChangeInfo)

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, "group", userAddr)
				suite.Require().NoError(err)
			},
			subspaceID: 1,
			target:     "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionWrite,
			expError:   false,
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

				suite.k.SaveUserGroup(ctx, 1, "group", types.PermissionWrite|types.PermissionChangeInfo)

				userAddr, err := sdk.AccAddressFromBech32("cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, "group", userAddr)
				suite.Require().NoError(err)

				suite.k.SetPermissions(ctx, 1, userAddr.String(), types.PermissionAddLink)
			},
			subspaceID: 1,
			target:     "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionAddLink,
			expError:   false,
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

			result, err := suite.k.HasPermission(ctx, tc.subspaceID, tc.target, tc.permission)
			if tc.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResult, result)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestSetPermissions() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		target     string
		permission types.Permission
		check      func(ctx sdk.Context)
	}{
		{
			name:       "permission is set properly for group",
			subspaceID: 1,
			target:     "group",
			permission: types.PermissionWrite,
			check: func(ctx sdk.Context) {
				permission := suite.k.GetPermissions(ctx, 1, "group")
				suite.Require().Equal(types.PermissionWrite, permission)
			},
		},
		{
			name:       "permission is set properly for user",
			subspaceID: 1,
			target:     "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permission: types.PermissionChangeInfo,
			check: func(ctx sdk.Context) {
				permission := suite.k.GetPermissions(ctx, 1, "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn")
				suite.Require().Equal(types.PermissionChangeInfo, permission)
			},
		},
		{
			name: "existing permission is overridden",
			store: func(ctx sdk.Context) {
				suite.k.SetPermissions(ctx, 1, "group", types.PermissionAddLink)
			},
			subspaceID: 1,
			target:     "group",
			permission: types.PermissionWrite,
			check: func(ctx sdk.Context) {
				permission := suite.k.GetPermissions(ctx, 1, "group")
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

			suite.k.SetPermissions(ctx, tc.subspaceID, tc.target, tc.permission)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
