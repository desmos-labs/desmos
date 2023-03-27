package keeper

import (
	"context"

	errors "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

var _ types.QueryServer = &Keeper{}

// Reactions implements the QueryReactions gRPC method
func (k Keeper) Reactions(ctx context.Context, request *types.QueryReactionsRequest) (*types.QueryReactionsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	if request.PostId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	storePrefix := types.PostReactionsPrefix(request.SubspaceId, request.PostId)
	reactionsStore := prefix.NewStore(store, storePrefix)

	var reactions []types.Reaction
	pageRes, err := query.FilteredPaginate(reactionsStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var reaction types.Reaction
		if err := k.cdc.Unmarshal(value, &reaction); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}

		// Skip all the reactions from a different author if the user is specified
		if request.User != "" && request.User != reaction.Author {
			return false, nil
		}

		if accumulate {
			reactions = append(reactions, reaction)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryReactionsResponse{Reactions: reactions, Pagination: pageRes}, nil
}

// Reaction implements the QueryReaction gRPC method
func (k Keeper) Reaction(ctx context.Context, request *types.QueryReactionRequest) (*types.QueryReactionResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	if request.PostId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid post id")
	}

	if request.ReactionId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid reaction id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reaction, found := k.GetReaction(sdkCtx, request.SubspaceId, request.PostId, request.ReactionId)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "reaction with id %d not found inside subspace %d", request.ReactionId, request.SubspaceId)
	}

	return &types.QueryReactionResponse{Reaction: reaction}, nil
}

// RegisteredReactions implements the QueryRegisteredReactions gRPC method
func (k Keeper) RegisteredReactions(ctx context.Context, request *types.QueryRegisteredReactionsRequest) (*types.QueryRegisteredReactionsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	storePrefix := types.SubspaceRegisteredReactionsPrefix(request.SubspaceId)
	reactionsStore := prefix.NewStore(store, storePrefix)

	var reactions []types.RegisteredReaction
	pageRes, err := query.Paginate(reactionsStore, request.Pagination, func(key []byte, value []byte) error {
		var reaction types.RegisteredReaction
		if err := k.cdc.Unmarshal(value, &reaction); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		reactions = append(reactions, reaction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRegisteredReactionsResponse{RegisteredReactions: reactions, Pagination: pageRes}, nil
}

// RegisteredReaction implements the QueryRegisteredReaction gRPC method
func (k Keeper) RegisteredReaction(ctx context.Context, request *types.QueryRegisteredReactionRequest) (*types.QueryRegisteredReactionResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	if request.ReactionId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid registered reaction id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reaction, found := k.GetRegisteredReaction(sdkCtx, request.SubspaceId, request.ReactionId)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrNotFound, "registered reaction with id %d not found inside subspace %d", request.ReactionId, request.SubspaceId)
	}

	return &types.QueryRegisteredReactionResponse{RegisteredReaction: reaction}, nil
}

// ReactionsParams implements the QueryReactionsParams gRPC method
func (k Keeper) ReactionsParams(ctx context.Context, request *types.QueryReactionsParamsRequest) (*types.QueryReactionsParamsResponse, error) {
	if request.SubspaceId == 0 {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid subspace id")
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	params, err := k.GetSubspaceReactionsParams(sdkCtx, request.SubspaceId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryReactionsParamsResponse{Params: params}, nil
}
