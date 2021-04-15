package v0150_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v0130posts "github.com/desmos-labs/desmos/x/staging/posts/legacy/v0.13.0"
	v0150posts "github.com/desmos-labs/desmos/x/staging/posts/legacy/v0.15.0"
)

func TestMigrate(t *testing.T) {
	parentPostCreator, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	parentCreationTime := time.Now().UTC()
	postCreationTime := parentCreationTime.Add(time.Hour)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	v0130genesisState := v0130posts.GenesisState{
		Posts: []v0130posts.Post{
			{
				PostId:         "parent_id",
				ParentId:       "",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   []v0130posts.OptionalDataEntry{{Key: "optional", Value: "data"}},
				Created:        parentCreationTime,
				LastEdited:     time.Time{},
				Creator:        parentPostCreator,
				Attachments:    []v0130posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
			},
			{
				PostId:         "post_id_1",
				ParentId:       "parent_id",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   []v0130posts.OptionalDataEntry{{Key: "optional", Value: "data"}},
				Created:        postCreationTime,
				LastEdited:     time.Time{},
				Creator:        postCreator,
				Attachments:    []v0130posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
			},
		},
		UsersPollAnswers: map[string][]v0130posts.UserAnswer{
			"post_id_1": {
				v0130posts.UserAnswer{
					Answers: []uint64{1, 2},
					User:    postCreator,
				},
			},
		},
		PostReactions: map[string][]v0130posts.PostReaction{
			"post_id_1": {
				v0130posts.PostReaction{
					Owner:     postCreator,
					Shortcode: ":fire:",
					Value:     "🔥",
				},
				v0130posts.PostReaction{
					Owner:     postCreator,
					Shortcode: ":my_house:",
					Value:     "https://myHouse.png",
				},
			},
		},
		RegisteredReactions: []v0130posts.RegisteredReaction{
			{
				ShortCode: ":my_house:",
				Value:     "https://myHouse.png",
				Subspace:  subspace,
				Creator:   postCreator,
			},
		},
		Params: v0130posts.Params{
			MaxPostMessageLength:            sdk.NewInt(10),
			MaxOptionalDataFieldsNumber:     sdk.NewInt(10),
			MaxOptionalDataFieldValueLength: sdk.NewInt(10),
		},
	}

	expectedGenState := v0150posts.GenesisState{
		Posts: []v0150posts.Post{
			{
				PostId:         "parent_id",
				ParentId:       "",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []v0150posts.OptionalDataEntry{
					{
						Key:   "optional",
						Value: "data",
					},
				},
				Created:    parentCreationTime,
				LastEdited: time.Time{},
				Creator:    parentPostCreator.String(),
				Attachments: []v0150posts.Attachment{
					{
						URI:      "https://uri.com",
						MimeType: "text/plain",
						Tags:     []string{},
					},
				},
			},
			{
				PostId:         "post_id_1",
				ParentId:       "parent_id",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []v0150posts.OptionalDataEntry{
					{
						Key:   "optional",
						Value: "data",
					},
				},
				Created:    postCreationTime,
				LastEdited: time.Time{},
				Creator:    postCreator.String(),
				Attachments: []v0150posts.Attachment{
					{
						URI:      "https://uri.com",
						MimeType: "text/plain",
						Tags:     []string{},
					},
				},
			},
		},
		UsersPollAnswers: []v0150posts.UserAnswersEntry{
			{
				PostId: "post_id_1",
				UserAnswers: []v0150posts.UserAnswer{
					{
						User:    postCreator.String(),
						Answers: []string{"1", "2"},
					},
				},
			},
		},
		PostsReactions: []v0150posts.PostReactionsEntry{
			{
				PostId: "post_id_1",
				Reactions: []v0150posts.PostReaction{
					{
						ShortCode: ":fire:",
						Value:     "🔥",
						Owner:     postCreator.String(),
					},
					{
						ShortCode: ":my_house:",
						Value:     "https://myHouse.png",
						Owner:     postCreator.String(),
					},
				},
			},
		},
		RegisteredReactions: []v0150posts.RegisteredReaction{
			{
				ShortCode: ":my_house:",
				Value:     "https://myHouse.png",
				Subspace:  subspace,
				Creator:   postCreator.String(),
			},
		},
		Params: v0150posts.Params{
			MaxPostMessageLength:            sdk.NewInt(10),
			MaxOptionalDataFieldsNumber:     sdk.NewInt(10),
			MaxOptionalDataFieldValueLength: sdk.NewInt(10),
		},
	}

	migrated := v0150posts.Migrate(v0130genesisState)

	// Check for posts
	require.Len(t, expectedGenState.Posts, len(migrated.Posts))
	for index, post := range migrated.Posts {
		require.Equal(t, expectedGenState.Posts[index], post)
	}

	// Check for users poll answers
	require.Len(t, expectedGenState.UsersPollAnswers, len(migrated.UsersPollAnswers))
	for index, userAnswersEntry := range migrated.UsersPollAnswers {
		require.Equal(t, expectedGenState.UsersPollAnswers[index].PostId, userAnswersEntry.PostId)
		for idx, answers := range userAnswersEntry.UserAnswers {
			require.Equal(t, expectedGenState.UsersPollAnswers[index].UserAnswers[idx], answers)
		}
	}

	// Check for post reactions
	require.Len(t, expectedGenState.PostsReactions, len(migrated.PostsReactions))
	for index, postReactionEntry := range migrated.PostsReactions {
		require.Equal(t, expectedGenState.PostsReactions[index].PostId, postReactionEntry.PostId)
		for idx, postReaction := range postReactionEntry.Reactions {
			require.Equal(t, expectedGenState.PostsReactions[index].Reactions[idx], postReaction)
		}
	}

	// Check for registered reactions
	require.Len(t, expectedGenState.RegisteredReactions, len(migrated.RegisteredReactions))
	for index, regReaction := range migrated.RegisteredReactions {
		require.Equal(t, expectedGenState.RegisteredReactions[index], regReaction)
	}

	// Check for params
	require.Equal(t, migrated.Params, expectedGenState.Params)
}
