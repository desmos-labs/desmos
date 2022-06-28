package v3

import (
	"fmt"

	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v2"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// MigrateStore migrates the store from version 2 to version 3.
// The migration includes the following:
//
// - migrate all the group permissions from the old system to the new one
// - migrate all the user permissions from the old system to the new one
// - set the NextSectionID to 1 for all the existing subspaces
// - create the default section for all the existing subspaces
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	err := setupSubspacesSections(store, cdc)
	if err != nil {
		return err
	}

	err = migrateUserGroupsPermissions(store, cdc)
	if err != nil {
		return err
	}

	err = migrateUserPermissions(store, cdc)
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
		v3Permissions, err := MigratePermissions(v2Group.Permissions)
		if err != nil {
			return err
		}

		// Store the new group
		v3Group := types.NewUserGroup(v2Group.SubspaceID, types.RootSectionID, v2Group.ID, v2Group.Name, v2Group.Description, v3Permissions)
		store.Set(types.GroupStoreKey(v3Group.SubspaceID, v3Group.SectionID, v3Group.ID), cdc.MustMarshal(&v3Group))
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
		v3Permissions, err := MigratePermissions(v2Permissions)
		if err != nil {
			return err
		}

		// If there are no permissions, just skip this migration
		if len(v3Permissions) == 0 {
			continue
		}

		// Store the new permissions
		userPermission := types.NewUserPermission(subspaceID, types.RootSectionID, user.String(), v3Permissions)
		store.Set(types.UserPermissionStoreKey(subspaceID, types.RootSectionID, user.String()), cdc.MustMarshal(&userPermission))
	}

	return nil
}

// MigratePermissions migrates the given v2 permissions to the new system
func MigratePermissions(permissions v2.Permission) (types.Permissions, error) {
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

	case v2.PermissionWrite:
		return poststypes.PermissionWrite, nil
	case v2.PermissionModerateContent:
		return poststypes.PermissionModerateContent, nil

	default:
		return "", fmt.Errorf("permission not supported: %d", permission)
	}
}
