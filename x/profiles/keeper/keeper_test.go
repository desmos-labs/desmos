package keeper_test

import (
	"strings"
	"time"

	"github.com/desmos-labs/desmos/v4/testutil/profilestesting"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveProfile() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		profile   *types.Profile
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "existent profile with different creator returns error",
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				profile.DTag = "DTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
			},
			profile: suite.CheckProfileNoError(types.NewProfile(
				"DTag",
				"",
				"",
				types.NewPictures("", ""),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
			),
			shouldErr: true,
		},
		{
			name: "existing profile is updated correctly",
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				profile.DTag = "Old DTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))

				// Save a DTag transfer request
				request := types.NewDTagTransferRequest(
					"DTag",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			profile:   profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				profile, found, err := suite.k.GetProfile(ctx, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Equal("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773-dtag", profile.DTag)

				// Verify the DTag transfer requests have been deleted
				suite.Require().Empty(suite.k.GetDTagTransferRequests(ctx))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.SaveProfile(ctx, tc.profile)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetProfile() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		address    string
		shouldErr  bool
		expFound   bool
		expProfile *types.Profile
	}{
		{
			name:      "invalid address returns error",
			address:   "",
			shouldErr: true,
		},
		{
			name: "found profile is returned properly",
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
			},
			address:    "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			shouldErr:  false,
			expFound:   true,
			expProfile: profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
		},
		{
			name:      "not found profile returns no error",
			address:   "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			shouldErr: false,
			expFound:  false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			profile, found, err := suite.k.GetProfile(ctx, tc.address)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expFound, found)

				if found {
					suite.Require().Equal(tc.expProfile, profile)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAddressFromDTag() {
	testCases := []struct {
		name    string
		store   func(ctx sdk.Context)
		DTag    string
		expAddr string
	}{
		{
			name: "valid profile returns correct address",
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				profile.DTag = "DTag"
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
			},
			DTag:    "dtag",
			expAddr: "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
		},
		{
			name:    "non existing profile returns empty address",
			DTag:    "test",
			expAddr: "",
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			addr := suite.k.GetAddressFromDTag(ctx, tc.DTag)
			suite.Require().Equal(tc.expAddr, addr)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveProfile() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		address   string
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name:      "non existent profile returns error",
			address:   "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			shouldErr: true,
		},
		{
			name: "found profile is deleted properly",
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
			},
			address:   "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			shouldErr: false,
			check: func(ctx sdk.Context) {
				_, found, err := suite.k.GetProfile(ctx, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(err)
				suite.Require().False(found)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.RemoveProfile(ctx, tc.address)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_ValidateProfile() {
	testCases := []struct {
		name      string
		profile   *types.Profile
		shouldErr bool
	}{
		{
			name: "max nickname length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				strings.Repeat("A", 1005),
				"my-bio",
				types.NewPictures(
					"https://tc.com/profile-picture",
					"https://tc.com/cover-pic",
				),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			)),
			shouldErr: true,
		},
		{
			name: "min nickname length not reached",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				"m",
				"my-bio",
				types.NewPictures(
					"https://tc.com/profile-picture",
					"https://tc.com/cover-pic",
				),

				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			)),
			shouldErr: true,
		},
		{
			name: "max bio length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				"nickname",
				strings.Repeat("A", 1005),
				types.NewPictures(
					"https://tc.com/profile-picture",
					"https://tc.com/cover-pic",
				),

				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			)),
			shouldErr: true,
		},
		{
			name: "invalid DTag doesn't match regEx",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom.",
				"nickname",
				strings.Repeat("A", 1000),
				types.NewPictures(
					"https://tc.com/profile-picture",
					"https://tc.com/cover-pic",
				),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			)),
			shouldErr: true,
		},
		{
			name: "min DTag length not reached",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"d",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://tc.com/profile-picture",
					"https://tc.com/cover-pic",
				),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			)),
			shouldErr: true,
		},
		{
			name: "max DTag length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"9YfrVVi3UEI1ymN7n6isSct30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEyHNS",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://tc.com/profile-picture",
					"https://tc.com/cover-pic",
				),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			)),
			shouldErr: true,
		},
		{
			name: "invalid profile pictures",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"dtag",
				"nickname",
				"my-bio",
				types.NewPictures(
					"pic",
					"htts://tc.com/cover-pic",
				),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			)),
			shouldErr: true,
		},
		{
			name: "valid profile",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"dtag",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://tc.com/profile-picture",
					"https://tc.com/cover-pic",
				),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			)),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			suite.k.SetParams(ctx, types.DefaultParams())

			err := suite.k.ValidateProfile(ctx, tc.profile)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
