package v040_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	"github.com/stretchr/testify/require"
)

func TestMigratePostReactions(t *testing.T) {
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

	v030postReactions := map[string][]v030posts.Reaction{post.PostID.String(): {v030posts.Reaction{
		Owner: postCreator,
		Value: ":fire:",
	}}}

	expectedPostReactions := map[string][]v040posts.PostReaction{string(postID): {v040posts.PostReaction{
		Owner: postCreator,
		Value: ":fire:",
	}}}

	actualReactions, err := v040posts.MigratePostReactions(v030postReactions, []v030posts.Post{post})
	require.NoError(t, err)
	require.Equal(t, expectedPostReactions, actualReactions)
}

func TestGetReactionsToRegister(t *testing.T) {
	posts := []v040posts.Post{
		{PostID: "1", Subspace: "desmos"},
		{PostID: "2", Subspace: "mooncake"},
		{PostID: "3", Subspace: "random"},
		{PostID: "4", Subspace: "mooncake"},
	}

	user1, err := sdk.AccAddressFromBech32("cosmos1xvgje8eay7569pn6ps4rrqv0kn6zkkxj5t45fn")
	require.NoError(t, err)

	user2, err := sdk.AccAddressFromBech32("cosmos1gnp50apn896dstusk940rg5gllx2pnrp8jzfgc")
	require.NoError(t, err)

	reactions := map[string][]v040posts.PostReaction{
		"1": {
			v040posts.PostReaction{
				Value: ":+1:",
				Owner: user1,
			},
			v040posts.PostReaction{
				Value: ":thumbsup:",
				Owner: user2,
			},
		},
		"2": {},
		"3": {},
		"4": {
			v040posts.PostReaction{
				Value: ":heart:",
				Owner: user2,
			},
		},
	}

	reactionsToRegister, err := v040posts.GetReactionsToRegister(posts, reactions)
	require.NoError(t, err)

	expected := []v040posts.Reaction{
		{
			ShortCode: ":+1:",
			Value:     "👍",
			Subspace:  "desmos",
			Creator:   supply.NewModuleAddress("posts"),
		},
		{
			ShortCode: ":thumbsup:",
			Value:     "👍",
			Subspace:  "desmos",
			Creator:   supply.NewModuleAddress("posts"),
		},
		{
			ShortCode: ":heart:",
			Value:     "❤",
			Subspace:  "mooncake",
			Creator:   supply.NewModuleAddress("posts"),
		},
	}

	require.Len(t, reactionsToRegister, len(expected))
	for _, reaction := range reactionsToRegister {
		require.Contains(t, expected, reaction)
	}
}
