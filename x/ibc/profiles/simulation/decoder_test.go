package simulation_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/kv"
	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/x/ibc/profiles/simulation"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	src := ed25519.GenPrivKey().PubKey().Address().String()
	dst := ed25519.GenPrivKey().PubKey().Address().String()

	link := types.NewLink(dst, src)

	linkBz, err := cdc.MarshalBinaryBare(&link)
	require.NoError(t, err)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{Key: types.LinkStoreKey(dst), Value: linkBz},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Link", fmt.Sprintf("Link: %s\n", link)},
		{"other", ""},
	}

	for i, test := range tests {
		i, test := i, test
		t.Run(test.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, test.name)
			default:
				require.Equal(t, test.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), test.name)
			}
		})
	}
}
