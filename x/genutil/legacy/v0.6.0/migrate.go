package v060

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
)

// Migrate migrates exported state from v0.5.0 to a v0.6.0 genesis state.
func Migrate(appState genutil.AppMap, _ ...interface{}) genutil.AppMap {
	v050Codec := codec.New()
	codec.RegisterCrypto(v050Codec)

	v060Codec := codec.New()
	codec.RegisterCrypto(v060Codec)

	// Migrate posts state
	if appState[v040posts.ModuleName] != nil {
		var genDocs v040posts.GenesisState
		v050Codec.MustUnmarshalJSON(appState[v040posts.ModuleName], &genDocs)

		appState[v040posts.ModuleName] = v060Codec.MustMarshalJSON(
			v060posts.Migrate(genDocs),
		)
	}

	return appState
}
