package v080_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	"github.com/stretchr/testify/require"

	v040 "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060 "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	v080 "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
)

func TestMigrate(t *testing.T) {
	parentPostCreator, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	parentCreationTime := time.Now().UTC()
	postCreationTime := parentCreationTime.Add(time.Hour)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	parentID := v040.ComputeID(parentCreationTime, parentPostCreator, subspace)
	postID := v040.ComputeID(postCreationTime, postCreator, subspace)

	v060GenState := v060.GenesisState{
		Posts: []v040.Post{
			{
				PostID:         parentID,
				ParentID:       "",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Created:        parentCreationTime,
				LastEdited:     time.Time{},
				Creator:        parentPostCreator,
				Medias:         []v040.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
			{
				PostID:         postID,
				ParentID:       parentID,
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Created:        postCreationTime,
				LastEdited:     time.Time{},
				Creator:        postCreator,
				Medias:         []v040.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
		},
		UsersPollAnswers: map[string][]v040.UserAnswer{string(postID): {v040.UserAnswer{
			Answers: []v040.AnswerID{1, 2},
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
		RegisteredReactions: []v040.Reaction{
			{
				ShortCode: ":my_house:",
				Value:     "https://myHouse.png",
				Subspace:  subspace,
				Creator:   postCreator,
			},
		},
	}

	expected := v060.GenesisState{
		Posts: []v040.Post{
			{
				PostID:         parentID,
				ParentID:       "",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Created:        parentCreationTime,
				LastEdited:     time.Time{},
				Creator:        parentPostCreator,
				Medias:         []v040.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
			{
				PostID:         postID,
				ParentID:       parentID,
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{},
				Created:        postCreationTime,
				LastEdited:     time.Time{},
				Creator:        postCreator,
				Medias:         []v040.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
		},
		UsersPollAnswers: map[string][]v040.UserAnswer{string(postID): {v040.UserAnswer{
			Answers: []v040.AnswerID{1, 2},
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
		RegisteredReactions: []v040.Reaction{
			{
				ShortCode: ":my_house:",
				Value:     "https://myHouse.png",
				Subspace:  subspace,
				Creator:   postCreator,
			},
		},
	}

	migrated := v080.Migrate(v060GenState)

	// Check for posts
	require.Len(t, migrated.Posts, len(expected.Posts))
	for index, post := range migrated.Posts {
		require.Equal(t, expected.Posts[index].PostID, post.PostID)
		require.Equal(t, expected.Posts[index].ParentID, post.ParentID)
	}

	// Check for users poll answers
	require.Len(t, migrated.UsersPollAnswers, len(expected.UsersPollAnswers))
	for key := range expected.UsersPollAnswers {
		require.Equal(t, expected.UsersPollAnswers[key], migrated.UsersPollAnswers[key])
	}

	// Check for posts reactions
	require.Len(t, migrated.PostReactions, len(expected.PostReactions))
	for key := range expected.PostReactions {
		require.Equal(t, expected.PostReactions[key], migrated.PostReactions[key])
	}

	require.Len(t, migrated.RegisteredReactions, len(expected.RegisteredReactions))
	for index, reaction := range migrated.RegisteredReactions {
		require.Equal(t, expected.RegisteredReactions[index], reaction)
	}

}

func TestMigrate080(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	content, err := ioutil.ReadFile("v060state.json")
	require.NoError(t, err)

	var v060state v060.GenesisState
	err = json.Unmarshal(content, &v060state)
	require.NoError(t, err)

	v080state := v080.Migrate(v060state)
	for _, reaction := range v080state.RegisteredReactions {
		// Make sure each reaction shortcode does not represent an emoji
		_, err := emoji.LookupEmojiByCode(reaction.ShortCode)
		require.Error(t, err)

		// Make sure no reaction value is an emoji
		_, err = emoji.LookupEmoji(reaction.Value)
		require.Error(t, err)
	}

	// Make sure that all the posts are migrated
	require.Equal(t, len(v080state.Posts), len(v060state.Posts))

	// Make sure all the reactions are migrated
	require.Equal(t, len(v080state.PostReactions), len(v060state.PostReactions))

	// Make sure all the poll answers are migrated
	require.Equal(t, len(v080state.UsersPollAnswers), len(v060state.UsersPollAnswers))

	params := v080.Params{
		MaxPostMessageLength:            sdk.NewInt(500),
		MaxOptionalDataFieldsNumber:     sdk.NewInt(10),
		MaxOptionalDataFieldValueLength: sdk.NewInt(200),
	}

	// make sure params are properly set
	require.Equal(t, params, v080state.Params)
}

func TestRemoveInvalidEmojiRegisteredReactions(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("desmos1mmeu5t0j5284p7jkergq9hyejlhdwkzp25y84l")
	require.NoError(t, err)

	reactions := []v040.Reaction{
		{
			ShortCode: ":cool:",
			Value:     "https://test.com/example",
			Subspace:  "d4d5e4e8ac7fce379301602dc9c5614dd6fc49f042b1276db226e9de38776a5c",
			Creator:   user,
		},
		{
			ShortCode: ":new-shortcode:",
			Value:     "👍",
			Subspace:  "d4d5e4e8ac7fce379301602dc9c5614dd6fc49f042b1276db226e9de38776a5c",
			Creator:   user,
		},
		{
			ShortCode: ":new-shortcode:",
			Value:     "https://test.com/example",
			Subspace:  "d4d5e4e8ac7fce379301602dc9c5614dd6fc49f042b1276db226e9de38776a5c",
			Creator:   user,
		},
	}

	cleanedReactions := v080.RemoveInvalidEmojiRegisteredReactions(reactions)
	require.Len(t, cleanedReactions, 1)
	require.Equal(t, cleanedReactions[0], v040.Reaction{
		ShortCode: ":new-shortcode:",
		Value:     "https://test.com/example",
		Subspace:  "d4d5e4e8ac7fce379301602dc9c5614dd6fc49f042b1276db226e9de38776a5c",
		Creator:   user,
	})
}
