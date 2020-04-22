package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"
)

// ---------------------------
// --- handleMsgCreatePost
// ---------------------------

func Test_handleMsgCreatePost(t *testing.T) {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")

	computedID := types.ComputeID(testPost.Created, testPost.Creator, testPost.Subspace)

	tests := []struct {
		name        string
		storedPosts types.Posts
		msg         types.MsgCreatePost
		expPost     types.Post
		expError    error
	}{
		{
			name: "Trying to store post with same id returns expError",
			storedPosts: types.Posts{
				types.NewPost(
					computedID,
					testPost.ParentID,
					testPost.Message,
					testPost.AllowsComments,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
				),
			},
			msg: types.NewMsgCreatePost(
				testPost.Message,
				testPost.ParentID,
				testPost.AllowsComments,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
				testPost.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"the provided post conflicts with the one having id 46e61c7ac7016e8dd1d7270b114ecb7d1cf45cc85caa0308de540ccc15676fc7"),
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
				computedID,
				testPost.ParentID,
				testPost.Message,
				testPost.AllowsComments,
				testPost.Subspace,
				testPost.OptionalData,
				testPost.Created,
				testPost.Creator,
			).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
		},
		{
			name: "Storing a valid post with missing parent id returns expError",
			msg: types.NewMsgCreatePost(
				testPost.Message,
				id2,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Creator,
				testPost.Created,
				testPost.Medias,
				testPost.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "parent post with id f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd not found"),
		},
		{
			name: "Storing a valid post with parent stored but not accepting comments returns expError",
			storedPosts: types.Posts{
				types.NewPost(
					id,
					id2,
					"Parent post",
					false,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					map[string]string{},
					testPost.Created,
					testPost.Creator,
				),
			},
			msg: types.NewMsgCreatePost(
				testPost.Message,
				id,
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				map[string]string{},
				testPost.Creator,
				testPost.Created.AddDate(0, 0, 1),
				testPost.Medias,
				testPost.PollData,
			),
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af does not allow comments"),
		},
		{
			name: "Post with exact same data is not posted again",
			storedPosts: []types.Post{
				types.NewPost(
					computedID,
					testPost.ParentID,
					testPost.Message,
					testPost.AllowsComments,
					testPost.Subspace,
					testPost.OptionalData,
					testPost.Created,
					testPost.Creator,
				).WithMedias(testPost.Medias).WithPollData(*testPost.PollData),
			},
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
			expError: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"the provided post conflicts with the one having id 46e61c7ac7016e8dd1d7270b114ecb7d1cf45cc85caa0308de540ccc15676fc7"),
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

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				// Check the post
				var stored types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.expPost.PostID)), &stored)
				require.True(t, stored.Equals(test.expPost), "Expected: %s, actual: %s", test.expPost, stored)

				// Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.expPost.PostID), res.Data)

				// Check the events
				creationEvent := sdk.NewEvent(
					types.EventTypePostCreated,
					sdk.NewAttribute(types.AttributeKeyPostID, test.expPost.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostParentID, test.expPost.ParentID.String()),
					sdk.NewAttribute(types.AttributeKeyCreationTime, test.expPost.Created.String()),
					sdk.NewAttribute(types.AttributeKeyPostOwner, test.expPost.Creator.String()),
				)
				require.Len(t, ctx.EventManager().Events(), 1)
				require.Contains(t, ctx.EventManager().Events(), creationEvent)
			}

			// Invalid response
			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expError.Error(), err.Error())
			}
		})
	}

}

func Test_handleMsgEditPost(t *testing.T) {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	editor, err := sdk.AccAddressFromBech32("cosmos1z427v6xdc8jgn5yznfzhwuvetpzzcnusut3z63")
	require.NoError(t, err)

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
			msg:        types.NewMsgEditPost(id, "Edited message", testPostOwner, testPost.Created),
			expError:   sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af not found"),
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
				store.Set(types.PostStoreKey(test.storedPost.PostID), k.Cdc.MustMarshalBinaryBare(&test.storedPost))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				require.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypePostEdited,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostEditTime, test.msg.EditDate.String()),
				))

				var stored types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(test.storedPost.PostID)), &stored)
				require.True(t, test.expPost.Equals(stored))
			}

			// Invalid response
			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expError.Error(), err.Error())
			}
		})
	}
}

