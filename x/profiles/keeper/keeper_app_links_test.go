package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v2/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_SaveApplicationLink() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		link      types.ApplicationLink
		shouldErr bool
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
			),
			shouldErr: true,
		},
		{
			name: "correct requests returns no error",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
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
			),
			shouldErr: false,
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

				store := ctx.KVStore(suite.storeKey)
				suite.Require().True(store.Has(types.UserApplicationLinkKey(tc.link.User, tc.link.Data.Application, tc.link.Data.Username)))
				suite.Require().True(store.Has(types.ApplicationLinkClientIDKey(tc.link.OracleRequest.ClientID)))
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))

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
		name        string
		store       func(store sdk.Context)
		user        string
		application string
		username    string
		shouldErr   bool
	}{
		{
			name: "wrong user returns error",
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			user:        "user",
			application: "twitter",
			username:    "twitteruser",
			shouldErr:   true,
		},
		{
			name: "wrong application returns error",
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "github",
			username:    "twitteruser",
			shouldErr:   true,
		},
		{
			name: "wrong username returns error",
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "twitter",
			username:    "twitter-user",
			shouldErr:   true,
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
				)

				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(address)))
				err := suite.k.SaveApplicationLink(ctx, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "twitter",
			username:    "twitteruser",
			shouldErr:   false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			err := suite.k.DeleteApplicationLink(ctx, tc.user, tc.application, tc.username)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				_, found, err := suite.k.GetApplicationLink(ctx, tc.user, tc.application, tc.username)
				suite.Require().NoError(err)
				suite.Require().False(found)
			}
		})
	}
}
