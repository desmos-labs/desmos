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

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) getPostResponse(ctx sdk.Context, post types.Post) types.QueryPostResponse {
	// Get the reactions
	postReactions := k.GetPostReactions(ctx, post.PostID)
	if postReactions == nil {
		postReactions = []types.PostReaction{}
	}

	// Get the children
	childrenIDs := k.GetPostCommentIDs(ctx, post.PostID)
	if childrenIDs == nil {
		childrenIDs = []string{}
	}

	//Get the user answers if poll exist
	var answers []types.UserAnswer
	if post.PollData != nil {
		answers = k.GetUserAnswersByPost(ctx, post.PostID)
	}

	// Crete the response object
	return types.QueryPostResponse{
		Post:        post,
		UserAnswers: answers,
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
		if types.IsValidPostID(req.ParentId) {
			matchParentID = req.ParentId == post.ParentID
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

func (k Keeper) UserAnswers(goCtx context.Context, req *types.QueryUserAnswersRequest) (*types.QueryUserAnswersResponse, error) {
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

	var answers []types.UserAnswer
	store := ctx.KVStore(k.storeKey)
	answersStore := prefix.NewStore(store, types.UserAnswersStoreKey(req.PostId, req.User))
	pageRes, err := query.FilteredPaginate(answersStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		answer := types.MustUnmarshalUserAnswer(k.cdc, value)
		if accumulate {
			answers = append(answers, answer)
		}
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryUserAnswersResponse{Answers: answers, Pagination: pageRes}, nil
}

func (k Keeper) RegisteredReactions(goCtx context.Context, req *types.QueryRegisteredReactionsRequest) (*types.QueryRegisteredReactionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var reactions []types.RegisteredReaction

	store := ctx.KVStore(k.storeKey)
	reactionsStore := prefix.NewStore(store, types.RegisteredReactionsPrefix(req.Subspace))

	pageRes, err := query.FilteredPaginate(reactionsStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var reaction types.RegisteredReaction
		k.cdc.MustUnmarshalBinaryBare(value, &reaction)
		if accumulate {
			reactions = append(reactions, reaction)
		}
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryRegisteredReactionsResponse{RegisteredReactions: reactions, Pagination: pageRes}, nil
}

// Reports implements the Query/Reports gRPC method
func (k Keeper) Reports(
	ctx context.Context, request *types.QueryReportsRequest,
) (*types.QueryReportsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reports := k.GetPostReports(sdkCtx, request.PostId)
	return &types.QueryReportsResponse{Reports: reports}, nil
}

func (k Keeper) Params(goCtx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// PostComments implements the Query/PostComments gRPC method
func (k Keeper) PostComments(
	goCtx context.Context, req *types.QueryPostCommentsRequest,
) (*types.QueryPostCommentsResponse, error) {
	if !types.IsValidPostID(req.PostId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %s", req.PostId)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetPost(ctx, req.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", req.PostId)
	}

	store := ctx.KVStore(k.storeKey)
	commetsStore := prefix.NewStore(store, types.PostCommentsPrefix(req.PostId))

	var comments []types.Post
	pageRes, err := query.FilteredPaginate(commetsStore, req.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		// it assumes that the comment must exist
		comment, _ := k.GetPost(ctx, string(value))
		if accumulate {
			comments = append(comments, comment)
		}
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPostCommentsResponse{Comments: comments, Pagination: pageRes}, nil
}
