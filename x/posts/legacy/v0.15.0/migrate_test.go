package v0150_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v0100posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.10.0"
	v0130 "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v0150posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.15.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060 "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	v080posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
	"github.com/stretchr/testify/require"
)

func TestMigrate(t *testing.T) {
	parentPostCreator, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	parentCreationTime := time.Now().UTC()
	postCreationTime := parentCreationTime.Add(time.Hour)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	parentID := v040posts.ComputeID(parentCreationTime, parentPostCreator, subspace)
	postID := v040posts.ComputeID(postCreationTime, postCreator, subspace)

	v0130genesisState := v0130posts.GenesisState{
		Posts: []v0130posts.Post{
			{
				PostID:         parentID,
				ParentID:       "",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   []v0130.OptionalDataEntry{{Key: "optional", Value: "data"}},
				Created:        parentCreationTime,
				LastEdited:     time.Time{},
				Creator:        parentPostCreator,
				Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
			},
			{
				PostID:         postID,
				ParentID:       parentID,
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   []v0130.OptionalDataEntry{{Key: "optional", Value: "data"}},
				Created:        postCreationTime,
				LastEdited:     time.Time{},
				Creator:        postCreator,
				Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
			},
		},
		UsersPollAnswers: map[string][]v040posts.UserAnswer{string(postID): {v040posts.UserAnswer{
			Answers: []v040posts.AnswerID{1, 2},
			User:    postCreator,
		}}},
		PostReactions: map[string][]v060.PostReaction{string(postID): {
			v060.PostReaction{
				Owner:     postCreator,
				Shortcode: ":fire:",
				Value:     "🔥",
			},
			v060.PostReaction{
				Owner:     postCreator,
				Shortcode: ":my_house:",
				Value:     "https://myHouse.png",
			},
		}},
		RegisteredReactions: []v040posts.Reaction{
			{
				ShortCode: ":my_house:",
				Value:     "https://myHouse.png",
				Subspace:  subspace,
				Creator:   postCreator,
			},
		},
		Params: v080posts.Params{
			MaxPostMessageLength:            sdk.NewInt(10),
			MaxOptionalDataFieldsNumber:     sdk.NewInt(10),
			MaxOptionalDataFieldValueLength: sdk.NewInt(10),
		},
	}

	expectedGenState := v0150posts.GenesisState{
		Posts: []v0150posts.Post{
			{
				PostID:         string(parentID),
				ParentID:       "",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   []v0130.OptionalDataEntry{{Key: "optional", Value: "data"}},
				Created:        parentCreationTime,
				LastEdited:     time.Time{},
				Creator:        parentPostCreator.String(),
				Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
			},
			{
				PostID:         string(postID),
				ParentID:       string(parentID),
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   []v0130.OptionalDataEntry{{Key: "optional", Value: "data"}},
				Created:        postCreationTime,
				LastEdited:     time.Time{},
				Creator:        postCreator.String(),
				Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
			},
		},
		UsersPollAnswers: []v0150posts.UserAnswersEntry{
			{
				PostId: string(postID),
				UserAnswers: []v0150posts.UserAnswer{
					{
						User:    postCreator.String(),
						Answers: []string{"1", "2"},
					},
				},
			},
		},
		PostReactions: []v0150posts.PostReactionsEntry{
			{
				PostId: string(postID),
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
		Params: v080posts.Params{
			MaxPostMessageLength:            sdk.NewInt(10),
			MaxOptionalDataFieldsNumber:     sdk.NewInt(10),
			MaxOptionalDataFieldValueLength: sdk.NewInt(10),
		},
	}

	migrated := v0150posts.Migrate(v0130genesisState)

	// Check for posts
	require.Len(t, expectedGenState.Posts, len(migrated.Posts))
	for index, post := range migrated.Posts {
		require.Equal(t, expectedGenState.Posts[index].PostID, post.PostID)
		require.Equal(t, expectedGenState.Posts[index].ParentID, post.ParentID)
		require.Equal(t, expectedGenState.Posts[index].OptionalData, post.OptionalData)
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
	require.Len(t, expectedGenState.PostReactions, len(migrated.PostReactions))
	for index, postReactionEntry := range migrated.PostReactions {
		require.Equal(t, expectedGenState.PostReactions[index].PostId, postReactionEntry.PostId)
		for idx, postReaction := range postReactionEntry.Reactions {
			require.Equal(t, t, expectedGenState.PostReactions[index].Reactions[idx], postReaction)
		}
	}

	// Check for registered reactions
	require.Len(t, expectedGenState.RegisteredReactions, len(migrated.RegisteredReactions))
	for index, regReaction := range migrated.RegisteredReactions {
		require.Equal(t, expectedGenState.RegisteredReactions[index].Creator, regReaction.Creator)
		require.Equal(t, expectedGenState.RegisteredReactions[index].ShortCode, regReaction.ShortCode)
		require.Equal(t, expectedGenState.RegisteredReactions[index].Subspace, regReaction.Subspace)
		require.Equal(t, expectedGenState.RegisteredReactions[index].Value, regReaction.Value)
	}

	// Check for params
	require.Equal(t, migrated.Params, expectedGenState.Params)
}
