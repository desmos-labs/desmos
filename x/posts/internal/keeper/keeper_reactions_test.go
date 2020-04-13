package keeper_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -------------
// --- PostReactions
// -------------

func TestKeeper_SaveReaction(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name               string
		storedReaction     types.PostReactions
		postID             types.PostID
		reaction           types.PostReaction
		storedPost         types.Post
		registeredReaction types.Reaction
		error              error
		expectedStored     types.PostReactions
	}{
		{
			name:           "PostReaction from same user already present returns expError",
			storedReaction: types.PostReactions{types.NewPostReaction(":like:", liker)},
			postID:         types.PostID(id),
			reaction:       types.NewPostReaction(":like:", liker),
			storedPost: types.NewPost(
				id,
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			registeredReaction: types.NewReaction(liker, ":like:", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error:          fmt.Errorf("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 has already reacted with :like: to the post with id 10"),
			expectedStored: types.PostReactions{types.NewPostReaction(":like:", liker)},
		},
		{
			name:           "PostReaction is not a registered reaction and returns error",
			storedReaction: types.PostReactions{},
			postID:         types.PostID(id),
			reaction:       types.NewPostReaction(":like:", liker),
			storedPost: types.NewPost(
				id,
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			registeredReaction: types.NewReaction(liker, ":smile:", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: fmt.Errorf("reaction with short code :like: isn't registered yet and can't be used to react to the post with ID 10 and sub 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e, please register it before use"),
		},
		{
			name:           "First liker is stored properly",
			storedReaction: types.PostReactions{},
			postID:         types.PostID(id),
			reaction:       types.NewPostReaction(":like:", liker),
			storedPost: types.NewPost(
				id,
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			registeredReaction: types.NewReaction(liker, ":like:", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error:          nil,
			expectedStored: types.PostReactions{types.NewPostReaction(":like:", liker)},
		},
		{
			name:           "Second liker is stored properly",
			storedReaction: types.PostReactions{types.NewPostReaction(":like:", liker)},
			postID:         types.PostID(id),
			reaction:       types.NewPostReaction(":like:", otherLiker),
			storedPost: types.NewPost(
				id,
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Created,
				testPost.Creator,
			),
			registeredReaction: types.NewReaction(liker, ":like:", "https://smile.jpg",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			error: nil,
			expectedStored: types.PostReactions{
				types.NewPostReaction(":like:", liker),
				types.NewPostReaction(":like:", otherLiker),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if len(test.storedReaction) != 0 {
				store.Set(types.PostReactionsStoreKey(test.postID), k.Cdc.MustMarshalBinaryBare(&test.storedReaction))
			}

			k.SavePost(ctx, test.storedPost)
			k.RegisterReaction(ctx, test.registeredReaction)

			err := k.SavePostReaction(ctx, test.postID, test.reaction)
			require.Equal(t, test.error, err)

			var stored types.PostReactions
			k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(test.postID)), &stored)
			require.Equal(t, test.expectedStored, stored)
		})
	}
}

func TestKeeper_RemoveReaction(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name           string
		storedLikes    types.PostReactions
		postID         types.PostID
		liker          sdk.AccAddress
		value          string
		error          error
		expectedStored types.PostReactions
	}{
		{
			name:           "PostReaction from same liker is removed properly",
			storedLikes:    types.PostReactions{types.NewPostReaction(":like:", liker)},
			postID:         types.PostID(id),
			liker:          liker,
			value:          ":like:",
			error:          nil,
			expectedStored: types.PostReactions{},
		},
		{
			name:           "Non existing reaction returns error - Creator",
			storedLikes:    types.PostReactions{},
			postID:         types.PostID(id),
			liker:          liker,
			value:          ":like:",
			error:          fmt.Errorf("cannot remove the reaction with value :like: from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.PostReactions{},
		},
		{
			name:           "Non existing reaction returns error - Value",
			storedLikes:    types.PostReactions{types.NewPostReaction(":like:", liker)},
			postID:         types.PostID(id),
			liker:          liker,
			value:          ":smile:",
			error:          fmt.Errorf("cannot remove the reaction with value :smile: from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.PostReactions{types.NewPostReaction(":like:", liker)},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if len(test.storedLikes) != 0 {
				store.Set(types.PostReactionsStoreKey(test.postID), k.Cdc.MustMarshalBinaryBare(&test.storedLikes))
			}

			err := k.RemovePostReaction(ctx, test.postID, test.liker, test.value)
			require.Equal(t, test.error, err)

			var stored types.PostReactions
			k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(test.postID)), &stored)

			require.Len(t, stored, len(test.expectedStored))
			for index, like := range test.expectedStored {
				require.Equal(t, like, stored[index])
			}
		})
	}
}

func TestKeeper_GetPostReactions(t *testing.T) {
	id := []byte("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name               string
		likes              types.PostReactions
		postID             types.PostID
		storedPost         types.Post
		registeredReaction types.Reaction
	}{
		{
			name:   "Empty list are returned properly",
			likes:  types.PostReactions{},
			postID: types.PostID(id),
		},
		{
			name: "Valid list of likes is returned properly",
			likes: types.PostReactions{
				types.NewPostReaction(":smile:", otherLiker),
				types.NewPostReaction(":smile:", liker),
			},
			postID:             types.PostID(id),
			storedPost:         testPost,
			registeredReaction: testRegisteredReaction,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, l := range test.likes {
				k.SavePost(ctx, test.storedPost)
				k.RegisterReaction(ctx, test.registeredReaction)
				err := k.SavePostReaction(ctx, test.postID, l)
				require.NoError(t, err)
			}

			stored := k.GetPostReactions(ctx, test.postID)

			require.Len(t, stored, len(test.likes))
			for _, l := range test.likes {
				require.Contains(t, stored, l)
			}
		})
	}
}

