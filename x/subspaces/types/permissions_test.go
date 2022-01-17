package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
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
			permission: types.PermissionAddLink,
			expected:   []byte{0, 0, 0, 4},
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
			bz:       []byte{0, 0, 0, 8},
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
