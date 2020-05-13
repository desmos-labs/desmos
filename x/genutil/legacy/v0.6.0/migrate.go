package v060

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v040profiles "github.com/desmos-labs/desmos/x/profile/legacy/v0.4.0"
	v060profiles "github.com/desmos-labs/desmos/x/profile/legacy/v0.6.0"
)

// migrateProfilesModule migrates the state of profiles from v0.4.0 to a v0.6.0 genesis state.
func migrateProfilesModule(cdc *codec.Codec, appState genutil.AppMap) genutil.AppMap {
	v040Codec := codec.New()
	codec.RegisterCrypto(v040Codec)

	// Migrate profile state
	if appState[v040profiles.ModuleName] != nil {
		var genDocs v040profiles.GenesisState
		v040Codec.MustUnmarshalJSON(appState[v040profiles.ModuleName], &genDocs)

		delete(appState, v040profiles.ModuleName)
		appState[v060profiles.ModuleName] = cdc.MustMarshalJSON(
			genDocs,
		)
	}

	return appState
}

// Migrate migrates exported state from v0.5.0 to a v0.6.0 genesis state.
func Migrate(appState genutil.AppMap, _ ...interface{}) genutil.AppMap {
	v060Codec := codec.New()
	codec.RegisterCrypto(v060Codec)

	appState = migrateProfilesModule(v060Codec, appState)

	return appState
}
