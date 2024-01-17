package ibctesting

import (
	"testing"
	"time"

	ibctesting "github.com/cosmos/ibc-go/v8/testing"
)

var (
	globalStartTime = time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
)

func NewCoordinator(t *testing.T, n int) *ibctesting.Coordinator {
	chains := make(map[string]*ibctesting.TestChain)
	coord := &ibctesting.Coordinator{
		T:           t,
		CurrentTime: globalStartTime,
	}

	for i := 1; i <= n; i++ {
		chainID := ibctesting.GetChainID(i)
		chains[chainID] = NewTestChain(t, coord, chainID)
	}
	coord.Chains = chains

	return coord
}
