package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func TestSplitGroupMemberStoreKey(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
	require.NoError(t, err)

	testCases := []struct {
		name          string
		key           []byte
		shouldErr     bool
		expSubspaceID uint64
		expGroupID    uint32
		expUserAddr   sdk.AccAddress
	}{
		{
			name:      "invalid key returns error",
			key:       []byte{0x01},
			shouldErr: true,
		},
		{
			name:          "valid key is split accordingly",
			key:           types.GroupMemberStoreKey(1, 2, user),
			shouldErr:     false,
			expSubspaceID: 1,
			expGroupID:    2,
			expUserAddr:   user,
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
				require.True(t, tc.expUserAddr.Equals(user))
			}
		})
	}

}
