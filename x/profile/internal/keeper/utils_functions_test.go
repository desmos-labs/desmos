package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_IterateProfile(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	creator2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	creator3, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	creator4, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	profiles := types.Profiles{
		types.NewProfile("first", creator),
		types.NewProfile("second", creator2),
		types.NewProfile("not", creator3),
		types.NewProfile("third", creator4),
	}

	expProfiles := types.Profiles{
		profiles[0],
		profiles[1],
		profiles[3],
	}

	ctx, k := SetupTestInput()

	for _, profile := range profiles {
		err := k.SaveProfile(ctx, profile)
		require.NoError(t, err)
	}

	var validProfiles types.Profiles
	k.IterateProfiles(ctx, func(_ int64, profile types.Profile) (stop bool) {
		if profile.DTag == "not" {
			return false
		}
		validProfiles = append(validProfiles, profile)
		return false
	})

	require.Len(t, expProfiles, len(validProfiles))

	for _, profile := range validProfiles {
		require.Contains(t, expProfiles, profile)
	}

}
