//nolint
package clitest

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestDesmosCLIProfileCreate_noFlags(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave("mrBrown", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.DTag, "mrBrown")

	// Test --dry-run
	success, _, _ = f.TxProfileSave("mrBrown", fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileSave("mrBrown", fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedProfiles = f.QueryProfiles()
	require.Len(t, storedProfiles, 1)

	f.Cleanup()
}

func TestDesmosCLIProfileCreate_withFlags(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave("mrBrown", fooAddr, "-y",
		"--moniker Leonardo",
		"--bio biography",
		"--profile-pic https://profilePic.jpg",
		"--cover-pic https://profileCover.jpg")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	require.Equal(t, *storedProfiles[0].Moniker, "Leonardo")

	// Test --dry-run
	success, _, _ = f.TxProfileSave("mrBrown", fooAddr, "--dry-run",
		"--moniker Leonardo",
		"--bio biography",
		"--profile-pic https://profilePic.jpg",
		"--cover-pic https://profileCover.jpg")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileSave("mrBrown", fooAddr, "--generate-only=true",
		"--moniker Leonardo",
		"--bio biography",
		"--profile-pic https://profilePic.jpg",
		"--cover-pic https://profileCover.jpg")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedProfiles = f.QueryProfiles()
	require.Len(t, storedProfiles, 1)

	f.Cleanup()
}

func TestDesmosCLIProfileEdit_noFlags(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	fooAcc := f.QueryAccount(fooAddr)
	dTag := "mrBrown"
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave(dTag, fooAddr, "-y",
		"--moniker Leonardo",
		"--bio biography",
		"--profile-pic https://profilePic.jpg",
		"--cover-pic https://profileCover.jpg",
	)
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, *profile.Moniker, "Leonardo")

	// Edit the profile
	success, _, sterr = f.TxProfileSave("mrPink", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is edited
	editedProfiles := f.QueryProfiles()
	require.NotEmpty(t, editedProfiles)

	// Make sure the profile has been edited
	require.Nil(t, editedProfiles[0].Moniker)
	require.Nil(t, editedProfiles[0].Bio)
	require.Equal(t, "mrPink", editedProfiles[0].DTag)

	// Test --dry-run
	success, _, _ = f.TxProfileSave("mrPink", fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileSave("mrPink", fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedProfiles = f.QueryProfiles()
	require.Len(t, storedProfiles, 1)

	f.Cleanup()
}

func TestDesmosCLIProfileEdit_withFlags(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	fooAcc := f.QueryAccount(fooAddr)
	dTag := "mrBrown"
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave(dTag, fooAddr, "-y",
		"--moniker Leonardo",
		"--bio biography",
		"--profile-pic https://profilePic.jpg",
		"--cover-pic https://profileCover.jpg")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, *profile.Moniker, "Leonardo")

	// Edit the profile
	success, _, sterr = f.TxProfileSave(dTag, fooAddr, "-y",
		"--moniker Leo",
		"--bio HollywoodActor",
		"--profile-pic https://profilePic.jpg",
		"--cover-pic https://profileCover.jpg")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is edited
	editedProfiles := f.QueryProfiles()
	require.NotEmpty(t, editedProfiles)

	editedProfile := editedProfiles[0]
	require.Equal(t, *editedProfile.Moniker, "Leo")

	// Make sure the profile has been edited
	require.NotEqual(t, storedProfiles[0].Moniker, editedProfiles[0].Moniker)
	require.NotEqual(t, storedProfiles[0].Bio, editedProfiles[0].Bio)

	// Test --dry-run
	success, _, _ = f.TxProfileSave(dTag, fooAddr, "--dry-run",
		"--moniker Leo",
		"--bio HollywoodActor",
		"--profile-pic https://profilePic.jpg",
		"--cover-pic https://profileCover.jpg")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileSave("mrPink", fooAddr, "--generate-only=true",
		"--moniker Leo",
		"--bio HollywoodActor",
		"--profile-pic https://profilePic.jpg",
		"--cover-pic https://profileCover.jpg")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedProfiles = f.QueryProfiles()
	require.Len(t, storedProfiles, 1)

	f.Cleanup()
}

func TestDesmosCLIProfileDelete(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)

	// Later usage variables
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave("mrBrown", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.DTag, "mrBrown")

	// Test --dry-run
	// This is run before the actual no dry-run call due to the fact that even using --dry-run the checks
	// are performed anyway, and this would fail if the profile didn't exist
	success, _, _ = f.TxProfileDelete(fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileDelete(fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	// Check state didn't change
	storedProfiles = f.QueryProfiles()
	require.Len(t, storedProfiles, 1)

	// Delete the profile
	success, _, sterr = f.TxProfileDelete(fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is deleted
	storedProfiles = f.QueryProfiles()
	require.Empty(t, storedProfiles)

	f.Cleanup()
}

func TestDesmosCLIRequestDTagTransfer(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	barAddr := f.KeyAddress(keyBar)
	calAddr := f.KeyAddress(keyBaz)
	f.TxSend(fooAddr.String(), barAddr, sdk.NewCoin(denom, sdk.NewInt(1000)), "-y")
	f.TxSend(fooAddr.String(), calAddr, sdk.NewCoin(denom, sdk.NewInt(1000)), "-y")

	// Create the profile of the DTag owner
	success, _, sterr := f.TxProfileSave("mrBrown", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.DTag, "mrBrown")

	// Create a request
	success, _, sterr = f.TxProfileRequestDTagTransfer(fooAddr, barAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the request is saved
	storedRequests := f.QueryUserDTagRequests(fooAddr)
	require.NotEmpty(t, storedRequests)

	// Test --dry-run

	// Create the profile of the dTag owner
	success, _, sterr = f.TxProfileSave("mrPink", calAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	success, _, _ = f.TxProfileRequestDTagTransfer(calAddr, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileRequestDTagTransfer(calAddr, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}

func TestDesmosCLIAcceptDTagTransferRequest(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	barAddr := f.KeyAddress(keyBar)
	f.TxSend(fooAddr.String(), barAddr, sdk.NewCoin(denom, sdk.NewInt(1000)), "-y")

	// Create a profile
	success, _, sterr := f.TxProfileSave("mrBrown", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.DTag, "mrBrown")

	// Create a request
	success, _, sterr = f.TxProfileRequestDTagTransfer(fooAddr, barAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the request is saved
	storedRequests := f.QueryUserDTagRequests(fooAddr)
	require.NotEmpty(t, storedRequests)

	// Accept the request
	success, _, sterr = f.TxProfileAcceptDTagTransfer("newDtag", barAddr, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Create a request
	success, _, sterr = f.TxProfileRequestDTagTransfer(fooAddr, barAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Test --dry-run
	success, _, _ = f.TxProfileAcceptDTagTransfer("otherDtag", barAddr, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileAcceptDTagTransfer("otherDtag", barAddr, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}

func TestDesmosCLIMultipleDTagTransferRequest(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	barAddr := f.KeyAddress(keyBar)
	f.TxSend(fooAddr.String(), barAddr, sdk.NewCoin(denom, sdk.NewInt(1000)), "-y")

	// Create a profile for the DTag owner
	success, _, sterr := f.TxProfileSave("mrBrown", fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved and the DTag isn't empty
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.DTag, "mrBrown")

	// Create a request from a user without a profile
	success, _, sterr = f.TxProfileRequestDTagTransfer(fooAddr, barAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the request is saved
	storedRequests := f.QueryUserDTagRequests(fooAddr)
	require.NotEmpty(t, storedRequests)

	// Accept the request
	success, _, sterr = f.TxProfileAcceptDTagTransfer("mrPink", barAddr, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the requests are now empty
	storedRequests = f.QueryUserDTagRequests(fooAddr)
	require.Empty(t, storedRequests)

	receiverRequests := f.QueryUserDTagRequests(barAddr)
	require.Empty(t, receiverRequests)

	// Make sure that the DTag has been transferred properly and the profile for receiver created
	storedProfiles = f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile = storedProfiles[0]
	require.Equal(t, profile.DTag, "mrPink")
	receiverProfile := storedProfiles[1]
	require.Equal(t, receiverProfile.DTag, "mrBrown")

	// Create another request
	success, _, sterr = f.TxProfileRequestDTagTransfer(fooAddr, barAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the request is saved
	storedRequests = f.QueryUserDTagRequests(fooAddr)
	require.NotEmpty(t, storedRequests)
	require.Len(t, storedRequests, 1)

	f.Cleanup()
}
