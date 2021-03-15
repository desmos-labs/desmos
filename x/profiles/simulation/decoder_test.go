package simulation_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/profiles/keeper"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/simulation"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	requests := keeper.NewWrappedDTagTransferRequests([]types.DTagTransferRequest{
		types.NewDTagTransferRequest(
			"dtag",
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		),
	})

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.DtagTransferRequestStoreKey("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			Value: cdc.MustMarshalBinaryBare(&requests),
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Requests", fmt.Sprintf("RequestsA: %s\nRequestsB: %s\n", requests.Requests, requests.Requests)},
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
