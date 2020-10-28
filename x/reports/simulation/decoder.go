package simulation

import (
	"bytes"
	"fmt"
	"github.com/desmos-labs/desmos/x/reports/keeper"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// NewDecodeStore unmarshalls the KVPair's Reports to the corresponding reports type
func NewDecodeStore(k keeper.Keeper) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.ReportsStorePrefix):
			reportsA, err := k.UnmarshalReports(kvA.Value)
			if err != nil {
				panic(err)
			}

			reportsB, err := k.UnmarshalReports(kvB.Value)
			if err != nil {
				panic(err)
			}

			return fmt.Sprintf("ReportsA: %s\nReportsB: %s\n", reportsA, reportsB)

		default:
			panic(fmt.Sprintf("invalid account key %X", kvA.Key))
		}
	}
}
