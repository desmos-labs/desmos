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

func Test_handleMsgCreatePost_returns_error_with_existing_post_id(t *testing.T) {
	ctx, k := SetupTestInput()

	msg := types.MsgCreatePost{
		ParentID:      testPost.ParentID,
		Message:       testPost.Message,
		Owner:         testPost.Owner,
		Namespace:     testPost.Namespace,
		ExternalOwner: testPost.ExternalOwner,
	}

	existing := testPost
	existing.PostID = types.PostID(1)

	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(types.PostID(0)))
	store.Set([]byte(types.PostStorePrefix+existing.PostID.String()), k.Cdc.MustMarshalBinaryBare(&existing))

	handler := keeper.NewHandler(k)
	res := handler(ctx, msg)

	// Check the response
	assert.False(t, res.IsOK())
	assert.Contains(t, res.Log, "Post with id 1 already exists")
	assert.Empty(t, res.Events, 0)

	// Check the events
	assert.Len(t, ctx.EventManager().Events(), 1)
	expected := sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyAction, types.ActionCreatePost),
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
	)
	assert.Contains(t, ctx.EventManager().Events(), expected)
}

func Test_handleMsgCreatePost_valid_request(t *testing.T) {
	ctx, k := SetupTestInput()

	expectedPostID := types.PostID(1)
	msg := types.MsgCreatePost{
		ParentID:      testPost.ParentID,
		Message:       testPost.Message,
		Owner:         testPost.Owner,
		Namespace:     testPost.Namespace,
		ExternalOwner: testPost.ExternalOwner,
	}

	handler := keeper.NewHandler(k)
	res := handler(ctx, msg)

	// Check the response
	assert.True(t, res.IsOK())
	assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(expectedPostID), res.Data)

	// Check the events
	creationEvent := sdk.NewEvent(
		types.EventTypeCreatePost,
		sdk.NewAttribute(types.AttributeKeyPostID, expectedPostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostParentID, msg.ParentID.String()),
		sdk.NewAttribute(types.AttributeKeyCreationTime, strconv.FormatInt(ctx.BlockHeight(), 10)),
		sdk.NewAttribute(types.AttributeKeyPostOwner, msg.Owner.String()),
		sdk.NewAttribute(types.AttributeKeyNamespace, msg.Namespace),
		sdk.NewAttribute(types.AttributeKeyExternalOwner, msg.ExternalOwner),
	)
	assert.Len(t, ctx.EventManager().Events(), 2)
	assert.Equal(t, ctx.EventManager().Events(), res.Events)
	assert.Contains(t, ctx.EventManager().Events(), creationEvent)

	// Check the stored post
	expected := types.Post{
		PostID:        expectedPostID,
		ParentID:      msg.ParentID,
		Message:       msg.Message,
		LastEdited:    0,
		Owner:         msg.Owner,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
	}

	var stored types.Post
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+expectedPostID.String())), &stored)
	assert.Equal(t, expected, stored)
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
			blockHeight: testPost.Created - 1,
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
			assert.Len(t, ctx.EventManager().Events(), 1)
			expectedEvent := sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(sdk.AttributeKeyAction, types.ActionEditPost),
				sdk.NewAttribute(sdk.AttributeKeySender, test.msg.Editor.String()),
			)
			assert.Contains(t, ctx.EventManager().Events(), expectedEvent)
		})
	}
}

func Test_handleMsgEditPost_valid_request(t *testing.T) {
	ctx, k := SetupTestInput()
	ctx = ctx.WithBlockHeight(testPost.Created + 1)

	// Insert the post
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.PostStorePrefix+testPost.PostID.String()), k.Cdc.MustMarshalBinaryBare(&testPost))

	// Handle the message
	msg := types.MsgEditPost{
		PostID:  testPost.PostID,
		Message: "Edited message",
		Editor:  testPost.Owner,
	}

	handler := keeper.NewHandler(k)
	res := handler(ctx, msg)

	// Check the response
	assert.True(t, res.IsOK())
	assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(testPost.PostID), res.Data)

	// Check the events
	editEvent := sdk.NewEvent(
		types.EventTypeEditPost,
		sdk.NewAttribute(types.AttributeKeyPostID, testPost.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostEditTime, strconv.FormatInt(ctx.BlockHeight(), 10)),
	)
	assert.Len(t, ctx.EventManager().Events(), 2)
	assert.Equal(t, ctx.EventManager().Events(), res.Events)
	assert.Contains(t, ctx.EventManager().Events(), editEvent)

	// Check the stored post
	expected := types.Post{
		PostID:        testPost.PostID,
		ParentID:      testPost.ParentID,
		Message:       msg.Message,
		Owner:         testPost.Owner,
		Namespace:     testPost.Namespace,
		ExternalOwner: testPost.ExternalOwner,
		Created:       testPost.Created,
		LastEdited:    ctx.BlockHeight(),
	}

	var stored types.Post
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+testPost.PostID.String())), &stored)
	assert.Equal(t, expected, stored)
}

