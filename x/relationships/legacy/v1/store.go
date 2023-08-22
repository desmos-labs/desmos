package v1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"

	profilesv4 "github.com/desmos-labs/desmos/v6/x/profiles/legacy/v4"
	"github.com/desmos-labs/desmos/v6/x/relationships/types"
)

// MigrateStore performs in-place store migrations from v0 to v1
// The migration includes:
//
// - store all the relationships present inside x/profiles into this module
// - store all the user blocks present inside x/profiles into this module
//
// NOTE: This method must be called BEFORE the migration from v4 to v5 of the profiles module.
//
//	If this order is not preserved, all relationships and blocks WILL BE DELETED.
func MigrateStore(ctx sdk.Context, pk profilesv4.Keeper, relationshipsStoreKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(relationshipsStoreKey)

	err := migrateUserBlocks(ctx, pk, store, cdc)
	if err != nil {
		return err
	}

	err = migrateRelationships(ctx, pk, store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// migrateUserBlocks migrates the user blocks stored to the new type, converting the subspace from string to uint64
func migrateUserBlocks(ctx sdk.Context, pk profilesv4.Keeper, store sdk.KVStore, cdc codec.BinaryCodec) error {
	for _, v230Block := range pk.GetBlocks(ctx) {
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

		// Store the new value inside the relationships store
		store.Set(types.UserBlockStoreKey(v300Block.Blocker, v300Block.Blocked, v300Block.SubspaceID), blockBz)
	}

	return nil
}

// migrateRelationships migrates the relationships stored to the new type, converting the subspace from string to uint64
func migrateRelationships(ctx sdk.Context, pk profilesv4.Keeper, store sdk.KVStore, cdc codec.BinaryCodec) error {
	for _, v230Relationship := range pk.GetRelationships(ctx) {
		// Get the subspace id
		subspaceID, err := subspacestypes.ParseSubspaceID(v230Relationship.SubspaceID)
		if err != nil {
			return err
		}

		// Do not migrate the relationships for the subspace with ID 0
		if subspaceID == 0 {
			continue
		}

		// Serialize the relationship as the new type
		v300Relationship := types.NewRelationship(v230Relationship.Creator, v230Relationship.Recipient, subspaceID)
		relationshipBz, err := cdc.Marshal(&v300Relationship)
		if err != nil {
			return err
		}

		// Store the new relationship inside the relationships store
		store.Set(types.RelationshipsStoreKey(v300Relationship.Creator, v300Relationship.Counterparty, v300Relationship.SubspaceID), relationshipBz)
	}

	return nil
}
