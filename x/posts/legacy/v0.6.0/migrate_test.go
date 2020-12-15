package v060_test

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

	v040GenState := v040.GenesisState{
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
		PostReactions: map[string][]v040.PostReaction{string(postID): {
			v040.PostReaction{
				Owner: postCreator,
				Value: ":fire:",
			},
			v040.PostReaction{
				Owner: postCreator,
				Value: ":my_house:",
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

	migrated := v060.Migrate(v040GenState)

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

func Test_GetConvertedPostReaction(t *testing.T) {
	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	registeredReactions := []v040.Reaction{
		{
			ShortCode: ":my_house:",
			Value:     "https://myHouse.jpeg",
			Subspace:  "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			Creator:   postCreator,
		},
	}

	tests := []struct {
		name                string
		postReaction        v040.PostReaction
		expMigratedReaction v060.PostReaction
		registeredReactions []v040.Reaction
	}{
		{
			name: "Migrate and emoji based reaction correctly",
			postReaction: v040.PostReaction{
				Owner: postCreator,
				Value: ":smile:",
			},
			expMigratedReaction: v060.PostReaction{
				Owner:     postCreator,
				Shortcode: ":smile:",
				Value:     "😄",
			},
			registeredReactions: nil,
		},
		{
			name: "Migrate and emoji based reaction correctly",
			postReaction: v040.PostReaction{
				Owner: postCreator,
				Value: ":my_house:",
			},
			expMigratedReaction: v060.PostReaction{
				Owner:     postCreator,
				Shortcode: ":my_house:",
				Value:     "https://myHouse.jpeg",
			},
			registeredReactions: registeredReactions,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := v060.GetConvertedPostReaction(test.postReaction, test.registeredReactions)
			require.Equal(t, test.expMigratedReaction, actual)
		})
	}

}

func TestMigrate050(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	content, err := ioutil.ReadFile("v050state.json")
	require.NoError(t, err)

	var v050state v040.GenesisState
	err = json.Unmarshal(content, &v050state)
	require.NoError(t, err)

	v060state := v060.Migrate(v050state)
	for _, reaction := range v060state.RegisteredReactions {
		// Make sure each reaction shortcode does not represent an emoji
		_, err := emoji.LookupEmojiByCode(reaction.ShortCode)
		require.Error(t, err)

		// Make sure no reaction value is an emoji
		_, err = emoji.LookupEmoji(reaction.Value)
		require.Error(t, err)
	}

	// Make sure the posts are all the same
	require.Equal(t, len(v060state.Posts), len(v050state.Posts))
	for index, post := range v060state.Posts {
		require.Equal(t, post, v050state.Posts[index])
	}

	// Make sure the reactions are all the same
	require.Equal(t, len(v060state.PostReactions), len(v050state.PostReactions))
	for index, postReactions := range v060state.PostReactions {
		require.Contains(t, postReactions, v050state.PostReactions[index])
	}

	// Make sure the poll answers are all the same
	require.Equal(t, len(v060state.UsersPollAnswers), len(v050state.UsersPollAnswers))
	for index, answer := range v060state.UsersPollAnswers {
		require.Equal(t, answer, v050state.UsersPollAnswers[index])
	}
}
