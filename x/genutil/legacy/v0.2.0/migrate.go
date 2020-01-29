package v020

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/genutil"

	v010posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.1.0"
	v020posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.2.0"
)

// Migrate migrates exported state from v0.1.0 to a v0.2.0 genesis state.
// It requires args to contain an integer value representing the block interval that should be considered when
// converting block height-based timestamps into time.Time timestamps.
func Migrate(appState genutil.AppMap, args ...interface{}) genutil.AppMap {
	v010Codec := codec.New()
	codec.RegisterCrypto(v010Codec)

	v020Codec := codec.New()
	codec.RegisterCrypto(v020Codec)

	// Migrate posts state
	if appState[v010posts.ModuleName] != nil {
		var genDocs v010posts.GenesisState
		v010Codec.MustUnmarshalJSON(appState[v010posts.ModuleName], &genDocs)

		delete(appState, v010posts.ModuleName) // delete old key in case the name changed

		// Get genesis time and block interval
		genesisTime := args[0].(time.Time)
		blockInterval := args[1].(int)
		if blockInterval == 0 {
			panic(fmt.Errorf("block interval cannot be 0 when migrating to v0.2.0, please set it using the --block-interval flag"))
		}

		appState[v020posts.ModuleName] = v020Codec.MustMarshalJSON(
			v020posts.Migrate(genDocs, genesisTime, blockInterval),
		)
	}

	return appState
}
