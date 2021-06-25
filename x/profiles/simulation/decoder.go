package simulation

import (
	"bytes"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding relationships type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.DTagPrefix):
			addressA := sdk.AccAddress(bytes.TrimPrefix(kvA.Value, types.DTagPrefix)).String()
			addressB := sdk.AccAddress(bytes.TrimPrefix(kvB.Value, types.DTagPrefix)).String()
			return fmt.Sprintf("DTagAddressA: %s\nDTagAddressB: %s\n", addressA, addressB)

		case bytes.HasPrefix(kvA.Key, types.DTagTransferRequestPrefix):
			var requestA, requestB types.DTagTransferRequest
			cdc.MustUnmarshalBinaryBare(kvA.Value, &requestA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &requestB)
			return fmt.Sprintf("RequestA: %s\nRequestB: %s\n", requestA, requestB)

		case bytes.HasPrefix(kvA.Key, types.RelationshipsStorePrefix):
			var relationshipA, relationshipB types.Relationship
			cdc.MustUnmarshalBinaryBare(kvA.Value, &relationshipA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &relationshipB)
			return fmt.Sprintf("Relationships A: %s\nRelationships B: %s\n",
				relationshipA, relationshipB)

		case bytes.HasPrefix(kvA.Key, types.UsersBlocksStorePrefix):
			var userBlocksA, userBlocksB types.UserBlocks
			cdc.MustUnmarshalBinaryBare(kvA.Value, &userBlocksA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &userBlocksB)
			return fmt.Sprintf("User blocks A: %s\nUser blocks B: %s\n",
				userBlocksA.Blocks, userBlocksB.Blocks)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
