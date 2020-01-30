package v0_3_0

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	cosmosv0380 "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v0_38"
	v020posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.2.0"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
)

// Migrate migrates exported state from v0.1.0 to a v0.2.0 genesis state.
// It requires args to contain an integer value representing the block interval that should be considered when
// converting block height-based timestamps into time.Time timestamps.
func Migrate(appState genutil.AppMap, _ ...interface{}) genutil.AppMap {
	// Perform the Cosmos SDK migration first
	appState = cosmosv0380.Migrate(appState)

	v010Codec := codec.New()
	codec.RegisterCrypto(v010Codec)

	v020Codec := codec.New()
	codec.RegisterCrypto(v020Codec)

	// Migrate posts state
	if appState[v020posts.ModuleName] != nil {
		var genDocs v020posts.GenesisState
		v010Codec.MustUnmarshalJSON(appState[v020posts.ModuleName], &genDocs)

		delete(appState, v020posts.ModuleName) // delete old key in case the name changed
		appState[v030posts.ModuleName] = v020Codec.MustMarshalJSON(
			v030posts.Migrate(genDocs),
		)
	}

	return appState
}
