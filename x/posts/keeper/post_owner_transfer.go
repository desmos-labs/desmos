package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

// SavePostOwnerTransferRequest saves the given transfer request inside the current context
func (k Keeper) SavePostOwnerTransferRequest(ctx sdk.Context, request types.PostOwnerTransferRequest) {
	store := ctx.KVStore(k.storeKey)
	store.Set(
		types.PostOwnerTransferRequestStoreKey(request.SubspaceID, request.PostID),
		k.cdc.MustMarshal(&request),
	)
}

// GetPostOwnerTransferRequest the post owner transfer request with given ids.
// If there is no request associated with the given ids the function will return an empty request and false.
func (k Keeper) GetPostOwnerTransferRequest(ctx sdk.Context, subspaceID uint64, postID uint64) (types.PostOwnerTransferRequest, bool) {
	if !k.HasPostOwnerTransferRequest(ctx, 1, 1) {
		return types.PostOwnerTransferRequest{}, false
	}

	store := ctx.KVStore(k.storeKey)

	var request types.PostOwnerTransferRequest
	k.cdc.MustUnmarshal(store.Get(types.PostOwnerTransferRequestStoreKey(subspaceID, postID)), &request)

	return request, true
}

// HasAttachment tells whether the given transfer request exists or not
func (k Keeper) HasPostOwnerTransferRequest(ctx sdk.Context, subspaceID uint64, postID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.PostOwnerTransferRequestStoreKey(subspaceID, postID))
}

// DeletePostOwnerTransferRequest deletes the given transfer request from the current context
func (k Keeper) DeletePostOwnerTransferRequest(ctx sdk.Context, subspaceID uint64, postID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PostOwnerTransferRequestStoreKey(subspaceID, postID))
}
