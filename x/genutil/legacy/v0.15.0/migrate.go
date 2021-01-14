package v0150

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	v040 "github.com/cosmos/cosmos-sdk/x/genutil/legacy/v040"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	v0150fees "github.com/desmos-labs/desmos/x/fees/types"

	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"

	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v0150posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.15.0"
	v0130profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
	v0150profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.15.0"
	v0130relationships "github.com/desmos-labs/desmos/x/relationships/legacy/v0.13.0"
	v0150relationships "github.com/desmos-labs/desmos/x/relationships/legacy/v0.15.0"
	v0130reports "github.com/desmos-labs/desmos/x/reports/legacy/v0.13.0"
	v0150reports "github.com/desmos-labs/desmos/x/reports/legacy/v0.15.0"
)

// Migrate migrates exported state from v0.13.0 to a v0.15.0 genesis state.
func Migrate(appState genutiltypes.AppMap, cliCtx client.Context) genutiltypes.AppMap {
	v040.Migrate(appState, cliCtx)

	v0130Codec := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(v0130Codec)

	v0150Codec := cliCtx.LegacyAmino

	// Migrate posts state
	if appState[v0120posts.ModuleName] != nil {
		var genDocs v0130posts.GenesisState
		v0130Codec.MustUnmarshalJSON(appState[v0120posts.ModuleName], &genDocs)

		appState[v0150posts.ModuleName] = v0150Codec.MustMarshalJSON(v0150posts.Migrate(genDocs))
	}

	// Migrate profiles state
	if appState[v0130profiles.ModuleName] != nil {
		var genDocs v0130profiles.GenesisState
		v0130Codec.MustUnmarshalJSON(appState[v0130profiles.ModuleName], &genDocs)

		appState[v0150profiles.ModuleName] = v0150Codec.MustMarshalJSON(v0150profiles.Migrate(genDocs))
	}

	// Migrate relationships state
	if appState[v0130relationships.ModuleName] != nil {
		var genDocs v0130relationships.GenesisState
		v0130Codec.MustUnmarshalJSON(appState[v0130relationships.ModuleName], &genDocs)

		appState[v0150relationships.ModuleName] = v0150Codec.MustMarshalJSON(v0150relationships.Migrate(genDocs))
	}

	// Migrate reports state
	if appState[v0130reports.ModuleName] != nil {
		var genDocs v0130reports.GenesisState
		v0130Codec.MustUnmarshalJSON(appState[v0130reports.ModuleName], &genDocs)

		appState[v0150reports.ModuleName] = v0150Codec.MustMarshalJSON(v0150reports.Migrate(genDocs))
	}

	// Add fees if non existing
	if appState[v0150fees.ModuleName] == nil {
		appState[v0150fees.ModuleName] = v0150Codec.MustMarshalJSON(v0150fees.DefaultGenesisState())
	}

	return appState
}
