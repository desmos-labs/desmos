package keeper_test

import (
	"fmt"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_StoreProfile() {
	addr, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	accountAny, err := codectypes.NewAnyWithValue(authtypes.NewBaseAccountWithAddress(addr))
	suite.Require().NoError(err)

	updatedProfile, err := types.NewProfile(
		"updated-dtag",
		"updated-nickname",
		suite.testData.profile.Bio,
		suite.testData.profile.Pictures,
		suite.testData.profile.CreationDate,
		suite.testData.profile.GetAccount(),
	)
	suite.Require().NoError(err)

	tests := []struct {
		name           string
		account        *types.Profile
		storedProfiles []*types.Profile
		shouldErr      bool
	}{
		{
			name:      "Non existent profile is saved correctly",
			account:   suite.testData.profile,
			shouldErr: false,
		},
		{
			name: "Edited profile is saved correctly",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			account:   updatedProfile,
			shouldErr: false,
		},
		{
			name: "Existent account with different creator returns error",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			account: &types.Profile{
				DTag:     suite.testData.profile.DTag,
				Bio:      suite.testData.profile.Bio,
				Pictures: suite.testData.profile.Pictures,
				Account:  accountAny,
			},
			shouldErr: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, profile := range test.storedProfiles {
				err = suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			err = suite.k.StoreProfile(suite.ctx, test.account)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				// Verify the DTag -> Address association
				store := suite.ctx.KVStore(suite.storeKey)
				suite.Require().Equal(test.account.GetAddress().Bytes(), store.Get(types.DTagStoreKey(test.account.DTag)),
					"DTag -> Address association not correct")

				for _, stored := range test.storedProfiles {
					// Make sure that if the DTag has been edited, the old association has been removed
					if stored.GetAddress().Equals(test.account.GetAddress()) && stored.DTag != test.account.DTag {
						suite.Require().Nil(store.Get(types.DTagStoreKey(stored.DTag)),
							"Old DTag -> Address association still exists")
					}
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_StoreProfile_Update() {
	// Store the initial profile
	suite.Require().NoError(suite.k.StoreProfile(suite.ctx, suite.testData.profile))

	// Verify the store keys
	store := suite.ctx.KVStore(suite.storeKey)
	suite.Require().Equal(
		suite.testData.profile.GetAddress().String(),
		sdk.AccAddress(store.Get(types.DTagStoreKey(suite.testData.profile.DTag))).String(),
	)

	oldAccounts := suite.ak.GetAllAccounts(suite.ctx)
	suite.Require().Len(oldAccounts, 1)

	// Update the profile
	updatedProfile, err := types.NewProfile(
		suite.testData.profile.DTag+"-update",
		"",
		"",
		types.NewPictures("", ""),
		suite.testData.profile.CreationDate,
		suite.ak.GetAccount(suite.ctx, suite.testData.profile.GetAddress()),
	)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.k.StoreProfile(suite.ctx, updatedProfile))

	// Verify the store keys
	suite.Require().Nil(
		store.Get(types.DTagStoreKey(suite.testData.profile.DTag)),
	)
	suite.Require().Equal(
		suite.testData.profile.GetAddress().String(),
		sdk.AccAddress(store.Get(types.DTagStoreKey(suite.testData.profile.DTag+"-update"))).String(),
	)

	newAccounts := suite.ak.GetAllAccounts(suite.ctx)
	suite.Require().Len(newAccounts, 1)

	for _, account := range newAccounts {
		suite.Require().NotContains(oldAccounts, account)
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetProfile() {
	tests := []struct {
		name           string
		storedProfiles []*types.Profile
		address        string
		shouldErr      bool
		expFound       bool
		expProfile     *types.Profile
	}{
		{
			name:           "Invalid address",
			storedProfiles: nil,
			address:        "",
			shouldErr:      true,
		},
		{
			name: "Profile found",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			address:    suite.testData.profile.GetAddress().String(),
			shouldErr:  false,
			expFound:   true,
			expProfile: suite.testData.profile,
		},
		{
			name:           "Profile not found",
			storedProfiles: []*types.Profile{},
			address:        suite.testData.profile.GetAddress().String(),
			expFound:       false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, profile := range test.storedProfiles {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			res, found, err := suite.k.GetProfile(suite.ctx, test.address)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expFound, found)

				if found {
					suite.Require().Equal(test.expProfile, res)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAddressFromDTag() {
	tests := []struct {
		name    string
		profile *types.Profile
		dTag    string
		expAddr string
	}{
		{
			name:    "valid profile returns correct address",
			profile: suite.testData.profile,
			dTag:    suite.testData.profile.DTag,
			expAddr: suite.testData.profile.GetAddress().String(),
		},
		{
			name:    "non existing profile returns empty address",
			profile: nil,
			dTag:    "test",
			expAddr: "",
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.profile != nil {
				err := suite.k.StoreProfile(suite.ctx, test.profile)
				suite.Require().NoError(err)
			}

			addr := suite.k.GetAddressFromDTag(suite.ctx, test.dTag)
			suite.Require().Equal(test.expAddr, addr)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveProfile() {
	err := suite.k.StoreProfile(suite.ctx, suite.testData.profile)
	suite.Require().Nil(err)

	_, found, _ := suite.k.GetProfile(suite.ctx, suite.testData.profile.GetAddress().String())
	suite.True(found)

	err = suite.k.RemoveProfile(suite.ctx, suite.testData.profile.GetAddress().String())
	suite.Require().NoError(err)

	_, found, _ = suite.k.GetProfile(suite.ctx, suite.testData.profile.GetAddress().String())
	suite.Require().False(found)

	addr := suite.k.GetAddressFromDTag(suite.ctx, suite.testData.profile.DTag)
	suite.Require().Equal("", addr)
}

func (suite *KeeperTestSuite) TestKeeper_ValidateProfile() {
	tests := []struct {
		name    string
		profile *types.Profile
		expErr  error
	}{
		{
			name: "Max nickname length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				strings.Repeat("A", 1005),
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile nickname cannot exceed 1000 characters"),
		},
		{
			name: "Min nickname length not reached",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				"m",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile nickname cannot be less than 2 characters"),
		},
		{
			name: "Max bio length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				"nickname",
				strings.Repeat("A", 1005),
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Invalid dtag doesn't match regEx",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom.",
				"nickname",
				strings.Repeat("A", 1000),
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("invalid profile dtag, it should match the following regEx ^[A-Za-z0-9_]+$"),
		},
		{
			name: "Min dtag length not reached",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"d",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile dtag cannot be less than 3 characters"),
		},
		{
			name: "Max dtag length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"9YfrVVi3UEI1ymN7n6isSct30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEyHNS",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile dtag cannot exceed 30 characters"),
		},
		{
			name: "Invalid profile pictures returns error",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"dtag",
				"nickname",
				"my-bio",
				types.NewPictures(
					"pic",
					"htts://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Valid profile returns no error",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"dtag",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.k.SetParams(suite.ctx, types.DefaultParams())

			actual := suite.k.ValidateProfile(suite.ctx, test.profile)
			suite.Require().Equal(test.expErr, actual)
		})
	}
}
