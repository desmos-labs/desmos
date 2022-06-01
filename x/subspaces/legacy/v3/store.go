package v3

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// MigrateStore migrates the store to version 3.
// The migration process includes the following operations:
// - set the NextSectionID to 1 for all the existing subspaces
// - create the default section for all the existing subspaces
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	err := setupSubspacesSections(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// setupSubspacesSections sets the NextSectionID to 1 for all the current subspaces and creates the default sections
func setupSubspacesSections(store sdk.KVStore, cdc codec.BinaryCodec) error {
	subspacesStore := prefix.NewStore(store, types.SubspacePrefix)
	iterator := subspacesStore.Iterator(nil, nil)

	for ; iterator.Valid(); iterator.Next() {
		var subspace types.Subspace
		if err := cdc.Unmarshal(iterator.Value(), &subspace); err != nil {
			return err
		}

		// Set the initial section id
		store.Set(types.NextSectionIDStoreKey(subspace.ID), types.GetSectionIDBytes(1))

		// Create the default section
		defaultSection := types.DefaultSection(subspace.ID)
		store.Set(types.SectionStoreKey(subspace.ID, 0), cdc.MustMarshal(&defaultSection))
	}

	return nil
}
