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

		case bytes.HasPrefix(kvA.Key, types.DTagTransferRequestsPrefix):
			var requestsA, requestsB types.DTagTransferRequests
			cdc.MustUnmarshalBinaryBare(kvA.Value, &requestsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &requestsB)
			return fmt.Sprintf("RequestsA: %s\nRequestsB: %s\n", requestsA.Requests, requestsB.Requests)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
