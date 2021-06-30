package keeper_test

import (
	"strings"
	"time"

	keeper2 "github.com/desmos-labs/desmos/x/posts/keeper"
	types2 "github.com/desmos-labs/desmos/x/posts/types"

	profilestypes "github.com/desmos-labs/desmos/x/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestMsgServer_CreatePost() {
	tests := []struct {
		name             string
		storedPosts      []types2.Post
		storedUserBlocks []profilestypes.UserBlock
		msg              *types2.MsgCreatePost
		expError         bool
		expPosts         []types2.Post
	}{
		{
			name: "Trying to store post with same id returns error",
			storedPosts: []types2.Post{
				{
					PostID:               "b7c1193823638c3a65f4f1933e1c22928f710919fb86d01364024e407b3634af",
					ParentID:             suite.testData.post.ParentID,
					Message:              suite.testData.post.Message,
					Created:              suite.testData.post.Created,
					CommentsState:        suite.testData.post.CommentsState,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: suite.testData.post.AdditionalAttributes,
					Creator:              suite.testData.post.Creator,
					PollData:             suite.testData.post.PollData,
				},
			},
			msg: types2.NewMsgCreatePost(
				suite.testData.post.Message,
				suite.testData.post.ParentID,
				suite.testData.post.CommentsState,
				suite.testData.post.Subspace,
				suite.testData.post.AdditionalAttributes,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: true,
		},
		{
			name:        "Post with new id is stored properly",
			storedPosts: nil,
			msg: types2.NewMsgCreatePost(
				suite.testData.post.Message,
				suite.testData.post.ParentID,
				suite.testData.post.CommentsState,
				suite.testData.post.Subspace,
				suite.testData.post.AdditionalAttributes,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: false,
			expPosts: []types2.Post{
				types2.NewPost(
					"b7c1193823638c3a65f4f1933e1c22928f710919fb86d01364024e407b3634af",
					suite.testData.post.ParentID,
					suite.testData.post.Message,
					suite.testData.post.CommentsState,
					suite.testData.post.Subspace,
					suite.testData.post.AdditionalAttributes,
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
			msg: types2.NewMsgCreatePost(
				suite.testData.post.Message,
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				suite.testData.post.CommentsState,
				suite.testData.post.Subspace,
				suite.testData.post.AdditionalAttributes,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: true,
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns error",
			storedPosts: []types2.Post{
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "1234",
					Message:              "Parent post",
					Created:              suite.testData.post.Created,
					CommentsState:        suite.testData.post.CommentsState,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: suite.testData.post.AdditionalAttributes,
					Creator:              suite.testData.post.Creator,
				},
			},
			msg: types2.NewMsgCreatePost(
				suite.testData.post.Message,
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				suite.testData.post.CommentsState,
				suite.testData.post.Subspace,
				suite.testData.post.AdditionalAttributes,
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: true,
		},
		{
			name: "Post with the exact same data is not posted again",
			storedPosts: []types2.Post{
				types2.NewPost(
					"b7c1193823638c3a65f4f1933e1c22928f710919fb86d01364024e407b3634af",
					suite.testData.post.ParentID,
					suite.testData.post.Message,
					suite.testData.post.CommentsState,
					suite.testData.post.Subspace,
					suite.testData.post.AdditionalAttributes,
					suite.testData.post.Attachments,
					suite.testData.post.PollData,
					time.Time{},
					suite.ctx.BlockTime(),
					suite.testData.post.Creator,
				),
			},
			msg: types2.NewMsgCreatePost(
				suite.testData.post.Message,
				suite.testData.post.ParentID,
				suite.testData.post.CommentsState,
				suite.testData.post.Subspace,
				suite.testData.post.AdditionalAttributes,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: true,
		},
		{
			name: "Post message cannot be longer than 500 characters",
			msg: types2.NewMsgCreatePost(
				strings.Repeat("a", 550),
				suite.testData.post.ParentID,
				suite.testData.post.CommentsState,
				suite.testData.post.Subspace,
				suite.testData.post.AdditionalAttributes,
				suite.testData.post.Creator,
				suite.testData.post.Attachments,
				suite.testData.post.PollData,
			),
			expError: true,
		},
		{
			name: "Post tag blocked the post creator",
			storedUserBlocks: []profilestypes.UserBlock{
				profilestypes.NewUserBlock(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					suite.testData.post.Creator,
					"test",
					suite.testData.post.Subspace,
				),
			},
			msg: types2.NewMsgCreatePost(
				"blocked",
				suite.testData.post.ParentID,
				suite.testData.post.CommentsState,
				suite.testData.post.Subspace,
				suite.testData.post.AdditionalAttributes,
				suite.testData.post.Creator,
				types2.NewAttachments(
					types2.NewAttachment(
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
			suite.k.SetParams(suite.ctx, types2.DefaultParams())

			err := suite.sk.SaveSubspace(suite.ctx, suite.testData.subspace, suite.testData.postOwner)
			suite.Require().NoError(err)

			for _, post := range test.storedPosts {
				suite.k.SavePost(suite.ctx, post)
			}

			for _, block := range test.storedUserBlocks {
				err := suite.rk.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err = handler.CreatePost(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Len(suite.ctx.EventManager().Events(), 1)
				suite.Require().Equal(test.expPosts, suite.k.GetPosts(suite.ctx))
			}
		})
	}

}

func (suite *KeeperTestSuite) TestMsgServer_EditPost() {
	testData := []struct {
		name             string
		storedPosts      []types2.Post
		storedUserBlocks []profilestypes.UserBlock
		timeDifference   time.Duration
		msg              *types2.MsgEditPost
		expError         bool
		expPosts         []types2.Post
	}{
		{
			name: "Post not found",
			msg: types2.NewMsgEditPost(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				"Edited message",
				types2.CommentsStateUnspecified,
				nil,
				nil,
				suite.testData.post.Creator,
			),
			expError: true,
		},
		{
			name: "Invalid editor",
			storedPosts: []types2.Post{
				suite.testData.post,
			},
			msg: types2.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				types2.CommentsStateUnspecified,
				nil,
				nil,
				"cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63",
			),
			expError: true,
		},
		{
			name: "Edit date before creation date",
			storedPosts: []types2.Post{
				suite.testData.post,
			},
			timeDifference: -10,
			msg: types2.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				types2.CommentsStateUnspecified,
				nil,
				nil,
				suite.testData.post.Creator,
			),
			expError: true,
		},
		{
			name: "Blocked creator from tags",
			storedUserBlocks: []profilestypes.UserBlock{
				profilestypes.NewUserBlock(
					"cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63",
					suite.testData.post.Creator,
					"test",
					suite.testData.post.Subspace,
				)},
			storedPosts: []types2.Post{
				suite.testData.post,
			},
			msg: types2.NewMsgEditPost(
				suite.testData.post.PostID,
				"blocked",
				types2.CommentsStateUnspecified,
				types2.NewAttachments(
					types2.NewAttachment(
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
			storedPosts: []types2.Post{
				{
					PostID:               suite.testData.post.PostID,
					ParentID:             suite.testData.post.ParentID,
					Message:              "Message",
					Created:              suite.ctx.BlockTime(),
					LastEdited:           suite.testData.post.Created.AddDate(0, 0, 1),
					CommentsState:        suite.testData.post.CommentsState,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: suite.testData.post.AdditionalAttributes,
					Creator:              suite.testData.post.Creator,
					Attachments: types2.NewAttachments(
						types2.NewAttachment("https://edited.com", "text/plain", nil),
					),
					PollData: types2.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types2.NewPollAnswers(
							types2.NewPollAnswer("1", "No"),
							types2.NewPollAnswer("2", "No"),
						),
						false,
						true,
					),
				},
			},
			timeDifference: time.Hour * 24,
			msg: types2.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				types2.CommentsStateAllowed,
				nil,
				nil,
				suite.testData.post.Creator,
			),
			expPosts: []types2.Post{
				{
					PostID:               suite.testData.post.PostID,
					ParentID:             suite.testData.post.ParentID,
					Message:              "Edited message",
					Created:              suite.ctx.BlockTime(),
					LastEdited:           suite.testData.post.Created.AddDate(0, 0, 1),
					CommentsState:        types2.CommentsStateAllowed,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: suite.testData.post.AdditionalAttributes,
					Creator:              suite.testData.post.Creator,
					Attachments: types2.NewAttachments(
						types2.NewAttachment("https://edited.com", "text/plain", nil),
					),
					PollData: types2.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types2.NewPollAnswers(
							types2.NewPollAnswer("1", "No"),
							types2.NewPollAnswer("2", "No"),
						),
						false,
						true,
					),
				},
			},
		},
		{
			name: "Valid request is handled properly with attachments and poll data",
			storedPosts: []types2.Post{
				suite.testData.post,
			},
			timeDifference: time.Hour * 24,
			msg: types2.NewMsgEditPost(
				suite.testData.post.PostID,
				"Edited message",
				types2.CommentsStateUnspecified,
				types2.NewAttachments(
					types2.NewAttachment("https://edited.com", "text/plain", nil),
				),
				types2.NewPollData(
					"poll?",
					time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
					types2.NewPollAnswers(
						types2.NewPollAnswer("1", "No"),
						types2.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				suite.testData.post.Creator,
			),
			expPosts: []types2.Post{
				{
					PostID:               suite.testData.post.PostID,
					ParentID:             suite.testData.post.ParentID,
					Message:              "Edited message",
					Created:              suite.ctx.BlockTime(),
					LastEdited:           suite.testData.post.Created.AddDate(0, 0, 1),
					CommentsState:        suite.testData.post.CommentsState,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: suite.testData.post.AdditionalAttributes,
					Creator:              suite.testData.post.Creator,
					Attachments: types2.NewAttachments(
						types2.NewAttachment("https://edited.com", "text/plain", nil),
					),
					PollData: types2.NewPollData(
						"poll?",
						time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
						types2.NewPollAnswers(
							types2.NewPollAnswer("1", "No"),
							types2.NewPollAnswer("2", "No"),
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
			suite.k.SetParams(suite.ctx, types2.DefaultParams())

			err := suite.sk.SaveSubspace(suite.ctx, suite.testData.subspace, suite.testData.postOwner)
			suite.Require().NoError(err)

			suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(test.timeDifference))
			for _, post := range test.storedPosts {
				suite.k.SavePost(suite.ctx, post)
			}

			for _, block := range test.storedUserBlocks {
				err := suite.rk.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err = handler.EditPost(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Len(suite.ctx.EventManager().Events(), 1)
				suite.Require().Equal(test.expPosts, suite.k.GetPosts(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_AddPostReaction() {
	tests := []struct {
		name                   string
		storedPosts            []types2.Post
		registeredReactions    []types2.RegisteredReaction
		msg                    *types2.MsgAddPostReaction
		expError               bool
		expEvents              sdk.Events
		expPostReactionEntries []types2.PostReactionsEntry
	}{
		{
			name: "Post not found",
			msg: types2.NewMsgAddPostReaction(
				"invalid",
				":smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			expError: true,
		},
		{
			name: "Registered Reaction not found",
			storedPosts: []types2.Post{
				suite.testData.post,
			},
			msg: types2.NewMsgAddPostReaction(
				suite.testData.post.PostID,
				":super-smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			expError: true,
		},
		{
			name: "Valid message works properly (shortcode)",
			registeredReactions: []types2.RegisteredReaction{
				suite.testData.registeredReaction,
			},
			storedPosts: []types2.Post{
				suite.testData.post,
			},
			msg: types2.NewMsgAddPostReaction(
				suite.testData.post.PostID,
				":smile:",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypePostReactionAdded,
					sdk.NewAttribute(types2.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types2.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types2.AttributeKeyPostReactionValue, "😄"),
					sdk.NewAttribute(types2.AttributeKeyReactionShortCode, ":smile:"),
				),
			},
			expPostReactionEntries: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					suite.testData.post.PostID,
					[]types2.PostReaction{
						types2.NewPostReaction(
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
			storedPosts: []types2.Post{
				suite.testData.post,
			},
			msg: types2.NewMsgAddPostReaction(
				suite.testData.post.PostID,
				"🙂",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypePostReactionAdded,
					sdk.NewAttribute(types2.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types2.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types2.AttributeKeyPostReactionValue, "🙂"),
					sdk.NewAttribute(types2.AttributeKeyReactionShortCode, ":slightly_smiling_face:"),
				),
			},
			expPostReactionEntries: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(
					suite.testData.post.PostID,
					[]types2.PostReaction{
						types2.NewPostReaction(
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

			err := suite.sk.SaveSubspace(suite.ctx, suite.testData.subspace, suite.testData.postOwner)
			suite.Require().NoError(err)

			for _, reaction := range test.registeredReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, post := range test.storedPosts {
				suite.k.SavePost(suite.ctx, post)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err = handler.AddPostReaction(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Len(suite.ctx.EventManager().Events(), 1)
				suite.Require().Equal(test.expPostReactionEntries, suite.k.GetPostReactionsEntries(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_RemovePostReaction() {
	tests := []struct {
		name                string
		storedPosts         []types2.Post
		registeredReactions []types2.RegisteredReaction
		existingReactions   []types2.PostReactionsEntry
		msg                 *types2.MsgRemovePostReaction
		expError            bool
		expEvents           sdk.Events
		expReactions        []types2.PostReactionsEntry
	}{
		{
			name: "Post not found",
			msg: types2.NewMsgRemovePostReaction(
				"invalid",
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				"registeredReactions",
			),
			expError: true,
		},
		{
			name: "Reaction not found",
			storedPosts: []types2.Post{
				{
					PostID:               suite.testData.postID,
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			msg: types2.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				"😄",
			),
			expError: true,
		},
		{
			name: "Removing a registeredReactions using the code works properly (registered registeredReactions)",
			storedPosts: []types2.Post{
				{
					PostID:               suite.testData.postID,
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			existingReactions: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(suite.testData.postID, []types2.PostReaction{
					types2.NewPostReaction(
						":registeredReactions:",
						"react",
						"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					),
				}),
			},
			registeredReactions: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					":registeredReactions:",
					"react",
					suite.testData.post.Subspace,
				),
			},
			msg: types2.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":registeredReactions:",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypePostReactionRemoved,
					sdk.NewAttribute(types2.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types2.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types2.AttributeKeyPostReactionValue, "react"),
					sdk.NewAttribute(types2.AttributeKeyReactionShortCode, ":registeredReactions:"),
				),
			},
		},
		{
			name: "Removing a registeredReactions using the code works properly (emoji shortcode)",
			storedPosts: []types2.Post{
				{
					PostID:               suite.testData.postID,
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			existingReactions: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(suite.testData.postID, []types2.PostReaction{
					types2.NewPostReaction(
						":smile:",
						"😄",
						"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					),
				}),
			},
			msg: types2.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":smile:",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypePostReactionRemoved,
					sdk.NewAttribute(types2.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types2.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types2.AttributeKeyPostReactionValue, "😄"),
					sdk.NewAttribute(types2.AttributeKeyReactionShortCode, ":smile:"),
				),
			},
		},
		{
			name: "Removing a registeredReactions using the emoji works properly",
			storedPosts: []types2.Post{
				{
					PostID:               suite.testData.postID,
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			existingReactions: []types2.PostReactionsEntry{
				types2.NewPostReactionsEntry(suite.testData.postID, []types2.PostReaction{
					types2.NewPostReaction(
						":+1:",
						"👍",
						"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					),
				}),
			},
			msg: types2.NewMsgRemovePostReaction(
				suite.testData.postID,
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				"👍",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypePostReactionRemoved,
					sdk.NewAttribute(types2.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types2.AttributeKeyPostReactionOwner, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types2.AttributeKeyPostReactionValue, "👍"),
					sdk.NewAttribute(types2.AttributeKeyReactionShortCode, ":+1:"),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			err := suite.sk.SaveSubspace(suite.ctx, suite.testData.subspace, suite.testData.postOwner)
			suite.Require().NoError(err)

			for _, reaction := range test.registeredReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, post := range test.storedPosts {
				suite.k.SavePost(suite.ctx, post)
			}

			for _, entry := range test.existingReactions {
				for _, reaction := range entry.Reactions {
					err := suite.k.SavePostReaction(suite.ctx, entry.PostID, reaction)
					suite.Require().NoError(err)
				}
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err = handler.RemovePostReaction(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
				suite.Require().Equal(test.expReactions, suite.k.GetPostReactionsEntries(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_AnswerPoll() {
	tests := []struct {
		name                string
		storedPosts         []types2.Post
		storedAnswers       []types2.UserAnswer
		blockTimeDifference time.Duration
		msg                 *types2.MsgAnswerPoll
		expErr              bool
		expEvents           sdk.Events
	}{
		{
			name:   "Post not found",
			msg:    types2.NewMsgAnswerPoll(suite.testData.postID, []string{"1"}, suite.testData.post.Creator),
			expErr: true,
		},
		{
			name: "No poll associated with post",
			storedPosts: []types2.Post{
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
				},
			},
			msg: types2.NewMsgAnswerPoll(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Answer after poll ending",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					PollData: &types2.PollData{
						Question:          "poll?",
						ProvidedAnswers:   types2.PollAnswers{suite.testData.answers[0]},
						EndDate:           suite.testData.postEndPollDateExpired,
						AllowsAnswerEdits: true,
					},
				},
			},
			blockTimeDifference: -time.Second,
			msg: types2.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Poll doesn't allow multiple answers",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					PollData: &types2.PollData{
						Question:              "poll?",
						ProvidedAnswers:       types2.PollAnswers{suite.testData.answers[0]},
						EndDate:               suite.testData.postEndPollDate,
						AllowsAnswerEdits:     true,
						AllowsMultipleAnswers: false,
					},
				},
			},
			msg: types2.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Too many answers provided",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             suite.testData.post.Subspace,
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					PollData: &types2.PollData{
						Question:              "poll?",
						ProvidedAnswers:       suite.testData.answers,
						EndDate:               suite.testData.postEndPollDate,
						AllowsAnswerEdits:     true,
						AllowsMultipleAnswers: true,
					},
				},
			},
			msg: types2.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2", "3"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Provided answers are not the ones provided by the poll",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             "desmos",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					PollData: &types2.PollData{
						Question:              "poll?",
						ProvidedAnswers:       suite.testData.answers,
						EndDate:               suite.testData.postEndPollDate,
						AllowsMultipleAnswers: true,
						AllowsAnswerEdits:     true,
					},
				},
			},
			msg: types2.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "3"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Poll doesn't allow answers' edits",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					Subspace:             "desmos",
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					PollData: &types2.PollData{
						Question:              "poll?",
						ProvidedAnswers:       suite.testData.answers,
						EndDate:               suite.testData.postEndPollDate,
						AllowsMultipleAnswers: true,
					},
				},
			},
			storedAnswers: []types2.UserAnswer{
				types2.NewUserAnswer("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af", suite.testData.post.Creator, []string{"1", "2"}),
			},
			msg: types2.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			expErr: true,
		},
		{
			name: "Answered correctly to post's poll",
			storedPosts: []types2.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post message",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             suite.testData.subspace.ID,
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					PollData: &types2.PollData{
						Question:              "poll?",
						ProvidedAnswers:       suite.testData.answers,
						EndDate:               suite.testData.postEndPollDate,
						AllowsMultipleAnswers: true,
						AllowsAnswerEdits:     true,
					},
				},
			},
			msg: types2.NewMsgAnswerPoll(
				"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				[]string{"1", "2"},
				suite.testData.post.Creator,
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeAnsweredPoll,
					sdk.NewAttribute(types2.AttributeKeyPostID, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
					sdk.NewAttribute(types2.AttributeKeyPollAnswerer, suite.testData.post.Creator),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			err := suite.sk.SaveSubspace(suite.ctx, suite.testData.subspace, suite.testData.postOwner)
			suite.Require().NoError(err)

			suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(test.blockTimeDifference))
			for _, post := range test.storedPosts {
				suite.k.SavePost(suite.ctx, post)
			}

			for _, answer := range test.storedAnswers {
				suite.k.SaveUserAnswer(suite.ctx, answer)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err = handler.AnswerPoll(sdk.WrapSDKContext(suite.ctx), test.msg)

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
		existingReactions []types2.RegisteredReaction
		msg               *types2.MsgRegisterReaction
		expError          bool
		expEvents         sdk.Events
		expReactions      []types2.RegisteredReaction
	}{
		{
			name: "Reaction registered without error",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":test:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expError: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeRegisterReaction,
					sdk.NewAttribute(types2.AttributeKeyReactionCreator, "cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg"),
					sdk.NewAttribute(types2.AttributeKeyReactionShortCode, ":test:"),
					sdk.NewAttribute(types2.AttributeKeyPostReactionValue, "https://smile.jpg"),
					sdk.NewAttribute(types2.AttributeKeyReactionSubSpace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				),
			},
			expReactions: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					":test:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
		{
			name: "Emoji registeredReactions returns error",
			msg: types2.NewMsgRegisterReaction(
				"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
				":smile:",
				"https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expError: true,
		},
		{
			name: "Already registered registeredReactions returns error",
			existingReactions: []types2.RegisteredReaction{
				types2.NewRegisteredReaction(
					"cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg",
					":test:",
					"https://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types2.NewMsgRegisterReaction(
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

			err := suite.sk.SaveSubspace(suite.ctx, suite.testData.subspace, suite.testData.postOwner)
			suite.Require().NoError(err)

			for _, reaction := range test.existingReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err = handler.RegisterReaction(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
				suite.Require().Equal(test.expReactions, suite.k.GetRegisteredReactions(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_ReportPost() {
	tests := []struct {
		name          string
		storedPosts   []types2.Post
		storedReports []types2.Report
		msg           *types2.MsgReportPost
		expErr        bool
		expEvents     sdk.Events
		expReports    []types2.Report
	}{
		{
			name: "invalid post id",
			msg: types2.NewMsgReportPost(
				"post_id",
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name:        "post not found",
			storedPosts: nil,
			msg: types2.NewMsgReportPost(
				suite.testData.postID,
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "double report",
			storedPosts: []types2.Post{
				{
					PostID:               suite.testData.postID,
					Message:              "Post",
					Created:              suite.testData.postCreationDate,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
				},
			},
			storedReports: []types2.Report{
				types2.NewReport(
					suite.testData.postID,
					"type",
					"message",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				),
			},
			msg: types2.NewMsgReportPost(
				suite.testData.postID,
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "message handled correctly",
			storedPosts: []types2.Post{
				{
					PostID:               suite.testData.postID,
					Message:              "Post",
					Created:              suite.testData.postCreationDate,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
				},
			},
			msg: types2.NewMsgReportPost(
				suite.testData.postID,
				"type",
				"message",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypePostReported,
					sdk.NewAttribute(types2.AttributeKeyPostID, suite.testData.postID),
					sdk.NewAttribute(types2.AttributeKeyReportOwner, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				),
			},
			expReports: []types2.Report{
				types2.NewReport(
					suite.testData.postID,
					"type",
					"message",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			err := suite.sk.SaveSubspace(suite.ctx, suite.testData.subspace, suite.testData.postOwner)
			suite.Require().NoError(err)

			for _, post := range test.storedPosts {
				suite.k.SavePost(suite.ctx, post)
			}

			for _, report := range test.storedReports {
				err := suite.k.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			server := keeper2.NewMsgServerImpl(suite.k)
			_, err = server.ReportPost(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				reports := suite.k.GetAllReports(suite.ctx)
				suite.Require().Equal(test.expReports, reports)
			}
		})
	}
}
