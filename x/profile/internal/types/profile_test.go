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

func TestProfile_WithMoniker(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	profile := types.NewProfile(owner)

	profileWithMoniker := profile.WithMoniker("monik")

	require.Equal(t, types.NewProfile(owner).WithMoniker("monik"), profileWithMoniker)
}

func TestProfile_WithName(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	moniker := "monik"
	profile := types.NewProfile(owner).WithMoniker(moniker)
	name := "name"

	tests := []struct {
		name       string
		profile    types.Profile
		profName   string
		expProfile types.Profile
	}{
		{
			name:       "not nil name value",
			profile:    profile,
			profName:   name,
			expProfile: types.Profile{Moniker: moniker, Creator: owner, Name: &name},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actProf := test.profile.WithName(&test.profName)
			require.Equal(t, test.expProfile, actProf)
		})
	}
}

func TestProfile_WithSurname(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	moniker := "monker"
	profile := types.NewProfile(owner).WithMoniker(moniker)
	surname := "surname"

	tests := []struct {
		name        string
		profile     types.Profile
		profSurname string
		expProfile  types.Profile
	}{
		{
			name:        "not nil surname",
			profile:     profile,
			profSurname: surname,
			expProfile:  types.Profile{Moniker: moniker, Creator: owner, Surname: &surname},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actProf := test.profile.WithSurname(&test.profSurname)
			require.Equal(t, test.expProfile, actProf)
		})
	}
}

func TestProfile_WithBio(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	moniker := "moniker"
	profile := types.NewProfile(owner).WithMoniker(moniker)
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
			expProfile: types.Profile{Moniker: moniker, Creator: owner, Bio: &bio},
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
	moniker := "moniker"
	profile := types.NewProfile(owner).WithMoniker(moniker)
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
			expProfile: types.Profile{Moniker: moniker, Creator: owner, Pictures: types.NewPictures(&pic, &cov)},
		},
		{
			name:       "nil pics",
			profile:    profile,
			pic:        nil,
			cov:        nil,
			expProfile: types.Profile{Moniker: moniker, Creator: owner},
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

	var name = "name"
	var surname = "surname"
	var bio = "biography"

	var testAccount = types.Profile{
		Name:     &name,
		Surname:  &surname,
		Moniker:  "moniker",
		Bio:      &bio,
		Pictures: testPictures,
		Creator:  owner,
	}

	require.Equal(t,
		`{"moniker":"moniker","name":"name","surname":"surname","bio":"biography","pictures":{"profile":"https://shorturl.at/adnX3","cover":"https://shorturl.at/cgpyF"},"creator":"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"}`,
		testAccount.String(),
	)
}

func TestProfile_Equals(t *testing.T) {
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	var pic = "profile"
	var cov = "cover"
	var testPictures = types.NewPictures(&pic, &cov)

	var testAccount = types.Profile{
		Name:     &name,
		Surname:  &surname,
		Moniker:  "moniker",
		Bio:      &bio,
		Pictures: testPictures,
		Creator:  testPostOwner,
	}

	var testAccount2 = types.Profile{
		Name:     &name,
		Surname:  &surname,
		Moniker:  "oniker",
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

//TODO add tests for chainLink and verifiedServices when implemented
func TestProfile_Validate(t *testing.T) {
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	var name = "name"
	var surname = "surname"
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
				Name:     &name,
				Surname:  &surname,
				Moniker:  "moniker",
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  nil,
			},
			expErr: fmt.Errorf("profile creator cannot be empty or blank"),
		},
		{
			name: "Empty profile moniker returns error",
			account: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  "",
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  testPostOwner,
			},
			expErr: fmt.Errorf("profile moniker cannot be empty or blank"),
		},
		{
			name: "Valid profile returns no error",
			account: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  "moniker",
				Bio:      &bio,
				Pictures: testPictures,
				Creator:  testPostOwner,
			},
			expErr: nil,
		},
		{
			name: "Invalid profile pictures returns error",
			account: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  "moniker",
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
