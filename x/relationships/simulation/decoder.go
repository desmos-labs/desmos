package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

// RelationshipsUnmarshaler defines the expected encoding store functions.
type RelationshipsUnmarshaler interface {
	UnmarshalRelationships([]byte) ([]types.Relationship, error)
	UnmarshalUserBlocks([]byte) ([]types.UserBlock, error)
}

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding relationships type
func NewDecodeStore(cdc RelationshipsUnmarshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.RelationshipsStorePrefix):
			relationshipsA, err := cdc.UnmarshalRelationships(kvA.Value)
			if err != nil {
				panic(err)
			}

			relationshipsB, err := cdc.UnmarshalRelationships(kvB.Value)
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("Relationships: %s\nRelationships: %s\n", relationshipsA, relationshipsB)

		case bytes.HasPrefix(kvA.Key, types.UsersBlocksStorePrefix):
			userBlocksA, err := cdc.UnmarshalUserBlocks(kvA.Value)
			if err != nil {
				panic(err)
			}

			userBlocksB, err := cdc.UnmarshalUserBlocks(kvB.Value)
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("UsersBlocks: %s\nUsersBlocks: %s\n", userBlocksA, userBlocksB)

		default:
			panic(fmt.Sprintf("invalid relationships key %X", kvA.Key))
		}
	}
}