func TestKeeper_GetReactions(t *testing.T) {
	id := "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"
	id2 := "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"
	liker1, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	liker2, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name  string
		likes map[string]types.PostReactions
	}{
		{
			name:  "Empty likes data are returned correctly",
			likes: map[string]types.PostReactions{},
		},
		{
			name: "Non empty likes data are returned correcly",
			likes: map[string]types.PostReactions{
				id: {
					types.NewPostReaction("reaction", liker1),
					types.NewPostReaction("reaction", liker2),
				},
				id2: {
					types.NewPostReaction("reaction", liker1),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			for postID, likes := range test.likes {
				store.Set(types.PostReactionsStoreKey([]byte(postID)), k.Cdc.MustMarshalBinaryBare(likes))
			}

			likesData := k.GetReactions(ctx)
			require.Equal(t, test.likes, likesData)
		})
	}
}

// -------------
// --- Reactions
// -------------

func TestKeeper_RegisterReaction(t *testing.T) {
	ctx, k := SetupTestInput()
	var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	reaction := types.NewReaction(
		testOwner,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	store := ctx.KVStore(k.StoreKey)
	key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)

	k.RegisterReaction(ctx, reaction)

	var actualReaction types.Reaction

	bz := store.Get(key)
	k.Cdc.MustUnmarshalBinaryBare(bz, &actualReaction)

	require.Equal(t, reaction, actualReaction)
}

func TestKeeper_DoesReactionForShortcodeExist(t *testing.T) {
	var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

	reaction := types.NewReaction(
		testOwner,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	tests := []struct {
		name           string
		storedReaction types.Reaction
		shortCode      string
		expBool        bool
	}{
		{
			name:           "reaction for given short code exists",
			storedReaction: reaction,
			shortCode:      ":smile:",
			expBool:        true,
		},
		{
			name:           "reaction for the given short code doesn't exist",
			storedReaction: reaction,
			shortCode:      ":test:",
			expBool:        false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)
			store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.storedReaction))

			actualReaction, exist := k.DoesReactionForShortCodeExist(ctx, test.shortCode, reaction.Subspace)
			if test.shortCode == reaction.ShortCode {
				require.True(t, exist)
				require.Equal(t, test.storedReaction, actualReaction)
			} else {
				require.False(t, exist)
				require.Equal(t, types.Reaction{}, actualReaction)
			}
		})
	}
}

func TestKeeper_ListReactions(t *testing.T) {
	ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	var testOwner, _ = sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	reaction := types.NewReaction(
		testOwner,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	reaction2 := types.NewReaction(
		testOwner,
		":thumbs_up:",
		"https://thumbs_up.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)
	reactions := types.Reactions{reaction, reaction2}

	for _, reaction := range reactions {
		key := types.ReactionsStoreKey(reaction.ShortCode, reaction.Subspace)
		store.Set(key, k.Cdc.MustMarshalBinaryBare(reaction))
	}

	actualReactions := k.ListReactions(ctx)

	require.Equal(t, reactions, actualReactions)

}
