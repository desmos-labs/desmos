package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_AssociateDtagWithAddress() {
	store := suite.ctx.KVStore(suite.keeper.storeKey)

	suite.keeper.AssociateDtagWithAddress(suite.ctx, "dtag", suite.testData.profile.Creator)

	var acc sdk.AccAddress
	key := types.DtagStoreKey("dtag")
	bz := store.Get(key)
	suite.keeper.cdc.MustUnmarshalBinaryBare(bz, &acc)

	suite.Require().Equal(suite.testData.profile.Creator, acc)
}

func (suite *KeeperTestSuite) TestKeeper_GetDtagRelatedAddress() {
	suite.keeper.AssociateDtagWithAddress(suite.ctx, "moner", suite.testData.profile.Creator)

	addr := suite.keeper.GetDtagRelatedAddress(suite.ctx, "moner")
	suite.Require().Equal(suite.testData.profile.Creator, addr)
}

func (suite *KeeperTestSuite) TestKeeper_DeleteDtagAddressAssociation() {
	suite.keeper.AssociateDtagWithAddress(suite.ctx, "monik", suite.testData.profile.Creator)
	suite.keeper.DeleteDtagAddressAssociation(suite.ctx, "monik")

	addr := suite.keeper.GetDtagRelatedAddress(suite.ctx, "monik")
	suite.Require().Nil(addr)
}

