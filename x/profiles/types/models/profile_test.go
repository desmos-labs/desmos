package models_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/desmos-labs/desmos/x/profiles/types/common"
	"github.com/desmos-labs/desmos/x/profiles/types/models"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewProfile(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	expProfile := models.NewProfile("test", owner, date)
	actProfile := models.NewProfile("test", owner, date)

	require.True(t, expProfile.Equals(actProfile))
}

func TestProfile_WithMoniker(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := models.NewProfile("monik", owner, date)

	profileWithMoniker := profile.WithMoniker(common.NewStrPtr("test-moniker"))
	require.Equal(t, "test-moniker", *profileWithMoniker.Moniker)
}

func TestProfile_WithBio(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := models.NewProfile("dtag", owner, date)

	profileWithBio := profile.WithBio(common.NewStrPtr("new-biography"))
	require.Equal(t, "new-biography", *profileWithBio.Bio)
}

func TestProfile_WithPics(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := models.NewProfile("dtag", owner, date)

	tests := []struct {
		name       string
		profile    models.Profile
		pic        *string
		cov        *string
		expProfile models.Profile
	}{
		{
			name:    "not nil pics",
			profile: profile,
			pic:     common.NewStrPtr("pic"),
			cov:     common.NewStrPtr("cov"),
			expProfile: models.NewProfile("dtag", owner, date).
				WithPictures(common.NewStrPtr("pic"), common.NewStrPtr("cov")),
		},
		{
			name:       "nil pics",
			profile:    profile,
			pic:        nil,
			cov:        nil,
			expProfile: models.NewProfile("dtag", owner, date),
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

	tests := []struct {
		name      string
		profile   models.Profile
		expString string
	}{
		{
			name:      "profile without moniker, bio and pictures",
			profile:   models.NewProfile("my_Tag", owner, date),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC",
		},
		{
			name:      "profile with moniker",
			profile:   models.NewProfile("my_Tag", owner, date).WithMoniker(common.NewStrPtr("moniker")),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC [Moniker] moniker",
		},
		{
			name:      "profile with bio",
			profile:   models.NewProfile("my_Tag", owner, date).WithBio(common.NewStrPtr("bio")),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC [Biography] bio",
		},
		{
			name:      "profile with profile pic",
			profile:   models.NewProfile("my_Tag", owner, date).WithPictures(common.NewStrPtr("pic"), nil),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC Pictures:\n[Profile] pic ",
		},
		{
			name:      "profile with profile cov",
			profile:   models.NewProfile("my_Tag", owner, date).WithPictures(nil, common.NewStrPtr("cov")),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC Pictures:\n[Cover] cov",
		},
		{
			name:      "profile with moniker, bio, pictures",
			profile:   models.NewProfile("my_Tag", owner, date).WithMoniker(common.NewStrPtr("moniker")).WithBio(common.NewStrPtr("bio")).WithPictures(common.NewStrPtr("pic"), common.NewStrPtr("cov")),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC [Moniker] moniker [Biography] bio Pictures:\n[Profile] pic [Cover] cov",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expString, test.profile.String())
		})
	}
}

func TestProfile_Equals(t *testing.T) {
	user1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	user2, err := sdk.AccAddressFromBech32("cosmos1a8z4rzhd00eqvknv9dfga5rrh8fxwfs86kesv2")
	require.NoError(t, err)

	time1, err := time.Parse(time.RFC3339, "2020-01-01T01:01:01Z")
	require.NoError(t, err)

	time2, err := time.Parse(time.RFC3339, "2020-02-02T02:02:02Z")
	require.NoError(t, err)

	tests := []struct {
		name    string
		first   models.Profile
		second  models.Profile
		expBool bool
	}{
		{
			name:    "Different DTag returns false",
			first:   models.NewProfile("dtag-1", user1, time1),
			second:  models.NewProfile("dtag-2", user1, time1),
			expBool: false,
		},
		{
			name: "Different moniker returns false",
			first: models.NewProfile("dtag", user1, time1).
				WithMoniker(common.NewStrPtr("moniker-1")),
			second: models.NewProfile("dtag", user1, time1).
				WithMoniker(common.NewStrPtr("moniker-2")),
			expBool: false,
		},
		{
			name: "Different bio returns false",
			first: models.NewProfile("dtag", user1, time1).
				WithBio(common.NewStrPtr("bio-1")),
			second: models.NewProfile("dtag", user1, time1).
				WithBio(common.NewStrPtr("bio-2")),
			expBool: false,
		},
		{
			name: "Different pictures returns false",
			first: models.NewProfile("dtag", user1, time1).
				WithPictures(common.NewStrPtr("profile-1"), common.NewStrPtr("cover-1")),
			second: models.NewProfile("dtag", user1, time1).
				WithPictures(common.NewStrPtr("profile-2"), common.NewStrPtr("cover-2")),
			expBool: false,
		},
		{
			name:    "Different creation dates returns false",
			first:   models.NewProfile("dtag", user1, time1),
			second:  models.NewProfile("dtag", user1, time2),
			expBool: false,
		},
		{
			name:    "Different creators returns false",
			first:   models.NewProfile("dtag", user1, time1),
			second:  models.NewProfile("dtag", user2, time1),
			expBool: false,
		},
		{
			name: "Same profiles return true",
			first: models.NewProfile("dtag-1", user1, time1).
				WithMoniker(common.NewStrPtr("moniker-1")).
				WithBio(common.NewStrPtr("bio-1")).
				WithPictures(common.NewStrPtr("profile-1"), common.NewStrPtr("cover-1")),
			second: models.NewProfile("dtag-1", user1, time1).
				WithMoniker(common.NewStrPtr("moniker-1")).
				WithBio(common.NewStrPtr("bio-1")).
				WithPictures(common.NewStrPtr("profile-1"), common.NewStrPtr("cover-1")),
			expBool: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.first.Equals(test.second))
		})
	}

}

