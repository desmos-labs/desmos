package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_SaveApplicationLink() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		link      types.ApplicationLink
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "user without profile returns error",
			link: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "correct requests returns no error",
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
			},
			link: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasApplicationLink(ctx,
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"twitter",
					"twitteruser",
				))

				// Check the additional keys
				store := ctx.KVStore(suite.storeKey)
				suite.Require().True(store.Has(types.ApplicationLinkClientIDKey("client_id")))
				suite.Require().True(store.Has(types.ApplicationLinkOwnerKey(
					"twitter",
					"twitteruser",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				)))
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

			err := suite.k.SaveApplicationLink(ctx, tc.link)
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

func (suite *KeeperTestSuite) Test_GetApplicationLink() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		user        string
		application string
		username    string
		expFound    bool
		expLink     types.ApplicationLink
	}{
		{
			name: "different user does not find link",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			application: "twitter",
			username:    "twitteruser",
			expFound:    false,
		},
		{
			name: "different application does not find link",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "github",
			username:    "twitteruser",
			expFound:    false,
		},
		{
			name: "different application username does not find link",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "twitter",
			username:    "twitter-user",
			expFound:    false,
		},
		{
			name: "correct data returns proper link",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "twitter",
			username:    "twitteruser",
			expFound:    true,
			expLink: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			link, found, err := suite.k.GetApplicationLink(ctx, tc.user, tc.application, tc.username)
			suite.Require().Equal(tc.expFound, found)
			suite.Require().NoError(err)

			if tc.expFound {
				suite.Require().Equal(tc.expLink, link)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_GetApplicationLinkByClientID() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		clientID  string
		expFound  bool
		shouldErr bool
		expLink   types.ApplicationLink
	}{
		{
			name:      "invalid client id returns false",
			clientID:  "client_id",
			expFound:  false,
			shouldErr: false,
		},
		{
			name: "valid client id returns proper data",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))

				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			expFound:  true,
			shouldErr: false,
			clientID:  "client_id",
			expLink: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			link, found, err := suite.k.GetApplicationLinkByClientID(ctx, tc.clientID)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expFound, found)
				if tc.expFound {
					suite.Require().Equal(tc.expLink, link)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_DeleteApplicationLink() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		link  types.ApplicationLink
		check func(ctx sdk.Context)
	}{
		{
			name: "different user does not delete link",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			link: types.NewApplicationLink(
				"cosmos1xvvggrlgjkhu4rva9j500rc52za2smxhluvftc",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasApplicationLink(ctx,
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"twitter",
					"twitteruser",
				))
			},
		},
		{
			name: "different application does not delete the link",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			link: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("github", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasApplicationLink(ctx,
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"twitter",
					"twitteruser",
				))
			},
		},
		{
			name: "different username does not delete the link",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			link: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("twitter", "another-user"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasApplicationLink(ctx,
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"twitter",
					"twitteruser",
				))
			},
		},
		{
			name: "valid request deletes link",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				suite.Require().NoError(suite.k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			link: types.NewApplicationLink(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasApplicationLink(ctx,
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"twitter",
					"twitteruser",
				))

				// Check the additional keys
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.ApplicationLinkClientIDKey("client_id")))
				suite.Require().False(store.Has(types.ApplicationLinkOwnerKey(
					"twitter",
					"twitteruser",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				)))
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

			suite.k.DeleteApplicationLink(ctx, tc.link)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
