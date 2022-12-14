package types_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
	"github.com/stretchr/testify/require"
)

func TestUserGrant_Validate(t *testing.T) {
	validGrant, err := types.NewUserGrant(1, "granter", "grantee", &feegrant.BasicAllowance{})
	require.NoError(t, err)
	testCases := []struct {
		name      string
		grant     types.UserGrant
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			grant: types.UserGrant{
				SubspaceID: 0,
				Granter:    "granter",
				Grantee:    "grantee",
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid granter returns error - empty",
			grant: types.UserGrant{
				SubspaceID: 1,
				Granter:    "",
				Grantee:    "grantee",
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error - empty",
			grant: types.UserGrant{
				SubspaceID: 1,
				Granter:    "granter",
				Grantee:    "",
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error - blank",
			grant: types.UserGrant{
				SubspaceID: 1,
				Granter:    "granter",
				Grantee:    "   ",
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "granter self-grant returns error",
			grant: types.UserGrant{
				SubspaceID: 1,
				Granter:    "granter",
				Grantee:    "granter",
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid allowance returns error",
			grant: types.UserGrant{
				SubspaceID: 1,
				Granter:    "granter",
				Grantee:    "grantee",
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name:      "valid grant returns no error",
			grant:     validGrant,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.grant.ValidateBasic()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// --------------------------------------------------------------------------------------------------------------------

func TestGroupGrant_Validate(t *testing.T) {
	validGrant, err := types.NewGroupGrant(1, "granter", 1, &feegrant.BasicAllowance{})
	require.NoError(t, err)
	testCases := []struct {
		name      string
		grant     types.GroupGrant
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			grant: types.GroupGrant{
				SubspaceID: 0,
				Granter:    "granter",
				GroupID:    1,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid granter returns error - empty",
			grant: types.GroupGrant{
				SubspaceID: 1,
				Granter:    "",
				GroupID:    1,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid grantee returns error - empty",
			grant: types.GroupGrant{
				SubspaceID: 1,
				Granter:    "  ",
				GroupID:    1,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid group id returns error - blank",
			grant: types.GroupGrant{
				SubspaceID: 1,
				Granter:    "granter",
				GroupID:    0,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid allowance returns error",
			grant: types.GroupGrant{
				SubspaceID: 1,
				Granter:    "granter",
				GroupID:    1,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name:      "valid grant returns no error",
			grant:     validGrant,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.grant.ValidateBasic()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
