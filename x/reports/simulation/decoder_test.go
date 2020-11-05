package simulation_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/reports/simulation"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/desmos-labs/desmos/x/reports/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	address := ed25519.GenPrivKey().PubKey().Address().String()
	reports := []types.Report{
		types.NewReport(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			"offense",
			"it offends me",
			address,
		),
		types.NewReport(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			"scam",
			"it's a scam",
			address,
		),
	}

	wrapped := types.Reports{Reports: reports}
	bz, err := cdc.MarshalBinaryBare(&wrapped)
	require.NoError(t, err)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.ReportStoreKey("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
			Value: bz,
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Report", fmt.Sprintf("ReportsA: %s\nReportsB: %s\n", reports, reports)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
