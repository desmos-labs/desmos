package keeper_test

import (
	"fmt"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	relationshipstypes "github.com/desmos-labs/desmos/x/relationships/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_IsUserBlocked() {
	tests := []struct {
		name       string
		blocker    string
		blocked    string
		userBlocks []relationshipstypes.UserBlock
		expBool    bool
	}{
		{
			name:    "blocked user found returns true",
			blocker: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			blocked: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			userBlocks: []relationshipstypes.UserBlock{
				relationshipstypes.NewUserBlock(
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
				_ = suite.rk.SaveUserBlock(suite.ctx, test.userBlocks[0])
			}
			res := suite.k.IsUserBlocked(suite.ctx, test.blocker, test.blocked)
			suite.Equal(test.expBool, res)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDtagFromAddress() {
	tests := []struct {
		name           string
		storedProfiles []types.Profile
		address        string
		expDTag        string
	}{
		{
			name: "found right dtag",
			storedProfiles: []types.Profile{
				suite.testData.profile,
			},
			address: suite.testData.profile.Creator,
			expDTag: suite.testData.profile.Dtag,
		},
		{
			name: "no dtag found",
			storedProfiles: []types.Profile{
				suite.testData.profile,
			},
			address: "non_existent",
			expDTag: "",
		},
	}

	for _, test := range tests {
		suite.SetupTest() //reset
		test := test
		suite.Run(test.name, func() {
			for _, profile := range test.storedProfiles {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			dTag := suite.k.GetDtagFromAddress(suite.ctx, test.address)
			suite.Require().Equal(test.expDTag, dTag)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_StoreProfile() {
	tests := []struct {
		name           string
		account        types.Profile
		storedProfiles []types.Profile
		expError       error
	}{
		{
			name:           "Non existent Profile saved correctly",
			account:        suite.testData.profile,
			storedProfiles: nil,
			expError:       nil,
		},
		{
			name: "Existent account with different creator returns error",
			account: types.Profile{
				Dtag:     suite.testData.profile.Dtag,
				Bio:      suite.testData.profile.Bio,
				Pictures: suite.testData.profile.Pictures,
				Creator:  "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			storedProfiles: []types.Profile{suite.testData.profile},
			expError: sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"a profile with dtag dtag has already been created"),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, profile := range test.storedProfiles {
				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)
			}

			err := suite.k.StoreProfile(suite.ctx, test.account)
			suite.RequireErrorsEqual(test.expError, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetProfile() {
	tests := []struct {
		name           string
		storedProfiles []types.Profile
		address        string
		expFound       bool
		expProfile     *types.Profile
	}{
		{
			name: "Profile founded",
			storedProfiles: []types.Profile{
				suite.testData.profile,
			},
			address:    suite.testData.profile.Creator,
			expFound:   true,
			expProfile: &suite.testData.profile,
		},
		{
			name:           "Profile not found",
			storedProfiles: []types.Profile{},
			address:        suite.testData.profile.Creator,
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

			res, found := suite.k.GetProfile(suite.ctx, test.address)
			suite.Require().Equal(test.expFound, found)

			if found {
				suite.Require().True(res.Equal(test.expProfile))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_RemoveProfile() {
	err := suite.k.StoreProfile(suite.ctx, suite.testData.profile)
	suite.Require().Nil(err)

	_, found := suite.k.GetProfile(suite.ctx, suite.testData.profile.Creator)
	suite.True(found)

	err = suite.k.RemoveProfile(suite.ctx, suite.testData.profile.Creator)
	suite.Require().NoError(err)

	_, found = suite.k.GetProfile(suite.ctx, suite.testData.profile.Creator)
	suite.Require().False(found)
}

func (suite *KeeperTestSuite) TestKeeper_GetProfiles() {
	tests := []struct {
		name     string
		accounts []types.Profile
	}{
		{
			name:     "Non empty Profiles list returned",
			accounts: []types.Profile{suite.testData.profile},
		},
		{
			name:     "Profile not found",
			accounts: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if len(test.accounts) != 0 {
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ProfileStoreKey(test.accounts[0].Creator)
				store.Set(key, suite.cdc.MustMarshalBinaryBare(&test.accounts[0]))
			}

			res := suite.k.GetProfiles(suite.ctx)
			suite.Require().Equal(test.accounts, res)
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
		expErr                error
		expStoredTransferReqs []types.DTagTransferRequest
	}{
		{
			name: "Already present request returns error",
			storedTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			transferReq: types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			expErr: sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the transfer request from %s to %s has already been made",
				suite.testData.user, suite.testData.otherUser,
			),
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
			expErr:      nil,
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
			expErr:      nil,
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
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"the transfer request from %s to %s has already been made",
				suite.testData.user, suite.testData.otherUser,
			),
			expStoredTransferReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
		},
		{
			name:               "Not already present request is saved correctly",
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
			for _, req := range test.storedTransferReqs {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			err := suite.k.SaveDTagTransferRequest(suite.ctx, test.transferReq)
			suite.RequireErrorsEqual(test.expErr, err)

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
				reqs := keeper.NewWrappedDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DtagTransferRequestStoreKey(suite.testData.user),
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
				reqs := keeper.NewWrappedDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DtagTransferRequestStoreKey(suite.testData.user),
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
				reqs := keeper.NewWrappedDTagTransferRequests(test.storedReqs)
				store.Set(
					types.DtagTransferRequestStoreKey(suite.testData.user),
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
		expErr          error
		storedReqsAfter []types.DTagTransferRequest
	}{
		{
			name:       "Empty requests array returns error",
			storedReqs: nil,
			sender:     suite.testData.user,
			receiver:   suite.testData.otherUser,
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"request from %s to %s not found",
				suite.testData.user,
				suite.testData.otherUser,
			),
		},
		{
			name: "Deleting non existent request returns an error and doesn't change the store",
			storedReqs: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.user),
			},
			sender:   suite.testData.user,
			receiver: suite.testData.otherUser,
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"request from %s to %s not found",
				suite.testData.user,
				suite.testData.otherUser,
			),
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
			suite.RequireErrorsEqual(test.expErr, err)

			reqs := suite.k.GetDTagTransferRequests(suite.ctx)
			suite.Require().Equal(test.storedReqsAfter, reqs)
		})
	}
}
