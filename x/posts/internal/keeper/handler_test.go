package keeper_test

import (
	"fmt"
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
			name: "Trying to store post with same id returns expError",
			storedPosts: types.Posts{
				types.NewTextPost(types.PostID(1), testPost.ParentID, testPost.Message, testPost.AllowsComments, "desmos", map[string]string{}, testPost.Created.Int64(), testPost.Creator),
			},
			lastPostID: types.PostID(0),
			msg:        types.NewMsgCreatePost(testPost.Message, testPost.ParentID, testPost.AllowsComments, "desmos", map[string]string{}, testPost.Creator),
			expError:   "Post with id 1 already exists",
		},
		{
			name:    "Text Post with new id is stored properly",
			msg:     types.NewMsgCreatePost(testPost.Message, testPost.ParentID, false, "desmos", map[string]string{}, testPost.Creator),
			expPost: types.NewTextPost(types.PostID(1), testPost.ParentID, testPost.Message, testPost.AllowsComments, "desmos", map[string]string{}, 0, testPost.Creator),
		},
		{
			name:     "Storing a valid post with missing parent id returns expError",
			msg:      types.NewMsgCreatePost(testPost.Message, types.PostID(50), false, "desmos", map[string]string{}, testPost.Creator),
			expError: "Parent post with id 50 not found",
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns expError",
			storedPosts: types.Posts{
				types.NewTextPost(types.PostID(50), types.PostID(50), "Parent post", false, "desmos", map[string]string{}, 0, testPost.Creator),
			},
			msg:      types.NewMsgCreatePost(testPost.Message, types.PostID(50), false, "desmos", map[string]string{}, testPost.Creator),
			expError: "Post with id 50 does not allow comments",
		},
		{
			name: "Media Post with new id is stored properly",
			msg: types.NewMsgCreateMediaPost(testPost.Message, testPost.ParentID, false, "desmos", map[string]string{}, testPost.Creator, types.PostMedias{types.PostMedia{
				Provider: "provider",
				URI:      "uri",
				MimeType: "text/plain",
			}}),
			expPost: types.NewMediaPost(
				types.NewTextPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, 0, testPost.Creator),
				types.PostMedias{
					types.PostMedia{
						Provider: "provider",
						URI:      "uri",
						MimeType: "text/plain",
					},
				},
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			for _, p := range test.storedPosts {
				store.Set([]byte(types.PostStorePrefix+p.GetID().String()), k.Cdc.MustMarshalBinaryBare(p))
			}

			if test.lastPostID.Valid() {
				store.Set([]byte(types.LastPostIDStoreKey), k.Cdc.MustMarshalBinaryBare(&test.lastPostID))
			}

			handler := keeper.NewHandler(k)
			res := handler(ctx, test.msg)

			// Valid response
			if len(test.expError) == 0 {
				assert.True(t, res.IsOK())

				// Check the post
				var stored types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+test.expPost.GetID().String())), &stored)
				assert.True(t, stored.Equals(test.expPost))

				// Check the data
				assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.expPost.GetID()), res.Data)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypePostCreated,
					sdk.NewAttribute(types.AttributeKeyPostID, test.expPost.GetID().String()),
					sdk.NewAttribute(types.AttributeKeyPostParentID, test.expPost.GetParentID().String()),
					sdk.NewAttribute(types.AttributeKeyCreationTime, test.expPost.CreationTime().String()),
					sdk.NewAttribute(types.AttributeKeyPostOwner, test.expPost.Owner().String()),
				)
				assert.Len(t, ctx.EventManager().Events(), 1)
				assert.Contains(t, ctx.EventManager().Events(), creationEvent)
			}

			// Invalid response
			if len(test.expError) != 0 {
				assert.False(t, res.IsOK())
				assert.Contains(t, res.Log, test.expError)
			}
		})
	}

}

func Test_handleMsgEditPost(t *testing.T) {
	editor, _ := sdk.AccAddressFromBech32("cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63")
	testData := []struct {
		name        string
		storedPost  *types.TextPost
		msg         types.MsgEditPost
		blockHeight int64
		expError    string
		expPost     types.TextPost
	}{
		{
			name:       "Text Post not found",
			storedPost: nil,
			msg:        types.NewMsgEditPost(types.PostID(0), "Edited message", testPostOwner),
			expError:   "Post with id 0 not found",
		},
		{
			name:       "Invalid editor",
			storedPost: &testPost,
			msg:        types.NewMsgEditPost(testPost.PostID, "Edited message", editor),
			expError:   "Incorrect owner",
		},
		{
			name:        "Edit date before creation date",
			storedPost:  &testPost,
			blockHeight: testPost.Created.Int64() - 1,
			msg:         types.NewMsgEditPost(testPost.PostID, "Edited message", testPost.Creator),
			expError:    "Edit date cannot be before creation date",
		},
		{
			name:        "Valid request is handled properly",
			storedPost:  &testPost,
			blockHeight: testPost.Created.Int64() + 1,
			msg:         types.NewMsgEditPost(testPost.PostID, "Edited message", testPost.Creator),
			expPost: types.TextPost{
				PostID:         testPost.PostID,
				ParentID:       testPost.ParentID,
				Message:        "Edited message",
				Created:        testPost.Created,
				LastEdited:     testPost.Created.Add(sdk.NewInt(1)),
				AllowsComments: testPost.AllowsComments,
				Subspace:       testPost.Subspace,
				OptionalData:   testPost.OptionalData,
				Creator:        testPost.Creator,
			},
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

			// Valid response
			if len(test.expError) == 0 {
				assert.True(t, res.IsOK())
				assert.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypePostEdited,
					sdk.NewAttribute(types.AttributeKeyPostID, testPost.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostEditTime, strconv.FormatInt(ctx.BlockHeight(), 10)),
				))

				var stored types.TextPost
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+testPost.PostID.String())), &stored)
				assert.True(t, test.expPost.Equals(stored))
			}

			// Invalid response
			if len(test.expError) != 0 {
				assert.False(t, res.IsOK())
				assert.Contains(t, res.Log, test.expError)
				assert.Empty(t, res.Events)
			}
		})
	}
}