func (suite *KeeperTestSuite) TestKeeper_GetDtagFromAddress() {
	creator, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	tests := []struct {
		name      string
		dtags     []string
		addresses []sdk.AccAddress
		expDtag   string
	}{
		{
			name:      "found right dtag",
			dtags:     []string{"lol", "oink"},
			addresses: []sdk.AccAddress{creator, suite.testData.profile.Creator},
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

			suite.Require().Equal(test.expDtag, monk)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveProfile() {
	// nolint - errcheck
	diffCreator, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	tests := []struct {
		name             string
		account          types.Profile
		existentAccounts types.Profiles
		expError         error
	}{
		{
			name:             "Non existent Profile saved correctly",
			account:          suite.testData.profile,
			existentAccounts: nil,
			expError:         nil,
		},
		{
			name: "Existent account with different creator returns error",
			account: types.Profile{
				DTag:     suite.testData.profile.DTag,
				Bio:      suite.testData.profile.Bio,
				Pictures: suite.testData.profile.Pictures,
				Creator:  diffCreator,
			},
			existentAccounts: types.Profiles{suite.testData.profile},
			expError:         fmt.Errorf("a profile with dtag: dtag has already been created"),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, profile := range test.existentAccounts {
				store := suite.ctx.KVStore(suite.keeper.storeKey)
				key := types.ProfileStoreKey(profile.Creator)
				store.Set(key, suite.keeper.cdc.MustMarshalBinaryBare(profile))
				suite.keeper.AssociateDtagWithAddress(suite.ctx, profile.DTag, profile.Creator)
			}

			err := suite.keeper.StoreProfile(suite.ctx, test.account)

			suite.Require().Equal(test.expError, err)

		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteProfile() {
	err := suite.keeper.StoreProfile(suite.ctx, suite.testData.profile)
	suite.Require().Nil(err)

	res, found := suite.keeper.GetProfile(suite.ctx, suite.testData.profile.Creator)
	suite.Require().Equal(suite.testData.profile, res)
	suite.True(found)

	suite.keeper.RemoveProfile(suite.ctx, suite.testData.profile.Creator, suite.testData.profile.DTag)

	res, found = suite.keeper.GetProfile(suite.ctx, suite.testData.profile.Creator)
	suite.Require().Equal(types.Profile{}, res)
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
			existentAccount: &suite.testData.profile,
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
				store := suite.ctx.KVStore(suite.keeper.storeKey)
				key := types.ProfileStoreKey(test.existentAccount.Creator)
				store.Set(key, suite.keeper.cdc.MustMarshalBinaryBare(&test.existentAccount))
				suite.keeper.AssociateDtagWithAddress(suite.ctx, test.existentAccount.DTag, test.existentAccount.Creator)
			}

			res, found := suite.keeper.GetProfile(suite.ctx, testPostOwner)

			if test.existentAccount != nil {
				suite.Require().Equal(*test.existentAccount, res)
				suite.True(found)
			} else {
				suite.Require().Equal(types.Profile{}, res)
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
			existentAccounts: types.Profiles{suite.testData.profile},
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
				store := suite.ctx.KVStore(suite.keeper.storeKey)
				key := types.ProfileStoreKey(test.existentAccounts[0].Creator)
				store.Set(key, suite.keeper.cdc.MustMarshalBinaryBare(&test.existentAccounts[0]))
			}

			res := suite.keeper.GetProfiles(suite.ctx)

			if len(test.existentAccounts) != 0 {
				suite.Require().Equal(test.existentAccounts, res)
			} else {
				suite.Require().Equal(types.Profiles{}, res)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveDTagTransferRequest() {
	tests := []struct {
		name                  string
		storedTransferReqs    []types.DTagTransferRequest
		transferReq           types.DTagTransferRequest
		expErr                error
		expStoredTransferReqs []types.DTagTransferRequest
	}{
		{
			name: "already present request returns error",
			storedTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			expErr: fmt.Errorf("the transfer request from %s to %s has already been made",
				suite.testData.otherUser, suite.testData.user),
			expStoredTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser)},
		},
		{
			name: "different current owner request saved correctly",
			storedTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.otherUser, suite.testData.user),
			expErr:      nil,
			expStoredTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.otherUser, suite.testData.user)},
		},
		{
			name: "different receiver request saved correctly",
			storedTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			expErr:      nil,
			expStoredTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser), types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.user)},
		},
		{
			name: "different dtag request saved correctly",
			storedTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag1", suite.testData.user, suite.testData.otherUser),
			expErr:      nil,
			expStoredTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser), types.NewDTagTransferRequest(
				"dtag1", suite.testData.user, suite.testData.otherUser)},
		},
		{
			name:               "not already present request saved correctly",
			storedTransferReqs: nil,
			transferReq:        types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			expErr:             nil,
			expStoredTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser)},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			if test.storedTransferReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(test.storedTransferReqs[0].Receiver),
					suite.keeper.cdc.MustMarshalBinaryBare(&test.storedTransferReqs),
				)
			}

			actualErr := suite.keeper.SaveDTagTransferRequest(suite.ctx, test.transferReq)
			suite.Require().Equal(test.expErr, actualErr)

			var actualReqs []types.DTagTransferRequest
			suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.DtagTransferRequestStoreKey(test.transferReq.Receiver)), &actualReqs)
			suite.Require().Equal(test.expStoredTransferReqs, actualReqs)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserDTagTransferRequests() {
	tests := []struct {
		name       string
		storedReqs []types.DTagTransferRequest
		expReqs    []types.DTagTransferRequest
	}{
		{
			name: "returns a non-empty array of dTag requests",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
			expReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name:       "returns an empty array of dTag requests",
			storedReqs: nil,
			expReqs:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			if test.storedReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.cdc.MustMarshalBinaryBare(&test.storedReqs),
				)
			}

			suite.Require().Equal(test.expReqs, suite.keeper.GetUserDTagTransferRequests(suite.ctx, suite.testData.user))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDTagTransferRequests() {
	tests := []struct {
		name       string
		storedReqs []types.DTagTransferRequest
		expReqs    []types.DTagTransferRequest
	}{
		{
			name: "returns a non-empty array of dTag requests",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
			expReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name:       "returns an empty array of dTag requests",
			storedReqs: nil,
			expReqs:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			if test.storedReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.cdc.MustMarshalBinaryBare(&test.storedReqs),
				)
			}

			suite.Require().Equal(test.expReqs, suite.keeper.GetDTagTransferRequests(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteAllDTagTransferRequests() {
	tests := []struct {
		name       string
		storedReqs []types.DTagTransferRequest
		expReqs    []types.DTagTransferRequest
	}{
		{
			name: "returns a non-empty array of dTag requests",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
			expReqs: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			if test.storedReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.cdc.MustMarshalBinaryBare(&test.storedReqs),
				)
			}

			suite.keeper.DeleteAllDTagTransferRequests(suite.ctx, suite.testData.user)
			suite.Require().Equal(test.expReqs, suite.keeper.GetDTagTransferRequests(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteDTagTransferRequest() {
	tests := []struct {
		name       string
		storedReqs []types.DTagTransferRequest
		sender     sdk.AccAddress
		expReqs    []types.DTagTransferRequest
		error      error
	}{
		{
			name:       "empty requests array returns error",
			storedReqs: nil,
			error:      fmt.Errorf("no requests to be deleted"),
		},
		{
			name: "no request made by the sender returns error",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.user),
			},
			sender:  suite.testData.otherUser,
			expReqs: nil,
			error:   fmt.Errorf("no request made by %s", suite.testData.otherUser),
		},
		{
			name: "request removed properly (remaining requests array)",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.user),
			},
			sender: suite.testData.otherUser,
			expReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.user),
			},
			error: nil,
		},
		{
			name: "request removed properly (no remaining requests)",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag", suite.testData.user, suite.testData.otherUser),
			},
			sender:  suite.testData.otherUser,
			expReqs: nil,
			error:   nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			if test.storedReqs != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.cdc.MustMarshalBinaryBare(&test.storedReqs),
				)
			}

			err := suite.keeper.DeleteDTagTransferRequest(suite.ctx, suite.testData.user, suite.testData.otherUser)
			if err != nil {
				suite.Require().Equal(test.error, err)
			} else {
				suite.Require().Equal(test.expReqs, suite.keeper.GetDTagTransferRequests(suite.ctx))
			}
		})
	}
}
