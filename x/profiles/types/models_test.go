package types_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/profiles/types"

	"github.com/stretchr/testify/require"
)

func TestPictures_Validate(t *testing.T) {
	tests := []struct {
		name     string
		pictures types.Pictures
		expErr   error
	}{
		{
			name:     "Valid Pictures",
			pictures: types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			expErr:   nil,
		},
		{
			name:     "Invalid Pictures profile uri",
			pictures: types.NewPictures("invalid", "https://shorturl.at/cgpyF"),
			expErr:   fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name:     "Invalid Pictures cover uri",
			pictures: types.NewPictures("https://shorturl.at/adnX3", "invalid"),
			expErr:   fmt.Errorf("invalid profile cover uri provided"),
		},
	}

	for _, test := range tests {
		actErr := test.pictures.Validate()
		require.Equal(t, test.expErr, actErr)
	}
}

// ___________________________________________________________________________________________________________________

func TestProfile_Update(t *testing.T) {
	tests := []struct {
		name       string
		original   types.Profile
		update     types.Profile
		expError   bool
		expProfile types.Profile
	}{
		{
			name: "DoNotModify and empty fields do not update original values",
			original: types.NewProfile(
				"dtag",
				"moniker",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			update: types.NewProfile(
				types.DoNotModify,
				types.DoNotModify,
				types.DoNotModify,
				types.NewPictures(types.DoNotModify, types.DoNotModify),
				time.Time{},
				types.DoNotModify,
			),
			expError: false,
			expProfile: types.NewProfile(
				"dtag",
				"moniker",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
		},
		{
			name: "Update works properly with all fields",
			original: types.NewProfile(
				"dtag",
				"moniker",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			update: types.NewProfile(
				"dtag-2",
				"moniker-2",
				"bio-2",
				types.NewPictures(
					"https://example.com/2",
					"https://example.com/2",
				),
				time.Unix(200, 0),
				"cosmos1pqcac4w0k8z4elysqppgce5vauzu5krew7jegg",
			),
			expError: false,
			expProfile: types.NewProfile(
				"dtag-2",
				"moniker-2",
				"bio-2",
				types.NewPictures(
					"https://example.com/2",
					"https://example.com/2",
				),
				time.Unix(200, 0),
				"cosmos1pqcac4w0k8z4elysqppgce5vauzu5krew7jegg",
			),
		},
		{
			name: "Update does not allow setting invalid fields",
			original: types.NewProfile(
				"dtag",
				"moniker",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			update: types.NewProfile(
				"dtag-2",
				"",
				"",
				types.NewPictures("", ""),
				time.Time{},
				"invalid-address",
			),
			expError: true,
		},
		{
			name: "Update allows to set empty fields",
			original: types.NewProfile(
				"dtag",
				"moniker",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			update: types.NewProfile(
				types.DoNotModify,
				"",
				"",
				types.NewPictures("", ""),
				time.Time{},
				types.DoNotModify,
			),
			expError: false,
			expProfile: types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures("", ""),
				time.Unix(100, 0),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			updated, err := test.original.Update(test.update)

			if test.expError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, test.expProfile, updated)
			}
		})
	}
}

func TestProfile_Validate(t *testing.T) {
	tests := []struct {
		name    string
		account types.Profile
		expErr  error
	}{
		{
			name: "Empty profile creator returns error",
			account: types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Now(),
				"",
			),
			expErr: fmt.Errorf("invalid creator address: "),
		},
		{
			name: "Empty profile DTag returns error",
			account: types.NewProfile(
				"",
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile DTag: "),
		},
		{
			name: "Invalid profile picture returns error",
			account: types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("pic", "https://example.com"),
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Invalid cover picture returns error",
			account: types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("https://example.com", "cov"),
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile cover uri provided"),
		},
		{
			name: "Do not modify moniker returns error",
			account: types.NewProfile(
				"dtag",
				types.DoNotModify,
				"",
				types.Pictures{},
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile moniker: %s", types.DoNotModify),
		},
		{
			name: "Do not modify bio returns error",
			account: types.NewProfile(
				"dtag",
				"",
				types.DoNotModify,
				types.Pictures{},
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile bio: %s", types.DoNotModify),
		},
		{
			name: "Do not modify profile picture returns error",
			account: types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures(types.DoNotModify, ""),
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile picture: %s", types.DoNotModify),
		},
		{
			name: "Do not modify profile cover returns error",
			account: types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures("", types.DoNotModify),
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("invalid profile cover: %s", types.DoNotModify),
		},
		{
			name: "Profile with only DTag does not error",
			account: types.NewProfile(
				"dtag",
				"",
				"",
				types.Pictures{},
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: nil,
		},
		{
			name: "Valid profile returns no error",
			account: types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
				time.Now(),
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.account.Validate())
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestDTagTransferRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		request types.DTagTransferRequest
		expErr  error
	}{
		{
			name: "Empty DTag to trade returns error",
			request: types.NewDTagTransferRequest(
				"",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("invalid DTag to trade "),
		},
		{
			name: "Empty request sender returns error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: fmt.Errorf("invalid sender address: "),
		},
		{
			name: "Empty request receiver returns error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
			),
			expErr: fmt.Errorf("invalid receiver address: "),
		},
		{
			name: "Equals request receiver and request sender addresses return error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			),
			expErr: fmt.Errorf("the sender and receiver must be different"),
		},
		{
			name: "Valid request returns no error",
			request: types.NewDTagTransferRequest(
				"dtag",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.request.Validate())
		})
	}
}
