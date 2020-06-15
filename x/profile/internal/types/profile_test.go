package types_test

import (
	"fmt"
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

	moniker := "test-moniker"
	profileWithMoniker := profile.WithMoniker(&moniker)

	require.Equal(t, "test-moniker", *profileWithMoniker.Moniker)
}

func TestProfile_WithBio(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	profile := types.NewProfile("dtag", owner)

	bio := "surname"
	profileWithBio := profile.WithBio(&bio)

	require.Equal(t, "surname", *profileWithBio.Bio)
}

func TestProfile_WithPics(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	dtag := "dtag"
	profile := types.NewProfile(dtag, owner)
	var pic = "profile"
	var cov = "cover"

	tests := []struct {
		name       string
		profile    types.Profile
		pic        *string
		cov        *string
		expProfile types.Profile
	}{
		{
			name:       "not nil pics",
			profile:    profile,
			pic:        &pic,
			cov:        &cov,
			expProfile: types.Profile{DTag: dtag, Creator: owner, Pictures: types.NewPictures(&pic, &cov)},
		},
		{
			name:       "nil pics",
			profile:    profile,
			pic:        nil,
			cov:        nil,
			expProfile: types.Profile{DTag: dtag, Creator: owner},
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
		DTag:     "dtag",
		Bio:      &bio,
		Pictures: testPictures,
		Creator:  owner,
	}

	require.Equal(t,
		`{"dtag":"dtag","bio":"biography","pictures":{"profile":"https://shorturl.at/adnX3","cover":"https://shorturl.at/cgpyF"},"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`,
		testAccount.String(),
	)
}

func TestProfile_Equals(t *testing.T) {
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	var pic = "profile"
	var cov = "cover"
	var testPictures = types.NewPictures(&pic, &cov)

	var testAccount = types.Profile{
		DTag:     "dtag",
		Bio:      &bio,
		Pictures: testPictures,
		Creator:  testPostOwner,
	}

	var testAccount2 = types.Profile{
		DTag:     "oniker",
		Bio:      &bio,
		Pictures: testPictures,
		Creator:  testPostOwner,
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
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	shortMoniker := "1"
	longMoniker := "012345678901234567890123456789012346578901234657890123457890123456789"

	bio := "biography"
	pic := "pic"
	cov := "cov"
	invalidPics := types.NewPictures(&pic, &cov)

	tests := []struct {
		name    string
		account types.Profile
		expErr  error
	}{
		{
			name: "Empty profile creator returns error",
			account: types.Profile{
				DTag:     "dtag",
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  nil,
			},
			expErr: fmt.Errorf("profile creator cannot be empty or blank"),
		},
		{
			name: "Empty profile dtag returns error",
			account: types.Profile{
				DTag:     "",
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  testPostOwner,
			},
			expErr: fmt.Errorf("invalid profile dtag"),
		},
		{
			name: "Short moniker profile returns error",
			account: types.Profile{
				DTag:     "dtag",
				Moniker:  &shortMoniker,
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  testPostOwner,
			},
			expErr: fmt.Errorf("invalid profile moniker. Length should be between 2 and 50"),
		},
		{
			name: "Long moniker profile returns error",
			account: types.Profile{
				DTag:     "dtag",
				Moniker:  &longMoniker,
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  testPostOwner,
			},
			expErr: fmt.Errorf("invalid profile moniker. Length should be between 2 and 50"),
		},
		{
			name: "Valid profile returns no error",
			account: types.Profile{
				DTag:     "dtag",
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  testPostOwner,
			},
			expErr: nil,
		},
		{
			name: "Invalid profile pictures returns error",
			account: types.Profile{
				DTag:     "dtag",
				Bio:      &bio,
				Pictures: invalidPics,
				Creator:  testPostOwner,
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
