package keeper_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/posts"
	"strings"
	"time"

	"github.com/desmos-labs/desmos/x/relationships"

	"github.com/desmos-labs/desmos/x/posts/types/models"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) Test_handleMsgCreatePost() {
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")

	otherCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	createPostMessage := types.NewMsgCreatePost(
		suite.testData.post.Message,
		suite.testData.post.ParentID,
		suite.testData.post.AllowsComments,
		suite.testData.post.Subspace,
		suite.testData.post.OptionalData,
		suite.testData.post.Creator,
		suite.testData.post.Attachments,
		suite.testData.post.PollData,
	)

	postID := types.PostID("040b0c16cd541101d24100e4a9c90e4dbaebbee977a94d673f79591cbb5f4465")

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
				types.Post{
					PostID:         postID,
					ParentID:       suite.testData.post.ParentID,
					Message:        suite.testData.post.Message,
					Created:        suite.testData.post.Created,
					AllowsComments: suite.testData.post.AllowsComments,
					Subspace:       suite.testData.post.Subspace,
					OptionalData:   suite.testData.post.OptionalData,
					Creator:        suite.testData.post.Creator,
				},
			},
			msg: createPostMessage,
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"the provided post conflicts with the one having id 040b0c16cd541101d24100e4a9c90e4dbaebbee977a94d673f79591cbb5f4465"),
		},
		{
			name: "Post with new id is stored properly",
			msg:  createPostMessage,
			expPost: types.NewPost(
				suite.testData.post.ParentID,
				suite.testData.post.Message,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.ctx.BlockTime(),
				suite.testData.post.Creator,
			).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
		},
		{
			name: "Storing a valid post with missing parent id returns expError",
			msg: types.NewMsgCreatePost(
				suite.testData.post.Message,
				id2,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("parent post with id %s not found", id2)),
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns expError",
			storedPosts: types.Posts{
				types.Post{
					PostID:         id2,
					ParentID:       "1234",
					Message:        "Parent post",
					Created:        suite.testData.post.Created,
					AllowsComments: false,
					Subspace:       suite.testData.post.Subspace,
					OptionalData:   suite.testData.post.OptionalData,
					Creator:        suite.testData.post.Creator,
				},
			},
			msg: types.NewMsgCreatePost(
				suite.testData.post.Message,
				id2,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				otherCreator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd does not allow comments"),
		},
		{
			name: "Post with exact same data is not posted again",
			storedPosts: []types.Post{
				types.NewPost(
					suite.testData.post.ParentID,
					suite.testData.post.Message,
					suite.testData.post.AllowsComments,
					suite.testData.post.Subspace,
					suite.testData.post.OptionalData,
					suite.ctx.BlockTime(),
					suite.testData.post.Creator,
				).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
			},
			msg: createPostMessage,
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"the provided post conflicts with the one having id 040b0c16cd541101d24100e4a9c90e4dbaebbee977a94d673f79591cbb5f4465"),
		},
		{
			name: "Post message cannot be longer than 500 characters",
			msg: types.NewMsgCreatePost(
				strings.Repeat("a", 550),
				suite.testData.post.ParentID,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"post with id 38caeb754684d0173f3e47e45831bd15a23056caa9b64b498a61b67739f6f8a0 has more than 500 characters"),
		},
		{
			name: "post tag blocked the post creator",
			msg: types.NewMsgCreatePost(
				"blocked",
				suite.testData.post.ParentID,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.testData.post.Creator,
				[]types.Attachment{types.NewAttachment("http://uri.com", "text/plain",
					[]sdk.AccAddress{otherCreator})},
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("The user with address %s has blocked you", otherCreator)),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())
			store := suite.ctx.KVStore(suite.keeper.storeKey)

			for _, p := range test.storedPosts {
				store.Set(types.PostStoreKey(p.PostID), suite.keeper.cdc.MustMarshalBinaryBare(p))
			}

			if test.msg.Message == "blocked" {
				_ = suite.relationshipsKeeper.SaveUserBlock(suite.ctx,
					relationships.NewUserBlock(otherCreator, suite.testData.post.Creator, "test", ""))
			}

			handler := posts.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				// Check the post
				var stored types.Post
				suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(postID)), &stored)

				suite.True(stored.Equals(test.expPost), "Expected: %s, actual: %s", test.expPost, stored)

				// Check the data
				suite.Require().Equal(suite.keeper.cdc.MustMarshalBinaryLengthPrefixed(test.expPost.PostID), res.Data)

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
				suite.Require().Equal(test.expError.Error(), err.Error())
			}
		})
	}

}

