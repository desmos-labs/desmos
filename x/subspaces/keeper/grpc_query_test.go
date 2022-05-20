package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (suite *KeeperTestsuite) TestQueryServer_Subspaces() {
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
				))
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					2,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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

func (suite *KeeperTestsuite) TestQueryServer_Subspace() {
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

func (suite *KeeperTestsuite) TestQueryServer_UserGroups() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryUserGroupsRequest
		shouldErr bool
		expGroups []types.UserGroup
	}{
		{
			name:      "non existing subspace returns error",
			req:       types.NewQueryUserGroupsRequest(1, nil),
			shouldErr: true,
		},
		{
			name: "existing groups are returned properly",
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

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"First test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					2,
					"Second test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					3,
					"Third test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			req: types.NewQueryUserGroupsRequest(1, &query.PageRequest{
				Offset: 2,
				Limit:  2,
			}),
			shouldErr: false,
			expGroups: []types.UserGroup{
				types.NewUserGroup(
					1,
					2,
					"Second test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				),
				types.NewUserGroup(
					1,
					3,
					"Third test group",
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

func (suite *KeeperTestsuite) TestQueryServer_UserGroup() {
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
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
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

func (suite *KeeperTestsuite) TestQueryServer_UserGroupMembers() {
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
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, userAddr)
				suite.Require().NoError(err)
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

func (suite *KeeperTestsuite) TestQueryServer_UserPermissions() {
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
				))
			},
			req: types.NewQueryUserPermissionsRequest(
				1,
				"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			),
			shouldErr: false,
			expResponse: types.QueryUserPermissionsResponse{
				Permissions: nil,
				Details: []types.PermissionDetail{
					types.NewPermissionDetailGroup(0, nil),
				},
			},
		},
		{
			name: "existing permissions are returned correctly",
			store: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e")
				suite.Require().NoError(err)

				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionManageGroups),
				))
				err = suite.k.AddUserToGroup(ctx, 1, 1, sdkAddr)
				suite.Require().NoError(err)

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					2,
					"Another test group",
					"This is another test group",
					types.NewPermissions(types.PermissionSetPermissions),
				))
				err = suite.k.AddUserToGroup(ctx, 1, 2, sdkAddr)
				suite.Require().NoError(err)

				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.NewPermissions(types.PermissionEditSubspace))
			},
			req: types.NewQueryUserPermissionsRequest(
				1,
				"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
			),
			shouldErr: false,
			expResponse: types.QueryUserPermissionsResponse{
				Permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionManageGroups, types.PermissionSetPermissions),
				Details: []types.PermissionDetail{
					types.NewPermissionDetailUser("cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e", types.NewPermissions(types.PermissionEditSubspace)),
					types.NewPermissionDetailGroup(0, nil),
					types.NewPermissionDetailGroup(1, types.NewPermissions(types.PermissionManageGroups)),
					types.NewPermissionDetailGroup(2, types.NewPermissions(types.PermissionSetPermissions)),
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
