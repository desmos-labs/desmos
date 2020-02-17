package v030_test

import (
	"testing"

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
