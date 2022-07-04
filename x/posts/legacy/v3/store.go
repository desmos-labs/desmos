package v3

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	v2 "github.com/desmos-labs/desmos/v4/x/posts/legacy/v2"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

// MigrateStore performs in-place store migrations from v2 to v3
// The only thing done here is migrating the posts from v2 to v3 to add the tags.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	// Migrate all the posts
	return migratePosts(store, cdc)
}

// migratePosts migrates the posts preset inside the store from v2 to v3
func migratePosts(store sdk.KVStore, cdc codec.BinaryCodec) error {
	prefixStore := prefix.NewStore(store, types.PostPrefix)
	iterator := prefixStore.Iterator(nil, nil)

	// Get all the posts
	var v2Posts []v2.Post
	for ; iterator.Valid(); iterator.Next() {
		var v2Post v2.Post
		cdc.MustUnmarshal(iterator.Value(), &v2Post)
		v2Posts = append(v2Posts, v2Post)
	}

	// Close the iterator
	err := iterator.Close()
	if err != nil {
		return err
	}

	// Convert the posts
	for _, v2Post := range v2Posts {
		v3Post := types.NewPost(
			v2Post.SubspaceID,
			v2Post.SectionID,
			v2Post.ID,
			v2Post.ExternalID,
			v2Post.Text,
			v2Post.Author,
			v2Post.ConversationID,
			migrateEntities(v2Post.Entities),
			nil,
			migratePostReferences(v2Post.ReferencedPosts),
			migrateReplySettings(v2Post.ReplySettings),
			v2Post.CreationDate,
			v2Post.LastEditedDate,
		)

		// Save the post
		store.Set(types.PostStoreKey(v3Post.SubspaceID, v3Post.ID), cdc.MustMarshal(&v3Post))
	}

	return nil
}

// migrateEntities migrates the given entities from v2 to v3
func migrateEntities(v2Entities *v2.Entities) *types.Entities {
	if v2Entities == nil {
		return nil
	}

	return types.NewEntities(
		migrateTags(v2Entities.Hashtags),
		migrateTags(v2Entities.Mentions),
		migrateUrls(v2Entities.Urls),
	)
}

// migrateTags migrates the given tags from v2 to v3
func migrateTags(v2Tags []v2.Tag) []types.TextTag {
	if v2Tags == nil {
		return nil
	}

	v3Tags := make([]types.TextTag, len(v2Tags))
	for i, v2Tag := range v2Tags {
		v3Tags[i] = types.NewTextTag(v2Tag.Start, v2Tag.End, v2Tag.Tag)
	}
	return v3Tags
}

// migrateUrls migrates the given urls from v2 to v3
func migrateUrls(v2Urls []v2.Url) []types.Url {
	if v2Urls == nil {
		return nil
	}

	v3Urls := make([]types.Url, len(v2Urls))
	for i, v2Url := range v2Urls {
		v3Urls[i] = types.NewURL(v2Url.Start, v2Url.End, v2Url.Url, v2Url.DisplayUrl)
	}
	return v3Urls
}

// migratePostReferences migrates the given references from v2 to v3
func migratePostReferences(v2References []v2.PostReference) []types.PostReference {
	if v2References == nil {
		return nil
	}

	v3References := make([]types.PostReference, len(v2References))
	for i, v2Reference := range v2References {
		v3References[i] = types.NewPostReference(
			migratePostReferenceType(v2Reference.Type),
			v2Reference.PostID,
			v2Reference.Position,
		)
	}
	return v3References
}

// migratePostReferenceType migrates the given post reference type from v2 to v3
func migratePostReferenceType(v2Type v2.PostReferenceType) types.PostReferenceType {
	switch v2Type {
	case v2.POST_REFERENCE_TYPE_UNSPECIFIED:
		return types.POST_REFERENCE_TYPE_UNSPECIFIED
	case v2.POST_REFERENCE_TYPE_REPLY:
		return types.POST_REFERENCE_TYPE_REPLY
	case v2.POST_REFERENCE_TYPE_QUOTE:
		return types.POST_REFERENCE_TYPE_QUOTE
	case v2.POST_REFERENCE_TYPE_REPOST:
		return types.POST_REFERENCE_TYPE_REPOST
	default:
		panic(fmt.Errorf("invalid post reference type: %s", v2Type))
	}
}

// migrateReplySettings migrates the given reply setting from v2 to v3
func migrateReplySettings(settings v2.ReplySetting) types.ReplySetting {
	switch settings {
	case v2.REPLY_SETTING_UNSPECIFIED:
		return types.REPLY_SETTING_UNSPECIFIED
	case v2.REPLY_SETTING_EVERYONE:
		return types.REPLY_SETTING_EVERYONE
	case v2.REPLY_SETTING_FOLLOWERS:
		return types.REPLY_SETTING_FOLLOWERS
	case v2.REPLY_SETTING_MUTUAL:
		return types.REPLY_SETTING_MUTUAL
	case v2.REPLY_SETTING_MENTIONS:
		return types.REPLY_SETTING_MENTIONS
	default:
		panic(fmt.Errorf("invalid reply settings value: %s", settings))
	}
}
