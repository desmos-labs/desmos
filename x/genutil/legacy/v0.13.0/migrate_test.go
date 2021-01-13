package v0130_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/stretchr/testify/require"
	tm "github.com/tendermint/tendermint/types"

	v0130 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.13.0"
	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
)

func TestMigrate0130(t *testing.T) {
	cdc := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(cdc)

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	// Read the genesis
	genesis, err := tm.GenesisDocFromFile("v0120state.json")
	require.NoError(t, err)

	// Read the whole app state
	var v012state genutiltypes.AppMap
	err = cdc.UnmarshalJSON(genesis.AppState, &v012state)
	require.NoError(t, err)

	// Make sure that all the posts are migrated
	var v012postsState v0120posts.GenesisState
	err = cdc.UnmarshalJSON(v012state[v0120posts.ModuleName], &v012postsState)
	require.NoError(t, err)

	// Migrate everything
	v0130state := v0130.Migrate(v012state, client.Context{})

	var v0130postsState v0130posts.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v0130posts.ModuleName], &v0130postsState)
	require.NoError(t, err)

	verifyPostsStateMigrated(t, v0130postsState, v012postsState)
}

func verifyPostsStateMigrated(
	t *testing.T, v0130postsState v0130posts.GenesisState, v012postsState v0120posts.GenesisState,
) {
	// Make sure that all the posts are migrated correctly
	require.Len(t, v0130postsState.Posts, len(v012postsState.Posts))
	for index, migrated := range v0130postsState.Posts {
		original := v012postsState.Posts[index]
		verifyMigratedPost(t, original, migrated)
	}

	// Make sure the poll answers are migrated properly
	require.Len(t, v0130postsState.UsersPollAnswers, len(v012postsState.UsersPollAnswers))
	for key, migratedValue := range v0130postsState.UsersPollAnswers {
		originalValue := v012postsState.UsersPollAnswers[key]
		for index, original := range originalValue {
			migrated := migratedValue[index]
			verifyMigratedUserAnswer(t, original, migrated)
		}
	}

	// Make sure the post reactions are migrated properly
	require.Len(t, v0130postsState.PostReactions, len(v012postsState.PostReactions))
	for key, migratedValue := range v0130postsState.PostReactions {
		originalValue := v012postsState.PostReactions[key]
		for index, original := range originalValue {
			migrated := migratedValue[index]
			verifyPostReaction(t, original, migrated)
		}
	}

	// Make sure the registered reactions are migrated properly
	require.Len(t, v0130postsState.RegisteredReactions, len(v012postsState.RegisteredReactions))
	for key, migrated := range v0130postsState.RegisteredReactions {
		original := v012postsState.RegisteredReactions[key]
		verifyRegisteredReaction(t, original, migrated)
	}

	// Make sure the params are migrated properly
	verifyParams(t, v012postsState.Params, v0130postsState.Params)
}

func verifyMigratedPost(t *testing.T, original v0120posts.Post, migrated v0130posts.Post) {
	require.Equal(t, original.PostID, migrated.PostID)
	require.Equal(t, original.ParentID, migrated.ParentID)
	require.Equal(t, original.Message, migrated.Message)
	require.True(t, original.Created.Equal(migrated.Created))
	require.True(t, original.LastEdited.Equal(migrated.LastEdited))
	require.Equal(t, original.AllowsComments, migrated.AllowsComments)
	require.Equal(t, original.Subspace, migrated.Subspace)
	require.Equal(t, original.Creator, migrated.Creator)

	// Check optional data
	if original.OptionalData == nil {
		require.Nil(t, migrated.OptionalData)
	} else {
		require.NotNil(t, migrated.OptionalData)
		i := 0
		for key, value := range original.OptionalData {
			require.Equal(t, key, migrated.OptionalData[i].Key)
			require.Equal(t, value, migrated.OptionalData[i].Value)
			i++
		}
	}

	// Check attachments
	require.Len(t, migrated.Attachments, len(original.Attachments))
	for index, originalAttachment := range original.Attachments {
		migratedAttachment := migrated.Attachments[index]

		require.Equal(t, originalAttachment.URI, migratedAttachment.URI)
		require.Equal(t, originalAttachment.MimeType, migratedAttachment.MimeType)
		require.Equal(t, originalAttachment.Tags, migratedAttachment.Tags)
	}

	// Check the poll data
	if original.PollData == nil {
		require.Nil(t, migrated.PollData)
	} else {
		originalPoll := original.PollData
		migratedPoll := migrated.PollData

		require.Equal(t, originalPoll.Question, migrated.PollData.Question)

		require.Len(t, originalPoll.ProvidedAnswers, len(migratedPoll.ProvidedAnswers))
		for index, originalAnswer := range originalPoll.ProvidedAnswers {
			migratedAnswer := migratedPoll.ProvidedAnswers[index]
			require.Equal(t, originalAnswer.ID, migratedAnswer.ID)
			require.Equal(t, originalAnswer.Text, migratedAnswer.Text)
		}

		require.Equal(t, originalPoll.EndDate, migratedPoll.EndDate)
		require.Equal(t, originalPoll.AllowsMultipleAnswers, migratedPoll.AllowsMultipleAnswers)
		require.Equal(t, originalPoll.AllowsAnswerEdits, migratedPoll.AllowsAnswerEdits)
	}
}

func verifyMigratedUserAnswer(t *testing.T, original v0120posts.UserAnswer, migrated v0130posts.UserAnswer) {
	require.Equal(t, original.Answers, migrated.Answers)
	require.True(t, original.User.Equals(migrated.User))
}

func verifyPostReaction(t *testing.T, original v0120posts.PostReaction, migrated v0130posts.PostReaction) {
	require.True(t, original.Owner.Equals(migrated.Owner))
	require.Equal(t, original.Shortcode, migrated.Shortcode)
	require.Equal(t, original.Value, migrated.Value)
}

func verifyRegisteredReaction(t *testing.T, original v0120posts.RegisteredReaction, migrated v0130posts.RegisteredReaction) {
	require.Equal(t, original.ShortCode, migrated.ShortCode)
	require.Equal(t, original.Value, migrated.Value)
	require.Equal(t, original.Subspace, migrated.Subspace)
	require.True(t, original.Creator.Equals(migrated.Creator))
}

func verifyParams(t *testing.T, original v0120posts.Params, migrated v0130posts.Params) {
	require.Equal(t, original.MaxPostMessageLength, migrated.MaxPostMessageLength)
	require.Equal(t, original.MaxOptionalDataFieldsNumber, migrated.MaxOptionalDataFieldsNumber)
	require.Equal(t, original.MaxOptionalDataFieldValueLength, migrated.MaxOptionalDataFieldValueLength)
}
