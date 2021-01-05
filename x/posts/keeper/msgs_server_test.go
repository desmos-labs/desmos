package keeper_test

import (
	"strings"
	"time"

	"github.com/desmos-labs/desmos/x/posts/keeper"

	relationshipstypes "github.com/desmos-labs/desmos/x/relationships/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestMsgServer_CreatePost() {
	tests := []struct {
		name             string
		storedPosts      []types.Post
		storedUserBlocks []relationshipstypes.UserBlock
		msg              *types.MsgCreatePost
		expError         bool
		expPosts         []types.Post
	}{
		{
			name: "Trying to store post with same id returns error",
			storedPosts: []types.Post{
				{
					PostID:         "1be6efe53a22ecf6e501ab6c8963b6dc3ae2259eacf9d012816c8bd2c8d9b73f",
					ParentID:       suite.testData.post.ParentID,
					Message:        suite.testData.post.Message,
					Created:        suite.testData.post.Created,
					AllowsComments: suite.testData.post.AllowsComments,
					Subspace:       suite.testData.post.Subspace,
					OptionalData:   suite.testData.post.OptionalData,
					Creator:        suite.testData.post.Creator,
					PollData:       suite.testData.post.PollData,
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
			expError: true,
		},
		{
			name:        "Post with new id is stored properly",
			storedPosts: nil,
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
			expError: false,
			expPosts: []types.Post{
				types.NewPost(
					"1be6efe53a22ecf6e501ab6c8963b6dc3ae2259eacf9d012816c8bd2c8d9b73f",
					suite.testData.post.ParentID,
					suite.testData.post.Message,
					suite.testData.post.AllowsComments,
					suite.testData.post.Subspace,
					suite.testData.post.OptionalData,
					suite.testData.post.Attachments,
					suite.testData.post.PollData,
					time.Time{},
					suite.ctx.BlockTime(),
					suite.testData.post.Creator,
				),
			},
		},
		{
			name: "Storing a valid post with missing parent id returns error",
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
			expError: true,
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns error",
			storedPosts: []types.Post{
				{
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
			expError: true,
		},
		{
			name: "Post with the exact same data is not posted again",
			storedPosts: []types.Post{
				types.NewPost(
					"1be6efe53a22ecf6e501ab6c8963b6dc3ae2259eacf9d012816c8bd2c8d9b73f",
					suite.testData.post.ParentID,
					suite.testData.post.Message,
					suite.testData.post.AllowsComments,
					suite.testData.post.Subspace,
					suite.testData.post.OptionalData,
					suite.testData.post.Attachments,
					suite.testData.post.PollData,
					time.Time{},
					suite.ctx.BlockTime(),
					suite.testData.post.Creator,
				),
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
			expError: true,
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
			expError: true,
		},
		{
			name: "Post tag blocked the post creator",
			storedUserBlocks: []relationshipstypes.UserBlock{
				relationshipstypes.NewUserBlock(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					suite.testData.post.Creator,
					"test",
					suite.testData.post.Subspace,
				),
			},
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
			expError: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())

			for _, post := range test.storedPosts {
				suite.keeper.SavePost(suite.ctx, post)
			}

			for _, block := range test.storedUserBlocks {
				err := suite.rk.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			handler := keeper.NewMsgServerImpl(suite.keeper)
			_, err := handler.CreatePost(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Len(suite.ctx.EventManager().Events(), 1)
				suite.Require().Equal(test.expPosts, suite.keeper.GetPosts(suite.ctx))
			}
		})
	}

}

func (suite *KeeperTestSuite) TestMsgServer_EditPost() {
	testData := []struct {
		name             string
		storedPosts      []types.Post
		storedUserBlocks []relationshipstypes.UserBlock
		timeDifference   time.Duration
		msg              *types.MsgEditPost
		expError         bool
		expPosts         []types.Post
	}{
		{
			name: "Post not found",
			msg: types.NewMsgEditPost(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"Edited message",
				nil,
				nil,
				suite.testData.post.Creator,
			),
			expError: true,
		},
		{
			name: "Invalid editor",
			storedPosts: []types.Post{
				suite.testData.post,
			},
			msg: types.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				nil,
				nil,
				"cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63",
			),
			expError: true,
		},
		{
			name: "Edit date before creation date",
			storedPosts: []types.Post{
				suite.testData.post,
			},
			timeDifference: -10,
			msg: types.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				nil,
				nil,
				suite.testData.post.Creator,
			),
			expError: true,
		},
		{
			name: "Blocked creator from tags",
			storedUserBlocks: []relationshipstypes.UserBlock{
				relationshipstypes.NewUserBlock(
					"cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63",
					suite.testData.post.Creator,
					"test",
					suite.testData.post.Subspace,
				)},
			storedPosts: []types.Post{
				suite.testData.post,
			},
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
			expError: true,
		},
		{
			name: "Valid request is handled properly without attachments and poll data",
			storedPosts: []types.Post{
				{
					PostID:         suite.testData.post.PostID,
					ParentID:       suite.testData.post.ParentID,
					Message:        "Message",
					Created:        suite.ctx.BlockTime(),
					LastEdited:     suite.testData.post.Created.AddDate(0, 0, 1),
					AllowsComments: suite.testData.post.AllowsComments,
					Subspace:       suite.testData.post.Subspace,
					OptionalData:   suite.testData.post.OptionalData,
					Creator:        suite.testData.post.Creator,
					Attachments: types.NewAttachments(
						types.NewAttachment("https://edited.com", "text/plain", nil),
					),
					PollData: types.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types.NewPollAnswers(
							types.NewPollAnswer("1", "No"),
							types.NewPollAnswer("2", "No"),
						),
						false,
						true,
					),
				},
			},
			timeDifference: time.Hour * 24,
			msg: types.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				nil,
				nil,
				suite.testData.post.Creator,
			),
			expPosts: []types.Post{
				{
					PostID:         suite.testData.post.PostID,
					ParentID:       suite.testData.post.ParentID,
					Message:        "Edited message",
					Created:        suite.ctx.BlockTime(),
					LastEdited:     suite.testData.post.Created.AddDate(0, 0, 1),
					AllowsComments: suite.testData.post.AllowsComments,
					Subspace:       suite.testData.post.Subspace,
					OptionalData:   suite.testData.post.OptionalData,
					Creator:        suite.testData.post.Creator,
					Attachments: types.NewAttachments(
						types.NewAttachment("https://edited.com", "text/plain", nil),
					),
					PollData: types.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types.NewPollAnswers(
							types.NewPollAnswer("1", "No"),
							types.NewPollAnswer("2", "No"),
						),
						false,
						true,
					),
				},
			},
		},
		{
			name: "Valid request is handled properly with attachments and poll data",
			storedPosts: []types.Post{
				suite.testData.post,
			},
			timeDifference: time.Hour * 24,
			msg: types.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				types.NewAttachments(
					types.NewAttachment("https://edited.com", "text/plain", nil),
				),
				types.NewPollData(
					"poll?",
					time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
					types.NewPollAnswers(
						types.NewPollAnswer("1", "No"),
						types.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				suite.testData.post.Creator,
			),
			expPosts: []types.Post{
				{
					PostID:         suite.testData.post.PostID,
					ParentID:       suite.testData.post.ParentID,
					Message:        "Edited message",
					Created:        suite.ctx.BlockTime(),
					LastEdited:     suite.testData.post.Created.AddDate(0, 0, 1),
					AllowsComments: suite.testData.post.AllowsComments,
					Subspace:       suite.testData.post.Subspace,
					OptionalData:   suite.testData.post.OptionalData,
					Creator:        suite.testData.post.Creator,
					Attachments: types.NewAttachments(
						types.NewAttachment("https://edited.com", "text/plain", nil),
					),
					PollData: types.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types.NewPollAnswers(
							types.NewPollAnswer("1", "No"),
							types.NewPollAnswer("2", "No"),
						),
						false,
						true,
					),
				},
			},
		},
	}

	for _, test := range testData {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())

			suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(test.timeDifference))
			for _, post := range test.storedPosts {
				suite.keeper.SavePost(suite.ctx, post)
			}

			for _, block := range test.storedUserBlocks {
				err := suite.rk.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			handler := keeper.NewMsgServerImpl(suite.keeper)
			_, err := handler.EditPost(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Len(suite.ctx.EventManager().Events(), 1)
				suite.Require().Equal(test.expPosts, suite.keeper.GetPosts(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_AddPostReaction() {
	tests := []struct {
		name                   string
		storedPosts            []types.Post
		registeredReactions    []types.RegisteredReaction
		msg                    *types.MsgAddPostReaction
		expError               bool
		expEvents              sdk.Events
		expPostReactionEntries []types.PostReactionsEntry
	}{
		{
			name: "Post not found",
			msg: types.NewMsgAddPostReaction(
				"invalid",
				":smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			expError: true,
		},
		{
			name: "Registered Reaction not found",
			storedPosts: []types.Post{
				suite.testData.post,
			},
			msg: types.NewMsgAddPostReaction(
				suite.testData.post.PostID,
				":super-smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			expError: true,
		},
		{
			name: "Valid message works properly (shortcode)",
			registeredReactions: []types.RegisteredReaction{
				suite.testData.registeredReaction,
			},
			storedPosts: []types.Post{
				suite.testData.post,
			},
			msg: types.NewMsgAddPostReaction(
				suite.testData.post.PostID,
				":smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypePostReactionAdded,
					sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, "😄"),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":smile:"),
				),
			},
			expPostReactionEntries: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(
					suite.testData.post.PostID,
					[]types.PostReaction{
						types.NewPostReaction(
							":smile:",
							"😄",
							"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
						),
					},
				),
			},
		},
		{
			name: "Valid message works properly (emoji)",
			storedPosts: []types.Post{
				suite.testData.post,
			},
			msg: types.NewMsgAddPostReaction(
				suite.testData.post.PostID,
				"🙂",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypePostReactionAdded,
					sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, "🙂"),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":slightly_smiling_face:"),
				),
			},
			expPostReactionEntries: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(
					suite.testData.post.PostID,
					[]types.PostReaction{
						types.NewPostReaction(
							":slightly_smiling_face:",
							"🙂",
							"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
						),
					}),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, reaction := range test.registeredReactions {
				suite.keeper.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, post := range test.storedPosts {
				suite.keeper.SavePost(suite.ctx, post)
			}

			handler := keeper.NewMsgServerImpl(suite.keeper)
			_, err := handler.AddPostReaction(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Len(suite.ctx.EventManager().Events(), 1)
				suite.Require().Equal(test.expPostReactionEntries, suite.keeper.GetPostReactionsEntries(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_RemovePostReaction() {
	tests := []struct {
		name                string
		storedPosts         []types.Post
		registeredReactions []types.RegisteredReaction
		existingReactions   []types.PostReactionsEntry
		msg                 *types.MsgRemovePostReaction
		expError            bool
		expEvents           sdk.Events
		expReactions        []types.PostReactionsEntry
	}{
		{
			name: "Post not found",
			msg: types.NewMsgRemovePostReaction(
				"invalid",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				"registeredReactions",
			),
			expError: true,
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
			expError: true,
		},
		{
			name: "Removing a registeredReactions using the code works properly (registered registeredReactions)",
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
			existingReactions: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(suite.testData.postID, []types.PostReaction{
					types.NewPostReaction(
						":registeredReactions:",
						"react",
						"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					),
				}),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					":registeredReactions:",
					"react",
					suite.testData.post.Subspace,
				),
			},
			msg: types.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":registeredReactions:",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypePostReactionRemoved,
					sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, "react"),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":registeredReactions:"),
				),
			},
		},
		{
			name: "Removing a registeredReactions using the code works properly (emoji shortcode)",
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
			existingReactions: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(suite.testData.postID, []types.PostReaction{
					types.NewPostReaction(
						":smile:",
						"😄",
						"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					),
				}),
			},
			msg: types.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":smile:",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypePostReactionRemoved,
					sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, "😄"),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":smile:"),
				),
			},
		},
		{
			name: "Removing a registeredReactions using the emoji works properly",
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
			existingReactions: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(suite.testData.postID, []types.PostReaction{
					types.NewPostReaction(
						":+1:",
						"👍",
						"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					),
				}),
			},
			msg: types.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				"👍",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypePostReactionRemoved,
					sdk.NewAttribute(types.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, "👍"),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":+1:"),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, reaction := range test.registeredReactions {
				suite.keeper.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, post := range test.storedPosts {
				suite.keeper.SavePost(suite.ctx, post)
			}

			for _, entry := range test.existingReactions {
				for _, reaction := range entry.Reactions {
					err := suite.keeper.SavePostReaction(suite.ctx, entry.PostId, reaction)
					suite.Require().NoError(err)
				}
			}

			handler := keeper.NewMsgServerImpl(suite.keeper)
			_, err := handler.RemovePostReaction(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
				suite.Require().Equal(test.expReactions, suite.keeper.GetPostReactionsEntries(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_AnswerPoll() {
	tests := []struct {
		name                string
		storedPosts         []types.Post
		storedAnswers       []types.UserAnswersEntry
		blockTimeDifference time.Duration
		msg                 *types.MsgAnswerPoll
		expErr              bool
		expEvents           sdk.Events
	}{
		{
			name:   "Post not found",
			msg:    types.NewMsgAnswerPoll(suite.testData.postID, []string{"1"}, suite.testData.post.Creator),
			expErr: true,
		},
		{
			name: "No poll associated with post",
			storedPosts: []types.Post{
				{
					PostID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					Message:      "Post message",
					Created:      suite.testData.post.Created,
					Subspace:     suite.testData.post.Subspace,
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			msg: types.NewMsgAnswerPoll(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Answer after poll ending",
			storedPosts: []types.Post{
				{
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
			},
			blockTimeDifference: -time.Second,
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Poll doesn't allow multiple answers",
			storedPosts: []types.Post{
				{
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
			},
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Too many answers provided",
			storedPosts: []types.Post{
				{
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
			},
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2", "3"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Provided answers are not the ones provided by the poll",
			storedPosts: []types.Post{
				{
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
			},
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "3"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Poll doesn't allow answers' edits",
			storedPosts: []types.Post{
				{
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
			},
			storedAnswers: []types.UserAnswersEntry{
				types.NewUserAnswersEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types.UserAnswer{
						types.NewUserAnswer([]string{"1", "2"}, suite.testData.post.Creator),
					},
				),
			},
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Answered correctly to post's poll",
			storedPosts: []types.Post{
				{
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
			msg: types.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeAnsweredPoll,
					sdk.NewAttribute(types.AttributeKeyPostID, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
					sdk.NewAttribute(types.AttributeKeyPollAnswerer, suite.testData.post.Creator),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(test.blockTimeDifference))
			for _, post := range test.storedPosts {
				suite.keeper.SavePost(suite.ctx, post)
			}

			for _, entry := range test.storedAnswers {
				for _, answer := range entry.UserAnswers {
					suite.keeper.SavePollAnswers(suite.ctx, entry.PostId, answer)
				}
			}

			handler := keeper.NewMsgServerImpl(suite.keeper)
			_, err := handler.AnswerPoll(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}
		})
	}

}

func (suite *KeeperTestSuite) TestMsgServer_RegisterReaction() {
	tests := []struct {
		name              string
		existingReactions []types.RegisteredReaction
		msg               *types.MsgRegisterReaction
		expError          bool
		expEvents         sdk.Events
		expReactions      []types.RegisteredReaction
	}{
		{
			name: "Reaction registered without error",
			msg: types.NewMsgRegisterReaction(
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":test:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRegisterReaction,
					sdk.NewAttribute(types.AttributeKeyReactionCreator, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, ":test:"),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, "https://smile.jpg"),
					sdk.NewAttribute(types.AttributeKeyReactionSubSpace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				),
			},
			expReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					":test:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
		{
			name: "Emoji registeredReactions returns error",
			msg: types.NewMsgRegisterReaction(
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expError: true,
		},
		{
			name: "Already registered registeredReactions returns error",
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
			expError: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, reaction := range test.existingReactions {
				suite.keeper.SaveRegisteredReaction(suite.ctx, reaction)
			}

			handler := keeper.NewMsgServerImpl(suite.keeper)
			_, err := handler.RegisterReaction(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
				suite.Require().Equal(test.expReactions, suite.keeper.GetRegisteredReactions(suite.ctx))
			}
		})
	}
}
