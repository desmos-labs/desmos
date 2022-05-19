package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/keeper"
	"github.com/desmos-labs/desmos/v3/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestMsgServer_CreatePost() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		setupCtx    func(ctx sdk.Context) sdk.Context
		msg         *types.MsgCreatePost
		shouldErr   bool
		expResponse *types.MsgCreatePostResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgCreatePost(
				1,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgCreatePost(
				1,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid conversation id returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionWrite)

				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid reference returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionWrite)

				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				[]types.AttachmentContent{
					types.NewMedia("", ""),
				},
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "initial post id not set returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionWrite)
			},
			msg: types.NewMsgCreatePost(
				1,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid post returns error",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionWrite)

				suite.k.SetNextPostID(ctx, 1, 1)

				// Set the max post length to 1 character
				suite.k.SetParams(ctx, types.NewParams(1))
			},
			msg: types.NewMsgCreatePost(
				1,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid attachment returns error",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionWrite)

				suite.k.SetNextPostID(ctx, 1, 1)

				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				[]types.AttachmentContent{
					types.NewMedia("", ""),
				},
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "valid post is stored correctly",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionWrite)

				suite.k.SetNextPostID(ctx, 1, 1)

				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				[]types.AttachmentContent{
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				},
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: false,
			expResponse: &types.MsgCreatePostResponse{
				PostID:       1,
				CreationDate: time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreatePost{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeCreatePost,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyAuthor, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
					sdk.NewAttribute(types.AttributeKeyCreationTime, time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				// Check the post
				stored, found := suite.k.GetPost(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				), stored)

				// Check the attachments
				attachments := suite.k.GetPostAttachments(ctx, 1, 1)
				suite.Require().Equal([]types.Attachment{
					types.NewAttachment(
						1,
						1,
						1,
						types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
					),
				}, attachments)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.CreatePost(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_EditPost() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		setupCtx    func(ctx sdk.Context) sdk.Context
		msg         *types.MsgEditPost
		shouldErr   bool
		expResponse *types.MsgEditPostResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "not found post returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"This is a new post",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"This is a new post",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid update returns error",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"This is a new post",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				// Set max post text length to 1 character
				suite.k.SetParams(ctx, types.NewParams(1))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "post is updated correctly",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)

				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a new post",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: false,
			expResponse: &types.MsgEditPostResponse{
				EditDate: time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgEditPost{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeEditPost,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyLastEditTime, time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the post is what we are expecting
				stored, found := suite.k.GetPost(ctx, 1, 1)
				suite.Require().True(found)

				editDate := time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC)
				suite.Require().Equal(types.NewPost(
					1,
					1,
					"External ID",
					"This is my new text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					&editDate,
				), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.EditPost(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_DeletePost() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		setupCtx  func(ctx sdk.Context) sdk.Context
		msg       *types.MsgDeletePost
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name:      "non existing subspace returns error",
			msg:       types.NewMsgDeletePost(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
			shouldErr: true,
		},
		{
			name: "not found post returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionWrite&subspacestypes.PermissionEditOwnContent)
			},
			msg:       types.NewMsgDeletePost(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg:       types.NewMsgDeletePost(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
			shouldErr: true,
		},
		{
			name: "author cannot delete other user post",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg:       types.NewMsgDeletePost(1, 1, "cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4"),
			shouldErr: true,
		},
		{
			name: "moderator can delete post",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionModerateContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg:       types.NewMsgDeletePost(1, 1, "cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4"),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgDeletePost{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4"),
				),
				sdk.NewEvent(
					types.EventTypeDeletePost,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPost(ctx, 1, 1))
			},
		},
		{
			name: "author can delete post",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgDeletePost(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgDeletePost{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeDeletePost,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasPost(ctx, 1, 1))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.DeletePost(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestsuite) TestMsgServer_AddPostAttachment() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		setupCtx    func(ctx sdk.Context) sdk.Context
		msg         *types.MsgAddPostAttachment
		shouldErr   bool
		expResponse *types.MsgAddPostAttachmentResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgAddPostAttachment(
				1,
				1,
				types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "not found post returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgAddPostAttachment(
				1,
				1,
				types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgAddPostAttachment(
				1,
				1,
				types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "user without permissions returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgAddPostAttachment(
				1,
				1,
				types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid attachment returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)
			},
			msg: types.NewMsgAddPostAttachment(
				1,
				1,
				types.NewMedia("", ""),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "correct data is stored properly",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)
			},
			msg: types.NewMsgAddPostAttachment(
				1,
				1,
				types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: false,
			expResponse: &types.MsgAddPostAttachmentResponse{
				AttachmentID: 1,
				EditDate:     time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAddPostAttachment{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeAddPostAttachment,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyAttachmentID, "1"),
					sdk.NewAttribute(types.AttributeKeyLastEditTime, time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the post is updated properly
				post, found := suite.k.GetPost(ctx, 1, 1)
				suite.Require().True(found)

				updateDate := time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC)
				suite.Require().Equal(&updateDate, post.LastEditedDate)

				// Make sure the attachment is there
				stored, found := suite.k.GetAttachment(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewAttachment(
					1,
					1,
					1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.AddPostAttachment(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_RemovePostAttachment() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		setupCtx    func(ctx sdk.Context) sdk.Context
		msg         *types.MsgRemovePostAttachment
		shouldErr   bool
		expResponse *types.MsgRemovePostAttachmentResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "not found subspace returns error",
			msg: types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "not found post returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "user without permissions returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "user with permissions cannot delete other author attachment",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "not found attachment returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "moderator can delete attachment",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionModerateContent)

				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))
			},
			msg: types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
			),
			shouldErr: false,
			expResponse: &types.MsgRemovePostAttachmentResponse{
				EditDate: time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRemovePostAttachment{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4"),
				),
				sdk.NewEvent(
					types.EventTypeRemovePostAttachment,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyAttachmentID, "1"),
					sdk.NewAttribute(types.AttributeKeyLastEditTime, time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the post is updated properly
				post, found := suite.k.GetPost(ctx, 1, 1)
				suite.Require().True(found)

				updateDate := time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC)
				suite.Require().Equal(&updateDate, post.LastEditedDate)

				// Make sure the attachment is no longer there
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 1, 1))
			},
		},
		{
			name: "author can delete attachment",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionEditOwnContent)

				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))
			},
			msg: types.NewMsgRemovePostAttachment(
				1,
				1,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: false,
			expResponse: &types.MsgRemovePostAttachmentResponse{
				EditDate: time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC),
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRemovePostAttachment{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeRemovePostAttachment,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyAttachmentID, "1"),
					sdk.NewAttribute(types.AttributeKeyLastEditTime, time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the post is updated properly
				post, found := suite.k.GetPost(ctx, 1, 1)
				suite.Require().True(found)

				updateDate := time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC)
				suite.Require().Equal(&updateDate, post.LastEditedDate)

				// Make sure the attachment is no longer there
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 1, 1))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.RemovePostAttachment(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_AnswerPoll() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		setupCtx  func(ctx sdk.Context) sdk.Context
		msg       *types.MsgAnswerPoll
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "not found subspace returns error",
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{1, 2, 3},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{1, 2, 3},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "not found post returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionInteractWithContent)
			},
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{1, 2, 3},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "not found poll returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionInteractWithContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{1, 2, 3},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "already answered poll returns error if no answer edits are allowed",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionInteractWithContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					1,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						nil,
					),
				))

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{0, 1},
					user,
				))
			},
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{1},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "multiple answers return error if they are not allowed",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionInteractWithContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					1,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						nil,
					),
				))
			},
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{0, 1},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid answer indexes return error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionInteractWithContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					1,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						true,
						true,
						nil,
					),
				))
			},
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{0, 1, 2},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "editing an answer works correctly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionInteractWithContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					1,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						true,
						nil,
					),
				))

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{1},
					user,
				))
			},
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{0},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAnswerPoll{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeAnswerPoll,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyPollID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				// Check the user answer
				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)

				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, user)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{0},
					user,
				), stored)
			},
		},
		{
			name: "new answer is stored correctly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test",
					"Testing subspace",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					"cosmos1sg2j68v5n8qvehew6ml0etun3lmv7zg7r49s67",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)
				suite.sk.SetUserPermissions(ctx, 1, user, subspacestypes.PermissionInteractWithContent)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					1,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						true,
						false,
						nil,
					),
				))
			},
			msg: types.NewMsgAnswerPoll(
				1,
				1,
				1,
				[]uint32{0, 1},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAnswerPoll{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeAnswerPoll,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyPollID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				// Check the user answer
				user, err := sdk.AccAddressFromBech32("cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().NoError(err)

				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, user)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{0, 1},
					user,
				), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.AnswerPoll(sdk.WrapSDKContext(ctx), tc.msg)
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
