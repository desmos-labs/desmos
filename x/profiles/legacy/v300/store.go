package v300

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v230 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v230"
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
	var keys [][]byte
	var values []v230.UserBlock

	userBlocksStore := prefix.NewStore(store, types.UsersBlocksStorePrefix)
	iterator := userBlocksStore.Iterator(nil, nil)

	for ; iterator.Valid(); iterator.Next() {
		// Get the keys
		keys = append(keys, iterator.Key())

		// Get the associated values
		var block v230.UserBlock
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

	for i := 0; i < len(keys); i++ {
		// Serialize the block as the new type
		blockBz, err := cdc.Marshal(&types.UserBlock{
			Blocker:  values[i].Blocker,
			Blocked:  values[i].Blocked,
			Reason:   values[i].Reason,
			Subspace: 0,
		})
		if err != nil {
			return err
		}

		// Set the key inside the store
		store.Set(keys[i], blockBz)
	}

	return nil
}

// migrateRelationships migrates the relationships stored to the new type, converting the subspace from string to uint64
func migrateRelationships(store sdk.KVStore, cdc codec.BinaryCodec) error {
	var keys [][]byte
	var values []v230.Relationship

	relationshipsStore := prefix.NewStore(store, types.RelationshipsStorePrefix)
	iterator := relationshipsStore.Iterator(nil, nil)

	for ; iterator.Valid(); iterator.Next() {
		// Get the keys
		keys = append(keys, iterator.Key())

		// Get the associated values
		var block v230.Relationship
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

	for i := 0; i < len(keys); i++ {
		// Serialize the relationship as the new type
		blockBz, err := cdc.Marshal(&types.Relationship{
			Creator:   values[i].Creator,
			Recipient: values[i].Recipient,
			Subspace:  0,
		})
		if err != nil {
			return err
		}

		// Set the key inside the store
		store.Set(keys[i], blockBz)
	}

	return nil
}
