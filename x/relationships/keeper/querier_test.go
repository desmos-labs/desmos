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

	expRelationships := types.Relationships{
		types.NewRelationship(addr1, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
		types.NewRelationship(addr2, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
	}

	tests := []struct {
		name          string
		path          []string
		relationships types.Relationships
		expResult     types.Relationships
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
			name: "User Relationships returned correctly",
			path: []string{types.QueryUserRelationships, suite.testData.user.String()},
			relationships: types.Relationships{
				types.NewRelationship(addr1, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewRelationship(addr2, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expResult: expRelationships,
			expErr:    nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest() // reset
		suite.Run(test.name, func() {
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
		relationships types.Relationships
		expResult     map[string]types.Relationships
		expErr        error
	}{
		{
			name: "Relationships returned correctly",
			path: []string{types.QueryRelationships},
			relationships: types.Relationships{
				types.NewRelationship(addr1, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewRelationship(addr2, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expResult: map[string]types.Relationships{
				suite.testData.user.String(): {
					types.NewRelationship(addr1, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					types.NewRelationship(addr2, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				},
				suite.testData.otherUser.String(): {
					types.NewRelationship(addr1, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					types.NewRelationship(addr2, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest() // reset
		suite.Run(test.name, func() {
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
