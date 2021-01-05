package v0100_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v0100posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.10.0"
	v040 "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060 "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	v080posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
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

	v080GenState := v080posts.GenesisState{
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

	expected := v0100posts.GenesisState{
		Posts: []v0100posts.Post{
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
				Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
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
				Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
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

	migrated := v0100posts.Migrate(v080GenState)

	// Check for posts
	require.Len(t, migrated.Posts, len(expected.Posts))
	for index, post := range migrated.Posts {
		require.Equal(t, expected.Posts[index].PostID, post.PostID)
		require.Equal(t, expected.Posts[index].ParentID, post.ParentID)

		// Check for attachments
		require.Equal(t, expected.Posts[index].Attachments, post.Attachments)
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

func TestMigrate0100(t *testing.T) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("desmos", "desmos"+sdk.PrefixPublic)
	config.Seal()

	content, err := ioutil.ReadFile("v080state.json")
	require.NoError(t, err)

	var v080state v080posts.GenesisState
	err = json.Unmarshal(content, &v080state)
	require.NoError(t, err)

	v0100state := v0100posts.Migrate(v080state)

	// Make sure that all the posts are migrated
	require.Equal(t, len(v0100state.Posts), len(v080state.Posts))

	// Make sure that all the posts' medias are migrated correctly
	for index, post := range v080state.Posts {
		require.Equal(t, len(post.Medias), len(v0100state.Posts[index].Attachments))
	}

	// Make sure all the reactions are migrated
	require.Equal(t, len(v0100state.PostReactions), len(v080state.PostReactions))

	// Make sure all the poll answers are migrated
	require.Equal(t, len(v0100state.UsersPollAnswers), len(v080state.UsersPollAnswers))

	// make sure params are properly set
	require.Equal(t, v080state.Params, v0100state.Params)
}

func TestConvertMediasToAttachments(t *testing.T) {
	postMedias := []v040.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}}
	attachments := []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}}

	actual := v0100posts.ConvertMediasToAttachments(postMedias)

	require.Equal(t, attachments, actual)
}

func TestConvertPosts(t *testing.T) {
	parentPostCreator, err := sdk.AccAddressFromBech32("desmos1mmeu5t0j5284p7jkergq9hyejlhdwkzp25y84l")
	require.NoError(t, err)

	postCreator, err := sdk.AccAddressFromBech32("desmos1mmeu5t0j5284p7jkergq9hyejlhdwkzp25y84l")
	require.NoError(t, err)

	parentCreationTime := time.Now().UTC()
	postCreationTime := parentCreationTime.Add(time.Hour)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	parentID := v040.ComputeID(parentCreationTime, parentPostCreator, subspace)
	postID := v040.ComputeID(postCreationTime, postCreator, subspace)

	posts := []v040.Post{
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
	}

	expectedPosts := []v0100posts.Post{
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
			Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
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
			Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
		},
	}

	actualPosts := v0100posts.ConvertPosts(posts)

	require.Equal(t, expectedPosts, actualPosts)
}
