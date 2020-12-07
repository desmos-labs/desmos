package v0130_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v0100posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.10.0"
	v0120posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
	v0130 "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v0130posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
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

	parentID := v040posts.ComputeID(parentCreationTime, parentPostCreator, subspace)
	postID := v040posts.ComputeID(postCreationTime, postCreator, subspace)

	v0120GenState := v0120posts.GenesisState{
		Posts: []v0120posts.Post{
			{
				PostID:         parentID,
				ParentID:       "",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{"optional": "data"},
				Created:        parentCreationTime,
				LastEdited:     time.Time{},
				Creator:        parentPostCreator,
				Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
			{
				PostID:         postID,
				ParentID:       parentID,
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{"optional": "data"},
				Created:        postCreationTime,
				LastEdited:     time.Time{},
				Creator:        postCreator,
				Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain"}},
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
	}

	expected := v0130posts.GenesisState{
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
	}

	migrated := v0130posts.Migrate(v0120GenState)

	// Check for posts
	require.Len(t, migrated.Posts, len(expected.Posts))
	for index, post := range migrated.Posts {
		require.Equal(t, expected.Posts[index].PostID, post.PostID)
		require.Equal(t, expected.Posts[index].ParentID, post.ParentID)
		require.Equal(t, expected.Posts[index].OptionalData, post.OptionalData)
	}
}

func TestConvertPosts(t *testing.T) {
	parentPostCreator, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	parentCreationTime := time.Now().UTC()
	postCreationTime := parentCreationTime.Add(time.Hour)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	parentID := v040posts.ComputeID(parentCreationTime, parentPostCreator, subspace)
	postID := v040posts.ComputeID(postCreationTime, postCreator, subspace)

	var posts = []v0120posts.Post{
		{
			PostID:         parentID,
			ParentID:       "",
			Message:        "Message",
			AllowsComments: true,
			Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData:   map[string]string{"optional": "data"},
			Created:        parentCreationTime,
			LastEdited:     time.Time{},
			Creator:        parentPostCreator,
			Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain"}},
		},
		{
			PostID:         postID,
			ParentID:       parentID,
			Message:        "Message",
			AllowsComments: true,
			Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData:   map[string]string{"optional": "data"},
			Created:        postCreationTime,
			LastEdited:     time.Time{},
			Creator:        postCreator,
			Attachments:    []v0100posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain"}},
		},
	}

	var expectedPosts = []v0130posts.Post{
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
	}

	convertedPosts := v0130.ConvertPosts(posts)

	require.Equal(t, expectedPosts, convertedPosts)
	require.Equal(t, len(expectedPosts), len(convertedPosts))
}

func TestConvertOptionalData(t *testing.T) {
	oldOptionalData := v040posts.OptionalData{
		"optional": "data",
		"old":      "version",
		"another":  "data",
	}

	actualOptionalData := v0130.ConvertOptionalData(oldOptionalData)

	require.Equal(t, len(oldOptionalData), len(actualOptionalData))
	for key, value := range oldOptionalData {
		found := false
		for _, entry := range actualOptionalData {
			found = entry.Key == key && entry.Value == value
			if found {
				break
			}
		}
		require.True(t, found)
	}
}
