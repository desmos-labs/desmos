package v300

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	v2 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v2"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v2.3 to v3.0.
// The migration includes:
//
// - replace all relationship subspace id from string to uint64
// - replace all user blocks subspace id from string to uint64
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	err := migrateUserBlocks(store, cdc)
	if err != nil {
		return err
	}

	err = migrateRelationships(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// migrateUserBlocks migrates the user blocks stored to the new type, converting the subspace from string to uint64
func migrateUserBlocks(store sdk.KVStore, cdc codec.BinaryCodec) error {
	var values []v2.UserBlock

	userBlocksStore := prefix.NewStore(store, v2.UsersBlocksStorePrefix)
	iterator := userBlocksStore.Iterator(nil, nil)

	for ; iterator.Valid(); iterator.Next() {
		var block v2.UserBlock
		err := cdc.Unmarshal(iterator.Value(), &block)
		if err != nil {
			return err
		}
		values = append(values, block)
	}

	// Close the iterator
	err := iterator.Close()
	if err != nil {
		return err
	}

	for _, v230Block := range values {
		// Delete the previous key
		store.Delete(v2.UserBlockStoreKey(v230Block.Blocker, v230Block.SubspaceID, v230Block.Blocked))

		// Get the subspace id
		subspaceID, err := subspacestypes.ParseSubspaceID(v230Block.SubspaceID)
		if err != nil {
			return err
		}

		// Serialize the block as the new type
		v300Block := types.NewUserBlock(v230Block.Blocker, v230Block.Blocked, v230Block.Reason, subspaceID)
		blockBz, err := cdc.Marshal(&v300Block)
		if err != nil {
			return err
		}

		// Store the new value inside the store
		store.Set(types.UserBlockStoreKey(v300Block.Blocker, v300Block.SubspaceID, v300Block.Blocked), blockBz)
	}

	return nil
}

// migrateRelationships migrates the relationships stored to the new type, converting the subspace from string to uint64
func migrateRelationships(store sdk.KVStore, cdc codec.BinaryCodec) error {
	var values []v2.Relationship

	relationshipsStore := prefix.NewStore(store, types.RelationshipsStorePrefix)
	iterator := relationshipsStore.Iterator(nil, nil)

	for ; iterator.Valid(); iterator.Next() {
		var relationship v2.Relationship
		err := cdc.Unmarshal(iterator.Value(), &relationship)
		if err != nil {
			return err
		}
		values = append(values, relationship)
	}

	// Close the iterator
	err := iterator.Close()
	if err != nil {
		return err
	}

	for _, v230Relationship := range values {
		// Delete the previous key
		store.Delete(v2.RelationshipsStoreKey(v230Relationship.Creator, v230Relationship.Subspace, v230Relationship.Recipient))

		// Get the subspace id
		subspaceID, err := subspacestypes.ParseSubspaceID(v230Relationship.Subspace)
		if err != nil {
			return err
		}

		// Serialize the relationship as the new type
		v300Relationship := types.NewRelationship(v230Relationship.Creator, v230Relationship.Recipient, subspaceID)
		relationshipBz, err := cdc.Marshal(&v300Relationship)
		if err != nil {
			return err
		}

		// Store the new relationship inside the store
		store.Set(types.RelationshipsStoreKey(v300Relationship.Creator, v300Relationship.SubspaceID, v300Relationship.Recipient), relationshipBz)
	}

	return nil
}
