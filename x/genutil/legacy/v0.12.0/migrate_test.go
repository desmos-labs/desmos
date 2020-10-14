package v0120_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/stretchr/testify/require"
	tm "github.com/tendermint/tendermint/types"

	v0120 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.12.0"
	v0100posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.10.0"
	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
	v060posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	"github.com/desmos-labs/desmos/x/relationships"
)

func TestMigrate0100(t *testing.T) {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	// Read the genesis
	genesis, err := tm.GenesisDocFromFile("v0100state.json")
	require.NoError(t, err)

	// Read the whole app state
	var v010state genutil.AppMap
	err = cdc.UnmarshalJSON(genesis.AppState, &v010state)
	require.NoError(t, err)

	// Make sure that all the posts are migrated
	var v010postsState v0100posts.GenesisState
	err = cdc.UnmarshalJSON(v010state[v060posts.ModuleName], &v010postsState)

	// Migrate everything
	v0120state := v0120.Migrate(v010state)

	var v0120postsState v0120posts.GenesisState
	err = cdc.UnmarshalJSON(v0120state[v060posts.ModuleName], &v0120postsState)
	require.Equal(t, len(v0120postsState.Posts), len(v010postsState.Posts))

	// Make sure that all the posts' polls are migrated correctly
	for index, post := range v010postsState.Posts {
		if post.PollData == nil {
			require.Nil(t, v010postsState.Posts[index].PollData)
		} else {
			require.NotNil(t, v010postsState.Posts[index].PollData)
		}
	}

	// Make sure the relationships genesis state is not nil
	require.NotNil(t, v0120state[relationships.ModuleName])
}
