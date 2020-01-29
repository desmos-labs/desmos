package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/tendermint/abci/types"
)

func Test_queryPost(t *testing.T) {
	creator, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	otherCreator, _ := sdk.AccAddressFromBech32("cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l")
	tests := []struct {
		name            string
		path            []string
		storedPosts     types.Posts
		storedReactions map[types.PostID]types.Reactions
		expResult       types.PostQueryResponse
		expError        sdk.Error
	}{
		{
			name:     "Invalid ID returns error",
			path:     []string{types.QueryPost, ""},
			expError: sdk.ErrUnknownRequest("Invalid post id: "),
		},
		{
			name:     "Post not found returns error",
			path:     []string{types.QueryPost, "1"},
			expError: sdk.ErrUnknownRequest("Post with id 1 not found"),
		},
		{
			name: "Post without likes is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
			},
			path: []string{types.QueryPost, "1"},
			expResult: types.NewPostResponse(
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.Reactions{},
				types.PostIDs{types.PostID(2)},
			),
		},
		{
			name: "Post without children is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
			},
			path: []string{types.QueryPost, "1"},
			expResult: types.NewPostResponse(
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.Reactions{},
				types.PostIDs{},
			),
		},
		{
			name: "Post with all data is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
			},
			storedReactions: map[types.PostID]types.Reactions{
				types.PostID(1): {
					types.NewReaction("Like", creator),
					types.NewReaction("Like", otherCreator),
				},
			},
			path: []string{types.QueryPost, "1"},
			expResult: types.NewPostResponse(
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.Reactions{
					types.NewReaction("Like", creator),
					types.NewReaction("Like", otherCreator),
				},
				types.PostIDs{types.PostID(2)},
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, p := range test.storedPosts {
				k.SavePost(ctx, p)
			}

			for postID, reactions := range test.storedReactions {
				for _, reaction := range reactions {
					_ = k.SaveReaction(ctx, postID, reaction)
				}
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if test.expError != nil {
				assert.Equal(t, test.expError, err)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				expectedIndented, _ := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				assert.Equal(t, string(expectedIndented), string(result))
			}
		})
	}
}

func Test_queryPosts(t *testing.T) {
	creator, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	tests := []struct {
		name        string
		storedPosts types.Posts
		params      types.QueryPostsParams
		expResponse []types.PostQueryResponse
	}{
		{
			name: "Empty params returns all",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
			},
			params: types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
					types.Reactions{},
					types.PostIDs{types.PostID(2)},
				),
				types.NewPostResponse(
					types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
					types.Reactions{},
					types.PostIDs{},
				),
			},
		},
		{
			name: "Empty params returns all posts without medias",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, nil, testPost.PollData),
			},
			params: types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, nil, testPost.PollData),
					types.Reactions{},
					types.PostIDs{},
				),
			},
		},
		{
			name: "Non empty params return proper posts",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
			},
			params: types.DefaultQueryPostsParams(1, 1),
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
					types.Reactions{},
					types.PostIDs{types.PostID(2)},
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			for _, p := range test.storedPosts {
				k.SavePost(ctx, p)
			}

			querier := keeper.NewQuerier(k)
			request := abci.RequestQuery{Data: k.Cdc.MustMarshalJSON(&test.params)}
			result, err := querier(ctx, []string{types.QueryPosts}, request)

			expSerialized, _ := codec.MarshalJSONIndent(k.Cdc, &test.expResponse)
			assert.Nil(t, err)
			assert.Equal(t, string(expSerialized), string(result))
		})
	}
}

