package v5

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

// MigrateStore migrates the store from version 5 to version 6.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	return moveDefaultUserGroupToRootSection(ctx, storeKey, cdc)
}

func moveDefaultUserGroupToRootSection(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	groupsStore := prefix.NewStore(store, types.GroupsPrefix)
	iterator := groupsStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var oldGroup types.UserGroup
		cdc.MustUnmarshal(iterator.Value(), &oldGroup)

		if oldGroup.ID != types.DefaultGroupID {
			// Skip because the group is not default
			continue
		}

		if oldGroup.SectionID == types.RootSectionID {
			// Skip because the default group has not been moved
			continue
		}

		// Delete moved old default group key
		store.Delete(types.GroupStoreKey(oldGroup.SubspaceID, oldGroup.SectionID, oldGroup.ID))

		// Save new default group key
		newGroup := types.NewUserGroup(
			oldGroup.SubspaceID,
			types.RootSectionID,
			types.DefaultGroupID,
			oldGroup.Name,
			oldGroup.Description,
			oldGroup.Permissions,
		)

		store.Set(
			types.GroupStoreKey(newGroup.SubspaceID, newGroup.SectionID, newGroup.ID),
			cdc.MustMarshal(&newGroup),
		)
	}
	return nil
}
