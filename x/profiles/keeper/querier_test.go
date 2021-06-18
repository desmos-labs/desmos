package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_queryProfile() {
	tests := []struct {
		name          string
		path          []string
		storedAccount *types.Profile
		expErr        error
	}{
		{
			name:          "Profile doesnt exist (address given)",
			path:          []string{types.QueryProfile, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			storedAccount: suite.testData.profile.Profile,
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"Profile with address cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 doesn't exists",
			),
		},
		{
			name:          "Profile doesnt exist (blank path given)",
			path:          []string{types.QueryProfile, ""},
			storedAccount: suite.testData.profile.Profile,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "DTag or address cannot be empty or blank"),
		},
		{
			name:          "Profile doesnt exist (DTag given)",
			path:          []string{types.QueryProfile, "monk"},
			storedAccount: suite.testData.profile.Profile,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No address related to this DTag: monk"),
		},
		{
			name:          "Profile returned correctly (address given)",
			path:          []string{types.QueryProfile, suite.testData.profile.GetAddress().String()},
			storedAccount: suite.testData.profile.Profile,
			expErr:        nil,
		},
		{
			name:          "Profile returned correctly (dtag given)",
			path:          []string{types.QueryProfile, "dtag"},
			storedAccount: suite.testData.profile.Profile,
			expErr:        nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			err := suite.k.StoreProfile(suite.ctx, test.storedAccount)
			suite.Require().Nil(err)

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.legacyAminoCdc, &test.storedAccount)
				suite.Require().NoError(err)
				suite.Require().Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
				suite.Require().Nil(result)
			}
		})
	}

}

func (suite *KeeperTestSuite) Test_queryParams() {
	tests := []struct {
		name                 string
		path                 []string
		nicknameParamsStored types.NicknameParams
		dTagParamsStored     types.DTagParams
		bioParamStored       sdk.Int
		expResult            types.Params
	}{
		{
			name:                 "Returning profile parameters correctly",
			path:                 []string{types.QueryParams},
			nicknameParamsStored: types.NewNicknameParams(sdk.NewInt(3), sdk.NewInt(30)),
			dTagParamsStored:     types.NewDTagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(30)),
			bioParamStored:       sdk.NewInt(30),
			expResult: types.NewParams(
				types.NewNicknameParams(sdk.NewInt(3), sdk.NewInt(30)),
				types.NewDTagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(30)),
				sdk.NewInt(30),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, types.NewParams(test.nicknameParamsStored, test.dTagParamsStored, test.bioParamStored))

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)

				suite.Require().NoError(err)
				suite.Require().Equal(string(expectedIndented), string(result))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryDTagRequests() {
	tests := []struct {
		name           string
		path           []string
		storedRequests []types.DTagTransferRequest
		expResult      []types.DTagTransferRequest
		shouldErr      bool
	}{
		{
			name:           "Invalid address returns error",
			path:           []string{types.QueryIncomingDTagRequests, "invalid"},
			storedRequests: nil,
			expResult:      nil,
			shouldErr:      true,
		},
		{
			name:           "Empty DTag requests returns correctly",
			path:           []string{types.QueryIncomingDTagRequests, suite.testData.user},
			storedRequests: nil,
			expResult:      nil,
			shouldErr:      false,
		},
		{
			name: "Stored dTag requests returns correctly",
			storedRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			path: []string{types.QueryIncomingDTagRequests, suite.testData.otherUser},
			expResult: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			shouldErr: false,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, req := range test.storedRequests {
				err := suite.k.SaveDTagTransferRequest(suite.ctx, req)
				suite.Require().NoError(err)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				var requests []types.DTagTransferRequest
				err = suite.legacyAminoCdc.UnmarshalJSON(result, &requests)
				suite.Require().NoError(err)
				suite.Require().Equal(test.expResult, requests)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryUserRelationships() {
	tests := []struct {
		name          string
		path          []string
		relationships []types.Relationship
		expErr        error
		expResult     []types.Relationship
	}{
		{
			name:          "Invalid bech32 address returns error",
			path:          []string{types.QueryUserRelationships, "invalidAddress"},
			relationships: nil,
			expResult:     nil,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid bech32 address: invalidAddress"),
		},
		{
			name: "User relationships returned correctly",
			path: []string{types.QueryUserRelationships, suite.testData.user},
			relationships: []types.Relationship{
				types.NewRelationship(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expResult: []types.Relationship{
				types.NewRelationship(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.relationships {
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Require().Nil(err)

				var actual []types.Relationship
				err := suite.legacyAminoCdc.UnmarshalJSON(result, &actual)
				suite.Require().NoError(err)

				suite.Require().Len(actual, len(test.expResult))
				for _, relationship := range actual {
					suite.Require().Contains(test.expResult, relationship)
				}
			}

			if result == nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
				suite.Require().Nil(result)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryUserBlocks() {
	tests := []struct {
		name       string
		path       []string
		userBlocks []types.UserBlock
		expResult  []types.UserBlock
		expErr     error
	}{
		{
			name:       "Invalid bech32 address returns error",
			path:       []string{types.QueryUserBlocks, "invalidAddress"},
			userBlocks: nil,
			expResult:  nil,
			expErr:     sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid bech32 address: invalidAddress"),
		},
		{
			name: "User Relationships returned correctly",
			path: []string{types.QueryUserBlocks, suite.testData.user},
			userBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expResult: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest() // reset
		suite.Run(test.name, func() {
			for _, ub := range test.userBlocks {
				_ = suite.k.SaveUserBlock(suite.ctx, ub)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
				suite.Require().NoError(err)
				suite.Require().Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.Require().Error(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
				suite.Require().Nil(result)
			}
		})
	}
}
