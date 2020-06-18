package types_test

import (
	"fmt"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestNewProfile(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	expProfile := types.Profile{DTag: "test", Creator: owner}
	actProfile := types.NewProfile("test", owner)

	require.True(t, expProfile.Equals(actProfile))
}

func TestProfile_WithDTag(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	profile := types.NewProfile("monik", owner)
	profileWithDtag := profile.WithDTag("new-dtag")

	require.Equal(t, "new-dtag", profileWithDtag.DTag)
}

func TestProfile_WithMoniker(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	profile := types.NewProfile("monik", owner)

	profileWithMoniker := profile.WithMoniker(newStrPtr("test-moniker"))
	require.Equal(t, "test-moniker", *profileWithMoniker.Moniker)
}

func TestProfile_WithBio(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	profile := types.NewProfile("dtag", owner)

	profileWithBio := profile.WithBio(newStrPtr("new-biography"))
	require.Equal(t, "new-biography", *profileWithBio.Bio)
}

func TestProfile_WithPics(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	profile := types.NewProfile("dtag", owner)

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
			expProfile: types.NewProfile("dtag", owner).
				WithPictures(newStrPtr("pic"), newStrPtr("cov")),
		},
		{
			name:       "nil pics",
			profile:    profile,
			pic:        nil,
			cov:        nil,
			expProfile: types.NewProfile("dtag", owner),
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

	var bio = "biography"
	var testAccount = types.Profile{
		DTag: "dtag",
		Bio:  &bio,
		Pictures: types.NewPictures(
			newStrPtr("https://shorturl.at/adnX3"),
			newStrPtr("https://shorturl.at/cgpyF"),
		),
		Creator: owner,
	}

	require.Equal(t,
		`{"dtag":"dtag","bio":"biography","pictures":{"profile":"https://shorturl.at/adnX3","cover":"https://shorturl.at/cgpyF"},"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`,
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
			expErr: fmt.Errorf("invalid profile dtag"),
		},
		{
			name: "Short moniker profile returns error",
			account: types.Profile{
				DTag:    "dtag",
				Moniker: newStrPtr("1"),
				Bio:     newStrPtr("bio"),
				Pictures: types.NewPictures(
					newStrPtr("https://shorturl.at/adnX3"),
					newStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: user,
			},
			expErr: fmt.Errorf("invalid profile moniker. Length should be between 2 and 50"),
		},
		{
			name: "Long moniker profile returns error",
			account: types.Profile{
				DTag:    "dtag",
				Moniker: newStrPtr(strings.Repeat("1", 100)),
				Bio:     newStrPtr("bio"),
				Pictures: types.NewPictures(
					newStrPtr("https://shorturl.at/adnX3"),
					newStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: user,
			},
			expErr: fmt.Errorf("invalid profile moniker. Length should be between 2 and 50"),
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
