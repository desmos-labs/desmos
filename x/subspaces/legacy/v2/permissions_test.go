package v2_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	v2 "github.com/desmos-labs/desmos/v7/x/subspaces/legacy/v2"
)

func TestSplitPermissions(t *testing.T) {
	testCases := []struct {
		name           string
		permission     v2.Permission
		expPermissions []v2.Permission
	}{
		{
			name:           "single permission is split properly",
			permission:     v2.PermissionWrite,
			expPermissions: []v2.Permission{v2.PermissionWrite},
		},
		{
			name:           "combined permission is split properly",
			permission:     v2.PermissionWrite | v2.PermissionChangeInfo | v2.PermissionSetPermissions | v2.PermissionDeleteSubspace,
			expPermissions: []v2.Permission{v2.PermissionWrite, v2.PermissionChangeInfo, v2.PermissionSetPermissions, v2.PermissionDeleteSubspace},
		},
		{
			name:           "permission everything is returned properly",
			permission:     v2.PermissionEverything | v2.PermissionChangeInfo,
			expPermissions: []v2.Permission{v2.PermissionEverything},
		},
		{
			name:           "permission nothing returns an empty slice",
			permission:     v2.PermissionNothing,
			expPermissions: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			permissions := v2.SplitPermissions(tc.permission)
			require.Equal(t, tc.expPermissions, permissions)
		})
	}
}
