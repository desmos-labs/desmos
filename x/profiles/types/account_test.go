package types_test

import (
	"testing"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func TestProfile_TrackDelegation(t *testing.T) {
	// Create a profile with vesting account
	vestingAccount := vestingtypes.NewContinuousVestingAccount(
		&authtypes.BaseAccount{
			Address: "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k",
			PubKey: codectypes.UnsafePackAny(
				profilestesting.PubKeyFromBech32(
					"cosmospub1addwnpepqtkndttcutq2sehejxs2x3jl2uhxzuds4705u8nkgayuct0khqkzjd0vvln",
				),
			),
			AccountNumber: 0,
			Sequence:      0,
		},
		sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000))),
		time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC).Unix(),
		time.Date(2100, 9, 1, 0, 0, 0, 0, time.UTC).Unix(),
	)
	profile := profilestesting.AssertNoProfileError(
		types.NewProfile(
			"dtag",
			"nickname",
			"bio",
			types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
			vestingAccount,
		),
	)

	// Execute track delegation
	profile.TrackDelegation(time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(2000))), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))))

	// Setup expected vesting account
	expected := vestingtypes.NewContinuousVestingAccount(
		&authtypes.BaseAccount{
			Address: "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k",
			PubKey: codectypes.UnsafePackAny(
				profilestesting.PubKeyFromBech32(
					"cosmospub1addwnpepqtkndttcutq2sehejxs2x3jl2uhxzuds4705u8nkgayuct0khqkzjd0vvln",
				),
			),
			AccountNumber: 0,
			Sequence:      0,
		},
		sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000))),
		time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC).Unix(),
		time.Date(2100, 9, 1, 0, 0, 0, 0, time.UTC).Unix(),
	)
	expected.TrackDelegation(time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(2000))), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))))

	require.Equal(t, expected, profile.GetAccount())
}

func TestProfile_TrackUndelegation(t *testing.T) {
	// Create a profile with vesting account
	vestingAccount := vestingtypes.NewContinuousVestingAccount(
		&authtypes.BaseAccount{
			Address: "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k",
			PubKey: codectypes.UnsafePackAny(
				profilestesting.PubKeyFromBech32(
					"cosmospub1addwnpepqtkndttcutq2sehejxs2x3jl2uhxzuds4705u8nkgayuct0khqkzjd0vvln",
				),
			),
			AccountNumber: 0,
			Sequence:      0,
		},
		sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000))),
		time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC).Unix(),
		time.Date(2100, 9, 1, 0, 0, 0, 0, time.UTC).Unix(),
	)
	profile := profilestesting.AssertNoProfileError(
		types.NewProfile(
			"dtag",
			"nickname",
			"bio",
			types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
			vestingAccount,
		),
	)

	// Execute track delegation
	profile.TrackDelegation(time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(2000))), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(200))))
	profile.TrackUndelegation(sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))))

	// Setup expected vesting account
	expected := vestingtypes.NewContinuousVestingAccount(
		&authtypes.BaseAccount{
			Address: "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k",
			PubKey: codectypes.UnsafePackAny(
				profilestesting.PubKeyFromBech32(
					"cosmospub1addwnpepqtkndttcutq2sehejxs2x3jl2uhxzuds4705u8nkgayuct0khqkzjd0vvln",
				),
			),
			AccountNumber: 0,
			Sequence:      0,
		},
		sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1000))),
		time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC).Unix(),
		time.Date(2100, 9, 1, 0, 0, 0, 0, time.UTC).Unix(),
	)
	expected.TrackDelegation(time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(2000))), sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(200))))
	expected.TrackUndelegation(sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(100))))

	require.Equal(t, expected, profile.GetAccount())
}

