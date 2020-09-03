package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

func (suite *KeeperTestSuite) TestKeeper_StoreRelationship() {
	tests := []struct {
		name                string
		storedRelationships types.Relationships
		user                sdk.AccAddress
		relationship        types.Relationship
		expErr              error
	}{
		{
			name: "already existent relationship returns error",
			storedRelationships: types.Relationships{
				types.NewRelationship(suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			user: suite.testData.user,
			relationship: types.NewRelationship(suite.testData.otherUser,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expErr: fmt.Errorf("relationship already exists with %s", suite.testData.otherUser),
		},
		{
			name:                "relationship added correctly",
			storedRelationships: nil,
			user:                suite.testData.user,
			relationship: types.NewRelationship(suite.testData.otherUser,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expErr: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationships != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.RelationshipsStoreKey(test.user), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))
			}
			err := suite.keeper.StoreRelationship(suite.ctx, test.user, test.relationship)
			suite.Equal(test.expErr, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUsersRelationships() {
	tests := []struct {
		name                string
		storedRelationships types.Relationships
		expMap              map[string]types.Relationships
	}{
		{
			name: "Return a non-empty address -> relationships map",
			storedRelationships: types.Relationships{
				types.NewRelationship(suite.testData.user,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewRelationship(suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expMap: map[string]types.Relationships{
				suite.testData.user.String(): {
					types.NewRelationship(suite.testData.otherUser,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				},
				suite.testData.otherUser.String(): {
					types.NewRelationship(suite.testData.user,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				},
			},
		},
		{
			name:                "Return an empty address -> relationships map",
			storedRelationships: nil,
			expMap:              map[string]types.Relationships{},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationships != nil {
				_ = suite.keeper.StoreRelationship(suite.ctx, suite.testData.user, test.storedRelationships[1])
				_ = suite.keeper.StoreRelationship(suite.ctx, suite.testData.otherUser, test.storedRelationships[0])
			}

			actualIDsMap := suite.keeper.GetUsersRelationships(suite.ctx)
			suite.Equal(test.expMap, actualIDsMap)
		})
	}

}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	tests := []struct {
		name                string
		storedRelationships types.Relationships
		expRelationships    types.Relationships
	}{
		{
			name: "Returns non empty relationships slice",
			storedRelationships: types.Relationships{
				types.NewRelationship(addr1,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewRelationship(addr2,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expRelationships: types.Relationships{
				types.NewRelationship(addr1,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewRelationship(addr2,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
		},
		{
			name:                "Returns empty relationships slice",
			storedRelationships: nil,
			expRelationships:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationships != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.RelationshipsStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))
			}

			suite.Equal(test.expRelationships, suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRelationship() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	addr3, err := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	suite.NoError(err)

	tests := []struct {
		name                 string
		storedRelationships  types.Relationships
		expRelationships     types.Relationships
		relationshipToDelete types.Relationship
	}{
		{
			name: "Delete a relationship with len(relationships) > 1",
			storedRelationships: types.Relationships{
				types.NewRelationship(addr1,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewRelationship(addr2,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewRelationship(addr3,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expRelationships: types.Relationships{
				types.NewRelationship(addr1,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewRelationship(addr3,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			relationshipToDelete: types.NewRelationship(addr2,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
		},
		{
			name: "Delete a relationship with len(relationships) == 1",
			storedRelationships: types.Relationships{
				types.NewRelationship(addr1,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expRelationships: nil,
			relationshipToDelete: types.NewRelationship(addr1,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
		},
		{
			name:                "Delete a relationship with len(relationships) == 0",
			storedRelationships: nil,
			expRelationships:    nil,
			relationshipToDelete: types.NewRelationship(addr1,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedRelationships != nil {
				store.Set(types.RelationshipsStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))
			}

			suite.keeper.DeleteRelationship(suite.ctx, suite.testData.user, test.relationshipToDelete)
			rel := suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user)
			suite.Equal(test.expRelationships, rel)
		})
	}
}