// --------------------
// --- handleMsgLike
// --------------------

func Test_handleMsgLikePost_invalid_requests(t *testing.T) {

	liker, _ := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	tests := []struct {
		name         string
		existingPost *types.Post
		msg          types.MsgLikePost
		blockHeight  int64
		error        string
	}{
		{
			name: "Post not found",
			msg: types.MsgLikePost{
				PostID:        types.PostID(0),
				Liker:         liker,
				Namespace:     "cosmos",
				ExternalLiker: "cosmos14xf748kl34mhn54zymlnppvg7pq58f0q0u968d",
			},
			error: "Post with id 0 not found",
		},
		{
			name:         "Like date before post date",
			existingPost: &testPost,
			blockHeight:  testPost.Created - 1,
			msg: types.MsgLikePost{
				PostID:        testPost.PostID,
				Liker:         liker,
				Namespace:     "cosmos",
				ExternalLiker: "cosmos14xf748kl34mhn54zymlnppvg7pq58f0q0u968d",
			},
			error: "Like cannot have a creation time before the post itself",
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
			assert.Len(t, ctx.EventManager().Events(), 1)
			expectedEvent := sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
				sdk.NewAttribute(sdk.AttributeKeyAction, types.ActionLikePost),
				sdk.NewAttribute(sdk.AttributeKeySender, test.msg.Liker.String()),
			)
			assert.Contains(t, ctx.EventManager().Events(), expectedEvent)
		})
	}
}

func Test_handleMsgLikePost_valid_request(t *testing.T) {
	ctx, k := SetupTestInput()
	ctx = ctx.WithBlockHeight(testPost.Created)

	// Insert the post
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.PostStorePrefix+testPost.PostID.String()), k.Cdc.MustMarshalBinaryBare(&testPost))

	// Handle the message
	expectedLikeID := types.LikeID(1)
	liker, _ := sdk.AccAddressFromBech32("cosmos1dshanwvhmq4c5jk9a3ywtuyex426cflq5l4mqp")
	msg := types.MsgLikePost{
		PostID:        testPost.PostID,
		Liker:         liker,
		Namespace:     "cosmos",
		ExternalLiker: "cosmos14xf748kl34mhn54zymlnppvg7pq58f0q0u968d",
	}

	handler := keeper.NewHandler(k)
	res := handler(ctx, msg)

	// Check the response
	assert.True(t, res.IsOK())
	assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(expectedLikeID), res.Data)

	// Check the events
	creationEvent := sdk.NewEvent(
		types.EventTypeLikePost,
		sdk.NewAttribute(types.AttributeKeyLikeID, expectedLikeID.String()),
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyNamespace, msg.Namespace),
		sdk.NewAttribute(types.AttributeKeyLikeOwner, msg.Liker.String()),
	)
	assert.Len(t, ctx.EventManager().Events(), 2)
	assert.Equal(t, ctx.EventManager().Events(), res.Events)
	assert.Contains(t, ctx.EventManager().Events(), creationEvent)

	// Check that the post has a new like
	expectedPost := types.Post{
		PostID:        testPost.PostID,
		ParentID:      testPost.ParentID,
		Message:       testPost.Message,
		LastEdited:    testPost.LastEdited,
		Owner:         testPost.Owner,
		Namespace:     testPost.Namespace,
		ExternalOwner: testPost.ExternalOwner,
		Created:       testPost.Created,
	}

	var storedPost types.Post
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+testPost.PostID.String())), &storedPost)
	assert.Equal(t, expectedPost, storedPost)

	// Check the stored like
	expectedLike := types.Like{
		LikeID:        expectedLikeID,
		PostID:        msg.PostID,
		Owner:         msg.Liker,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalLiker,
		Created:       ctx.BlockHeight(),
	}

	var storedLike types.Like
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LikesStorePrefix+expectedLikeID.String())), &storedLike)
	assert.Equal(t, expectedLike, storedLike)
}
