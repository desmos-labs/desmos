// +build cli_test

//nolint
package clitest

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestDesmosCLIReportPost(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	repType := "scam"
	repMess := "message"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a post
	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"
	message := "message#test"
	success, _, sterr := f.TxPostsCreate(subspace, message, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the post is saved
	storedPosts := f.QueryPosts()
	require.NotEmpty(t, storedPosts)
	post := storedPosts[0]
	require.Nil(t, post.PollData)
	require.Nil(t, post.Attachments)

	//Report a post
	success, _, sterr = f.TxReportPost(post.PostID.String(), repType, repMess, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the report is saved
	storedReports := f.QueryReports(post.PostID.String())
	require.NotEmpty(t, storedReports.Reports)
	report := storedReports.Reports[0]
	require.Equal(t, string(report.Type), repType)
	require.Equal(t, report.Message, repMess)
	require.Equal(t, report.User, fooAcc.Address)

	// Test --dry-run
	success, _, _ = f.TxReportPost(post.PostID.String(), repType, repMess, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxReportPost(post.PostID.String(), repType, repMess, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedReports = f.QueryReports(post.PostID.String())
	require.Len(t, storedReports.Reports, 1)

	f.Cleanup()
}
