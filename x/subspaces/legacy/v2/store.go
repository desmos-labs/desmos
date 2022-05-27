package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	err := fixGroupsPermissions(store, cdc)
	if err != nil {
		return err
	}

	fixUsersPermissions(store)

	return nil
}

func fixGroupsPermissions(store sdk.KVStore, cdc codec.BinaryCodec) error {
	groupsStore := prefix.NewStore(store, GroupsPrefix)
	iterator := groupsStore.Iterator(nil, nil)

	var groups []UserGroup
	for ; iterator.Valid(); iterator.Next() {
		var group UserGroup
		err := cdc.Unmarshal(iterator.Value(), &group)
		if err != nil {
			return err
		}

		// Sanitize the permissions
		group.Permissions = SanitizePermission(group.Permissions)
		groups = append(groups, group)
	}

	iterator.Close()

	// Store the new groups
	for i, group := range groups {
		bz, err := cdc.Marshal(&groups[i])
		if err != nil {
			return err
		}

		store.Set(GroupStoreKey(group.SubspaceID, group.ID), bz)
	}

	return nil
}

type userPermissionDetails struct {
	subspaceID  uint64
	user        sdk.AccAddress
	permissions types.Permission
}

func fixUsersPermissions(store sdk.KVStore) {
	permissionsStore := prefix.NewStore(store, UserPermissionsStorePrefix)
	iterator := permissionsStore.Iterator(nil, nil)

	var permissions []userPermissionDetails
	for ; iterator.Valid(); iterator.Next() {
		// The first 8 bytes are the subspace id (uint64 takes up 8 bytes)
		// The remaining bytes are the user address
		subspaceBz, addressBz := iterator.Key()[:8], iterator.Key()[8:]

		permissions = append(permissions, userPermissionDetails{
			subspaceID: GetSubspaceIDFromBytes(subspaceBz),
			user:       GetAddressBytes(addressBz),

			// Sanitize the permission
			permissions: SanitizePermission(types.UnmarshalPermission(iterator.Value())),
		})
	}

	iterator.Close()

	// Store the new permissions
	for _, entry := range permissions {
		store.Set(UserPermissionStoreKey(entry.subspaceID, entry.user), types.MarshalPermission(entry.permissions))
	}
}
