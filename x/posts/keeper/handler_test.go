package keeper_test

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) Test_handleMsgCreatePost() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")

	computedID := types.ComputeID(suite.testData.post.Created, suite.testData.post.Creator, suite.testData.post.Subspace)

	otherCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	tests := []struct {
		name        string
		storedPosts types.Posts
		msg         types.MsgCreatePost
		expPost     types.Post
		expError    error
	}{
		{
			name: "Trying to store post with same id returns expError",
			storedPosts: types.Posts{
				types.NewPost(
					computedID,
					suite.testData.post.ParentID,
					suite.testData.post.Message,
					suite.testData.post.AllowsComments,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					suite.testData.post.Created,
					suite.testData.post.Creator,
				),
			},
			msg: types.NewMsgCreatePost(
				suite.testData.post.Message,
				suite.testData.post.ParentID,
				suite.testData.post.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"the provided post conflicts with the one having id 46e61c7ac7016e8dd1d7270b114ecb7d1cf45cc85caa0308de540ccc15676fc7"),
		},
		{
			name: "Post with new id is stored properly",
			msg: types.NewMsgCreatePost(
				suite.testData.post.Message,
				suite.testData.post.ParentID,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expPost: types.NewPost(
				computedID,
				suite.testData.post.ParentID,
				suite.testData.post.Message,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.testData.post.Created,
				suite.testData.post.Creator,
			).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
		},
		{
			name: "Storing a valid post with missing parent id returns expError",
			msg: types.NewMsgCreatePost(
				suite.testData.post.Message,
				id2,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "parent post with id f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd not found"),
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns expError",
			storedPosts: types.Posts{
				types.NewPost(
					id,
					id2,
					"Parent post",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					suite.testData.post.Created,
					suite.testData.post.Creator,
				),
			},
			msg: types.NewMsgCreatePost(
				suite.testData.post.Message,
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				otherCreator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af does not allow comments"),
		},
		{
			name: "Post with exact same data is not posted again",
			storedPosts: []types.Post{
				types.NewPost(
					computedID,
					suite.testData.post.ParentID,
					suite.testData.post.Message,
					suite.testData.post.AllowsComments,
					suite.testData.post.Subspace,
					suite.testData.post.OptionalData,
					suite.testData.post.Created,
					suite.testData.post.Creator,
				).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
			},
			msg: types.NewMsgCreatePost(
				suite.testData.post.Message,
				suite.testData.post.ParentID,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"the provided post conflicts with the one having id 46e61c7ac7016e8dd1d7270b114ecb7d1cf45cc85caa0308de540ccc15676fc7"),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())
			store := suite.ctx.KVStore(suite.keeper.StoreKey)

			for _, p := range test.storedPosts {
				computedID := types.ComputeID(suite.ctx.BlockTime(), p.Creator, p.Subspace)
				test.msg.ParentID = computedID
				store.Set(types.PostStoreKey(computedID), suite.keeper.Cdc.MustMarshalBinaryBare(p))
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				// Check the post
				var stored types.Post
				computedID := types.ComputeID(suite.ctx.BlockTime(), test.expPost.Creator, test.expPost.Subspace)
				suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(computedID)), &stored)
				test.expPost.Created = suite.ctx.BlockTime() // make sure that the two posts has the same BlockTime
				test.expPost.PostID = computedID             // make sure that the two posts has the same ID calculated using the blockTime

				suite.True(stored.Equals(test.expPost), "Expected: %s, actual: %s", test.expPost, stored)

				// Check the data
				suite.Equal(suite.keeper.Cdc.MustMarshalBinaryLengthPrefixed(test.expPost.PostID), res.Data)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypePostCreated,
					sdk.NewAttribute(types.AttributeKeyPostID, test.expPost.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostParentID, test.expPost.ParentID.String()),
					sdk.NewAttribute(types.AttributeKeyPostCreationTime, test.expPost.Created.Format(time.RFC3339)),
					sdk.NewAttribute(types.AttributeKeyPostOwner, test.expPost.Creator.String()),
				)
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, creationEvent)
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expError.Error(), err.Error())
			}
		})
	}

}