func Test_handleMsgAddPostReaction(t *testing.T) {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	require.NoError(t, err)

	var testRegisteredReaction2 = types.NewReaction(types.ModuleAddress, ":slightly_smiling_face:", "🙂",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")

	tests := []struct {
		name               string
		existingPost       *types.Post
		msg                types.MsgAddPostReaction
		registeredReaction *types.Reaction
		error              error
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgAddPostReaction(id, ":smile:", user),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af not found"),
		},
		{
			name:               "Valid message works properly (shortcode)",
			existingPost:       &testPost,
			msg:                types.NewMsgAddPostReaction(testPost.PostID, ":smile:", user),
			registeredReaction: &testRegisteredReaction,
			error:              nil,
		},
		{
			name:               "Valid message works properly (emoji)",
			existingPost:       &testPost,
			msg:                types.NewMsgAddPostReaction(testPost.PostID, "🙂", user),
			registeredReaction: &testRegisteredReaction2,
			error:              nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if test.existingPost != nil {
				store.Set(types.PostStoreKey(test.existingPost.PostID), k.Cdc.MustMarshalBinaryBare(&test.existingPost))
				if test.registeredReaction != nil && test.registeredReaction.Value != "🙂" {
					k.RegisterReaction(ctx, *test.registeredReaction)
				}
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				require.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypePostReactionAdded,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostReactionOwner, test.msg.User.String()),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, test.registeredReaction.ShortCode),
				))

				var storedPost types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(testPost.PostID)), &storedPost)
				require.True(t, test.existingPost.Equals(storedPost))

				if test.msg.Value == "🙂" {
					var regReaction types.Reaction
					bytes := store.Get(types.ReactionsStoreKey(testRegisteredReaction2.ShortCode, testRegisteredReaction2.Subspace))
					k.Cdc.MustUnmarshalBinaryBare(bytes, &regReaction)
					require.Equal(t, testRegisteredReaction2, regReaction)
				}

				var storedReactions types.PostReactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				require.Contains(t, storedReactions, types.NewPostReaction(test.registeredReaction.ShortCode, test.msg.User))
			}

			// Invalid response
			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.error.Error(), err.Error())
			}
		})
	}
}

func Test_handleMsgRemovePostReaction(t *testing.T) {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	require.NoError(t, err)

	reaction := types.NewPostReaction("reaction", user)
	tests := []struct {
		name             string
		existingPost     *types.Post
		existingReaction *types.PostReaction
		msg              types.MsgRemovePostReaction
		error            error
	}{
		{
			name:  "Post not found",
			msg:   types.NewMsgRemovePostReaction(id, user, "reaction"),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af not found"),
		},
		{
			name:         "PostReaction not found",
			existingPost: &testPost,
			msg:          types.NewMsgRemovePostReaction(testPost.PostID, user, "reaction"),
			error:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("cannot remove the reaction with value reaction from user %s as it does not exist", user)),
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
					k.Cdc.MustMarshalBinaryBare(&types.PostReactions{*test.existingReaction}),
				)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				require.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypePostReactionRemoved,
					sdk.NewAttribute(types.AttributeKeyPostID, test.msg.PostID.String()),
					sdk.NewAttribute(types.AttributeKeyPostReactionOwner, test.msg.User.String()),
					sdk.NewAttribute(types.AttributeKeyPostReactionValue, test.msg.Reaction),
				))

				var storedPost types.Post
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostStoreKey(testPost.PostID)), &storedPost)
				require.True(t, test.existingPost.Equals(storedPost))

				var storedReactions types.PostReactions
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.PostReactionsStoreKey(storedPost.PostID)), &storedReactions)
				require.NotContains(t, storedReactions, test.existingReaction)
			}

			// Invalid response
			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.error.Error(), err.Error())
			}
		})
	}
}

