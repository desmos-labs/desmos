package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"
)

// ---------------------------
// --- handleMsgCreatePost
// ---------------------------

func Test_handleMsgCreatePost(t *testing.T) {
	tests := []struct {
		name        string
		storedPosts types.Posts
		lastPostID  types.PostID
		msg         types.MsgCreatePost
		expPost     types.Post
		expError    string
	}{
		{
			name: "Trying to store post with same id returns error",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(1), testPost.ParentID, testPost.Message, testPost.AllowsComments, "desmos", map[string]string{}, testPost.Created.Int64(), testPost.Owner),
			},
			lastPostID: types.PostID(0),
			msg:        types.NewMsgCreatePost(testPost.Message, testPost.ParentID, testPost.AllowsComments, "desmos", map[string]string{}, testPost.Owner),
			expError:   "Post with id 1 already exists",
		},
		{
			name:    "Post with new id is stored properly",
			msg:     types.NewMsgCreatePost(testPost.Message, testPost.ParentID, false, "desmos", map[string]string{}, testPost.Owner),
			expPost: types.NewPost(types.PostID(1), testPost.ParentID, testPost.Message, testPost.AllowsComments, "desmos", map[string]string{}, 0, testPost.Owner),
		},
		{
			name:     "Storing a valid post with missing parent id returns error",
			msg:      types.NewMsgCreatePost(testPost.Message, types.PostID(50), false, "desmos", map[string]string{}, testPost.Owner),
			expError: "Parent post with id 50 not found",
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns error",
			storedPosts: types.Posts{
				types.NewPost(types.PostID(50), types.PostID(50), "Parent post", false, "desmos", map[string]string{}, 0, testPost.Owner),
			},
			msg:      types.NewMsgCreatePost(testPost.Message, types.PostID(50), false, "desmos", map[string]string{}, testPost.Owner),
			expError: "Post with id 50 does not allow comments",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			for _, p := range test.storedPosts {
				store.Set([]byte(types.PostStorePrefix+p.PostID.String()), k.Cdc.MustMarshalBinaryBare(p))
			}

			if test.lastPostID.Valid() {
				store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(&test.lastPostID))
			}

			handler := keeper.NewHandler(k)
			res := handler(ctx, test.msg)

			if len(test.expError) != 0 {
				assert.False(t, res.IsOK())
				assert.Contains(t, res.Log, test.expError)
			}

			if len(test.expError) == 0 {
				assert.True(t, res.IsOK())

				// Check the post
				var stored types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+test.expPost.PostID.String())), &stored)
				assert.True(t, stored.Equals(test.expPost))

				// Check the data
				assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.expPost.PostID), res.Data)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypePostCreated,
					sdk.NewAttribute(types.AttributeKeyPostID, test.expPost.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostParentID, test.expPost.ParentID.String()),
					sdk.NewAttribute(types.AttributeKeyCreationTime, test.expPost.Created.String()),
					sdk.NewAttribute(types.AttributeKeyPostOwner, test.expPost.Owner.String()),
				)
				assert.Len(t, ctx.EventManager().Events(), 1)
				assert.Contains(t, ctx.EventManager().Events(), creationEvent)
			}
		})
	}

}

// --------------------------
// --- handleMsgEditPost
// --------------------------

func Test_handleMSgEditPost_invalid_requests(t *testing.T) {
	editor, _ := sdk.AccAddressFromBech32("cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63")
	testData := []struct {
		name        string
		storedPost  *types.Post
		msg         types.MsgEditPost
		blockHeight int64
		error       string
	}{
		{
			name:       "Post not found",
			storedPost: nil,
			msg: types.MsgEditPost{
				PostID:  types.PostID(0),
				Message: "Edited message",
				Editor:  testPostOwner,
			},
			error: "Post with id 0 not found",
		},
		{
			name:       "Invalid editor",
			storedPost: &testPost,
			msg: types.MsgEditPost{
				PostID:  testPost.PostID,
				Message: "Edited message",
				Editor:  editor,
			},
			error: "Incorrect owner",
		}, {
			name:        "Edit date before creation date",
			storedPost:  &testPost,
			blockHeight: testPost.Created.Int64() - 1,
			msg: types.MsgEditPost{
				PostID:  testPost.PostID,
				Message: "Edited message",
				Editor:  testPost.Owner,
			},
			error: "Edit date cannot be before creation date",
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.blockHeight != 0 {
				ctx = ctx.WithBlockHeight(test.blockHeight)
			}

			store := ctx.KVStore(k.StoreKey)

			if test.storedPost != nil {
				store.Set(
					[]byte(types.PostStorePrefix+test.storedPost.PostID.String()),
					k.Cdc.MustMarshalBinaryBare(&test.storedPost),
				)
			}

			handler := keeper.NewHandler(k)
			res := handler(ctx, test.msg)

			// Check the response
			assert.False(t, res.IsOK())
			assert.Contains(t, res.Log, test.error)
			assert.Empty(t, res.Events, 0)

			// Check the events
			assert.Empty(t, ctx.EventManager().Events())
		})
	}
}

