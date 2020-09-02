//nolint
package clitest

import (
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/stretchr/testify/require"
	"testing"
)

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
	success, _, sterr := f.TxCreateRelationship(receiver, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is created
	storedRelationships := f.QueryRelationships(fooAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship := types.NewRelationshipResponse([]sdk.AccAddress{receiver})
	require.Equal(t, expRelationship, storedRelationships)

	// Delete the relationship to perform other tests
	success, _, sterr = f.TxDeleteUserRelationship(receiver, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Test --dry-tun
	success, _, _ = f.TxCreateRelationship(receiver, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxCreateRelationship(receiver, fooAddr, "--generate-only=true")
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
	success, _, sterr := f.TxCreateRelationship(receiver, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is created
	storedRelationships := f.QueryRelationships(fooAddr)
	require.NotEmpty(t, storedRelationships)
	expRelationship := types.NewRelationshipResponse([]sdk.AccAddress{receiver})
	require.Equal(t, expRelationship, storedRelationships)

	// Delete the relationship to perform other tests
	success, _, sterr = f.TxDeleteUserRelationship(receiver, fooAddr, "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is deleted
	storedRelationships = f.QueryRelationships(fooAddr)
	require.Empty(t, storedRelationships)

	// Create mono directional relationship
	success, _, sterr = f.TxCreateRelationship(receiver, fooAddr, "-y")

	// Test --dry-tun
	success, _, _ = f.TxDeleteUserRelationship(receiver, fooAddr, "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxDeleteUserRelationship(receiver, fooAddr, "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}

func TestDesmosCLIBlockUser(t *testing.T) {
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
	userToBlock, err := sdk.AccAddressFromBech32("desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru")
	require.NoError(t, err)

	// Block user
	success, _, sterr := f.TxBlockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "reason", "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is created
	userBlocks := f.QueryUserBlocks(fooAddr)
	require.NotEmpty(t, userBlocks)
	expRelationship := []types.UserBlock{types.NewUserBlock(fooAddr, userToBlock, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")}
	require.Equal(t, expRelationship, userBlocks)

	// Unblock the user to perform other tests
	success, _, sterr = f.TxUnblockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Test --dry-tun
	success, _, _ = f.TxBlockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "reason", "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxBlockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "reason", "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}

func TestDesmosCLIUnblockUser(t *testing.T) {
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
	userToBlock, err := sdk.AccAddressFromBech32("desmos15ux5mc98jlhsg30dzwwv06ftjs82uy4g3t99ru")
	require.NoError(t, err)

	// Block user
	success, _, sterr := f.TxBlockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "reason", "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Make sure relationship is created
	userBlocks := f.QueryUserBlocks(fooAddr)
	require.NotEmpty(t, userBlocks)
	expRelationship := []types.UserBlock{types.NewUserBlock(fooAddr, userToBlock, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")}
	require.Equal(t, expRelationship, userBlocks)

	// Unblock the user
	success, _, sterr = f.TxUnblockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Block user again for other tests
	success, _, sterr = f.TxBlockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "reason", "-y")
	require.True(t, success)
	require.Empty(t, sterr)
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Test --dry-tun
	success, _, _ = f.TxUnblockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "--dry-run")
	require.True(t, success)

	// Test --generate-only
	success, stdout, stderr := f.TxUnblockUser(userToBlock, fooAddr, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", "--generate-only=true")
	require.Empty(t, stderr)
	require.True(t, success)
	msg := unmarshalStdTx(f.T, stdout)
	require.NotZero(t, msg.Fee.Gas)
	require.Len(t, msg.Msgs, 1)
	require.Len(t, msg.GetSignatures(), 0)

	f.Cleanup()
}
