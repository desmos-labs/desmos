package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/x/reports/types"
)

// NewDecodeStore unmarshalls the KVPair's Reports to the corresponding reports type
func NewDecodeStore(cdc codec.BinaryMarshaler) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.ReportsStorePrefix):
			var reportsA, reportsB types.Reports
			cdc.MustUnmarshalBinaryBare(kvA.Value, &reportsA)
			cdc.MustUnmarshalBinaryBare(kvB.Value, &reportsB)
			return fmt.Sprintf("ReportsA: %s\nReportsB: %s\n", reportsA.Reports, reportsB.Reports)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
