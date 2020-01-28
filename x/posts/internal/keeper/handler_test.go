package keeper_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"
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
				types.NewPost(
					types.PostID(1),
					testPost.ParentID,
					testPost.Message,
					testPost.AllowsComments,
					"desmos",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
					testPost.PollData,
				),
			},
			lastPostID: types.PostID(0),
			msg: types.NewMsgCreatePost(
				testPost.Message,
				testPost.ParentID,
				testPost.AllowsComments,
				"desmos",
				map[string]string{},
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
				testPost.PollData,
			),
			expError: "Post with id 1 already exists",
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
				testPost.PollData,
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
				testPost.PollData,
			),
		},
		{
			name: "Storing a valid post with missing parent id returns expError",
			msg: types.NewMsgCreatePost(
				testPost.Message,
				types.PostID(50),
				false,
				"desmos",
				map[string]string{},
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
				testPost.PollData,
			),
			expError: "Parent post with id 50 not found",
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns expError",
			storedPosts: types.Posts{
				types.NewPost(
					types.PostID(50),
					types.PostID(50),
					"Parent post",
					false,
					"desmos",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
					testPost.Medias,
					testPost.PollData,
				),
			},
			msg: types.NewMsgCreatePost(
				testPost.Message,
				types.PostID(50),
				false,
				"desmos",
				map[string]string{},
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
				testPost.PollData,
			),
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

			// Valid response
			if len(test.expError) == 0 {
				assert.True(t, res.IsOK())

				// Check the post
				var stored types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+test.expPost.PostID.String())), &stored)
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
		name       string
		storedPost *types.Post
		msg        types.MsgEditPost
		expError   string
		expPost    types.Post
	}{
		{
			name:       "Post not found",
			storedPost: nil,
			msg:        types.NewMsgEditPost(types.PostID(0), "Edited message", testPostOwner, testPost.Created),
			expError:   "Post with id 0 not found",
		},
		{
			name:       "Invalid editor",
			storedPost: &testPost,
			msg:        types.NewMsgEditPost(testPost.PostID, "Edited message", editor, testPost.Created),
			expError:   "Incorrect owner",
		},
		{
			name:       "Edit date before creation date",
			storedPost: &testPost,
			msg:        types.NewMsgEditPost(testPost.PostID, "Edited message", testPost.Creator, testPost.Created.AddDate(0, 0, -1)),
			expError:   "Edit date cannot be before creation date",
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
				PollData:       testPost.PollData,
			},
		},
	}

	for _, test := range testData {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

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
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostEditTime, test.msg.EditDate.String()),
				))

				var stored types.Post
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
		existingPost *types.Post
		msg          types.MsgAddPostReaction
		error        string
	}{
		{
			name:  "Post not found",
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

				var storedPost types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostStorePrefix+testPost.PostID.String())), &storedPost)
				assert.True(t, test.existingPost.Equals(storedPost))

				var storedReactions types.Reactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.PostReactionsStorePrefix+storedPost.PostID.String())), &storedReactions)
				assert.Contains(t, storedReactions, types.NewReaction(test.msg.Value, test.msg.User))
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
	reaction := types.NewReaction("like", user)
	tests := []struct {
		name             string
		existingPost     *types.Post
		existingReaction *types.Reaction
		msg              types.MsgRemovePostReaction
		error            string
	}{
		{
			name:  "Post not found",
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

				var storedPost types.Post
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

func Test_handleMsgAnswerPollPosition(t *testing.T) {
	tests := []struct {
		name          string
		msg           types.MsgAnswerPollPost
		storedPost    types.Post
		storedAnswers []uint64
		expErr        string
	}{
		{
			name: "Post not found",
			msg:  types.NewMsgAnswerPollPost(types.PostID(1), []uint64{1, 2}, testPostOwner),
			storedPost: types.NewPost(types.PostID(2), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			),
			expErr: "Post with id 1 doesn't exist",
		},
		{
			name: "No poll associated with post",
			msg:  types.NewMsgAnswerPollPost(types.PostID(1), []uint64{1, 2}, testPostOwner),
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil, nil),
			expErr: "No poll associated with ID: 1",
		},
		{
			name: "Answer after poll closure",
			msg:  types.NewMsgAnswerPollPost(types.PostID(1), []uint64{1}, testPostOwner),
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDateExpired, types.PollAnswers{answer}, true, false, true),
			),
			expErr: fmt.Sprintf("The poll associated with ID %s was closed at %s", types.PostID(1), testPostEndPollDateExpired),
		},
		{
			name: "Poll doesn't allow multiple answers",
			msg:  types.NewMsgAnswerPollPost(types.PostID(1), []uint64{1, 2}, testPostOwner),
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer}, true, false, true),
			),
			expErr: "The poll associated with ID 1 doesn't allow multiple answers",
		},
		{
			name: "User provide too many answers",
			msg:  types.NewMsgAnswerPollPost(types.PostID(1), []uint64{1, 2, 3}, testPostOwner),
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			),
			expErr: "User's answers are more than the available ones in Poll",
		},
		{
			name: "Poll doesn't allow answers' edits",
			msg:  types.NewMsgAnswerPollPost(types.PostID(1), []uint64{1, 2}, testPostOwner),
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, false),
			),
			storedAnswers: []uint64{1, 2},
			expErr:        "Post with ID 1 doesn't allow answers' edits",
		},
		{
			name: "Answered correctly to post's poll",
			msg:  types.NewMsgAnswerPollPost(types.PostID(1), []uint64{1, 2}, testPostOwner),
			storedPost: types.NewPost(types.PostID(1), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			store.Set([]byte(types.PostStorePrefix+test.storedPost.PostID.String()), k.Cdc.MustMarshalBinaryBare(test.storedPost))

			if len(test.storedAnswers) != 0 {
				store.Set([]byte(types.PollAnswersStorePrefix+test.storedPost.PostID.String()+testPostOwner.String()), k.Cdc.MustMarshalBinaryBare(&test.storedAnswers))
			}

			handler := keeper.NewHandler(k)
			res := handler(ctx, test.msg)

			// Invalid response
			if len(test.expErr) != 0 {
				assert.False(t, res.IsOK())
				assert.Contains(t, res.Log, test.expErr)
			}

			// Valid response
			if len(test.expErr) == 0 {
				assert.True(t, res.IsOK())

				// Check the data
				assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed("Answered to poll correctly"), res.Data)

				// Check the events
				answerEvent := sdk.NewEvent(
					types.EventTypeAnsweredPoll,
					sdk.NewAttribute(types.AttributeKeyPostID, test.storedPost.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPollAnswerer, testPostOwner.String()),
				)

				assert.Len(t, ctx.EventManager().Events(), 1)
				assert.Contains(t, ctx.EventManager().Events(), answerEvent)
			}
		})
	}

}