func Test_queryPollUserAnswers(t *testing.T) {

	creator, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")

	tests := []struct {
		name          string
		path          []string
		storedAnswers []uint64
		storedPost    types.Post
		expResult     types.PollUserAnswersQueryResponse
		expError      sdk.Error
	}{
		{
			name:     "Invalid post id returns error",
			path:     []string{types.QueryPollUserAnswer, "", ""},
			expError: sdk.ErrUnknownRequest("Invalid post id: "),
		},
		{
			name:     "Empty address returns error",
			path:     []string{types.QueryPollUserAnswer, "1", ""},
			expError: sdk.ErrInvalidAddress("Address cannot be empty"),
		},
		{
			name:     "Invalid address returns error",
			path:     []string{types.QueryPollUserAnswer, "1", "invalid"},
			expError: sdk.ErrUnknownRequest("Invalid bech32 addr: invalid"),
		},
		{
			name:     "Post not found returns error",
			path:     []string{types.QueryPollUserAnswer, "1", creator.String()},
			expError: sdk.ErrUnknownRequest("Post with id 1 not found"),
		},
		{
			name:          "User's answers are returned properly",
			path:          []string{types.QueryPollUserAnswer, "1", creator.String()},
			storedAnswers: []uint64{1, 2},
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			),
			expResult: types.PollUserAnswersQueryResponse{
				PostID:  types.PostID(1),
				User:    creator,
				Answers: []uint64{1, 2},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			k.SavePost(ctx, test.storedPost)
			if len(test.storedAnswers) != 0 {
				k.SavePollPostAnswers(ctx, test.storedPost.PostID, test.storedAnswers, creator)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if test.expError != nil {
				assert.Equal(t, test.expError, err)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				expectedIndented, _ := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				assert.Equal(t, string(expectedIndented), string(result))
			}
		})
	}
}

func Test_queryPollAnswersAmount(t *testing.T) {
	creator, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")

	tests := []struct {
		name          string
		path          []string
		storedAnswers []uint64
		storedPost    types.Post
		expResult     types.PollAnswersAmountResponse
		expError      sdk.Error
	}{
		{
			name:     "Invalid post id returns error",
			path:     []string{types.QueryAnswersAmount, ""},
			expError: sdk.ErrUnknownRequest("Invalid post id: "),
		},
		{
			name:     "Post with ID not found returns error",
			path:     []string{types.QueryAnswersAmount, "1"},
			expError: sdk.ErrUnknownRequest("Post with ID 1 not found"),
		},
		{
			name:          "Total answers amount returned properly",
			path:          []string{types.QueryAnswersAmount, "1"},
			storedAnswers: []uint64{1, 2},
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			),
			expResult: types.PollAnswersAmountResponse{
				PostID:        types.PostID(1),
				AnswersAmount: sdk.NewInt(2),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			k.SavePost(ctx, test.storedPost)

			if len(test.storedAnswers) != 0 {
				k.SavePollPostAnswers(ctx, test.storedPost.PostID, test.storedAnswers, creator)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if test.expError != nil {
				assert.Equal(t, test.expError, err)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				expectedIndented, _ := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				assert.Equal(t, string(expectedIndented), string(result))
			}
		})
	}

}

func Test_queryPollAnswerVotes(t *testing.T) {
	creator, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")

	tests := []struct {
		name          string
		path          []string
		storedAnswers []uint64
		storedPost    types.Post
		expResult     types.PollAnswerVotesResponse
		expError      sdk.Error
	}{
		{
			name:     "Invalid post id returns error",
			path:     []string{types.QueryAnswerVotes, ""},
			expError: sdk.ErrUnknownRequest("Invalid post id: "),
		},
		{
			name:     "Invalid answer id returns error",
			path:     []string{types.QueryAnswerVotes, "1", ""},
			expError: sdk.ErrUnknownRequest("Unable to parse answer id: "),
		},
		{
			name:     "Post with ID not found returns error",
			path:     []string{types.QueryAnswerVotes, "1", "2"},
			expError: sdk.ErrUnknownRequest("Post with ID 1 not found"),
		},
		{
			name:          "Answer's votes returns correctly",
			path:          []string{types.QueryAnswerVotes, "1", "2"},
			storedAnswers: []uint64{1, 2},
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			),
			expResult: types.PollAnswerVotesResponse{
				PostID:      types.PostID(1),
				AnswerID:    uint64(2),
				VotesAmount: sdk.NewInt(1),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			k.SavePost(ctx, test.storedPost)

			if len(test.storedAnswers) != 0 {
				k.SavePollPostAnswers(ctx, test.storedPost.PostID, test.storedAnswers, creator)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if test.expError != nil {
				assert.Equal(t, test.expError, err)
				assert.Nil(t, result)
			} else {
				assert.Nil(t, err)
				expectedIndented, _ := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				assert.Equal(t, string(expectedIndented), string(result))
			}
		})
	}
}