func TestProfile_Update(t *testing.T) {
	testCases := []struct {
		name       string
		original   *types.Profile
		update     *types.ProfileUpdate
		shouldErr  bool
		expProfile *types.Profile
	}{
		{
			name: "DoNotModify does not update original values",
			original: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"nickname",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			update: types.NewProfileUpdate(
				types.DoNotModify,
				types.DoNotModify,
				types.DoNotModify,
				types.NewPictures(
					types.DoNotModify,
					types.DoNotModify,
				),
			),
			shouldErr: false,
			expProfile: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"nickname",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
		},
		{
			name: "empty fields are allowed",
			original: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"nickname",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			update: types.NewProfileUpdate(
				types.DoNotModify,
				"",
				"",
				types.NewPictures("", ""),
			),
			shouldErr: false,
			expProfile: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures("", ""),
				time.Unix(100, 0),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
		},
		{
			name: "all fields are updated correctly",
			original: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"nickname",
				"bio",
				types.NewPictures(
					"https://example.com",
					"https://example.com",
				),
				time.Unix(100, 0),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
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
			shouldErr: false,
			expProfile: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag-2",
				"nickname-2",
				"bio-2",
				types.NewPictures(
					"https://example.com/2",
					"https://example.com/2",
				),
				time.Unix(100, 0),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			updated, err := tc.original.Update(tc.update)

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expProfile, updated)
			}
		})
	}
}

