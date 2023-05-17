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