func TestProfile_Validate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	tests := []struct {
		name    string
		account models.Profile
		expErr  error
	}{
		{
			name: "Empty profile creator returns error",
			account: models.Profile{
				DTag: "dtag",
				Bio:  common.NewStrPtr("bio"),
				Pictures: models.NewPictures(
					common.NewStrPtr("https://shorturl.at/adnX3"),
					common.NewStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: nil,
			},
			expErr: fmt.Errorf("profile creator cannot be empty or blank"),
		},
		{
			name: "Empty profile dtag returns error",
			account: models.Profile{
				DTag: "",
				Bio:  common.NewStrPtr("bio"),
				Pictures: models.NewPictures(
					common.NewStrPtr("https://shorturl.at/adnX3"),
					common.NewStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: user,
			},
			expErr: fmt.Errorf("profile dtag cannot be empty or blank"),
		},
		{
			name: "Valid profile returns no error",
			account: models.Profile{
				DTag: "dtag",
				Bio:  common.NewStrPtr("bio"),
				Pictures: models.NewPictures(
					common.NewStrPtr("https://shorturl.at/adnX3"),
					common.NewStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: user,
			},
			expErr: nil,
		},
		{
			name: "Invalid profile pictures returns error",
			account: models.Profile{
				DTag:     "dtag",
				Bio:      common.NewStrPtr("bio"),
				Pictures: models.NewPictures(common.NewStrPtr("pic"), common.NewStrPtr("cov")),
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

func TestNewProfiles(t *testing.T) {
	profile := models.Profile{
		DTag: "dtag",
		Bio:  common.NewStrPtr("bio"),
		Pictures: models.NewPictures(
			common.NewStrPtr("https://shorturl.at/adnX3"),
			common.NewStrPtr("https://shorturl.at/cgpyF"),
		),
		Creator: nil,
	}

	profiles := types.NewProfiles(profile)

	require.Equal(t, types.Profiles{profile}, profiles)
}