func (suite *KeeperTestSuite) Test_handleMsgEditPost() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	editor, err := sdk.AccAddressFromBech32("cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63")
	suite.NoError(err)

	testData := []struct {
		name       string
		storedPost *types.Post
		msg        types.MsgEditPost
		expError   error
		expPost    *types.Post
	}{
		{
			name:       "Post not found",
			storedPost: nil,
			msg:        types.NewMsgEditPost(id, "Edited message", suite.testData.post.Creator),
			expError:   sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af not found"),
		},
		{
			name:       "Invalid editor",
			storedPost: &suite.testData.post,
			msg:        types.NewMsgEditPost(suite.testData.post.PostID, "Edited message", editor),
			expError:   sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner"),
		},
		{
			name:       "Edit date before creation date",
			storedPost: &suite.testData.post,
			msg:        types.NewMsgEditPost(suite.testData.post.PostID, "Edited message", suite.testData.post.Creator),
			expError:   sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "edit date cannot be before creation date"),
		},
		{
			name:       "Valid request is handled properly",
			storedPost: &suite.testData.post,
			msg:        types.NewMsgEditPost(suite.testData.post.PostID, "Edited message", suite.testData.post.Creator),
			expPost: &types.Post{
				PostID:         suite.testData.post.PostID,
				ParentID:       suite.testData.post.ParentID,
				Message:        "Edited message",
				Created:        suite.testData.post.Created,
				LastEdited:     suite.testData.post.Created.AddDate(0, 0, 1),
				AllowsComments: suite.testData.post.AllowsComments,
				Subspace:       suite.testData.post.Subspace,
				OptionalData:   suite.testData.post.OptionalData,
				Creator:        suite.testData.post.Creator,
				Attachments:    suite.testData.post.Attachments,
				PollData:       suite.testData.post.PollData,
			},
		},
	}

	for _, test := range testData {
		test := test
		suite.Run(test.name, func() {
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())

			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedPost != nil {
				computedID := types.ComputeID(suite.ctx.BlockTime(), test.storedPost.Creator, test.storedPost.Subspace)
				test.storedPost.PostID = computedID
				test.msg.PostID = computedID
				if test.expPost != nil {
					test.storedPost.Created = suite.ctx.BlockTime()
					test.expPost.Created = test.storedPost.Created
					test.expPost.PostID = test.storedPost.PostID
					suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().AddDate(0, 0, 1))
					test.expPost.LastEdited = suite.ctx.BlockTime()
				}
				store.Set(types.PostStoreKey(computedID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedPost))
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				suite.Contains(res.Events, sdk.NewEvent(
					types.EventTypePostEdited,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostEditTime, test.expPost.LastEdited.Format(time.RFC3339)),
				))

				var stored types.Post
				suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.storedPost.PostID)), &stored)
				suite.True(test.expPost.Equals(stored))
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expError.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAddPostReaction() {
	post := types.NewPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"",
		"Post message",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		suite.testData.post.Created,
		suite.testData.post.Creator,
	)

	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	suite.NoError(err)

	tests := []struct {
		name               string
		existingPost       *types.Post
		msg                types.MsgAddPostReaction
		registeredReaction *types.Reaction
		error              error
		expEvent           sdk.Event
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgAddPostReaction("invalid", ":smile:", user),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id invalid not found"),
		},
		{
			name:         "Registered Reaction not found",
			existingPost: &post,
			msg:          types.NewMsgAddPostReaction(post.PostID, ":super-smile:", user),
			error:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "short code :super-smile: must be registered before using it"),
		},
		{
			name:               "Valid message works properly (shortcode)",
			existingPost:       &post,
			msg:                types.NewMsgAddPostReaction(post.PostID, ":smile:", user),
			registeredReaction: &suite.testData.registeredReaction,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionAdded,
				sdk.NewAttribute(types.AttributeKeyPostID, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, "😄"),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":smile:"),
			),
		},
		{
			name:         "Valid message works properly (emoji)",
			existingPost: &post,
			msg:          types.NewMsgAddPostReaction(post.PostID, "🙂", user),
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionAdded,
				sdk.NewAttribute(types.AttributeKeyPostID, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, "🙂"),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":slightly_smiling_face:"),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.existingPost))
			}

			if test.registeredReaction != nil {
				suite.keeper.RegisterReaction(suite.ctx, *test.registeredReaction)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				suite.Contains(res.Events, test.expEvent)

				// Check the post
				var storedPost types.Post
				suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.msg.PostID)), &storedPost)
				suite.True(test.existingPost.Equals(storedPost))

				// Check the post reactions
				var reactValue, reactShortcode string
				if e, err := emoji.LookupEmoji(test.msg.Reaction); err == nil {
					reactShortcode = e.Shortcodes[0]
					reactValue = e.Value
				} else {
					e, err := emoji.LookupEmojiByCode(test.msg.Reaction)
					if err != nil {
						panic(err)
					}
					reactShortcode = e.Shortcodes[0]
					reactValue = e.Value
				}

				var storedReactions types.PostReactions
				suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				suite.Contains(storedReactions, types.NewPostReaction(reactShortcode, reactValue, test.msg.User))

				// Check the registered reactions
				registeredReactions := suite.keeper.GetRegisteredReactions(suite.ctx)
				if test.registeredReaction != nil {
					found := false
					for _, reaction := range registeredReactions {
						found = found || reaction.Equals(*test.registeredReaction)
					}
					suite.True(found)
				} else {
					suite.Empty(registeredReactions)
				}
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.error.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRemovePostReaction() {
	post := types.NewPost(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"",
		"Post message",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		suite.testData.post.Created,
		suite.testData.post.Creator,
	)

	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	suite.NoError(err)

	regReaction := types.NewReaction(user, ":reaction:", "react", suite.testData.post.Subspace)
	reaction := types.NewPostReaction(":reaction:", "react", user)
	emojiShortcodeReaction := types.NewPostReaction(":smile:", "😄", user)

	emoji, err := emoji.LookupEmojiByCode(":+1:")
	suite.NoError(err)

	emojiReaction := types.NewPostReaction(emoji.Shortcodes[0], emoji.Value, user)

	tests := []struct {
		name               string
		existingPost       *types.Post
		registeredReaction *types.Reaction
		existingReaction   *types.PostReaction
		msg                types.MsgRemovePostReaction
		error              error
		expEvent           sdk.Event
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgRemovePostReaction("invalid", user, "reaction"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id invalid not found"),
		},
		{
			name:         "Reaction not found",
			existingPost: &post,
			msg:          types.NewMsgRemovePostReaction(post.PostID, user, "😄"),
			error:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("cannot remove the reaction with value :smile: from user %s as it does not exist", user)),
		},
		{
			name:               "Removing a reaction using the code works properly (registered reaction)",
			existingPost:       &post,
			existingReaction:   &reaction,
			registeredReaction: &regReaction,
			msg:                types.NewMsgRemovePostReaction(post.PostID, user, reaction.Shortcode),
			error:              nil,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionRemoved,
				sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, user.String()),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, regReaction.Value),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, regReaction.ShortCode),
			),
		},
		{
			name:             "Removing a reaction using the code works properly (emoji shortcode)",
			existingPost:     &post,
			existingReaction: &emojiShortcodeReaction,
			msg:              types.NewMsgRemovePostReaction(post.PostID, user, emojiShortcodeReaction.Shortcode),
			error:            nil,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionRemoved,
				sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, user.String()),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, "😄"),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, emojiShortcodeReaction.Shortcode),
			),
		},
		{
			name:             "Removing a reaction using the emoji works properly",
			existingPost:     &post,
			existingReaction: &emojiReaction,
			msg:              types.NewMsgRemovePostReaction(post.PostID, user, emoji.Value),
			error:            nil,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionRemoved,
				sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, user.String()),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, emoji.Value),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, emojiReaction.Shortcode),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.existingPost))
			}

			if test.registeredReaction != nil {
				store.Set(types.ReactionsStoreKey(test.registeredReaction.ShortCode, test.registeredReaction.Subspace),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.registeredReaction))
			}

			if test.existingReaction != nil {
				store.Set(
					types.PostReactionsStoreKey(test.existingPost.PostID),
					suite.keeper.Cdc.MustMarshalBinaryBare(&types.PostReactions{*test.existingReaction}),
				)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				suite.Contains(res.Events, test.expEvent)

				var storedPost types.Post
				suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(suite.testData.post.PostID)), &storedPost)
				suite.True(test.existingPost.Equals(storedPost))

				var storedReactions types.PostReactions
				suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				suite.NotContains(storedReactions, test.existingReaction)
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.error.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAnswerPollPost() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}
	userPollAnswers := types.NewUserAnswer(answers, suite.testData.post.Creator)

	tests := []struct {
		name          string
		msg           types.MsgAnswerPoll
		storedPost    types.Post
		storedAnswers *types.UserAnswer
		expErr        error
	}{
		{
			name: "Post not found",
			msg:  types.NewMsgAnswerPoll(id2, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				suite.testData.post.Created,
				suite.testData.post.Creator,
			).WithPollData(types.NewPollData(
				"poll?",
				suite.testData.postEndPollDate,
				types.PollAnswers{suite.testData.answers[0], suite.testData.answers[1]},
				true,
				true,
				true,
			)),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd doesn't exist"),
		},
		{
			name: "No poll associated with post",
			msg:  types.NewMsgAnswerPoll(id2, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.NewPost(
				id2,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				suite.testData.post.Created,
				suite.testData.post.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no poll associated with ID: f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
		},
		{
			name: "Answer after poll closure",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1}, suite.testData.post.Creator),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				suite.testData.post.Created,
				suite.testData.post.Creator,
			).WithPollData(types.NewPollData(
				"poll?",
				suite.testData.postEndPollDate,
				types.PollAnswers{suite.testData.answers[0]},
				true,
				false,
				true,
			)),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("the poll associated with ID %s was closed at %s", id, suite.testData.postEndPollDateExpired)),
		},
		{
			name: "Poll doesn't allow multiple answers",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				suite.testData.postCreationDate,
				suite.testData.post.Creator,
			).WithPollData(types.NewPollData(
				"poll?",
				suite.testData.postEndPollDate,
				types.PollAnswers{suite.testData.answers[0]},
				true,
				false,
				true,
			),
			),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "the poll associated with ID 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af doesn't allow multiple answers"),
		},
		{
			name: "Creator provide too many answers",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2, 3}, suite.testData.post.Creator),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				suite.testData.post.Created,
				suite.testData.post.Creator,
			).WithPollData(types.NewPollData(
				"poll?",
				suite.testData.postEndPollDate,
				suite.testData.answers,
				true,
				true,
				true,
			),
			),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "user's answers are more than the available ones in Poll"),
		},
		{
			name: "Creator provide answers that are not the ones provided by the poll",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 3}, suite.testData.post.Creator),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				suite.testData.post.Created,
				suite.testData.post.Creator,
			).WithPollData(types.NewPollData(
				"poll?",
				suite.testData.postEndPollDate,
				suite.testData.answers,
				true,
				true,
				true,
			)),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "answer with ID 3 isn't one of the poll's provided answers"),
		},
		{
			name: "Poll doesn't allow answers' edits",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				suite.testData.post.Created,
				suite.testData.post.Creator,
			).WithPollData(types.NewPollData(
				"poll?",
				suite.testData.postEndPollDate,
				suite.testData.answers,
				true,
				true,
				false,
			)),
			storedAnswers: &userPollAnswers,
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "post with ID 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af doesn't allow answers' edits"),
		},
		{
			name: "Answered correctly to post's poll",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				suite.testData.post.Created,
				suite.testData.post.Creator,
			).WithPollData(types.NewPollData(
				"poll?",
				suite.testData.postEndPollDate,
				suite.testData.answers,
				true,
				true,
				true,
			)),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			store.Set(types.PostStoreKey(test.storedPost.PostID), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedPost))

			if test.storedAnswers != nil {
				suite.keeper.SavePollAnswers(suite.ctx, test.storedPost.PostID, *test.storedAnswers)
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}

			// Valid response
			if res != nil {
				{
					// Check the data
					suite.Equal(suite.keeper.Cdc.MustMarshalBinaryLengthPrefixed("Answered to poll correctly"), res.Data)

					// Check the events
					answerEvent := sdk.NewEvent(
						types.EventTypeAnsweredPoll,
						sdk.NewAttribute(types.AttributeKeyPostID, test.storedPost.PostID.String()),
						sdk.NewAttribute(types.AttributeKeyPollAnswerer, suite.testData.post.Creator.String()),
					)

					suite.Len(res.Events, 1)
					suite.Contains(res.Events, answerEvent)
				}
			}
		})
	}

}