func Test_handleMsgAddPostReaction(t *testing.T) {

	user, _ := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	tests := []struct {
		name         string
		existingPost *types.TextPost
		msg          types.MsgAddPostReaction
		error        string
	}{
		{
			name:  "Text Post not found",
			msg:   types.NewMsgAddPostReaction(types.PostID(0), "like", user),
			error: "Post with id 0 not found",
		},
		{
			name:         "Valid message works properly",
			existingPost: &testPost,
			msg:          types.NewMsgAddPostReaction(testPost.PostID, "like", user),
			error:        "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			if test.existingPost != nil {
				ctx = ctx.WithBlockHeight(test.existingPost.Created.Int64() + 1)
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

			// Valid response
			if len(test.error) == 0 {
				assert.True(t, res.IsOK())
				assert.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypeReactionAdded,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyReactionOwner, test.msg.User.String()),
					sdk.NewAttribute(types.AttributeKeyReactionValue, test.msg.Value),
				))

				var storedPost types.TextPost
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+testPost.PostID.String())), &storedPost)
				assert.True(t, test.existingPost.Equals(storedPost))

				var storedReactions types.Reactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostReactionsStorePrefix+storedPost.PostID.String())), &storedReactions)
				assert.Contains(t, storedReactions, types.NewReaction(test.msg.Value, ctx.BlockHeight(), test.msg.User))
			}

			// Invalid response
			if len(test.error) != 0 {
				assert.Contains(t, res.Log, test.error)
				assert.Empty(t, res.Events)
			}
		})
	}
}

func Test_handleMsgRemovePostReaction(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	reaction := types.NewReaction("like", testPost.Created.Int64()+1, user)
	tests := []struct {
		name             string
		existingPost     *types.TextPost
		existingReaction *types.Reaction
		msg              types.MsgRemovePostReaction
		error            string
	}{
		{
			name:  "Text Post not found",
			msg:   types.NewMsgRemovePostReaction(types.PostID(0), user, "like"),
			error: "Post with id 0 not found",
		},
		{
			name:         "Reaction not found",
			existingPost: &testPost,
			msg:          types.NewMsgRemovePostReaction(testPost.PostID, user, "like"),
			error:        fmt.Sprintf("Cannot remove the reaction with value like from user %s as it does not exist", user),
		},
		{
			name:             "Valid message works properly",
			existingPost:     &testPost,
			existingReaction: &reaction,
			msg:              types.NewMsgRemovePostReaction(testPost.PostID, user, reaction.Value),
			error:            "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			if test.existingPost != nil {
				ctx = ctx.WithBlockHeight(test.existingPost.Created.Int64() + 1)
			}

			store := ctx.KVStore(k.StoreKey)
			if test.existingPost != nil {
				store.Set(
					[]byte(types.PostStorePrefix+test.existingPost.PostID.String()),
					k.Cdc.MustMarshalBinaryBare(&test.existingPost),
				)
			}

			if test.existingReaction != nil {
				store.Set(
					[]byte(types.PostReactionsStorePrefix+test.existingPost.PostID.String()),
					k.Cdc.MustMarshalBinaryBare(&types.Reactions{*test.existingReaction}),
				)
			}

			handler := keeper.NewHandler(k)
			res := handler(ctx, test.msg)

			// Valid response
			if len(test.error) == 0 {
				assert.True(t, res.IsOK())
				assert.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypePostReactionRemoved,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyReactionOwner, test.msg.User.String()),
					sdk.NewAttribute(types.AttributeKeyReactionValue, test.msg.Reaction),
				))

				var storedPost types.TextPost
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+testPost.PostID.String())), &storedPost)
				assert.True(t, test.existingPost.Equals(storedPost))

				var storedReactions types.Reactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostReactionsStorePrefix+storedPost.PostID.String())), &storedReactions)
				assert.NotContains(t, storedReactions, test.existingReaction)
			}

			// Invalid response
			if len(test.error) != 0 {
				assert.Contains(t, res.Log, test.error)
				assert.Empty(t, res.Events)
			}
		})
	}
}
