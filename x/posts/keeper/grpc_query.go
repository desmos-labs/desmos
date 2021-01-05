package keeper

import (
	"context"
	"sort"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/desmos-labs/desmos/x/posts/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) getPostResponse(ctx sdk.Context, post types.Post) types.QueryPostResponse {
	// Get the reactions
	postReactions := k.GetPostReactions(ctx, post.PostID)
	if postReactions == nil {
		postReactions = []types.PostReaction{}
	}

	// Get the children
	childrenIDs := k.GetPostChildrenIDs(ctx, post.PostID)
	if childrenIDs == nil {
		childrenIDs = []string{}
	}

	//Get the poll answers if poll exist
	var answers []types.UserAnswer
	if post.PollData != nil {
		answers = k.GetPollAnswers(ctx, post.PostID)
	}

	// Crete the response object
	return types.QueryPostResponse{
		Post:        post,
		PollAnswers: answers,
		Reactions:   postReactions,
		Children:    childrenIDs,
	}
}

func (k Keeper) Posts(goCtx context.Context, req *types.QueryPostsRequest) (*types.QueryPostsResponse, error) {
	var filteredPosts []types.QueryPostResponse
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	postsStore := prefix.NewStore(store, types.PostStorePrefix)

	pageRes, err := query.FilteredPaginate(postsStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var post types.Post
		if err := k.cdc.UnmarshalBinaryBare(value, &post); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}

		matchParentID, matchCreationTime, matchSubspace, matchCreator, matchHashtags := true, true, true, true, true

		// match parent id if valid
		if types.IsValidPostID(req.ParentID) {
			matchParentID = req.ParentID == post.ParentID
		}

		// match creation time if valid height
		if req.CreationTime != nil {
			matchCreationTime = req.CreationTime.Equal(post.Created)
		}

		// match subspace if provided
		if req.Subspace != "" {
			matchSubspace = req.Subspace == post.Subspace
		}

		// match creator address (if supplied)
		if req.Creator != "" {
			matchCreator = req.Creator == post.Creator
		}

		// match hashtags if provided
		if req.Hashtags != nil {
			postHashtags := post.GetPostHashtags()
			matchHashtags = len(postHashtags) == len(req.Hashtags)
			sort.Strings(postHashtags)
			sort.Strings(req.Hashtags)
			for index := 0; index < len(req.Hashtags) && matchHashtags; index++ {
				matchHashtags = postHashtags[index] == req.Hashtags[index]
			}
		}

		if matchParentID && matchCreationTime && matchSubspace && matchCreator && matchHashtags {
			if accumulate {
				filteredPosts = append(filteredPosts, k.getPostResponse(ctx, post))
			}

			return true, nil
		}

		return false, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPostsResponse{Posts: filteredPosts, Pagination: pageRes}, nil
}

func (k Keeper) Post(goCtx context.Context, req *types.QueryPostRequest) (*types.QueryPostResponse, error) {
	if !types.IsValidPostID(req.PostId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %s", req.PostId)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	post, found := k.GetPost(ctx, req.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", req.PostId)
	}

	response := k.getPostResponse(ctx, post)
	return &response, nil
}

func (k Keeper) PollAnswers(goCtx context.Context, req *types.QueryPollAnswersRequest) (*types.QueryPollAnswersResponse, error) {
	if !types.IsValidPostID(req.PostId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %s", req.PostId)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	post, found := k.GetPost(ctx, req.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", req.PostId)
	}

	if post.PollData == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s has no poll associated", req.PostId)
	}

	pollAnswers := k.GetPollAnswers(ctx, req.PostId)
	return &types.QueryPollAnswersResponse{PostId: req.PostId, Answers: pollAnswers}, nil
}

func (k Keeper) RegisteredReactions(goCtx context.Context, _ *types.QueryRegisteredReactionsRequest) (*types.QueryRegisteredReactionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	reactions := k.GetRegisteredReactions(ctx)
	return &types.QueryRegisteredReactionsResponse{RegisteredReactions: reactions}, nil
}

func (k Keeper) Params(goCtx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}
