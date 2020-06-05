package keeper_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_AssociateMonikerWithAddress(t *testing.T) {
	ctx, k := SetupTestInput()

	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	moniker := "moniker"

	k.AssociateMonikerWithAddress(ctx, moniker, creator)

	store := ctx.KVStore(k.StoreKey)

	var acc sdk.AccAddress
	key := models.MonikerStoreKey(moniker)
	bz := store.Get(key)
	k.Cdc.MustUnmarshalBinaryBare(bz, &acc)

	require.Equal(t, creator, acc)
}

func TestKeeper_GetMonikerRelatedAddress(t *testing.T) {
	ctx, k := SetupTestInput()

	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	moniker := "moner"

	k.AssociateMonikerWithAddress(ctx, moniker, creator)

	addr := k.GetMonikerRelatedAddress(ctx, moniker)

	require.Equal(t, creator, addr)
}

func TestKeeper_DeleteMonikerAddressAssociation(t *testing.T) {
	ctx, k := SetupTestInput()

	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	moniker := "monik"

	k.AssociateMonikerWithAddress(ctx, moniker, creator)

	k.DeleteMonikerAddressAssociation(ctx, moniker)

	addr := k.GetMonikerRelatedAddress(ctx, moniker)

	require.Nil(t, addr)

}

func TestKeeper_GetMonikerFromAddress(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)
	creator2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	tests := []struct {
		name       string
		monikers   []string
		addresses  []sdk.AccAddress
		expMoniker string
	}{
		{
			name:       "found right moniker",
			monikers:   []string{"lol", "oink"},
			addresses:  []sdk.AccAddress{creator, creator2},
			expMoniker: "lol",
		},
		{
			name:       "no moniker found",
			monikers:   []string{"lol", "oink"},
			addresses:  []sdk.AccAddress{creator},
			expMoniker: "",
		},
	}

	for _, test := range tests {
		ctx, k := SetupTestInput()
		if len(test.addresses) == len(test.monikers) {
			for i, moniker := range test.monikers {
				k.AssociateMonikerWithAddress(ctx, moniker, test.addresses[i])
			}
		}

		monk := k.GetMonikerFromAddress(ctx, test.addresses[0])

		require.Equal(t, test.expMoniker, monk)
	}

}

func TestKeeper_SaveProfile(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

	require.NoError(t, err)

	tests := []struct {
		name             string
		account          models.Profile
		existentAccounts models.Profiles
		expError         error
	}{
		{
			name:             "Non existent Profile saved correctly",
			account:          testProfile,
			existentAccounts: nil,
			expError:         nil,
		},
		{
			name: "Existent account with different creator returns error",
			account: models.Profile{
				Name:     testProfile.Name,
				Surname:  testProfile.Surname,
				Moniker:  testProfile.Moniker,
				Bio:      testProfile.Bio,
				Pictures: testProfile.Pictures,
				Creator:  creator,
			},
			existentAccounts: models.Profiles{testProfile},
			expError:         fmt.Errorf("a profile with moniker: moniker has already been created"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, profile := range test.existentAccounts {
				store := ctx.KVStore(k.StoreKey)
				key := models.ProfileStoreKey(profile.Creator)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(profile))
				k.AssociateMonikerWithAddress(ctx, profile.Moniker, profile.Creator)
			}

			err := k.SaveProfile(ctx, test.account)

			require.Equal(t, test.expError, err)

		})
	}
}

func TestKeeper_DeleteProfile(t *testing.T) {
	ctx, k := SetupTestInput()

	err := k.SaveProfile(ctx, testProfile)
	require.Nil(t, err)

	res, found := k.GetProfile(ctx, testProfile.Creator)

	require.Equal(t, testProfile, res)
	require.True(t, found)

	k.DeleteProfile(ctx, testProfile.Creator, testProfile.Moniker)

	res, found = k.GetProfile(ctx, testProfile.Creator)

	require.Equal(t, models.Profile{}, res)
	require.False(t, found)
}

func TestKeeper_GetProfile(t *testing.T) {
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name            string
		existentAccount *models.Profile
		expFound        bool
	}{
		{
			name:            "Profile founded",
			existentAccount: &testProfile,
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
				key := models.ProfileStoreKey(test.existentAccount.Creator)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				k.AssociateMonikerWithAddress(ctx, test.existentAccount.Moniker, test.existentAccount.Creator)
			}

			res, found := k.GetProfile(ctx, testPostOwner)

			if test.existentAccount != nil {
				require.Equal(t, *test.existentAccount, res)
				require.True(t, found)
			} else {
				require.Equal(t, models.Profile{}, res)
				require.False(t, found)
			}

		})
	}
}

func TestKeeper_GetProfiles(t *testing.T) {
	tests := []struct {
		name             string
		existentAccounts models.Profiles
	}{
		{
			name:             "Non empty Profiles list returned",
			existentAccounts: models.Profiles{testProfile},
		},
		{
			name:             "Profile not found",
			existentAccounts: models.Profiles{},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if len(test.existentAccounts) != 0 {
				store := ctx.KVStore(k.StoreKey)
				key := models.ProfileStoreKey(test.existentAccounts[0].Creator)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccounts[0]))
			}

			res := k.GetProfiles(ctx)

			if len(test.existentAccounts) != 0 {
				require.Equal(t, test.existentAccounts, res)
			} else {
				require.Equal(t, models.Profiles{}, res)
			}

		})
	}
}
