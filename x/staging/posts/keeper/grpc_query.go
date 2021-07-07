package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

var _ types.QueryServer = Keeper{}

// Posts implements the Query/Posts gRPC method
func (k Keeper) Posts(goCtx context.Context, req *types.QueryPostsRequest) (*types.QueryPostsResponse, error) {
	var posts []types.Post
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !subspacestypes.IsValidSubspace(req.SubspaceId) {
		return nil, sdkerrors.Wrapf(subspacestypes.ErrInvalidSubspaceID, req.SubspaceId)
	}

	store := ctx.KVStore(k.storeKey)
	postsStore := prefix.NewStore(store, types.SubspacePostsPrefix(req.SubspaceId))

	pageRes, err := query.Paginate(postsStore, req.Pagination, func(key []byte, value []byte) error {
		store := ctx.KVStore(k.storeKey)
		bz := store.Get(types.PostStoreKey(string(value)))

		var post types.Post
		if err := k.cdc.UnmarshalBinaryBare(bz, &post); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		posts = append(posts, post)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPostsResponse{Posts: posts, Pagination: pageRes}, nil
}

// Post implements the Query/Post gRPC method
func (k Keeper) Post(goCtx context.Context, req *types.QueryPostRequest) (*types.QueryPostResponse, error) {
	if !types.IsValidPostID(req.PostId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %s", req.PostId)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	post, found := k.GetPost(ctx, req.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", req.PostId)
	}
	return &types.QueryPostResponse{Post: post}, nil
}

// UserAnswers implements the Query/UserAnswers gRPC method
func (k Keeper) UserAnswers(goCtx context.Context, req *types.QueryUserAnswersRequest) (*types.QueryUserAnswersResponse, error) {
	if !types.IsValidPostID(req.PostId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %s", req.PostId)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	post, found := k.GetPost(ctx, req.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", req.PostId)
	}

	if post.Poll == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s has no poll associated", req.PostId)
	}

	var answers []types.UserAnswer
	store := ctx.KVStore(k.storeKey)
	answersStore := prefix.NewStore(store, types.UserAnswersStoreKey(req.PostId, req.User))
	pageRes, err := query.Paginate(answersStore, req.Pagination, func(key []byte, value []byte) error {
		var answer types.UserAnswer
		if err := k.cdc.UnmarshalBinaryBare(value, &answer); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		answers = append(answers, answer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryUserAnswersResponse{Answers: answers, Pagination: pageRes}, nil
}

// RegisteredReactions implements the Query/RegisteredReactions gRPC method
func (k Keeper) RegisteredReactions(goCtx context.Context, req *types.QueryRegisteredReactionsRequest) (*types.QueryRegisteredReactionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var reactions []types.RegisteredReaction

	store := ctx.KVStore(k.storeKey)
	reactionsStore := prefix.NewStore(store, types.RegisteredReactionsPrefix(req.SubspaceId))

	pageRes, err := query.Paginate(reactionsStore, req.Pagination, func(key []byte, value []byte) error {
		var reaction types.RegisteredReaction
		if err := k.cdc.UnmarshalBinaryBare(value, &reaction); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		reactions = append(reactions, reaction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryRegisteredReactionsResponse{Reactions: reactions, Pagination: pageRes}, nil
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

// PostReactions implements the Query/PostReactions gRPC method
func (k Keeper) PostReactions(goCtx context.Context, req *types.QueryPostReactionsRequest) (*types.QueryPostReactionsResponse, error) {
	if !types.IsValidPostID(req.PostId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id: %s", req.PostId)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_, found := k.GetPost(ctx, req.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", req.PostId)
	}

	var reactions []types.PostReaction

	store := ctx.KVStore(k.storeKey)
	reactionsStore := prefix.NewStore(store, types.PostReactionsPrefix(req.PostId))

	pageRes, err := query.Paginate(reactionsStore, req.Pagination, func(key []byte, value []byte) error {
		var reaction types.PostReaction
		if err := k.cdc.UnmarshalBinaryBare(value, &reaction); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		reactions = append(reactions, reaction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPostReactionsResponse{Reactions: reactions, Pagination: pageRes}, nil
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
	pageRes, err := query.Paginate(commetsStore, req.Pagination, func(key []byte, value []byte) error {
		// it assumes that the comment must exist
		comment, _ := k.GetPost(ctx, string(value))
		comments = append(comments, comment)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPostCommentsResponse{Comments: comments, Pagination: pageRes}, nil
}
