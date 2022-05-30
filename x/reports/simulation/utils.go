package simulation

// DONTCOVER

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

// RandomStandardReason returns a random standard reason from the given slice
func RandomStandardReason(r *rand.Rand, reasons []types.StandardReason) types.StandardReason {
	return reasons[r.Intn(len(reasons))]
}

// RandomReason returns a random reason from the given slice
func RandomReason(r *rand.Rand, reasons []types.Reason) types.Reason {
	return reasons[r.Intn(len(reasons))]
}

// GetRandomReasonTitle returns a random reason title
func GetRandomReasonTitle(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 30)
}

// GetRandomReasonDescription returns a random reason description
func GetRandomReasonDescription(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 50)
}

// GetRandomMessage returns a random reporting message
func GetRandomMessage(r *rand.Rand) string {
	return simtypes.RandStringOfLength(r, 30)
}

// RandomReport returns a random report from the given slice
func RandomReport(r *rand.Rand, reports []types.Report) types.Report {
	return reports[r.Intn(len(reports))]
}

// GetRandomStandardReasons returns a randomly generated slice of standard reason
func GetRandomStandardReasons(r *rand.Rand, num int) []types.StandardReason {
	standardReasons := make([]types.StandardReason, num)
	for i := 0; i < num; i++ {
		standardReasons[i] = types.NewStandardReason(
			uint32(i+1),
			GetRandomReasonTitle(r),
			GetRandomReasonDescription(r),
		)
	}
	return standardReasons
}
