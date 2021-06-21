package simulation_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"

	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/desmos-labs/desmos/x/profiles/simulation"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	requests := types.NewDTagTransferRequests([]types.DTagTransferRequest{
		types.NewDTagTransferRequest(
			"dtag",
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		),
	})

	addr, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	firstAddr := ed25519.GenPrivKey().PubKey().Address().String()
	secondAddr := ed25519.GenPrivKey().PubKey().Address().String()

	relationship := types.NewRelationship(
		firstAddr,
		secondAddr,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	relBz, err := cdc.MarshalBinaryBare(&relationship)
	require.NoError(t, err)

	usersBlocks := []types.UserBlock{
		types.NewUserBlock(
			firstAddr,
			secondAddr,
			"reason",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewUserBlock(
			secondAddr,
			firstAddr,
			"reason",
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}
	blocksBz, err := cdc.MarshalBinaryBare(&types.UserBlocks{Blocks: usersBlocks})
	require.NoError(t, err)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.DTagStoreKey("AAkvohxhflhXsuyMg"),
			Value: addr,
		},
		{
			Key:   types.DTagTransferRequestStoreKey("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			Value: cdc.MustMarshalBinaryBare(&requests),
		},
		{
			Key:   types.RelationshipsStoreKey(firstAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", secondAddr),
			Value: relBz,
		},
		{
			Key:   types.UsersBlocksStoreKey(firstAddr,"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", secondAddr),
			Value: blocksBz,
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"DTags", fmt.Sprintf("DTagAddressA: %s\nDTagAddressB: %s\n", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")},
		{"Requests", fmt.Sprintf("RequestsA: %s\nRequestsB: %s\n", requests.Requests, requests.Requests)},
		{"Relationships", fmt.Sprintf("Relationships A: %s\nRelationships B: %s\n", relationship, relationship)},
		{"UsersBlocks", fmt.Sprintf("User blocks A: %s\nUser blocks B: %s\n", usersBlocks, usersBlocks)},
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
