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
			existentAccount: &testProfile,
			msg: types.NewMsgCreateProfile(
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "An account with moniker moniker already exists"),
		},
		{
			name:            "Profile doesnt exists",
			existentAccount: nil,
			msg: types.NewMsgCreateProfile(
				testProfile.Moniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures,
				testProfile.Creator,
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
				key := types.ProfileStoreKey(test.existentAccount.Creator)
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

				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, createAccountEv)
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
	var newMoniker = "newMoniker"
	var invalidPic = "pic"

	testAcc2 := types.Profile{
		Name:     &name,
		Surname:  &surname,
		Moniker:  "newMoniker",
		Bio:      &bio,
		Pictures: testPictures,
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
			existentAccounts: types.Profiles{testProfile},
			msg: types.NewMsgEditProfile(
				&newMoniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			expErr: nil,
		},
		{
			name:             "Profile not edited because no profile with given account found",
			existentAccounts: nil,
			msg: types.NewMsgEditProfile(
				&newMoniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testProfile.Creator,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("No existent profile to edit for address: cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")),
		},
		{
			name:             "Profile not edited because the new moniker already exists",
			existentAccounts: types.Profiles{testProfile, testAcc2},
			msg: types.NewMsgEditProfile(
				&newMoniker,
				testProfile.Name,
				testProfile.Surname,
				testProfile.Bio,
				testProfile.Pictures.Profile,
				testProfile.Pictures.Cover,
				testPostOwner,
			),
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "A profile with moniker: newMoniker has already been created"),
		},
		{
			name:             "Profile not edited because of the invalid pics uri",
			existentAccounts: types.Profiles{testProfile},
			msg: types.NewMsgEditProfile(
				&newMoniker,
				testProfile.Name,
				testProfile.Surname,
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
					sdk.NewAttribute(types.AttributeProfileMoniker, *test.msg.NewMoniker),
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

				require.Len(t, res.Events, 1)
				require.Contains(t, res.Events, createAccountEv)
			}

		})
	}
}

func TestHandler_GetEditedProfile(t *testing.T) {
	name := "mame"
	surname := "habe"
	bio := "bioh"
	pic := "pic"
	cov := "cov"
	pictures := types.NewPictures(&pic, &cov)
	picEdited := types.NewPictures(&pic, testProfile.Pictures.Cover)
	covEdited := types.NewPictures(&pic, &cov)
	newMoniker := "omn"

	tests := []struct {
		name       string
		profile    types.Profile
		msg        types.MsgEditProfile
		expProfile types.Profile
	}{
		{
			name:    "edited profile correctly",
			profile: testProfile,
			msg: types.NewMsgEditProfile(
				&newMoniker,
				&name,
				&surname,
				&bio,
				&pic,
				&cov,
				testProfile.Creator,
			),
			expProfile: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  newMoniker,
				Bio:      &bio,
				Pictures: pictures,
				Creator:  testProfile.Creator,
			},
		},
		{
			name:    "edited profile correctly 2",
			profile: testProfile,
			msg: types.NewMsgEditProfile(
				&newMoniker,
				&name,
				&surname,
				&bio,
				&pic,
				nil,
				testProfile.Creator,
			),
			expProfile: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  "omn",
				Bio:      &bio,
				Pictures: picEdited,
				Creator:  testProfile.Creator,
			},
		},
		{
			name:    "edited profile correctly 3",
			profile: testProfile,
			msg: types.NewMsgEditProfile(
				&newMoniker,
				&name,
				&surname,
				&bio,
				nil,
				&cov,
				testProfile.Creator,
			),
			expProfile: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  "omn",
				Bio:      &bio,
				Pictures: covEdited,
				Creator:  testProfile.Creator,
			},
		},
		{
			name:    "edited profile correctly 4",
			profile: testProfile,
			msg: types.NewMsgEditProfile(
				nil,
				&name,
				&surname,
				&bio,
				nil,
				&cov,
				testProfile.Creator,
			),
			expProfile: types.Profile{
				Name:     &name,
				Surname:  &surname,
				Moniker:  testProfile.Moniker,
				Bio:      &bio,
				Pictures: covEdited,
				Creator:  testProfile.Creator,
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			edited := keeper.GetEditedProfile(test.profile, test.msg)
			require.Equal(t, test.expProfile, edited)
		})
	}
}
