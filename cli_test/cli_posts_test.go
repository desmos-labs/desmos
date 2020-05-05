// +build cli_test

// nolint
package clitest

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts"
	"github.com/stretchr/testify/require"
)

func TestDesmosCLIPostsCreateNoMediasNoPollData(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	message := "message#test"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a post
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	computedID := posts.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.PollData)
	require.Nil(t, post.Medias)

	// Test --dry-run
	success, _, _ = f.TxPostsCreate(subspace, message, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPosts = f.QueryPosts()
	require.Len(t, storedPosts, 1)

	f.Cleanup()
}

func TestDesmosCLIPostsCreateAllowsCommentFalse(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	message := "message#test"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a post
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y",
		"--allows-comments=false")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	computedID := posts.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.False(t, post.AllowsComments)
	require.Nil(t, post.PollData)
	require.Nil(t, post.Medias)

	// Test --dry-run
	success, _, _ = f.TxPostsCreate(subspace, message, fooAddr, "--dry-run",
		"--allows-comments=false")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--generate-only=true",
		"--allows-comments=false")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPosts = f.QueryPosts()
	require.Len(t, storedPosts, 1)

	f.Cleanup()
}

func TestDesmosCLIPostsCreateWithMediasAndEmptyMessage(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	message := ""
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a post
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y",
		"--media https://example.com/media1,text/plain",
		"--media https://example.com/media2,application/json")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	computedID := posts.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.PollData)
	require.Len(t, post.Medias, 2)
	require.Equal(t, post.Medias, posts.NewPostMedias(
		posts.NewPostMedia("https://example.com/media1", "text/plain"),
		posts.NewPostMedia("https://example.com/media2", "application/json")))

	// Test --dry-run
	success, _, _ = f.TxPostsCreate(subspace, message, fooAddr, "--dry-run",
		"--media https://second.example.com/media1,text/plain",
		"--media https://second.example.com/media2,application/json")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--generate-only",
		"--media https://third.example.com/media1,text/plain",
		"--media https://third.example.com/media2,application/json")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPosts = f.QueryPosts()
	require.Len(t, storedPosts, 1)

	f.Cleanup()
}

func TestDesmosCLIPostsCreateWithMediasAndNonEmptyMessage(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	message := "message"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a post
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y",
		"--media https://example.com/media1,text/plain",
		"--media https://example.com/media2,application/json")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	computedID := posts.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.PollData)
	require.Len(t, post.Medias, 2)
	require.Equal(t, post.Medias, posts.NewPostMedias(
		posts.NewPostMedia("https://example.com/media1", "text/plain"),
		posts.NewPostMedia("https://example.com/media2", "application/json")))

	// Test --dry-run
	success, _, _ = f.TxPostsCreate(subspace, message, fooAddr, "--dry-run",
		"--media https://second.example.com/media1,text/plain",
		"--media https://second.example.com/media2,application/json")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--generate-only",
		"--media https://third.example.com/media1,text/plain",
		"--media https://third.example.com/media2,application/json")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPosts = f.QueryPosts()
	require.Len(t, storedPosts, 1)

	f.Cleanup()
}

func TestDesmosCLIPostsCreateWithNoMediasAndNonEmptyMessage(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	message := "message"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a post
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	computedID := posts.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.PollData)
	require.Len(t, post.Medias, 0)

	// Test --dry-run
	success, _, _ = f.TxPostsCreate(subspace, message, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--generate-only")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPosts = f.QueryPosts()
	require.Len(t, storedPosts, 1)

	f.Cleanup()
}

func TestDesmosCLIPostsCreateWithPoll(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	message := "message"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a post
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y",
		"--poll-details question=Dog?,multiple-answers=false,allows-answer-edits=true,end-date=2100-01-01T15:00:00.000Z",
		"--poll-answer Beagle",
		"--poll-answer Pug",
		"--poll-answer Shiba")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	computedID := posts.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.Medias)
	require.NotNil(t, post.PollData)

	// Check poll data
	pollData := post.PollData
	require.Equal(t, "Dog?", pollData.Question)
	require.False(t, pollData.AllowsMultipleAnswers)
	require.True(t, pollData.AllowsAnswerEdits)
	location, _ := time.LoadLocation("UTC")
	date := time.Date(2100, 1, 1, 15, 0, 0, 0, location)
	require.Equal(t, pollData.EndDate, date)

	// Check poll answers
	require.Len(t, pollData.ProvidedAnswers, 3)
	require.Equal(t, posts.NewPollAnswer(0, "Beagle"), pollData.ProvidedAnswers[0])
	require.Equal(t, posts.NewPollAnswer(1, "Pug"), pollData.ProvidedAnswers[1])
	require.Equal(t, posts.NewPollAnswer(2, "Shiba"), pollData.ProvidedAnswers[2])

	// Test --dry-run
	success, _, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--dry-run",
		"--poll-details question=Dog?,multiple-answers=false,allows-answer-edits=true,end-date=2100-01-01T15:00:00.000Z",
		"--poll-answer Beagle",
		"--poll-answer Pug",
		"--poll-answer Shiba")
	require.Empty(t, sterr)
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--generate-only",
		"--poll-details question=Dog?,multiple-answers=false,allows-answer-edits=true,end-date=2100-01-01T15:00:00.000Z",
		"--poll-answer Beagle",
		"--poll-answer Pug",
		"--poll-answer Shiba")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPosts = f.QueryPosts()
	require.Len(t, storedPosts, 1)

	f.Cleanup()
}

