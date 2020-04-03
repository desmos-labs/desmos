// +build cli_test

// nolint
package clitest

import (
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"testing"
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
	success, _, sterr := f.TxProfileCreate(moniker, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.Moniker, moniker)

	// Test --dry-run
	success, _, _ = f.TxProfileCreate(moniker, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileCreate(moniker, fooAddr, "--generate-only=true")
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
	success, _, sterr := f.TxProfileCreate(moniker, fooAddr, "-y",
		"--name Leonardo",
		"--surname DiCaprio",
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
	success, _, _ = f.TxProfileCreate(moniker, fooAddr, "--dry-run",
		"--name Leonardo",
		"--surname DiCaprio",
		"--bio biography",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileCreate(moniker, fooAddr, "--generate-only=true",
		"--name Leonardo",
		"--surname DiCaprio",
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

	// Create an profile
	success, _, sterr := f.TxProfileCreate(moniker, fooAddr, "-y",
		"--name Leonardo",
		"--surname DiCaprio",
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

	// Edit the profile
	success, _, sterr = f.TxProfileEdit(moniker, fooAddr, "-y",
		"--moniker mrPink")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the profile is edited
	editedProfiles := f.QueryProfiles()
	require.NotEmpty(t, editedProfiles)
	editedProfile := editedProfiles[0]
	require.Equal(t, editedProfile.Moniker, newMoniker)

	//Make sure the profile has been edited
	require.Equal(t, storedProfiles[0].Name, editedProfiles[0].Name)
	require.Equal(t, storedProfiles[0].Surname, editedProfiles[0].Surname)

	// Test --dry-run
	success, _, _ = f.TxProfileEdit(moniker, fooAddr, "--dry-run",
		"--moniker mrPink")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileEdit(moniker, fooAddr, "--generate-only=true",
		"--moniker mrPink")
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

	// Create an profile
	success, _, sterr := f.TxProfileCreate(moniker, fooAddr, "-y",
		"--name Leonardo",
		"--surname DiCaprio",
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

	// Edit the profile
	success, _, sterr = f.TxProfileEdit(moniker, fooAddr, "-y",
		"--moniker mrPink",
		"--name Leo",
		"--surname DiCap",
		"--bio HollywoodActor",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the profile is edited
	editedProfiles := f.QueryProfiles()
	require.NotEmpty(t, editedProfiles)
	editedProfile := editedProfiles[0]
	require.Equal(t, editedProfile.Moniker, newMoniker)

	//Make sure the profile has been edited
	require.NotEqual(t, storedProfiles[0].Name, editedProfiles[0].Name)
	require.NotEqual(t, storedProfiles[0].Surname, editedProfiles[0].Surname)

	// Test --dry-run
	success, _, _ = f.TxProfileEdit(moniker, fooAddr, "--dry-run",
		"--moniker mrPink",
		"--name Leo",
		"--surname DiCap",
		"--bio HollywoodActor",
		"--picture https://profilePic.jpg",
		"--cover https://profileCover.jpg")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileEdit(moniker, fooAddr, "--generate-only=true",
		"--moniker mrPink",
		"--name Leo",
		"--surname DiCap",
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

	// Create an profile
	success, _, sterr := f.TxProfileCreate(moniker, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the profile is saved
	storedProfiles := f.QueryProfiles()
	require.NotEmpty(t, storedProfiles)
	profile := storedProfiles[0]
	require.Equal(t, profile.Moniker, moniker)

	// Delete the profile
	success, _, sterr = f.TxProfileDelete(moniker, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	//Make sure the profile is deleted
	storedProfiles = f.QueryProfiles()
	require.Empty(t, storedProfiles)

	// Test --dry-run
	success, _, _ = f.TxProfileDelete(moniker, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxProfileDelete(moniker, fooAddr, "--generate-only=true")
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