func (suite *KeeperTestSuite) Test_handleMsgEditPost() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	editor, err := sdk.AccAddressFromBech32("cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63")
	suite.Require().NoError(err)
	timeZone, _ := time.LoadLocation("UTC")

	editedPollData := types.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
		models.NewPollAnswers(
			types.NewPollAnswer(models.AnswerID(1), "No"),
			types.NewPollAnswer(models.AnswerID(2), "No"),
		),
		false,
		true,
	)

	editedAttachments := models.NewAttachments(types.NewAttachment("https://edited.com", "text/plain", nil))

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
			msg:        types.NewMsgEditPost(id, "Edited message", nil, nil, suite.testData.post.Creator),
			expError:   sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af not found"),
		},
		{
			name:       "Invalid editor",
			storedPost: &suite.testData.post,
			msg:        types.NewMsgEditPost(suite.testData.post.PostID, "Edited message", nil, nil, editor),
			expError:   sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner"),
		},
		{
			name:       "Edit date before creation date",
			storedPost: &suite.testData.post,
			msg:        types.NewMsgEditPost(suite.testData.post.PostID, "Edited message", nil, nil, suite.testData.post.Creator),
			expError:   sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "edit date cannot be before creation date"),
		},
		{
			name:       "Blocked creator from tags",
			storedPost: &suite.testData.post,
			msg: types.NewMsgEditPost(suite.testData.post.PostID, "blocked",
				models.NewAttachments(types.NewAttachment("https://edited.com", "text/plain",
					[]sdk.AccAddress{otherCreator})), nil, suite.testData.post.Creator),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf(fmt.Sprintf("The user with address %s has blocked you", otherCreator))),
		},
		{
			name:       "Valid request is handled properly without attachments and pollData",
			storedPost: &suite.testData.post,
			msg: types.NewMsgEditPost(suite.testData.post.PostID, "Edited message",
				editedAttachments, &editedPollData, suite.testData.post.Creator),
			expPost: &types.Post{
				PostID:         suite.testData.post.PostID,
				ParentID:       suite.testData.post.ParentID,
				Message:        "Edited message",
				Created:        suite.ctx.BlockTime(),
				LastEdited:     suite.testData.post.Created.AddDate(0, 0, 1),
				AllowsComments: suite.testData.post.AllowsComments,
				Subspace:       suite.testData.post.Subspace,
				OptionalData:   suite.testData.post.OptionalData,
				Creator:        suite.testData.post.Creator,
				Attachments:    editedAttachments,
				PollData:       &editedPollData,
			},
		},
	}

	for _, test := range testData {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())

			store := suite.ctx.KVStore(suite.keeper.storeKey)

			if test.msg.Message == "blocked" {
				_ = suite.relationshipsKeeper.SaveUserBlock(suite.ctx,
					relationships.NewUserBlock(otherCreator, suite.testData.post.Creator, "test", ""))
				test.storedPost.Created = suite.ctx.BlockTime()
				suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().AddDate(0, 0, 1))
			}

			if test.storedPost != nil {
				if test.expPost != nil {
					test.storedPost.Created = suite.ctx.BlockTime()
					suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().AddDate(0, 0, 1))
					test.expPost.LastEdited = suite.ctx.BlockTime()
				}
				store.Set(types.PostStoreKey(test.storedPost.PostID), suite.keeper.cdc.MustMarshalBinaryBare(&test.storedPost))
			}

			handler := posts.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				suite.Contains(res.Events, sdk.NewEvent(
					types.EventTypePostEdited,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostEditTime, test.expPost.LastEdited.Format(time.RFC3339)),
				))

				var stored types.Post
				suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.storedPost.PostID)), &stored)
				suite.True(test.expPost.Equals(stored))
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expError.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAddPostReaction() {
	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	suite.Require().NoError(err)

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
			existingPost: &suite.testData.post,
			msg:          types.NewMsgAddPostReaction(suite.testData.post.PostID, ":super-smile:", user),
			error:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "short code :super-smile: must be registered before using it"),
		},
		{
			name:               "Valid message works properly (shortcode)",
			existingPost:       &suite.testData.post,
			msg:                types.NewMsgAddPostReaction(suite.testData.post.PostID, ":smile:", user),
			registeredReaction: &suite.testData.registeredReaction,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionAdded,
				sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID.String()),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, "😄"),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":smile:"),
			),
		},
		{
			name:         "Valid message works properly (emoji)",
			existingPost: &suite.testData.post,
			msg:          types.NewMsgAddPostReaction(suite.testData.post.PostID, "🙂", user),
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionAdded,
				sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID.String()),
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
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), suite.keeper.cdc.MustMarshalBinaryBare(&test.existingPost))
			}

			if test.registeredReaction != nil {
				suite.keeper.SaveRegisteredReaction(suite.ctx, *test.registeredReaction)
			}

			handler := posts.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				suite.Contains(res.Events, test.expEvent)

				// Check the post
				var storedPost types.Post
				suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.msg.PostID)), &storedPost)
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
				suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
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
				suite.Require().Equal(test.error.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgRemovePostReaction() {
	post := types.Post{
		PostID:       suite.testData.postID,
		Message:      "Post message",
		Created:      suite.testData.post.Created,
		Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData: nil,
		Creator:      suite.testData.post.Creator,
	}

	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	suite.Require().NoError(err)

	regReaction := types.NewReaction(user, ":reaction:", "react", suite.testData.post.Subspace)
	reaction := types.NewPostReaction(":reaction:", "react", user)
	emojiShortcodeReaction := types.NewPostReaction(":smile:", "😄", user)

	emoji, err := emoji.LookupEmojiByCode(":+1:")
	suite.Require().NoError(err)

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
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), suite.keeper.cdc.MustMarshalBinaryBare(&test.existingPost))
			}

			if test.registeredReaction != nil {
				store.Set(types.ReactionsStoreKey(test.registeredReaction.ShortCode, test.registeredReaction.Subspace),
					suite.keeper.cdc.MustMarshalBinaryBare(&test.registeredReaction))
			}

			if test.existingReaction != nil {
				store.Set(
					types.PostReactionsStoreKey(test.existingPost.PostID),
					suite.keeper.cdc.MustMarshalBinaryBare(&types.PostReactions{*test.existingReaction}),
				)
			}

			handler := posts.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				suite.Contains(res.Events, test.expEvent)

				var storedPost types.Post
				suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(suite.testData.post.PostID)), &storedPost)
				suite.True(test.existingPost.Equals(storedPost))

				var storedReactions types.PostReactions
				suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				suite.NotContains(storedReactions, test.existingReaction)
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.error.Error(), err.Error())
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
				"",
				"Post message",
				false,
				"desmos",
				nil,
				suite.testData.post.Created,
				suite.testData.post.Creator,
			).WithPollData(types.NewPollData(
				"poll?",
				suite.testData.postEndPollDate,
				types.PollAnswers{suite.testData.answers[0], suite.testData.answers[1]},
				true,
				true,
			)),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd doesn't exist"),
		},
		{
			name: "No poll associated with post",
			msg:  types.NewMsgAnswerPoll(id2, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.Post{
				PostID:       id2,
				Message:      "Post message",
				Created:      suite.testData.post.Created,
				Subspace:     suite.testData.post.Subspace,
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
			},
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no poll associated with ID: f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
		},
		{
			name: "Answer after poll closure",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1}, suite.testData.post.Creator),
			storedPost: types.Post{
				PostID:       id,
				Message:      "Post message",
				Created:      suite.testData.post.Created,
				Subspace:     suite.testData.post.Subspace,
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
				PollData: &types.PollData{
					Question:          "poll?",
					ProvidedAnswers:   types.PollAnswers{suite.testData.answers[0]},
					EndDate:           suite.testData.postEndPollDateExpired,
					AllowsAnswerEdits: true,
				},
			},
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("the poll associated with ID %s was closed at %s", id, suite.testData.postEndPollDateExpired)),
		},
		{
			name: "Poll doesn't allow multiple answers",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.Post{
				PostID:       id,
				Message:      "Post message",
				Created:      suite.testData.post.Created,
				Subspace:     suite.testData.post.Subspace,
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
				PollData: &types.PollData{
					Question:              "poll?",
					ProvidedAnswers:       types.PollAnswers{suite.testData.answers[0]},
					EndDate:               suite.testData.postEndPollDate,
					AllowsAnswerEdits:     true,
					AllowsMultipleAnswers: false,
				},
			},
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "the poll associated with ID 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af doesn't allow multiple answers"),
		},
		{
			name: "Creator provide too many answers",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2, 3}, suite.testData.post.Creator),
			storedPost: types.Post{
				PostID:       id,
				Message:      "Post message",
				Created:      suite.testData.post.Created,
				Subspace:     suite.testData.post.Subspace,
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
				PollData: &types.PollData{
					Question:              "poll?",
					ProvidedAnswers:       suite.testData.answers,
					EndDate:               suite.testData.postEndPollDate,
					AllowsAnswerEdits:     true,
					AllowsMultipleAnswers: true,
				},
			},
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "user's answers are more than the available ones in Poll"),
		},
		{
			name: "Creator provide answers that are not the ones provided by the poll",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 3}, suite.testData.post.Creator),
			storedPost: types.Post{
				PostID:       id,
				Message:      "Post message",
				Created:      suite.testData.post.Created,
				Subspace:     "desmos",
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
				PollData: &types.PollData{
					Question:              "poll?",
					ProvidedAnswers:       suite.testData.answers,
					EndDate:               suite.testData.postEndPollDate,
					AllowsMultipleAnswers: true,
					AllowsAnswerEdits:     true,
				},
			},
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "answer with ID 3 isn't one of the poll's provided answers"),
		},
		{
			name: "Poll doesn't allow answers' edits",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.Post{
				PostID:       id,
				Message:      "Post message",
				Created:      suite.testData.post.Created,
				Subspace:     "desmos",
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
				PollData: &types.PollData{
					Question:              "poll?",
					ProvidedAnswers:       suite.testData.answers,
					EndDate:               suite.testData.postEndPollDate,
					AllowsMultipleAnswers: true,
				},
			},
			storedAnswers: &userPollAnswers,
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "post with ID 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af doesn't allow answers' edits"),
		},
		{
			name: "Answered correctly to post's poll",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, suite.testData.post.Creator),
			storedPost: types.Post{
				PostID:       id,
				Message:      "Post message",
				Created:      suite.testData.post.Created,
				LastEdited:   suite.testData.post.LastEdited,
				Subspace:     "desmos",
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
				PollData: &types.PollData{
					Question:              "poll?",
					ProvidedAnswers:       suite.testData.answers,
					EndDate:               suite.testData.postEndPollDate,
					AllowsMultipleAnswers: true,
					AllowsAnswerEdits:     true,
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			store.Set(types.PostStoreKey(test.storedPost.PostID), suite.keeper.cdc.MustMarshalBinaryBare(&test.storedPost))

			if test.storedAnswers != nil {
				suite.keeper.SavePollAnswers(suite.ctx, test.storedPost.PostID, *test.storedAnswers)
			}

			if test.storedPost.PollData != nil && test.storedPost.PollData.EndDate == suite.testData.postEndPollDateExpired {
				suite.ctx = suite.ctx.WithBlockTime(suite.testData.postEndPollDate)
			}

			handler := posts.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}

			// Valid response
			if res != nil {
				{
					// Check the data
					suite.Require().Equal(suite.keeper.cdc.MustMarshalBinaryLengthPrefixed("Answered to poll correctly"), res.Data)

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
	suite.Require().NoError(err)

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
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			for _, react := range test.existingReactions {
				react := react
				store.Set(types.ReactionsStoreKey(react.ShortCode, react.Subspace), suite.keeper.cdc.MustMarshalBinaryBare(&react))
			}

			handler := posts.NewHandler(suite.keeper)
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
				suite.keeper.cdc.MustUnmarshalBinaryBare(
					store.Get(types.ReactionsStoreKey(test.msg.ShortCode, test.msg.Subspace)),
					&storedReaction,
				)

				expected := types.NewReaction(user, test.msg.ShortCode, test.msg.Value, test.msg.Subspace)
				suite.True(expected.Equals(storedReaction))
			}

			// Invalid response
			if res == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.error.Error(), err.Error())
			}
		})
	}
}
