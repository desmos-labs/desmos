package keeper_test

import (
	"fmt"

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
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("Profile with address %s doesn't exists", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
			),
		},
		{
			name:          "Profile doesnt exist (blank path given)",
			path:          []string{types.QueryProfile, ""},
			storedAccount: suite.testData.profile,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				"DTag or address cannot be empty or blank",
			),
		},
		{
			name:          "Profile doesnt exist (dtag given)",
			path:          []string{types.QueryProfile, "monk"},
			storedAccount: suite.testData.profile,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "No address related to this dtag: monk"),
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
			suite.SetupTest() // reset
			err := suite.keeper.SaveProfile(suite.ctx, test.storedAccount)
			suite.Require().Nil(err)

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.storedAccount)
				suite.Require().NoError(err)
				suite.Require().Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
				suite.Require().Nil(result)
			}

		})
	}

}

func (suite *KeeperTestSuite) Test_queryProfiles() {

	tests := []struct {
		name          string
		path          []string
		storedAccount *types.Profile
		expResult     types.Profiles
	}{
		{
			name:          "Empty Profiles",
			path:          []string{types.QueryProfiles},
			storedAccount: nil,
			expResult:     types.Profiles{},
		},
		{
			name:          "Profile returned correctly",
			path:          []string{types.QueryProfiles},
			storedAccount: &suite.testData.profile,
			expResult:     types.Profiles{suite.testData.profile},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset

			if test.storedAccount != nil {
				err := suite.keeper.SaveProfile(suite.ctx, *test.storedAccount)
				suite.Require().Nil(err)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.Require().NoError(err)
				suite.Require().Equal(string(expectedIndented), string(result))
			}

		})
	}

}

func (suite *KeeperTestSuite) Test_queryParams() {
	validMin := sdk.NewInt(3)
	validMax := sdk.NewInt(30)

	nsParams := types.NewMonikerParams(validMin, validMax)
	monikerParams := types.NewDtagParams("^[A-Za-z0-9_]+$", validMin, validMax)

	tests := []struct {
		name                string
		path                []string
		nsParamsStored      types.MonikerParams
		monikerParamsStored types.DtagParams
		bioParamStored      sdk.Int
		expResult           types.Params
	}{
		{
			name:                "Returning profile parameters correctly",
			path:                []string{types.QueryParams},
			nsParamsStored:      nsParams,
			monikerParamsStored: monikerParams,
			bioParamStored:      validMax,
			expResult:           types.NewParams(nsParams, monikerParams, validMax),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			suite.keeper.SetParams(suite.ctx, types.NewParams(test.nsParamsStored, test.monikerParamsStored, test.bioParamStored))
			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if result != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
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
			path:           []string{types.QueryDTagRequests, "invalid"},
			storedRequests: nil,
			expResult:      nil,
			expErr:         sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid bech32 address: invalid"),
		},
		{
			name:           "Empty dTag requests returns correctly",
			path:           []string{types.QueryDTagRequests, suite.testData.user.String()},
			storedRequests: nil,
			expResult:      nil,
			expErr:         nil,
		},
		{
			name: "Stored dTag requests returns correctly",
			path: []string{types.QueryDTagRequests, suite.testData.user.String()},
			storedRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			expResult: []types.DTagTransferRequest{
				types.NewDTagTransferRequest("dtag", suite.testData.user, suite.testData.otherUser),
			},
			expErr: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest() // reset
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedRequests != nil {
				store.Set(types.DtagTransferRequestStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRequests),
				)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.Require().NoError(err)
				suite.Require().Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Require().Equal(test.expErr.Error(), err.Error())
				suite.Require().Nil(result)
			}
		})
	}

}
