package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// SaveDTagTransferRequest save the given request associating it to the request recipient.
// It returns an error if the same request already exists.
func (k Keeper) SaveDTagTransferRequest(ctx sdk.Context, request types.DTagTransferRequest) error {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(request.Sender, request.Receiver)
	if store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the transfer request from %s to %s has already been made",
			request.Sender, request.Receiver)
	}

	store.Set(key, k.cdc.MustMarshalBinaryBare(&request))
	k.Logger(ctx).Info("DTag transfer request", "sender", request.Sender, "receiver", request.Receiver)
	return nil
}

// GetDTagTransferRequest retries the DTag transfer request made from the specified sender to the given receiver.
// If the request was not found, returns false instead.
func (k Keeper) GetDTagTransferRequest(ctx sdk.Context, sender, receiver string) (types.DTagTransferRequest, bool, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(sender, receiver)
	if !store.Has(key) {
		return types.DTagTransferRequest{}, false, nil
	}

	var request types.DTagTransferRequest
	err := k.cdc.UnmarshalBinaryBare(store.Get(key), &request)
	if err != nil {
		return types.DTagTransferRequest{}, false, err
	}

	return request, true, nil
}

// GetDTagTransferRequests returns all the requests inside the given context
func (k Keeper) GetDTagTransferRequests(ctx sdk.Context) (requests []types.DTagTransferRequest) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DTagTransferRequestPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		request := types.MustUnmarshalDTagTransferRequest(k.cdc, iterator.Value())
		requests = append(requests, request)
	}

	return requests
}

// DeleteDTagTransferRequest deletes the transfer request made from the sender towards the recipient
func (k Keeper) DeleteDTagTransferRequest(ctx sdk.Context, sender, recipient string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(sender, recipient)
	if !store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "request from %s to %s not found", sender, recipient)
	}

	store.Delete(key)
	return nil
}

// DeleteAllDTagTransferRequests delete all the requests made to the given user
func (k Keeper) DeleteAllDTagTransferRequests(ctx sdk.Context, receiver string) {
	var requests []types.DTagTransferRequest
	k.IterateUserIncomingDTagTransferRequests(ctx, receiver, func(index int64, request types.DTagTransferRequest) (stop bool) {
		requests = append(requests, request)
		return false
	})

	store := ctx.KVStore(k.storeKey)
	for _, request := range requests {
		store.Delete(types.DTagTransferRequestStoreKey(request.Sender, request.Receiver))
	}
}
