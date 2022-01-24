package types_test

import (
	"testing"
	"time"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"

	"github.com/stretchr/testify/require"
)

func TestACLEntry_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		entry     types.ACLEntry
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			entry: types.NewACLEntry(
				0,
				"group",
				types.PermissionWrite,
			),
			shouldErr: true,
		},
		{
			name: "valid group entry returns no error",
			entry: types.NewACLEntry(
				1,
				"group",
				types.PermissionWrite,
			),
			shouldErr: false,
		},
		{
			name: "valid user entry returns no error",
			entry: types.NewACLEntry(
				1,
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

func TestUserGroup_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		group     types.UserGroup
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			group:     types.NewUserGroup(0, "group", nil),
			shouldErr: true,
		},
		{
			name:      "invalid group name returns error",
			group:     types.NewUserGroup(1, "", nil),
			shouldErr: true,
		},
		{
			name: "invalid member returns error",
			group: types.NewUserGroup(1, "group", []string{
				"another-group",
			}),
			shouldErr: true,
		},
		{
			name: "valid group returns no error",
			group: types.NewUserGroup(1, "group", []string{
				"cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
			}),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.group.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name      string
		genesis   *types.GenesisState
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error (zero)",
			genesis:   types.NewGenesisState(0, nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "invalid subspace id returns error (too low)",
			genesis: types.NewGenesisState(
				1,
				[]types.Subspace{
					types.NewSubspace(
						1,
						"",
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
				},
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			genesis: types.NewGenesisState(
				1,
				[]types.Subspace{
					types.NewSubspace(
						1,
						"",
						"This is a test subspace",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				},
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "duplicated subspace returns error",
			genesis: types.NewGenesisState(
				3,
				[]types.Subspace{
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
				},
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid ACL entry returns error",
			genesis: types.NewGenesisState(
				1,
				nil,
				nil,
				[]types.ACLEntry{
					types.NewACLEntry(0, "group", types.PermissionWrite),
				},
			),
			shouldErr: true,
		},
		{
			name: "duplicated ACL entry returns error",
			genesis: types.NewGenesisState(
				1,
				nil,
				nil,
				[]types.ACLEntry{
					types.NewACLEntry(1, "group", types.PermissionWrite),
					types.NewACLEntry(1, "group", types.PermissionSetPermissions),
				},
			),
			shouldErr: true,
		},
		{
			name: "invalid group returns error",
			genesis: types.NewGenesisState(
				1,
				nil,
				[]types.UserGroup{
					types.NewUserGroup(0, "group", nil),
				},
				nil,
			),
			shouldErr: true,
		},
		{
			name: "duplicated group returns error",
			genesis: types.NewGenesisState(
				1,
				nil,
				[]types.UserGroup{
					types.NewUserGroup(1, "group", nil),
					types.NewUserGroup(1, "group", nil),
				},
				nil,
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
				[]types.UserGroup{
					types.NewUserGroup(1, "group", []string{
						"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
					}),
					types.NewUserGroup(2, "another-group", []string{
						"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					}),
				},
				[]types.ACLEntry{
					types.NewACLEntry(1, "group", types.PermissionWrite),
					types.NewACLEntry(2, "another-group", types.PermissionManageGroups),
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
