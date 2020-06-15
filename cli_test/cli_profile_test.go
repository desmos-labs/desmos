// +build cli_test

// nolint
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
	moniker := "mrBrown"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave(fooAddr, "-y", "--dtag mrBrown")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.Moniker, moniker)

	// Test --dry-run
	success, _, _ = f.TxProfileSave(fooAddr, "--dry-run", "--dtag mrBrown")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileSave(fooAddr, "--generate-only=true",
		"--moniker mrBrown")
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
	moniker := "mrBrown"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave(fooAddr, "-y",
		"--dtag mrBrown",
		"--moniker Leonardo",
		"--bio biography",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.Moniker, moniker)

	// Test --dry-run
	success, _, _ = f.TxProfileSave(fooAddr, "--dry-run",
		"--dtag mrBrown",
		"--moniker Leonardo",
		"--bio biography",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileSave(fooAddr, "--generate-only=true",
		"--dtag mrBrown",
		"--moniker Leonardo",
		"--bio biography",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
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
	moniker := "mrBrown"
	newMoniker := "mrPink"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave(fooAddr, "-y",
		"--dtag mrBrown",
		"--moniker Leonardo",
		"--surname DiCaprio",
		"--bio biography",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.Moniker, moniker)

	// Edit the profile
	success, _, sterr = f.TxProfileSave(fooAddr, "-y", "--dtag mrPink")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is edited
	editedProfiles := f.QueryProfiles()
	require.NotEmpty(t, editedProfiles)
	editedProfile := editedProfiles[0]
	require.Equal(t, editedProfile.Moniker, newMoniker)

	// Make sure the profile has been edited
	var emptyField *string
	require.Equal(t, emptyField, editedProfiles[0].Moniker)
	require.Equal(t, emptyField, editedProfiles[0].Bio)

	// Test --dry-run
	success, _, _ = f.TxProfileSave(fooAddr, "--dry-run", "--dtag mrPink")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileSave(fooAddr, "--moniker mrPink", "--generate-only=true")
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
	moniker := "mrBrown"
	newMoniker := "mrPink"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave(fooAddr, "-y",
		"--dtag mrBrown",
		"--moniker Leonardo",
		"--bio biography",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.Moniker, moniker)

	// Edit the profile
	success, _, sterr = f.TxProfileSave(fooAddr, "-y",
		"--dtag mrPink",
		"--moniker Leo",
		"--bio HollywoodActor",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is edited
	editedProfiles := f.QueryProfiles()
	require.NotEmpty(t, editedProfiles)

	editedProfile := editedProfiles[0]
	require.Equal(t, editedProfile.Moniker, newMoniker)

	// Make sure the profile has been edited
	require.NotEqual(t, storedProfiles[0].Moniker, editedProfiles[0].Moniker)
	require.NotEqual(t, storedProfiles[0].Bio, editedProfiles[0].Bio)

	// Test --dry-run
	success, _, _ = f.TxProfileSave(fooAddr, "--dry-run",
		"--drat mrPink",
		"--moniker Leo",
		"--bio HollywoodActor",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileSave(fooAddr, "--generate-only=true",
		"--dtag mrPink",
		"--moniker Leo",
		"--bio HollywoodActor",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
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
	moniker := "mrBrown"
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	// Create a profile
	success, _, sterr := f.TxProfileSave(fooAddr, "-y", "--dtag mrBrown")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.Moniker, moniker)

	// Delete the profile
	success, _, sterr = f.TxProfileDelete(fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is deleted
	storedProfiles = f.QueryProfiles()
	require.Empty(t, storedProfiles)

	// Test --dry-run
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
	require.Len(t, storedProfiles, 0)

	f.Cleanup()
}
