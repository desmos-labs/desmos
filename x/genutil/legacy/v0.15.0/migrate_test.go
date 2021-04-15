package v0150_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/stretchr/testify/require"
	tm "github.com/tendermint/tendermint/types"

	"github.com/desmos-labs/desmos/app"
	v0150 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.15.0"
	v0130profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.13.0"
	v0150profiles "github.com/desmos-labs/desmos/x/profiles/legacy/v0.15.0"
	v0130posts "github.com/desmos-labs/desmos/x/staging/posts/legacy/v0.13.0"
	v0150posts "github.com/desmos-labs/desmos/x/staging/posts/legacy/v0.15.0"
	v0130relationships "github.com/desmos-labs/desmos/x/staging/relationships/legacy/v0.13.0"
	v0150relationships "github.com/desmos-labs/desmos/x/staging/relationships/legacy/v0.15.0"
	v0130reports "github.com/desmos-labs/desmos/x/staging/reports/legacy/v0.13.0"
	v0150reports "github.com/desmos-labs/desmos/x/staging/reports/legacy/v0.15.0"
)

func TestMigrate0150(t *testing.T) {
	encodingConfig := app.MakeTestEncodingConfig()
	clientCtx := client.Context{}.
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithJSONMarshaler(encodingConfig.Marshaler)

	cdc := codec.NewLegacyAmino()
	cryptocodec.RegisterCrypto(cdc)

	config := sdk.GetConfig()
	app.SetupConfig(config)
	config.Seal()

	// ---------------------------------------------------------------

	// Read the genesis
	genesis, err := tm.GenesisDocFromFile("v0130state.json")
	require.NoError(t, err)

	// Read the whole app state
	var v0130state genutiltypes.AppMap
	err = cdc.UnmarshalJSON(genesis.AppState, &v0130state)
	require.NoError(t, err)

	// Deserialize the various genesis states

	var v0130postsState v0130posts.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v0130posts.ModuleName], &v0130postsState)
	require.NoError(t, err)

	var v0130profilesState v0130profiles.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v0130profiles.ModuleName], &v0130profilesState)
	require.NoError(t, err)

	var v0130relationshipsState v0130relationships.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v0130relationships.ModuleName], &v0130relationshipsState)
	require.NoError(t, err)

	var v0130reportsState v0130reports.GenesisState
	err = cdc.UnmarshalJSON(v0130state[v0130reports.ModuleName], &v0130reportsState)
	require.NoError(t, err)

	// ---------------------------------------------------------------

	// Migrate everything
	v0150state := v0150.Migrate(v0130state, clientCtx)

	// Deserialize the various genesis states

	var v0150postsState v0150posts.GenesisState
	err = cdc.UnmarshalJSON(v0150state[v0150posts.ModuleName], &v0150postsState)
	require.NoError(t, err)

	var v0150profilesState v0150profiles.GenesisState
	err = cdc.UnmarshalJSON(v0150state[v0150profiles.ModuleName], &v0150profilesState)
	require.NoError(t, err)

	var v0150relationshipsState v0150relationships.GenesisState
	err = cdc.UnmarshalJSON(v0150state[v0150relationships.ModuleName], &v0150relationshipsState)
	require.NoError(t, err)

	var v0150reportsState v0150reports.GenesisState
	err = cdc.UnmarshalJSON(v0150state[v0150reports.ModuleName], &v0150reportsState)
	require.NoError(t, err)

	// ---------------------------------------------------------------

	// Verify the migrations
	verifyPostsStateMigrated(t, v0130postsState, v0150postsState)
	verifyProfilesStateMigrated(t, v0130profilesState, v0150profilesState)
	verifyRelationshipsStateMigrated(t, v0130relationshipsState, v0150relationshipsState)
	verifyReportsStateMigrated(t, v0130reportsState, v0150reportsState)
}

// -----------------------------------------------------------------------------------------------------------------

func verifyPostsStateMigrated(
	t *testing.T, original v0130posts.GenesisState, migrated v0150posts.GenesisState,
) {
	// Make sure that all the posts are migrated correctly
	require.Len(t, migrated.Posts, len(original.Posts))
	for index, migrated := range migrated.Posts {
		original := original.Posts[index]
		verifyMigratedPost(t, original, migrated)
	}

	// Make sure the poll answers are migrated properly
	for postID, originalValue := range original.UsersPollAnswers {
		found, migratedEntry := v0150posts.FindUserAnswerEntryForPostID(migrated, postID)
		require.True(t, found)

		for index, original := range originalValue {
			migrated := migratedEntry.UserAnswers[index]
			verifyMigratedUserAnswer(t, original, migrated)
		}
	}

	// Make sure the post reactions are migrated properly
	require.Len(t, migrated.PostsReactions, len(original.PostReactions))
	for postID, originalValue := range original.PostReactions {
		found, migratedEntry := v0150posts.FindPostReactionEntryForPostID(migrated, postID)
		require.True(t, found)

		for index, original := range originalValue {
			migrated := migratedEntry.Reactions[index]
			verifyMigratedPostReaction(t, original, migrated)
		}
	}

	// Make sure the registered reactions are migrated properly
	require.Len(t, migrated.RegisteredReactions, len(original.RegisteredReactions))
	for key, migrated := range migrated.RegisteredReactions {
		original := original.RegisteredReactions[key]
		verifyMigratedRegisteredReaction(t, original, migrated)
	}

	// Make sure the params are migrated properly
	verifyMigratedPostsParams(t, original.Params, migrated.Params)
}

