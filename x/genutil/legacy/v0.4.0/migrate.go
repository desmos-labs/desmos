package v040

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
)

// Migrate migrates exported state from v0.3.0 to a v0.4.0 genesis state.
// It requires args to contain an integer value representing the block interval that should be considered when
// converting block height-based timestamps into time.Time timestamps.
func Migrate(appState genutil.AppMap, _ ...interface{}) genutil.AppMap {
	v030Codec := codec.New()
	codec.RegisterCrypto(v030Codec)

	v040Codec := codec.New()
	codec.RegisterCrypto(v040Codec)

	// Migrate posts state
	if appState[v030posts.ModuleName] != nil {
		var genDocs v030posts.GenesisState
		v030Codec.MustUnmarshalJSON(appState[v030posts.ModuleName], &genDocs)

		delete(appState, v030posts.ModuleName) // delete old key in case the name changed
		appState[v040posts.ModuleName] = v040Codec.MustMarshalJSON(
			v040posts.Migrate(genDocs),
		)
	}

	return appState
}
