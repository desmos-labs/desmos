package simulation_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/staging/subspaces/simulation"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	decoder := simulation.NewDecodeStore(cdc)

	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	require.NoError(t, err)

	subspace := types.Subspace{
		ID:              "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		Name:            "test",
		Owner:           "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		Creator:         "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		CreationTime:    date,
		Open:            true,
		Admins:          []string{},
		BlockedUsers:    []string{},
		RegisteredUsers: []string{},
	}

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.SubspaceStoreKey(subspace.ID),
			Value: cdc.MustMarshalBinaryBare(&subspace),
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Subspace", fmt.Sprintf("SubspaceA: %s\nSubspaceB: %s\n", subspace.String(), subspace.String())},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
