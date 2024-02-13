package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v1 "github.com/desmos-labs/desmos/v7/x/relationships/legacy/v1"
	"github.com/desmos-labs/desmos/v7/x/relationships/types"
)

// MigrateStore performs in-place store migrations from v1 to v2.
// The migration includes:
//
// - migrate all relationships keys to the new ones
// - migrate all user blocks keys to the new ones
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	err := migrateRelationships(store, cdc)
	if err != nil {
		return err
	}

	err = migrateUserBlocks(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// migrateRelationships migrates all the relationships from using the old key to using the new key
func migrateRelationships(store sdk.KVStore, cdc codec.BinaryCodec) error {
	prefixStore := prefix.NewStore(store, v1.RelationshipsStorePrefix)
	iterator := prefixStore.Iterator(nil, nil)

	// Read all the existing relationships and the associated keys
	var keys [][]byte
	var relationships []types.Relationship
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, append(v1.RelationshipsStorePrefix, iterator.Key()...))

		var relationship types.Relationship
		err := cdc.Unmarshal(iterator.Value(), &relationship)
		if err != nil {
			return err
		}
		relationships = append(relationships, relationship)
	}
	iterator.Close()

	for i, relationship := range relationships {
		// Delete the old key
		store.Delete(keys[i])

		bz, err := cdc.Marshal(&relationships[i])
		if err != nil {
			return err
		}

		// Store the relationship with the new key
		store.Set(types.RelationshipsStoreKey(relationship.Creator, relationship.Counterparty, relationship.SubspaceID), bz)
	}

	return nil
}

// migrateUserBlocks migrates all the user blocks from using the old key to using the new key
func migrateUserBlocks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	prefixStore := prefix.NewStore(store, v1.UsersBlocksStorePrefix)
	iterator := prefixStore.Iterator(nil, nil)

	// Read all the existing blocks and the associated keys
	var keys [][]byte
	var blocks []types.UserBlock
	for ; iterator.Valid(); iterator.Next() {
		keys = append(keys, append(v1.UsersBlocksStorePrefix, iterator.Key()...))

		var block types.UserBlock
		err := cdc.Unmarshal(iterator.Value(), &block)
		if err != nil {
			return err
		}
		blocks = append(blocks, block)
	}
	iterator.Close()

	for i, block := range blocks {
		// Delete the old key
		store.Delete(keys[i])

		bz, err := cdc.Marshal(&blocks[i])
		if err != nil {
			return err
		}

		// Store the block with the new key
		store.Set(types.UserBlockStoreKey(block.Blocker, block.Blocked, block.SubspaceID), bz)
	}

	return nil
}
