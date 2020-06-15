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
	expProfile := types.Profile{Creator: owner}
	actProfile := types.NewProfile(owner)

	require.Equal(t, expProfile, actProfile)
}

func TestProfile_WithDtag(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	profile := types.NewProfile(owner)

	profileWithDtag := profile.WithDtag("monik")

	require.Equal(t, types.NewProfile(owner).WithDtag("monik"), profileWithDtag)
}

func TestProfile_WithBio(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	dtag := "dtag"
	profile := types.NewProfile(owner).WithDtag(dtag)
	bio := "surname"

	tests := []struct {
		name       string
		profile    types.Profile
		profBio    string
		expProfile types.Profile
	}{
		{
			name:       "not nil bio",
			profile:    profile,
			profBio:    bio,
			expProfile: types.Profile{DTag: dtag, Creator: owner, Bio: &bio},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actProf := test.profile.WithBio(&test.profBio)
			require.Equal(t, test.expProfile, actProf)
		})
	}
}

func TestProfile_WithPics(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	dtag := "dtag"
	profile := types.NewProfile(owner).WithDtag(dtag)
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
			require.Equal(t, test.expProfile, actProf)
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
			actual := test.account.Equals(test.otherAcc)
			require.Equal(t, test.expBool, actual)
		})
	}

}

func TestProfile_Validate(t *testing.T) {
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	var bio = "biography"
	var pic = "pic"
	var cov = "cov"
	var invalidPics = types.NewPictures(&pic, &cov)

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
			actual := test.account.Validate()
			require.Equal(t, test.expErr, actual)
		})
	}
}
