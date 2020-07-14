package simulation_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/kv"

	posts "github.com/desmos-labs/desmos/x/posts/types"
	sim "github.com/desmos-labs/desmos/x/reports/simulation"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// nolint
var (
	privKey           = ed25519.GenPrivKey().PubKey()
	reportCreatorAddr = sdk.AccAddress(privKey.Address())
	id                = posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
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

	reports := types.Reports{
		types.NewReport("offense", "it offends me", reportCreatorAddr),
		types.NewReport("scam", "it's a scam", reportCreatorAddr),
	}

	kvPairs := kv.Pairs{
		kv.Pair{Key: types.ReportStoreKey(id), Value: cdc.MustMarshalBinaryBare(&reports)},
	}

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
				require.Panics(t, func() { sim.DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, sim.DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
