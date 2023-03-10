package keeper_test

import (
	"time"

	"github.com/golang/mock/gomock"

	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"
	"github.com/desmos-labs/desmos/v4/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

func (suite *KeeperTestSuite) TestMsgServer_AddReaction() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgAddReaction
		shouldErr   bool
		expResponse *types.MsgAddReactionResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "user without profile returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(false)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "non existing post returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.Post{}, false)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "blocked user returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						uint64(1)).
					Return(true)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						uint64(1)).
					Return(false)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionsReact).
					Return(false)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "already existing reaction returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						uint64(1)).
					Return(false)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionsReact).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "not set next reaction id returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						uint64(1)).
					Return(false)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionsReact).
					Return(true)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "invalid reaction value returns error",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionsReact).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReactionID(ctx, 1, 1, 1)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly - registered reaction value",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						uint64(1)).
					Return(false)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionsReact).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.DefaultReactionsParams(1))
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
				suite.k.SetNextReactionID(ctx, 1, 1, 1)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewRegisteredReactionValue(1),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: false,
			expResponse: &types.MsgAddReactionResponse{
				ReactionID: 1,
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAddReaction{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
				sdk.NewEvent(
					types.EventTypeAddReaction,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyReactionID, "1"),
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the reaction is saved
				stored, found := suite.k.GetReaction(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				), stored)

				// Make sure the next reaction id is updated properly
				storedID, err := suite.k.GetNextReactionID(ctx, 1, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), storedID)
			},
		},
		{
			name: "valid request works properly - free text reaction value",
			setup: func() {
				suite.ak.EXPECT().
					HasProfile(gomock.Any(), "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f").
					Return(true)

				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.rk.EXPECT().
					HasUserBlocked(gomock.Any(),
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						uint64(1)).
					Return(false)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionsReact).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.DefaultReactionsParams(1))
				suite.k.SetNextReactionID(ctx, 1, 1, 1)
			},
			msg: types.NewMsgAddReaction(
				1,
				1,
				types.NewFreeTextValue("🚀"),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: false,
			expResponse: &types.MsgAddReactionResponse{
				ReactionID: 1,
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAddReaction{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
				sdk.NewEvent(
					types.EventTypeAddReaction,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyReactionID, "1"),
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the reaction is saved
				stored, found := suite.k.GetReaction(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReaction(
					1,
					1,
					1,
					types.NewFreeTextValue("🚀"),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				), stored)

				// Make sure the next reaction id is updated properly
				storedID, err := suite.k.GetNextReactionID(ctx, 1, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), storedID)
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.AddReaction(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_RemoveReaction() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgRemoveReaction
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgRemoveReaction(
				1,
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "non existing post returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.Post{}, false)
			},
			msg: types.NewMsgRemoveReaction(
				1,
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "non existing reaction returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)
			},
			msg: types.NewMsgRemoveReaction(
				1,
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "different user and reaction author return error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)
			},
			msg: types.NewMsgRemoveReaction(
				1,
				1,
				1,
				"cosmos1gh5af0v83nvxupch3zsg3cktqu6wwgye3cm5ra",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionsReact).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			msg: types.NewMsgRemoveReaction(
				1,
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.pk.EXPECT().
					GetPost(gomock.Any(), uint64(1), uint64(1)).
					Return(poststypes.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					), true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionsReact).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
				))
			},
			msg: types.NewMsgRemoveReaction(
				1,
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRemoveReaction{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
				sdk.NewEvent(
					types.EventTypeRemoveReaction,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyPostID, "1"),
					sdk.NewAttribute(types.AttributeKeyReactionID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the reaction has been deleted
				suite.Require().False(suite.k.HasReaction(ctx, 1, 1, 1))
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.RemoveReaction(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_AddRegisteredReaction() {
	testCases := []struct {
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		msg         *types.MsgAddRegisteredReaction
		shouldErr   bool
		expResponse *types.MsgAddRegisteredReactionResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgAddRegisteredReaction(
				1,
				":hello:",
				"https://example.com?image=hello.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(false)
			},
			msg: types.NewMsgAddRegisteredReaction(
				1,
				":hello:",
				"https://example.com?image=hello.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "not set next registered reaction id returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(true)
			},
			msg: types.NewMsgAddRegisteredReaction(
				1,
				":hello:",
				"https://example.com?image=hello.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "invalid registered reaction data returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextRegisteredReactionID(ctx, 1, 1)
			},
			msg: types.NewMsgAddRegisteredReaction(
				1,
				"",
				"https://example.com?image=hello.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextRegisteredReactionID(ctx, 1, 1)
			},
			msg: types.NewMsgAddRegisteredReaction(
				1,
				":hello:",
				"https://example.com?image=hello.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: false,
			expResponse: &types.MsgAddRegisteredReactionResponse{
				RegisteredReactionID: 1,
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAddRegisteredReaction{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
				sdk.NewEvent(
					types.EventTypeAddRegisteredReaction,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyRegisteredReactionID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the registered reaction is saved properly
				stored, found := suite.k.GetRegisteredReaction(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				), stored)

				// Make sure the next registered reaction id is updated properly
				storedID, err := suite.k.GetNextRegisteredReactionID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), storedID)
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.AddRegisteredReaction(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_EditRegisteredReaction() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgEditRegisteredReaction
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgEditRegisteredReaction(
				1,
				1,
				":wave:",
				"https://example.com?image=wave.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "non existing registered reaction returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			msg: types.NewMsgEditRegisteredReaction(
				1,
				1,
				":wave:",
				"https://example.com?image=wave.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			msg: types.NewMsgEditRegisteredReaction(
				1,
				1,
				":wave:",
				"https://example.com?image=wave.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "invalid update returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			msg: types.NewMsgEditRegisteredReaction(
				1,
				1,
				"",
				"",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			msg: types.NewMsgEditRegisteredReaction(
				1,
				1,
				":wave:",
				"https://example.com?image=wave.png",
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgEditRegisteredReaction{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
				sdk.NewEvent(
					types.ActionEditRegisteredReaction,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyRegisteredReactionID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the reaction has been edited
				stored, found := suite.k.GetRegisteredReaction(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewRegisteredReaction(
					1,
					1,
					":wave:",
					"https://example.com?image=wave.png",
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.EditRegisteredReaction(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_RemoveRegisteredReaction() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgRemoveRegisteredReaction
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgRemoveRegisteredReaction(
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "non existing registered reaction returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)
			},
			msg: types.NewMsgRemoveRegisteredReaction(
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			msg: types.NewMsgRemoveRegisteredReaction(
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageRegisteredReactions).
					Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			msg: types.NewMsgRemoveRegisteredReaction(
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRemoveRegisteredReaction{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
				sdk.NewEvent(
					types.EventTypeRemoveRegisteredReaction,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyRegisteredReactionID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the reaction has been deleted
				suite.Require().False(suite.k.HasRegisteredReaction(ctx, 1, 1))
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.RemoveRegisteredReaction(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_SetReactionsParams() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgSetReactionsParams
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(false)
			},
			msg: types.NewMsgSetReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, "[a-z]"),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageReactionParams).
					Return(false)
			},
			msg: types.NewMsgSetReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, "[a-z]"),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "invalid params return error",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageReactionParams).
					Return(true)
			},
			msg: types.NewMsgSetReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 0, ""),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "valid request works properly",
			setup: func() {
				suite.sk.EXPECT().
					HasSubspace(gomock.Any(), uint64(1)).
					Return(true)

				suite.sk.EXPECT().
					HasPermission(gomock.Any(),
						uint64(1),
						uint32(0),
						"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
						types.PermissionManageReactionParams).
					Return(true)
			},
			msg: types.NewMsgSetReactionsParams(
				1,
				types.NewRegisteredReactionValueParams(true),
				types.NewFreeTextValueParams(true, 100, "[a-z]"),
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgSetReactionsParams{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"),
				),
				sdk.NewEvent(
					types.EventTypeSetReactionsParams,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the params have been stored
				stored, err := suite.k.GetSubspaceReactionsParams(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 100, "[a-z]"),
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
			if tc.store != nil {
				tc.store(ctx)
			}

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.SetReactionsParams(sdk.WrapSDKContext(ctx), tc.msg)
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
