package v3

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	v2 "github.com/desmos-labs/desmos/v3/x/subspaces/legacy/v2"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// MigrateStore migrates the store from version 2 to version 3.
// The migration includes the following:
//
// - migrate all the group permissions from the old system to the new one
// - migrate all the user permissions from the old system to the new one
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	err := migrateUserGroupsPermissions(store, cdc)
	if err != nil {
		return err
	}

	err = migrateUserPermissions(store, cdc)
	if err != nil {
		return err
	}

	return nil
}

// migrateUserGroupsPermissions migrates all the user groups permissions from the old system to the new one
func migrateUserGroupsPermissions(store sdk.KVStore, cdc codec.BinaryCodec) error {
	groupsStore := prefix.NewStore(store, v2.GroupsPrefix)
	iterator := groupsStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var v2Group v2.UserGroup
		err := cdc.Unmarshal(iterator.Value(), &v2Group)
		if err != nil {
			return err
		}

		// Delete the old key
		store.Delete(v2.GroupStoreKey(v2Group.SubspaceID, v2Group.ID))

		// Migrate the permission
		v3Permissions, err := migratePermissions(v2Group.Permissions)
		if err != nil {
			return err
		}

		// Store the new group
		v3Group := types.NewUserGroup(v2Group.SubspaceID, v2Group.ID, v2Group.Name, v2Group.Description, v3Permissions)
		store.Set(types.GroupStoreKey(v3Group.SubspaceID, v3Group.ID), cdc.MustMarshal(&v3Group))
	}

	return nil
}

// migrateUserPermissions migrates all the user permissions from the old system to the new one
func migrateUserPermissions(store sdk.KVStore, cdc codec.BinaryCodec) error {
	permissionsStore := prefix.NewStore(store, v2.UserPermissionsStorePrefix)
	iterator := permissionsStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		// Get the permissions value
		subspaceID, user := v2.SplitUserPermissionKey(append(v2.UserPermissionsStorePrefix, iterator.Key()...))
		v2Permissions := v2.UnmarshalPermission(iterator.Value())

		// Delete the old key
		store.Delete(v2.UserPermissionStoreKey(subspaceID, user))

		// Migrate the permissions
		v3Permissions, err := migratePermissions(v2Permissions)
		if err != nil {
			return err
		}

		// If there are no permissions, just skip this migration
		if len(v3Permissions) == 0 {
			continue
		}

		// Store the new permissions
		userPermission := types.NewUserPermission(subspaceID, user.String(), v3Permissions)
		store.Set(types.UserPermissionStoreKey(subspaceID, user), cdc.MustMarshal(&userPermission))
	}

	return nil
}

// migratePermissions migrates the given v2 permissions to the new system
func migratePermissions(permissions v2.Permission) (types.Permissions, error) {
	v2PermissionsSlice := v2.SplitPermissions(permissions)
	v3Permissions := make([]types.Permission, len(v2PermissionsSlice))
	for i, permission := range v2PermissionsSlice {
		v3Permission, err := migratePermission(permission)
		if err != nil {
			return nil, err
		}
		v3Permissions[i] = v3Permission
	}
	return v3Permissions, nil
}

// migratePermission migrates the given permission value to the corresponding new one
func migratePermission(permission v2.Permission) (types.Permission, error) {
	switch permission {
	case v2.PermissionWrite:
		return poststypes.PermissionWrite, nil
	case v2.PermissionModerateContent:
		return poststypes.PermissionModerateContent, nil
	case v2.PermissionChangeInfo:
		return types.PermissionEditSubspace, nil
	case v2.PermissionManageGroups:
		return types.PermissionManageGroups, nil
	case v2.PermissionSetPermissions:
		return types.PermissionSetPermissions, nil
	case v2.PermissionDeleteSubspace:
		return types.PermissionDeleteSubspace, nil
	case v2.PermissionEverything:
		return types.PermissionEverything, nil
	default:
		return "", fmt.Errorf("permission not supported: %d", permission)
	}
}
