package v016

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_migrateProfile(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	addr, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	// create a new profile with nickname field
	profile, err := types.NewProfile(
		"dtag",
		"nickname",
		"bio",
		types.NewPictures(
			"https://example.com",
			"https://example.com",
		),
		time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
		authtypes.NewBaseAccountWithAddress(addr),
	)

	require.NoError(t, err)

	// create a legacy v0160 profile with the same fields
	v0160Profile := types.Profile016{
		Account:      profile.Account,
		DTag:         profile.DTag,
		Moniker:      profile.Nickname,
		Bio:          profile.Bio,
		Pictures:     profile.Pictures,
		CreationDate: profile.CreationDate,
	}

	// encode the legacy v0160 profile
	binaryProfile016 := cdc.MustMarshalBinaryBare(&v0160Profile)

	// Try unmarshaling it into actual profile
	var actualProfile types.Profile
	err = cdc.UnmarshalBinaryBare(binaryProfile016, &actualProfile)
	require.NoError(t, err)

	require.Equal(t, *profile, actualProfile)
}
