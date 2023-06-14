package v7

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

// MigrateStore migrates the x/posts module state from the consensus version 6 to version 7.
// This migration set new owner field to author address for all the posts.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	prefixStore := prefix.NewStore(store, types.PostPrefix)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()

	// Get all the posts
	var posts []types.Post
	for ; iterator.Valid(); iterator.Next() {
		var post types.Post
		cdc.MustUnmarshal(iterator.Value(), &post)
		posts = append(posts, post)
	}

	// Set the owner to its author for posts
	for _, post := range posts {
		post.Owner = post.Author

		// Save the post
		store.Set(types.PostStoreKey(post.SubspaceID, post.ID), cdc.MustMarshal(&post))
	}

	return nil
}