func Test_handleMsgClosePoll(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos15lt0mflt6j9a9auj7yl3p20xec4xvljge0zhae")
	tests := []struct {
		name       string
		msg        types.MsgClosePollPost
		storedPost types.Post
		expErr     string
	}{
		{
			name:   "Post not found",
			msg:    types.NewMsgClosePollPost(types.PostID(1), "message", user),
			expErr: "Post with id 1 doesn't exists",
		},
		{
			name: "Try closing post from another user than creator",
			msg:  types.NewMsgClosePollPost(types.PostID(10), "message", user),
			storedPost: types.NewPost(types.PostID(10), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			),
			expErr: fmt.Sprintf("Only the poll creator can close it, %s", testPostOwner),
		},
		{
			name: "Try closing a post poll that doesn't exist",
			msg:  types.NewMsgClosePollPost(types.PostID(10), "message", testPostOwner),
			storedPost: types.NewPost(types.PostID(10), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil, nil),
			expErr: "No poll associated with this post ID: 10",
		},
		{
			name: "Try closing a post that is already closed",
			msg:  types.NewMsgClosePollPost(types.PostID(10), "message", testPostOwner),
			storedPost: types.NewPost(types.PostID(10), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDateExpired, types.PollAnswers{answer, answer2}, true, true, true),
			),
			expErr: "The poll associated with this post ID is already closed",
		},
		{
			name: "Closing post's poll correctly",
			msg:  types.NewMsgClosePollPost(types.PostID(10), "message", testPostOwner),
			storedPost: types.NewPost(types.PostID(10), types.PostID(0), "Post message", false, "desmos", map[string]string{}, testPostCreationDate,
				testPostOwner, nil,
				types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			store.Set([]byte(types.PostStorePrefix+test.storedPost.PostID.String()), k.Cdc.MustMarshalBinaryBare(test.storedPost))

			handler := keeper.NewHandler(k)
			res := handler(ctx, test.msg)

			// Invalid response
			if len(test.expErr) != 0 {
				assert.False(t, res.IsOK())
				assert.Contains(t, res.Log, test.expErr)
			}

			// Valid response
			if len(test.expErr) == 0 {
				assert.True(t, res.IsOK())

				// Check the data
				assert.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(fmt.Sprintf("Poll closed correctly, %s", test.msg.Message)), res.Data)

				// Check the events
				closeEvent := sdk.NewEvent(
					types.EventTypeClosePoll,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostOwner, test.msg.Creator.String()),
				)

				assert.Len(t, ctx.EventManager().Events(), 1)
				assert.Contains(t, ctx.EventManager().Events(), closeEvent)
			}
		})
	}

}
