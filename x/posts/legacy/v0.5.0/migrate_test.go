package v050_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	emoji "github.com/desmos-labs/Go-Emoji-Utils"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v050posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.5.0"
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

	v040GenState := v040posts.GenesisState{
		Posts: []v040posts.Post{
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
				Medias:         []v040posts.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
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
				Medias:         []v040posts.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
		},
		UsersPollAnswers: map[string][]v040posts.UserAnswer{string(postID): {v040posts.UserAnswer{
			Answers: []v040posts.AnswerID{1, 2},
			User:    postCreator,
		}}},
		PostReactions: map[string][]v040posts.PostReaction{string(postID): {
			v040posts.PostReaction{
				Owner: postCreator,
				Value: ":fire:",
			},
		}},
		RegisteredReactions: []v040posts.Reaction{
			{
				ShortCode: ":fire:",
				Value:     "🔥",
				Subspace:  subspace,
				Creator:   postCreator,
			},
			{
				ShortCode: ":my_house:",
				Value:     "https://myHouse.png",
				Subspace:  subspace,
				Creator:   postCreator,
			},
		},
	}

	expected := v040posts.GenesisState{
		Posts: []v040posts.Post{
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
				Medias:         []v040posts.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
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
				Medias:         []v040posts.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
		},
		UsersPollAnswers: map[string][]v040posts.UserAnswer{string(postID): {v040posts.UserAnswer{
			Answers: []v040posts.AnswerID{1, 2},
			User:    postCreator,
		}}},
		PostReactions: map[string][]v040posts.PostReaction{string(postID): {
			v040posts.PostReaction{
				Owner: postCreator,
				Value: ":fire:",
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
	}

	migrated := v050posts.Migrate(v040GenState)

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

func TestGetReactionsToRegister(t *testing.T) {
	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	registeredReactions := []v040posts.Reaction{
		{
			ShortCode: ":fire:",
			Value:     "🔥",
			Subspace:  subspace,
			Creator:   postCreator,
		},
		{
			ShortCode: ":my_house:",
			Value:     "https://myHouse.png",
			Subspace:  subspace,
			Creator:   postCreator,
		},
	}

	expected := []v040posts.Reaction{
		{
			ShortCode: ":my_house:",
			Value:     "https://myHouse.png",
			Subspace:  subspace,
			Creator:   postCreator,
		},
	}

	actual := v050posts.GetReactionsToRegister(registeredReactions)

	require.Equal(t, expected, actual)
}

func TestMigrate040(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	content, err := ioutil.ReadFile("v040state.json")
	require.NoError(t, err)

	var v040state v040posts.GenesisState
	err = json.Unmarshal(content, &v040state)
	require.NoError(t, err)

	v050state := v050posts.Migrate(v040state)
	for _, reaction := range v050state.RegisteredReactions {
		// Make sure each reaction shrotcode does not represent an emoji
		_, err := emoji.LookupEmojiByCode(reaction.ShortCode)
		require.Error(t, err)

		// Make sure no reaction value is an emoji
		_, err = emoji.LookupEmoji(reaction.Value)
		require.Error(t, err)
	}

	// Make sure the posts are all the same
	require.Equal(t, len(v050state.Posts), len(v040state.Posts))
	for index, post := range v050state.Posts {
		require.Equal(t, post, v040state.Posts[index])
	}

	// Make sure the reactions are all the same
	require.Equal(t, len(v050state.PostReactions), len(v040state.PostReactions))
	for index, postReaction := range v050state.PostReactions {
		require.Equal(t, postReaction, v040state.PostReactions[index])
	}

	// Make sure the poll answers are all the same
	require.Equal(t, len(v050state.UsersPollAnswers), len(v040state.UsersPollAnswers))
	for index, answer := range v050state.UsersPollAnswers {
		require.Equal(t, answer, v040state.UsersPollAnswers[index])
	}
}
