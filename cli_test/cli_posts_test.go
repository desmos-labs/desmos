// +build cli_test

// nolint
package clitest

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/posts/types"
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
	computedID := types.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.PollData)
	require.Nil(t, post.Attachments)

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
	computedID := types.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.False(t, post.AllowsComments)
	require.Nil(t, post.PollData)
	require.Nil(t, post.Attachments)

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

func TestDesmosCLIPostsCreateWithAttachmentsAndEmptyMessage(t *testing.T) {
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

	tag, err := sdk.AccAddressFromBech32("desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru")
	require.NoError(t, err)
	tag2, err2 := sdk.AccAddressFromBech32("desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr")
	require.NoError(t, err2)

	// Create a post
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y",
		"--attachment https://example.com/media1,text/plain,desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru",
		"--attachment https://example.com/media2,application/json,desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	computedID := types.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.PollData)
	require.Len(t, post.Attachments, 2)
	require.Equal(t, post.Attachments, types.NewAttachments(
		types.NewAttachment("https://example.com/media1", "text/plain", []sdk.AccAddress{tag}),
		types.NewAttachment("https://example.com/media2", "application/json", []sdk.AccAddress{tag2})))

	// Test --dry-run
	success, _, _ = f.TxPostsCreate(subspace, message, fooAddr, "--dry-run",
		"--attachment https://example.com/media1,text/plain,desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru",
		"--attachment https://example.com/media2,application/json,desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--generate-only",
		"--attachment https://example.com/media1,text/plain,desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru",
		"--attachment https://example.com/media2,application/json,desmos1ulmv2dyc8zjmhk9zlsq4ajpudwc8zjfm82aysr")
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

func TestDesmosCLIPostsCreateWithAttachmentsAndNonEmptyMessage(t *testing.T) {
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
		"--attachment https://example.com/media1,text/plain",
		"--attachment https://example.com/media2,application/json",
		"--attachment https://example.com/media3,text/plain",
	)
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	computedID := types.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.PollData)
	require.Len(t, post.Attachments, 3)
	require.Equal(t, types.NewAttachments(
		types.NewAttachment("https://example.com/media1", "text/plain", nil),
		types.NewAttachment("https://example.com/media2", "application/json", nil),
		types.NewAttachment("https://example.com/media3", "text/plain", nil),
	), post.Attachments)

	// Test --dry-run
	success, _, _ = f.TxPostsCreate(subspace, message, fooAddr, "--dry-run",
		"--attachment https://example.com/media1,text/plain",
		"--attachment https://example.com/media2,application/json",
		"--attachment https://example.com/media3,text/plain",
	)
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsCreate(subspace, message, fooAddr, "--generate-only",
		"--attachment https://example.com/media1,text/plain",
		"--attachment https://example.com/media2,application/json",
		"--attachment https://example.com/media3,text/plain",
	)
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
	computedID := types.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.PollData)
	require.Len(t, post.Attachments, 0)

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
	computedID := types.ComputeID(post.Created, post.Creator, post.Subspace)
	require.Equal(t, computedID, post.PostID)
	require.Nil(t, post.Attachments)
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
	require.Equal(t, types.NewPollAnswer(0, "Beagle"), pollData.ProvidedAnswers[0])
	require.Equal(t, types.NewPollAnswer(1, "Pug"), pollData.ProvidedAnswers[1])
	require.Equal(t, types.NewPollAnswer(2, "Shiba"), pollData.ProvidedAnswers[2])

	// Test --dry-run
	success, _, _ = f.TxPostsCreate(subspace, message, fooAddr, "--dry-run",
		"--poll-details question=Dog?,multiple-answers=false,allows-answer-edits=true,end-date=2100-01-01T15:00:00.000Z",
		"--poll-answer Beagle",
		"--poll-answer Pug",
		"--poll-answer Shiba")
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
	success, _, sterr = f.TxPostsAnswerPoll(post.PostID, []types.AnswerID{types.AnswerID(1)}, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Check that answers have been inserted
	postQueryResponse := f.QueryPost(post.PostID.String())
	require.NotEmpty(t, postQueryResponse.PollAnswers)
	require.Equal(t, types.NewUserAnswer([]types.AnswerID{types.AnswerID(1)}, fooAddr), postQueryResponse.PollAnswers[0])

	// Test --dry-run
	success, _, _ = f.TxPostsAnswerPoll(post.PostID, []types.AnswerID{types.AnswerID(1)}, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsAnswerPoll(post.PostID, []types.AnswerID{types.AnswerID(1)}, fooAddr, "--generate-only")
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
	reactions := types.Reactions{
		types.Reaction{":earth:", "http://earth.jpg", subspace, fooAddr},
		types.Reaction{":blush:", "https://gph.is/2p19Zai", subspace, fooAddr},
		types.Reaction{":thumbsdown:", "https://gph.is/2phybnt", subspace, fooAddr},
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
	require.Equal(t, types.NewPostReaction(":+1:", "👍", fooAddr), storedPost.Reactions[0])

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

	// Test --dry-run
	// This is executed before the actual delete since the dry-run performs the proper checks and would fail
	// telling there is no such added reaction otherwise
	success, _, _ = f.TxPostsRemoveReaction(post.PostID.String(), ":+1:", fooAddr, "--dry-run")
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
	require.Len(t, storedPost.Reactions, 1)

	// Remove a reaction
	success, _, sterr = f.TxPostsRemoveReaction(post.PostID.String(), ":+1:", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the reaction has been removed
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
	require.Equal(t, registeredReactions, types.Reactions{types.NewReaction(fooAddr, shortCode, value, subspace)})

	// Test --dry-run
	success, _, _ = f.TxPostsRegisterReaction(":second:", value, subspace, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxPostsRegisterReaction(":third:", value, subspace, fooAddr, "--generate-only")
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
	require.Equal(t, registeredReactions, types.Reactions{types.NewReaction(fooAddr, shortCode, value, subspace)})

	// Test --dry-run
	success, _, _ = f.TxPostsRegisterReaction(":second:", value, subspace, fooAddr, "--dry-run")
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
