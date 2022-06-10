package simulation_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/app"

	"github.com/desmos-labs/desmos/v3/x/relationships/simulation"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	relationship := types.NewRelationship(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		1,
	)

	userBlock := types.NewUserBlock(
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		"reason",
		1,
	)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{

		{
			Key: types.RelationshipsStoreKey(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			Value: cdc.MustMarshal(&relationship),
		},
		{
			Key: types.UserBlockStoreKey(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				1,
			),
			Value: cdc.MustMarshal(&userBlock),
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Relationship", fmt.Sprintf("Relationships A: %s\nRelationships B: %s\n", &relationship, &relationship)},
		{"User block", fmt.Sprintf("User block A: %s\nUser block B: %s\n", &userBlock, &userBlock)},
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
