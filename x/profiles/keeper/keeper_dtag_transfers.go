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
	key := types.DTagTransferRequestStoreKey(request.Receiver)

	var requests types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &requests)
	for _, req := range requests.Requests {
		if req.Sender == request.Sender && req.Receiver == request.Receiver {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the transfer request from %s to %s has already been made",
				request.Sender, request.Receiver)
		}
	}

	requests = types.NewDTagTransferRequests(append(requests.Requests, request))
	store.Set(key, k.cdc.MustMarshalBinaryBare(&requests))
	return nil
}

// GetUserIncomingDTagTransferRequests returns all the request made to the given user inside the current context.
func (k Keeper) GetUserIncomingDTagTransferRequests(ctx sdk.Context, user string) []types.DTagTransferRequest {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(user)

	var requests types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &requests)
	return requests.Requests
}

// GetDTagTransferRequests returns all the requests inside the given context
func (k Keeper) GetDTagTransferRequests(ctx sdk.Context) (requests []types.DTagTransferRequest) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DTagTransferRequestsPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var userRequests types.DTagTransferRequests
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &userRequests)
		requests = append(requests, userRequests.Requests...)
	}

	return requests
}

// DeleteAllDTagTransferRequests delete all the requests made to the given user
func (k Keeper) DeleteAllDTagTransferRequests(ctx sdk.Context, user string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DTagTransferRequestStoreKey(user))
}

// DeleteDTagTransferRequest deletes the transfer requests made from the sender towards the recipient
func (k Keeper) DeleteDTagTransferRequest(ctx sdk.Context, sender, recipient string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(recipient)

	var wrapped types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &wrapped)

	for index, request := range wrapped.Requests {
		if request.Sender == sender {
			requests := append(wrapped.Requests[:index], wrapped.Requests[index+1:]...)
			if len(requests) == 0 {
				store.Delete(key)
			} else {
				store.Set(key, k.cdc.MustMarshalBinaryBare(&types.DTagTransferRequests{Requests: requests}))
			}
			return nil
		}
	}

	return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "request from %s to %s not found", sender, recipient)
}
