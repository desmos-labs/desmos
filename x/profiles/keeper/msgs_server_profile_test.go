package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/testutil"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestMsgServer_SaveProfile() {
	blockTime := time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC)
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgSaveProfile
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "profile saved (with no previous profile created)",
			store: func(ctx sdk.Context) {
				suite.ak.SetAccount(ctx, testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"))
			},
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				"my-nickname",
				"my-bio",
				"https://tc.com/profile-picture",
				"https://tc.com/cover-pic",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			),
			check: func(ctx sdk.Context) {
				_, found, err := suite.k.GetProfile(ctx, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(err)
				suite.Require().True(found)
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDTag, "custom_dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeProfileCreationTime, blockTime.Format(time.RFC3339)),
				),
			},
		},
		{
			name: "profile saved (with previous profile created)",
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				profile := suite.CheckProfileNoError(types.NewProfile(
					"test_dtag",
					"old-nickname",
					"old-biography",
					types.NewPictures(
						"https://tc.com/old-profile-pic",
						"https://tc.com/old-cover-pic",
					),
					blockTime,
					testutil.AccountFromAddr(address),
				))
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			msg: types.NewMsgSaveProfile(
				"other_dtag",
				"nickname",
				"biography",
				"https://tc.com/profile-pic",
				"https://tc.com/cover-pic",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDTag, "other_dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeProfileCreationTime, blockTime.Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				profile, found, err := suite.k.GetProfile(ctx, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Equal("other_dtag", profile.DTag)
				suite.Require().Equal("nickname", profile.Nickname)
			},
		},
		{
			name: "profile saved with same DTag but capital first letter (with previous profile created)",
			store: func(ctx sdk.Context) {
				profile := suite.CheckProfileNoError(types.NewProfile(
					"tc",
					"old-nickname",
					"old-biography",
					types.NewPictures(
						"https://tc.com/old-profile-pic",
						"https://tc.com/old-cover-pic",
					),
					blockTime,
					testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
				))
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			msg: types.NewMsgSaveProfile(
				"Test",
				"nickname",
				"biography",
				"https://tc.com/profile-pic",
				"https://tc.com/cover-pic",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDTag, "Test"),
					sdk.NewAttribute(types.AttributeProfileCreator, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeProfileCreationTime, blockTime.Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				profile, found, err := suite.k.GetProfile(ctx, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Equal("Test", profile.DTag)
				suite.Require().Equal("nickname", profile.Nickname)
			},
		},
		{
			name: "profile not saved because of the same DTag",
			store: func(ctx sdk.Context) {
				profile := suite.CheckProfileNoError(types.NewProfile(
					"tc",
					"nickname",
					"biography",
					types.NewPictures(
						"https://tc.com/profile-pic",
						"https://tc.com/cover-pic",
					),
					blockTime,
					testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
				))
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			msg: types.NewMsgSaveProfile(
				"Test",
				"another-one",
				"biography",
				"https://tc.com/profile-pic",
				"https://tc.com/cover-pic",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "profile not edited because of the invalid profile picture",
			store: func(ctx sdk.Context) {
				profile := suite.CheckProfileNoError(types.NewProfile(
					"custom_dtag",
					"biography",
					"",
					types.NewPictures("", ""),
					blockTime,
					testutil.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
				))
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			msg: types.NewMsgSaveProfile(
				"custom_dtag",
				"",
				"",
				"invalid-pic",
				"",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			),
			shouldErr: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			ctx = ctx.WithBlockTime(blockTime)

			suite.k.SetParams(ctx, types.DefaultParams())
			if tc.store != nil {
				tc.store(ctx)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.SaveProfile(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_DeleteProfile() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgDeleteProfile
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name:      "non existent profile returns error",
			msg:       types.NewMsgDeleteProfile("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			shouldErr: true,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "existent profile is deleted successfully",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.StoreProfile(ctx, profile))
			},
			msg:       types.NewMsgDeleteProfile("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileCreator, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
				),
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

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.DeleteProfile(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}
