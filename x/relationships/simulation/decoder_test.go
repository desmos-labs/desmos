package simulation_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/relationships/simulation"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	firstAddr := ed25519.GenPrivKey().PubKey().Address().String()
	secondAddr := ed25519.GenPrivKey().PubKey().Address().String()

	relationships := []types.Relationship{
		types.NewRelationship(
			firstAddr,
			secondAddr,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
		types.NewRelationship(
			secondAddr,
			firstAddr,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}
	relBz, err := cdc.MarshalBinaryBare(&types.Relationships{Relationships: relationships})
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
		{Key: types.RelationshipsStoreKey(firstAddr), Value: relBz},
		{Key: types.UsersBlocksStoreKey(firstAddr), Value: blocksBz},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Relationships", fmt.Sprintf("Relationships A: %s\nRelationships B: %s\n", relationships, relationships)},
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
