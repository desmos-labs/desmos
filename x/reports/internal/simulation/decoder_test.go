package simulation

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/kv"
	"testing"
)

var (
	privKey           = ed25519.GenPrivKey().PubKey()
	reportCreatorAddr = sdk.AccAddress(privKey.Address())

	id  = posts.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 = posts.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
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
				require.Panics(t, func() { DecodeStore(cdc, kvPairs[i], kvPairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, DecodeStore(cdc, kvPairs[i], kvPairs[i]), tt.name)
			}
		})
	}
}