func verifyMigratedPost(t *testing.T, original v0130posts.Post, migrated v0150posts.Post) {
	require.Equal(t, original.PostId, migrated.PostId)
	require.Equal(t, original.ParentId, migrated.ParentId)
	require.Equal(t, original.Message, migrated.Message)
	require.True(t, original.Created.Equal(migrated.Created))
	require.True(t, original.LastEdited.Equal(migrated.LastEdited))
	require.Equal(t, original.AllowsComments, migrated.AllowsComments)
	require.Equal(t, original.Subspace, migrated.Subspace)
	require.Equal(t, original.Creator.String(), migrated.Creator)

	// Check optional data
	if original.OptionalData == nil {
		require.Nil(t, migrated.OptionalData)
	} else {
		require.NotNil(t, migrated.OptionalData)
		for index, value := range original.OptionalData {
			require.Equal(t, value.Key, migrated.OptionalData[index].Key)
			require.Equal(t, value.Value, migrated.OptionalData[index].Value)
		}
	}

	// Check attachments
	require.Len(t, migrated.Attachments, len(original.Attachments))
	for index, originalAttachment := range original.Attachments {
		migratedAttachment := migrated.Attachments[index]

		require.Equal(t, originalAttachment.URI, migratedAttachment.URI)
		require.Equal(t, originalAttachment.MimeType, migratedAttachment.MimeType)

		require.Len(t, migratedAttachment.Tags, len(originalAttachment.Tags))
		for index, originalTag := range originalAttachment.Tags {
			require.Equal(t, originalTag.String(), migratedAttachment.Tags[index])
		}
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
			require.Equal(t, strconv.FormatUint(originalAnswer.ID, 10), migratedAnswer.ID)
			require.Equal(t, originalAnswer.Text, migratedAnswer.Text)
		}

		require.Equal(t, originalPoll.EndDate, migratedPoll.EndDate)
		require.Equal(t, originalPoll.AllowsMultipleAnswers, migratedPoll.AllowsMultipleAnswers)
		require.Equal(t, originalPoll.AllowsAnswerEdits, migratedPoll.AllowsAnswerEdits)
	}
}

func verifyMigratedUserAnswer(t *testing.T, original v0130posts.UserAnswer, migrated v0150posts.UserAnswer) {
	require.Len(t, migrated.Answers, len(original.Answers))
	for index, originalAnswer := range original.Answers {
		require.Equal(t, strconv.FormatUint(originalAnswer, 10), migrated.Answers[index])
	}

	require.Equal(t, original.User.String(), migrated.User)
}

func verifyMigratedPostReaction(t *testing.T, original v0130posts.PostReaction, migrated v0150posts.PostReaction) {
	require.Equal(t, original.Owner.String(), migrated.Owner)
	require.Equal(t, original.Shortcode, migrated.ShortCode)
	require.Equal(t, original.Value, migrated.Value)
}

func verifyMigratedRegisteredReaction(t *testing.T, original v0130posts.RegisteredReaction, migrated v0150posts.RegisteredReaction) {
	require.Equal(t, original.ShortCode, migrated.ShortCode)
	require.Equal(t, original.Value, migrated.Value)
	require.Equal(t, original.Subspace, migrated.Subspace)
	require.Equal(t, original.Creator.String(), migrated.Creator)
}

func verifyMigratedPostsParams(t *testing.T, original v0130posts.Params, migrated v0150posts.Params) {
	require.Equal(t, original.MaxPostMessageLength, migrated.MaxPostMessageLength)
	require.Equal(t, original.MaxOptionalDataFieldsNumber, migrated.MaxOptionalDataFieldsNumber)
	require.Equal(t, original.MaxOptionalDataFieldValueLength, migrated.MaxOptionalDataFieldValueLength)
}

// -----------------------------------------------------------------------------------------------------------------

func verifyProfilesStateMigrated(
	t *testing.T, original v0130profiles.GenesisState, migrated v0150profiles.GenesisState,
) {
	require.Len(t, migrated.Profiles, len(original.Profiles))
	for index, originalProfile := range original.Profiles {
		migratedProfile := migrated.Profiles[index]
		verifyMigratedProfile(t, originalProfile, migratedProfile)
	}

	require.Len(t, migrated.DtagTransferRequests, len(original.DTagTransferRequests))
	for index, originalRequest := range original.DTagTransferRequests {
		migrated := migrated.DtagTransferRequests[index]
		verifyMigratedDTagTransferRequest(t, originalRequest, migrated)
	}

	verifyMigratedProfilesParams(t, original.Params, migrated.Params)
}

