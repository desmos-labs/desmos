package types_test

import (
	"testing"
	"time"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/stretchr/testify/require"
)

func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name      string
		genesis   *types.GenesisState
		shouldErr bool
	}{
		{
			name: "invalid initial subspace id returns error",
			genesis: types.NewGenesisState(
				0,
				nil,
				nil,
				nil,
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "duplicated subspace data returns error",
			genesis: types.NewGenesisState(1, []types.SubspaceData{
				types.NewSubspaceData(1, 1, 1),
				types.NewSubspaceData(1, 1, 1),
			}, nil, nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "invalid subspace data returns error",
			genesis: types.NewGenesisState(1, []types.SubspaceData{
				types.NewSubspaceData(1, 1, 0),
			}, nil, nil, nil, nil),
		},
		{
			name: "duplicated subspace returns error",
			genesis: types.NewGenesisState(1, nil, []types.Subspace{
				types.NewSubspace(
					1,
					"This is a test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
				types.NewSubspace(
					1,
					"Another test subspace",
					"This is another test subspace",
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
					"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			}, nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			genesis: types.NewGenesisState(2, nil, []types.Subspace{
				types.NewSubspace(
					0,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				),
			}, nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "duplicated user permission returns error",
			genesis: types.NewGenesisState(1, nil, nil, []types.UserPermission{
				types.NewUserPermission(1, 1, "cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd", types.PermissionWrite),
				types.NewUserPermission(1, 1, "cosmos15p3m7a93luselt80ffzpf4jwtn9ama34ray0nd", types.PermissionSetPermissions),
			}, nil, nil),
			shouldErr: true,
		},
		{
			name: "invalid user permission returns error",
			genesis: types.NewGenesisState(1, nil, nil, []types.UserPermission{
				types.NewUserPermission(0, 0, "group", types.PermissionWrite),
			}, nil, nil),
			shouldErr: true,
		},
		{
			name: "duplicated group returns error",
			genesis: types.NewGenesisState(1, nil, nil, nil, []types.UserGroup{
				types.NewUserGroup(
					1,
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				),
				types.NewUserGroup(
					1,
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				),
			}, nil),
			shouldErr: true,
		},
		{
			name: "invalid group returns error",
			genesis: types.NewGenesisState(1, nil, nil, nil, []types.UserGroup{
				types.NewUserGroup(
					1,
					1,
					0,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				),
			}, nil),
			shouldErr: true,
		},
		{
			name: "duplicated group members entry returns error",
			genesis: types.NewGenesisState(1, nil, nil, nil, nil, []types.UserGroupMembersEntry{
				types.NewUserGroupMembersEntry(1, 1, nil),
				types.NewUserGroupMembersEntry(1, 1, nil),
			}),
			shouldErr: true,
		},
		{
			name: "invalid group members entry returns error",
			genesis: types.NewGenesisState(1, nil, nil, nil, nil, []types.UserGroupMembersEntry{
				types.NewUserGroupMembersEntry(1, 0, nil),
			},
			),
			shouldErr: true,
		},
		{
			name:      "default genesis returns no error",
			genesis:   types.DefaultGenesisState(),
			shouldErr: false,
		},
		{
			name: "valid genesis state returns no error",
			genesis: types.NewGenesisState(
				3,
				[]types.SubspaceData{
					types.NewSubspaceData(1, 1, 3),
					types.NewSubspaceData(2, 1, 1),
				},
				[]types.Subspace{
					types.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
					types.NewSubspace(
						2,
						"Another test subspace",
						"This is another test subspace",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
						time.Date(2020, 1, 2, 12, 00, 00, 000, time.UTC),
					),
				},
				[]types.UserPermission{
					types.NewUserPermission(1, 0, "cosmos19gz9jn5pl6ke6qg5s4gt9ga9my7w8a0x3ar0qy", types.PermissionWrite),
					types.NewUserPermission(2, 0, "cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e", types.PermissionManageGroups),
				},
				[]types.UserGroup{
					types.NewUserGroup(
						1,
						1,
						1,
						"Test group",
						"This is a test group",
						types.PermissionWrite,
					),
					types.NewUserGroup(
						2,
						1,
						1,
						"Another test group",
						"This is another test group",
						types.PermissionWrite,
					),
				},
				[]types.UserGroupMembersEntry{
					types.NewUserGroupMembersEntry(1, 1, []string{
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					}),
					types.NewUserGroupMembersEntry(2, 1, []string{
						"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					}),
				},
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateGenesis(tc.genesis)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// -------------------------------------------------------------------------------------------------------------------

func TestSubspaceData_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		data      types.SubspaceData
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			data:      types.NewSubspaceData(0, 1, 1),
			shouldErr: true,
		},
		{
			name:      "invalid next section id returns error",
			data:      types.NewSubspaceData(1, 0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid next group id returns error",
			data:      types.NewSubspaceData(1, 1, 0),
			shouldErr: true,
		},
		{
			name:      "valid data returns no error",
			data:      types.NewSubspaceData(1, 1, 1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.data.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// -------------------------------------------------------------------------------------------------------------------

func TestUserPermission_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		entry     types.UserPermission
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			entry: types.NewUserPermission(
				0,
				1,
				"cosmos19gz9jn5pl6ke6qg5s4gt9ga9my7w8a0x3ar0qy",
				types.PermissionWrite,
			),
			shouldErr: true,
		},
		{
			name: "invalid user returns no error",
			entry: types.NewUserPermission(
				1,
				0,
				"cosmos19gz9jn5pl6ke6",
				types.PermissionWrite,
			),
			shouldErr: true,
		},
		{
			name: "valid user entry returns no error",
			entry: types.NewUserPermission(
				1,
				0,
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				types.PermissionEverything,
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entry.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// -------------------------------------------------------------------------------------------------------------------

func TestUserGroupMembersEntry_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		entry     types.UserGroupMembersEntry
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			entry:     types.NewUserGroupMembersEntry(0, 1, nil),
			shouldErr: true,
		},
		{
			name:      "invalid group id returns error",
			entry:     types.NewUserGroupMembersEntry(1, 0, nil),
			shouldErr: true,
		},
		{
			name: "invalid member returns error",
			entry: types.NewUserGroupMembersEntry(1, 1, []string{
				"invalid-user",
			}),
			shouldErr: true,
		},
		{
			name: "valid entry returns no error",
			entry: types.NewUserGroupMembersEntry(1, 1, []string{
				"cosmos1nv9kkuads7f627q2zf4k9kwdudx709rjck3s7e",
				"cosmos19gz9jn5pl6ke6qg5s4gt9ga9my7w8a0x3ar0qy",
			}),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entry.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
