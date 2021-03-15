package simulation

import (
	"bytes"
	"fmt"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding relationships type
func NewDecodeStore(cdc codec.Marshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.DTagTransferRequestsPrefix):
			var requestsA, requestsB keeper.WrappedDTagTransferRequests
			cdc.MustUnmarshalBinaryBare(kvA.Value, &requestsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &requestsB)
			return fmt.Sprintf("RequestsA: %s\nRequestsB: %s\n", requestsA.Requests, requestsB.Requests)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
