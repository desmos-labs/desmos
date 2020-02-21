package v030_test

import (
	"testing"
	"time"

	v020posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.2.0"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
	"github.com/stretchr/testify/require"
)

func TestMigrate(t *testing.T) {
	v020GenesisState := v020posts.GenesisState{
		Posts: []v020posts.Post{
			{Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"},
			{Subspace: "106a5842fc5fce6f663176285ed1516dbb1e3d15c05abab12fdca46d60b539b7"},
		},
		Reactions: map[string][]v020posts.Reaction{},
	}

	migrated := v030posts.Migrate(v020GenesisState)
	expected := v030posts.GenesisState{
		Posts: []v030posts.Post{
			{Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"},
			{Subspace: "106a5842fc5fce6f663176285ed1516dbb1e3d15c05abab12fdca46d60b539b7"},
		},
		Reactions: nil,
	}

	// Check for posts
	require.Len(t, migrated.Posts, len(expected.Posts))
	for index, post := range migrated.Posts {
		require.Equal(t, expected.Posts[index].Subspace, post.Subspace)
	}
}

func TestRemoveDoublePosts(t *testing.T) {
	location, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	created := time.Date(2020, 1, 1, 15, 0, 0, 0, location)
	posts := []v020posts.Post{
		{PostID: 1, ParentID: 0, Message: "Post 1", Created: created},
		{PostID: 2, ParentID: 0, Message: "Post 1", Created: created},
		{PostID: 10, ParentID: 1, Message: "Post 3", Created: created.AddDate(0, 0, 1)},
		{PostID: 11, ParentID: 3, Message: "Post 4", Created: created.AddDate(0, 0, 2)},
		{PostID: 12, ParentID: 3, Message: "Post 4", Created: created.AddDate(0, 0, 2)},
		{PostID: 13, ParentID: 3, Message: "Post 4", Created: created.AddDate(0, 0, 2)},
	}
	reactions := map[string][]v020posts.Reaction{
		"2":  {{Value: "like"}},
		"11": {{Value: ":smile:"}, {Value: ":like:"}},
		"13": {{Value: ":heart:"}},
	}

	nResult, nReactions := v030posts.RemoveDoublePosts(posts, reactions)

	expPosts := []v020posts.Post{
		{PostID: 1, ParentID: 0, Message: "Post 1", Created: created},
		{PostID: 10, ParentID: 1, Message: "Post 3", Created: created.AddDate(0, 0, 1)},
		{PostID: 11, ParentID: 3, Message: "Post 4", Created: created.AddDate(0, 0, 2)},
	}
	require.Equal(t, expPosts, nResult)

	expReactions := map[string][]v020posts.Reaction{
		"1":  {{Value: "like"}},
		"11": {{Value: ":smile:"}, {Value: ":like:"}, {Value: ":heart:"}},
	}
	require.Equal(t, expReactions, nReactions)
}

func TestSquashPostIDs(t *testing.T) {
	posts := []v020posts.Post{
		{PostID: 15, ParentID: 1, Message: "Post 3"},
		{PostID: 1, ParentID: 0, Message: "Post 1"},
		{PostID: 30, ParentID: 12, Message: "Post 4"},
		{PostID: 12, ParentID: 0, Message: "Post 2"},
		{PostID: 50, ParentID: 15, Message: "Post 5"},
	}

	reactions := map[string][]v020posts.Reaction{
		"15": nil,
		"1":  {v020posts.Reaction{Value: "like"}},
		"12": {v020posts.Reaction{Value: "like"}, v020posts.Reaction{Value: "smile"}},
	}

	newPost, newReactions := v030posts.SquashPostIDs(posts, reactions)
	require.Len(t, newPost, len(posts))
	require.Len(t, newReactions, 2)

	expected := []v020posts.Post{
		{PostID: 1, ParentID: 0, Message: "Post 1"},
		{PostID: 2, ParentID: 0, Message: "Post 2"},
		{PostID: 3, ParentID: 1, Message: "Post 3"},
		{PostID: 4, ParentID: 2, Message: "Post 4"},
		{PostID: 5, ParentID: 3, Message: "Post 5"},
	}

	require.Equal(t, expected, newPost)
}