func TestProfile_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		account   *types.Profile
		shouldErr bool
	}{
		{
			name: "empty profile creator returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
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
			shouldErr: true,
		},
		{
			name: "empty profile DTag returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"",
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: true,
		},
		{
			name: "setting DTag to DoNotModify returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				types.DoNotModify,
				"",
				"bio",
				types.NewPictures(
					"https://shorturl.at/adnX3",
					"https://shorturl.at/cgpyF",
				),
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: true,
		},
		{
			name: "invalid profile picture returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("pic", "https://example.com"),
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: true,
		},
		{
			name: "invalid cover picture returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("https://example.com", "cov"),
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: true,
		},
		{
			name: "setting the nickname to DoNotModify returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				types.DoNotModify,
				"",
				types.Pictures{},
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: true,
		},
		{
			name: "setting the bio to DoNotModify returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"",
				types.DoNotModify,
				types.Pictures{},
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: true,
		},
		{
			name: "setting the profile picture to DoNotModify returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures(types.DoNotModify, ""),
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: true,
		},
		{
			name: "setting the profile cover to DoNotModify returns error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"",
				types.NewPictures("", types.DoNotModify),
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: true,
		},
		{
			name: "profile with only DTag does not error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"",
				types.Pictures{},
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: false,
		},
		{
			name: "valid profile returns no error",
			account: profilestesting.AssertNoProfileError(types.NewProfile(
				"dtag",
				"",
				"bio",
				types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
				time.Now(),
				profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			)),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.account.Validate()

			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestProfileSerialization(t *testing.T) {

	cdc := app.MakeEncodingConfig().Codec

	// Create a profile
	protoAccount := &authtypes.BaseAccount{
		Address:       "",
		PubKey:        nil,
		AccountNumber: 0,
		Sequence:      0,
	}
	accountAny, err := codectypes.NewAnyWithValue(protoAccount)
	require.NoError(t, err)

	profile := &types.Profile{
		Account: accountAny,
	}

	bz, err := cdc.MarshalInterface(profile)
	require.NoError(t, err)

	var original sdk.AccountI
	err = cdc.UnmarshalInterface(bz, &original)
	require.NoError(t, err)

	// Update the data
	addr2, err := sdk.AccAddressFromBech32("cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k")
	err = profile.SetAddress(addr2)
	require.NoError(t, err)

	pubKey := profilestesting.PubKeyFromBech32(
		"cosmospub1addwnpepqtkndttcutq2sehejxs2x3jl2uhxzuds4705u8nkgayuct0khqkzjd0vvln",
	)
	err = profile.SetPubKey(pubKey)
	require.NoError(t, err)

	err = profile.SetAccountNumber(100)
	require.NoError(t, err)

	err = profile.SetSequence(20)
	require.NoError(t, err)

	// Serialize
	bz, err = cdc.MarshalInterface(profile)
	require.NoError(t, err)

	// Deserialize
	var serialized sdk.AccountI
	err = cdc.UnmarshalInterface(bz, &serialized)
	require.NoError(t, err)

	// Check the data
	require.False(t, serialized.GetAddress().Equals(original.GetAddress()), "address not updated")
	require.NotEqual(t, serialized.GetPubKey(), original.GetPubKey(), "pub key not updated")
	require.NotEqual(t, serialized.GetAccountNumber(), original.GetAccountNumber(), "account number not updated")
	require.NotEqual(t, serialized.GetSequence(), original.GetSequence(), "sequence not updated")

	require.True(t, profile.GetAddress().Equals(serialized.GetAddress()), "addresses do not match")
	require.Equal(t, profile.GetPubKey(), serialized.GetPubKey(), "pub keys do not match")
	require.Equal(t, profile.GetAccountNumber(), serialized.GetAccountNumber(), "account numbers do not match")
	require.Equal(t, profile.GetSequence(), serialized.GetSequence(), "sequences do not match")
}

func TestProfile_MarshalYAML(t *testing.T) {
	// Create a profile
	account := &authtypes.BaseAccount{
		Address: "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k",
		PubKey: codectypes.UnsafePackAny(
			profilestesting.PubKeyFromBech32(
				"cosmospub1addwnpepqtkndttcutq2sehejxs2x3jl2uhxzuds4705u8nkgayuct0khqkzjd0vvln",
			),
		),
		AccountNumber: 0,
		Sequence:      0,
	}
	profile := profilestesting.AssertNoProfileError(
		types.NewProfile(
			"dtag",
			"nickname",
			"bio",
			types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
			account,
		),
	)

	bz, err := profile.MarshalYAML()
	require.NoError(t, err)

	require.Equal(
		t,
		`account_number: 0
address: cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k
bio: bio
creation_date: "2023-09-01T00:00:00Z"
dtag: dtag
nickname: nickname
pictures:
  cover: https://shorturl.at/cgpyF
  profile: https://shorturl.at/adnX3
public_key: PubKeySecp256k1{02ED36AD78E2C0A866F991A0A3465F572E6171B0AF9F4E1E764749CC2DF6B82C29}
sequence: 0
`,
		bz,
	)
}

func TestProfile_MarshalJSON(t *testing.T) {
	// Create a profile
	account := &authtypes.BaseAccount{
		Address: "cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k",
		PubKey: codectypes.UnsafePackAny(
			profilestesting.PubKeyFromBech32(
				"cosmospub1addwnpepqtkndttcutq2sehejxs2x3jl2uhxzuds4705u8nkgayuct0khqkzjd0vvln",
			),
		),
		AccountNumber: 0,
		Sequence:      0,
	}
	profile := profilestesting.AssertNoProfileError(
		types.NewProfile(
			"dtag",
			"nickname",
			"bio",
			types.NewPictures("https://shorturl.at/adnX3", "https://shorturl.at/cgpyF"),
			time.Date(2023, 9, 1, 0, 0, 0, 0, time.UTC),
			account,
		),
	)

	bz, err := profile.MarshalJSON()
	require.NoError(t, err)

	require.Equal(
		t,
		"{\"address\":\"cosmos1tdgrkvx2qgjk0uqsmdhm6dcz6wvwh9f8t37x0k\",\"public_key\":\"PubKeySecp256k1{02ED36AD78E2C0A866F991A0A3465F572E6171B0AF9F4E1E764749CC2DF6B82C29}\",\"account_number\":0,\"sequence\":0,\"dtag\":\"dtag\",\"nickname\":\"nickname\",\"bio\":\"bio\",\"pictures\":{\"profile\":\"https://shorturl.at/adnX3\",\"cover\":\"https://shorturl.at/cgpyF\"},\"creation_date\":\"2023-09-01T00:00:00Z\"}",
		string(bz),
	)
}

func BenchmarkProfile_Update(b *testing.B) {
	profile := profilestesting.AssertNoProfileError(types.NewProfile(
		"dtag",
		"nickname",
		"bio",
		types.NewPictures(
			"https://example.com",
			"https://example.com",
		),
		time.Unix(100, 0),
		profilestesting.AccountFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
	))

	update := types.NewProfileUpdate(
		"dtag-2",
		"nickname-2",
		"bio-2",
		types.NewPictures(
			"https://example.com/2",
			"https://example.com/2",
		),
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = profile.Update(update)
	}
}
