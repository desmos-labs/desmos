package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v6/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
)

func (suite *KeeperTestSuite) TestQueryServer_Subspaces() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		req          *types.QuerySubspacesRequest
		expSubspaces []types.Subspace
	}{

		{
			name: "invalid pagination returns empty slice",
			req: types.NewQuerySubspacesRequest(&query.PageRequest{
				Limit:  1,
				Offset: 1,
			}),
			expSubspaces: nil,
		},
		{
			name: "valid pagination returns result properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			req: &types.QuerySubspacesRequest{
				Pagination: &query.PageRequest{
					Offset:     1,
					Limit:      1,
					CountTotal: true,
				},
			},
			expSubspaces: []types.Subspace{
				types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				),
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

			res, err := suite.k.Subspaces(sdk.WrapSDKContext(ctx), tc.req)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expSubspaces, res.Subspaces)
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Subspace() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		request     *types.QuerySubspaceRequest
		shouldErr   bool
		expResponse *types.QuerySubspaceResponse
	}{
		{
			name:      "not found subspace returns error",
			request:   types.NewQuerySubspaceRequest(1),
			shouldErr: true,
		},
		{
			name: "found subspace is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			request:   types.NewQuerySubspaceRequest(1),
			shouldErr: false,
			expResponse: &types.QuerySubspaceResponse{
				Subspace: types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				),
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

			response, err := suite.k.Subspace(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, response)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Sections() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QuerySectionsRequest
		shouldErr   bool
		expSections []types.Section
	}{
		{
			name:      "not found subspace returns error",
			req:       types.NewQuerySectionsRequest(1, nil),
			shouldErr: true,
		},
		{
			name: "request without pagination returns correct data",
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

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					2,
					0,
					"Another test section",
					"Another test section",
				))
			},
			req:       types.NewQuerySectionsRequest(1, nil),
			shouldErr: false,
			expSections: []types.Section{
				types.DefaultSection(1),
				types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				),
				types.NewSection(
					1,
					2,
					0,
					"Another test section",
					"Another test section",
				),
			},
		},
		{
			name: "request with pagination works properly",
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

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					2,
					0,
					"Another test section",
					"Another test section",
				))
			},
			req: types.NewQuerySectionsRequest(1, &query.PageRequest{
				Limit:  1,
				Offset: 1, // Skip the default section
			}),
			shouldErr: false,
			expSections: []types.Section{
				types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				),
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

			res, err := suite.k.Sections(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expSections, res.Sections)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Section() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		req        *types.QuerySectionRequest
		shouldErr  bool
		expSection types.Section
	}{
		{
			name:      "non existing subspace returns error",
			req:       types.NewQuerySectionRequest(1, 1),
			shouldErr: true,
		},
		{
			name: "non existing section returns error",
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
			req:       types.NewQuerySectionRequest(1, 1),
			shouldErr: true,
		},
		{
			name: "existing section is returned properly",
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

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"Test section",
				))
			},
			req:       types.NewQuerySectionRequest(1, 1),
			shouldErr: false,
			expSection: types.NewSection(
				1,
				1,
				0,
				"Test section",
				"Test section",
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

			res, err := suite.k.Section(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expSection, res.Section)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserGroups() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryUserGroupsRequest
		shouldErr bool
		expGroups []types.UserGroup
	}{
		{
			name:      "non existing subspace returns error",
			req:       types.NewQueryUserGroupsRequest(1, 0, nil),
			shouldErr: true,
		},
		{
			name: "existing groups are returned properly with section id ≠ 0",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"First test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					2,
					"Second test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					3,
					"Third test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			req:       types.NewQueryUserGroupsRequest(1, 1, nil),
			shouldErr: false,
			expGroups: []types.UserGroup{
				types.NewUserGroup(
					1,
					1,
					2,
					"Second test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				),
			},
		},
		{
			name: "existing groups are returned properly with section id = 0",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"First test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					2,
					"Second test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					3,
					"Third test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			req: types.NewQueryUserGroupsRequest(1, 0, &query.PageRequest{
				Offset: 2, // Skip default user group and the first custom one
				Limit:  2,
			}),
			shouldErr: false,
			expGroups: []types.UserGroup{
				types.NewUserGroup(
					1,
					0,
					3,
					"Third test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				),
				types.NewUserGroup(
					1,
					1,
					2,
					"Second test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				),
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

			res, err := suite.k.UserGroups(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expGroups, res.Groups)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserGroup() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryUserGroupRequest
		shouldErr bool
		expGroup  types.UserGroup
	}{
		{
			name:      "not found group returns error",
			req:       types.NewQueryUserGroupRequest(1, 1),
			shouldErr: true,
		},
		{
			name: "found group is returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
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
			},
			req:       types.NewQueryUserGroupRequest(1, 1),
			shouldErr: false,
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

			res, err := suite.k.UserGroup(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expGroup, res.Group)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserGroupMembers() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		req        *types.QueryUserGroupMembersRequest
		shouldErr  bool
		expMembers []string
	}{
		{
			name:      "non existing subspace returns error",
			req:       types.NewQueryUserGroupMembersRequest(1, 1, nil),
			shouldErr: true,
		},
		{
			name: "non existing group returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			req:       types.NewQueryUserGroupMembersRequest(1, 1, nil),
			shouldErr: true,
		},
		{
			name: "existing group members are returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
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

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
			},
			req: types.NewQueryUserGroupMembersRequest(1, 1, &query.PageRequest{
				Offset: 1,
				Limit:  1,
			}),
			shouldErr:  false,
			expMembers: []string{"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.UserGroupMembers(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expMembers, res.Members)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserPermissions() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryUserPermissionsRequest
		shouldErr   bool
		expResponse types.QueryUserPermissionsResponse
	}{
		{
			name:      "not found subspace returns error",
			shouldErr: true,
			req: types.NewQueryUserPermissionsRequest(
				1,
				types.RootSectionID,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
		},
		{
			name: "not found user returns the permission from the default group",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			req: types.NewQueryUserPermissionsRequest(
				1,
				types.RootSectionID,
				"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			),
			shouldErr: false,
			expResponse: types.QueryUserPermissionsResponse{
				Permissions: nil,
				Details: []types.PermissionDetail{
					types.NewPermissionDetailGroup(1, 0, 0, nil),
				},
			},
		},
		{
			name: "existing permissions are returned correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
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
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e")

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					2,
					"Another test group",
					"This is another test group",
					types.NewPermissions(types.PermissionSetPermissions),
				))
				suite.k.AddUserToGroup(ctx,
					1,
					2,
					"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
				)

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
					types.NewPermissions(types.PermissionDeleteSubspace),
				)
			},
			req: types.NewQueryUserPermissionsRequest(
				1,
				types.RootSectionID,
				"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			),
			shouldErr: false,
			expResponse: types.QueryUserPermissionsResponse{
				Permissions: types.CombinePermissions(types.PermissionDeleteSubspace, types.PermissionEditSubspace, types.PermissionSetPermissions),
				Details: []types.PermissionDetail{
					types.NewPermissionDetailUser(1, 0, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e", types.NewPermissions(types.PermissionDeleteSubspace)),
					types.NewPermissionDetailGroup(1, 0, 0, nil),
					types.NewPermissionDetailGroup(1, 0, 1, types.NewPermissions(types.PermissionEditSubspace)),
					types.NewPermissionDetailGroup(1, 0, 2, types.NewPermissions(types.PermissionSetPermissions)),
				},
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

			res, err := suite.k.UserPermissions(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse.Permissions, res.Permissions)
				suite.Require().Equal(tc.expResponse.Details, res.Details)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserAllowances() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryUserAllowancesRequest
		shouldErr bool
		expGrants []types.Grant
	}{
		{
			name:      "invalid subspace id returns error",
			req:       types.NewQueryUserAllowancesRequest(0, "", nil),
			shouldErr: true,
		},
		{
			name:      "not found subspace returns error",
			req:       types.NewQueryUserAllowancesRequest(1, "", nil),
			shouldErr: true,
		},
		{
			name: "user grants query without grantee returns the correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1, "test", "test", "owner", "treasury", "creator", time.Now(), nil))

				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))

				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewUserGrantee("cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			req:       types.NewQueryUserAllowancesRequest(1, "", nil),
			shouldErr: false,
			expGrants: []types.Grant{
				types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				),
				types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewUserGrantee("cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				),
			},
		},
		{
			name: "valid query returns the correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1, "test", "test", "owner", "treasury", "creator", time.Now(), nil))

				suite.k.SaveGrant(ctx, types.NewGrant(1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))

				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewUserGrantee("cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			req:       types.NewQueryUserAllowancesRequest(1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5", nil),
			shouldErr: false,
			expGrants: []types.Grant{
				types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				),
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

			res, err := suite.k.UserAllowances(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expGrants, res.Grants)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_GroupAllowances() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryGroupAllowancesRequest
		shouldErr bool
		expGrants []types.Grant
	}{
		{
			name:      "invalid subspace id returns error",
			req:       types.NewQueryGroupAllowancesRequest(0, 1, nil),
			shouldErr: true,
		},
		{
			name:      "not found subspace returns error",
			req:       types.NewQueryGroupAllowancesRequest(1, 1, nil),
			shouldErr: true,
		},
		{
			name: "group grants query without group id returns the correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1, "test", "test", "owner", "treasury", "creator", time.Now(), nil))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 1, "test", "tets", nil))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 2, "test", "tets", nil))

				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewGroupGrantee(2),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			req:       types.NewQueryGroupAllowancesRequest(1, 0, nil),
			shouldErr: false,
			expGrants: []types.Grant{
				types.NewGrant(1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				),
				types.NewGrant(1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewGroupGrantee(2),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				),
			},
		},
		{
			name: "valid group grants query returns the correct data",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(1, "test", "test", "owner", "treasury", "creator", time.Now(), nil))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 1, "test", "tets", nil))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(1, 0, 2, "test", "tets", nil))

				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
				suite.k.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewGroupGrantee(2),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				))
			},
			req:       types.NewQueryGroupAllowancesRequest(1, 1, nil),
			shouldErr: false,
			expGrants: []types.Grant{
				types.NewGrant(
					1,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(100)))},
				),
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

			res, err := suite.k.GroupAllowances(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expGrants, res.Grants)
			}
		})
	}
}
