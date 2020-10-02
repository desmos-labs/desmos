package v0120

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"

	"github.com/desmos-labs/desmos/x/relationships"
	relationshipsTypes "github.com/desmos-labs/desmos/x/relationships/types"

	v0100posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.10.0"
	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
)

// Migrate migrates exported state from v0.10.0 to a v0.12.0 genesis state.
func Migrate(appState genutil.AppMap, values ...interface{}) genutil.AppMap {
	v0100Codec := codec.New()
	codec.RegisterCrypto(v0100Codec)

	v0120Codec := codec.New()
	codec.RegisterCrypto(v0120Codec)

	// Migrate posts state
	if appState[v060posts.ModuleName] != nil {
		var genDocs v0100posts.GenesisState
		v0120Codec.MustUnmarshalJSON(appState[v060posts.ModuleName], &genDocs)

		appState[v060posts.ModuleName] = v0120Codec.MustMarshalJSON(
			v0120posts.Migrate(genDocs),
		)
	}

	// Add the relationships default state
	if appState[relationships.ModuleName] == nil {
		appState[relationships.ModuleName] = v0120Codec.MustMarshalJSON(
			relationshipsTypes.DefaultGenesisState(),
		)
	}

	return appState
}
