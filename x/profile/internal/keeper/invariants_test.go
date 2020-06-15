package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestInvariants(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	dtag := "dtag"

	tests := []struct {
		name        string
		profile     types.Profile
		expResponse string
		expBool     bool
	}{
		{
			name:        "Invariants not violated",
			profile:     types.NewProfile(owner).WithDtag(dtag),
			expResponse: "Every invariant condition is fulfilled correctly",
			expBool:     true,
		},
		{
			name:        "ValidProfile invariant violated",
			profile:     types.NewProfile(owner).WithDtag(""),
			expResponse: "profiles: invalid profiles invariant\nThe following list contains invalid profiles:\n Invalid profiles:\n[DTag]: , [Creator]: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\n\n",
			expBool:     true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			// nolint: errcheck
			k.SaveProfile(ctx, test.profile)

			res, stop := keeper.AllInvariants(k)(ctx)

			require.Equal(t, test.expResponse, res)
			require.Equal(t, test.expBool, stop)
		})
	}

}
