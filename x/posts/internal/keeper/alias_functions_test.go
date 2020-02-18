package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_IsPostDuplicate(t *testing.T) {
	firstUser, err := sdk.AccAddressFromBech32("cosmos10vu52vwmv5gn8k4xp9gcuz2an4my73lct9fnms")
	require.NoError(t, err)

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
			name: "Existing list returns true when it does (content)",
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
			name: "Existing list returns true when it does (ID)",
			posts: types.Posts{
				types.Post{
					PostID:         testPost.PostID,
					ParentID:       testPost.ParentID,
					Message:        testPost.Message + "other",
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
			name: "Existing list returns false when it does not",
			posts: types.Posts{
				types.Post{
					PostID:         testPost.PostID + 1,
					ParentID:       testPost.ParentID,
					Message:        testPost.Message + "other",
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
			expContains: false,
		},
		{
			name: "Same post but from different users give no problem",
			posts: types.Posts{
				types.Post{
					PostID:         testPost.PostID,
					ParentID:       testPost.ParentID,
					Message:        testPost.Message + "other",
					Created:        testPost.Created,
					LastEdited:     testPost.LastEdited,
					AllowsComments: testPost.AllowsComments,
					Subspace:       testPost.Subspace,
					OptionalData:   testPost.OptionalData,
					Creator:        firstUser,
					Medias:         testPost.Medias,
					PollData:       testPost.PollData,
				},
			},
			post: types.Post{
				PostID:         testPost.PostID + 1,
				ParentID:       testPost.ParentID,
				Message:        testPost.Message + "other",
				Created:        testPost.Created,
				LastEdited:     testPost.LastEdited,
				AllowsComments: testPost.AllowsComments,
				Subspace:       testPost.Subspace,
				OptionalData:   testPost.OptionalData,
				Creator:        secondUser,
				Medias:         testPost.Medias,
				PollData:       testPost.PollData,
			},
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

			post, contains := k.IsPostDuplicate(ctx, test.post)
			require.Equal(t, test.expContains, contains)

			if test.expContains {
				require.True(t, test.expPost.IsDuplicate(*post))
			} else {
				require.Nil(t, post)
			}
		})
	}
}
