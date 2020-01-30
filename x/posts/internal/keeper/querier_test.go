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

	answers := []uint64{uint64(1)}

	tests := []struct {
		name            string
		path            []string
		storedPosts     types.Posts
		storedReactions map[types.PostID]types.Reactions
		storedAnswers   []types.AnswersDetails
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
				k.SavePollPostAnswers(ctx, test.storedPosts[index].PostID, ans)
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
	answers := []uint64{uint64(1)}

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
				k.SavePollPostAnswers(ctx, test.storedPosts[index].PostID, ans)
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
