package v0130

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
)

// Migrate migrates exported state from v0.12.0 to a v0.13.0 genesis state.
func Migrate(appState genutiltypes.AppMap, values ...interface{}) genutiltypes.AppMap {
	v0120Codec := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(v0120Codec)

	v0130Codec := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(v0130Codec)

	// Migrate posts state
	if appState[v060posts.ModuleName] != nil {
		var genDocs v0120posts.GenesisState
		v0120Codec.MustUnmarshalJSON(appState[v060posts.ModuleName], &genDocs)

		appState[v060posts.ModuleName] = v0130Codec.MustMarshalJSON(
			v0130posts.Migrate(genDocs),
		)
	}

	return appState
}
