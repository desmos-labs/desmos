package keeper_test

import (
	"strings"
	"time"

	"github.com/desmos-labs/desmos/x/posts/keeper"

	"github.com/desmos-labs/desmos/x/posts"
	relationshipstypes "github.com/desmos-labs/desmos/x/relationships/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) Test_handleMsgCreatePost() {
	tests := []struct {
		name        string
		storedPosts types.Posts
		msg         *types.MsgCreatePost
		expPost     types.Post
		expError    error
	}{
		{
			name: "Trying to store post with same id returns expError",
			storedPosts: types.Posts{
				types.Post{
					PostID:         "040b0c16cd541101d24100e4a9c90e4dbaebbee977a94d673f79591cbb5f4465",
					ParentID:       suite.testData.post.ParentID,
					Message:        suite.testData.post.Message,
					Created:        suite.testData.post.Created,
					AllowsComments: suite.testData.post.AllowsComments,
					Subspace:       suite.testData.post.Subspace,
					OptionalData:   suite.testData.post.OptionalData,
					Creator:        suite.testData.post.Creator,
				},
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
				"the provided post conflicts with the one having id 040b0c16cd541101d24100e4a9c90e4dbaebbee977a94d673f79591cbb5f4465"),
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
				suite.testData.post.ParentID,
				suite.testData.post.Message,
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.ctx.BlockTime(),
				suite.testData.post.Creator,
			).WithAttachments(
				suite.testData.post.Attachments,
			).WithPollData(
				*suite.testData.post.PollData,
			),
		},
		{
			name: "Storing a valid post with missing parent id returns expError",
			msg: types.NewMsgCreatePost(
				suite.testData.post.Message,
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"parent post with id f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd not found",
			),
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns expError",
			storedPosts: types.Posts{
				types.Post{
					PostID:         "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
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
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				suite.testData.post.AllowsComments,
				suite.testData.post.Subspace,
				suite.testData.post.OptionalData,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
				types.NewAttachments(
					types.NewAttachment(
						"http://uri.com",
						"text/plain",
						[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
					),
				),
				suite.testData.post.PollData,
			),
			expError: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"The user with address cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns has blocked you",
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.k.SetParams(suite.ctx, types.DefaultParams())
			store := suite.ctx.KVStore(suite.storeKey)

			for _, p := range test.storedPosts {
				store.Set(types.PostStoreKey(p.PostID), suite.cdc.MustMarshalBinaryBare(&p))
			}

			if test.msg.Message == "blocked" {
				for _, attachment := range test.msg.Attachments {
					for _, tag := range attachment.Tags {
						block := relationshipstypes.NewUserBlock(tag, suite.testData.post.Creator, "test", "")
						err := suite.rk.SaveUserBlock(suite.ctx, block)
						suite.Require().NoError(err)
					}
				}
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.CreatePost(sdk.WrapSDKContext(suite.ctx), test.msg)

			// Valid response
			if test.expError == nil {
				// Check the post
				var stored types.Post
				suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.expPost.PostID)), &stored)

				suite.True(stored.Equal(test.expPost), "Expected: %s, actual: %s", test.expPost, stored)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypePostCreated,
					sdk.NewAttribute(types.AttributeKeyPostID, test.expPost.PostID),
					sdk.NewAttribute(types.AttributeKeyPostParentID, test.expPost.ParentID),
					sdk.NewAttribute(types.AttributeKeyPostCreationTime, test.expPost.Created.Format(time.RFC3339)),
					sdk.NewAttribute(types.AttributeKeyPostOwner, test.expPost.Creator),
				)
				suite.Len(suite.ctx.EventManager(), 1)
				suite.Contains(suite.ctx.EventManager(), creationEvent)
			}

			// Invalid response
			if test.expError != nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expError.Error(), err.Error())
			}
		})
	}

}

