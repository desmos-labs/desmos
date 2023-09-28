package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"

	"github.com/desmos-labs/desmos/v6/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func (suite *KeeperTestSuite) TestMsgServer_NewMsgServerImpl() {
	suite.k.ChannelKeeper = nil
	server := keeper.NewMsgServerImpl(suite.k)
	suite.k.ChannelKeeper = suite.channelKeeper

	// Make sure keeper set ChannelKeeper properly after msg server initialized.
	suite.Require().Equal(suite.channelKeeper, server.(*keeper.MsgServer).Keeper.ChannelKeeper)
}

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
			name: "account not found returns error (with no previous profile created)",
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
			shouldErr: true,
		},
		{
			name: "profile saved (with no previous profile created)",
			store: func(ctx sdk.Context) {
				suite.ak.SetAccount(ctx, profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"))
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
					sdk.NewAttribute(types.AttributeKeyProfileDTag, "custom_dtag"),
					sdk.NewAttribute(types.AttributeKeyProfileCreator, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyProfileCreationTime, blockTime.Format(time.RFC3339)),
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
					profilestesting.AccountFromAddr(address),
				))
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
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
					sdk.NewAttribute(types.AttributeKeyProfileDTag, "other_dtag"),
					sdk.NewAttribute(types.AttributeKeyProfileCreator, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyProfileCreationTime, blockTime.Format(time.RFC3339)),
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
					"tomtom",
					"old-nickname",
					"old-biography",
					types.NewPictures(
						"https://tc.com/old-profile-pic",
						"https://tc.com/old-cover-pic",
					),
					blockTime,
					profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
				))
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
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
					sdk.NewAttribute(types.AttributeKeyProfileDTag, "Test"),
					sdk.NewAttribute(types.AttributeKeyProfileCreator, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyProfileCreationTime, blockTime.Format(time.RFC3339)),
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
					"Test",
					"nickname",
					"biography",
					types.NewPictures(
						"https://tc.com/profile-pic",
						"https://tc.com/cover-pic",
					),
					blockTime,
					profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
				))
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
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
			name: "profile not created because DTag is set to DoNotModify",
			msg: types.NewMsgSaveProfile(
				types.DoNotModify,
				"another-one",
				"biography",
				"https://tc.com/profile-pic",
				"https://tc.com/cover-pic",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
		},
		{
			name: "profile updated correctly with DTag set to DoNotModify",
			store: func(ctx sdk.Context) {
				profile := suite.CheckProfileNoError(types.NewProfile(
					"tomtom",
					"nickname",
					"biography",
					types.NewPictures(
						"https://tc.com/profile-pic",
						"https://tc.com/cover-pic",
					),
					blockTime,
					profilestesting.AccountFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				))
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
			},
			msg: types.NewMsgSaveProfile(
				types.DoNotModify,
				"another-one",
				"biography",
				"https://tc.com/profile-pic",
				"https://tc.com/cover-pic",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeKeyProfileDTag, "tomtom"),
					sdk.NewAttribute(types.AttributeKeyProfileCreator, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeyProfileCreationTime, blockTime.Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				profile, found, err := suite.k.GetProfile(ctx, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Equal("tomtom", profile.DTag)
				suite.Require().Equal("another-one", profile.Nickname)
			},
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
					profilestesting.AccountFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
				))
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
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
			_, err := server.SaveProfile(ctx, tc.msg)

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
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.Require().NoError(suite.k.SaveProfile(ctx, profile))
			},
			msg:       types.NewMsgDeleteProfile("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeKeyProfileCreator, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
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
			_, err := server.DeleteProfile(ctx, tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())
			}
		})
	}
}
