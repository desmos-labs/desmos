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

func Test_handleMsgCreateProfile(t *testing.T) {
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
				*testAccount.Name,
				*testAccount.Surname,
				testAccount.Moniker,
				*testAccount.Bio,
				testAccount.Pictures,
				testAccount.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "An account with moniker moniker already exists"),
		},
		{
			name:            "Profile doesnt exists",
			existentAccount: nil,
			msg: types.NewMsgCreateProfile(
				*testAccount.Name,
				*testAccount.Surname,
				testAccount.Moniker,
				*testAccount.Bio,
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
				key := types.ProfileStoreKey(test.existentAccount.Creator.String())
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				k.AssociateMonikerWithAddress(ctx, test.existentAccount.Moniker, test.existentAccount.Creator)
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
					types.EventTypeProfileCreated,
					sdk.NewAttribute(types.AttributeProfileMoniker, test.msg.Moniker),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				require.Len(t, ctx.EventManager().Events(), 1)
				require.Contains(t, ctx.EventManager().Events(), createAccountEv)
			}

		})
	}
}

func Test_handleMsgEditProfile(t *testing.T) {
	editor, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	var name = "name"
	var surname = "surname"
	var bio = "biography"

	testAcc2 := types.Profile{
		Name:     &name,
		Surname:  &surname,
		Moniker:  "newMoniker",
		Bio:      &bio,
		Pictures: &testPictures,
		Creator:  editor,
	}

	tests := []struct {
		name             string
		existentAccounts types.Profiles
		expAccount       *types.Profile
		msg              types.MsgEditProfile
		expErr           error
	}{
		{
			name:             "Profile edited",
			existentAccounts: types.Profiles{testAccount},
			msg: types.NewMsgEditProfile(
				testAccount.Moniker,
				"newMoniker",
				*testAccount.Name,
				*testAccount.Surname,
				*testAccount.Bio,
				testAccount.Pictures.Profile,
				testAccount.Pictures.Cover,
				testAccount.Creator,
			),
			expErr: nil,
		},
		{
			name:             "Profile not edited because no profile with given account found",
			existentAccounts: nil,
			msg: types.NewMsgEditProfile(
				testAccount.Moniker,
				"newMoniker",
				*testAccount.Name,
				*testAccount.Surname,
				*testAccount.Bio,
				testAccount.Pictures.Profile,
				testAccount.Pictures.Cover,
				testAccount.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("No existent profile to edit for address: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")),
		},
		{
			name:             "Profile not edited because the new moniker already exists",
			existentAccounts: types.Profiles{testAccount, testAcc2},
			msg: types.NewMsgEditProfile(
				testAccount.Moniker,
				"newMoniker",
				*testAccount.Name,
				*testAccount.Surname,
				*testAccount.Bio,
				testAccount.Pictures.Profile,
				testAccount.Pictures.Cover,
				testPostOwner,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "An account with moniker: newMoniker has already been created"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccounts != nil {
				for _, acc := range test.existentAccounts {
					key := types.ProfileStoreKey(acc.Creator.String())
					store.Set(key, k.Cdc.MustMarshalBinaryBare(acc))
					k.AssociateMonikerWithAddress(ctx, acc.Moniker, acc.Creator)
				}
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {

				profiles := k.GetProfiles(ctx)
				require.Len(t, profiles, 1)

				//Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.msg.NewMoniker), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileEdited,
					sdk.NewAttribute(types.AttributeProfileMoniker, test.msg.NewMoniker),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				require.Len(t, ctx.EventManager().Events(), 1)
				require.Contains(t, ctx.EventManager().Events(), createAccountEv)
			}

		})
	}
}

func Test_handleMsgDeleteProfile(t *testing.T) {
	tests := []struct {
		name            string
		existentAccount *types.Profile
		msg             types.MsgDeleteProfile
		expErr          error
	}{
		{
			name:            "Profile doesnt exists",
			existentAccount: nil,
			msg:             types.NewMsgDeleteProfile(testAccount.Creator),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"No profile associated with this address: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
		},
		{
			name:            "Profile deleted successfully",
			existentAccount: &testAccount,
			msg:             types.NewMsgDeleteProfile(testAccount.Creator),
			expErr:          nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccount != nil {
				key := types.ProfileStoreKey(test.existentAccount.Creator.String())
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				k.AssociateMonikerWithAddress(ctx, test.existentAccount.Moniker, test.existentAccount.Creator)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {
				//Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed("moniker"), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileMoniker, "moniker"),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				require.Len(t, ctx.EventManager().Events(), 1)
				require.Contains(t, ctx.EventManager().Events(), createAccountEv)
			}

		})
	}
}
