package keeper_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func Test_handleMsgCreateAccount(t *testing.T) {
	tests := []struct {
		name            string
		existentAccount *types.Profile
		msg             types.MsgCreateProfile
		expErr          error
	}{
		{
			name:            "Profile already exists",
			existentAccount: &testAccount,
			msg: types.NewMsgCreateProfile(
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
			name:            "Profile doesnt exists",
			existentAccount: nil,
			msg: types.NewMsgCreateProfile(
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
				key := types.ProfileStoreKey(test.existentAccount.Moniker)
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
		existentAccount *types.Profile
		msg             types.MsgEditProfile
		expErr          error
	}{
		{
			name:            "Profile edited",
			existentAccount: &testAccount,
			msg: types.NewMsgEditProfile(
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
			name:            "Profile not edited because the new moniker already exists",
			existentAccount: &testAccount,
			msg: types.NewMsgEditProfile(
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
				key := types.ProfileStoreKey(test.existentAccount.Moniker)
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
		existentAccount *types.Profile
		msg             types.MsgDeleteProfile
		expErr          error
	}{
		{
			name:            "Profile doesnt exists",
			existentAccount: nil,
			msg:             types.NewMsgDeleteProfile("moniker", testAccount.Creator),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("An account with %s moniker doesn't exist", "moniker")),
		},
		{
			name:            "Profile not owned by user",
			existentAccount: &testAccount,
			msg:             types.NewMsgDeleteProfile("moniker", user),
			expErr:          sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("You cannot delete an account that is not yours")),
		},
		{
			name:            "Profile deleted successfully",
			existentAccount: &testAccount,
			msg:             types.NewMsgDeleteProfile("moniker", testAccount.Creator),
			expErr:          nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccount != nil {
				key := types.ProfileStoreKey(test.existentAccount.Moniker)
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
