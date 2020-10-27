package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (suite *KeeperTestSuite) Test_queryUserRelationships() {
	tests := []struct {
		name          string
		path          []string
		relationships []types.Relationship
		expResult     []types.Relationship
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
			path: []string{types.QueryUserRelationships, suite.testData.user},
			relationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expResult: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
			for _, rel := range test.relationships {
				err := suite.keeper.StoreRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			querier := keeper.NewQuerier(suite.keeper, suite.legacyAmino)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.legacyAmino, &test.expResult)
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

func (suite *KeeperTestSuite) Test_queryRelationships() {
	tests := []struct {
		name          string
		path          []string
		relationships []types.Relationship
		expResult     []types.Relationship
		expErr        error
	}{
		{
			name: "Relationships returned correctly",
			path: []string{types.QueryRelationships},
			relationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest() // reset
		suite.Run(test.name, func() {
			for _, rel := range test.relationships {
				err := suite.keeper.StoreRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			querier := keeper.NewQuerier(suite.keeper, suite.legacyAmino)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.legacyAmino, &test.expResult)
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
				_ = suite.keeper.SaveUserBlock(suite.ctx, ub)
			}

			querier := keeper.NewQuerier(suite.keeper, suite.legacyAmino)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expResult != nil {
				suite.Require().Nil(err)
				expectedIndented, err := codec.MarshalJSONIndent(suite.legacyAmino, &test.expResult)
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
