package v080

import (
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/desmos-labs/desmos/x/reports"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	v080posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
	v060profile "github.com/desmos-labs/desmos/x/profiles/legacy/v0.6.0"
	v080profile "github.com/desmos-labs/desmos/x/profiles/legacy/v0.8.0"
)

// Migrate migrates exported state from v0.6.0 to a v0.8.0 genesis state.
func Migrate(appState genutil.AppMap, values ...interface{}) genutil.AppMap {
	v060Codec := codec.New()
	codec.RegisterCrypto(v060Codec)

	v080Codec := codec.New()
	evidence.RegisterCodec(v080Codec)
	codec.RegisterCrypto(v080Codec)

	genesisTime, ok := values[0].(time.Time)
	if !ok || genesisTime.IsZero() {
		panic("no genesis time provided")
	}

	// Add default modules
	appState[crisis.ModuleName] = v080Codec.MustMarshalJSON(crisis.DefaultGenesisState())
	appState[evidence.ModuleName] = v080Codec.MustMarshalJSON(evidence.DefaultGenesisState())
	appState[gov.ModuleName] = v080Codec.MustMarshalJSON(gov.DefaultGenesisState())
	appState[reports.ModuleName] = v080Codec.MustMarshalJSON(reports.DefaultGenesisState())

	// Migrate posts state
	if appState[v060posts.ModuleName] != nil {
		var genDocs v060posts.GenesisState
		v080Codec.MustUnmarshalJSON(appState[v060posts.ModuleName], &genDocs)

		appState[v060posts.ModuleName] = v080Codec.MustMarshalJSON(
			v080posts.Migrate(genDocs),
		)
	}

	// Migrate profile state
	if appState[v060profile.ModuleName] != nil {
		var genDocs v060profile.GenesisState
		v060Codec.MustUnmarshalJSON(appState[v060profile.ModuleName], &genDocs)

		delete(appState, v060profile.ModuleName) // delete old key in case the name changed
		appState[v080profile.ModuleName] = v080Codec.MustMarshalJSON(
			v080profile.Migrate(genDocs, genesisTime),
		)
	}

	return appState
}
