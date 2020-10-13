package v0130_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	v0130 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.13.0"
	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	"github.com/stretchr/testify/require"
	tm "github.com/tendermint/tendermint/types"
	"reflect"
	"testing"
)

func TestMigrate0120(t *testing.T) {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	// Read the genesis
	genesis, err := tm.GenesisDocFromFile("v0120state.json")
	require.NoError(t, err)

	// Read the whole app state
	var v012state genutil.AppMap
	err = cdc.UnmarshalJSON(genesis.AppState, &v012state)
	require.NoError(t, err)

	// Migrate everything
	v0130state := v0130.Migrate(v012state)

	// Make sure that all the posts are migrated
	var v012postsState v0120posts.GenesisState
	err = cdc.UnmarshalJSON(v012state[v060posts.ModuleName], &v012postsState)
	//require.NoError(t, err)

	var v0130postsState v0130posts.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v060posts.ModuleName], &v0130postsState)
	require.NoError(t, err)
	require.Equal(t, len(v012postsState.Posts), len(v0130postsState.Posts))

	// Make sure that all the posts' polls are migrated correctly
	for _, post := range v0130postsState.Posts {
		require.Equal(t, "OptionalData", reflect.TypeOf(post).String())
	}
}