func verifyMigratedProfile(t *testing.T, original v0130profiles.Profile, migrated v0150profiles.Profile) {
	require.Equal(t, original.DTag, migrated.Dtag)

	if original.Moniker == nil {
		require.Empty(t, migrated.Moniker)
	} else {
		require.Equal(t, *original.Moniker, migrated.Moniker)
	}

	if original.Bio == nil {
		require.Empty(t, migrated.Bio)
	} else {
		require.Equal(t, *original.Bio, migrated.Bio)
	}

	if original.Pictures == nil {
		require.Equal(t, "", migrated.Pictures.Profile)
		require.Equal(t, "", migrated.Pictures.Cover)
	} else {
		if original.Pictures.Profile == nil {
			require.Empty(t, migrated.Pictures.Profile)
		} else {
			require.Equal(t, *original.Pictures.Profile, migrated.Pictures.Profile)
		}

		if original.Pictures.Cover == nil {
			require.Empty(t, migrated.Pictures.Cover)
		} else {
			require.Equal(t, *original.Pictures.Cover, migrated.Pictures.Cover)
		}
	}

	require.Equal(t, original.Creator.String(), migrated.Creator)
	require.True(t, original.CreationDate.Equal(migrated.CreationDate))
}

func verifyMigratedDTagTransferRequest(t *testing.T, original v0130profiles.DTagTransferRequest, migrated v0150profiles.DTagTransferRequest) {
	require.Equal(t, original.DTagToTrade, migrated.DtagToTrade)
	require.Equal(t, original.Receiver.String(), migrated.Receiver)
	require.Equal(t, original.Sender.String(), migrated.Sender)
}

func verifyMigratedProfilesParams(t *testing.T, original v0130profiles.Params, migrated v0150profiles.Params) {
	require.Equal(t, original.MonikerParams.MinMonikerLen, migrated.MonikerParams.MinMonikerLength)
	require.Equal(t, original.MonikerParams.MaxMonikerLen, migrated.MonikerParams.MaxMonikerLength)

	require.Equal(t, original.DtagParams.RegEx, migrated.DtagParams.RegEx)
	require.Equal(t, original.DtagParams.MinDtagLen, migrated.DtagParams.MinDtagLength)
	require.Equal(t, original.DtagParams.MaxDtagLen, migrated.DtagParams.MaxDtagLength)

	require.Equal(t, original.MaxBioLen, migrated.MaxBioLength)
}

// -----------------------------------------------------------------------------------------------------------------

func verifyRelationshipsStateMigrated(
	t *testing.T, original v0130relationships.GenesisState, migrated v0150relationships.GenesisState,
) {
	for user, originalRelationships := range original.UsersRelationships {
		migratedRelationships := v0150relationships.FindRelationshipsForUser(migrated, user)

		require.Len(t, migratedRelationships, len(originalRelationships))
		for index, originalRelationship := range originalRelationships {
			migratedRelationship := migratedRelationships[index]
			verifyMigratedRelationship(t, user, originalRelationship, migratedRelationship)
		}
	}

	require.Len(t, migrated.Blocks, len(original.UsersBlocks))
	for index, originalBlock := range original.UsersBlocks {
		migratedBlock := migrated.Blocks[index]
		verifyMigratedUserBlock(t, originalBlock, migratedBlock)
	}
}

func verifyMigratedRelationship(t *testing.T, user string, original v0130relationships.Relationship, migrated v0150relationships.Relationship) {
	require.Equal(t, user, migrated.Creator)
	require.Equal(t, original.Recipient.String(), migrated.Recipient)
	require.Equal(t, original.Subspace, migrated.Subspace)
}

func verifyMigratedUserBlock(t *testing.T, original v0130relationships.UserBlock, migrated v0150relationships.UserBlock) {
	require.Equal(t, original.Blocker.String(), migrated.Blocker)
	require.Equal(t, original.Blocked.String(), migrated.Blocked)
	require.Equal(t, original.Reason, migrated.Reason)
	require.Equal(t, original.Subspace, migrated.Subspace)
}

// -----------------------------------------------------------------------------------------------------------------

func verifyReportsStateMigrated(
	t *testing.T, original v0130reports.GenesisState, migrated v0150reports.GenesisState,
) {
	for postID, originalReports := range original.Reports {
		migratedReports := v0150reports.FindReportsForPostWithID(migrated, postID)

		require.Len(t, migratedReports, len(originalReports))
		for index, originalReport := range originalReports {
			migratedReport := migratedReports[index]
			verifyMigratedReport(t, postID, originalReport, migratedReport)
		}
	}
}

func verifyMigratedReport(t *testing.T, postID string, original v0130reports.Report, migrated v0150reports.Report) {
	require.Equal(t, postID, migrated.PostId)
	require.Equal(t, original.Type, migrated.Type)
	require.Equal(t, original.Message, migrated.Message)
	require.Equal(t, original.User.String(), migrated.User)
}
