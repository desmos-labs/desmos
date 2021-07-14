package keeper_test

import (
	"strings"
	"time"

	"github.com/desmos-labs/desmos/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/x/staging/subspaces/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_StoreProfile() {
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
				profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				profile.DTag = "DTag"
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			profile: suite.CheckProfileNoError(types.NewProfile(
				"DTag",
				"",
				"",
				types.NewPictures("", ""),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				testutil.AccountFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
			),
			shouldErr: true,
		},
		{
			name: "existing profile is updated correctly",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				profile.DTag = "Old DTag"
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))

				// Save a DTag transfer request
				request := types.NewDTagTransferRequest(
					"DTag",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			profile:   testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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

			err := suite.k.StoreProfile(ctx, tc.profile)
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
				profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			address:    "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			shouldErr:  false,
			expFound:   true,
			expProfile: testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
				profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				profile.DTag = "DTag"
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
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
				profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
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
				testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
				testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
				testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
				testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
				testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
				testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
				testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
				testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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

func (suite *KeeperTestSuite) TestKeeper_DeleteUnregisteredUserFromSubspace() {
	ctx, _ := suite.ctx.CacheContext()

	suite.sk.AddSubspaceUnregisteredPair(ctx, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	// Init relationships
	relationships := []types.Relationship{
		types.NewRelationship(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewRelationship(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewRelationship(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}
	suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
	for _, rel := range relationships {
		suite.Require().NoError(suite.k.SaveRelationship(ctx, rel))
	}

	// Init blocks
	blocks := []types.UserBlock{
		types.NewUserBlock(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			"reason",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewUserBlock(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			"reason",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewUserBlock(
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
			"reason",
			"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}
	suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
	for _, block := range blocks {
		suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
	}

	suite.k.DeleteUnregisteredUserFromSubspace(ctx)

	// Check result
	suite.Require().Equal(
		[]types.Relationship{
			types.NewRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
				"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
		}, suite.k.GetAllRelationships(ctx))

	suite.Require().Equal(
		[]types.UserBlock{
			types.NewUserBlock(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
				"reason",
				"5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
		}, suite.k.GetAllUsersBlocks(ctx))

	var pairs []subspacestypes.UnregisteredPair
	suite.sk.IterateUnregisteredPairs(ctx, func(_ int64, pair subspacestypes.UnregisteredPair) (stop bool) {
		pairs = append(pairs, pair)
		return false
	})
	suite.Require().Empty(pairs)
}
