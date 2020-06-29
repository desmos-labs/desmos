package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/internal/keeper"
	"github.com/desmos-labs/desmos/x/profiles/internal/types"
	"github.com/stretchr/testify/require"
)

func TestInvariants(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	tests := []struct {
		name        string
		profile     types.Profile
		expResponse string
		expBool     bool
	}{
		{
			name:        "Invariants not violated",
			profile:     types.NewProfile("dtag", owner, date),
			expResponse: "Every invariant condition is fulfilled correctly",
			expBool:     true,
		},
		{
			name:        "ValidProfile invariant violated",
			profile:     types.NewProfile("", owner, date),
			expResponse: "profiles: invalid profiles invariant\nThe following list contains invalid profiles:\n Invalid profiles:\n[DTag]: , [Creator]: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\n\n",
			expBool:     true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			err := k.SaveProfile(ctx, test.profile)
			require.NoError(t, err)

			res, stop := keeper.AllInvariants(k)(ctx)

			require.Equal(t, test.expResponse, res)
			require.Equal(t, test.expBool, stop)
		})
	}

}