func (suite *KeeperTestSuite) Test_handleMsgRegisterReaction() {
	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	suite.NoError(err)

	tests := []struct {
		name              string
		existingReactions []types.Reaction
		msg               types.MsgRegisterReaction
		error             error
	}{
		{
			name: "Reaction registered without error",
			msg: types.NewMsgRegisterReaction(
				user,
				":test:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
		{
			name: "Emoji reaction returns error",
			msg: types.NewMsgRegisterReaction(
				user,
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				`shortcode :smile: represents an emoji and thus can't be used to register a new reaction`,
			),
		},
		{
			name: "Already registered reaction returns error",
			existingReactions: []types.Reaction{
				types.NewReaction(
					user,
					":test:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types.NewMsgRegisterReaction(
				user,
				":test:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"reaction with shortcode :test: and subspace 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e has already been registered",
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			for _, react := range test.existingReactions {
				react := react
				store.Set(types.ReactionsStoreKey(react.ShortCode, react.Subspace), suite.keeper.Cdc.MustMarshalBinaryBare(&react))
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				suite.Contains(res.Events, sdk.NewEvent(
					types.EventTypeRegisterReaction,
					sdk.NewAttribute(types.AttributeKeyReactionCreator, test.msg.Creator.String()),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, test.msg.ShortCode),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, test.msg.Value),
					sdk.NewAttribute(types.AttributeKeyReactionSubSpace, test.msg.Subspace),
				))

				var storedReaction types.Reaction
				suite.keeper.Cdc.MustUnmarshalBinaryBare(
					store.Get(types.ReactionsStoreKey(test.msg.ShortCode, test.msg.Subspace)),
					&storedReaction,
				)

				expected := types.NewReaction(user, test.msg.ShortCode, test.msg.Value, test.msg.Subspace)
				suite.True(expected.Equals(storedReaction))
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Equal(test.error.Error(), err.Error())
			}
		})
	}
}
