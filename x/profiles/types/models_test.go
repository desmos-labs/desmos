package types_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types/common"
	"github.com/stretchr/testify/require"
)

func TestPictures_Equals(t *testing.T) {
	tests := []struct {
		name      string
		pictures  *types.Pictures
		otherPics *types.Pictures
		expBool   bool
	}{
		{
			name:      "Different pictures returns false",
			pictures:  types.NewPictures(common.NewStrPtr("cover"), common.NewStrPtr("profile")),
			otherPics: types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "First picture with nil value returns false (profile)",
			pictures:  types.NewPictures(nil, common.NewStrPtr("cover")),
			otherPics: types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "First picture with nil value returns false (cover)",
			pictures:  types.NewPictures(common.NewStrPtr("profile"), nil),
			otherPics: types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "Second picture with nil value returns false (profile)",
			pictures:  types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			otherPics: types.NewPictures(nil, common.NewStrPtr("cover")),
			expBool:   false,
		},
		{
			name:      "Second picture with nil value returns false (cover)",
			pictures:  types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			otherPics: types.NewPictures(common.NewStrPtr("profile"), nil),
			expBool:   false,
		},
		{
			name:      "Equals pictures returns true",
			pictures:  types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			otherPics: types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   true,
		},
		{
			name:      "Same values but different pointers return true",
			pictures:  types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			otherPics: types.NewPictures(common.NewStrPtr("profile"), common.NewStrPtr("cover")),
			expBool:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.pictures.Equals(test.otherPics)
			require.Equal(t, test.expBool, actual)
		})
	}
}

func TestPictures_Validate(t *testing.T) {
	profilePic := "https://shorturl.at/adnX3"
	profileCov := "https://shorturl.at/cgpyF"
	invalidURI := "invalid"
	tests := []struct {
		name     string
		pictures *types.Pictures
		expErr   error
	}{
		{
			name:     "Valid Pictures",
			pictures: types.NewPictures(&profilePic, &profileCov),
			expErr:   nil,
		},
		{
			name:     "Invalid Pictures profile uri",
			pictures: types.NewPictures(&invalidURI, &profileCov),
			expErr:   fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name:     "Invalid Pictures cover uri",
			pictures: types.NewPictures(&profilePic, &invalidURI),
			expErr:   fmt.Errorf("invalid profile cover uri provided"),
		},
	}

	for _, test := range tests {
		actErr := test.pictures.Validate()
		require.Equal(t, test.expErr, actErr)
	}
}

// ___________________________________________________________________________________________________________________

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

func TestProfile_WithDTag(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := types.NewProfile("dtag", owner, date)

	profileWithDtag := profile.WithDTag("test-dtag")
	require.Equal(t, "test-dtag", profileWithDtag.DTag)
}

func TestProfile_WithMoniker(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := types.NewProfile("monik", owner, date)

	profileWithMoniker := profile.WithMoniker(common.NewStrPtr("test-moniker"))
	require.Equal(t, "test-moniker", *profileWithMoniker.Moniker)
}

