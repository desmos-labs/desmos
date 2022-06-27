package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding relationships type
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.DTagPrefix):
			addressA := sdk.AccAddress(kvA.Value).String()
			addressB := sdk.AccAddress(kvB.Value).String()
			return fmt.Sprintf("DTagAddressA: %s\nDTagAddressB: %s\n", addressA, addressB)

		case bytes.HasPrefix(kvA.Key, types.DTagTransferRequestPrefix):
			var requestA, requestB types.DTagTransferRequest
			cdc.MustUnmarshal(kvA.Value, &requestA)
			cdc.MustUnmarshal(kvB.Value, &requestB)
			return fmt.Sprintf("RequestA: %s\nRequestB: %s\n", requestA, requestB)

		case bytes.HasPrefix(kvA.Key, types.ChainLinksPrefix):
			var chainLinkA, chainLinkB types.ChainLink
			cdc.MustUnmarshal(kvA.Value, &chainLinkA)
			cdc.MustUnmarshal(kvB.Value, &chainLinkB)
			return fmt.Sprintf("Chain link A: %s\nChain link B: %s\n", chainLinkA, chainLinkB)

		case bytes.HasPrefix(kvA.Key, types.ApplicationLinkPrefix):
			var applicationLinkA, applicationLinkB types.ApplicationLink
			cdc.MustUnmarshal(kvA.Value, &applicationLinkA)
			cdc.MustUnmarshal(kvB.Value, &applicationLinkB)
			return fmt.Sprintf("Application link A: %s\nApplication link B: %s\n", &applicationLinkA, &applicationLinkB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
