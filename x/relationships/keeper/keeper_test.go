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
		{
			name: "relationship added correctly (another subspace)",
			storedRelationships: types.Relationships{
				types.NewRelationship(suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			user: suite.testData.user,
			relationship: types.NewRelationship(suite.testData.otherUser,
				"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88"),
			expErr: nil,
		},
		{
			name: "relationship added correctly (another receiver)",
			storedRelationships: types.Relationships{
				types.NewRelationship(suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			user: suite.testData.user,
			relationship: types.NewRelationship(suite.testData.user,
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

func (suite *KeeperTestSuite) TestKeeper_SaveUserBlock() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		userBlock        types.UserBlock
		expErr           error
	}{
		{
			name: "already blocked user returns error",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, suite.testData.otherUser, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			userBlock: types.NewUserBlock(suite.testData.user, suite.testData.otherUser, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expErr:    fmt.Errorf("the user with %s address has already been blocked", suite.testData.otherUser),
		},
		{
			name:             "user block added correctly",
			storedUserBlocks: nil,
			userBlock:        types.NewUserBlock(suite.testData.user, suite.testData.otherUser, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expErr:           nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedUserBlocks != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.UsersBlocksStoreKey(suite.testData.user), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedUserBlocks))
			}
			err := suite.keeper.SaveUserBlock(suite.ctx, test.userBlock)
			suite.Equal(test.expErr, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_UnblockUser() {
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	addr3, err := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	suite.NoError(err)

	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		expBlocks        []types.UserBlock
		expError         error
		userToUnblock    sdk.AccAddress
	}{
		{
			name: "Unblock user with len(storedUserBlocks) > 1",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewUserBlock(suite.testData.user, addr3, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr3, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			userToUnblock: addr2,
			expError:      nil,
		},
		{
			name:             "Unblock user with len(storedUserBlocks) == 1",
			storedUserBlocks: []types.UserBlock{types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
			expBlocks:        nil,
			userToUnblock:    addr2,
			expError:         nil,
		},
		{
			name:             "Delete a relationship with len(userBlocks) == 0",
			storedUserBlocks: nil,
			expBlocks:        nil,
			userToUnblock:    addr2,
			expError:         fmt.Errorf("blocked user with address %s not found", addr2),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedUserBlocks != nil {
				store.Set(types.UsersBlocksStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedUserBlocks))
			}

			err := suite.keeper.UnblockUser(suite.ctx, suite.testData.user, test.userToUnblock, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
			suite.Equal(test.expError, err)
			rel := suite.keeper.GetUserBlocks(suite.ctx, suite.testData.user)
			suite.Equal(test.expBlocks, rel)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserBlocks() {
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		expUserBlocks    []types.UserBlock
	}{
		{
			name: "Returns non empty user blocks slice",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expUserBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
		},
		{
			name:             "Returns empty user blocks slice",
			storedUserBlocks: nil,
			expUserBlocks:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedUserBlocks != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.UsersBlocksStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&[]types.UserBlock{
						types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					}))
			}

			suite.Equal(test.expUserBlocks, suite.keeper.GetUserBlocks(suite.ctx, suite.testData.user))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUsersBlocks() {
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	addr3, err := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	suite.NoError(err)

	tests := []struct {
		name              string
		storedUsersBlocks []types.UserBlock
		expUsersBlocks    []types.UserBlock
	}{
		{
			name: "Returns a non-empty users blocks slice",
			storedUsersBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewUserBlock(suite.testData.otherUser, addr3, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expUsersBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewUserBlock(suite.testData.otherUser, addr3, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, userBlock := range test.storedUsersBlocks {
				err := suite.keeper.SaveUserBlock(suite.ctx, userBlock)
				suite.NoError(err)
			}

			actualBlocks := suite.keeper.GetUsersBlocks(suite.ctx)
			suite.Equal(test.expUsersBlocks, actualBlocks)
		})
	}
}
