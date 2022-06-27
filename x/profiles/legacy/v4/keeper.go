package v4

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec
}

func NewKeeper(storeKey sdk.StoreKey, cdc codec.BinaryCodec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

func (k Keeper) IterateDTags(ctx sdk.Context, fn func(index int64, dTag string, value []byte) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	dTagsStore := prefix.NewStore(store, DTagPrefix)
	iterator := dTagsStore.Iterator(nil, nil)
	defer iterator.Close()

	var stop = false
	var index = int64(0)
	for ; iterator.Valid() && !stop; iterator.Next() {
		stop = fn(index, string(iterator.Key()), iterator.Value())
		index++
	}
}

func (k Keeper) IterateDTagTransferRequests(ctx sdk.Context, fn func(index int64, request types.DTagTransferRequest) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	requestsStore := prefix.NewStore(store, DTagTransferRequestPrefix)
	iterator := requestsStore.Iterator(nil, nil)
	defer iterator.Close()

	var stop = false
	var index = int64(0)
	for ; iterator.Valid() && !stop; iterator.Next() {
		var request types.DTagTransferRequest
		err := k.cdc.Unmarshal(iterator.Value(), &request)
		if err != nil {
			panic(err)
		}
		stop = fn(index, request)
		index++
	}
}

func (k Keeper) IterateRelationships(ctx sdk.Context, fn func(index int64, relationship Relationship) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	relationshipsStore := prefix.NewStore(store, RelationshipsStorePrefix)
	iterator := relationshipsStore.Iterator(nil, nil)
	defer iterator.Close()

	var stop = false
	var index = int64(0)
	for ; iterator.Valid() && !stop; iterator.Next() {
		var relationship Relationship
		err := k.cdc.Unmarshal(iterator.Value(), &relationship)
		if err != nil {
			panic(err)
		}
		stop = fn(index, relationship)
		index++
	}
}

func (k Keeper) GetRelationships(ctx sdk.Context) []Relationship {
	var values []Relationship
	k.IterateRelationships(ctx, func(_ int64, relationship Relationship) (stop bool) {
		values = append(values, relationship)
		return false
	})
	return values
}

func (k Keeper) DeleteRelationships(ctx sdk.Context) {
	var keys [][]byte
	k.IterateRelationships(ctx, func(_ int64, relationship Relationship) (stop bool) {
		keys = append(keys, RelationshipsStoreKey(relationship.Creator, relationship.SubspaceID, relationship.Recipient))
		return false
	})

	store := ctx.KVStore(k.storeKey)
	for _, key := range keys {
		store.Delete(key)
	}
}

func (k Keeper) IterateBlocks(ctx sdk.Context, fn func(index int64, relationship UserBlock) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	userBlocksStore := prefix.NewStore(store, UsersBlocksStorePrefix)
	iterator := userBlocksStore.Iterator(nil, nil)
	defer iterator.Close()

	var stop = false
	var index = int64(0)
	for ; iterator.Valid() && !stop; iterator.Next() {
		var block UserBlock
		err := k.cdc.Unmarshal(iterator.Value(), &block)
		if err != nil {
			panic(err)
		}
		stop = fn(index, block)
		index++
	}
}

func (k Keeper) GetBlocks(ctx sdk.Context) []UserBlock {
	var values []UserBlock
	k.IterateBlocks(ctx, func(_ int64, block UserBlock) (stop bool) {
		values = append(values, block)
		return false
	})

	return values
}

func (k Keeper) DeleteBlocks(ctx sdk.Context) {
	var keys [][]byte
	k.IterateBlocks(ctx, func(_ int64, block UserBlock) (stop bool) {
		keys = append(keys, UserBlockStoreKey(block.Blocker, block.SubspaceID, block.Blocked))
		return false
	})

	store := ctx.KVStore(k.storeKey)
	for _, key := range keys {
		store.Delete(key)
	}
}

func (k Keeper) IterateChainLinks(ctx sdk.Context, fn func(index int64, chainLink ChainLink) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	chainLinksStore := prefix.NewStore(store, ChainLinksPrefix)
	iterator := chainLinksStore.Iterator(nil, nil)
	defer iterator.Close()

	var stop = false
	var index = int64(0)
	for ; iterator.Valid() && !stop; iterator.Next() {
		var chainLink ChainLink
		err := k.cdc.Unmarshal(iterator.Value(), &chainLink)
		if err != nil {
			panic(err)
		}
		stop = fn(index, chainLink)
		index++
	}
}

func (k Keeper) IterateApplicationLinks(ctx sdk.Context, fn func(index int64, applicationLink types.ApplicationLink) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	applicationLinksStore := prefix.NewStore(store, UserApplicationLinkPrefix)
	iterator := applicationLinksStore.Iterator(nil, nil)
	defer iterator.Close()

	var stop = false
	var index = int64(0)
	for ; iterator.Valid() && !stop; iterator.Next() {
		var applicationLink types.ApplicationLink
		err := k.cdc.Unmarshal(iterator.Value(), &applicationLink)
		if err != nil {
			panic(err)
		}
		stop = fn(index, applicationLink)
		index++
	}
}

func (k Keeper) IterateApplicationLinkClientIDKeys(ctx sdk.Context, fn func(index int64, key []byte, value []byte) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	clientIDsStore := prefix.NewStore(store, ApplicationLinkClientIDPrefix)
	iterator := clientIDsStore.Iterator(nil, nil)
	defer iterator.Close()

	var stop = false
	var index = int64(0)
	for ; iterator.Valid() && !stop; iterator.Next() {
		stop = fn(index, append(ApplicationLinkClientIDPrefix, iterator.Key()...), iterator.Value())
		index++
	}
}
