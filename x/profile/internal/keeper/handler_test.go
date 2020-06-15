package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func Test_handleMsgSaveProfile(t *testing.T) {
	editor, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	var bio = "biography"
	var newDtag = "newDtag"
	var invalidPic = "pic"

	testAcc2 := types.Profile{
		DTag:     "newDtag",
		Bio:      &bio,
		Pictures: testPictures,
		Creator:  editor,
	}

	tests := []struct {
		name             string
		existentAccounts types.Profiles
		expAccount       *types.Profile
		msg              types.MsgSaveProfile
		expErr           error
	}{
		{
			name:             "Profile saved (with previous profile created)",
			existentAccounts: types.Profiles{testProfile},
			msg: types.NewMsgSaveProfile(
				newDtag,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			expErr: nil,
		},
		{
			name:             "Profile saved (with no previous profile created)",
			existentAccounts: nil,
			msg: types.NewMsgSaveProfile(
				newDtag,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			expErr: nil,
		},
		{
			name:             "Profile not edited because the new dtag already exists",
			existentAccounts: types.Profiles{testProfile, testAcc2},
			msg: types.NewMsgSaveProfile(
				newDtag,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testPostOwner,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "a profile with dtag: newDtag has already been created"),
		},
		{
			name:             "Profile not edited because of the invalid pics uri",
			existentAccounts: types.Profiles{testProfile},
			msg: types.NewMsgSaveProfile(
				newDtag,
				testProfile.Bio,
				&invalidPic,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid profile picture uri provided"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccounts != nil {
				for _, acc := range test.existentAccounts {
					key := types.ProfileStoreKey(acc.Creator)
					store.Set(key, k.Cdc.MustMarshalBinaryBare(acc))
					k.AssociateDtagWithAddress(ctx, acc.DTag, acc.Creator)
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
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed(test.msg.Dtag), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileSaved,
					sdk.NewAttribute(types.AttributeProfileDtag, test.msg.Dtag),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, createAccountEv)
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
			msg:             types.NewMsgDeleteProfile(testProfile.Creator),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"No profile associated with this address: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
		},
		{
			name:            "Profile deleted successfully",
			existentAccount: &testProfile,
			msg:             types.NewMsgDeleteProfile(testProfile.Creator),
			expErr:          nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			if test.existentAccount != nil {
				key := types.ProfileStoreKey(test.existentAccount.Creator)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				k.AssociateDtagWithAddress(ctx, test.existentAccount.DTag, test.existentAccount.Creator)
			}

			handler := keeper.NewHandler(k)
			res, err := handler(ctx, test.msg)

			if res == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
			}
			if res != nil {
				//Check the data
				require.Equal(t, k.Cdc.MustMarshalBinaryLengthPrefixed("dtag"), res.Data)

				//Check the events
				createAccountEv := sdk.NewEvent(
					types.EventTypeProfileDeleted,
					sdk.NewAttribute(types.AttributeProfileDtag, "dtag"),
					sdk.NewAttribute(types.AttributeProfileCreator, test.msg.Creator.String()),
				)

				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, createAccountEv)
			}

		})
	}
}
