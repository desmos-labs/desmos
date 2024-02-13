package simulation

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/v7/x/reports/types"
)

// NewDecodeStore returns a new decoder that unmarshals the KVPair's Value
// to the corresponding reports type
func NewDecodeStore(cdc codec.Codec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.HasPrefix(kvA.Key, types.NextReportIDPrefix):
			var idA, idB uint64
			idA = types.GetReportIDFromBytes(kvA.Value)
			idB = types.GetReportIDFromBytes(kvB.Value)
			return fmt.Sprintf("NextReportIDA: %d\nNextReportIDB: %d\n", idA, idB)

		case bytes.HasPrefix(kvA.Key, types.ReportPrefix):
			var reportA, reportB types.Report
			cdc.MustUnmarshal(kvA.Value, &reportA)
			cdc.MustUnmarshal(kvB.Value, &reportB)
			return fmt.Sprintf("ReportA: %s\nReportB: %s\n", &reportA, &reportB)

		case bytes.HasPrefix(kvA.Key, types.PostsReportsPrefix):
			var idA, idB uint64
			idA = types.GetReportIDFromBytes(kvA.Value)
			idB = types.GetReportIDFromBytes(kvB.Value)
			return fmt.Sprintf("PostReportIDA: %d\nPostReportIDB: %d\n", idA, idB)

		case bytes.HasPrefix(kvA.Key, types.UsersReportsPrefix):
			var idA, idB uint64
			idA = types.GetReportIDFromBytes(kvA.Value)
			idB = types.GetReportIDFromBytes(kvB.Value)
			return fmt.Sprintf("UserReportIDA: %d\nUserReportIDB: %d\n", idA, idB)

		case bytes.HasPrefix(kvA.Key, types.NextReasonIDPrefix):
			var idA, idB uint32
			idA = types.GetReasonIDFromBytes(kvA.Value)
			idB = types.GetReasonIDFromBytes(kvB.Value)
			return fmt.Sprintf("NextReasonIDA: %d\nNextReasonIDB: %d\n", idA, idB)

		case bytes.HasPrefix(kvA.Key, types.ReasonPrefix):
			var reasonA, reasonB types.Reason
			cdc.MustUnmarshal(kvA.Value, &reasonA)
			cdc.MustUnmarshal(kvB.Value, &reasonB)
			return fmt.Sprintf("ReasonA: %s\nReasonB: %s\n", &reasonA, &reasonB)

		default:
			panic(fmt.Sprintf("unexpected %s key %X (%s)", types.ModuleName, kvA.Key, kvA.Key))
		}
	}
}