func (suite *KeeperTestSuite) Test_handleMsgEditPost() {
	editedPollData := types.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, time.FixedZone("UTC", 0)),
		types.NewPollAnswers(
			types.NewPollAnswer("1", "No"),
			types.NewPollAnswer("2", "No"),
		),
		false,
		true,
	)
	editedAttachments := types.NewAttachments(types.NewAttachment("https://edited.com", "text/plain", nil))

	testData := []struct {
		name       string
		storedPost *types.Post
		msg        *types.MsgEditPost
		expError   error
		expPost    *types.Post
	}{
		{
			name:       "Post not found",
			storedPost: nil,
			msg: types.NewMsgEditPost(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"Edited message",
				nil,
				nil,
				suite.testData.post.Creator,
			),
			expError: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af not found",
			),
		},
		{
			name:       "Invalid editor",
			storedPost: &suite.testData.post,
			msg: types.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				nil,
				nil,
				"cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63",
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner"),
		},
		{
			name:       "Edit date before creation date",
			storedPost: &suite.testData.post,
			msg: types.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				nil,
				nil,
				suite.testData.post.Creator,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "edit date cannot be before creation date"),
		},
		{
			name:       "Blocked creator from tags",
			storedPost: &suite.testData.post,
			msg: types.NewMsgEditPost(
				suite.testData.post.PostID,
				"blocked",
				types.NewAttachments(
					types.NewAttachment(
						"https://edited.com",
						"text/plain",
						[]string{"cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63"},
					),
				),
				nil,
				suite.testData.post.Creator,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"The user with address cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63 has blocked you"),
		},
		{
			name:       "Valid request is handled properly without attachments and pollData",
			storedPost: &suite.testData.post,
			msg: types.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				editedAttachments,
				&editedPollData,
				suite.testData.post.Creator,
			),
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
			suite.k.SetParams(suite.ctx, types.DefaultParams())

			store := suite.ctx.KVStore(suite.storeKey)

			if test.msg.Message == "blocked" {
				for _, attachment := range test.msg.Attachments {
					for _, tag := range attachment.Tags {
						block := relationshipstypes.NewUserBlock(tag, suite.testData.post.Creator, "test", "")
						err := suite.rk.SaveUserBlock(suite.ctx, block)
						suite.Require().NoError(err)
					}
				}
				test.storedPost.Created = suite.ctx.BlockTime()
				suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().AddDate(0, 0, 1))
			}

			if test.storedPost != nil {
				if test.expPost != nil {
					test.storedPost.Created = suite.ctx.BlockTime()
					suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().AddDate(0, 0, 1))
					test.expPost.LastEdited = suite.ctx.BlockTime()
				}
				store.Set(types.PostStoreKey(test.storedPost.PostID), suite.cdc.MustMarshalBinaryBare(test.storedPost))
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.EditPost(sdk.WrapSDKContext(suite.ctx), test.msg)

			// Valid response
			if err == nil {
				suite.Contains(suite.ctx.EventManager().Events(), sdk.NewEvent(
					types.EventTypePostEdited,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID),
					sdk.NewAttribute(types.AttributeKeyPostEditTime, test.expPost.LastEdited.Format(time.RFC3339)),
				))

				var stored types.Post
				suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.storedPost.PostID)), &stored)
				suite.True(test.expPost.Equal(stored))
			}

			// Invalid response
			if err != nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expError.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAddPostReaction() {
	tests := []struct {
		name               string
		existingPost       *types.Post
		msg                *types.MsgAddPostReaction
		registeredReaction *types.RegisteredReaction
		error              error
		expEvent           sdk.Event
	}{
		{
			name: "Post not found",
			msg: types.NewMsgAddPostReaction(
				"invalid",
				":smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id invalid not found"),
		},
		{
			name:         "Registered Reaction not found",
			existingPost: &suite.testData.post,
			msg: types.NewMsgAddPostReaction(
				suite.testData.post.PostID,
				":super-smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "short code :super-smile: must be registered before using it"),
		},
		{
			name:         "Valid message works properly (shortcode)",
			existingPost: &suite.testData.post,
			msg: types.NewMsgAddPostReaction(
				suite.testData.post.PostID,
				":smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			registeredReaction: &suite.testData.registeredReaction,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionAdded,
				sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, "😄"),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":smile:"),
			),
		},
		{
			name:         "Valid message works properly (emoji)",
			existingPost: &suite.testData.post,
			msg:          types.NewMsgAddPostReaction(suite.testData.post.PostID, "🙂", "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionAdded,
				sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
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
			store := suite.ctx.KVStore(suite.storeKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), suite.cdc.MustMarshalBinaryBare(test.existingPost))
			}

			if test.registeredReaction != nil {
				suite.k.SaveRegisteredReaction(suite.ctx, *test.registeredReaction)
			}

			handler := posts.NewHandler(suite.k)
			res, err := handler(suite.ctx, test.msg)

			// Valid response
			if res != nil {
				suite.Contains(res.Events, test.expEvent)

				// Check the post
				var storedPost types.Post
				suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.msg.PostID)), &storedPost)
				suite.True(test.existingPost.Equal(storedPost))

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

				var storedReactions keeper.WrappedPostReactions
				suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				suite.Contains(storedReactions.Reactions, types.NewPostReaction(reactShortcode, reactValue, test.msg.User))

				// Check the registered reactions
				registeredReactions := suite.k.GetRegisteredReactions(suite.ctx)
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
	tests := []struct {
		name                string
		storedPosts         []types.Post
		registeredReactions []types.RegisteredReaction
		existingReactions   []types.PostReaction
		msg                 *types.MsgRemovePostReaction
		error               error
		expEvent            sdk.Event
	}{
		{
			name: "Post not found",
			msg: types.NewMsgRemovePostReaction(
				"invalid",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				"reaction",
			),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id invalid not found"),
		},
		{
			name: "Reaction not found",
			storedPosts: []types.Post{
				{
					PostID:       suite.testData.postID,
					Message:      "Post message",
					Created:      suite.testData.post.Created,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			msg: types.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				"😄",
			),
			error: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"cannot remove the reaction with value :smile: from user %s as it does not exist",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
		},
		{
			name: "Removing a reaction using the code works properly (registered reaction)",
			storedPosts: []types.Post{
				{
					PostID:       suite.testData.postID,
					Message:      "Post message",
					Created:      suite.testData.post.Created,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			existingReactions: []types.PostReaction{
				types.NewPostReaction(
					":reaction:",
					"react",
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					":reaction:",
					"react",
					suite.testData.post.Subspace,
				),
			},
			msg: types.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":reaction:",
			),
			error: nil,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionRemoved,
				sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, "react"),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":reaction:"),
			),
		},
		{
			name: "Removing a reaction using the code works properly (emoji shortcode)",
			storedPosts: []types.Post{
				{
					PostID:       suite.testData.postID,
					Message:      "Post message",
					Created:      suite.testData.post.Created,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			existingReactions: []types.PostReaction{
				types.NewPostReaction(
					":smile:",
					"😄",
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				),
			},
			msg: types.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":smile:",
			),
			error: nil,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionRemoved,
				sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, "😄"),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":smile:"),
			),
		},
		{
			name: "Removing a reaction using the emoji works properly",
			storedPosts: []types.Post{
				{
					PostID:       suite.testData.postID,
					Message:      "Post message",
					Created:      suite.testData.post.Created,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			existingReactions: []types.PostReaction{
				types.NewPostReaction(
					":+1:",
					"👍",
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				),
			},
			msg: types.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				"👍",
			),
			error: nil,
			expEvent: sdk.NewEvent(
				types.EventTypePostReactionRemoved,
				sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
				sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
				sdk.NewAttribute(types.AttributeKeyPostReactionValue, "👍"),
				sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":+1:"),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)

			for _, reaction := range test.registeredReactions {
				key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)
				store.Set(key, suite.cdc.MustMarshalBinaryBare(&reaction))
			}

			for _, post := range test.storedPosts {
				store.Set(types.PostStoreKey(post.PostID), suite.cdc.MustMarshalBinaryBare(&post))

				key := types.PostReactionsStoreKey(post.PostID)
				wrapped := keeper.WrappedPostReactions{Reactions: test.existingReactions}
				store.Set(key, suite.cdc.MustMarshalBinaryBare(&wrapped))
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.RemovePostReaction(sdk.WrapSDKContext(suite.ctx), test.msg)

			// Valid response
			if err == nil {
				suite.Contains(suite.ctx.EventManager().Events(), test.expEvent)

				var storedPost types.Post
				suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(suite.testData.post.PostID)), &storedPost)
				suite.Require().Contains(test.storedPosts, storedPost)

				var storedReactions keeper.WrappedPostReactions
				suite.cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				suite.NotContains(storedReactions.Reactions, test.existingReactions)
			}

			// Invalid response
			if err != nil {
				suite.NotNil(err)
				suite.Require().Equal(test.error.Error(), err.Error())
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgAnswerPollPost() {
	tests := []struct {
		name          string
		storedPost    types.Post
		storedAnswers []types.UserAnswer
		msg           *types.MsgAnswerPoll
		expErr        error
	}{
		{
			name: "Post not found",
			msg: types.NewMsgAnswerPoll(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
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
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"post with id f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd doesn't exist",
			),
		},
		{
			name: "No poll associated with post",
			msg: types.NewMsgAnswerPoll(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			storedPost: types.Post{
				PostID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:      "Post message",
				Created:      suite.testData.post.Created,
				Subspace:     suite.testData.post.Subspace,
				OptionalData: nil,
				Creator:      suite.testData.post.Creator,
			},
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				"no poll associated with ID: f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			),
		},
		{
			name: "Answer after poll closure",
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1"},
				suite.testData.post.Creator,
			),
			storedPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
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
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"the poll associated with ID %s was closed at %s",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				suite.testData.postEndPollDateExpired,
			),
		},
		{
			name: "Poll doesn't allow multiple answers",
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			storedPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
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
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"the poll associated with ID %s doesn't allow multiple answers",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			),
		},
		{
			name: "Creator provide too many answers",
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2", "3"},
				suite.testData.post.Creator,
			),
			storedPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
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
				sdkerrors.ErrInvalidRequest,
				"user's answers are more than the available ones in Poll",
			),
		},
		{
			name: "Creator provide answers that are not the ones provided by the poll",
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "3"},
				suite.testData.post.Creator,
			),
			storedPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
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
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			storedPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
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
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, suite.testData.post.Creator),
			},
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"post with ID %s doesn't allow answers' edits",
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			),
		},
		{
			name: "Answered correctly to post's poll",
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			storedPost: types.Post{
				PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
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
			store := suite.ctx.KVStore(suite.storeKey)
			store.Set(types.PostStoreKey(test.storedPost.PostID), suite.cdc.MustMarshalBinaryBare(&test.storedPost))

			for _, answer := range test.storedAnswers {
				suite.k.SavePollAnswers(suite.ctx, test.storedPost.PostID, answer)
			}

			if test.storedPost.PollData != nil && test.storedPost.PollData.EndDate == suite.testData.postEndPollDateExpired {
				suite.ctx = suite.ctx.WithBlockTime(suite.testData.postEndPollDate)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.AnswerPoll(sdk.WrapSDKContext(suite.ctx), test.msg)

			// Valid response
			if err == nil {

				// Check the events
				answerEvent := sdk.NewEvent(
					types.EventTypeAnsweredPoll,
					sdk.NewAttribute(types.AttributeKeyPostID, test.storedPost.PostID),
					sdk.NewAttribute(types.AttributeKeyPollAnswerer, suite.testData.post.Creator),
				)

				suite.Len(suite.ctx.EventManager().Events(), 1)
				suite.Contains(suite.ctx.EventManager().Events(), answerEvent)
			}

			// Invalid response
			if err != nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
			}
		})
	}

}

