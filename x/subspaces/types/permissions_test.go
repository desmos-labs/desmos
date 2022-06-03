package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func TestCheckPermission(t *testing.T) {
	testCases := []struct {
		name        string
		permissions types.Permissions
		permission  types.Permission
		expResult   bool
	}{
		{
			name:        "same permission returns true",
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			permission:  types.PermissionEditSubspace,
			expResult:   true,
		},
		{
			name:        "different permission returns false",
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			permission:  types.PermissionSetPermissions,
			expResult:   false,
		},
		{
			name:        "combined permission returns true when contains",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace, types.PermissionManageGroups),
			permission:  types.PermissionDeleteSubspace,
			expResult:   true,
		},
		{
			name:        "combined permission returns false when does not contain",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace, types.PermissionManageGroups),
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
		expResult   types.Permissions
	}{
		{
			name:        "combining the same permission returns the permission itself",
			permissions: []types.Permission{types.PermissionEditSubspace, types.PermissionEditSubspace},
			expResult:   types.NewPermissions(types.PermissionEditSubspace),
		},
		{
			name:        "combining anything with PermissionEverything returns PermissionEverything",
			permissions: []types.Permission{types.PermissionEditSubspace, types.PermissionEverything},
			expResult:   types.NewPermissions(types.PermissionEverything),
		},
		{
			name:        "combining different permissions returns the correct result",
			permissions: []types.Permission{types.PermissionEditSubspace, types.PermissionManageGroups, types.PermissionSetPermissions},
			expResult:   types.CombinePermissions(types.PermissionEditSubspace, types.PermissionManageGroups, types.PermissionSetPermissions),
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
		name        string
		permissions types.Permissions
		expResult   types.Permissions
	}{
		{
			name:        "valid permission returns the same value",
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			expResult:   types.NewPermissions(types.PermissionEditSubspace),
		},
		{
			name:        "combined permission returns the same value",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace),
			expResult:   types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace),
		},
		{
			name:        "invalid permission returns permission nothing",
			permissions: types.NewPermissions(""),
			expResult:   nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := types.SanitizePermissions(tc.permissions)
			require.Equal(t, tc.expResult, result)
		})
	}
}

func TestArePermissionsValid(t *testing.T) {
	testCases := []struct {
		name        string
		permissions types.Permissions
		expValid    bool
	}{
		{
			name:        "valid permission returns true",
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			expValid:    true,
		},
		{
			name:        "valid combined permission returns true",
			permissions: types.NewPermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace),
			expValid:    true,
		},
		{
			name:        "invalid permission returns false",
			permissions: types.NewPermissions(""),
			expValid:    false,
		},
		{
			name:        "invalid combined permission returns false",
			permissions: types.NewPermissions(types.PermissionEditSubspace, ""),
			expValid:    false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			valid := types.ArePermissionsValid(tc.permissions)
			require.Equal(t, tc.expValid, valid)
		})
	}
}
