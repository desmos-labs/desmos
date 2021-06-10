package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding subspaces type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.SubspaceStorePrefix):
			var subspaceA, subspaceB types.Subspace
			cdc.MustUnmarshalBinaryBare(kvA.Value, &subspaceA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &subspaceB)
			return fmt.Sprintf("SubspaceA: %s\nSubspaceB: %s\n", subspaceA.String(), subspaceB.String())
		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
