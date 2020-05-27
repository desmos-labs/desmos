package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/tendermint/tendermint/libs/kv"
)

// DecodeStore unmarshals the KVPair's Reports to the corresponding posts type
func DecodeStore(cdc *codec.Codec, kvA, kvB kv.Pair) string {
	switch {
	case bytes.HasPrefix(kvA.Key, types.ReportsStorePrefix):
		var reportsA, reportsB types.Reports
		cdc.MustUnmarshalBinaryBare(kvA.Value, &reportsA)
		cdc.MustUnmarshalBinaryBare(kvB.Value, &reportsB)
		return fmt.Sprintf("ReportsA: %s\nReportsB: %s\n", reportsA, reportsB)
	default:
		panic(fmt.Sprintf("invalid account key %X", kvA.Key))
	}
}
