package keeper_test

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) Test_queryPost() {
	creationDate := time.Date(2100, 1, 1, 10, 0, 0, 0, suite.testData.timeZone)
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	computedID := types.ComputeID(creationDate, creator, subspace)
	stringID := computedID.String()

	otherCreator, err := sdk.AccAddressFromBech32("cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l")
	suite.NoError(err)

	computedID2 := types.ComputeID(creationDate, otherCreator, subspace)

	answers := []types.AnswerID{types.AnswerID(1)}

	reaction := types.NewReaction(suite.testData.postOwner, ":like:", "https://smile.jpg", "")

	tests := []struct {
		name               string
		path               []string
		storedPosts        types.Posts
		storedReactions    map[string]types.PostReactions
		registeredReaction *types.Reaction
		storedAnswers      []types.UserAnswer
		expResult          types.PostQueryResponse
		expError           error
	}{
		{
			name:               "Invalid query endpoint",
			path:               []string{"invalid", ""},
			registeredReaction: nil,
			expError:           fmt.Errorf("unknown post query endpoint"),
		},
		{
			name:               "Invalid ID returns error",
			path:               []string{types.QueryPost, ""},
			registeredReaction: nil,
			expError:           sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid postID: "),
		},
		{
			name:               "Post not found returns error",
			path:               []string{types.QueryPost, computedID.String()},
			registeredReaction: nil,
			expError:           sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Post with id f29d559522dd14484d38330a4df10115c04814d23cf35aabd9c35c80dbd5268f not found"),
		},
		{
			name: "Post without reactions is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
				types.NewPost(computedID2, computedID, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
			},
			storedAnswers:      []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			registeredReaction: nil,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
				[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
				[]types.PostReaction{},
				types.PostIDs{computedID2},
			),
		},
		{
			name: "Post without children is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
			},
			storedAnswers:      []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			registeredReaction: nil,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
				[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
				[]types.PostReaction{},
				types.PostIDs{},
			),
		},
		{
			name: "Post without medias is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithPollData(*suite.testData.post.PollData),
				types.NewPost(computedID2, computedID, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			storedReactions: map[string]types.PostReactions{
				stringID: {
					types.NewPostReaction(":like:", "https://smile.jpg", creator),
					types.NewPostReaction(":like:", "https://smile.jpg", otherCreator),
				},
			},
			registeredReaction: &reaction,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithPollData(*suite.testData.post.PollData),
				[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
				[]types.PostReaction{
					types.NewPostReaction(reaction.ShortCode, reaction.Value, creator),
					types.NewPostReaction(reaction.ShortCode, reaction.Value, otherCreator),
				},
				types.PostIDs{computedID2},
			),
		},
		{
			name: "Post without poll and poll answers is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
				types.NewPost(computedID2, computedID, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
			},
			storedReactions: map[string]types.PostReactions{
				stringID: {
					types.NewPostReaction(":like:", "https://smile.jpg", creator),
					types.NewPostReaction(":like:", "https://smile.jpg", otherCreator),
				},
			},
			registeredReaction: &reaction,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
				nil,
				[]types.PostReaction{
					types.NewPostReaction(reaction.ShortCode, reaction.Value, creator),
					types.NewPostReaction(reaction.ShortCode, reaction.Value, otherCreator),
				},
				types.PostIDs{computedID2},
			),
		},
		{
			name: "Post with all data is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
				types.NewPost(computedID2, computedID, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
			},
			storedReactions: map[string]types.PostReactions{
				stringID: {
					types.NewPostReaction(":like:", "https://smile.jpg", creator),
					types.NewPostReaction(":like:", "https://smile.jpg", otherCreator),
				},
			},
			storedAnswers:      []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			registeredReaction: &reaction,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
				[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
				[]types.PostReaction{
					types.NewPostReaction(reaction.ShortCode, reaction.Value, creator),
					types.NewPostReaction(reaction.ShortCode, reaction.Value, otherCreator),
				},
				types.PostIDs{computedID2},
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			for _, p := range test.storedPosts {
				suite.keeper.SavePost(suite.ctx, p)
			}

			if test.registeredReaction != nil {
				suite.keeper.RegisterReaction(suite.ctx, *test.registeredReaction)
			}

			for index, ans := range test.storedAnswers {
				suite.keeper.SavePollAnswers(suite.ctx, test.storedPosts[index].PostID, ans)
			}

			for postID, reactions := range test.storedReactions {
				for _, reaction := range reactions {
					err = suite.keeper.SavePostReaction(suite.ctx, types.PostID(postID), reaction)
					suite.NoError(err)
				}
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)
				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expError.Error(), err.Error())
				suite.Nil(result)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryPosts() {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)

	answers := []types.AnswerID{types.AnswerID(1)}

	tests := []struct {
		name          string
		storedPosts   types.Posts
		storedAnswers []types.UserAnswer
		params        types.QueryPostsParams
		expResponse   []types.PostQueryResponse
	}{
		{
			name: "Empty params returns all",
			storedPosts: types.Posts{
				types.NewPost(id, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
				types.NewPost(id2, id, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			params:        types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(id, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
					[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
					[]types.PostReaction{},
					types.PostIDs{id2},
				),
				types.NewPostResponse(
					types.NewPost(id2, id, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
					nil,
					[]types.PostReaction{},
					types.PostIDs{},
				),
			},
		},
		{
			name: "Empty params returns all posts without medias",
			storedPosts: types.Posts{
				types.NewPost(id2, id, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithPollData(*suite.testData.post.PollData),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			params:        types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(id2, id, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithPollData(*suite.testData.post.PollData),
					[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
					[]types.PostReaction{},
					types.PostIDs{},
				),
			},
		},
		{
			name: "Empty params returns all posts without poll data and poll answers",
			storedPosts: types.Posts{
				types.NewPost(id, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
				types.NewPost(id2, id, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
			},
			params: types.DefaultQueryPostsParams(1, 1),
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(id, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
					nil,
					[]types.PostReaction{},
					types.PostIDs{id2},
				),
			},
		},
		{
			name: "Non empty params return proper posts",
			storedPosts: types.Posts{
				types.NewPost(id, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
				types.NewPost(id2, id, "Child", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			params:        types.DefaultQueryPostsParams(1, 1),
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(id, "", "Parent", false, "", map[string]string{}, suite.testData.post.Created, creator).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
					[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
					[]types.PostReaction{},
					types.PostIDs{id2},
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			for _, p := range test.storedPosts {
				suite.keeper.SavePost(suite.ctx, p)
			}

			for index, ans := range test.storedAnswers {
				suite.keeper.SavePollAnswers(suite.ctx, test.storedPosts[index].PostID, ans)
			}

			querier := keeper.NewQuerier(suite.keeper)
			request := abci.RequestQuery{Data: suite.keeper.Cdc.MustMarshalJSON(&test.params)}
			result, err := querier(suite.ctx, []string{types.QueryPosts}, request)
			suite.NoError(err)

			expSerialized, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResponse)
			suite.NoError(err)
			suite.Equal(string(expSerialized), string(result))
		})
	}
}

func (suite *KeeperTestSuite) Test_queryPollAnswers() {
	creationDate := time.Date(2100, 1, 1, 10, 0, 0, 0, suite.testData.timeZone)
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	computedID := types.ComputeID(creationDate, creator, subspace)
	stringID := computedID.String()

	answers := []types.AnswerID{types.AnswerID(1)}

	tests := []struct {
		name          string
		path          []string
		storedPosts   types.Posts
		storedAnswers []types.UserAnswer
		expResult     types.PollAnswersQueryResponse
		expError      error
	}{
		{
			name:     "Invalid post id returns error",
			path:     []string{types.QueryPollAnswers, ""},
			expError: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid postID: "),
		},
		{
			name:     "Post not found returns error",
			path:     []string{types.QueryPollAnswers, "1"},
			expError: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "invalid postID: 1"),
		},
		{
			name:     "No post associated with ID given",
			path:     []string{types.QueryPollAnswers, stringID},
			expError: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Post with id %s not found", stringID)),
		},
		{
			name: "Post without poll returns error",
			path: []string{types.QueryPollAnswers, stringID},
			storedPosts: types.Posts{
				types.NewPost(
					computedID,
					"",
					"post with poll",
					false,
					"",
					map[string]string{},
					suite.testData.post.Created,
					suite.testData.post.Creator,
				).WithAttachments(suite.testData.post.Attachments),
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post with id f29d559522dd14484d38330a4df10115c04814d23cf35aabd9c35c80dbd5268f has no poll associated"),
		},
		{
			name: "Returns answers details of the post correctly",
			path: []string{types.QueryPollAnswers, stringID},
			storedPosts: types.Posts{
				types.NewPost(
					computedID,
					"",
					"post with poll",
					false,
					"",
					map[string]string{},
					suite.testData.post.Created,
					suite.testData.post.Creator,
				).WithAttachments(suite.testData.post.Attachments).WithPollData(*suite.testData.post.PollData),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			expResult: types.PollAnswersQueryResponse{
				PostID:         computedID,
				AnswersDetails: types.UserAnswers{types.NewUserAnswer(answers, creator)}},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, p := range test.storedPosts {
				suite.keeper.SavePost(suite.ctx, p)
			}

			for index, ans := range test.storedAnswers {
				suite.keeper.SavePollAnswers(suite.ctx, test.storedPosts[index].PostID, ans)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)

				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expError.Error(), err.Error())
				suite.Nil(result)
			}
		})
	}

}

func (suite *KeeperTestSuite) Test_queryRegisteredReactions() {
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.NoError(err)

	tests := []struct {
		name            string
		path            []string
		storedReactions types.Reactions
		expError        error
		expResult       types.Reactions
	}{
		{
			name: "PostReactions returned properly",
			path: []string{types.QueryRegisteredReactions},
			storedReactions: types.Reactions{
				types.NewReaction(creator, ":smile:", "http://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewReaction(creator, ":sad:", "http://sad.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expError: nil,
			expResult: types.Reactions{
				types.NewReaction(creator, ":sad:", "http://sad.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewReaction(creator, ":smile:", "http://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, r := range test.storedReactions {
				suite.keeper.RegisterReaction(suite.ctx, r)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)

				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expError.Error(), err.Error())
				suite.Nil(result)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryParams() {
	tests := []struct {
		name      string
		path      []string
		expResult types.Params
	}{
		{
			name:      "Returning posts parameters correctly",
			path:      []string{types.QueryParams},
			expResult: types.DefaultParams(),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.keeper.SetParams(suite.ctx, types.DefaultParams())
			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)
				suite.Equal(string(expectedIndented), string(result))
			}

		})
	}
}
