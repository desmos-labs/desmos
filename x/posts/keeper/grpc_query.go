package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/desmos-labs/desmos/v7/x/posts/types"
)

var _ types.QueryServer = &Keeper{}

// SubspacePosts implements the QuerySubspacePosts gRPC method
func (k Keeper) SubspacePosts(ctx context.Context, request *types.QuerySubspacePostsRequest) (*types.QuerySubspacePostsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	store := sdkCtx.KVStore(k.storeKey)
	postsStore := prefix.NewStore(store, types.SubspacePostsPrefix(request.SubspaceId))

	var posts []types.Post
	pageRes, err := query.Paginate(postsStore, request.Pagination, func(key []byte, value []byte) error {
		var post types.Post
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		posts = append(posts, post)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubspacePostsResponse{
		Posts:      posts,
		Pagination: pageRes,
	}, nil
}

// SectionPosts implements the QuerySectionPosts gRPC method
func (k Keeper) SectionPosts(ctx context.Context, request *types.QuerySectionPostsRequest) (*types.QuerySectionPostsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	store := sdkCtx.KVStore(k.storeKey)
	postsPrefix := types.SectionPostsPrefix(request.SubspaceId, request.SectionId)
	sectionsPostsStore := prefix.NewStore(store, postsPrefix)

	var posts []types.Post
	pageRes, err := query.Paginate(sectionsPostsStore, request.Pagination, func(key []byte, value []byte) error {
		subspaceID, _, postID := types.SplitPostSectionStoreKey(append(postsPrefix, key...))
		post, found := k.GetPost(sdkCtx, subspaceID, postID)
		if !found {
			return fmt.Errorf("post not found: subspace id %d, post id %d", subspaceID, postID)
		}

		posts = append(posts, post)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySectionPostsResponse{
		Posts:      posts,
		Pagination: pageRes,
	}, nil
}

// Post implements the QueryPost gRPC method
func (k Keeper) Post(ctx context.Context, request *types.QueryPostRequest) (*types.QueryPostResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}
	if request.PostId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	post, found := k.GetPost(sdkCtx, request.SubspaceId, request.PostId)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "post with id %d not found", request.PostId)
	}

	return &types.QueryPostResponse{
		Post: post,
	}, nil
}

// PostAttachments implements the QueryPostAttachments gRPC method
func (k Keeper) PostAttachments(ctx context.Context, request *types.QueryPostAttachmentsRequest) (*types.QueryPostAttachmentsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}
	if request.PostId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	store := sdkCtx.KVStore(k.storeKey)
	attachmentsStore := prefix.NewStore(store, types.PostAttachmentsPrefix(request.SubspaceId, request.PostId))

	var attachments []types.Attachment
	pageRes, err := query.Paginate(attachmentsStore, request.Pagination, func(key []byte, value []byte) error {
		var attachment types.Attachment
		if err := k.cdc.Unmarshal(value, &attachment); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		attachments = append(attachments, attachment)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPostAttachmentsResponse{
		Attachments: attachments,
		Pagination:  pageRes,
	}, nil
}

// PollAnswers implements the QueryPollAnswers gRPC method
func (k Keeper) PollAnswers(ctx context.Context, request *types.QueryPollAnswersRequest) (*types.QueryPollAnswersResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}
	if request.PostId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id")
	}
	if request.PollId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid poll id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	store := sdkCtx.KVStore(k.storeKey)
	answersPrefix := types.PollAnswersPrefix(request.SubspaceId, request.PostId, request.PollId)
	if request.User != "" {
		answersPrefix = types.PollAnswerStoreKey(request.SubspaceId, request.PostId, request.PollId, request.User)
	}
	answersStore := prefix.NewStore(store, answersPrefix)

	var answers []types.UserAnswer
	pageRes, err := query.Paginate(answersStore, request.Pagination, func(key []byte, value []byte) error {
		var answer types.UserAnswer
		if err := k.cdc.Unmarshal(value, &answer); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		answers = append(answers, answer)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPollAnswersResponse{
		Answers:    answers,
		Pagination: pageRes,
	}, nil
}

// Params implements the QueryParams gRPC method
func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// IncomingPostOwnerTransferRequests implements the QueryIncomingPostOwnerTransferRequests gRPC method
func (k Keeper) IncomingPostOwnerTransferRequests(ctx context.Context, request *types.QueryIncomingPostOwnerTransferRequestsRequest) (*types.QueryIncomingPostOwnerTransferRequestsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	store := sdkCtx.KVStore(k.storeKey)
	transferRequestsPrefix := types.SubspacePostOwnerTransferRequestPrefix(request.SubspaceId)
	transferRequestsStore := prefix.NewStore(store, transferRequestsPrefix)

	var transferRequests []types.PostOwnerTransferRequest
	pageRes, err := query.FilteredPaginate(transferRequestsStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var transferRequest types.PostOwnerTransferRequest
		if err := k.cdc.Unmarshal(value, &transferRequest); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}

		// Filter out the request whose receiver does not match the given receiver
		if request.Receiver != "" && request.Receiver != transferRequest.Receiver {
			return false, nil
		}

		if accumulate {
			transferRequests = append(transferRequests, transferRequest)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryIncomingPostOwnerTransferRequestsResponse{
		Requests:   transferRequests,
		Pagination: pageRes,
	}, nil
}
