package keeper_test

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_IsUserBlocked() {
	user, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	otherUser, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

	tests := []struct {
		name       string
		blocker    sdk.AccAddress
		blocked    sdk.AccAddress
		userBlocks []relationships.UserBlock
		expBool    bool
	}{
		{
			name:       "blocked user found returns true",
			blocker:    user,
			blocked:    otherUser,
			userBlocks: []relationships.UserBlock{relationships.NewUserBlock(user, otherUser, "test", "")},
			expBool:    true,
		},
		{
			name:       "non blocked user not found returns false",
			blocker:    user,
			blocked:    otherUser,
			userBlocks: nil,
			expBool:    false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.userBlocks != nil {
				_ = suite.relationshipsKeeper.SaveUserBlock(suite.ctx, test.userBlocks[0])
			}
			res := suite.keeper.IsUserBlocked(suite.ctx, test.blocker, test.blocked)
			suite.Equal(test.expBool, res)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_AssociateDtagWithAddress() {
	store := suite.ctx.KVStore(suite.storeKey)

	suite.keeper.AssociateDtagWithAddress(suite.ctx, "dtag", suite.testData.profile.Creator)

	var owner keeper.DTagOwner
	suite.cdc.MustUnmarshalBinaryBare(store.Get(types.DtagStoreKey("dtag")), &owner)

	suite.Require().Equal(suite.testData.profile.Creator, owner.Address)
}

func (suite *KeeperTestSuite) TestKeeper_GetDtagRelatedAddress() {
	suite.keeper.AssociateDtagWithAddress(suite.ctx, "moner", suite.testData.profile.Creator)

	addr := suite.keeper.GetDtagRelatedAddress(suite.ctx, "moner")
	suite.Require().Equal(suite.testData.profile.Creator, addr)
}

func (suite *KeeperTestSuite) TestKeeper_GetDtagFromAddress() {
	tests := []struct {
		name      string
		dtags     []string
		addresses []string
		expDtag   string
	}{
		{
			name:      "found right dtag",
			dtags:     []string{"lol", "oink"},
			addresses: []string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", suite.testData.profile.Creator},
			expDtag:   "lol",
		},
		{
			name:      "no dtag found",
			dtags:     []string{"lol", "oink"},
			addresses: []string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
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

func (suite *KeeperTestSuite) TestKeeper_DeleteDtagAddressAssociation() {
	suite.keeper.AssociateDtagWithAddress(suite.ctx, "monik", suite.testData.profile.Creator)
	suite.keeper.DeleteDtagAddressAssociation(suite.ctx, "monik")

	addr := suite.keeper.GetDtagRelatedAddress(suite.ctx, "monik")
	suite.Require().Nil(addr)
}

func (suite *KeeperTestSuite) TestKeeper_StoreProfile() {
	tests := []struct {
		name             string
		account          types.Profile
		existentAccounts []types.Profile
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
				Dtag:     suite.testData.profile.Dtag,
				Bio:      suite.testData.profile.Bio,
				Pictures: suite.testData.profile.Pictures,
				Creator:  "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			existentAccounts: []types.Profile{suite.testData.profile},
			expError:         fmt.Errorf("a profile with dtag: dtag has already been created"),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			for _, profile := range test.existentAccounts {
				store.Set(types.ProfileStoreKey(profile.Creator), suite.cdc.MustMarshalBinaryBare(&profile))
				suite.keeper.AssociateDtagWithAddress(suite.ctx, profile.Dtag, profile.Creator)
			}

			err := suite.keeper.StoreProfile(suite.ctx, test.account)
			suite.Require().Equal(test.expError, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetProfile() {
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
			suite.SetupTest()

			store := suite.ctx.KVStore(suite.storeKey)
			if test.existentAccount != nil {
				store.Set(
					types.ProfileStoreKey(test.existentAccount.Creator),
					suite.cdc.MustMarshalBinaryBare(test.existentAccount),
				)
				suite.keeper.AssociateDtagWithAddress(suite.ctx, test.existentAccount.Dtag, test.existentAccount.Creator)
			}

			res, found := suite.keeper.GetProfile(suite.ctx, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

			if test.existentAccount != nil {
				suite.Require().Equal(*test.existentAccount, res)
				suite.True(found)
			} else {
				suite.Require().Equal(types.Profile{}, res)
				suite.Require().False(found)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveProfile() {
	err := suite.keeper.StoreProfile(suite.ctx, suite.testData.profile)
	suite.Require().Nil(err)

	_, found := suite.keeper.GetProfile(suite.ctx, suite.testData.profile.Creator)
	suite.True(found)

	err = suite.keeper.RemoveProfile(suite.ctx, suite.testData.profile.Creator)
	suite.Require().NoError(err)

	_, found = suite.keeper.GetProfile(suite.ctx, suite.testData.profile.Creator)
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestKeeper_GetProfiles() {
	tests := []struct {
		name             string
		existentAccounts []types.Profile
	}{
		{
			name:             "Non empty Profiles list returned",
			existentAccounts: []types.Profile{suite.testData.profile},
		},
		{
			name:             "Profile not found",
			existentAccounts: []types.Profile{},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			if len(test.existentAccounts) != 0 {
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ProfileStoreKey(test.existentAccounts[0].Creator)
				store.Set(key, suite.cdc.MustMarshalBinaryBare(&test.existentAccounts[0]))
			}

			res := suite.keeper.GetProfiles(suite.ctx)

			if len(test.existentAccounts) != 0 {
				suite.Require().Equal(test.existentAccounts, res)
			} else {
				suite.Require().Equal([]types.Profile{}, res)
			}

		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_ValidateProfile() {
	tests := []struct {
		name    string
		profile types.Profile
		expErr  error
	}{
		{
			name: "Max moniker length exceeded",
			profile: types.NewProfile(
				"custom_dtag",
				strings.Repeat("A", 1005),
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.Creator,
			),
			expErr: fmt.Errorf("profile moniker cannot exceed 1000 characters"),
		},
		{
			name: "Min moniker length not reached",
			profile: types.NewProfile(
				"custom_dtag",
				"m",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.Creator,
			),
			expErr: fmt.Errorf("profile moniker cannot be less than 2 characters"),
		},
		{
			name: "Max bio length exceeded",
			profile: types.NewProfile(
				"custom_dtag",
				"moniker",
				strings.Repeat("A", 1005),
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.Creator,
			),
			expErr: fmt.Errorf("profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Invalid dtag doesn't match regEx",
			profile: types.NewProfile(
				"custom.",
				"moniker",
				strings.Repeat("A", 1000),
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.Creator,
			),
			expErr: fmt.Errorf("invalid profile dtag, it should match the following regEx ^[A-Za-z0-9_]+$"),
		},
		{
			name: "Min dtag length not reached",
			profile: types.NewProfile(
				"d",
				"moniker",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.Creator,
			),
			expErr: fmt.Errorf("profile dtag cannot be less than 3 characters"),
		},
		{
			name: "Max dtag length exceeded",
			profile: types.NewProfile(
				"9YfrVVi3UEI1ymN7n6isSct30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEyHNS",
				"moniker",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.Creator,
			),
			expErr: fmt.Errorf("profile dtag cannot exceed 30 characters"),
		},
		{
			name: "Invalid profile pictures returns error",
			profile: types.NewProfile(
				"dtag",
				"moniker",
				"my-bio",
				types.NewPictures(
					"pic",
					"htts://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.Creator,
			),
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Valid profile returns no error",
			profile: types.NewProfile(
				"dtag",
				"moniker",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.Creator,
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.keeper.SetParams(suite.ctx, types.DefaultParams())

			actual := suite.keeper.ValidateProfile(suite.ctx, test.profile)
			suite.Require().Equal(test.expErr, actual)
		})
	}
}

// ___________________________________________________________________________________________________________________

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
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			expErr: fmt.Errorf("the transfer request from %s to %s has already been made",
				suite.testData.otherUser, suite.testData.user),
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name: "different current owner request saved correctly",
			storedTransferReqs: []types.DTagTransferRequest{types.NewDTagTransferRequest(
				"dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.otherUser, suite.testData.user),
			expErr:      nil,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.otherUser, suite.testData.user),
			},
		},
		{
			name: "different receiver request saved correctly",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			expErr:      nil,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
		},
		{
			name: "different dtag request saved correctly",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag1", suite.testData.user, suite.testData.otherUser),
			expErr:      nil,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
				types.NewDTagTransferRequest("dtag1", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name:               "not already present request saved correctly",
			storedTransferReqs: nil,
			transferReq:        types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			expErr:             nil,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedTransferReqs != nil {
				reqs := keeper.NewDTagRequests(test.storedTransferReqs)
				store.Set(
					types.DtagTransferRequestStoreKey(test.storedTransferReqs[0].Receiver),
					suite.cdc.MustMarshalBinaryBare(&reqs),
				)
			}

			actualErr := suite.keeper.SaveDTagTransferRequest(suite.ctx, test.transferReq)
			suite.Require().Equal(test.expErr, actualErr)

			var actualReqs keeper.DTagRequests
			suite.cdc.MustUnmarshalBinaryBare(
				store.Get(types.DtagTransferRequestStoreKey(test.transferReq.Receiver)),
				&actualReqs,
			)
			suite.Require().Equal(test.expStoredTransferReqs, actualReqs.Requests)
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
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			expReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
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
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedReqs != nil {
				reqs := keeper.NewDTagRequests(test.storedReqs)
				store.Set(
					types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
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
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedReqs != nil {
				reqs := keeper.NewDTagRequests(test.storedReqs)
				store.Set(
					types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
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
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedReqs != nil {
				reqs := keeper.NewDTagRequests(test.storedReqs)
				store.Set(
					types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
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
		sender     string
		expReqs    []types.DTagTransferRequest
	}{
		{
			name:       "empty requests array returns error",
			storedReqs: nil,
		},
		{
			name: "no request made by the sender returns error",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
			sender:  suite.testData.otherUser,
			expReqs: nil,
		},
		{
			name: "request removed properly (remaining requests array)",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
			sender: suite.testData.otherUser,
			expReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
		},
		{
			name: "request removed properly (no remaining requests)",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			sender:  suite.testData.otherUser,
			expReqs: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedReqs != nil {
				reqs := keeper.NewDTagRequests(test.storedReqs)
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
				)
			}

			suite.keeper.DeleteDTagTransferRequest(suite.ctx, suite.testData.user, suite.testData.otherUser)
			suite.Require().Equal(test.expReqs, suite.keeper.GetDTagTransferRequests(suite.ctx))
		})
	}
}
