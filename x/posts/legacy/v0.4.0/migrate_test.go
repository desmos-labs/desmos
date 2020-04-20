package v040_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	"github.com/stretchr/testify/require"
)

func TestMigrate(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	owner2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	creationTime := time.Now().UTC()

	parentPost030 := v030posts.Post{
		PostID:         v030posts.PostID(1),
		ParentID:       v030posts.PostID(0),
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        creationTime,
		LastEdited:     time.Now().UTC().Add(time.Hour),
		Creator:        owner2,
		Medias:         v030posts.PostMedias{v030posts.PostMedia{URI: "https://uri.com", MimeType: "text/plain"}},
	}

	post030 := v030posts.Post{
		PostID:         v030posts.PostID(2),
		ParentID:       v030posts.PostID(1),
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        creationTime,
		LastEdited:     time.Now().UTC().Add(time.Hour),
		Creator:        owner,
		Medias:         v030posts.PostMedias{v030posts.PostMedia{URI: "https://uri.com", MimeType: "text/plain"}},
	}

	id := v040posts.ComputeID(post030.Created, post030.Creator, post030.Subspace)
	id2 := v040posts.ComputeID(parentPost030.Created, parentPost030.Creator, parentPost030.Subspace)

	parentPost040 := v040posts.Post{
		PostID:         id2,
		ParentID:       "",
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        time.Now().UTC(),
		LastEdited:     time.Now().UTC().Add(time.Hour),
		Creator:        owner2,
		Medias:         []v040posts.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
	}

	post040 := v040posts.Post{
		PostID:         id,
		ParentID:       id2,
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        time.Now().UTC(),
		LastEdited:     time.Now().UTC().Add(time.Hour),
		Creator:        owner,
		Medias:         []v040posts.PostMedia{{URI: "https://uri.com", MimeType: "text/plain"}},
	}

	postReaction030 := v030posts.Reaction{
		Owner: owner,
		Value: ":fire:",
	}

	postReaction040 := v040posts.PostReaction{
		Owner: owner,
		Value: ":fire:",
	}

	pollAnswer030 := v030posts.UserAnswer{
		Answers: []v030posts.AnswerID{1, 2},
		User:    owner,
	}

	pollAnswer040 := v040posts.UserAnswer{
		Answers: []v040posts.AnswerID{1, 2},
		User:    owner,
	}

	v030GenesisState := v030posts.GenesisState{
		Posts: []v030posts.Post{
			parentPost030,
			post030,
		},
		PollAnswers: map[string][]v030posts.UserAnswer{post030.PostID.String(): {pollAnswer030}},
		Reactions:   map[string][]v030posts.Reaction{post030.PostID.String(): {postReaction030}},
	}

	migrated := v040posts.Migrate(v030GenesisState)

	expected := v040posts.GenesisState{
		Posts: []v040posts.Post{
			parentPost040,
			post040,
		},
		UsersPollAnswers: map[string][]v040posts.UserAnswer{string(id): {pollAnswer040}},
		PostReactions:    map[string][]v040posts.PostReaction{string(id): {postReaction040}},
	}

	// Check for posts
	require.Len(t, migrated.Posts, len(expected.Posts))
	for index, post := range migrated.Posts {
		require.Equal(t, expected.Posts[index].PostID, post.PostID)
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

func TestRemoveDoublePosts(t *testing.T) {
	location, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	created := time.Date(2020, 1, 1, 15, 0, 0, 0, location)
	posts := []v030posts.Post{
		{PostID: 1, ParentID: 0, Message: "Post 1", Created: created},
		{PostID: 2, ParentID: 0, Message: "Post 1", Created: created},
		{PostID: 10, ParentID: 1, Message: "Post 3", Created: created.AddDate(0, 0, 1)},
		{PostID: 11, ParentID: 3, Message: "Post 4", Created: created.AddDate(0, 0, 2)},
		{PostID: 12, ParentID: 3, Message: "Post 4", Created: created.AddDate(0, 0, 2)},
		{PostID: 13, ParentID: 3, Message: "Post 4", Created: created.AddDate(0, 0, 2)},
	}
	reactions := map[string][]v030posts.Reaction{
		"2":  {{Value: "like"}},
		"11": {{Value: ":smile:"}, {Value: ":like:"}},
		"13": {{Value: ":heart:"}},
	}

	nResult, nReactions := v040posts.RemoveDoublePosts(posts, reactions)

	expPosts := []v030posts.Post{
		{PostID: 1, ParentID: 0, Message: "Post 1", Created: created},
		{PostID: 10, ParentID: 1, Message: "Post 3", Created: created.AddDate(0, 0, 1)},
		{PostID: 11, ParentID: 3, Message: "Post 4", Created: created.AddDate(0, 0, 2)},
	}
	require.Equal(t, expPosts, nResult)

	expReactions := map[string][]v030posts.Reaction{
		"1":  {{Value: "like"}},
		"11": {{Value: ":smile:"}, {Value: ":like:"}, {Value: ":heart:"}},
	}
	require.Equal(t, expReactions, nReactions)
}

func TestSquashPostIDs(t *testing.T) {
	posts := []v030posts.Post{
		{PostID: 15, ParentID: 1, Message: "Post 3"},
		{PostID: 1, ParentID: 0, Message: "Post 1"},
		{PostID: 30, ParentID: 12, Message: "Post 4"},
		{PostID: 12, ParentID: 0, Message: "Post 2"},
		{PostID: 50, ParentID: 15, Message: "Post 5"},
	}

	reactions := map[string][]v030posts.Reaction{
		"15": nil,
		"1":  {v030posts.Reaction{Value: "like"}},
		"12": {v030posts.Reaction{Value: "like"}, v030posts.Reaction{Value: "smile"}},
	}

	newPost, newReactions := v040posts.SquashPostIDs(posts, reactions)
	require.Len(t, newPost, len(posts))
	require.Len(t, newReactions, 2)

	expected := []v030posts.Post{
		{PostID: 1, ParentID: 0, Message: "Post 1"},
		{PostID: 2, ParentID: 0, Message: "Post 2"},
		{PostID: 3, ParentID: 1, Message: "Post 3"},
		{PostID: 4, ParentID: 2, Message: "Post 4"},
		{PostID: 5, ParentID: 3, Message: "Post 5"},
	}

	require.Equal(t, expected, newPost)
}
