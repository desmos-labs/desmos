package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) Test_queryUserRelationships() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	expRelationships := types.NewRelationshipResponse([]sdk.AccAddress{addr1, addr2})

	tests := []struct {
		name          string
		path          []string
		relationships []sdk.AccAddress
		expResult     *types.RelationshipsResponse
		expErr        error
	}{
		{
			name:          "Invalid bech32 address returns error",
			path:          []string{types.QueryUserRelationships, "invalidAddress"},
			relationships: nil,
			expResult:     nil,
			expErr:        sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Invalid bech32 address: invalidAddress"),
		},
		{
			name:          "User Relationships returned correctly",
			path:          []string{types.QueryUserRelationships, suite.testData.user.String()},
			relationships: []sdk.AccAddress{addr1, addr2},
			expResult:     &expRelationships,
			expErr:        nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			for _, rel := range test.relationships {
				_ = suite.keeper.StoreRelationship(suite.ctx, suite.testData.user, rel)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)
				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
				suite.Nil(result)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryRelationships() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	tests := []struct {
		name          string
		path          []string
		relationships []sdk.AccAddress
		expResult     map[string][]sdk.AccAddress
		expErr        error
	}{
		{
			name:          "Relationships returned correctly",
			path:          []string{types.QueryRelationships},
			relationships: []sdk.AccAddress{addr1, addr2},
			expResult: map[string][]sdk.AccAddress{
				suite.testData.user.String():      {addr1, addr2},
				suite.testData.otherUser.String(): {addr1, addr2},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			for _, rel := range test.relationships {
				_ = suite.keeper.StoreRelationship(suite.ctx, suite.testData.user, rel)
				_ = suite.keeper.StoreRelationship(suite.ctx, suite.testData.otherUser, rel)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)
				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
				suite.Nil(result)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryUserBlocks() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

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
			path: []string{types.QueryUserBlocks, suite.testData.user.String()},
			userBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr1, "reason"),
				types.NewUserBlock(suite.testData.user, addr2, "reason"),
			},
			expResult: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr1, "reason"),
				types.NewUserBlock(suite.testData.user, addr2, "reason"),
			},
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			for _, ub := range test.userBlocks {
				_ = suite.keeper.SaveUserBlock(suite.ctx, ub)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expResult)
				suite.NoError(err)
				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
				suite.Nil(result)
			}
		})
	}
}
