package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_SaveApplicationLink() {
	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		user      string
		link      types.ApplicationLink
		shouldErr bool
	}{
		{
			name:      "user without profile returns error",
			user:      "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			shouldErr: true,
		},
		{
			name: "correct requests returns no error",
			store: func(ctx sdk.Context) {
				profile := suite.CreateProfileFromAddress("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.ak.SetAccount(ctx, profile)
			},
			user: "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			link: types.NewApplicationLink(
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					-1,
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

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			err := suite.k.SaveApplicationLink(ctx, uc.user, uc.link)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				store := ctx.KVStore(suite.storeKey)
				suite.Require().True(store.Has(types.ApplicationLinkKey(uc.user, uc.link.Data.Application, uc.link.Data.Username)))
				suite.Require().True(store.Has(types.ApplicationLinkClientIDKey(uc.link.OracleRequest.ClientID)))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_GetApplicationLink() {
	usecases := []struct {
		name        string
		store       func(ctx sdk.Context)
		user        string
		application string
		username    string
		shouldErr   bool
		expLink     types.ApplicationLink
	}{
		{
			name: "different user returns error",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))
				err := suite.k.SaveApplicationLink(ctx, address, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			application: "twitter",
			username:    "twitteruser",
			shouldErr:   true,
		},
		{
			name: "different application returns error",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))
				err := suite.k.SaveApplicationLink(ctx, address, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "github",
			username:    "twitteruser",
			shouldErr:   true,
		},
		{
			name: "different application username returns error",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))
				err := suite.k.SaveApplicationLink(ctx, address, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "twitter",
			username:    "twitter-user",
			shouldErr:   true,
		},
		{
			name: "correct data returns no error",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))
				err := suite.k.SaveApplicationLink(ctx, address, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "twitter",
			username:    "twitteruser",
			shouldErr:   false,
			expLink: types.NewApplicationLink(
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					-1,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			link, err := suite.k.GetApplicationLink(ctx, uc.user, uc.application, uc.username)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(uc.expLink, link)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_GetApplicationLinkByClientID() {
	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		clientID  string
		shouldErr bool
		expUser   string
		expLink   types.ApplicationLink
	}{
		{
			name:      "invalid client id returns error",
			clientID:  "client_id",
			shouldErr: true,
		},
		{
			name: "valid client id returns proper data",
			store: func(ctx sdk.Context) {
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))

				err := suite.k.SaveApplicationLink(ctx, address, link)
				suite.Require().NoError(err)
			},
			shouldErr: false,
			clientID:  "client_id",
			expUser:   "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			expLink: types.NewApplicationLink(
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					-1,
					1,
					types.NewOracleRequestCallData("twitter", "calldata"),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			user, link, err := suite.k.GetApplicationLinkByClientID(ctx, uc.clientID)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(uc.expUser, user)
				suite.Require().Equal(uc.expLink, link)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_DeleteApplicationLink() {
	usecases := []struct {
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
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))

				err := suite.k.SaveApplicationLink(ctx, address, link)
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
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))

				err := suite.k.SaveApplicationLink(ctx, address, link)
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
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))

				err := suite.k.SaveApplicationLink(ctx, address, link)
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
				link := types.NewApplicationLink(
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))

				err := suite.k.SaveApplicationLink(ctx, address, link)
				suite.Require().NoError(err)
			},
			user:        "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			application: "twitter",
			username:    "twitteruser",
			shouldErr:   false,
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			err := suite.k.DeleteApplicationLink(ctx, uc.user, uc.application, uc.username)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				_, err := suite.k.GetApplicationLink(ctx, uc.user, uc.application, uc.username)
				suite.Require().Error(err)
			}
		})
	}
}
