package simulation

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/kv"
	"testing"
)

var (
	privKey            = ed25519.GenPrivKey().PubKey()
	accountCreatorAddr = sdk.AccAddress(privKey.Address())

	anotherKey      = ed25519.GenPrivKey().PubKey()
	anotherUserAddr = sdk.AccAddress(anotherKey.Address())

	subspace = "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	relationships = types.Relationships{
		types.NewRelationship(accountCreatorAddr, subspace),
		types.NewRelationship(anotherUserAddr, subspace),
	}
)

func makeTestCodec() (cdc *codec.Codec) {
	cdc = codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	types.RegisterCodec(cdc)
	return
}

func TestDecodeStore(t *testing.T) {
	cdc := makeTestCodec()

	kvPairs := kv.Pairs{
		kv.Pair{Key: types.RelationshipsStoreKey(accountCreatorAddr), Value: cdc.MustMarshalBinaryBare(&relationships)},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Relationships", fmt.Sprintf("Relationships: %s\nRelationships: %s\n", relationships, relationships)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
