package keeper_test

import (
	"fmt"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

func (suite *KeeperTestSuite) TestKeeper_StoreRelationship() {
	tests := []struct {
		name         string
		stored       []types.Relationship
		user         string
		relationship types.Relationship
		expErr       error
	}{
		{
			name: "already existent relationship returns error",
			stored: []types.Relationship{
				types.NewRelationship(
					suite.testData.user,
					suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			user: suite.testData.user,
			relationship: types.NewRelationship(
				suite.testData.user,
				suite.testData.otherUser,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: fmt.Errorf("relationship already exists with %s", suite.testData.otherUser),
		},
		{
			name:   "relationship added correctly",
			stored: nil,
			user:   suite.testData.user,
			relationship: types.NewRelationship(
				suite.testData.user,
				suite.testData.otherUser,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: nil,
		},
		{
			name: "relationship added correctly (another subspace)",
			stored: []types.Relationship{
				types.NewRelationship(
					suite.testData.user,
					suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			user: suite.testData.user,
			relationship: types.NewRelationship(
				suite.testData.user,
				suite.testData.otherUser,
				"2bdf5932925584b9a86470bea60adce69041608a447f84a3317723aa5678ec88",
			),
			expErr: nil,
		},
		{
			name: "relationship added correctly (another receiver)",
			stored: []types.Relationship{
				types.NewRelationship(
					suite.testData.user,
					suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			user: suite.testData.user,
			relationship: types.NewRelationship(
				suite.testData.user,
				suite.testData.user,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.stored != nil {
				store := suite.ctx.KVStore(suite.storeKey)
				bz, err := suite.keeper.MarshalRelationships(test.stored)
				suite.Require().NoError(err)

				store.Set(types.RelationshipsStoreKey(test.user), bz)
			}
			err := suite.keeper.StoreRelationship(suite.ctx, test.relationship)
			suite.Require().Equal(test.expErr, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUsersRelationships() {
	tests := []struct {
		name     string
		stored   []types.Relationship
		expected []types.Relationship
	}{
		{
			name: "Return a non-empty address -> relationships map",
			stored: []types.Relationship{
				types.NewRelationship(
					suite.testData.user,
					suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					suite.testData.otherUser,
					suite.testData.user,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expected: []types.Relationship{
				types.NewRelationship(
					suite.testData.user,
					suite.testData.otherUser,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					suite.testData.otherUser,
					suite.testData.user,
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
		{
			name:     "Return an empty address -> relationships map",
			stored:   nil,
			expected: []types.Relationship{},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.stored {
				err := suite.keeper.StoreRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			relationships, err := suite.keeper.GetAllRelationships(suite.ctx)
			suite.Require().NoError(err)

			suite.Require().Equal(test.expected, relationships)
		})
	}

}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	tests := []struct {
		name     string
		stored   []types.Relationship
		expected []types.Relationship
	}{
		{
			name: "Returns non empty relationships slice",
			stored: []types.Relationship{
				types.NewRelationship(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expected: []types.Relationship{
				types.NewRelationship(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
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
			if test.stored != nil {
				store := suite.ctx.KVStore(suite.storeKey)
				bz, err := suite.keeper.MarshalRelationships(test.stored)
				suite.Require().NoError(err)

				store.Set(types.RelationshipsStoreKey(suite.testData.user), bz)
			}

			relationships, err := suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user)
			suite.Require().NoError(err)

			suite.Require().Equal(test.expected, relationships)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRelationship() {
	tests := []struct {
		name                 string
		stored               []types.Relationship
		expRelationships     []types.Relationship
		relationshipToDelete types.Relationship
	}{
		{
			name: "Delete a relationship with len(relationships) > 1",
			stored: []types.Relationship{
				types.NewRelationship(
					"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			relationshipToDelete: types.NewRelationship(
				"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
		{
			name: "Delete a relationship with len(relationships) == 1",
			stored: []types.Relationship{
				types.NewRelationship(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			relationshipToDelete: types.NewRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expRelationships: nil,
		},
		{
			name:   "Delete a relationship with len(relationships) == 0",
			stored: nil,
			relationshipToDelete: types.NewRelationship(
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expRelationships: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			if test.stored != nil {
				bz, err := suite.keeper.MarshalRelationships(test.stored)
				suite.Require().NoError(err)

				store.Set(types.RelationshipsStoreKey(suite.testData.user), bz)
			}

			err := suite.keeper.DeleteRelationship(suite.ctx, test.relationshipToDelete)
			suite.Require().NoError(err)

			rel, err := suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user)
			suite.Require().NoError(err)

			suite.Require().Equal(test.expRelationships, rel)
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
				store := suite.ctx.KVStore(suite.storeKey)
				bz, err := suite.keeper.MarshalUserBlocks(test.storedUserBlocks)
				suite.Require().NoError(err)

				store.Set(types.UsersBlocksStoreKey(suite.testData.user), bz)
			}
			err := suite.keeper.SaveUserBlock(suite.ctx, test.userBlock)
			suite.Require().Equal(test.expErr, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_UnblockUser() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		userToUnblock    string
		expBlocks        []types.UserBlock
		expError         error
	}{
		{
			name: "Unblock user with len(storedUserBlocks) > 1",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewUserBlock(
					suite.testData.user,
					"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			userToUnblock: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expError: nil,
		},
		{
			name: "Unblock user with len(storedUserBlocks) == 1",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			userToUnblock: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expBlocks:     nil,
			expError:      nil,
		},
		{
			name:             "Delete a relationship with len(userBlocks) == 0",
			storedUserBlocks: nil,
			userToUnblock:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expBlocks:        nil,
			expError:         fmt.Errorf("blocked user with address %s not found", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			if test.storedUserBlocks != nil {
				bz, err := suite.keeper.MarshalUserBlocks(test.storedUserBlocks)
				suite.Require().NoError(err)

				store.Set(types.UsersBlocksStoreKey(suite.testData.user), bz)
			}

			err := suite.keeper.DeleteUserBlock(
				suite.ctx,
				suite.testData.user,
				test.userToUnblock,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			)
			suite.Require().Equal(test.expError, err)

			blocks, err := suite.keeper.GetUserBlocks(suite.ctx, suite.testData.user)
			suite.Require().NoError(err)
			suite.Require().Equal(test.expBlocks, blocks)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserBlocks() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		expUserBlocks    []types.UserBlock
	}{
		{
			name: "Returns non empty user blocks slice",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expUserBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
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
				store := suite.ctx.KVStore(suite.storeKey)
				bz, err := suite.keeper.MarshalUserBlocks([]types.UserBlock{
					types.NewUserBlock(
						suite.testData.user,
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				})
				suite.Require().NoError(err)

				store.Set(types.UsersBlocksStoreKey(suite.testData.user), bz)
			}

			blocks, err := suite.keeper.GetUserBlocks(suite.ctx, suite.testData.user)
			suite.Require().NoError(err)

			suite.Require().Equal(test.expUserBlocks, blocks)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUsersBlocks() {
	tests := []struct {
		name              string
		storedUsersBlocks []types.UserBlock
		expUsersBlocks    []types.UserBlock
	}{
		{
			name: "Returns a non-empty users blocks slice",
			storedUsersBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewUserBlock(
					suite.testData.otherUser,
					"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expUsersBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewUserBlock(
					suite.testData.otherUser,
					"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, userBlock := range test.storedUsersBlocks {
				err := suite.keeper.SaveUserBlock(suite.ctx, userBlock)
				suite.Require().NoError(err)
			}

			actualBlocks, err := suite.keeper.GetUsersBlocks(suite.ctx)
			suite.Require().NoError(err)

			suite.Require().Equal(test.expUsersBlocks, actualBlocks)
		})
	}
}

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

			for _, block := range test.userBlocks {
				err := suite.keeper.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			res := suite.keeper.IsUserBlocked(suite.ctx, test.blocker, test.blocked)
			suite.Equal(test.expBool, res)
		})
	}
}
