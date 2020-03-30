package keeper_test

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func Test_queryAccount(t *testing.T) {

	tests := []struct {
		name          string
		path          []string
		storedAccount types.Account
		expErr        error
	}{
		{
			name:          "Account doesnt exist",
			path:          []string{types.QueryAccount, "monk"},
			storedAccount: testAccount,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("Account with moniker %s doesn't exists", "monk"),
			),
		},
		{
			name:          "Account returned correctly",
			path:          []string{types.QueryAccount, "moniker"},
			storedAccount: testAccount,
			expErr:        nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			err := k.SaveAccount(ctx, test.storedAccount)
			require.Nil(t, err)

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				require.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.storedAccount)
				require.NoError(t, err)
				require.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
				require.Nil(t, result)
			}

		})
	}

}

func Test_queryAccounts(t *testing.T) {

	tests := []struct {
		name          string
		path          []string
		storedAccount *types.Account
		expResult     types.Accounts
	}{
		{
			name:          "Empty Accounts",
			path:          []string{types.QueryAccounts},
			storedAccount: nil,
			expResult:     types.Accounts{},
		},
		{
			name:          "Account returned correctly",
			path:          []string{types.QueryAccounts},
			storedAccount: &testAccount,
			expResult:     types.Accounts{testAccount},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.storedAccount != nil {
				err := k.SaveAccount(ctx, *test.storedAccount)
				require.Nil(t, err)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.path, abci.RequestQuery{})

			if result != nil {
				require.Nil(t, err)
				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.expResult)
				require.NoError(t, err)
				require.Equal(t, string(expectedIndented), string(result))
			}

		})
	}

}
