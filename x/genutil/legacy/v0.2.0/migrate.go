package v0_2_0

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"

	v010posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.1.0"
	v020posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.2.0"
)

// Migrate migrates exported state from v0.1.0 to a v0.2.0 genesis state.
func Migrate(appState genutil.AppMap) genutil.AppMap {
	v010Codec := codec.New()
	codec.RegisterCrypto(v010Codec)

	v020Codec := codec.New()
	codec.RegisterCrypto(v020Codec)

	// Migrate docs state
	if appState[v010posts.ModuleName] != nil {
		var genDocs v010posts.GenesisState
		v010Codec.MustUnmarshalJSON(appState[v010posts.ModuleName], &genDocs)

		delete(appState, v010posts.ModuleName) // delete old key in case the name changed
		appState[v020posts.ModuleName] = v020Codec.MustMarshalJSON(
			v020posts.Migrate(genDocs),
		)
	}

	return appState
}
