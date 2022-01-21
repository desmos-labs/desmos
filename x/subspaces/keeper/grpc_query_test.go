package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

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

func (suite *KeeperTestsuite) TestQueryServer_UserGroups() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		req       *types.QueryUserGroupsRequest
		expGroups []string
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

				suite.k.SaveUserGroup(ctx, 1, "1-group", types.PermissionWrite)
				suite.k.SaveUserGroup(ctx, 1, "2-group", types.PermissionWrite)
				suite.k.SaveUserGroup(ctx, 1, "3-group", types.PermissionWrite)
			},
			req: types.NewQueryUserGroupsRequest(1, &query.PageRequest{
				Offset: 1,
				Limit:  2,
			}),
			shouldErr: false,
			expGroups: []string{"2-group", "3-group"},
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
			req:       types.NewQueryUserGroupMembersRequest(1, "group", nil),
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
			req:       types.NewQueryUserGroupMembersRequest(1, "group", nil),
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

				suite.k.SaveUserGroup(ctx, 1, "group", types.PermissionWrite)

				userAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, "group", userAddr)
				suite.Require().NoError(err)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, "group", userAddr)
				suite.Require().NoError(err)

				userAddr, err = sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, "group", userAddr)
				suite.Require().NoError(err)
			},
			req: types.NewQueryUserGroupMembersRequest(1, "group", &query.PageRequest{
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

func (suite *KeeperTestsuite) TestQueryServer_Permissions() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		req            *types.QueryPermissionsRequest
		shouldErr      bool
		expPermissions types.Permission
	}{
		{
			name:      "not found subspace returns error",
			shouldErr: true,
			req:       types.NewQueryPermissionsRequest(1, "group"),
		},
		{
			name: "not found target returns PermissionNothing",
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
			req:            types.NewQueryPermissionsRequest(1, "group"),
			shouldErr:      false,
			expPermissions: types.PermissionNothing,
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
				))
				suite.k.SetPermissions(ctx, 1, "group", types.PermissionWrite)
			},
			req:            types.NewQueryPermissionsRequest(1, "group"),
			shouldErr:      false,
			expPermissions: types.PermissionWrite,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.Permissions(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expPermissions, res.Permissions)
			}
		})
	}
}