func TestProfile_WithBio(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	profile := types.NewProfile("dtag", owner, date)

	profileWithBio := profile.WithBio(common.NewStrPtr("new-biography"))
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
			pic:     common.NewStrPtr("pic"),
			cov:     common.NewStrPtr("cov"),
			expProfile: types.NewProfile("dtag", owner, date).
				WithPictures(common.NewStrPtr("pic"), common.NewStrPtr("cov")),
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

	tests := []struct {
		name      string
		profile   types.Profile
		expString string
	}{
		{
			name:      "profile without moniker, bio and pictures",
			profile:   types.NewProfile("my_Tag", owner, date),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC",
		},
		{
			name:      "profile with moniker",
			profile:   types.NewProfile("my_Tag", owner, date).WithMoniker(common.NewStrPtr("moniker")),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC [Moniker] moniker",
		},
		{
			name:      "profile with bio",
			profile:   types.NewProfile("my_Tag", owner, date).WithBio(common.NewStrPtr("bio")),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC [Biography] bio",
		},
		{
			name:      "profile with profile pic",
			profile:   types.NewProfile("my_Tag", owner, date).WithPictures(common.NewStrPtr("pic"), nil),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC Pictures:\n[Profile] pic ",
		},
		{
			name:      "profile with profile cov",
			profile:   types.NewProfile("my_Tag", owner, date).WithPictures(nil, common.NewStrPtr("cov")),
			expString: "Profile:\n[Dtag] my_Tag [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Creation Time] 2010-10-02 12:10:00 +0000 UTC Pictures:\n[Cover] cov",
		},
		{
			name:      "profile with moniker, bio, pictures",
			profile:   types.NewProfile("my_Tag", owner, date).WithMoniker(common.NewStrPtr("moniker")).WithBio(common.NewStrPtr("bio")).WithPictures(common.NewStrPtr("pic"), common.NewStrPtr("cov")),
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
		first   types.Profile
		second  types.Profile
		expBool bool
	}{
		{
			name:    "Different DTag returns false",
			first:   types.NewProfile("dtag-1", user1, time1),
			second:  types.NewProfile("dtag-2", user1, time1),
			expBool: false,
		},
		{
			name: "Different moniker returns false",
			first: types.NewProfile("dtag", user1, time1).
				WithMoniker(common.NewStrPtr("moniker-1")),
			second: types.NewProfile("dtag", user1, time1).
				WithMoniker(common.NewStrPtr("moniker-2")),
			expBool: false,
		},
		{
			name: "Different bio returns false",
			first: types.NewProfile("dtag", user1, time1).
				WithBio(common.NewStrPtr("bio-1")),
			second: types.NewProfile("dtag", user1, time1).
				WithBio(common.NewStrPtr("bio-2")),
			expBool: false,
		},
		{
			name: "Different pictures returns false",
			first: types.NewProfile("dtag", user1, time1).
				WithPictures(common.NewStrPtr("profile-1"), common.NewStrPtr("cover-1")),
			second: types.NewProfile("dtag", user1, time1).
				WithPictures(common.NewStrPtr("profile-2"), common.NewStrPtr("cover-2")),
			expBool: false,
		},
		{
			name:    "Different creation dates returns false",
			first:   types.NewProfile("dtag", user1, time1),
			second:  types.NewProfile("dtag", user1, time2),
			expBool: false,
		},
		{
			name:    "Different creators returns false",
			first:   types.NewProfile("dtag", user1, time1),
			second:  types.NewProfile("dtag", user2, time1),
			expBool: false,
		},
		{
			name: "Same profiles return true",
			first: types.NewProfile("dtag-1", user1, time1).
				WithMoniker(common.NewStrPtr("moniker-1")).
				WithBio(common.NewStrPtr("bio-1")).
				WithPictures(common.NewStrPtr("profile-1"), common.NewStrPtr("cover-1")),
			second: types.NewProfile("dtag-1", user1, time1).
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
		account types.Profile
		expErr  error
	}{
		{
			name: "Empty profile creator returns error",
			account: types.Profile{
				DTag: "dtag",
				Bio:  common.NewStrPtr("bio"),
				Pictures: types.NewPictures(
					common.NewStrPtr("https://shorturl.at/adnX3"),
					common.NewStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: nil,
			},
			expErr: fmt.Errorf("profile creator cannot be empty or blank"),
		},
		{
			name: "Empty profile dtag returns error",
			account: types.Profile{
				DTag: "",
				Bio:  common.NewStrPtr("bio"),
				Pictures: types.NewPictures(
					common.NewStrPtr("https://shorturl.at/adnX3"),
					common.NewStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: user,
			},
			expErr: fmt.Errorf("profile dtag cannot be empty or blank"),
		},
		{
			name: "Valid profile returns no error",
			account: types.Profile{
				DTag: "dtag",
				Bio:  common.NewStrPtr("bio"),
				Pictures: types.NewPictures(
					common.NewStrPtr("https://shorturl.at/adnX3"),
					common.NewStrPtr("https://shorturl.at/cgpyF"),
				),
				Creator: user,
			},
			expErr: nil,
		},
		{
			name: "Invalid profile pictures returns error",
			account: types.Profile{
				DTag:     "dtag",
				Bio:      common.NewStrPtr("bio"),
				Pictures: types.NewPictures(common.NewStrPtr("pic"), common.NewStrPtr("cov")),
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
	profile := types.Profile{
		DTag: "dtag",
		Bio:  common.NewStrPtr("bio"),
		Pictures: types.NewPictures(
			common.NewStrPtr("https://shorturl.at/adnX3"),
			common.NewStrPtr("https://shorturl.at/cgpyF"),
		),
		Creator: nil,
	}

	profiles := NewProfiles(profile)

	require.Equal(t, Profiles{profile}, profiles)
}

// ___________________________________________________________________________________________________________________

func TestNewDTagTransferRequest(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	dTag := "dtag"

	require.Equal(t, DTagTransferRequest{DTagToTrade: dTag, Receiver: user, Sender: otherUser},
		NewDTagTransferRequest(dTag, user, otherUser))
}

func TestDTagTransferRequest_Equals(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name     string
		request  DTagTransferRequest
		otherReq DTagTransferRequest
		expBool  bool
	}{
		{
			name:     "Equals requests return true",
			request:  NewDTagTransferRequest("dtag", user, otherUser),
			otherReq: NewDTagTransferRequest("dtag", user, otherUser),
			expBool:  true,
		},
		{
			name:     "Non equals requests return false",
			request:  NewDTagTransferRequest("dtag", user, otherUser),
			otherReq: NewDTagTransferRequest("dtag", user, user),
			expBool:  false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.request.Equals(test.otherReq))
		})
	}
}

func TestDTagTransferRequest_String(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	require.Equal(t,
		"DTag transfer request:\n[DTag to trade] dtag [Request Receiver] cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 [Request Sender] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		NewDTagTransferRequest("dtag", user, otherUser).String(),
	)
}

func TestDTagTransferRequest_Validate(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name    string
		request DTagTransferRequest
		expErr  error
	}{
		{
			name:    "Empty DTag to trade returns error",
			request: NewDTagTransferRequest("", user, otherUser),
			expErr:  fmt.Errorf("invalid DTag to trade "),
		},
		{
			name:    "Empty request receiver returns error",
			request: NewDTagTransferRequest("dtag", nil, otherUser),
			expErr:  fmt.Errorf("receiver address cannot be empty"),
		},
		{
			name:    "Empty request sender returns error",
			request: NewDTagTransferRequest("dtag", user, nil),
			expErr:  fmt.Errorf("sender address cannot be empty"),
		},
		{
			name:    "Equals request receiver and request sender addresses return error",
			request: NewDTagTransferRequest("dtag", user, user),
			expErr:  fmt.Errorf("the sender and receiver must be different"),
		},
		{
			name:    "Valid request returns no error",
			request: NewDTagTransferRequest("dtag", user, otherUser),
			expErr:  nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.request.Validate())
		})
	}
}
