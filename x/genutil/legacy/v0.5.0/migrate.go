package v050

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v050posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.5.0"
)

// Migrate migrates exported state from v0.4.0 to a v0.5.0 genesis state.
func Migrate(appState genutil.AppMap, _ ...interface{}) genutil.AppMap {
	v040Codec := codec.New()
	codec.RegisterCrypto(v040Codec)

	v050Codec := codec.New()
	codec.RegisterCrypto(v050Codec)

	// Migrate posts state
	if appState[v040posts.ModuleName] != nil {
		var genDocs v040posts.GenesisState
		v040Codec.MustUnmarshalJSON(appState[v040posts.ModuleName], &genDocs)

		appState[v040posts.ModuleName] = v050Codec.MustMarshalJSON(
			v050posts.Migrate(genDocs),
		)
	}

	return appState
}