func Test_handleMsgAnswerPollPost(t *testing.T) {
	id := types.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := types.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}
	userPollAnswers := types.NewUserAnswer(answers, testPostOwner)

	tests := []struct {
		name          string
		msg           types.MsgAnswerPoll
		storedPost    types.Post
		storedAnswers *types.UserAnswer
		expErr        error
	}{
		{
			name: "Post not found",
			msg:  types.NewMsgAnswerPoll(id2, []types.AnswerID{1, 2}, testPostOwner),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				testPostCreationDate,
				testPostOwner,
			).WithPollData(types.NewPollData(
				"poll?",
				testPostEndPollDate,
				types.PollAnswers{answer, answer2},
				true,
				true,
				true,
			)),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "post with id f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd doesn't exist"),
		},
		{
			name: "No poll associated with post",
			msg:  types.NewMsgAnswerPoll(id2, []types.AnswerID{1, 2}, testPostOwner),
			storedPost: types.NewPost(
				id2,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				testPostCreationDate,
				testPostOwner,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no poll associated with ID: f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
		},
		{
			name: "Answer after poll closure",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1}, testPostOwner),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				testPostCreationDate,
				testPostOwner,
			).WithPollData(types.NewPollData(
				"poll?",
				testPostEndPollDateExpired,
				types.PollAnswers{answer},
				true,
				false,
				true,
			)),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("the poll associated with ID %s was closed at %s", id, testPostEndPollDateExpired)),
		},
		{
			name: "Poll doesn't allow multiple answers",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, testPostOwner),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				testPostCreationDate,
				testPostOwner,
			).WithPollData(types.NewPollData(
				"poll?",
				testPostEndPollDate,
				types.PollAnswers{answer},
				true,
				false,
				true,
			),
			),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "the poll associated with ID 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af doesn't allow multiple answers"),
		},
		{
			name: "Creator provide too many answers",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2, 3}, testPostOwner),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				testPostCreationDate,
				testPostOwner,
			).WithPollData(types.NewPollData(
				"poll?",
				testPostEndPollDate,
				types.PollAnswers{answer, answer2},
				true,
				true,
				true,
			),
			),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "user's answers are more than the available ones in Poll"),
		},
		{
			name: "Creator provide answers that are not the ones provided by the poll",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 3}, testPostOwner),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				testPostCreationDate,
				testPostOwner,
			).WithPollData(types.NewPollData(
				"poll?",
				testPostEndPollDate,
				types.PollAnswers{answer, answer2},
				true,
				true,
				true,
			)),
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "answer with ID 3 isn't one of the poll's provided answers"),
		},
		{
			name: "Poll doesn't allow answers' edits",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, testPostOwner),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				testPostCreationDate,
				testPostOwner,
			).WithPollData(types.NewPollData(
				"poll?",
				testPostEndPollDate,
				types.PollAnswers{answer, answer2},
				true,
				true,
				false,
			)),
			storedAnswers: &userPollAnswers,
			expErr: sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest, "post with ID 19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af doesn't allow answers' edits"),
		},
		{
			name: "Answered correctly to post's poll",
			msg:  types.NewMsgAnswerPoll(id, []types.AnswerID{1, 2}, testPostOwner),
			storedPost: types.NewPost(
				id,
				"",
				"Post message",
				false,
				"desmos",
				map[string]string{},
				testPostCreationDate,
				testPostOwner,
			).WithPollData(types.NewPollData(
				"poll?",
				testPostEndPollDate,
				types.PollAnswers{answer, answer2},
				true,
				true,
				true,
			)),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)
			store.Set(types.PostStoreKey(test.storedPost.PostID), k.Cdc.MustMarshalBinaryBare(&test.storedPost))

			if test.storedAnswers != nil {
				k.SavePollAnswers(ctx, test.storedPost.PostID, *test.storedAnswers)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Invalid response
			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}

			// Valid response
			if res != nil {
				{
					// Check the data
					require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed("Answered to poll correctly"), res.Data)

					// Check the events
					answerEvent := sdk.NewEvent(
						types.EventTypeAnsweredPoll,
						sdk.NewAttribute(types.AttributeKeyPostID, test.storedPost.PostID.String()),
						sdk.NewAttribute(types.AttributeKeyPollAnswerer, testPostOwner.String()),
					)

					require.Len(t, ctx.EventManager().Events(), 1)
					require.Contains(t, ctx.EventManager().Events(), answerEvent)
				}
			}
		})
	}

}

func Test_handleMsgRegisterReaction(t *testing.T) {

	user, err := sdk.AccAddressFromBech32("cosmos1q4hx350dh0843wr3csctxr87at3zcvd9qehqvg")
	require.NoError(t, err)

	shortCode := ":smile:"
	const (
		subspace = "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	)
	value := "https://smile.jpg"
	reaction := types.NewReaction(user, shortCode, value, subspace)

	tests := []struct {
		name             string
		existingReaction *types.Reaction
		msg              types.MsgRegisterReaction
		error            error
	}{
		{
			name:             "Reaction registered without error",
			existingReaction: nil,
			msg:              types.NewMsgRegisterReaction(user, shortCode, value, subspace),
			error:            nil,
		},
		{
			name:             "Already registered reaction returns error",
			existingReaction: &reaction,
			msg:              types.NewMsgRegisterReaction(user, shortCode, value, subspace),
			error: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf(
				"reaction with shortcode %s and subspace %s has already been registered", shortCode, subspace)),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if test.existingReaction != nil {
				store.Set(types.ReactionsStoreKey(test.existingReaction.ShortCode, test.existingReaction.Subspace),
					k.Cdc.MustMarshalBinaryBare(&test.existingReaction))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			// Valid response
			if res != nil {
				require.Contains(t, res.Events, sdk.NewEvent(
					types.EventTypeRegisterReaction,
					sdk.NewAttribute(types.AttributeKeyReactionCreator, test.msg.Creator.String()),
					sdk.NewAttribute(types.AttributeKeyReactionShortCode, test.msg.ShortCode),
					sdk.NewAttribute(types.AttributeKeyReactionSubSpace, test.msg.Subspace),
				))

				var storedReaction types.Reaction
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.ReactionsStoreKey(shortCode,
					subspace)), &storedReaction)
				require.True(t, reaction.Equals(storedReaction))
			}

			// Invalid response
			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.error.Error(), err.Error())
			}
		})
	}
}
