package types_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/internal/types"
	"github.com/stretchr/testify/require"
)

func TestNewProfile(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	expProfile := types.NewProfile("test", owner, date)
	actProfile := types.NewProfile("test", owner, date)

	require.True(t, expProfile.Equals(actProfile))
}

func TestProfile_WithMoniker(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := types.NewProfile("monik", owner, date)

	profileWithMoniker := profile.WithMoniker(newStrPtr("test-moniker"))
	require.Equal(t, "test-moniker", *profileWithMoniker.Moniker)
}

func TestProfile_WithBio(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := types.NewProfile("dtag", owner, date)

	profileWithBio := profile.WithBio(newStrPtr("new-biography"))
	require.Equal(t, "new-biography", *profileWithBio.Bio)
}

func TestProfile_WithPics(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := types.NewProfile("dtag", owner, date)

	tests := []struct {
		name       string
		profile    types.Profile
		pic        *string
		cov        *string
		expProfile types.Profile
	}{
		{
			name:    "not nil pics",
			profile: profile,
			pic:     newStrPtr("pic"),
			cov:     newStrPtr("cov"),
			expProfile: types.NewProfile("dtag", owner, date).
				WithPictures(newStrPtr("pic"), newStrPtr("cov")),
		},
		{
			name:       "nil pics",
			profile:    profile,
			pic:        nil,
			cov:        nil,
			expProfile: types.NewProfile("dtag", owner, date),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actProf := test.profile.WithPictures(test.pic, test.cov)
			require.True(t, test.expProfile.Equals(actProf))
		})
	}
}

func TestProfile_String(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	testAccount := types.NewProfile("dtag", owner, date).
		WithBio(newStrPtr("biography")).
		WithPictures(
			newStrPtr("https://shorturl.at/adnX3"),
			newStrPtr("https://shorturl.at/cgpyF"),
		)

	require.Equal(t,
		`{"dtag":"dtag","bio":"biography","pictures":{"profile":"https://shorturl.at/adnX3","cover":"https://shorturl.at/cgpyF"},"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns","creation_date":"2010-10-02T12:10:00Z"}`,
		testAccount.String(),
	)
}

func TestProfile_Equals(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	var testAccount = types.Profile{
		DTag:     "dtag",
		Bio:      newStrPtr("bio"),
		Pictures: types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
		Creator:  user,
	}

	var testAccount2 = types.Profile{
		DTag:     "oniker",
		Bio:      newStrPtr("bio"),
		Pictures: types.NewPictures(newStrPtr("profile"), newStrPtr("cover")),
		Creator:  user,
	}

	tests := []struct {
		name     string
		account  types.Profile
		otherAcc types.Profile
		expBool  bool
	}{
		{
			name:     "Equals accounts returns true",
			account:  testAccount,
			otherAcc: testAccount,
			expBool:  true,
		},
		{
			name:     "Non equals account returns false",
			account:  testAccount,
			otherAcc: testAccount2,
			expBool:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.account.Equals(test.otherAcc))
		})
	}

}

func TestProfile_Validate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	tests := []struct {
		name    string
		account types.Profile
		expErr  error
	}{
		{
			name: "Empty profile creator returns error",
			account: types.Profile{
				DTag: "dtag",
				Bio:  newStrPtr("bio"),
				Pictures: types.NewPictures(
					newStrPtr("https://shorturl.at/adnX3"),
					newStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: nil,
			},
			expErr: fmt.Errorf("profile creator cannot be empty or blank"),
		},
		{
			name: "Empty profile dtag returns error",
			account: types.Profile{
				DTag: "",
				Bio:  newStrPtr("bio"),
				Pictures: types.NewPictures(
					newStrPtr("https://shorturl.at/adnX3"),
					newStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: user,
			},
			expErr: fmt.Errorf("profile dtag cannot be empty or blank"),
		},
		{
			name: "Valid profile returns no error",
			account: types.Profile{
				DTag: "dtag",
				Bio:  newStrPtr("bio"),
				Pictures: types.NewPictures(
					newStrPtr("https://shorturl.at/adnX3"),
					newStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: user,
			},
			expErr: nil,
		},
		{
			name: "Invalid profile pictures returns error",
			account: types.Profile{
				DTag:     "dtag",
				Bio:      newStrPtr("bio"),
				Pictures: types.NewPictures(newStrPtr("pic"), newStrPtr("cov")),
				Creator:  user,
			},
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.account.Validate())
		})
	}
}
