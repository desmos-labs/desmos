package wasm_test

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
	"github.com/desmos-labs/desmos/v3/x/subspaces/wasm"
)

func (suite *Testsuite) TestSubspacesWasmQuerier_QueryCustom() {
	profilesQuery := profilestypes.QueryProfileRequest{User: ""}
	profilesQueryBz, err := profilesQuery.Marshal()
	suite.NoError(err)
	wrongQueryBz, err := json.Marshal(profilesQueryBz)
	suite.NoError(err)

	testCases := []struct {
		name        string
		request     json.RawMessage
		store       func(ctx sdk.Context)
		shouldErr   bool
		expResponse []byte
	}{
		{
			name:        "wrong request type returns error",
			request:     wrongQueryBz,
			shouldErr:   true,
			expResponse: nil,
		},
		{
			name:    "subspaces request is parsed correctly",
			request: buildSubspacesQueryRequest(suite.cdc, types.NewQuerySubspacesRequest(nil)),
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
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QuerySubspacesResponse{
					Subspaces: []types.Subspace{
						types.NewSubspace(
							1,
							"Test subspace",
							"This is a test subspace",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
							time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						),
					},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1},
				},
			),
		},
		{
			name:    "subspace request is parsed correctly",
			request: buildSubspaceQueryRequest(suite.cdc, types.NewQuerySubspaceRequest(1)),
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
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QuerySubspaceResponse{
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
			),
		},
		{
			name: "user groups query request is parsed correctly",
			request: buildUserGroupsQueryRequest(suite.cdc, types.NewQueryUserGroupsRequest(1, 0, &query.PageRequest{
				Offset: 1,
				Limit:  2,
			})),
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
					0,
					1,
					"First test group",
					"This is a test group",
					types.NewPermissions(poststypes.PermissionWrite),
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryUserGroupsResponse{
					Groups: []types.UserGroup{
						types.NewUserGroup(
							1,
							0,
							1,
							"First test group",
							"This is a test group",
							types.NewPermissions(poststypes.PermissionWrite),
						),
					},
					Pagination: &query.PageResponse{NextKey: nil, Total: 0},
				},
			),
		},
		{
			name:    "user group query request is parsed correctly",
			request: buildUserGroupQueryRequest(suite.cdc, types.NewQueryUserGroupRequest(1, 1)),
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
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(poststypes.PermissionWrite),
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryUserGroupResponse{
					Group: types.NewUserGroup(
						1,
						0,
						1,
						"Test group",
						"This is a test group",
						types.NewPermissions(poststypes.PermissionWrite),
					),
				},
			),
		},
		{
			name:    "user group members query request is parsed correctly",
			request: buildUserGroupMembersQueryRequest(suite.cdc, types.NewQueryUserGroupMembersRequest(1, 1, nil)),
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
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(poststypes.PermissionWrite),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryUserGroupMembersResponse{
					Members:    []string{"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1},
				},
			),
		},
		{
			name: "user permissions query request is parsed correctly",
			request: buildUserPermissionsQueryRequest(suite.cdc, types.NewQueryUserPermissionsRequest(1,
				"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e")),
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
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryUserPermissionsResponse{
					Permissions: nil,
					Details: []types.PermissionDetail{
						types.NewPermissionDetailGroup(1, 0, 0, nil),
					},
				},
			),
		},
	}

	querier := wasm.NewSubspacesWasmQuerier(suite.k, suite.cdc)

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}
			query, err := querier.QueryCustom(ctx, tc.request)
			if tc.shouldErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
			suite.Equal(tc.expResponse, query)
		})
	}

}
