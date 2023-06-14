package keeper_test

import (
	"time"

	"github.com/golang/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v5/x/posts/keeper"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestMsgServer_CreatePost() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		setupCtx    func(ctx sdk.Context) sdk.Context
		msg         *types.MsgCreatePost
		shouldErr   bool
		expResponse *types.MsgCreatePostResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "user without profile returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(false)
			},
			msg: types.NewMsgCreatePost(
				1,
				1,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			msg: types.NewMsgCreatePost(
				1,
				1,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "non existing section returns error",
			msg: types.NewMsgCreatePost(
				1,
				1,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(1)).Return(false)
			},
			shouldErr: true,
		},
		{
			name: "user without permissions returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionWrite),
				).Return(false)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionComment),
				).Return(false)
			},
			msg: types.NewMsgCreatePost(
				1,
				0,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid conversation id returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionWrite),
				).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionComment),
				).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				0,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid reference returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionWrite),
				).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				0,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				[]types.AttachmentContent{
					types.NewMedia("", ""),
				},
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
				},
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "initial post id not set returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionWrite),
				).Return(true)
			},
			msg: types.NewMsgCreatePost(
				1,
				0,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
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
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionWrite),
				).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)

				// Set the max post length to 1 character
				suite.k.SetParams(ctx, types.NewParams(1))
			},
			msg: types.NewMsgCreatePost(
				1,
				0,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				nil,
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid attachment returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionWrite),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)

				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				0,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
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
			name: "valid post is stored correctly with PermissionWrite",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionWrite),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)

				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				0,
				"External ID",
				"This is a text",
				0,
				types.REPLY_SETTING_EVERYONE,
				nil,
				[]string{"generic"},
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
					sdk.NewAttribute(types.AttributeKeySectionID, "0"),
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
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					[]string{"generic"},
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
		{
			name: "valid comment is stored correctly with PermissionComment",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionWrite),
				).Return(false)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionComment),
				).Return(true)

				suite.rk.EXPECT().HasUserBlocked(gomock.Any(),
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					uint64(1)).Return(false)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
				suite.k.SetNextPostID(ctx, 1, 2)

				suite.k.SetParams(ctx, types.DefaultParams())
			},
			msg: types.NewMsgCreatePost(
				1,
				0,
				"External ID",
				"This is a text",
				1,
				types.REPLY_SETTING_EVERYONE,
				nil,
				[]string{"generic"},
				[]types.AttachmentContent{
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				},
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: false,
			expResponse: &types.MsgCreatePostResponse{
				PostID:       2,
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
					sdk.NewAttribute(types.AttributeKeySectionID, "0"),
					sdk.NewAttribute(types.AttributeKeyPostID, "2"),
					sdk.NewAttribute(types.AttributeKeyAuthor, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
					sdk.NewAttribute(types.AttributeKeyCreationTime, time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				// Check the post
				stored, found := suite.k.GetPost(ctx, 1, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPost(
					1,
					0,
					2,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					[]string{"generic"},
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				), stored)

				// Check the attachments
				attachments := suite.k.GetPostAttachments(ctx, 1, 2)
				suite.Require().Equal([]types.Attachment{
					types.NewAttachment(
						1,
						2,
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
			if tc.setup != nil {
				tc.setup()
			}
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

func (suite *KeeperTestSuite) TestMsgServer_EditPost() {
	testCases := []struct {
		name        string
		setup       func()
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "not found post returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid editor returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"This is a new post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"This is a new post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid update returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"This is a new post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))

				// Set max post text length to 1 character
				suite.k.SetParams(ctx, types.NewParams(1))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				nil,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "post is updated correctly",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a new post",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			msg: types.NewMsgEditPost(
				1,
				1,
				"This is my new text",
				nil,
				[]string{"generic"},
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
					0,
					1,
					"External ID",
					"This is my new text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					[]string{"generic"},
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					&editDate,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
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

func (suite *KeeperTestSuite) TestMsgServer_DeletePost() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		setupCtx  func(ctx sdk.Context) sdk.Context
		msg       *types.MsgDeletePost
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			msg:       types.NewMsgDeletePost(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
			shouldErr: true,
		},
		{
			name: "not found post returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			msg:       types.NewMsgDeletePost(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(false)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			msg:       types.NewMsgDeletePost(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
			shouldErr: true,
		},
		{
			name: "author cannot delete other user post",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			msg:       types.NewMsgDeletePost(1, 1, "cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4"),
			shouldErr: true,
		},
		{
			name: "moderator can delete post",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(false)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
			if tc.setup != nil {
				tc.setup()
			}
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

func (suite *KeeperTestSuite) TestMsgServer_AddPostAttachment() {
	testCases := []struct {
		name        string
		setup       func()
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
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
			name: "not found post returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
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
			if tc.setup != nil {
				tc.setup()
			}
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

func (suite *KeeperTestSuite) TestMsgServer_RemovePostAttachment() {
	testCases := []struct {
		name        string
		setup       func()
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
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
			name: "not found post returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
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
			name: "user without PermissionModerateContent returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
			name: "user without PermissionEditOwnContent returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(false)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(false)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(false)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionModerateContent),
				).Return(false)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionEditOwnContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
			if tc.setup != nil {
				tc.setup()
			}
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

func (suite *KeeperTestSuite) TestMsgServer_AnswerPoll() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		setupCtx  func(ctx sdk.Context) sdk.Context
		msg       *types.MsgAnswerPoll
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "user without profile returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(false)
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
			name: "not found subspace returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
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
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
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
			name: "user without permission returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionInteractWithContent),
				).Return(false)
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
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionInteractWithContent),
				).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
			name: "voting after end time returns error",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionInteractWithContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2100, 1, 1, 00, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
			name: "already answered poll returns error if no answer edits are allowed",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionInteractWithContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2010, 1, 1, 00, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionInteractWithContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2010, 1, 1, 00, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionInteractWithContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2010, 1, 1, 00, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)

				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionInteractWithContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2010, 1, 1, 00, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
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
				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{0},
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				), stored)
			},
		},
		{
			name: "new answer is stored correctly",
			setup: func() {
				suite.ak.EXPECT().HasProfile(gomock.Any(), "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd").Return(true)

				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasPermission(
					gomock.Any(),
					uint64(1),
					uint32(0),
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					subspacestypes.NewPermission(types.PermissionInteractWithContent),
				).Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2010, 1, 1, 00, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos19mkklc8arp6phlg5eydu3v49syyqyfrq2sp4at",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd")
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{0, 1},
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
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

func (suite *KeeperTestSuite) TestMsgServer_UpdateParams() {
	testCases := []struct {
		name      string
		msg       *types.MsgUpdateParams
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "invalid authority return error",
			msg: types.NewMsgUpdateParams(
				types.DefaultParams(),
				"invalid",
			),
			shouldErr: true,
		},
		{
			name: "set params properly",
			msg: types.NewMsgUpdateParams(
				types.DefaultParams(),
				authtypes.NewModuleAddress("gov").String(),
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgUpdateParams{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos10d07y265gmmuvt4z0w9aw880jnsr700j6zn9kn"),
				),
			},
			check: func(ctx sdk.Context) {
				params := suite.k.GetParams(ctx)
				suite.Require().Equal(types.DefaultParams(), params)
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			// Reset any event that might have been emitted during the setup
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.UpdateParams(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_MovePost() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		setupCtx  func(ctx sdk.Context) sdk.Context
		msg       *types.MsgMovePost
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "post not found returns error",
			msg: types.NewMsgMovePost(
				1,
				1,
				2,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "target subspace not exist returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(2)).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			msg: types.NewMsgMovePost(
				1,
				1,
				2,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "target section not exist returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(2)).
					Return(true)
				suite.sk.EXPECT().
					HasSection(gomock.Any(), uint64(2), uint32(1)).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			msg: types.NewMsgMovePost(
				1,
				1,
				2,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "sender does not match post owner returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(2)).
					Return(true)
				suite.sk.EXPECT().
					HasSection(gomock.Any(), uint64(2), uint32(1)).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			msg: types.NewMsgMovePost(
				1,
				1,
				2,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "sender without permission inside target subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(2)).
					Return(true)
				suite.sk.EXPECT().
					HasSection(gomock.Any(), uint64(2), uint32(1)).
					Return(true)
				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(2),
						uint32(1),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						types.PermissionWrite,
					).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			msg: types.NewMsgMovePost(
				1,
				1,
				2,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "next id not set returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(2)).
					Return(true)
				suite.sk.EXPECT().
					HasSection(gomock.Any(), uint64(2), uint32(1)).
					Return(true)
				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(2),
						uint32(1),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						types.PermissionWrite,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(1))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			msg: types.NewMsgMovePost(
				1,
				1,
				2,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "invalid post returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(2)).
					Return(true)
				suite.sk.EXPECT().
					HasSection(gomock.Any(), uint64(2), uint32(1)).
					Return(true)
				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(2),
						uint32(1),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						types.PermissionWrite,
					).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(1))
				suite.k.SetNextPostID(ctx, 2, 1)
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			msg: types.NewMsgMovePost(
				1,
				1,
				2,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: true,
		},
		{
			name: "post moved properly",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(2)).
					Return(true)
				suite.sk.EXPECT().
					HasSection(gomock.Any(), uint64(2), uint32(1)).
					Return(true)
				suite.sk.EXPECT().
					HasPermission(
						gomock.Any(),
						uint64(2),
						uint32(1),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						types.PermissionWrite,
					).
					Return(true)
			},
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2021, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
				suite.k.SetNextPostID(ctx, 2, 2)
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))

				// Save a media
				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					2,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))

				// Save a tallied poll
				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					3,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						nil,
					),
				))

				// Save a tallied poll
				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					3,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2000, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
							types.NewAnswerResult(1, 100),
							types.NewAnswerResult(2, 50),
						}),
					),
				))

				// Save a active poll
				activePoll := types.NewAttachment(
					1,
					1,
					4,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						nil,
					),
				)
				suite.k.SaveAttachment(ctx, activePoll)
				suite.k.InsertActivePollQueue(ctx, activePoll)
			},
			msg: types.NewMsgMovePost(
				1,
				1,
				2,
				1,
				"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgMovePost{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd"),
				),
				sdk.NewEvent(
					types.EventTypeMovePost,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyNewSubspaceID, "2"),
					sdk.NewAttribute(types.AttributeKeyNewPostID, "2"),
				),
			},
			check: func(ctx sdk.Context) {
				// Check next id is updated
				nextID, err := suite.k.GetNextPostID(ctx, 2)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(3), nextID)

				// Check old post id is deleted
				suite.Require().False(suite.k.HasPost(ctx, 1, 1))

				// Check post is moved properly
				post, found := suite.k.GetPost(ctx, 2, 2)
				suite.Require().True(found)
				updateTime := ctx.BlockTime()
				suite.Require().Equal(types.NewPost(
					2,
					1,
					2,
					"External ID",
					"This is a text",
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					&updateTime,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				), post)

				// Check old attachments are deleted
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 1, 2))
				suite.Require().False(suite.k.HasAttachment(ctx, 1, 1, 4))

				// Check media active moved properly
				media, found := suite.k.GetAttachment(ctx, 2, 2, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewAttachment(
					2,
					2,
					2,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				), media)

				// Check tallied poll is moved properly
				talliedPoll, found := suite.k.GetAttachment(ctx, 2, 2, 3)
				suite.Require().True(found)
				suite.Require().Equal(types.NewAttachment(
					2,
					2,
					3,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2000, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
							types.NewAnswerResult(1, 100),
							types.NewAnswerResult(2, 50),
						}),
					),
				), talliedPoll)

				// Check active poll is moved properly
				activePoll, found := suite.k.GetAttachment(ctx, 2, 2, 4)
				suite.Require().True(found)
				suite.Require().Equal(types.NewAttachment(
					2,
					2,
					4,
					types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						nil,
					),
				), activePoll)

				// Check active poll is inside the queue
				suite.Require().True(ctx.KVStore(suite.storeKey).Has(
					types.ActivePollQueueKey(2, 2, 4, time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC))),
				)

				// Check next attachment id is set
				nextAttachmentID, err := suite.k.GetNextAttachmentID(ctx, 2, 2)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(5), nextAttachmentID)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			// Reset any event that might have been emitted during the setup
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.MovePost(sdk.WrapSDKContext(ctx), tc.msg)

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
