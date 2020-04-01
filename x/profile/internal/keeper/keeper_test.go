package keeper_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_SaveAccount(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name            string
		account         types.Profile
		existentAccount *types.Profile
		expError        error
	}{
		{
			name:            "Non existent Profile saved correctly",
			account:         testAccount,
			existentAccount: nil,
			expError:        nil,
		},
		{
			name: "Existent account with different creator returns error",
			account: types.Profile{
				Name:     testAccount.Name,
				Surname:  testAccount.Surname,
				Moniker:  testAccount.Moniker,
				Bio:      testAccount.Bio,
				Pictures: testAccount.Pictures,
				Creator:  creator,
			},
			existentAccount: &testAccount,
			expError:        fmt.Errorf("an account with moniker: moniker has already been created"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.existentAccount != nil {
				store := ctx.KVStore(k.StoreKey)
				key := types.ProfileStoreKey(test.existentAccount.Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
			}

			err := k.SaveProfile(ctx, test.account)

			require.Equal(t, test.expError, err)

		})
	}
}

func TestKeeper_DeleteAccount(t *testing.T) {
	ctx, k := SetupTestInput()

	err := k.SaveProfile(ctx, testAccount)
	require.Nil(t, err)

	res, found := k.GetProfile(ctx, testAccount.Moniker)

	require.Equal(t, testAccount, res)
	require.True(t, found)

	k.DeleteProfile(ctx, testAccount.Moniker)

	res, found = k.GetProfile(ctx, testAccount.Moniker)

	require.Equal(t, types.Profile{}, res)
	require.False(t, found)
}

func TestKeeper_GetAccount(t *testing.T) {

	tests := []struct {
		name            string
		existentAccount *types.Profile
		expFound        bool
	}{
		{
			name:            "Profile founded",
			existentAccount: &testAccount,
		},
		{
			name:            "Profile not found",
			existentAccount: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.existentAccount != nil {
				store := ctx.KVStore(k.StoreKey)
				key := types.ProfileStoreKey(test.existentAccount.Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
			}

			res, found := k.GetProfile(ctx, "moniker")

			if test.existentAccount != nil {
				require.Equal(t, *test.existentAccount, res)
				require.True(t, found)
			} else {
				require.Equal(t, types.Profile{}, res)
				require.False(t, found)
			}

		})
	}
}

func TestKeeper_GetAccounts(t *testing.T) {
	tests := []struct {
		name             string
		existentAccounts types.Profiles
	}{
		{
			name:             "Non empty Profiles list returned",
			existentAccounts: types.Profiles{testAccount},
		},
		{
			name:             "Profile not found",
			existentAccounts: types.Profiles{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if len(test.existentAccounts) != 0 {
				store := ctx.KVStore(k.StoreKey)
				key := types.ProfileStoreKey(test.existentAccounts[0].Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccounts[0]))
			}

			res := k.GetProfiles(ctx)

			if len(test.existentAccounts) != 0 {
				require.Equal(t, test.existentAccounts, res)
			} else {
				require.Equal(t, types.Profiles{}, res)
			}

		})
	}
}
