package v0150

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v0150posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.15.0"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	v0130profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
	v0150profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.15.0"
	v080profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.8.0"
	v0130relationships "github.com/desmos-labs/desmos/x/relationships/legacy/v0.13.0"
	v0150relationships "github.com/desmos-labs/desmos/x/relationships/legacy/v0.15.0"
	v0130reports "github.com/desmos-labs/desmos/x/reports/legacy/v0.13.0"
	v0150reports "github.com/desmos-labs/desmos/x/reports/legacy/v0.15.0"
)

// Migrate migrates exported state from v0.13.0 to a v0.15.0 genesis state.
func Migrate(appState genutiltypes.AppMap, values ...interface{}) genutiltypes.AppMap {
	v0130Codec := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(v0130Codec)

	v0150Codec := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(v0150Codec)

	// Migrate posts state
	if appState[v060posts.ModuleName] != nil {
		var genDocs v0130posts.GenesisState
		v0130Codec.MustUnmarshalJSON(appState[v060posts.ModuleName], &genDocs)

		appState[v060posts.ModuleName] = v0150Codec.MustMarshalJSON(
			v0150posts.Migrate(genDocs),
		)
	}

	// Migrate profiles state
	if appState[v080profiles.ModuleName] != nil {
		var genDocs v0130profiles.GenesisState
		v0130Codec.MustUnmarshalJSON(appState[v080profiles.ModuleName], &genDocs)

		appState[v080profiles.ModuleName] = v0150Codec.MustMarshalJSON(
			v0150profiles.Migrate(genDocs),
		)
	}

	// Migrate relationships state
	if appState[v0130relationships.ModuleName] != nil {
		var genDocs v0130relationships.GenesisState
		v0130Codec.MustUnmarshalJSON(appState[v0130relationships.ModuleName], &genDocs)

		appState[v0130relationships.ModuleName] = v0150Codec.MustMarshalJSON(
			v0150relationships.Migrate(genDocs),
		)
	}

	// Migrate reports state
	if appState[v0130reports.ModuleName] != nil {
		var genDocs v0130reports.GenesisState
		v0130Codec.MustUnmarshalJSON(appState[v0130reports.ModuleName], &genDocs)

		appState[v0130reports.ModuleName] = v0150Codec.MustMarshalJSON(
			v0150reports.Migrate(genDocs),
		)
	}

	return appState
}
