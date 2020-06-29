package keeper_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/profiles/internal/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestKeeper_AssociateDtagWithAddress(t *testing.T) {
	ctx, k := SetupTestInput()
	store := ctx.KVStore(k.StoreKey)

	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	k.AssociateDtagWithAddress(ctx, "dtag", creator)

	var acc sdk.AccAddress
	key := types.DtagStoreKey("dtag")
	bz := store.Get(key)
	k.Cdc.MustUnmarshalBinaryBare(bz, &acc)

	require.Equal(t, creator, acc)
}

func TestKeeper_GetDtagRelatedAddress(t *testing.T) {
	ctx, k := SetupTestInput()

	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	k.AssociateDtagWithAddress(ctx, "moner", creator)

	addr := k.GetDtagRelatedAddress(ctx, "moner")
	require.Equal(t, creator, addr)
}

func TestKeeper_DeleteDtagAddressAssociation(t *testing.T) {
	ctx, k := SetupTestInput()

	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	k.AssociateDtagWithAddress(ctx, "monik", creator)
	k.DeleteDtagAddressAssociation(ctx, "monik")

	addr := k.GetDtagRelatedAddress(ctx, "monik")
	require.Nil(t, addr)
}

func TestKeeper_GetDtagFromAddress(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	creator2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	tests := []struct {
		name      string
		dtags     []string
		addresses []sdk.AccAddress
		expDtag   string
	}{
		{
			name:      "found right dtag",
			dtags:     []string{"lol", "oink"},
			addresses: []sdk.AccAddress{creator, creator2},
			expDtag:   "lol",
		},
		{
			name:      "no dtag found",
			dtags:     []string{"lol", "oink"},
			addresses: []sdk.AccAddress{creator},
			expDtag:   "",
		},
	}

	for _, test := range tests {
		ctx, k := SetupTestInput()
		if len(test.addresses) == len(test.dtags) {
			for i, dtag := range test.dtags {
				k.AssociateDtagWithAddress(ctx, dtag, test.addresses[i])
			}
		}

		monk := k.GetDtagFromAddress(ctx, test.addresses[0])

		require.Equal(t, test.expDtag, monk)
	}

}

func TestKeeper_SaveProfile(t *testing.T) {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name             string
		account          types.Profile
		existentAccounts types.Profiles
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
			account: types.Profile{
				DTag:     testProfile.DTag,
				Bio:      testProfile.Bio,
				Pictures: testProfile.Pictures,
				Creator:  creator,
			},
			existentAccounts: types.Profiles{testProfile},
			expError:         fmt.Errorf("a profile with dtag: dtag has already been created"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, profile := range test.existentAccounts {
				store := ctx.KVStore(k.StoreKey)
				key := types.ProfileStoreKey(profile.Creator)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(profile))
				k.AssociateDtagWithAddress(ctx, profile.DTag, profile.Creator)
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

	k.DeleteProfile(ctx, testProfile.Creator, testProfile.DTag)

	res, found = k.GetProfile(ctx, testProfile.Creator)
	require.Equal(t, types.Profile{}, res)
	require.False(t, found)
}

func TestKeeper_GetProfile(t *testing.T) {
	var testPostOwner, _ = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name            string
		existentAccount *types.Profile
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
				key := types.ProfileStoreKey(test.existentAccount.Creator)
				store.Set(key, k.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				k.AssociateDtagWithAddress(ctx, test.existentAccount.DTag, test.existentAccount.Creator)
			}

			res, found := k.GetProfile(ctx, testPostOwner)

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

func TestKeeper_GetProfiles(t *testing.T) {
	tests := []struct {
		name             string
		existentAccounts types.Profiles
	}{
		{
			name:             "Non empty Profiles list returned",
			existentAccounts: types.Profiles{testProfile},
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
				key := types.ProfileStoreKey(test.existentAccounts[0].Creator)
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
