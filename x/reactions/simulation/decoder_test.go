package simulation_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/app"
	"github.com/desmos-labs/desmos/v5/x/reactions/simulation"
	"github.com/desmos-labs/desmos/v5/x/reactions/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	decoder := simulation.NewDecodeStore(cdc)

	registeredReaction := types.NewRegisteredReaction(
		1,
		1,
		":hello:",
		"https://example.com?image=hello.png",
	)
	reaction := types.NewReaction(
		1,
		1,
		1,
		types.NewRegisteredReactionValue(1),
		"cosmos1z0glns8fv5h0xgghg4nkq0jjy9gp0l682tcf79",
	)
	reactionsParams := types.DefaultReactionsParams(1)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.NextRegisteredReactionIDStoreKey(1),
			Value: types.GetRegisteredReactionIDBytes(1),
		},
		{
			Key:   types.RegisteredReactionStoreKey(1, 1),
			Value: cdc.MustMarshal(&registeredReaction),
		},
		{
			Key:   types.NextReactionIDStoreKey(1, 1),
			Value: types.GetReactionIDBytes(1),
		},
		{
			Key:   types.ReactionStoreKey(1, 1, 1),
			Value: cdc.MustMarshal(&reaction),
		},
		{
			Key:   types.SubspaceReactionsParamsStoreKey(1),
			Value: cdc.MustMarshal(&reactionsParams),
		},
		{
			Key:   []byte("Unknown key"),
			Value: nil,
		},
	}}

	testCases := []struct {
		name        string
		expectedLog string
	}{
		{"Next Registered Reaction ID", fmt.Sprintf("NextRegisteredReactionIDA: %d\nNextRegisteredReactionIDB: %d\n",
			1, 1)},
		{"Registered Reaction", fmt.Sprintf("RegisteredReactionA: %s\nRegisteredReactionB: %s\n",
			&registeredReaction, &registeredReaction)},
		{"Next Reaction ID", fmt.Sprintf("NextReactionIDA: %d\nNextReactionIDB: %d\n",
			1, 1)},
		{"Reaction", fmt.Sprintf("ReactionA: %s\nReactionB: %s\n",
			&reaction, &reaction)},
		{"Reactions Params", fmt.Sprintf("SubspaceParamsA: %s\nSubspaceParamsB: %s\n",
			&reactionsParams, &reactionsParams)},
		{"other", ""},
	}

	for i, tc := range testCases {
		i, tc := i, tc
		t.Run(tc.name, func(t *testing.T) {
			switch i {
			case len(testCases) - 1:
				require.Panics(t, func() { decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tc.name)
			default:
				require.Equal(t, tc.expectedLog, decoder(kvPairs.Pairs[i], kvPairs.Pairs[i]), tc.name)
			}
		})
	}
}