func Test_handleMsgEditPost_valid_request(t *testing.T) {
	ctx, k := SetupTestInput()
	ctx = ctx.WithBlockHeight(testPost.Created.Int64() + 1)

	// Insert the post
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.PostStorePrefix+testPost.PostID.String()), k.Cdc.MustMarshalBinaryBare(&testPost))

	// Handle the message
	msg := types.NewMsgEditPost(testPost.PostID, "Edited message", testPost.Owner)
	handler := keeper.NewHandler(k)
	res := handler(ctx, msg)

	// Check the response
	assert.True(t, res.IsOK())
	assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(testPost.PostID), res.Data)

	// Check the events
	editEvent := sdk.NewEvent(
		types.EventTypePostEdited,
		sdk.NewAttribute(types.AttributeKeyPostID, testPost.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostEditTime, strconv.FormatInt(ctx.BlockHeight(), 10)),
	)
	assert.Len(t, ctx.EventManager().Events(), 1)
	assert.Equal(t, ctx.EventManager().Events(), res.Events)
	assert.Contains(t, ctx.EventManager().Events(), editEvent)

	// Check the stored post
	expected := types.Post{
		PostID:       testPost.PostID,
		ParentID:     testPost.ParentID,
		Message:      msg.Message,
		Owner:        testPost.Owner,
		Created:      testPost.Created,
		Subspace:     "desmos",
		OptionalData: map[string]string{},
		LastEdited:   sdk.NewInt(ctx.BlockHeight()),
	}

	var stored types.Post
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+testPost.PostID.String())), &stored)
	assert.True(t, stored.Equals(expected))
}

// --------------------
// --- handleMsgAddPostReaction
// --------------------

func Test_handleMsgLikePost_invalid_requests(t *testing.T) {

	liker, _ := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	tests := []struct {
		name         string
		existingPost *types.Post
		msg          types.MsgAddPostReaction
		blockHeight  int64
		error        string
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgAddPostReaction(types.PostID(0), "like", liker),
			error: "Post with id 0 not found",
		},
		{
			name:         "Reaction date before post date",
			existingPost: &testPost,
			blockHeight:  testPost.Created.Int64() - 1,
			msg:          types.NewMsgAddPostReaction(testPost.PostID, "like", liker),
			error:        "Reaction cannot have a creation time before the post itself",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			if test.blockHeight != 0 {
				ctx = ctx.WithBlockHeight(test.blockHeight)
			}

			store := ctx.KVStore(k.StoreKey)
			if test.existingPost != nil {
				store.Set(
					[]byte(types.PostStorePrefix+test.existingPost.PostID.String()),
					k.Cdc.MustMarshalBinaryBare(&test.existingPost),
				)
			}

			handler := keeper.NewHandler(k)
			res := handler(ctx, test.msg)

			// Check response
			assert.False(t, res.IsOK())
			assert.Contains(t, res.Log, test.error)

			// Events
			assert.Empty(t, ctx.EventManager().Events())
		})
	}
}

func Test_handleMsgLikePost_valid_request(t *testing.T) {
	ctx, k := SetupTestInput()
	ctx = ctx.WithBlockHeight(testPost.Created.Int64())

	// Insert the post
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.PostStorePrefix+testPost.PostID.String()), k.Cdc.MustMarshalBinaryBare(&testPost))

	// Handle the message
	liker, _ := sdk.AccAddressFromBech32("cosmos1dshanwvhmq4c5jk9a3ywtuyex426cflq5l4mqp")
	msg := types.NewMsgAddPostReaction(testPost.PostID, "like", liker)

	handler := keeper.NewHandler(k)
	res := handler(ctx, msg)

	// Check the response
	assert.True(t, res.IsOK())
	//assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(expectedLikeID), res.Data)

	// Check the events
	creationEvent := sdk.NewEvent(
		types.EventTypeReactionAdded,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyReactionValue, msg.Value),
	)
	assert.Len(t, ctx.EventManager().Events(), 1)
	assert.Equal(t, ctx.EventManager().Events(), res.Events)
	assert.Contains(t, ctx.EventManager().Events(), creationEvent)

	// Check that the post has a new liker
	expectedPost := types.Post{
		PostID:       testPost.PostID,
		ParentID:     testPost.ParentID,
		Message:      testPost.Message,
		LastEdited:   testPost.LastEdited,
		Owner:        testPost.Owner,
		Subspace:     "desmos",
		OptionalData: map[string]string{},
		Created:      testPost.Created,
	}

	var storedPost types.Post
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+testPost.PostID.String())), &storedPost)
	assert.True(t, storedPost.Equals(expectedPost))

	// Check the stored liker
	expectedLikes := types.Reactions{types.NewReaction("like", ctx.BlockHeight(), msg.User)}

	var storedLikes types.Reactions
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostReactionsStorePrefix+storedPost.PostID.String())), &storedLikes)
	assert.Equal(t, expectedLikes, storedLikes)
}
