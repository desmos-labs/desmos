package v040_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	"github.com/stretchr/testify/require"
)

func TestMigrate(t *testing.T) {
	parentPostCreator, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	parentCreationTime := time.Now().UTC()
	postCreationTime := parentCreationTime.Add(time.Hour)

	parentPost := v030posts.Post{
		PostID:         v030posts.PostID(1),
		ParentID:       v030posts.PostID(0),
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        parentCreationTime,
		LastEdited:     time.Time{},
		Creator:        parentPostCreator,
		Medias:         v030posts.PostMedias{v030posts.PostMedia{URI: "https://uri.com", MimeType: "text/plain"}},
	}

	post := v030posts.Post{
		PostID:         v030posts.PostID(2),
		ParentID:       v030posts.PostID(1),
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        postCreationTime,
		LastEdited:     time.Time{},
		Creator:        postCreator,
		Medias:         v030posts.PostMedias{v030posts.PostMedia{URI: "https://uri.com", MimeType: "text/plain"}},
	}

	v030GenesisState := v030posts.GenesisState{
		Posts: []v030posts.Post{
			parentPost,
			post,
		},
		PollAnswers: map[string][]v030posts.UserAnswer{post.PostID.String(): {v030posts.UserAnswer{
			Answers: []v030posts.AnswerID{1, 2},
			User:    postCreator,
		}}},
		Reactions: map[string][]v030posts.Reaction{post.PostID.String(): {v030posts.Reaction{
			Owner: postCreator,
			Value: ":fire:",
		}}},
	}

	parentID := v040posts.ComputeID(parentPost.Created, parentPost.Creator, parentPost.Subspace)
	postID := v040posts.ComputeID(post.Created, post.Creator, post.Subspace)

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
		PostReactions: map[string][]v040posts.PostReaction{string(postID): {v040posts.PostReaction{
			Owner: postCreator,
			Value: ":fire:",
		}}},
	}

	// Migrate
	migrated := v040posts.Migrate(v030GenesisState)

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

}

func TestConvertID(t *testing.T) {
	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)
	creationDate := time.Date(2020, 4, 20, 0, 0, 0, 0, timeZone)

	post := v030posts.Post{
		PostID:         v030posts.PostID(2),
		ParentID:       v030posts.PostID(1),
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        creationDate,
		LastEdited:     time.Time{},
		Creator:        postCreator,
		Medias:         v030posts.PostMedias{v030posts.PostMedia{URI: "https://uri.com", MimeType: "text/plain"}},
	}

	actualID, err := v040posts.ConvertID("2", []v030posts.Post{post})
	require.NoError(t, err)
	require.Equal(t, v040posts.PostID("58c2b3f9cd00d04369595f51270acb5f382192293b540d8b2a7e1c917a4f8ad7"), actualID)
}

func TestMigrateUsersAnswers(t *testing.T) {
	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	parentCreationTime := time.Now().UTC()
	postCreationTime := parentCreationTime.Add(time.Hour)

	post := v030posts.Post{
		PostID:         v030posts.PostID(2),
		ParentID:       v030posts.PostID(1),
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        postCreationTime,
		LastEdited:     time.Time{},
		Creator:        postCreator,
		Medias:         v030posts.PostMedias{v030posts.PostMedia{URI: "https://uri.com", MimeType: "text/plain"}},
	}

	postID := v040posts.ComputeID(post.Created, post.Creator, post.Subspace)

	v030UserAnswers := map[string][]v030posts.UserAnswer{post.PostID.String(): {v030posts.UserAnswer{
		Answers: []v030posts.AnswerID{1, 2},
		User:    postCreator,
	}}}

	expectedAnswers := map[string][]v040posts.UserAnswer{string(postID): {v040posts.UserAnswer{
		Answers: []v040posts.AnswerID{1, 2},
		User:    postCreator,
	}}}

	actualAnswers, err := v040posts.MigrateUsersAnswers(v030UserAnswers, []v030posts.Post{post})
	require.NoError(t, err)
	require.Equal(t, expectedAnswers, actualAnswers)
}

func TestMigrate030(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	content, err := ioutil.ReadFile("v030state.json")
	require.NoError(t, err)

	var v030state v030posts.GenesisState
	err = json.Unmarshal(content, &v030state)
	require.NoError(t, err)

	posts := v040posts.MigratePosts(v030state.Posts)

	_, err = v040posts.MigrateUsersAnswers(v030state.PollAnswers, v030state.Posts)
	require.NoError(t, err)

	reactions, err := v040posts.MigratePostReactions(v030state.Reactions, v030state.Posts)
	require.NoError(t, err)

	_, err = v040posts.GetReactionsToRegister(posts, reactions)
	require.NoError(t, err)

	v040posts.Migrate(v030state)
}
