package types_test

import (
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func assertNoProfileError(profile *types.Profile, err error) *types.Profile {
	if err != nil {
		panic(err)
	}
	return profile
}

func TestProfile_Update(t *testing.T) {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	tests := []struct {
		name       string
		original   *types.Profile
		update     *types.ProfileUpdate
		expError   bool
		expProfile *types.Profile
	}{
		{
			name: "DoNotModify do not update original values",
			original: assertNoProfileError(types.NewProfile(
				"dtag",
				"nickname",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			update: types.NewProfileUpdate(
				types.DoNotModify,
				"",
				types.DoNotModify,
				types.NewPictures(
					types.DoNotModify,
					"",
				),
			),
			expError: false,
			expProfile: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures(
					"https://example.com",
					"",
				),
				time.Unix(100, 0),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
		},
		{
			name: "Update works properly with all fields",
			original: assertNoProfileError(types.NewProfile(
				"dtag",
				"nickname",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			update: types.NewProfileUpdate(
				"dtag-2",
				"nickname-2",
				"bio-2",
				types.NewPictures(
					"https://example.com/2",
					"https://example.com/2",
				),
			),
			expError: false,
			expProfile: assertNoProfileError(types.NewProfile(
				"dtag-2",
				"nickname-2",
				"bio-2",
				types.NewPictures(
					"https://example.com/2",
					"https://example.com/2",
				),
				time.Unix(100, 0),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
		},
		{
			name: "Update does not allow setting invalid fields",
			original: assertNoProfileError(types.NewProfile(
				"dtag",
				"nickname",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			update: types.NewProfileUpdate(
				"",
				"",
				"",
				types.NewPictures("", ""),
			),
			expError: true,
		},
		{
			name: "Update allows to set empty fields",
			original: assertNoProfileError(types.NewProfile(
				"dtag",
				"nickname",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			update: types.NewProfileUpdate(
				types.DoNotModify,
				"",
				"",
				types.NewPictures("", ""),
			),
			expError: false,
			expProfile: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures("", ""),
				time.Unix(100, 0),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
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
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	tests := []struct {
		name    string
		account *types.Profile
		expErr  error
	}{
		{
			name: "Empty profile creator returns error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Now(),
				authtypes.NewBaseAccountWithAddress(nil),
			)),
			expErr: fmt.Errorf("invalid address: "),
		},
		{
			name: "Empty profile DTag returns error",
			account: assertNoProfileError(types.NewProfile(
				"",
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			expErr: fmt.Errorf("invalid profile DTag: "),
		},
		{
			name: "Invalid profile picture returns error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("pic", "https://example.com"),
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Invalid cover picture returns error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("https://example.com", "cov"),
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			expErr: fmt.Errorf("invalid profile cover uri provided"),
		},
		{
			name: "Do not modify nickname returns error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				types.DoNotModify,
				"",
				types.Pictures{},
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			expErr: fmt.Errorf("invalid profile nickname: %s", types.DoNotModify),
		},
		{
			name: "Do not modify bio returns error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				types.DoNotModify,
				types.Pictures{},
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			expErr: fmt.Errorf("invalid profile bio: %s", types.DoNotModify),
		},
		{
			name: "Do not modify profile picture returns error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures(types.DoNotModify, ""),
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			expErr: fmt.Errorf("invalid profile picture: %s", types.DoNotModify),
		},
		{
			name: "Do not modify profile cover returns error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures("", types.DoNotModify),
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			expErr: fmt.Errorf("invalid profile cover: %s", types.DoNotModify),
		},
		{
			name: "Profile with only DTag does not error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"",
				types.Pictures{},
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
			expErr: nil,
		},
		{
			name: "Valid profile returns no error",
			account: assertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
				time.Now(),
				authtypes.NewBaseAccountWithAddress(addr1),
			)),
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
