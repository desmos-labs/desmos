package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

func (suite *KeeperTestSuite) TestInvariants() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")

	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	parentPost := types.NewPost(
		id,
		"",
		"Post without medias",
		false,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		map[string]string{},
		suite.testData.post.Created,
		suite.testData.post.Creator,
	).WithMedias(suite.testData.post.Medias).WithPollData(*suite.testData.post.PollData)

	commentPost := types.Post{
		PostID:         id2,
		ParentID:       id,
		Message:        "Post without medias",
		AllowsComments: false,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        suite.testData.post.Created.Add(time.Hour),
		Creator:        user,
	}

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}

	postReaction := types.NewPostReaction(":like:", "+1", user)
	reaction := types.NewReaction(suite.testData.post.Creator, ":like:", "+1", suite.testData.post.Subspace)
	answer := types.NewUserAnswer(answers, suite.testData.post.Creator)

	tests := []struct {
		name         string
		posts        types.Posts
		answers      *types.UserAnswer
		postReaction *types.PostReaction
		reaction     *types.Reaction
		expResponse  string
		expBool      bool
	}{
		{
			name:         "All invariants are not violated",
			posts:        types.Posts{parentPost, commentPost},
			answers:      &answer,
			postReaction: &postReaction,
			reaction:     &reaction,
			expResponse:  "Every invariant condition is fulfilled correctly",
			expBool:      true,
		},
		{
			name: "ValidPosts Invariants violated",
			posts: types.Posts{types.NewPost(
				"1234",
				"",
				"Message",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				suite.testData.post.Created,
				suite.testData.post.Creator,
			)},
			answers:      nil,
			postReaction: nil,
			reaction:     nil,
			expResponse:  "posts: invalid posts IDs invariant\nThe following posts are invalid:\n 1234\n\n",
			expBool:      true,
		},
		{
			name: "ValidCommentsDate Invariants violated",
			posts: types.Posts{parentPost,
				types.NewPost(
					commentPost.PostID,
					parentPost.PostID,
					"Message",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					suite.testData.postEndPollDateExpired,
					suite.testData.post.Creator,
				)},
			answers:      nil,
			postReaction: nil,
			reaction:     nil,
			expResponse:  "posts: comments dates before parent post date invariant\nThe following post IDs referred to posts which are invalid comments having creation date before parent post creation date:\n f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd\n\n",
			expBool:      true,
		},
		{
			name:         "ValidPostForReactions Invariants violated",
			posts:        types.Posts{},
			answers:      nil,
			postReaction: &postReaction,
			reaction:     &reaction,
			expResponse:  "posts: posts reactions refers to non existing posts invariant\nThe following reactions refer to posts that do not exist:\n [Shortcode] :like: [Value] +1 [Owner] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns\n\n",
			expBool:      true,
		},
		{
			name:         "ValidPollForPollAnswers Invariants violated",
			posts:        types.Posts{commentPost},
			answers:      &answer,
			postReaction: nil,
			reaction:     nil,
			expResponse:  "posts: poll answers refers to posts without poll invariant\nThe following answers refer to a post that either does not exist or has no poll associated to it:\n User: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 \nAnswers IDs: 1 2\n\n",
			expBool:      true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())
			for _, post := range test.posts {
				suite.keeper.SavePost(suite.ctx, post)
			}
			if test.reaction != nil && test.postReaction != nil {
				suite.keeper.RegisterReaction(suite.ctx, *test.reaction)
				// nolint: errcheck
				suite.keeper.SavePostReaction(suite.ctx, parentPost.PostID, *test.postReaction)
			}
			if test.answers != nil {
				suite.keeper.SavePollAnswers(suite.ctx, test.posts[0].PostID, *test.answers)
			}

			res, stop := keeper.AllInvariants(suite.keeper)(suite.ctx)

			suite.Equal(test.expResponse, res)
			suite.Equal(test.expBool, stop)

		})
	}
}