func TestDesmosCLIPostsAnswerPoll(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	message := "message"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a poll with single answer
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y",
		"--poll-details question=Dog?,multiple-answers=false,allows-answer-edits=true,end-date=2100-01-01T15:00:00.000Z",
		"--poll-answer Beagle",
		"--poll-answer Pug",
		"--poll-answer Shiba")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]

	// Insert an answer
	success, _, sterr = f.TxPostsAnswerPoll(post.PostID, []posts.AnswerID{posts.AnswerID(1)}, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Check that answers have been inserted
	postQueryResponse := f.QueryPost(post.PostID.String())
	require.NotEmpty(t, postQueryResponse.PollAnswers)
	require.Equal(t, posts.NewUserAnswer([]posts.AnswerID{posts.AnswerID(1)}, fooAddr), postQueryResponse.PollAnswers[0])

	// Test --dry-run
	success, _, stderr := f.TxPostsAnswerPoll(post.PostID, []posts.AnswerID{posts.AnswerID(1)}, fooAddr, "--dry-run")
	require.Empty(t, sterr)
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsAnswerPoll(post.PostID, []posts.AnswerID{posts.AnswerID(1)}, fooAddr, "--generate-only")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	postQueryResponse = f.QueryPost(post.PostID.String())
	require.Len(t, postQueryResponse.PollAnswers, 1)

	f.Cleanup()
}

func TestDesmosCLIPostsEdit(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Create a post
	success, _, sterr := f.TxPostsCreate("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "message", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]

	// Edit the message
	success, _, sterr = f.TxPostsEdit(post.PostID.String(), "NewMessage", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the message is edited
	storedPost := f.QueryPost(post.PostID.String())
	require.Equal(t, post.PostID, storedPost.PostID)
	require.Equal(t, "NewMessage", storedPost.Message)

	// Test --dry-run
	success, _, _ = f.TxPostsEdit(post.PostID.String(), "OtherMessage", fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsEdit(post.PostID.String(), "OtherMessage", fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPosts = f.QueryPosts()
	require.Len(t, storedPosts, 1)

	f.Cleanup()
}

func TestDesmosCLIPostsReactions(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	reactions := posts.Reactions{
		posts.Reaction{":earth:", "http://earth.jpg", subspace, fooAddr},
		posts.Reaction{":blush:", "https://gph.is/2p19Zai", subspace, fooAddr},
		posts.Reaction{":thumbsdown:", "https://gph.is/2phybnt", subspace, fooAddr},
	}

	// Create a post
	success, _, sterr := f.TxPostsCreate("4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		"message", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]

	// Register reactions
	for _, reaction := range reactions {
		success, _, sterr = f.TxPostsRegisterReaction(reaction.ShortCode, reaction.Value, reaction.Subspace, reaction.Creator, "-y")
		require.True(t, success)
		require.Empty(t, sterr)
		tests.WaitForNextNBlocksTM(1, f.Port)
	}

	// __________________________________________________________________________________
	// add-reaction

	// Add a reaction
	success, _, sterr = f.TxPostsAddReaction(post.PostID.String(), "👍", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the reaction is added
	storedPost := f.QueryPost(post.PostID.String())
	require.Len(t, storedPost.Reactions, 1)
	require.Equal(t, posts.NewReactionQueryResponse("👍", ":+1:", fooAddr), storedPost.Reactions[0])

	// Test --dry-run
	success, _, _ = f.TxPostsAddReaction(post.PostID.String(), ":blush:", fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsAddReaction(post.PostID.String(), ":thumbsdown:", fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPost = f.QueryPost(post.PostID.String())
	require.Len(t, storedPost.Reactions, 1)

	// __________________________________________________________________________________
	// remove-reaction

	// Remove a reaction
	success, _, sterr = f.TxPostsRemoveReaction(post.PostID.String(), ":+1:", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the reaction has been removed
	storedPost = f.QueryPost(post.PostID.String())
	require.Empty(t, storedPost.Reactions)

	// Test --dry-run
	success, _, _ = f.TxPostsRemoveReaction(post.PostID.String(), ":blush:", fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr = f.TxPostsRemoveReaction(post.PostID.String(), ":thumbsdown:", fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg = unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedPost = f.QueryPost(post.PostID.String())
	require.Empty(t, storedPost.Reactions)

	f.Cleanup()
}

func TestDesmosCLIRegisterReaction(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	shortCode := ":like:"
	value := "https://like.jpg"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Register a reaction
	success, _, sterr := f.TxPostsRegisterReaction(shortCode, value, subspace, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the reaction is registered
	registeredReactions := f.QueryReactions()
	require.NotEmpty(t, registeredReactions)
	require.Equal(t, registeredReactions, posts.Reactions{posts.NewReaction(fooAddr, shortCode, value, subspace)})

	// Test --dry-run
	success, _, _ = f.TxPostsRegisterReaction(shortCode, value, subspace, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsRegisterReaction(shortCode, value, subspace, fooAddr, "--generate-only")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	registeredReactions = f.QueryReactions()
	require.Len(t, registeredReactions, 1)

	f.Cleanup()

}

func TestDesmosCLIRegisterReactionEmojiValue(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	shortCode := ":like:"
	value := "https://gph.is/2phybnt"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Register a reaction
	success, _, sterr := f.TxPostsRegisterReaction(shortCode, value, subspace, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the reaction is registered
	registeredReactions := f.QueryReactions()
	require.NotEmpty(t, registeredReactions)
	require.Equal(t, registeredReactions, posts.Reactions{posts.NewReaction(fooAddr, shortCode, value, subspace)})

	// Test --dry-run
	success, _, _ = f.TxPostsRegisterReaction(shortCode, value, subspace, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsRegisterReaction(shortCode, value, subspace, fooAddr, "--generate-only")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	registeredReactions = f.QueryReactions()
	require.Len(t, registeredReactions, 1)

	f.Cleanup()

}
