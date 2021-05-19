package keeper_test

import (
	"fmt"
	"strings"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_IsUserBlocked() {
	tests := []struct {
		name       string
		blocker    string
		blocked    string
		userBlocks []types.UserBlock
		expBool    bool
	}{
		{
			name:    "blocked user found returns true",
			blocker: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			userBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"test",
					"",
				),
			},
			expBool: true,
		},
		{
			name:       "non blocked user not found returns false",
			blocker:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			userBlocks: nil,
			expBool:    false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.userBlocks != nil {
				_ = suite.k.SaveUserBlock(suite.ctx, test.userBlocks[0])
			}
			res := suite.k.IsUserBlocked(suite.ctx, test.blocker, test.blocked)
			suite.Equal(test.expBool, res)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_StoreProfile() {
	addr, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.Require().NoError(err)

	accountAny, err := codectypes.NewAnyWithValue(authtypes.NewBaseAccountWithAddress(addr))
	suite.Require().NoError(err)

	updatedProfile, err := types.NewProfile(
		"updated-dtag",
		"updated-nickname",
		suite.testData.profile.Bio,
		suite.testData.profile.Pictures,
		suite.testData.profile.CreationDate,
		suite.testData.profile.GetAccount(),
	)
	suite.Require().NoError(err)

	tests := []struct {
		name           string
		account        *types.Profile
		storedProfiles []*types.Profile
		shouldErr      bool
	}{
		{
			name:      "Non existent profile is saved correctly",
			account:   suite.testData.profile,
			shouldErr: false,
		},
		{
			name: "Edited profile is saved correctly",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			account:   updatedProfile,
			shouldErr: false,
		},
		{
			name: "Existent account with different creator returns error",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			account: &types.Profile{
				DTag:     suite.testData.profile.DTag,
				Bio:      suite.testData.profile.Bio,
				Pictures: suite.testData.profile.Pictures,
				Account:  accountAny,
			},
			shouldErr: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, profile := range test.storedProfiles {
				err = suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			err = suite.k.StoreProfile(suite.ctx, test.account)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				// Verify the DTag -> Address association
				store := suite.ctx.KVStore(suite.storeKey)
				suite.Require().Equal(test.account.GetAddress().Bytes(), store.Get(types.DTagStoreKey(test.account.DTag)),
					"DTag -> Address association not correct")

				for _, stored := range test.storedProfiles {
					// Make sure that if the DTag has been edited, the old association has been removed
					if stored.GetAddress().Equals(test.account.GetAddress()) && stored.DTag != test.account.DTag {
						suite.Require().Nil(store.Get(types.DTagStoreKey(stored.DTag)),
							"Old DTag -> Address association still exists")
					}
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_StoreProfile_Update() {
	// Store the initial profile
	suite.Require().NoError(suite.k.StoreProfile(suite.ctx, suite.testData.profile))

	// Verify the store keys
	store := suite.ctx.KVStore(suite.storeKey)
	suite.Require().Equal(
		suite.testData.profile.GetAddress().String(),
		sdk.AccAddress(store.Get(types.DTagStoreKey(suite.testData.profile.DTag))).String(),
	)

	oldAccounts := suite.ak.GetAllAccounts(suite.ctx)
	suite.Require().Len(oldAccounts, 1)

	// Update the profile
	updatedProfile, err := types.NewProfile(
		suite.testData.profile.DTag+"-update",
		"",
		"",
		types.NewPictures("", ""),
		suite.testData.profile.CreationDate,
		suite.ak.GetAccount(suite.ctx, suite.testData.profile.GetAddress()),
	)
	suite.Require().NoError(err)
	suite.Require().NoError(suite.k.StoreProfile(suite.ctx, updatedProfile))

	// Verify the store keys
	suite.Require().Nil(
		store.Get(types.DTagStoreKey(suite.testData.profile.DTag)),
	)
	suite.Require().Equal(
		suite.testData.profile.GetAddress().String(),
		sdk.AccAddress(store.Get(types.DTagStoreKey(suite.testData.profile.DTag+"-update"))).String(),
	)

	newAccounts := suite.ak.GetAllAccounts(suite.ctx)
	suite.Require().Len(newAccounts, 1)

	for _, account := range newAccounts {
		suite.Require().NotContains(oldAccounts, account)
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetProfile() {
	tests := []struct {
		name           string
		storedProfiles []*types.Profile
		address        string
		shouldErr      bool
		expFound       bool
		expProfile     *types.Profile
	}{
		{
			name:           "Invalid address",
			storedProfiles: nil,
			address:        "",
			shouldErr:      true,
		},
		{
			name: "Profile found",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			address:    suite.testData.profile.GetAddress().String(),
			shouldErr:  false,
			expFound:   true,
			expProfile: suite.testData.profile,
		},
		{
			name:           "Profile not found",
			storedProfiles: []*types.Profile{},
			address:        suite.testData.profile.GetAddress().String(),
			expFound:       false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, profile := range test.storedProfiles {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			res, found, err := suite.k.GetProfile(suite.ctx, test.address)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expFound, found)

				if found {
					suite.Require().Equal(test.expProfile, res)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAddressFromDTag() {
	tests := []struct {
		name    string
		profile *types.Profile
		dTag    string
		expAddr string
	}{
		{
			name:    "valid profile returns correct address",
			profile: suite.testData.profile,
			dTag:    suite.testData.profile.DTag,
			expAddr: suite.testData.profile.GetAddress().String(),
		},
		{
			name:    "non existing profile returns empty address",
			profile: nil,
			dTag:    "test",
			expAddr: "",
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.profile != nil {
				err := suite.k.StoreProfile(suite.ctx, test.profile)
				suite.Require().NoError(err)
			}

			addr := suite.k.GetAddressFromDTag(suite.ctx, test.dTag)
			suite.Require().Equal(test.expAddr, addr)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveProfile() {
	err := suite.k.StoreProfile(suite.ctx, suite.testData.profile)
	suite.Require().Nil(err)

	_, found, _ := suite.k.GetProfile(suite.ctx, suite.testData.profile.GetAddress().String())
	suite.True(found)

	err = suite.k.RemoveProfile(suite.ctx, suite.testData.profile.GetAddress().String())
	suite.Require().NoError(err)

	_, found, _ = suite.k.GetProfile(suite.ctx, suite.testData.profile.GetAddress().String())
	suite.Require().False(found)

	addr := suite.k.GetAddressFromDTag(suite.ctx, suite.testData.profile.DTag)
	suite.Require().Equal("", addr)
}

func (suite *KeeperTestSuite) TestKeeper_ValidateProfile() {
	tests := []struct {
		name    string
		profile *types.Profile
		expErr  error
	}{
		{
			name: "Max nickname length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				strings.Repeat("A", 1005),
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile nickname cannot exceed 1000 characters"),
		},
		{
			name: "Min nickname length not reached",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				"m",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile nickname cannot be less than 2 characters"),
		},
		{
			name: "Max bio length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom_dtag",
				"nickname",
				strings.Repeat("A", 1005),
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile biography cannot exceed 1000 characters"),
		},
		{
			name: "Invalid dtag doesn't match regEx",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"custom.",
				"nickname",
				strings.Repeat("A", 1000),
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("invalid profile dtag, it should match the following regEx ^[A-Za-z0-9_]+$"),
		},
		{
			name: "Min dtag length not reached",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"d",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile dtag cannot be less than 3 characters"),
		},
		{
			name: "Max dtag length exceeded",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"9YfrVVi3UEI1ymN7n6isSct30xG6Jn1EDxEXxWOn0voSMIKqLhHsBfnZoXEyHNS",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("profile dtag cannot exceed 30 characters"),
		},
		{
			name: "Invalid profile pictures returns error",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"dtag",
				"nickname",
				"my-bio",
				types.NewPictures(
					"pic",
					"htts://test.com/cover-pic",
				),
				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: fmt.Errorf("invalid profile picture uri provided"),
		},
		{
			name: "Valid profile returns no error",
			profile: suite.CheckProfileNoError(types.NewProfile(
				"dtag",
				"nickname",
				"my-bio",
				types.NewPictures(
					"https://test.com/profile-picture",
					"https://test.com/cover-pic",
				),

				suite.testData.profile.CreationDate,
				suite.testData.profile.GetAccount(),
			)),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.k.SetParams(suite.ctx, types.DefaultParams())

			actual := suite.k.ValidateProfile(suite.ctx, test.profile)
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
		shouldErr             bool
		expStoredTransferReqs []types.DTagTransferRequest
	}{
		{
			name: "Already present request returns error",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			shouldErr:   true,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name: "Different sender request is saved properly",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.otherUser, suite.testData.user),
			shouldErr:   false,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
				types.NewDTagTransferRequest("dtag", suite.testData.otherUser, suite.testData.user),
			},
		},
		{
			name: "Different receiver request is saved correctly",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			shouldErr:   false,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name: "Different DTag request returns an error",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag1", suite.testData.user, suite.testData.otherUser),
			shouldErr:   true,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name:               "Not already present request is saved correctly",
			storedTransferReqs: nil,
			transferReq:        types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			shouldErr:          false,
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedTransferReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			err := suite.k.SaveDTagTransferRequest(suite.ctx, test.transferReq)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			stored := suite.k.GetDTagTransferRequests(suite.ctx)
			suite.Require().Len(stored, len(test.expStoredTransferReqs))
			for _, req := range stored {
				suite.Require().Contains(test.expStoredTransferReqs, req)
			}
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
				reqs := types.NewDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DTagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
				)
			}

			suite.Require().Equal(test.expReqs, suite.k.GetUserIncomingDTagTransferRequests(suite.ctx, suite.testData.user))
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
				reqs := types.NewDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DTagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
				)
			}

			suite.Require().Equal(test.expReqs, suite.k.GetDTagTransferRequests(suite.ctx))
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
				reqs := types.NewDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DTagTransferRequestStoreKey(suite.testData.user),
					suite.cdc.MustMarshalBinaryBare(&reqs),
				)
			}

			suite.k.DeleteAllDTagTransferRequests(suite.ctx, suite.testData.user)
			suite.Require().Equal(test.expReqs, suite.k.GetDTagTransferRequests(suite.ctx))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteDTagTransferRequest() {
	tests := []struct {
		name            string
		storedReqs      []types.DTagTransferRequest
		sender          string
		receiver        string
		shouldErr       bool
		storedReqsAfter []types.DTagTransferRequest
	}{
		{
			name:       "Empty requests array returns error",
			storedReqs: nil,
			sender:     suite.testData.user,
			receiver:   suite.testData.otherUser,
			shouldErr:  true,
		},
		{
			name: "Deleting non existent request returns an error and doesn't change the store",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
			sender:    suite.testData.user,
			receiver:  suite.testData.otherUser,
			shouldErr: true,
			storedReqsAfter: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
		},
		{
			name: "Existing request gets removed properly and leaves an array",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
			sender:   suite.testData.user,
			receiver: suite.testData.otherUser,
			storedReqsAfter: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
		},
		{
			name: "Existing request gets removed properly and doesn't leave anything",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			sender:          suite.testData.user,
			receiver:        suite.testData.otherUser,
			storedReqsAfter: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			err := suite.k.DeleteDTagTransferRequest(suite.ctx, test.sender, test.receiver)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			reqs := suite.k.GetDTagTransferRequests(suite.ctx)
			suite.Require().Equal(test.storedReqsAfter, reqs)
		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestKeeper_SaveRelationship() {
	tests := []struct {
		name             string
		stored           []types.Relationship
		user             string
		relationship     types.Relationship
		expErr           bool
		expRelationships []types.Relationship
	}{
		{
			name: "already existent relationship returns error",
			stored: []types.Relationship{
				types.NewRelationship("user", "recipient", "subspace"),
			},
			user:         "user",
			relationship: types.NewRelationship("user", "recipient", "subspace"),
			expErr:       true,
		},
		{
			name:         "relationship added correctly",
			stored:       nil,
			user:         "user",
			relationship: types.NewRelationship("user", "recipient", "subspace"),
			expErr:       false,
			expRelationships: []types.Relationship{
				types.NewRelationship("user", "recipient", "subspace"),
			},
		},
		{
			name: "relationship added correctly (another subspace)",
			stored: []types.Relationship{
				types.NewRelationship("user", "recipient", "subspace"),
			},
			user:         "user",
			relationship: types.NewRelationship("user", "recipient", "subspace_2"),
			expErr:       false,
		},
		{
			name: "relationship added correctly (another receiver)",
			stored: []types.Relationship{
				types.NewRelationship("user", "recipient", "subspace"),
			},
			user:         "user",
			relationship: types.NewRelationship("user", "user", "subspace"),
			expErr:       false,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, relationship := range test.stored {
				err := suite.k.SaveRelationship(suite.ctx, relationship)
				suite.Require().NoError(err)
			}

			err := suite.k.SaveRelationship(suite.ctx, test.relationship)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllRelationships() {
	tests := []struct {
		name     string
		stored   []types.Relationship
		expected []types.Relationship
	}{
		{
			name: "non empty relationships slice is returned properly",
			stored: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
				types.NewRelationship("creator", "another_recipient", "subspace"),
				types.NewRelationship("recipient", "creator", "subspace"),
				types.NewRelationship("recipient", "creator", "subspace_2"),
			},
			expected: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
				types.NewRelationship("creator", "another_recipient", "subspace"),
				types.NewRelationship("recipient", "creator", "subspace"),
				types.NewRelationship("recipient", "creator", "subspace_2"),
			},
		},
		{
			name:     "empty relationships slice is returned properly",
			stored:   nil,
			expected: []types.Relationship{},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.stored {
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			relationships := suite.k.GetAllRelationships(suite.ctx)

			suite.Require().Len(relationships, len(test.expected))
			for _, rel := range relationships {
				suite.Require().Contains(test.expected, rel)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	tests := []struct {
		name     string
		stored   []types.Relationship
		user     string
		expected []types.Relationship
	}{
		{
			name: "Returns non empty relationships slice",
			stored: []types.Relationship{
				types.NewRelationship("user_1", "user_2", "subspace"),
				types.NewRelationship("user_2", "user_1", "subspace"),
			},
			user: "user_1",
			expected: []types.Relationship{
				types.NewRelationship("user_1", "user_2", "subspace"),
				types.NewRelationship("user_2", "user_1", "subspace"),
			},
		},
		{
			name:     "Returns empty relationships slice",
			stored:   nil,
			expected: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.stored {
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			relationships := suite.k.GetUserRelationships(suite.ctx, test.user)
			suite.Require().Equal(test.expected, relationships)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRelationship() {
	tests := []struct {
		name                 string
		stored               []types.Relationship
		relationshipToDelete types.Relationship
		expErr               bool
		expRelationships     []types.Relationship
	}{
		{
			name: "delete a relationship with len(relationships) > 1",
			stored: []types.Relationship{
				types.NewRelationship("user_1", "user_2", "subspace"),
				types.NewRelationship("user_2", "user_3", "subspace"),
				types.NewRelationship("user_1", "user_3", "subspace"),
			},
			relationshipToDelete: types.NewRelationship("user_1", "user_3", "subspace"),
			expErr:               false,
			expRelationships: []types.Relationship{
				types.NewRelationship("user_1", "user_2", "subspace"),
				types.NewRelationship("user_2", "user_3", "subspace"),
			},
		},
		{
			name: "delete a relationship with len(relationships) == 1",
			stored: []types.Relationship{
				types.NewRelationship("user_3", "user_2", "subspace"),
			},
			relationshipToDelete: types.NewRelationship("user_3", "user_2", "subspace"),
			expErr:               false,
		},
		{
			name:                 "deleting a non existing relationship returns an error",
			stored:               nil,
			relationshipToDelete: types.NewRelationship("user_3", "user_2", "subspace"),
			expErr:               true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.stored {
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			err := suite.k.RemoveRelationship(suite.ctx, test.relationshipToDelete)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				rel := suite.k.GetAllRelationships(suite.ctx)
				suite.Require().Equal(test.expRelationships, rel)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestKeeper_SaveUserBlock() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		userBlock        types.UserBlock
		expErr           bool
		expBlocks        []types.UserBlock
	}{
		{
			name: "already blocked user returns error",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock("user_1", "user_2", "reason", "subspace"),
			},
			userBlock: types.NewUserBlock("user_1", "user_2", "reason", "subspace"),
			expErr:    true,
		},
		{
			name:             "user block added correctly",
			storedUserBlocks: nil,
			userBlock:        types.NewUserBlock("user_1", "user_2", "reason", "subspace"),
			expErr:           false,
			expBlocks: []types.UserBlock{
				types.NewUserBlock("user_1", "user_2", "reason", "subspace"),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, block := range test.storedUserBlocks {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			err := suite.k.SaveUserBlock(suite.ctx, test.userBlock)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored := suite.k.GetAllUsersBlocks(suite.ctx)
				suite.Require().Equal(test.expBlocks, stored)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteUserBlock() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		data             struct {
			blocker  string
			blocked  string
			subspace string
		}
		expError  bool
		expBlocks []types.UserBlock
	}{
		{
			name: "delete user block with len(stored) > 1",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
			},
			expError: false,
		},
		{
			name: "delete user block with len(stored) == 1",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expError: false,
		},
		{
			name:             "deleting a user block that does not exist returns an error",
			storedUserBlocks: nil,
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expError: true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, block := range test.storedUserBlocks {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			err := suite.k.DeleteUserBlock(suite.ctx, test.data.blocker, test.data.blocked, test.data.subspace)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				blocks := suite.k.GetAllUsersBlocks(suite.ctx)
				suite.Require().Equal(test.expBlocks, blocks)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserBlocks() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		user             string
		expUserBlocks    []types.UserBlock
	}{
		{
			name: "non empty slice is returned properly",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			user: "blocker",
			expUserBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
		},
		{
			name:             "empty slice is returned properly",
			storedUserBlocks: nil,
			user:             "blocker",
			expUserBlocks:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, block := range test.storedUserBlocks {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			blocks := suite.k.GetUserBlocks(suite.ctx, test.user)
			suite.Require().Equal(test.expUserBlocks, blocks)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllUsersBlocks() {
	tests := []struct {
		name              string
		storedUsersBlocks []types.UserBlock
		expUsersBlocks    []types.UserBlock
	}{
		{
			name: "Returns a non-empty users blocks slice",
			storedUsersBlocks: []types.UserBlock{
				types.NewUserBlock("user_1", "user_2", "reason", "subspace_1"),
				types.NewUserBlock("user_1", "user_2", "reason", "subspace_2"),
				types.NewUserBlock("user_2", "user_1", "reason", "subspace_1"),
				types.NewUserBlock("user_2", "user_1", "reason", "subspace_2"),
			},
			expUsersBlocks: []types.UserBlock{
				types.NewUserBlock("user_1", "user_2", "reason", "subspace_1"),
				types.NewUserBlock("user_1", "user_2", "reason", "subspace_2"),
				types.NewUserBlock("user_2", "user_1", "reason", "subspace_1"),
				types.NewUserBlock("user_2", "user_1", "reason", "subspace_2"),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, userBlock := range test.storedUsersBlocks {
				err := suite.k.SaveUserBlock(suite.ctx, userBlock)
				suite.Require().NoError(err)
			}

			actualBlocks := suite.k.GetAllUsersBlocks(suite.ctx)

			suite.Require().Len(actualBlocks, len(test.expUsersBlocks))
			for _, block := range test.expUsersBlocks {
				suite.Require().Contains(actualBlocks, block)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasUserBlocked() {
	tests := []struct {
		name         string
		storedBlocks []types.UserBlock
		data         struct {
			blocker  string
			blocked  string
			subspace string
		}
		expBlocked bool
	}{
		{
			name: "blocked user found returns true",
			storedBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expBlocked: true,
		},
		{
			name: "blocked user not found returns false",
			storedBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace_2",
			},
			expBlocked: false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, block := range test.storedBlocks {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			blocked := suite.k.HasUserBlocked(suite.ctx, test.data.blocker, test.data.blocked, test.data.subspace)
			suite.Equal(test.expBlocked, blocked)
		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestKeeper_StoreLink() {
	link := types.NewLink(
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		types.NewProof("pubkey", "signature"),
		types.NewChainConfig("test-net", "cosmos"),
		time.Time{},
	)

	tests := []struct {
		name   string
		link   types.Link
		stored []types.Link
		expErr bool
	}{
		{
			name: "Non existent link is saved correctly",
			link: link,
			stored: []types.Link{
				types.NewLink(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewProof("pubkey", "signature"),
					types.NewChainConfig("test-net", "cosmos"),
					time.Time{},
				),
			},
			expErr: false,
		},
		{
			name: "already existent link returns error",
			link: link,
			stored: []types.Link{
				types.NewLink(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewProof("pubkey", "signature"),
					types.NewChainConfig("test-net", "cosmos"),
					time.Time{},
				),
				link,
			},
			expErr: true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, link := range test.stored {
				err := suite.k.StoreLink(suite.ctx, link)
				suite.Require().NoError(err)
			}
			err := suite.k.StoreLink(suite.ctx, test.link)

			if test.expErr {
				suite.Require().Error(err)
				suite.Require().Equal(fmt.Errorf("link already exists"), err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(append(test.stored, test.link), suite.k.GetAllLinks(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveLink() {
	link := types.NewLink(
		"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		types.NewProof("pubkey", "signature"),
		types.NewChainConfig("test-net", "cosmos"),
		time.Time{},
	)

	tests := []struct {
		name   string
		link   types.Link
		stored []types.Link
	}{
		{
			name: "Remove link successfully",
			link: link,
			stored: []types.Link{
				types.NewLink(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.NewProof("pubkey", "signature"),
					types.NewChainConfig("test-net", "cosmos"),
					time.Time{},
				),
				link,
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, link := range test.stored {
				err := suite.k.StoreLink(suite.ctx, link)
				suite.Require().NoError(err)
			}
			suite.k.RemoveLink(suite.ctx, test.link.ChainConfig.ID, link.Address)
			result := test.stored[:1]
			suite.Require().Equal(result, suite.k.GetAllLinks(suite.ctx))
		})
	}
}
