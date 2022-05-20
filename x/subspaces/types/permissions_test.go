package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func TestRegisterPermission(t *testing.T) {
	testCases := []struct {
		name       string
		permission types.Permission
		shouldErr  bool
	}{
		{
			name:       "already registered permissions returns error",
			permission: types.PermissionEverything,
			shouldErr:  true,
		},
		{
			name:       "new permissions does not return error",
			permission: "my custom permissions",
			shouldErr:  false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldErr {
				require.Panics(t, func() { types.RegisterPermission(tc.permission) })
			} else {
				require.NotPanics(t, func() { types.RegisterPermission(tc.permission) })
			}
		})
	}

}

func TestCheckPermission(t *testing.T) {
	testCases := []struct {
		name        string
		permissions types.Permissions
		permission  types.Permission
		expResult   bool
	}{
		{
			name:        "same permissions returns true",
			permissions: types.CombinePermissions(types.PermissionEditSubspace),
			permission:  types.PermissionEditSubspace,
			expResult:   true,
		},
		{
			name:        "different permissions returns false",
			permissions: types.CombinePermissions(types.PermissionEditSubspace),
			permission:  types.PermissionSetPermissions,
			expResult:   false,
		},
		{
			name:        "combined permissions returns true when contains",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionDeleteSubspace, types.PermissionManageGroups),
			permission:  types.PermissionEditSubspace,
			expResult:   true,
		},
		{
			name:        "combined permissions returns false when does not contain",
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
		permissions types.Permissions
		expResult   types.Permissions
	}{
		{
			name:        "combining the same permissions returns the permissions itself",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionEditSubspace),
			expResult:   types.NewPermissions(types.PermissionEditSubspace),
		},
		{
			name:        "combining anything with PermissionEverything returns PermissionEverything",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionEverything),
			expResult:   types.NewPermissions(types.PermissionEverything),
		},
		{
			name:        "combining different permissions returns the correct result",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionManageGroups, types.PermissionSetPermissions),
			expResult:   types.NewPermissions(types.PermissionEditSubspace, types.PermissionManageGroups, types.PermissionSetPermissions),
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
			name:        "valid permissions returns the same value",
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			expResult:   types.NewPermissions(types.PermissionEditSubspace),
		},
		{
			name:        "multiple permission returns the same value",
			permissions: types.NewPermissions(types.PermissionEditSubspace, types.PermissionEditSubspace),
			expResult:   types.NewPermissions(types.PermissionEditSubspace),
		},
		{
			name:        "combined permissions returns the same value",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionManageGroups),
			expResult:   types.NewPermissions(types.PermissionEditSubspace, types.PermissionManageGroups),
		},
		{
			name:        "invalid permissions returns permissions nothing",
			permissions: types.NewPermissions("invalid permissions"),
			expResult:   nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			result := types.SanitizePermission(tc.permissions)
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
			name:        "valid permissions returns true",
			permissions: types.NewPermissions(types.PermissionEditSubspace),
			expValid:    true,
		},
		{
			name:        "valid combined permissions returns true",
			permissions: types.CombinePermissions(types.PermissionEditSubspace, types.PermissionEditSubspace),
			expValid:    true,
		},
		{
			name:        "invalid permissions returns false",
			permissions: types.NewPermissions("invalid permission"),
			expValid:    false,
		},
		{
			name:        "invalid combined permissions returns false",
			permissions: types.NewPermissions(types.PermissionEditSubspace, "invalid permission"),
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
