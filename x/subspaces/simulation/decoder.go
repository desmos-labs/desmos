package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding subspaces type
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.SubspaceIDKey):
			var idA, idB uint64
			idA = types.GetSubspaceIDFromBytes(kvA.Value)
			idB = types.GetSubspaceIDFromBytes(kvB.Value)
			return fmt.Sprintf("SubspaceIDA: %d\nSubspaceIDB: %d\n", idA, idB)

		case bytes.HasPrefix(kvA.Key, types.SubspacePrefix):
			var subspaceA, subspaceB types.Subspace
			cdc.MustUnmarshal(kvA.Value, &subspaceA)
			cdc.MustUnmarshal(kvB.Value, &subspaceB)
			return fmt.Sprintf("SubspaceA: %s\nSubspaceB: %s\n", subspaceA.String(), subspaceB.String())

		case bytes.HasPrefix(kvA.Key, types.GroupIDPrefix):
			var groupIDA, groupIDB uint32
			groupIDA = types.GetGroupIDFromBytes(kvA.Value)
			groupIDB = types.GetGroupIDFromBytes(kvB.Value)
			return fmt.Sprintf("GroupIDA: %d\nGroupIDB: %d\n", groupIDA, groupIDB)

		case bytes.HasPrefix(kvA.Key, types.GroupsPrefix):
			var groupA, groupB types.UserGroup
			cdc.MustUnmarshal(kvA.Value, &groupA)
			cdc.MustUnmarshal(kvB.Value, &groupB)
			return fmt.Sprintf("GroupA: %s\nGroupB: %s\n", &groupA, &groupB)

		case bytes.HasPrefix(kvA.Key, types.GroupsMembersPrefix):
			return fmt.Sprintf("GroupMemberKeyA: %s\nGroupMemberKeyB: %s\n", kvA.Key, kvB.Key)

		case bytes.HasPrefix(kvA.Key, types.UserPermissionsStorePrefix):
			var permissionA, permissionB uint32
			permissionA = types.UnmarshalPermission(kvA.Value)
			permissionB = types.UnmarshalPermission(kvB.Value)
			return fmt.Sprintf("PermissionKeyA: %d\nPermissionKeyB: %d\n", permissionA, permissionB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
