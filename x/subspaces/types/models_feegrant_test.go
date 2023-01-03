package types_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
	"github.com/stretchr/testify/require"
)

func TestGroupTarget_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		target    *types.GroupTarget
		shouldErr bool
	}{
		{
			name:      "invalid group id returns error",
			target:    types.NewGroupTarget(0),
			shouldErr: true,
		},
		{
			name:      "valid grant returns no error",
			target:    types.NewGroupTarget(1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.target.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestUserTarget_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		target    *types.UserTarget
		shouldErr bool
	}{
		{
			name:      "invalid user address returns error",
			target:    types.NewUserTarget(""),
			shouldErr: true,
		},
		{
			name:      "valid grant returns no error",
			target:    types.NewUserTarget("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.target.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGrant_Validate(t *testing.T) {
	validUserTarget := types.NewUserTarget("cosmos1lv3e0l66rr68k5l74mnrv4j9kyny6cz27pvnez")
	validTargetAny, err := codectypes.NewAnyWithValue(types.NewUserTarget("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"))
	require.NoError(t, err)

	userGrant, err := types.NewGrant(1, "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0", validUserTarget, &feegrant.BasicAllowance{})
	require.NoError(t, err)

	invalidTargetAny, err := codectypes.NewAnyWithValue(types.NewUserTarget(""))
	require.NoError(t, err)

	invalidUserTargetAny, err := codectypes.NewAnyWithValue(types.NewUserTarget("cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0"))
	require.NoError(t, err)

	testCases := []struct {
		name      string
		grant     types.Grant
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			grant: types.Grant{
				SubspaceID: 0,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Target:     &codectypes.Any{},
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid granter returns error",
			grant: types.Grant{
				SubspaceID: 1,
				Granter:    "",
				Target:     &codectypes.Any{},
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid target returns error",
			grant: types.Grant{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Target:     invalidTargetAny,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "granter self-grant returns error",
			grant: types.Grant{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Target:     invalidUserTargetAny,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name: "invalid allowance returns error",
			grant: types.Grant{
				SubspaceID: 1,
				Granter:    "cosmos1vkuuth0rak58x36m7wuzj7ztttxh26fhqcfxm0",
				Target:     validTargetAny,
				Allowance:  &codectypes.Any{},
			},
			shouldErr: true,
		},
		{
			name:      "valid grant returns no error",
			grant:     userGrant,
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.grant.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
