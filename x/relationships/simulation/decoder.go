package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/v3/x/relationships/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding relationships type
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.RelationshipsStorePrefix):
			var relationshipA, relationshipB types.Relationship
			cdc.MustUnmarshal(kvA.Value, &relationshipA)
			cdc.MustUnmarshal(kvB.Value, &relationshipB)
			return fmt.Sprintf("RelationshipA: %s\nRelationshipB: %s\n",
				&relationshipA, &relationshipB)

		case bytes.HasPrefix(kvA.Key, types.UsersBlocksStorePrefix):
			var userBlockA, userBlockB types.UserBlock
			cdc.MustUnmarshal(kvA.Value, &userBlockA)
			cdc.MustUnmarshal(kvB.Value, &userBlockB)
			return fmt.Sprintf("UserBlockA: %s\nUserBlockB: %s\n",
				&userBlockA, &userBlockB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
