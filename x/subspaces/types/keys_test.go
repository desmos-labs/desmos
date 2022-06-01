package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func TestSplitGroupMemberStoreKey(t *testing.T) {
	testCases := []struct {
		name          string
		key           []byte
		shouldErr     bool
		expSubspaceID uint64
		expGroupID    uint32
		expUserAddr   string
	}{
		{
			name:      "invalid key returns error",
			key:       []byte{0x01},
			shouldErr: true,
		},
		{
			name:          "valid key is split accordingly",
			key:           types.GroupMemberStoreKey(1, 2, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			shouldErr:     false,
			expSubspaceID: 1,
			expGroupID:    2,
			expUserAddr:   "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldErr {
				require.Panics(t, func() { types.SplitGroupMemberStoreKey(tc.key) })
			} else {
				subspaceID, groupID, user := types.SplitGroupMemberStoreKey(tc.key)
				require.Equal(t, tc.expSubspaceID, subspaceID)
				require.Equal(t, tc.expGroupID, groupID)
				require.Equal(t, tc.expUserAddr, user)
			}
		})
	}

}

func TestSplitUserAddressPermissionKey(t *testing.T) {
	testCases := []struct {
		name          string
		key           []byte
		shouldErr     bool
		expSubspaceID uint64
		expSectionID  uint32
		expUser       string
	}{
		{
			name:      "invalid key returns error",
			key:       []byte{0x01},
			shouldErr: true,
		},
		{
			name:          "valid key returns proper data",
			key:           types.UserPermissionStoreKey(1, 2, "cosmos1vlknheepy5454pw4j6x53yeg57l7ec39rf8ffp"),
			expSubspaceID: 1,
			expSectionID:  2,
			expUser:       "cosmos1vlknheepy5454pw4j6x53yeg57l7ec39rf8ffp",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldErr {
				require.Panics(t, func() { types.SplitUserAddressPermissionKey(tc.key) })
			} else {
				subspaceID, sectionID, user := types.SplitUserAddressPermissionKey(tc.key)
				require.Equal(t, tc.expSubspaceID, subspaceID)
				require.Equal(t, tc.expSectionID, sectionID)
				require.Equal(t, tc.expUser, user)
			}
		})
	}
}