func (suite *KeeperTestSuite) Test_handleMsgRegisterReaction() {
	tests := []struct {
		name              string
		existingReactions []types.RegisteredReaction
		msg               *types.MsgRegisterReaction
		error             error
	}{
		{
			name: "Reaction registered without error",
			msg: types.NewMsgRegisterReaction(
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":test:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			error: nil,
		},
		{
			name: "Emoji reaction returns error",
			msg: types.NewMsgRegisterReaction(
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
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
			existingReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					":test:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types.NewMsgRegisterReaction(
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
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
			store := suite.ctx.KVStore(suite.storeKey)

			for _, react := range test.existingReactions {
				store.Set(types.ReactionsStoreKey(react.ShortCode, react.Subspace), suite.cdc.MustMarshalBinaryBare(&react))
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.RegisterReaction(sdk.WrapSDKContext(suite.ctx), test.msg)

			// Valid response
			if err == nil {
				suite.Contains(suite.ctx.EventManager().Events(), sdk.NewEvent(
					types.EventTypeRegisterReaction,
					sdk.NewAttribute(types.AttributeKeyReactionCreator, test.msg.Creator),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, test.msg.ShortCode),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, test.msg.Value),
					sdk.NewAttribute(types.AttributeKeyReactionSubSpace, test.msg.Subspace),
				))

				var storedReaction types.RegisteredReaction
				suite.cdc.MustUnmarshalBinaryBare(
					store.Get(types.ReactionsStoreKey(test.msg.ShortCode, test.msg.Subspace)),
					&storedReaction,
				)

				expected := types.NewRegisteredReaction(
					test.msg.Creator,
					test.msg.ShortCode,
					test.msg.Value,
					test.msg.Subspace,
				)
				suite.True(expected.Equals(storedReaction))
			}

			// Invalid response
			if err != nil {
				suite.NotNil(err)
				suite.Require().Equal(test.error.Error(), err.Error())
			}
		})
	}
}
