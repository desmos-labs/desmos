package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	tests := []struct {
		name                string
		posts               []types.Post
		answers             []types.UserAnswersEntry
		postReactions       []types.PostReactionsEntry
		registeredReactions []types.RegisteredReaction
		expStop             bool
	}{
		{
			name: "All invariants are not violated",
			posts: []types.Post{
				{
					PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Post without medias",
					Created:      suite.testData.post.Created,
					LastEdited:   time.Time{},
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
					Attachments:  suite.testData.post.Attachments,
					PollData:     suite.testData.post.PollData,
				},
				{
					PostID:         "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:        "Post without medias",
					AllowsComments: false,
					Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData:   nil,
					Created:        suite.testData.post.Created.Add(time.Hour),
					Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			answers: []types.UserAnswersEntry{
				types.NewUserAnswersEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types.UserAnswer{
						types.NewUserAnswer([]string{"1", "2"}, suite.testData.post.Creator),
					},
				),
			},
			postReactions: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types.PostReaction{
						types.NewPostReaction(
							":like:",
							"+1",
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						),
					},
				),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					suite.testData.post.Creator,
					":like:",
					"+1",
					suite.testData.post.Subspace),
			},
			expStop: true,
		},
		{
			name: "ValidPosts Invariants violated",
			posts: []types.Post{
				{
					PostID:       "1234",
					Message:      "Message",
					Created:      suite.testData.post.Created,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				}},
			answers:             nil,
			postReactions:       nil,
			registeredReactions: nil,
			expStop:             true,
		},
		{
			name: "ValidCommentsDate Invariants violated",
			posts: []types.Post{
				{
					PostID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Post without medias",
					Created:      suite.testData.post.Created,
					LastEdited:   time.Time{},
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
					Attachments:  suite.testData.post.Attachments,
					PollData:     suite.testData.post.PollData,
				},
				{
					PostID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:      "Message",
					Created:      suite.testData.postEndPollDateExpired,
					Subspace:     "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData: nil,
					Creator:      suite.testData.post.Creator,
				},
			},
			answers:             nil,
			postReactions:       nil,
			registeredReactions: nil,
			expStop:             true,
		},
		{
			name:    "ValidPostForReactions Invariants violated",
			posts:   []types.Post{},
			answers: nil,
			postReactions: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types.PostReaction{
						types.NewPostReaction(
							":like:",
							"+1",
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						),
					},
				),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					suite.testData.post.Creator,
					":like:",
					"+1",
					suite.testData.post.Subspace,
				),
			},
			expStop: true,
		},
		{
			name: "ValidPollForPollAnswers Invariants violated",
			posts: []types.Post{
				{
					PostID:         "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:        "Post without medias",
					AllowsComments: false,
					Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					OptionalData:   nil,
					Created:        suite.testData.post.Created.Add(time.Hour),
					Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			answers: []types.UserAnswersEntry{
				types.NewUserAnswersEntry(
					"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					[]types.UserAnswer{
						types.NewUserAnswer([]string{"1", "2"}, suite.testData.post.Creator),
					},
				),
			},
			postReactions:       nil,
			registeredReactions: nil,
			expStop:             true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())

			for _, post := range test.posts {
				suite.keeper.SavePost(suite.ctx, post)
			}

			for _, reaction := range test.registeredReactions {
				suite.keeper.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, entry := range test.postReactions {
				for _, reaction := range entry.Reactions {
					err := suite.keeper.SavePostReaction(suite.ctx, entry.PostId, reaction)
					suite.Require().NoError(err)
				}
			}

			for _, entry := range test.answers {
				for _, answer := range entry.UserAnswers {
					suite.keeper.SavePollAnswers(suite.ctx, entry.PostId, answer)
				}
			}

			_, stop := keeper.AllInvariants(suite.keeper)(suite.ctx)
			suite.Require().Equal(test.expStop, stop)
		})
	}
}
