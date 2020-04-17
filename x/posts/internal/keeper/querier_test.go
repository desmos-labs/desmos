package keeper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func Test_queryPost(t *testing.T) {
	creationDate := time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone)
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	computedID := types.ComputeID(creationDate, creator, subspace)
	stringID := computedID.String()

	otherCreator, err := sdk.AccAddressFromBech32("cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l")
	require.NoError(t, err)

	computedID2 := types.ComputeID(creationDate, otherCreator, subspace)

	answers := []types.AnswerID{types.AnswerID(1)}

	reaction := types.NewReaction(testPostOwner, ":like:", "https://smile.jpg",
		"")

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
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
				types.NewPost(computedID2, computedID, "Child", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
			},
			storedAnswers:      []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			registeredReaction: nil,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
				[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
				types.PostReactions{},
				types.PostIDs{computedID2},
			),
		},
		{
			name: "Post without children is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
			},
			storedAnswers:      []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			registeredReaction: nil,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
				[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
				types.PostReactions{},
				types.PostIDs{},
			),
		},
		{
			name: "Post without medias is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithPollData(*testPost.PollData),
				types.NewPost(computedID2, computedID, "Child", false, "", map[string]string{}, testPost.Created, creator),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			storedReactions: map[string]types.PostReactions{
				stringID: {
					types.NewPostReaction(":like:", creator),
					types.NewPostReaction(":like:", otherCreator),
				},
			},
			registeredReaction: &reaction,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithPollData(*testPost.PollData),
				[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
				types.PostReactions{
					types.NewPostReaction(":like:", creator),
					types.NewPostReaction(":like:", otherCreator),
				},
				types.PostIDs{computedID2},
			),
		},
		{
			name: "Post without poll and poll answers is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
				types.NewPost(computedID2, computedID, "Child", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
			},
			storedReactions: map[string]types.PostReactions{
				stringID: {
					types.NewPostReaction(":like:", creator),
					types.NewPostReaction(":like:", otherCreator),
				},
			},
			registeredReaction: &reaction,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
				nil,
				types.PostReactions{
					types.NewPostReaction(":like:", creator),
					types.NewPostReaction(":like:", otherCreator),
				},
				types.PostIDs{computedID2},
			),
		},
		{
			name: "Post with all data is returned properly",
			storedPosts: types.Posts{
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
				types.NewPost(computedID2, computedID, "Child", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
			},
			storedReactions: map[string]types.PostReactions{
				stringID: {
					types.NewPostReaction(":like:", creator),
					types.NewPostReaction(":like:", otherCreator),
				},
			},
			storedAnswers:      []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			registeredReaction: &reaction,
			path:               []string{types.QueryPost, stringID},
			expResult: types.NewPostResponse(
				types.NewPost(computedID, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
				[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
				types.PostReactions{
					types.NewPostReaction(":like:", creator),
					types.NewPostReaction(":like:", otherCreator),
				},
				types.PostIDs{computedID2},
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

			if test.registeredReaction != nil {
				k.RegisterReaction(ctx, *test.registeredReaction)
			}

			for index, ans := range test.storedAnswers {
				k.SavePollAnswers(ctx, test.storedPosts[index].PostID, ans)
			}

			for postID, reactions := range test.storedReactions {
				for _, reaction := range reactions {
					err = k.SavePostReaction(ctx, types.PostID(postID), reaction)
					require.NoError(t, err)
				}
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				require.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				require.NoError(t, err)
				require.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expError.Error(), err.Error())
				require.Nil(t, result)
			}
		})
	}
}

func Test_queryPosts(t *testing.T) {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

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
				types.NewPost(id, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
				types.NewPost(id2, id, "Child", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			params:        types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(id, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
					[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
					types.PostReactions{},
					types.PostIDs{id2},
				),
				types.NewPostResponse(
					types.NewPost(id2, id, "Child", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
					nil,
					types.PostReactions{},
					types.PostIDs{},
				),
			},
		},
		{
			name: "Empty params returns all posts without medias",
			storedPosts: types.Posts{
				types.NewPost(id2, id, "Child", false, "", map[string]string{}, testPost.Created, creator).WithPollData(*testPost.PollData),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			params:        types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(id2, id, "Child", false, "", map[string]string{}, testPost.Created, creator).WithPollData(*testPost.PollData),
					[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
					types.PostReactions{},
					types.PostIDs{},
				),
			},
		},
		{
			name: "Empty params returns all posts without poll data and poll answers",
			storedPosts: types.Posts{
				types.NewPost(id, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
				types.NewPost(id2, id, "Child", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
			},
			params: types.DefaultQueryPostsParams(1, 1),
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(id, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
					nil,
					types.PostReactions{},
					types.PostIDs{id2},
				),
			},
		},
		{
			name: "Non empty params return proper posts",
			storedPosts: types.Posts{
				types.NewPost(id, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
				types.NewPost(id2, id, "Child", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			params:        types.DefaultQueryPostsParams(1, 1),
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.NewPost(id, "", "Parent", false, "", map[string]string{}, testPost.Created, creator).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
					[]types.UserAnswer{types.NewUserAnswer(answers, creator)},
					types.PostReactions{},
					types.PostIDs{id2},
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
			require.NoError(t, err)

			expSerialized, err := codec.MarshalJSONIndent(k.Cdc, &test.expResponse)
			require.NoError(t, err)
			require.Equal(t, string(expSerialized), string(result))
		})
	}
}

func Test_queryPollAnswers(t *testing.T) {
	creationDate := time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone)
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)
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
					testPost.Created,
					testPost.Creator,
				).WithMedias(testPost.Medias),
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
					testPost.Created,
					testPost.Creator,
				).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
			},
			storedAnswers: []types.UserAnswer{types.NewUserAnswer(answers, creator)},
			expResult: types.PollAnswersQueryResponse{
				PostID:         computedID,
				AnswersDetails: types.UserAnswers{types.NewUserAnswer(answers, creator)}},
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
				require.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				require.NoError(t, err)

				require.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expError.Error(), err.Error())
				require.Nil(t, result)
			}
		})
	}

}

func Test_queryRegisteredReactions(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

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
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, r := range test.storedReactions {
				k.RegisterReaction(ctx, r)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				require.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				require.NoError(t, err)

				require.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expError.Error(), err.Error())
				require.Nil(t, result)
			}
		})
	}
}
