package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

type mockHooks struct {
	CalledMap map[string]bool
}

func newMockHooks() *mockHooks {
	return &mockHooks{CalledMap: make(map[string]bool)}
}

var _ types.SubspacesHooks = &mockHooks{}

func (h *mockHooks) AfterSubspaceSaved(ctx sdk.Context, subspaceID uint64) {
	h.CalledMap["AfterSubspaceSaved"] = true
}

func (h *mockHooks) AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) {
	h.CalledMap["AfterSubspaceDeleted"] = true
}

func (h *mockHooks) AfterSubspaceSectionSaved(ctx sdk.Context, subspaceID uint64, sectionID uint32) {
	h.CalledMap["AfterSubspaceSectionSaved"] = true
}

func (h *mockHooks) AfterSubspaceSectionDeleted(ctx sdk.Context, subspaceID uint64, sectionID uint32) {
	h.CalledMap["AfterSubspaceSectionDeleted"] = true
}

func (h *mockHooks) AfterSubspaceGroupSaved(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	h.CalledMap["AfterSubspaceGroupSaved"] = true
}

func (h *mockHooks) AfterSubspaceGroupMemberAdded(ctx sdk.Context, subspaceID uint64, groupID uint32, user string) {
	h.CalledMap["AfterSubspaceGroupMemberAdded"] = true
}

func (h *mockHooks) AfterSubspaceGroupMemberRemoved(ctx sdk.Context, subspaceID uint64, groupID uint32, user string) {
	h.CalledMap["AfterSubspaceGroupMemberRemoved"] = true
}

func (h *mockHooks) AfterSubspaceGroupDeleted(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	h.CalledMap["AfterSubspaceGroupDeleted"] = true
}

func (h *mockHooks) AfterUserPermissionSet(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string, permissions types.Permissions) {
	h.CalledMap["AfterUserPermissionSet"] = true
}

func (h *mockHooks) AfterUserPermissionRemoved(ctx sdk.Context, subspaceID uint64, sectionID uint32, user string) {
	h.CalledMap["AfterUserPermissionRemoved"] = true
}

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceSaved() {
	testCases := []struct {
		name     string
		subspace types.Subspace
	}{
		{
			name: "AfterSubspaceSaved is called properly",
			subspace: types.NewSubspace(
				1,
				"Test subspace",
				"This is a test subspace",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.SaveSubspace(ctx, tc.subspace)

			suite.Require().True(hooks.CalledMap["AfterSubspaceSaved"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
	}{
		{
			name: "AfterSubspaceDeleted is called properly",
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
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.DeleteSubspace(ctx, tc.subspaceID)

			suite.Require().True(hooks.CalledMap["AfterSubspaceDeleted"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceSectionSaved() {
	testCases := []struct {
		name    string
		section types.Section
	}{
		{
			name: "non existing section is stored properly",
			section: types.NewSection(
				1,
				1,
				0,
				"Test section",
				"This is a test section",
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.SaveSection(ctx, tc.section)

			suite.Require().True(hooks.CalledMap["AfterSubspaceSectionSaved"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceSectionDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
	}{
		{
			name: "AfterSubspaceSectionDeleted is called properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test edited section",
				))
			},
			subspaceID: 1,
			sectionID:  1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.DeleteSection(ctx, tc.subspaceID, tc.sectionID)

			suite.Require().True(hooks.CalledMap["AfterSubspaceSectionDeleted"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceGroupSaved() {
	testCases := []struct {
		name  string
		group types.UserGroup
	}{
		{
			name: "AfterSubspaceGroupSaved is called properly",
			group: types.NewUserGroup(
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

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.SaveUserGroup(ctx, tc.group)

			suite.Require().True(hooks.CalledMap["AfterSubspaceGroupSaved"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceGroupDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
	}{
		{
			name: "AfterSubspaceGroupDeleted is called properly",
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
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.DeleteUserGroup(ctx, tc.subspaceID, tc.groupID)

			suite.Require().True(hooks.CalledMap["AfterSubspaceGroupDeleted"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceGroupMemberAdded() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name: "AfterSubspaceGroupMemberAdded is called properly",
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
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.AddUserToGroup(ctx, tc.subspaceID, tc.groupID, tc.user)

			suite.Require().True(hooks.CalledMap["AfterSubspaceGroupMemberAdded"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterSubspaceGroupMemberRemoved() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		groupID    uint32
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name: "AfterSubspaceGroupMemberRemoved is called properly",
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
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.RemoveUserFromGroup(ctx, tc.subspaceID, tc.groupID, tc.user)

			suite.Require().True(hooks.CalledMap["AfterSubspaceGroupMemberRemoved"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterUserPermissionSet() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  uint64
		sectionID   uint32
		user        string
		permissions types.Permissions
	}{
		{
			name:        "AfterUserPermissionSet is called properly",
			subspaceID:  1,
			sectionID:   1,
			user:        "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
			permissions: types.NewPermissions(types.PermissionEditSubspace),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.SetUserPermissions(ctx, tc.subspaceID, tc.sectionID, tc.user, tc.permissions)

			suite.Require().True(hooks.CalledMap["AfterUserPermissionSet"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterUserPermissionRemoved() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		sectionID  uint32
		user       string
	}{
		{
			name:       "AfterUserPermissionRemoved is called properly",
			subspaceID: 1,
			sectionID:  0,
			user:       "cosmos1fz49f2njk28ue8geqm63g4zzsm97lahqa9vmwn",
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiSubspacesHooks(hooks))

			suite.k.RemoveUserPermissions(ctx, tc.subspaceID, tc.sectionID, tc.user)

			suite.Require().True(hooks.CalledMap["AfterUserPermissionRemoved"])
		})
	}
}
