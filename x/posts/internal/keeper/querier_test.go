package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/tendermint/abci/types"
)

func Test_queryPost(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	assert.NoError(t, err)

	otherCreator, err := sdk.AccAddressFromBech32("cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l")
	assert.NoError(t, err)

	answers := []uint{uint(1)}

	tests := []struct {
		name            string
		path            []string
		storedPosts     types.Posts
		storedReactions map[types.PostID]types.Reactions
		storedAnswers   []types.AnswersDetails
		expResult       types.PostQueryResponse
		expError        error
	}{
		{
			name:     "Invalid ID returns error",
			path:     []string{types.QueryPost, ""},
			expError: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Invalid post id: "),
		},
		{
			name:     "Post not found returns error",
			path:     []string{types.QueryPost, "1"},
			expError: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Post with id 1 not found"),
		},
		{
			name: "Post without likes is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
			},
			storedAnswers: []types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
			path:          []string{types.QueryPost, "1"},
			expResult: types.NewPostResponse(
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				[]types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
				types.Reactions{},
				types.PostIDs{types.PostID(2)},
			),
		},
		{
			name: "Post without children is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
			},
			storedAnswers: []types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
			path:          []string{types.QueryPost, "1"},
			expResult: types.NewPostResponse(
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				[]types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
				types.Reactions{},
				types.PostIDs{},
			),
		},
		{
			name: "Post without medias is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, nil, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, nil, nil),
			},
			storedAnswers: []types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
			storedReactions: map[types.PostID]types.Reactions{
				types.PostID(1): {
					types.NewReaction("Like", creator),
					types.NewReaction("Like", otherCreator),
				},
			},
			path: []string{types.QueryPost, "1"},
			expResult: types.NewPostResponse(
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, nil, testPost.PollData),
				[]types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
				types.Reactions{
					types.NewReaction("Like", creator),
					types.NewReaction("Like", otherCreator),
				},
				types.PostIDs{types.PostID(2)},
			),
		},
		{
			name: "Post without poll and poll answers is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
			},
			storedReactions: map[types.PostID]types.Reactions{
				types.PostID(1): {
					types.NewReaction("Like", creator),
					types.NewReaction("Like", otherCreator),
				},
			},
			path: []string{types.QueryPost, "1"},
			expResult: types.NewPostResponse(
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
				nil,
				types.Reactions{
					types.NewReaction("Like", creator),
					types.NewReaction("Like", otherCreator),
				},
				types.PostIDs{types.PostID(2)},
			),
		},
		{
			name: "Post with all data is returned properly",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
			},
			storedReactions: map[types.PostID]types.Reactions{
				types.PostID(1): {
					types.NewReaction("Like", creator),
					types.NewReaction("Like", otherCreator),
				},
			},
			storedAnswers: []types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
			path:          []string{types.QueryPost, "1"},
			expResult: types.NewPostResponse(
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				[]types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
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

			for index, ans := range test.storedAnswers {
				k.SavePollAnswers(ctx, test.storedPosts[index].PostID, ans)
			}

			for postID, reactions := range test.storedReactions {
				for _, reaction := range reactions {
					err := k.SaveReaction(ctx, postID, reaction)
					assert.NoError(t, err)
				}
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				assert.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				assert.NoError(t, err)
				assert.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.expError.Error(), err.Error())
				assert.Nil(t, result)
			}
		})
	}
}

func Test_queryPosts(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	assert.NoError(t, err)

	answers := []uint{uint(1)}

	tests := []struct {
		name          string
		storedPosts   types.Posts
		storedAnswers []types.AnswersDetails
		params        types.QueryPostsParams
		expResponse   []types.PostQueryResponse
	}{
		{
			name: "Empty params returns all",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
			},
			storedAnswers: []types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
			params:        types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
					[]types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
					types.Reactions{},
					types.PostIDs{types.PostID(2)},
				),
				types.NewPostResponse(
					types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
					nil,
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
			storedAnswers: []types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
			params:        types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, nil, testPost.PollData),
					[]types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
					types.Reactions{},
					types.PostIDs{},
				),
			},
		},
		{
			name: "Empty params returns all posts without poll data and poll answers",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
			},
			params: types.DefaultQueryPostsParams(1, 1),
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
					nil,
					types.Reactions{},
					types.PostIDs{types.PostID(2)},
				),
			},
		},
		{
			name: "Non empty params return proper posts",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
				types.NewPost(types.PostID(2), types.PostID(1), "Child", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, nil),
			},
			storedAnswers: []types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
			params:        types.DefaultQueryPostsParams(1, 1),
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(types.PostID(1), types.PostID(0), "Parent", false, "", map[string]string{}, testPost.Created, creator, testPost.Medias, testPost.PollData),
					[]types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
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

			for index, ans := range test.storedAnswers {
				k.SavePollAnswers(ctx, test.storedPosts[index].PostID, ans)
			}

			querier := keeper.NewQuerier(k)
			request := abci.RequestQuery{Data: k.Cdc.MustMarshalJSON(&test.params)}
			result, err := querier(ctx, []string{types.QueryPosts}, request)
			assert.NoError(t, err)

			expSerialized, err := codec.MarshalJSONIndent(k.Cdc, &test.expResponse)
			assert.NoError(t, err)
			assert.Equal(t, string(expSerialized), string(result))
		})
	}
}

func Test_queryPollAnswers(t *testing.T) {
	creator, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")

	answers := []uint{uint(1)}

	tests := []struct {
		name          string
		path          []string
		storedPosts   types.Posts
		storedAnswers []types.AnswersDetails
		expResult     types.PollAnswersQueryResponse
		expError      error
	}{
		{
			name:     "Invalid post id returns error",
			path:     []string{types.QueryPollAnswers, ""},
			expError: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Invalid post id: "),
		},
		{
			name:     "Post not found returns error",
			path:     []string{types.QueryPollAnswers, "1"},
			expError: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Post with id 1 not found"),
		},
		{
			name: "Post without poll returns error",
			path: []string{types.QueryPollAnswers, "1"},
			storedPosts: types.Posts{
				types.NewPost(
					types.PostID(1),
					types.PostID(0),
					"post with poll",
					false,
					"",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
					nil,
				),
			},
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Post with id 1 has no poll associated"),
		},
		{
			name: "Returns answers details of the post correctly",
			path: []string{types.QueryPollAnswers, "1"},
			storedPosts: types.Posts{
				types.NewPost(
					types.PostID(1),
					types.PostID(0),
					"post with poll",
					false,
					"",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
					testPost.PollData,
				),
			},
			storedAnswers: []types.AnswersDetails{types.NewAnswersDetails(answers, creator)},
			expResult: types.PollAnswersQueryResponse{
				PostID:         types.PostID(1),
				AnswersDetails: types.UsersAnswersDetails{types.NewAnswersDetails(answers, creator)}},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, p := range test.storedPosts {
				k.SavePost(ctx, p)
			}

			for index, ans := range test.storedAnswers {
				k.SavePollAnswers(ctx, test.storedPosts[index].PostID, ans)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				assert.Nil(t, err)
				expectedIndented, _ := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				assert.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.expError.Error(), err.Error())
				assert.Nil(t, result)
			}
		})
	}

}
