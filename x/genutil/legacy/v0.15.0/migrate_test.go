package v0150_test

/*
import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	v0150 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.15.0"
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
	"github.com/stretchr/testify/require"
	tm "github.com/tendermint/tendermint/types"
	"testing"
)

func TestMigrate0150(t *testing.T) {
	cdc := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(cdc)

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	// Read the genesis
	genesis, err := tm.GenesisDocFromFile("v0130state.json")
	require.NoError(t, err)

	// Read the whole app state
	var v0130state genutiltypes.AppMap
	err = cdc.UnmarshalJSON(genesis.AppState, &v0130state)
	require.NoError(t, err)

	// Make sure that all the posts are migrated
	var v0130postsState v0130posts.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v060posts.ModuleName], &v0130postsState)
	require.NoError(t, err)

	// Migrate everything
	v0150state := v0150.Migrate(v0130state)

	var v0150postsState v0150posts.GenesisState
	err = cdc.UnmarshalJSON(v0150state[v060posts.ModuleName], &v0150postsState)
	require.NoError(t, err)
	require.Len(t, v0130postsState.Posts, len(v0150postsState.Posts))
	require.Len(t, v0130postsState.UsersPollAnswers, len(v0150postsState.UsersPollAnswers))
	require.Len(t, v0130postsState.PostReactions, len(v0150postsState.PostReactions))
	require.Len(t, v0130postsState.RegisteredReactions, len(v0130postsState.RegisteredReactions))

	// Make sure that all the profiles are migrated correctly
	var v0130profilesState v0130profiles.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v080profiles.ModuleName], &v0130profilesState)
	require.NoError(t, err)

	var v0150profilesState v0150profiles.GenesisState
	err = cdc.UnmarshalJSON(v0150state[v080profiles.ModuleName], &v0150profilesState)
	require.NoError(t, err)
	require.Len(t, v0130profilesState.Profiles, len(v0150profilesState.Profiles))

	// Make sure that all the profiles are migrated correctly
	for index, profile := range v0150profilesState.Profiles {
		if profile.Moniker == "" {
			require.Nil(t, v0130profilesState.Profiles[index].Moniker)
		}
		if profile.Bio == "" {
			require.Nil(t, v0130profilesState.Profiles[index].Bio)
		}

		if profile.Pictures.Profile == "" && profile.Pictures.Cover == "" {
			require.Nil(t, v0130profilesState.Profiles[index].Pictures)
		}
	}

	require.Len(t, v0130profilesState.DTagTransferRequests, len(v0150profilesState.DtagTransferRequests))

	// Make sure that all the relationships are migrated correctly
	var v0130relationshipsState v0130relationships.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v0130relationships.ModuleName], &v0130relationshipsState)
	require.NoError(t, err)

	var v0150relationshipsState v0150relationships.GenesisState
	err = cdc.UnmarshalJSON(v0150state[v0130relationships.ModuleName], &v0150relationshipsState)
	require.NoError(t, err)
	require.Len(t, v0130relationshipsState.UsersRelationships, len(v0150relationshipsState.Relationships))
	require.Len(t, v0130relationshipsState.UsersBlocks, len(v0150relationshipsState.Blocks))

	// Make sure that all the reports are migrated correctly
	var v0130reportsState v0130reports.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v0130relationships.ModuleName], &v0130reportsState)
	require.NoError(t, err)

	var v0150reportsState v0150reports.GenesisState
	err = cdc.UnmarshalJSON(v0150state[v0130reports.ModuleName], &v0150reportsState)
	require.NoError(t, err)
	require.Len(t, v0130reportsState.Reports, len(v0150reportsState.Reports))
}

*/
