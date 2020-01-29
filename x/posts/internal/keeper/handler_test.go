package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		expError    error
	}{
		{
			name: "Trying to store post with same id returns expError",
			storedPosts: types.Posts{
				types.NewPost(
					types.PostID(1),
					testPost.ParentID,
					testPost.Message,
					testPost.AllowsComments,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
				),
			},
			lastPostID: types.PostID(0),
			msg: types.NewMsgCreatePost(
				testPost.Message,
				testPost.ParentID,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 1 already exists"),
		},
		{
			name: "Post with new id is stored properly",
			msg: types.NewMsgCreatePost(
				testPost.Message,
				testPost.ParentID,
				testPost.AllowsComments,
				testPost.Subspace,
				testPost.OptionalData,
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
			),
			expPost: types.NewPost(
				types.PostID(1),
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				testPost.Subspace,
				testPost.OptionalData,
				testPost.Created,
				testPost.Creator,
				testPost.Medias,
			),
		},
		{
			name: "Storing a valid post with missing parent id returns expError",
			msg: types.NewMsgCreatePost(
				testPost.Message,
				types.PostID(50),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "parent post with id 50 not found"),
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns expError",
			storedPosts: types.Posts{
				types.NewPost(
					types.PostID(50),
					types.PostID(50),
					"Parent post",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
				),
			},
			msg: types.NewMsgCreatePost(
				testPost.Message,
				types.PostID(50),
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 50 does not allow comments"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			for _, p := range test.storedPosts {
				store.Set(types.PostStoreKey(p.PostID), k.Cdc.MustMarshalBinaryBare(p))
			}

			if test.lastPostID.Valid() {
				store.Set(types.LastPostIDStoreKey, k.Cdc.MustMarshalBinaryBare(&test.lastPostID))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				// Check the post
				var stored types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.expPost.PostID)), &stored)
				assert.True(t, stored.Equals(test.expPost), "Expected: %s, actual: %s", test.expPost, stored)

				// Check the data
				assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.expPost.PostID), res.Data)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypePostCreated,
					sdk.NewAttribute(types.AttributeKeyPostID, test.expPost.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostParentID, test.expPost.ParentID.String()),
					sdk.NewAttribute(types.AttributeKeyCreationTime, test.expPost.Created.String()),
					sdk.NewAttribute(types.AttributeKeyPostOwner, test.expPost.Creator.String()),
				)
				assert.Len(t, ctx.EventManager().Events(), 1)
				assert.Contains(t, ctx.EventManager().Events(), creationEvent)
			}

			// Invalid response
			if res == nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.expError.Error(), err.Error())
			}
		})
	}

}

func Test_handleMsgEditPost(t *testing.T) {
	editor, err := sdk.AccAddressFromBech32("cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63")
	assert.NoError(t, err)

	testData := []struct {
		name       string
		storedPost *types.Post
		msg        types.MsgEditPost
		expError   error
		expPost    types.Post
	}{
		{
			name:       "Post not found",
			storedPost: nil,
			msg:        types.NewMsgEditPost(types.PostID(0), "Edited message", testPostOwner, testPost.Created),
			expError:   sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 0 not found"),
		},
		{
			name:       "Invalid editor",
			storedPost: &testPost,
			msg:        types.NewMsgEditPost(testPost.PostID, "Edited message", editor, testPost.Created),
			expError:   sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner"),
		},
		{
			name:       "Edit date before creation date",
			storedPost: &testPost,
			msg:        types.NewMsgEditPost(testPost.PostID, "Edited message", testPost.Creator, testPost.Created.AddDate(0, 0, -1)),
			expError:   sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "edit date cannot be before creation date"),
		},
		{
			name:       "Valid request is handled properly",
			storedPost: &testPost,
			msg:        types.NewMsgEditPost(testPost.PostID, "Edited message", testPost.Creator, testPost.Created.AddDate(0, 0, 1)),
			expPost: types.Post{
				PostID:         testPost.PostID,
				ParentID:       testPost.ParentID,
				Message:        "Edited message",
				Created:        testPost.Created,
				LastEdited:     testPost.Created.AddDate(0, 0, 1),
				AllowsComments: testPost.AllowsComments,
				Subspace:       testPost.Subspace,
				OptionalData:   testPost.OptionalData,
				Creator:        testPost.Creator,
				Medias:         testPost.Medias,
			},
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if test.storedPost != nil {
				store.Set(types.PostStoreKey(test.storedPost.PostID), k.Cdc.MustMarshalBinaryBare(&test.storedPost))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				assert.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypePostEdited,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostEditTime, test.msg.EditDate.String()),
				))

				var stored types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.storedPost.PostID)), &stored)
				assert.True(t, test.expPost.Equals(stored))
			}

			// Invalid response
			if res == nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.expError.Error(), err.Error())
			}
		})
	}
}

func Test_handleMsgAddPostReaction(t *testing.T) {

	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	assert.NoError(t, err)

	tests := []struct {
		name         string
		existingPost *types.Post
		msg          types.MsgAddPostReaction
		error        error
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgAddPostReaction(types.PostID(0), "like", user),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 0 not found"),
		},
		{
			name:         "Valid message works properly",
			existingPost: &testPost,
			msg:          types.NewMsgAddPostReaction(testPost.PostID, "like", user),
			error:        nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), k.Cdc.MustMarshalBinaryBare(&test.existingPost))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				assert.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypeReactionAdded,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyReactionOwner, test.msg.User.String()),
					sdk.NewAttribute(types.AttributeKeyReactionValue, test.msg.Value),
				))

				var storedPost types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(testPost.PostID)), &storedPost)
				assert.True(t, test.existingPost.Equals(storedPost))

				var storedReactions types.Reactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				assert.Contains(t, storedReactions, types.NewReaction(test.msg.Value, test.msg.User))
			}

			// Invalid response
			if res == nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.error.Error(), err.Error())
			}
		})
	}
}

func Test_handleMsgRemovePostReaction(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	assert.NoError(t, err)

	reaction := types.NewReaction("like", user)
	tests := []struct {
		name             string
		existingPost     *types.Post
		existingReaction *types.Reaction
		msg              types.MsgRemovePostReaction
		error            error
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgRemovePostReaction(types.PostID(0), user, "like"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 0 not found"),
		},
		{
			name:         "Reaction not found",
			existingPost: &testPost,
			msg:          types.NewMsgRemovePostReaction(testPost.PostID, user, "like"),
			error:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("cannot remove the reaction with value like from user %s as it does not exist", user)),
		},
		{
			name:             "Valid message works properly",
			existingPost:     &testPost,
			existingReaction: &reaction,
			msg:              types.NewMsgRemovePostReaction(testPost.PostID, user, reaction.Value),
			error:            nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), k.Cdc.MustMarshalBinaryBare(&test.existingPost))
			}

			if test.existingReaction != nil {
				store.Set(
					types.PostReactionsStoreKey(test.existingPost.PostID),
					k.Cdc.MustMarshalBinaryBare(&types.Reactions{*test.existingReaction}),
				)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				assert.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypePostReactionRemoved,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyReactionOwner, test.msg.User.String()),
					sdk.NewAttribute(types.AttributeKeyReactionValue, test.msg.Reaction),
				))

				var storedPost types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(testPost.PostID)), &storedPost)
				assert.True(t, test.existingPost.Equals(storedPost))

				var storedReactions types.Reactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				assert.NotContains(t, storedReactions, test.existingReaction)
			}

			// Invalid response
			if res == nil {
				assert.NotNil(t, err)
				assert.Equal(t, test.error.Error(), err.Error())
			}
		})
	}
}
