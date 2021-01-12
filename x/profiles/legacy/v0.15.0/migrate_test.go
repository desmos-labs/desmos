package v0150_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v0130profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
	v0150profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.15.0"
)

func TestMigrate(t *testing.T) {
	profileCreator, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	dTagReceiver, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	moniker := "moniker"
	bio := "bio"
	profile := "profile"
	cover := "cover"
	profileCreationTime := time.Now().UTC()

	pictures := v0130profiles.Pictures{
		Profile: &profile,
		Cover:   &cover,
	}

	v0130genesisState := v0130profiles.GenesisState{
		Profiles: []v0130profiles.Profile{
			{
				DTag:         "dtag",
				Moniker:      &moniker,
				Bio:          &bio,
				Pictures:     &pictures,
				Creator:      profileCreator,
				CreationDate: profileCreationTime,
			},
			{
				DTag:         "dtag",
				Moniker:      nil,
				Bio:          nil,
				Pictures:     nil,
				Creator:      profileCreator,
				CreationDate: profileCreationTime,
			},
		},
		DTagTransferRequests: []v0130profiles.DTagTransferRequest{
			{
				DTagToTrade: "dtagToTrade",
				Receiver:    dTagReceiver,
				Sender:      profileCreator,
			},
		},
		Params: v0130profiles.Params{
			MonikerParams: v0130profiles.MonikerParams{
				MinMonikerLen: sdk.NewIntFromUint64(2),
				MaxMonikerLen: sdk.NewIntFromUint64(50),
			},
			DtagParams: v0130profiles.DtagParams{
				RegEx:      "reg_ex",
				MinDtagLen: sdk.NewIntFromUint64(3),
				MaxDtagLen: sdk.NewIntFromUint64(15),
			},
			MaxBioLen: sdk.NewIntFromUint64(200),
		},
	}

	expectedGenState := v0150profiles.GenesisState{
		Profiles: []v0150profiles.Profile{
			{
				Dtag:    "dtag",
				Moniker: moniker,
				Bio:     bio,
				Pictures: v0150profiles.Pictures{
					Profile: *pictures.Profile,
					Cover:   *pictures.Cover,
				},
				Creator:      "",
				CreationDate: profileCreationTime,
			},
			{
				Dtag:    "dtag",
				Moniker: "",
				Bio:     "",
				Pictures: v0150profiles.Pictures{
					Profile: "",
					Cover:   "",
				},
				Creator:      profileCreator.String(),
				CreationDate: profileCreationTime,
			},
		},
		DtagTransferRequests: []v0150profiles.DTagTransferRequest{
			{
				DtagToTrade: "dtagToTrade",
				Receiver:    dTagReceiver.String(),
				Sender:      profileCreator.String(),
			},
		},
		Params: v0150profiles.Params{
			MonikerParams: v0150profiles.MonikerParams{
				MinMonikerLength: sdk.NewIntFromUint64(2),
				MaxMonikerLength: sdk.NewIntFromUint64(50),
			},
			DtagParams: v0150profiles.DTagParams{
				RegEx:         "reg_ex",
				MinDtagLength: sdk.NewIntFromUint64(3),
				MaxDtagLength: sdk.NewIntFromUint64(15),
			},
			MaxBioLength: sdk.NewIntFromUint64(200),
		},
	}

	migrated := v0150profiles.Migrate(v0130genesisState)

	// Check for profiles
	require.Len(t, migrated.Profiles, len(expectedGenState.Profiles))
	for index, profile := range migrated.Profiles {
		require.Equal(t, expectedGenState.Profiles[index].Dtag, profile.Dtag)
		require.Equal(t, expectedGenState.Profiles[index].Moniker, profile.Moniker)
		require.Equal(t, expectedGenState.Profiles[index].Bio, profile.Bio)
		require.Equal(t, expectedGenState.Profiles[index].Pictures.Profile, profile.Pictures.Profile)
		require.Equal(t, expectedGenState.Profiles[index].Pictures.Cover, profile.Pictures.Cover)
	}

	// Check fo
	//r dTag transfers requests
	require.Len(t, migrated.DtagTransferRequests, len(expectedGenState.DtagTransferRequests))
	for index, request := range migrated.DtagTransferRequests {
		require.Equal(t, expectedGenState.DtagTransferRequests[index], request)
	}

	// Check for params
	require.Equal(t, migrated.Params, expectedGenState.Params)
}
