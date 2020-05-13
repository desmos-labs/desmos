package v060

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"

// Migrate migrates exported state from v0.5.0 to a v0.6.0 genesis state.
func Migrate(appState genutil.AppMap, _ ...interface{}) genutil.AppMap {
  v050Codec := codec.New()
	codec.RegisterCrypto(v040Codec)
  
	v060Codec := codec.New()
	codec.RegisterCrypto(v060Codec)

  // Migrate posts state
	if appState[v040posts.ModuleName] != nil {
		var genDocs v040posts.GenesisState
		v050Codec.MustUnmarshalJSON(appState[v040posts.ModuleName], &genDocs)

		appState[v040posts.ModuleName] = cdc.MustMarshalJSON(
			v060posts.Migrate(genDocs),
		)
  }
  
  // Migrate profile state
	if appState[v040profiles.ModuleName] != nil {
		var genDocs v040profiles.GenesisState
		v050Codec.MustUnmarshalJSON(appState[v040profiles.ModuleName], &genDocs)

		delete(appState, v040profiles.ModuleName)
		appState[v060profiles.ModuleName] = cdc.MustMarshalJSON(
			genDocs,
		)
	}
  
	return appState
}
