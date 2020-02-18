package keeper

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		Cdc:      cdc,
	}
}

// -------------
// --- Posts
// -------------

// GetLastPostID returns the last post id that has been used
func (k Keeper) GetLastPostID(ctx sdk.Context) types.PostID {
	store := ctx.KVStore(k.StoreKey)
	if !store.Has(types.LastPostIDStoreKey) {
		return types.PostID(0)
	}

	var id types.PostID
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.LastPostIDStoreKey), &id)
	return id
}

// SavePostHashtag allows to save the hashtag association with the given postID.
// It assumes that the postID is associated with an existent post
func (k Keeper) SavePostHashtag(ctx sdk.Context, hashtag string, postID types.PostID) {
	store := ctx.KVStore(k.StoreKey)
	postIDs := k.GetHashtagAssociatedPosts(ctx, hashtag)
	postIDs, appended := postIDs.AppendIfMissing(postID)
	if appended {
		store.Set(types.HashtagStoreKey(hashtag), k.Cdc.MustMarshalBinaryBare(&postIDs))
	}
}

// RemovePostHashtags allows to remove all the hashtags associated with a postID.
// It assumes that there's already and association between the given postID and hashtags
func (k Keeper) RemovePostHashtags(ctx sdk.Context, postID types.PostID, hashtags []string) {
	store := ctx.KVStore(k.StoreKey)
	for _, hashtag := range hashtags {
		postIDs := k.GetHashtagAssociatedPosts(ctx, hashtag)
		postIDs, removed := postIDs.RemoveIfPresent(postID)
		if removed {
			store.Set(types.HashtagStoreKey(hashtag), k.Cdc.MustMarshalBinaryBare(&postIDs))
		}
	}
}

// GetHashtagAssociatedPosts returns the posts IDs associated with the given hashtag
func (k Keeper) GetHashtagAssociatedPosts(ctx sdk.Context, hashtag string) types.PostIDs {
	store := ctx.KVStore(k.StoreKey)
	var postIDs types.PostIDs
	bz := store.Get(types.HashtagStoreKey(hashtag))
	k.Cdc.MustUnmarshalBinaryBare(bz, &postIDs)

	return postIDs
}

// GetHashtags allows to returns the list of hashtags that have been stored inside the given context
func (k Keeper) GetHashtags(ctx sdk.Context) map[string]types.PostIDs {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.HashtagPrefix)

	hashtagsData := map[string]types.PostIDs{}
	for ; iterator.Valid(); iterator.Next() {
		var postIDs types.PostIDs
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &postIDs)
		hashtagsData[string(bytes.TrimPrefix(iterator.Key(), types.HashtagPrefix))] = postIDs
	}

	return hashtagsData
}

// SavePost allows to save the given post inside the current context.
// It assumes that the given post has already been validated.
// If another post has the same ID of the given post, the old post will be overridden
func (k Keeper) SavePost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.StoreKey)

	// Save the post
	store.Set(types.PostStoreKey(post.PostID), k.Cdc.MustMarshalBinaryBare(&post))

	// Set the last post id only if the current post has a greater one than the last one stored
	if id := post.PostID; id > k.GetLastPostID(ctx) {
		store.Set(types.LastPostIDStoreKey, k.Cdc.MustMarshalBinaryBare(&id))
	}

	// Save the comments to the parent post, if it is valid
	if post.ParentID.Valid() {
		parentCommentsKey := types.PostCommentsStoreKey(post.ParentID)

		var commentsIDs types.PostIDs
		k.Cdc.MustUnmarshalBinaryBare(store.Get(parentCommentsKey), &commentsIDs)

		if editedIDs, appended := commentsIDs.AppendIfMissing(post.PostID); appended {
			store.Set(parentCommentsKey, k.Cdc.MustMarshalBinaryBare(&editedIDs))
		}
	}
}

// GetPost returns the post having the given id inside the current context.
// If no post having the given id can be found inside the current context, false will be returned.
func (k Keeper) GetPost(ctx sdk.Context, id types.PostID) (post types.Post, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := types.PostStoreKey(id)
	if !store.Has(key) {
		return types.Post{}, false
	}

	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &post)
	return post, true
}

