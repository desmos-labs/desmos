//nolint
package clitest

import (
	"github.com/desmos-labs/desmos/x/profiles/types"
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
	success, _, sterr = f.TxProfileSave(dTag, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the profile is edited
	editedProfiles := f.QueryProfiles()
	require.NotEmpty(t, editedProfiles)

	// Make sure the profile has been edited
	require.Nil(t, editedProfiles[0].Moniker)
	require.Nil(t, editedProfiles[0].Bio)

	// Test --dry-run
	success, _, _ = f.TxProfileSave(dTag, fooAddr, "--dry-run")
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

func TestDesmosCLICreateMonoDirectionalRelationship(t *testing.T) {
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
	receiver, err := sdk.AccAddressFromBech32("desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru")
	require.NoError(t, err)

	// Create mono directional relationship
	success, _, sterr := f.TxCreateMonoDirectionalRelationship(receiver, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is created
	storedRelationships := f.QueryRelationships(fooAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship := types.NewMonodirectionalRelationship(fooAddr, receiver)
	require.Equal(t, expRelationship, storedRelationships[0])

	relationshipID := storedRelationships[0].RelationshipID()

	// Delete the relationship to perform other tests
	success, _, sterr = f.TxDeleteUserRelationship(relationshipID, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Test --dry-tun
	success, _, _ = f.TxCreateMonoDirectionalRelationship(receiver, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxCreateMonoDirectionalRelationship(receiver, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}

func TestDesmosCLIRequestBiDirectionalRelationship(t *testing.T) {
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
	receiver, err := sdk.AccAddressFromBech32("desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru")
	require.NoError(t, err)

	// Create request for bidirectional relationship
	success, _, sterr := f.TxRequestBiDirectionalRelationship(receiver, fooAddr, "hello", "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is created
	storedRelationships := f.QueryRelationships(fooAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship := types.NewBiDirectionalRelationship(fooAddr, receiver, types.Sent)
	require.Equal(t, expRelationship, storedRelationships[0])

	relationshipID := storedRelationships[0].RelationshipID()

	// Delete the relationship to perform other tests
	success, _, sterr = f.TxDeleteUserRelationship(relationshipID, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Test --dry-tun
	success, _, _ = f.TxRequestBiDirectionalRelationship(receiver, fooAddr, "hello", "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxRequestBiDirectionalRelationship(receiver, fooAddr, "hello", "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}

func TestDesmosCLIAcceptBiDirectionalRelationship(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	barAddr := f.KeyAddress(keyBar)

	// Later usage variables
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	sendTokens := sdk.TokensFromConsensusPower(10)
	success, _, sterr := f.TxSend(fooAddr.String(), barAddr, sdk.NewCoin(denom, sendTokens), "-y")
	require.True(t, success)
	require.Empty(t, sterr)

	// Create request for bidirectional relationship
	success, _, sterr = f.TxRequestBiDirectionalRelationship(fooAddr, barAddr, "hello", "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the request is made
	storedRelationships := f.QueryRelationships(barAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship := types.NewBiDirectionalRelationship(barAddr, fooAddr, types.Sent)
	require.Equal(t, expRelationship, storedRelationships[0])

	relationshipID := storedRelationships[0].RelationshipID()

	// Accept the previously made request
	success, _, sterr = f.TxAcceptBiDirectionalRelationship(relationshipID, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the request is accepted
	storedRelationships = f.QueryRelationships(fooAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship = types.NewBiDirectionalRelationship(barAddr, fooAddr, types.Accepted)
	require.Equal(t, expRelationship, storedRelationships[0])

	// Delete the relationship to perform other tests
	success, _, sterr = f.TxDeleteUserRelationship(relationshipID, barAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Create Request for dry-run test
	success, _, _ = f.TxRequestBiDirectionalRelationship(fooAddr, barAddr, "hello", "-y")
	require.True(t, success)

	// Test --dry-tun
	success, _, _ = f.TxAcceptBiDirectionalRelationship(relationshipID, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxAcceptBiDirectionalRelationship(relationshipID, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}

func TestDesmosCLIDenyBiDirectionalRelationship(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// Start Desmosd server
	proc := f.GDStart()
	defer proc.Stop(false)

	// Save key addresses for later use
	fooAddr := f.KeyAddress(keyFoo)
	barAddr := f.KeyAddress(keyBar)

	// Later usage variables
	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromConsensusPower(140)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(denom))

	sendTokens := sdk.TokensFromConsensusPower(10)
	success, _, sterr := f.TxSend(fooAddr.String(), barAddr, sdk.NewCoin(denom, sendTokens), "-y")
	require.True(t, success)
	require.Empty(t, sterr)

	// Create request for bidirectional relationship
	success, _, sterr = f.TxRequestBiDirectionalRelationship(fooAddr, barAddr, "hello", "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the request is made
	storedRelationships := f.QueryRelationships(barAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship := types.NewBiDirectionalRelationship(barAddr, fooAddr, types.Sent)
	require.Equal(t, expRelationship, storedRelationships[0])

	relationshipID := storedRelationships[0].RelationshipID()

	// Accept the previously made request
	success, _, sterr = f.TxDenyBiDirectionalRelationship(relationshipID, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure the request is accepted
	storedRelationships = f.QueryRelationships(fooAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship = types.NewBiDirectionalRelationship(barAddr, fooAddr, types.Denied)
	require.Equal(t, expRelationship, storedRelationships[0])

	// Delete the relationship to perform other tests
	success, _, sterr = f.TxDeleteUserRelationship(relationshipID, barAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Create Request for dry-run test
	success, _, _ = f.TxRequestBiDirectionalRelationship(fooAddr, barAddr, "hello", "-y")
	require.True(t, success)

	// Test --dry-tun
	success, _, _ = f.TxDenyBiDirectionalRelationship(relationshipID, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxDenyBiDirectionalRelationship(relationshipID, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}

func TestDesmosCLIDeleteRelationship(t *testing.T) {
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
	receiver, err := sdk.AccAddressFromBech32("desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru")
	require.NoError(t, err)

	// Create mono directional relationship
	success, _, sterr := f.TxCreateMonoDirectionalRelationship(receiver, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is created
	storedRelationships := f.QueryRelationships(fooAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship := types.NewMonodirectionalRelationship(fooAddr, receiver)
	require.Equal(t, expRelationship, storedRelationships[0])

	relationshipID := storedRelationships[0].RelationshipID()

	// Delete the relationship to perform other tests
	success, _, sterr = f.TxDeleteUserRelationship(relationshipID, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is deleted
	storedRelationships = f.QueryRelationships(fooAddr)
	require.Empty(t, storedRelationships)

	// Create mono directional relationship
	success, _, sterr = f.TxCreateMonoDirectionalRelationship(receiver, fooAddr, "-y")

	// Test --dry-tun
	success, _, _ = f.TxDeleteUserRelationship(relationshipID, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxDeleteUserRelationship(relationshipID, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}
