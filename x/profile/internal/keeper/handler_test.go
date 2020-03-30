package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_handleMsgCreateAccount(t *testing.T) {
	tests := []struct {
		name            string
		existentAccount *types.Account
		msg             types.MsgCreateAccount
		expErr          error
	}{
		{
			name:            "Account already exists",
			existentAccount: &testAccount,
			msg: types.NewMsgCreateAccount(
				testAccount.Name,
				testAccount.Surname,
				testAccount.Moniker,
				testAccount.Bio,
				testAccount.Pictures,
				testAccount.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "An account with moniker moniker already exist"),
		},
		{
			name:            "Account doesnt exists",
			existentAccount: nil,
			msg: types.NewMsgCreateAccount(
				testAccount.Name,
				testAccount.Surname,
				testAccount.Moniker,
				testAccount.Bio,
				testAccount.Pictures,
				testAccount.Creator,
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccount != nil {
				key := types.AccountStoreKey(test.existentAccount.Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {
				//Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.msg.Moniker), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeAccountCreated,
					sdk.NewAttribute(types.AttributeAccountMoniker, test.msg.Moniker),
					sdk.NewAttribute(types.AttributeAccountCreator, test.msg.Creator.String()),
				)

				require.Len(t, ctx.EventManager().Events(), 1)
				require.Contains(t, ctx.EventManager().Events(), createAccountEv)
			}

		})
	}
}

func Test_handleMsgEditAccount(t *testing.T) {
	editor, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name            string
		existentAccount *types.Account
		msg             types.MsgEditAccount
		expErr          error
	}{
		{
			name:            "Account edited",
			existentAccount: &testAccount,
			msg: types.NewMsgEditAccount(
				testAccount.Name,
				testAccount.Surname,
				testAccount.Moniker,
				testAccount.Bio,
				testAccount.Pictures,
				testAccount.Creator,
			),
			expErr: nil,
		},
		{
			name:            "Account not edited because the new moniker already exists",
			existentAccount: &testAccount,
			msg: types.NewMsgEditAccount(
				testAccount.Name,
				testAccount.Surname,
				testAccount.Moniker,
				testAccount.Bio,
				testAccount.Pictures,
				editor,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "an account with moniker: moniker has already been created"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccount != nil {
				key := types.AccountStoreKey(test.existentAccount.Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {
				//Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.msg.Moniker), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeAccountEdited,
					sdk.NewAttribute(types.AttributeAccountMoniker, test.msg.Moniker),
					sdk.NewAttribute(types.AttributeAccountCreator, test.msg.Creator.String()),
				)

				require.Len(t, ctx.EventManager().Events(), 1)
				require.Contains(t, ctx.EventManager().Events(), createAccountEv)
			}

		})
	}
}

func Test_handleMsgDeleteAccount(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name            string
		existentAccount *types.Account
		msg             types.MsgDeleteAccount
		expErr          error
	}{
		{
			name:            "Account doesnt exists",
			existentAccount: nil,
			msg:             types.NewMsgDeleteAccount("moniker", testAccount.Creator),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("An account with %s moniker doesn't exist", "moniker")),
		},
		{
			name:            "Account not owned by user",
			existentAccount: &testAccount,
			msg:             types.NewMsgDeleteAccount("moniker", user),
			expErr:          sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("You cannot delete an account that is not yours")),
		},
		{
			name:            "Account deleted successfully",
			existentAccount: &testAccount,
			msg:             types.NewMsgDeleteAccount("moniker", testAccount.Creator),
			expErr:          nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccount != nil {
				key := types.AccountStoreKey(test.existentAccount.Moniker)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {
				//Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.msg.Moniker), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeAccountDeleted,
					sdk.NewAttribute(types.AttributeAccountMoniker, test.msg.Moniker),
					sdk.NewAttribute(types.AttributeAccountCreator, test.msg.Creator.String()),
				)

				require.Len(t, ctx.EventManager().Events(), 1)
				require.Contains(t, ctx.EventManager().Events(), createAccountEv)
			}

		})
	}
}
