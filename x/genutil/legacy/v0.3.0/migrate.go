package v030

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	cosmosv0380 "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v0_38"
	v020magpie "github.com/desmos-labs/desmos/x/magpie/legacy/v0.2.0"
	v030magpie "github.com/desmos-labs/desmos/x/magpie/legacy/v0.3.0"
	v020posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.2.0"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
)

// Migrate migrates exported state from v0.1.0 to a v0.2.0 genesis state.
// It requires args to contain an integer value representing the block interval that should be considered when
// converting block height-based timestamps into time.Time timestamps.
func Migrate(appState genutil.AppMap, _ ...interface{}) genutil.AppMap {
	// Perform the Cosmos SDK migration first
	appState = cosmosv0380.Migrate(appState)

	v020Codec := codec.New()
	codec.RegisterCrypto(v020Codec)

	v030Codec := codec.New()
	codec.RegisterCrypto(v030Codec)

	// Migrate magpie state
	if appState[v020magpie.ModuleName] != nil {
		var genDocs v020magpie.GenesisState
		v020Codec.MustUnmarshalJSON(appState[v020magpie.ModuleName], &genDocs)

		delete(appState, v020magpie.ModuleName) // delete old key in case the name changed
		appState[v030magpie.ModuleName] = v030Codec.MustMarshalJSON(
			v030magpie.Migrate(genDocs),
		)
	}

	// Migrate posts state
	if appState[v020posts.ModuleName] != nil {
		var genDocs v020posts.GenesisState
		v020Codec.MustUnmarshalJSON(appState[v020posts.ModuleName], &genDocs)

		delete(appState, v020posts.ModuleName) // delete old key in case the name changed
		appState[v030posts.ModuleName] = v030Codec.MustMarshalJSON(
			v030posts.Migrate(genDocs),
		)
	}

	return appState
}
