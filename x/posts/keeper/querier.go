package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {

		case types.QueryPost:
			return queryPost(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case types.QueryPosts:
			return queryPosts(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryPollAnswers:
			return queryPollAnswers(ctx, path[1:], req, keeper, legacyQuerierCdc)

		case types.QueryRegisteredReactions:
			return queryRegisteredReactions(ctx, req, keeper, legacyQuerierCdc)

		case types.QueryParams:
			return queryParams(ctx, req, keeper, legacyQuerierCdc)

		default:
			return nil, fmt.Errorf("unknown post query endpoint")
		}
	}
}

// getPostResponse allows to get a PostQueryResponse from the given post retrieving the other information
// using the given Context and Keeper.
func getPostResponse(ctx sdk.Context, keeper Keeper, post types.Post) types.PostQueryResponse {
	// Get the reactions
	postReactions := keeper.GetPostReactions(ctx, post.PostID)
	if postReactions == nil {
		postReactions = []types.PostReaction{}
	}

	// Get the children
	childrenIDs := keeper.GetPostChildrenIDs(ctx, post.PostID)
	if childrenIDs == nil {
		childrenIDs = []string{}
	}

	//Get the poll answers if poll exist
	var answers []types.UserAnswer
	if post.PollData != nil {
		answers = keeper.GetPollAnswers(ctx, post.PostID)
	}

	// Crete the response object
	return types.NewPostResponse(post, answers, postReactions, childrenIDs)
}

// queryPost handles the request to get a post having a specific id
func queryPost(
	ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	id := path[0]
	if !types.IsValidPostID(id) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %s", id)
	}

	post, found := keeper.GetPost(ctx, id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", id)
	}

	postResponse := getPostResponse(ctx, keeper, post)
	bz, err2 := codec.MarshalJSONIndent(legacyQuerierCdc, &postResponse)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// queryPosts handles the request of listing all the posts that satisfy a specific filter
func queryPosts(
	ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	var params types.QueryPostsParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	posts := keeper.GetPostsFiltered(ctx, params)

	postResponses := make([]types.PostQueryResponse, len(posts))
	for index, post := range posts {
		postResponses[index] = getPostResponse(ctx, keeper, post)
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &postResponses)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

//queryPollAnswers handles the request to get poll answers related to a post with given id
func queryPollAnswers(
	ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	id := path[0]
	if !types.IsValidPostID(id) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %s", id)
	}

	post, found := keeper.GetPost(ctx, id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", id)
	}

	if post.PollData == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s has no poll associated", id)
	}

	pollAnswers := keeper.GetPollAnswers(ctx, id)

	pollAnswersResponse := types.QueryPollAnswersResponse{
		PostId:  id,
		Answers: pollAnswers,
	}
	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &pollAnswersResponse)

	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryRegisteredReactions(
	ctx sdk.Context, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	reactions := keeper.GetRegisteredReactions(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &reactions)

	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryParams(
	ctx sdk.Context, _ abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino,
) ([]byte, error) {
	params := keeper.GetParams(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, &params)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
