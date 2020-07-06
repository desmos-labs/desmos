package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/internal/types"
)

func (suite *KeeperTestSuite) TestKeeper_AssociateDtagWithAddress() {
	store := suite.ctx.KVStore(suite.keeper.StoreKey)

	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	suite.keeper.AssociateDtagWithAddress(suite.ctx, "dtag", creator)

	var acc sdk.AccAddress
	key := types.DtagStoreKey("dtag")
	bz := store.Get(key)
	suite.keeper.Cdc.MustUnmarshalBinaryBare(bz, &acc)

	suite.Equal(creator, acc)
}

func (suite *KeeperTestSuite) TestKeeper_GetDtagRelatedAddress() {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	suite.keeper.AssociateDtagWithAddress(suite.ctx, "moner", creator)

	addr := suite.keeper.GetDtagRelatedAddress(suite.ctx, "moner")
	suite.Equal(creator, addr)
}

func (suite *KeeperTestSuite) TestKeeper_DeleteDtagAddressAssociation() {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	suite.keeper.AssociateDtagWithAddress(suite.ctx, "monik", creator)
	suite.keeper.DeleteDtagAddressAssociation(suite.ctx, "monik")

	addr := suite.keeper.GetDtagRelatedAddress(suite.ctx, "monik")
	suite.Nil(addr)
}

func (suite *KeeperTestSuite) TestKeeper_GetDtagFromAddress() {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	creator2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)

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
		suite.SetupTest() //reset
		test := test
		suite.Run(test.name, func() {
			if len(test.addresses) == len(test.dtags) {
				for i, dtag := range test.dtags {
					suite.keeper.AssociateDtagWithAddress(suite.ctx, dtag, test.addresses[i])
				}
			}

			monk := suite.keeper.GetDtagFromAddress(suite.ctx, test.addresses[0])

			suite.Equal(test.expDtag, monk)
		})
	}

}

func (suite *KeeperTestSuite) TestKeeper_SaveProfile() {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

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
		suite.Run(test.name, func() {
			for _, profile := range test.existentAccounts {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				key := types.ProfileStoreKey(profile.Creator)
				store.Set(key, suite.keeper.Cdc.MustMarshalBinaryBare(profile))
				suite.keeper.AssociateDtagWithAddress(suite.ctx, profile.DTag, profile.Creator)
			}

			err := suite.keeper.SaveProfile(suite.ctx, test.account)

			suite.Equal(test.expError, err)

		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteProfile() {
	err := suite.keeper.SaveProfile(suite.ctx, testProfile)
	suite.Nil(err)

	res, found := suite.keeper.GetProfile(suite.ctx, testProfile.Creator)
	suite.Equal(testProfile, res)
	suite.True(found)

	suite.keeper.DeleteProfile(suite.ctx, testProfile.Creator, testProfile.DTag)

	res, found = suite.keeper.GetProfile(suite.ctx, testProfile.Creator)
	suite.Equal(types.Profile{}, res)
	suite.False(found)
}

func (suite *KeeperTestSuite) TestKeeper_GetProfile() {
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
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			if test.existentAccount != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				key := types.ProfileStoreKey(test.existentAccount.Creator)
				store.Set(key, suite.keeper.Cdc.MustMarshalBinaryBare(&test.existentAccount))
				suite.keeper.AssociateDtagWithAddress(suite.ctx, test.existentAccount.DTag, test.existentAccount.Creator)
			}

			res, found := suite.keeper.GetProfile(suite.ctx, testPostOwner)

			if test.existentAccount != nil {
				suite.Equal(*test.existentAccount, res)
				suite.True(found)
			} else {
				suite.Equal(types.Profile{}, res)
				suite.False(found)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetProfiles() {
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
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			if len(test.existentAccounts) != 0 {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				key := types.ProfileStoreKey(test.existentAccounts[0].Creator)
				store.Set(key, suite.keeper.Cdc.MustMarshalBinaryBare(&test.existentAccounts[0]))
			}

			res := suite.keeper.GetProfiles(suite.ctx)

			if len(test.existentAccounts) != 0 {
				suite.Equal(test.existentAccounts, res)
			} else {
				suite.Equal(types.Profiles{}, res)
			}

		})
	}
}
