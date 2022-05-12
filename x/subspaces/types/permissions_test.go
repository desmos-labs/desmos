package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func TestMarshalPermission(t *testing.T) {
	testCases := []struct {
		name       string
		permission types.Permission
		expected   []byte
	}{
		{
			name:       "zero permission",
			permission: types.PermissionNothing,
			expected:   []byte{0, 0, 0, 0},
		},
		{
			name:       "non-zero permission",
			permission: types.PermissionManageGroups,
			expected:   []byte{0, 0, 0, 8},
		},
		{
			name:       "high permission",
			permission: types.PermissionSetPermissions,
			expected:   []byte{0, 0, 0, 16},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			bz := types.MarshalPermission(tc.permission)
			require.Equal(t, tc.expected, bz)
		})
	}
}

func TestUnmarshalPermission(t *testing.T) {
	testCases := []struct {
		name     string
		bz       []byte
		expected types.Permission
	}{
		{
			name:     "empty byte array",
			bz:       []byte{},
			expected: types.PermissionNothing,
		},
		{
			name:     "nil bytes array",
			bz:       nil,
			expected: types.PermissionNothing,
		},
		{
			name:     "zero permission",
			bz:       []byte{0, 0, 0, 0},
			expected: types.PermissionNothing,
		},
		{
			name:     "non-zero permission",
			bz:       []byte{0, 0, 0, 4},
			expected: types.PermissionChangeInfo,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			permission := types.UnmarshalPermission(tc.bz)
			require.Equal(t, tc.expected, permission)
		})
	}
}

func TestCheckPermission(t *testing.T) {
	testCases := []struct {
		name        string
		permissions types.Permission
		permission  types.Permission
		expResult   bool
	}{
		{
			name:        "same permission returns true",
			permissions: types.PermissionWrite,
			permission:  types.PermissionWrite,
			expResult:   true,
		},
		{
			name:        "different permission returns false",
			permissions: types.PermissionWrite,
			permission:  types.PermissionSetPermissions,
			expResult:   false,
		},
		{
			name:        "combined permission returns true when contains",
			permissions: types.PermissionWrite | types.PermissionModerateContent | types.PermissionManageGroups,
			permission:  types.PermissionModerateContent,
			expResult:   true,
		},
		{
			name:        "combined permission returns false when does not contain",
			permissions: types.PermissionWrite | types.PermissionModerateContent | types.PermissionManageGroups,
			permission:  types.PermissionSetPermissions,
			expResult:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := types.CheckPermission(tc.permissions, tc.permission)
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestCombinePermissions(t *testing.T) {
	testCases := []struct {
		name        string
		permissions []types.Permission
		expResult   types.Permission
	}{
		{
			name:        "combining the same permission returns the permission itself",
			permissions: []types.Permission{types.PermissionWrite, types.PermissionWrite},
			expResult:   types.PermissionWrite,
		},
		{
			name:        "combining anything with PermissionNothing returns the permission itself",
			permissions: []types.Permission{types.PermissionNothing, types.PermissionWrite},
			expResult:   types.PermissionWrite,
		},
		{
			name:        "combining anything with PermissionEverything returns PermissionEverything",
			permissions: []types.Permission{types.PermissionWrite, types.PermissionEverything},
			expResult:   types.PermissionEverything,
		},
		{
			name:        "combining different permissions returns the correct result",
			permissions: []types.Permission{types.PermissionWrite, types.PermissionManageGroups, types.PermissionSetPermissions},
			expResult:   types.PermissionWrite | types.PermissionManageGroups | types.PermissionSetPermissions,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := types.CombinePermissions(tc.permissions...)
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestSanitizePermission(t *testing.T) {
	testCases := []struct {
		name       string
		permission types.Permission
		expResult  types.Permission
	}{
		{
			name:       "valid permission returns the same value",
			permission: types.PermissionWrite,
			expResult:  types.PermissionWrite,
		},
		{
			name:       "combined permission returns the same value",
			permission: types.PermissionWrite & types.PermissionChangeInfo,
			expResult:  types.PermissionWrite & types.PermissionChangeInfo,
		},
		{
			name:       "invalid permission returns permission nothing",
			permission: 256,
			expResult:  types.PermissionNothing,
		},
		{
			name:       "extra bits are set to 0",
			permission: 0b11111111111111111111111000000001,
			expResult:  types.PermissionWrite,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := types.SanitizePermission(tc.permission)
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestIsPermissionValid(t *testing.T) {
	testCases := []struct {
		name       string
		permission types.Permission
		expValid   bool
	}{
		{
			name:       "valid permission returns true",
			permission: types.PermissionWrite,
			expValid:   true,
		},
		{
			name:       "valid combined permission returns true",
			permission: types.PermissionWrite & types.PermissionChangeInfo,
			expValid:   true,
		},
		{
			name:       "invalid permission returns false",
			permission: 256,
			expValid:   false,
		},
		{
			name:       "invalid combined permission returns false",
			permission: 0b11111111111111111111111111000001,
			expValid:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			valid := types.IsPermissionValid(tc.permission)
			require.Equal(t, tc.expValid, valid)
		})
	}
}
