package simulation

import (
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/kv"
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

	reportsTypes := types.ReportTypes{
		"scam",
		"offense,",
		"nudity",
	}

	kvPairs := kv.Pairs{
		kv.Pair{Key: types.ReportStoreKey(id), Value: cdc.MustMarshalBinaryBare(&reports)},
		kv.Pair{Key: types.ReportsTypeStorePrefix, Value: cdc.MustMarshalBinaryBare(&reportsTypes)},
	}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Report", fmt.Sprintf("ReportsA: %s\nReportsB: %s\n", reports, reports)},
		{"Report types", fmt.Sprintf("ReportsTypeA: %s\nReportsTypeB: %s\n", reportsTypes, reportsTypes)},
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
