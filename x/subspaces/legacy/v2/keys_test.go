package v2_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v2 "github.com/desmos-labs/desmos/v6/x/subspaces/legacy/v2"
)

func TestSplitUserPermissionKey(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1ytez5aseztgvpwajnp6a9l69k8gd45mpv2j3jg")
	require.NoError(t, err)

	testCases := []struct {
		name          string
		key           []byte
		expSubspaceID uint64
		expUser       sdk.AccAddress
	}{
		{
			name:          "key is split properly",
			key:           v2.UserPermissionStoreKey(1, user),
			expSubspaceID: 1,
			expUser:       user,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			subspaceID, user := v2.SplitUserPermissionKey(tc.key)
			require.Equal(t, tc.expSubspaceID, subspaceID)
			require.True(t, tc.expUser.Equals(user))
		})
	}
}
