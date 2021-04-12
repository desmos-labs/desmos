package v0130_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	v0120posts "github.com/desmos-labs/desmos/x/staging/posts/legacy/v0.12.0"
	v0130posts "github.com/desmos-labs/desmos/x/staging/posts/legacy/v0.13.0"
)

func TestMigrate(t *testing.T) {
	parentPostCreator, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	postCreator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	parentCreationTime := time.Now().UTC()
	postCreationTime := parentCreationTime.Add(time.Hour)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	parentID := v0120posts.ComputeID(parentCreationTime, parentPostCreator, subspace)
	postID := v0120posts.ComputeID(postCreationTime, postCreator, subspace)

	v0120GenState := v0120posts.GenesisState{
		Posts: []v0120posts.Post{
			{
				PostId:         parentID,
				ParentId:       "",
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{"optional": "data"},
				Created:        parentCreationTime,
				LastEdited:     time.Time{},
				Creator:        parentPostCreator,
				Attachments:    []v0120posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
			{
				PostId:         postID,
				ParentId:       parentID,
				Message:        "Message",
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   map[string]string{"optional": "data"},
				Created:        postCreationTime,
				LastEdited:     time.Time{},
				Creator:        postCreator,
				Attachments:    []v0120posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain"}},
			},
		},
		UsersPollAnswers: map[string][]v0120posts.UserAnswer{string(postID): {v0120posts.UserAnswer{
			Answers: []uint64{1, 2},
			User:    postCreator,
		}}},
		PostReactions: map[string][]v0120posts.PostReaction{string(postID): {
			v0120posts.PostReaction{
				Owner:     postCreator,
				Shortcode: ":fire:",
				Value:     "🔥",
			},
			v0120posts.PostReaction{
				Owner:     postCreator,
				Shortcode: ":my_house:",
				Value:     "https://myHouse.png",
			},
		}},
		RegisteredReactions: []v0120posts.RegisteredReaction{
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
				PostId:         parentID,
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
				PostId:         postID,
				ParentId:       parentID,
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
			postID: {
				v0130posts.UserAnswer{
					Answers: []uint64{1, 2},
					User:    postCreator,
				},
			},
		},
		PostReactions: map[string][]v0130posts.PostReaction{
			postID: {
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
	}

	migrated := v0130posts.Migrate(v0120GenState)

	// Check for posts
	require.Len(t, migrated.Posts, len(expected.Posts))
	for index, post := range migrated.Posts {
		require.Equal(t, expected.Posts[index].PostId, post.PostId)
		require.Equal(t, expected.Posts[index].ParentId, post.ParentId)
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

	parentID := v0120posts.ComputeID(parentCreationTime, parentPostCreator, subspace)
	postID := v0120posts.ComputeID(postCreationTime, postCreator, subspace)

	var posts = []v0120posts.Post{
		{
			PostId:         parentID,
			ParentId:       "",
			Message:        "Message",
			AllowsComments: true,
			Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData:   map[string]string{"optional": "data"},
			Created:        parentCreationTime,
			LastEdited:     time.Time{},
			Creator:        parentPostCreator,
			Attachments:    []v0120posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain"}},
		},
		{
			PostId:         postID,
			ParentId:       parentID,
			Message:        "Message",
			AllowsComments: true,
			Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData:   map[string]string{"optional": "data"},
			Created:        postCreationTime,
			LastEdited:     time.Time{},
			Creator:        postCreator,
			Attachments:    []v0120posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain"}},
		},
	}

	var expectedPosts = []v0130posts.Post{
		{
			PostId:         parentID,
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
			PostId:         postID,
			ParentId:       parentID,
			Message:        "Message",
			AllowsComments: true,
			Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			OptionalData:   []v0130posts.OptionalDataEntry{{Key: "optional", Value: "data"}},
			Created:        postCreationTime,
			LastEdited:     time.Time{},
			Creator:        postCreator,
			Attachments:    []v0130posts.Attachment{{URI: "https://uri.com", MimeType: "text/plain", Tags: nil}},
		},
	}

	convertedPosts := v0130posts.ConvertPosts(posts)

	require.Equal(t, expectedPosts, convertedPosts)
	require.Equal(t, len(expectedPosts), len(convertedPosts))
}

func TestConvertOptionalData(t *testing.T) {
	oldOptionalData := v0120posts.OptionalData{
		"optional": "data",
		"old":      "version",
		"another":  "data",
	}

	actualOptionalData := v0130posts.ConvertOptionalData(oldOptionalData)

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