// GetPostChildrenIDs returns the IDs of all the children posts associated to the post
// having the given postID
// nolint: interfacer
func (k Keeper) GetPostChildrenIDs(ctx sdk.Context, postID types.PostID) types.PostIDs {
	store := ctx.KVStore(k.StoreKey)

	var postIDs types.PostIDs
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostCommentsStoreKey(postID)), &postIDs)
	return postIDs
}

// GetPosts returns the list of all the posts that are stored into the current state.
func (k Keeper) GetPosts(ctx sdk.Context) (posts types.Posts) {
	posts = types.Posts{}
	k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
		posts = append(posts, post)
		return false
	})

	return posts
}

// GetPostsFiltered retrieves posts filtered by a given set of params which
// include pagination parameters along with the creator address, the parent id and the creation time.
//
// NOTE: If no filters are provided, all posts will be returned in paginated
// form.
func (k Keeper) GetPostsFiltered(ctx sdk.Context, params types.QueryPostsParams) types.Posts {
	filteredPosts := types.Posts{}
	k.IteratePosts(ctx, func(_ int64, post types.Post) (stop bool) {
		matchParentID, matchCreationTime, matchAllowsComments, matchSubspace, matchCreator, matchHashtags := true, true, true, true, true, true

		// match parent id if valid
		if params.ParentID != nil {
			matchParentID = params.ParentID.Equals(post.ParentID)
		}

		// match creation time if valid height
		if params.CreationTime != nil {
			matchCreationTime = params.CreationTime.Equal(post.Created)
		}

		// match allows comments
		if params.AllowsComments != nil {
			matchAllowsComments = *params.AllowsComments == post.AllowsComments
		}

		// match subspace if provided
		if len(params.Subspace) > 0 {
			matchSubspace = params.Subspace == post.Subspace
		}

		// match creator address (if supplied)
		if len(params.Creator) > 0 {
			matchCreator = params.Creator.Equals(post.Creator)
		}

		// match hashtags if provided
		if len(params.Hashtags) > 0 {
			for _, hashtag := range params.Hashtags {
				postsIDs := k.GetHashtagAssociatedPosts(ctx, hashtag)
				if matchHashtags = postsIDs.Contains(post.PostID); !matchHashtags {
					break
				}
			}
		}

		if matchParentID && matchCreationTime && matchAllowsComments && matchSubspace && matchCreator && matchHashtags {
			filteredPosts = append(filteredPosts, post)
		}

		return false
	})

	// Sort the posts
	sort.Slice(filteredPosts, func(i, j int) bool {
		var result bool
		first, second := filteredPosts[i], filteredPosts[j]

		switch params.SortBy {
		case types.PostSortByCreationDate:
			result = first.Created.Before(second.Created)
			if params.SortOrder == types.PostSortOrderDescending {
				result = first.Created.After(second.Created)
			}

		default:
			result = first.PostID < second.PostID
			if params.SortOrder == types.PostSortOrderDescending {
				result = first.PostID > second.PostID
			}
		}

		// This should never be reached
		return result
	})

	// Default page
	page := params.Page
	if page == 0 {
		page = 1
	}

	start, end := client.Paginate(len(filteredPosts), page, params.Limit, 100)
	if start < 0 || end < 0 {
		filteredPosts = types.Posts{}
	} else {
		filteredPosts = filteredPosts[start:end]
	}

	return filteredPosts
}

// SavePollAnswers save the poll's answers associated with the given postID inside the current context
// It assumes that the post exists and has a Poll inside it.
// If userAnswersDetails are already present, the old ones will be overridden.
func (k Keeper) SavePollAnswers(ctx sdk.Context, postID types.PostID, userPollAnswers types.UserAnswer) {
	store := ctx.KVStore(k.StoreKey)

	sort.Slice(
		userPollAnswers.Answers,
		func(i, j int) bool { return userPollAnswers.Answers[i] < userPollAnswers.Answers[j] },
	)

	usersAnswersDetails := k.GetPollAnswers(ctx, postID)

	if usersAnswersDetails, appended := usersAnswersDetails.AppendIfMissingOrIfUsersEquals(userPollAnswers); appended {
		store.Set(types.PollAnswersStoreKey(postID), k.Cdc.MustMarshalBinaryBare(&usersAnswersDetails))
	}

}

