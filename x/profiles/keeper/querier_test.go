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
		storedAccount types.Profile
		expErr        error
	}{
		{
			name:          "Profile doesnt exist (address given)",
			path:          []string{types.QueryProfile, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			storedAccount: suite.testData.profile,
			expErr: sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest,
				"Profile with address cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47 doesn't exists",
			),
		},
		{
			name:          "Profile doesnt exist (blank path given)",
			path:          []string{types.QueryProfile, ""},
			storedAccount: suite.testData.profile,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "DTag or address cannot be empty or blank"),
		},
		{
			name:          "Profile doesnt exist (DTag given)",
			path:          []string{types.QueryProfile, "monk"},
			storedAccount: suite.testData.profile,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No address related to this DTag: monk"),
		},
		{
			name:          "Profile returned correctly (address given)",
			path:          []string{types.QueryProfile, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			storedAccount: suite.testData.profile,
			expErr:        nil,
		},
		{
			name:          "Profile returned correctly (dtag given)",
			path:          []string{types.QueryProfile, "dtag"},
			storedAccount: suite.testData.profile,
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
		name                string
		path                []string
		nsParamsStored      types.MonikerParams
		monikerParamsStored types.DTagParams
		bioParamStored      sdk.Int
		expResult           types.Params
	}{
		{
			name:                "Returning profile parameters correctly",
			path:                []string{types.QueryParams},
			nsParamsStored:      types.NewMonikerParams(sdk.NewInt(3), sdk.NewInt(30)),
			monikerParamsStored: types.NewDtagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(30)),
			bioParamStored:      sdk.NewInt(30),
			expResult: types.NewParams(
				types.NewMonikerParams(sdk.NewInt(3), sdk.NewInt(30)),
				types.NewDtagParams("^[A-Za-z0-9_]+$", sdk.NewInt(3), sdk.NewInt(30)),
				sdk.NewInt(30),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, types.NewParams(test.nsParamsStored, test.monikerParamsStored, test.bioParamStored))

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
		expErr         error
	}{
		{
			name:           "Invalid address returns error",
			path:           []string{types.QueryIncomingDTagRequests, "invalid"},
			storedRequests: nil,
			expResult:      nil,
			expErr:         sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid bech32 address: invalid"),
		},
		{
			name:           "Empty DTag requests returns correctly",
			path:           []string{types.QueryIncomingDTagRequests, suite.testData.user},
			storedRequests: nil,
			expResult:      nil,
			expErr:         nil,
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
			expErr: nil,
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
			suite.RequireErrorsEqual(test.expErr, err)

			if test.expErr == nil {
				var requests []types.DTagTransferRequest
				err = suite.legacyAminoCdc.UnmarshalJSON(result, &requests)
				suite.Require().NoError(err)
				suite.Require().Equal(test.expResult, requests)
			}
		})
	}
}
