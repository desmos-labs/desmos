package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_SaveAccount(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name            string
		account         types.Account
		existentAccount *types.Account
		expError        error
	}{
		{
			name:            "Non existent Account saved correctly",
			account:         testAccount,
			existentAccount: nil,
			expError:        nil,
		},
		{
			name: "Existent account with different creator returns error",
			account: types.Account{
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
				key := types.AccountStoreKey(test.existentAccount.Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
			}

			err := k.SaveAccount(ctx, test.account)

			require.Equal(t, test.expError, err)

		})
	}
}

func TestKeeper_DeleteAccount(t *testing.T) {
	ctx, k := SetupTestInput()

	err := k.SaveAccount(ctx, testAccount)
	require.Nil(t, err)

	res, found := k.GetAccount(ctx, testAccount.Moniker)

	require.Equal(t, testAccount, res)
	require.True(t, found)

	k.DeleteAccount(ctx, testAccount.Moniker)

	res, found = k.GetAccount(ctx, testAccount.Moniker)

	require.Equal(t, types.Account{}, res)
	require.False(t, found)
}

func TestKeeper_GetAccount(t *testing.T) {

	tests := []struct {
		name            string
		existentAccount *types.Account
		expFound        bool
	}{
		{
			name:            "Account founded",
			existentAccount: &testAccount,
		},
		{
			name:            "Account not found",
			existentAccount: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.existentAccount != nil {
				store := ctx.KVStore(k.StoreKey)
				key := types.AccountStoreKey(test.existentAccount.Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
			}

			res, found := k.GetAccount(ctx, "moniker")

			if test.existentAccount != nil {
				require.Equal(t, *test.existentAccount, res)
				require.True(t, found)
			} else {
				require.Equal(t, types.Account{}, res)
				require.False(t, found)
			}

		})
	}
}

func TestKeeper_GetAccounts(t *testing.T) {
	tests := []struct {
		name             string
		existentAccounts types.Accounts
	}{
		{
			name:             "Non empty Accounts list returned",
			existentAccounts: types.Accounts{testAccount},
		},
		{
			name:             "Account not found",
			existentAccounts: types.Accounts{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if len(test.existentAccounts) != 0 {
				store := ctx.KVStore(k.StoreKey)
				key := types.AccountStoreKey(test.existentAccounts[0].Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccounts[0]))
			}

			res := k.GetAccounts(ctx)

			if len(test.existentAccounts) != 0 {
				require.Equal(t, test.existentAccounts, res)
			} else {
				require.Equal(t, types.Accounts{}, res)
			}

		})
	}
}
