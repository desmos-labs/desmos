package keeper_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_SaveReaction(t *testing.T) {
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name           string
		storedLikes    types.Reactions
		postID         types.PostID
		like           types.PostReaction
		error          error
		expectedStored types.Reactions
	}{
		{
			name:           "PostReaction from same user already present returns expError",
			storedLikes:    types.Reactions{types.NewPostReaction("like", liker)},
			postID:         types.PostID(10),
			like:           types.NewPostReaction("like", liker),
			error:          fmt.Errorf("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 has already reacted with like to the post with id 10"),
			expectedStored: types.Reactions{types.NewPostReaction("like", liker)},
		},
		{
			name:           "First liker is stored properly",
			storedLikes:    types.Reactions{},
			postID:         types.PostID(15),
			like:           types.NewPostReaction("like", liker),
			error:          nil,
			expectedStored: types.Reactions{types.NewPostReaction("like", liker)},
		},
		{
			name:        "Second liker is stored properly",
			storedLikes: types.Reactions{types.NewPostReaction("like", liker)},
			postID:      types.PostID(87),
			like:        types.NewPostReaction("like", otherLiker),
			error:       nil,
			expectedStored: types.Reactions{
				types.NewPostReaction("like", liker),
				types.NewPostReaction("like", otherLiker),
			},
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

			err := k.SavePostReaction(ctx, test.postID, test.like)
			require.Equal(t, test.error, err)

			var stored types.Reactions
			k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(test.postID)), &stored)
			require.Equal(t, test.expectedStored, stored)
		})
	}
}

func TestKeeper_RemoveReaction(t *testing.T) {
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	tests := []struct {
		name           string
		storedLikes    types.Reactions
		postID         types.PostID
		liker          sdk.AccAddress
		value          string
		error          error
		expectedStored types.Reactions
	}{
		{
			name:           "PostReaction from same liker is removed properly",
			storedLikes:    types.Reactions{types.NewPostReaction("like", liker)},
			postID:         types.PostID(10),
			liker:          liker,
			value:          "like",
			error:          nil,
			expectedStored: types.Reactions{},
		},
		{
			name:           "Non existing reaction returns error - Creator",
			storedLikes:    types.Reactions{},
			postID:         types.PostID(15),
			liker:          liker,
			value:          "like",
			error:          fmt.Errorf("cannot remove the reaction with value like from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.Reactions{},
		},
		{
			name:           "Non existing reaction returns error - Value",
			storedLikes:    types.Reactions{types.NewPostReaction("like", liker)},
			postID:         types.PostID(15),
			liker:          liker,
			value:          "reaction",
			error:          fmt.Errorf("cannot remove the reaction with value reaction from user cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4 as it does not exist"),
			expectedStored: types.Reactions{types.NewPostReaction("like", liker)},
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

			var stored types.Reactions
			k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(test.postID)), &stored)

			require.Len(t, stored, len(test.expectedStored))
			for index, like := range test.expectedStored {
				require.Equal(t, like, stored[index])
			}
		})
	}
}

func TestKeeper_GetPostReactions(t *testing.T) {
	liker, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	otherLiker, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name   string
		likes  types.Reactions
		postID types.PostID
	}{
		{
			name:   "Empty list are returned properly",
			likes:  types.Reactions{},
			postID: types.PostID(10),
		},
		{
			name: "Valid list of likes is returned properly",
			likes: types.Reactions{
				types.NewPostReaction("like", otherLiker),
				types.NewPostReaction("like", liker),
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, l := range test.likes {
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
	liker1, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	liker2, err := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	require.NoError(t, err)

	tests := []struct {
		name  string
		likes map[types.PostID]types.Reactions
	}{
		{
			name:  "Empty likes data are returned correctly",
			likes: map[types.PostID]types.Reactions{},
		},
		{
			name: "Non empty likes data are returned correcly",
			likes: map[types.PostID]types.Reactions{
				types.PostID(5): {
					types.NewPostReaction("like", liker1),
					types.NewPostReaction("like", liker2),
				},
				types.PostID(10): {
					types.NewPostReaction("like", liker1),
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
				store.Set(types.PostReactionsStoreKey(postID), k.Cdc.MustMarshalBinaryBare(likes))
			}

			likesData := k.GetReactions(ctx)
			require.Equal(t, test.likes, likesData)
		})
	}
}
