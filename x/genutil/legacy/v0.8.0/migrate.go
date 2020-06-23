package v080

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v060profile "github.com/desmos-labs/desmos/x/profile/legacy/v0.6.0"
	v080profile "github.com/desmos-labs/desmos/x/profile/legacy/v0.8.0"
)

// Migrate migrates exported state from v0.6.0 to a v0.8.0 genesis state.
func Migrate(appState genutil.AppMap, values ...interface{}) genutil.AppMap {
	v050Codec := codec.New()
	codec.RegisterCrypto(v050Codec)

	v060Codec := codec.New()
	codec.RegisterCrypto(v060Codec)

	genesisTime, ok := values[0].(time.Time)
	if !ok || genesisTime.IsZero() {
		panic("no genesis time provided")
	}

	// Migrate posts state
	if appState[v060profile.ModuleName] != nil {
		var genDocs v060profile.GenesisState
		v050Codec.MustUnmarshalJSON(appState[v060profile.ModuleName], &genDocs)

		appState[v080profile.ModuleName] = v060Codec.MustMarshalJSON(
			v080profile.Migrate(genDocs, genesisTime),
		)
	}

	return appState
}
