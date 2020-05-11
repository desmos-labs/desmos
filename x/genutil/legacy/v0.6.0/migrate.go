package v060

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v040profiles "github.com/desmos-labs/desmos/x/profile/legacy/v0.4.0"
	v060profiles "github.com/desmos-labs/desmos/x/profile/legacy/v0.6.0"
)

// Migrate migrates exported state from v0.4.0 to a v0.6.0 genesis state.
func Migrate(appState genutil.AppMap, _ ...interface{}) genutil.AppMap {
	v040CodecProfiles := codec.New()
	codec.RegisterCrypto(v040CodecProfiles)

	v060Codec := codec.New()
	codec.RegisterCrypto(v060Codec)

	// Migrate profile state
	if appState[v040profiles.ModuleName] != nil {
		var genDocs v040profiles.GenesisState
		v040CodecProfiles.MustUnmarshalJSON(appState[v040profiles.ModuleName], &genDocs)

		delete(appState, v040profiles.ModuleName)
		appState[v060profiles.ModuleName] = v040CodecProfiles.MustMarshalJSON(
			genDocs,
		)
	}

	return appState
}
