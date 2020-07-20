package v090

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	v080posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
	v090posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.9.0"
	"time"
)

// Migrate migrates exported state from v0.8.0 to a v0.9.0 genesis state.
func Migrate(appState genutil.AppMap, values ...interface{}) genutil.AppMap {
	v080Codec := codec.New()
	codec.RegisterCrypto(v080Codec)

	v090Codec := codec.New()
	codec.RegisterCrypto(v090Codec)

	genesisTime, ok := values[0].(time.Time)
	if !ok || genesisTime.IsZero() {
		panic("no genesis time provided")
	}

	// Migrate posts state
	if appState[v060posts.ModuleName] != nil {
		var genDocs v080posts.GenesisState
		v090Codec.MustUnmarshalJSON(appState[v060posts.ModuleName], &genDocs)

		appState[v060posts.ModuleName] = v090Codec.MustMarshalJSON(
			v090posts.Migrate(genDocs),
		)
	}

	return appState
}