// GetPollAnswers returns the list of all the post polls answers associated with the given postID that are stored into the current state.
func (k Keeper) GetPollAnswers(ctx sdk.Context, postID types.PostID) types.UserAnswers {
	store := ctx.KVStore(k.StoreKey)

	var usersAnswersDetails types.UserAnswers
	answersBz := store.Get(types.PollAnswersStoreKey(postID))

	k.Cdc.MustUnmarshalBinaryBare(answersBz, &usersAnswersDetails)

	return usersAnswersDetails
}

// GetPollAnswersMap allows to returns the list of answers that have been stored inside the given context
func (k Keeper) GetPollAnswersMap(ctx sdk.Context) map[types.PostID]types.UserAnswers {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PollAnswersStorePrefix)

	usersAnswersData := map[types.PostID]types.UserAnswers{}
	for ; iterator.Valid(); iterator.Next() {
		var userAnswers types.UserAnswers
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &userAnswers)
		idBytes := bytes.TrimPrefix(iterator.Key(), types.PollAnswersStorePrefix)
		postID, err := types.ParsePostID(string(idBytes))
		if err != nil {
			// This should never happen
			panic(err)
		}

		usersAnswersData[postID] = userAnswers
	}

	return usersAnswersData
}

// GetPollAnswersByUser retrieves post poll answers associated to the given ID and filtered by user
func (k Keeper) GetPollAnswersByUser(ctx sdk.Context, postID types.PostID, user sdk.AccAddress) []types.AnswerID {
	postPollAnswers := k.GetPollAnswers(ctx, postID)

	for _, postPollAnswers := range postPollAnswers {
		if user.Equals(postPollAnswers.User) {
			return postPollAnswers.Answers
		}
	}
	return nil
}

// -------------
// --- Reactions
// -------------

// SaveReaction allows to save the given reaction inside the store.
// It assumes that the given reaction is valid.
// If another reaction from the same user for the same post and with the same value exists, returns an expError.
// nolint: interfacer
func (k Keeper) SaveReaction(ctx sdk.Context, postID types.PostID, reaction types.Reaction) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.PostReactionsStoreKey(postID)

	// Get the existent reactions
	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check for double reactions
	if reactions.ContainsReactionFrom(reaction.Owner, reaction.Value) {
		return fmt.Errorf("%s has already reacted with %s to the post with id %s",
			reaction.Owner, reaction.Value, postID)
	}

	// Save the new reaction
	reactions = append(reactions, reaction)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&reactions))

	return nil
}

// RemoveReaction removes the reaction from the given user from the post having the
// given postID. If no reaction with the same value was previously added from the given user, an expError
// is returned.
// nolint: interfacer
func (k Keeper) RemoveReaction(ctx sdk.Context, postID types.PostID, user sdk.AccAddress, value string) error {
	store := ctx.KVStore(k.StoreKey)
	key := types.PostReactionsStoreKey(postID)

	// Get the existing reactions
	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(key), &reactions)

	// Check if the user exists
	if !reactions.ContainsReactionFrom(user, value) {
		return fmt.Errorf("cannot remove the reaction with value %s from user %s as it does not exist",
			value, user)
	}

	// Remove and save the reactions list
	if newLikes, edited := reactions.RemoveReaction(user, value); edited {
		if len(newLikes) == 0 {
			store.Delete(key)
		} else {
			store.Set(key, k.Cdc.MustMarshalBinaryBare(&newLikes))
		}
	}

	return nil
}

// GetPostReactions returns the list of reactions that has been associated to the post having the given id
// nolint: interfacer
func (k Keeper) GetPostReactions(ctx sdk.Context, postID types.PostID) types.Reactions {
	store := ctx.KVStore(k.StoreKey)

	var reactions types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(postID)), &reactions)
	return reactions
}

// GetReactions allows to returns the list of reactions that have been stored inside the given context
func (k Keeper) GetReactions(ctx sdk.Context) map[types.PostID]types.Reactions {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.PostReactionsStorePrefix)

	reactionsData := map[types.PostID]types.Reactions{}
	for ; iterator.Valid(); iterator.Next() {
		var postLikes types.Reactions
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &postLikes)
		idBytes := bytes.TrimPrefix(iterator.Key(), types.PostReactionsStorePrefix)
		postID, err := types.ParsePostID(string(idBytes))
		if err != nil {
			// This should never verify
			panic(err)
		}

		reactionsData[postID] = postLikes
	}

	return reactionsData
}
