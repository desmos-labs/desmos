package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_IsPostConflicting(t *testing.T) {
	secondUser, err := sdk.AccAddressFromBech32("cosmos18438yx7re4hrdxxc64rcfdthl5qkfdejq24lta")
	require.NoError(t, err)

	tests := []struct {
		name        string
		posts       types.Posts
		post        types.Post
		expContains bool
		expPost     types.Post
	}{
		{
			name:        "Empty list returns false",
			posts:       types.Posts{},
			post:        testPost,
			expContains: false,
		},
		{
			name: "Same exact post returns true",
			posts: types.Posts{
				types.Post{
					PostID:         testPost.PostID + 1,
					ParentID:       testPost.ParentID,
					Message:        testPost.Message,
					Created:        testPost.Created,
					LastEdited:     testPost.LastEdited,
					AllowsComments: testPost.AllowsComments,
					Subspace:       testPost.Subspace,
					OptionalData:   testPost.OptionalData,
					Creator:        testPost.Creator,
					Medias:         testPost.Medias,
					PollData:       testPost.PollData,
				},
			},
			post:        testPost,
			expContains: true,
			expPost:     testPost,
		},
		{
			name: "Post with different creation date returns false",
			posts: types.Posts{
				types.Post{
					PostID:         testPost.PostID + 1,
					ParentID:       testPost.ParentID,
					Message:        testPost.Message,
					Created:        testPost.Created.AddDate(0, 0, 1),
					LastEdited:     testPost.LastEdited,
					AllowsComments: testPost.AllowsComments,
					Subspace:       testPost.Subspace,
					OptionalData:   testPost.OptionalData,
					Creator:        testPost.Creator,
					Medias:         testPost.Medias,
					PollData:       testPost.PollData,
				},
			},
			post:        testPost,
			expContains: false,
			expPost:     testPost,
		},
		{
			name: "Post with different subspace returns false",
			posts: types.Posts{
				types.Post{
					PostID:         testPost.PostID + 1,
					ParentID:       testPost.ParentID,
					Message:        testPost.Message,
					Created:        testPost.Created,
					LastEdited:     testPost.LastEdited,
					AllowsComments: testPost.AllowsComments,
					Subspace:       testPost.Subspace + "other",
					OptionalData:   testPost.OptionalData,
					Creator:        testPost.Creator,
					Medias:         testPost.Medias,
					PollData:       testPost.PollData,
				},
			},
			post:        testPost,
			expContains: false,
			expPost:     testPost,
		},
		{
			name: "Post with different creator returns false",
			posts: types.Posts{
				types.Post{
					PostID:         testPost.PostID + 1,
					ParentID:       testPost.ParentID,
					Message:        testPost.Message,
					Created:        testPost.Created,
					LastEdited:     testPost.LastEdited,
					AllowsComments: testPost.AllowsComments,
					Subspace:       testPost.Subspace,
					OptionalData:   testPost.OptionalData,
					Creator:        secondUser,
					Medias:         testPost.Medias,
					PollData:       testPost.PollData,
				},
			},
			post:        testPost,
			expContains: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			for _, post := range test.posts {
				k.SavePost(ctx, post)
			}

			post, contains := k.IsPostConflicting(ctx, test.post)
			require.Equal(t, test.expContains, contains)

			if test.expContains {
				require.True(t, test.expPost.IsConflictingWith(*post))
			} else {
				require.Nil(t, post)
			}
		})
	}
}
