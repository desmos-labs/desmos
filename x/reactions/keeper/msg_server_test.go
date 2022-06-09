package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/testutil"
	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	"github.com/desmos-labs/desmos/v3/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v3/x/reactions/types"
	relationshipstypes "github.com/desmos-labs/desmos/v3/x/relationships/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestMsgServer_AddReaction() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		msg         *types.MsgAddReaction
		shouldErr   bool
		expResponse *types.MsgAddReactionResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "user without profile returns error",
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
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)
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
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)

				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
			name: "blocked user returns error",
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)

				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.rk.SaveUserBlock(ctx, relationshipstypes.NewUserBlock(
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					"",
					1,
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
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)

				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
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
			name: "already existing reaction returns error",
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)

				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionsReact),
				)

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
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)

				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionsReact),
				)
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
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)

				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionsReact),
				)

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
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)

				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionsReact),
				)

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
			store: func(ctx sdk.Context) {
				err := suite.ak.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f"))
				suite.Require().NoError(err)

				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionsReact),
				)

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
		store     func(ctx sdk.Context)
		msg       *types.MsgRemoveReaction
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
			name: "non existing reaction returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionsReact),
				)
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionsReact),
				)
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionsReact),
				)

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
		store       func(ctx sdk.Context)
		msg         *types.MsgAddRegisteredReaction
		shouldErr   bool
		expResponse *types.MsgAddRegisteredReactionResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions),
				)
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions),
				)

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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions),
				)

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
		store     func(ctx sdk.Context)
		msg       *types.MsgEditRegisteredReaction
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions),
				)

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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions),
				)

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
		store     func(ctx sdk.Context)
		msg       *types.MsgRemoveRegisteredReaction
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgRemoveRegisteredReaction(
				1,
				1,
				"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
			),
			shouldErr: true,
		},
		{
			name: "non existing registered reaction returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionManageRegisteredReactions),
				)

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
					types.ActionRemoveRegisteredReaction,
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
		store     func(ctx sdk.Context)
		msg       *types.MsgSetReactionsParams
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionManageReactionParams),
				)
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
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.sk.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1efa8l9h4p6hmkps6vk8lu7nxydr46npr8qtg5f",
					subspacestypes.NewPermissions(types.PermissionManageReactionParams),
				)
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
					types.ActionSetReactionParams,
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
