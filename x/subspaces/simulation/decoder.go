package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
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
			return fmt.Sprintf("SubspaceA: %s\nSubspaceB: %s\n", &subspaceA, &subspaceB)

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
			var permissionA, permissionB types.UserPermission
			cdc.MustUnmarshal(kvA.Value, &permissionA)
			cdc.MustUnmarshal(kvB.Value, &permissionB)
			return fmt.Sprintf("PermissionA: %s\nPermissionB: %s\n", &permissionA, &permissionB)

		case bytes.HasPrefix(kvA.Key, types.SectionIDPrefix):
			var sectionIDA, sectionIDB uint32
			sectionIDA = types.GetSectionIDFromBytes(kvA.Value)
			sectionIDB = types.GetSectionIDFromBytes(kvB.Value)
			return fmt.Sprintf("SectionIDA: %d\nSectionIDB: %d\n", sectionIDA, sectionIDB)

		case bytes.HasPrefix(kvA.Key, types.SectionsPrefix):
			var sectionA, sectionB types.Section
			cdc.MustUnmarshal(kvA.Value, &sectionA)
			cdc.MustUnmarshal(kvB.Value, &sectionB)
			return fmt.Sprintf("SectionA: %s\nSectionB: %s\n", &sectionA, &sectionB)

		case bytes.HasPrefix(kvA.Key, types.UserAllowancePrefix) ||
			bytes.HasPrefix(kvA.Key, types.GroupAllowancePrefix):
			var grantA, grantB types.Grant
			cdc.MustUnmarshal(kvA.Value, &grantA)
			cdc.MustUnmarshal(kvB.Value, &grantB)
			return fmt.Sprintf("GrantA: %s\nGrantB: %s\n", &grantA, &grantB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
