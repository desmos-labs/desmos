package keeper_test

import (
	"github.com/desmos-labs/desmos/x/relationships/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveRelationship() {
	tests := []struct {
		name             string
		stored           []types.Relationship
		user             string
		relationship     types.Relationship
		expErr           bool
		expRelationships []types.Relationship
	}{
		{
			name: "already existent relationship returns error",
			stored: []types.Relationship{
				types.NewRelationship("user", "recipient", "subspace"),
			},
			user:         "user",
			relationship: types.NewRelationship("user", "recipient", "subspace"),
			expErr:       true,
		},
		{
			name:         "relationship added correctly",
			stored:       nil,
			user:         "user",
			relationship: types.NewRelationship("user", "recipient", "subspace"),
			expErr:       false,
			expRelationships: []types.Relationship{
				types.NewRelationship("user", "recipient", "subspace"),
			},
		},
		{
			name: "relationship added correctly (another subspace)",
			stored: []types.Relationship{
				types.NewRelationship("user", "recipient", "subspace"),
			},
			user:         "user",
			relationship: types.NewRelationship("user", "recipient", "subspace_2"),
			expErr:       false,
		},
		{
			name: "relationship added correctly (another receiver)",
			stored: []types.Relationship{
				types.NewRelationship("user", "recipient", "subspace"),
			},
			user:         "user",
			relationship: types.NewRelationship("user", "user", "subspace"),
			expErr:       false,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, relationship := range test.stored {
				err := suite.keeper.SaveRelationship(suite.ctx, relationship)
				suite.Require().NoError(err)
			}

			err := suite.keeper.SaveRelationship(suite.ctx, test.relationship)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllRelationships() {
	tests := []struct {
		name     string
		stored   []types.Relationship
		expected []types.Relationship
	}{
		{
			name: "non empty relationships slice is returned properly",
			stored: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
				types.NewRelationship("creator", "another_recipient", "subspace"),
				types.NewRelationship("recipient", "creator", "subspace"),
				types.NewRelationship("recipient", "creator", "subspace_2"),
			},
			expected: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
				types.NewRelationship("creator", "another_recipient", "subspace"),
				types.NewRelationship("recipient", "creator", "subspace"),
				types.NewRelationship("recipient", "creator", "subspace_2"),
			},
		},
		{
			name:     "empty relationships slice is returned properly",
			stored:   nil,
			expected: []types.Relationship{},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.stored {
				err := suite.keeper.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			relationships := suite.keeper.GetAllRelationships(suite.ctx)

			suite.Require().Len(relationships, len(test.expected))
			for _, rel := range relationships {
				suite.Require().Contains(test.expected, rel)
			}
		})
	}

}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	tests := []struct {
		name     string
		stored   []types.Relationship
		user     string
		expected []types.Relationship
	}{
		{
			name: "Returns non empty relationships slice",
			stored: []types.Relationship{
				types.NewRelationship("user_1", "user_2", "subspace"),
				types.NewRelationship("user_2", "user_1", "subspace"),
			},
			user: "user_1",
			expected: []types.Relationship{
				types.NewRelationship("user_1", "user_2", "subspace"),
				types.NewRelationship("user_2", "user_1", "subspace"),
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
			for _, rel := range test.stored {
				err := suite.keeper.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			relationships := suite.keeper.GetUserRelationships(suite.ctx, test.user)
			suite.Require().Equal(test.expected, relationships)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRelationship() {
	tests := []struct {
		name                 string
		stored               []types.Relationship
		relationshipToDelete types.Relationship
		expErr               bool
		expRelationships     []types.Relationship
	}{
		{
			name: "delete a relationship with len(relationships) > 1",
			stored: []types.Relationship{
				types.NewRelationship("user_1", "user_2", "subspace"),
				types.NewRelationship("user_2", "user_3", "subspace"),
				types.NewRelationship("user_1", "user_3", "subspace"),
			},
			relationshipToDelete: types.NewRelationship("user_1", "user_3", "subspace"),
			expErr:               false,
			expRelationships: []types.Relationship{
				types.NewRelationship("user_1", "user_2", "subspace"),
				types.NewRelationship("user_2", "user_3", "subspace"),
			},
		},
		{
			name: "delete a relationship with len(relationships) == 1",
			stored: []types.Relationship{
				types.NewRelationship("user_3", "user_2", "subspace"),
			},
			relationshipToDelete: types.NewRelationship("user_3", "user_2", "subspace"),
			expErr:               false,
		},
		{
			name:                 "deleting a non existing relationship returns an error",
			stored:               nil,
			relationshipToDelete: types.NewRelationship("user_3", "user_2", "subspace"),
			expErr:               true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.stored {
				err := suite.keeper.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			err := suite.keeper.RemoveRelationship(suite.ctx, test.relationshipToDelete)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				rel := suite.keeper.GetAllRelationships(suite.ctx)
				suite.Require().Equal(test.expRelationships, rel)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestKeeper_SaveUserBlock() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		userBlock        types.UserBlock
		expErr           bool
		expBlocks        []types.UserBlock
	}{
		{
			name: "already blocked user returns error",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock("user_1", "user_2", "reason", "subspace"),
			},
			userBlock: types.NewUserBlock("user_1", "user_2", "reason", "subspace"),
			expErr:    true,
		},
		{
			name:             "user block added correctly",
			storedUserBlocks: nil,
			userBlock:        types.NewUserBlock("user_1", "user_2", "reason", "subspace"),
			expErr:           false,
			expBlocks: []types.UserBlock{
				types.NewUserBlock("user_1", "user_2", "reason", "subspace"),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, block := range test.storedUserBlocks {
				err := suite.keeper.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			err := suite.keeper.SaveUserBlock(suite.ctx, test.userBlock)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored := suite.keeper.GetAllUsersBlocks(suite.ctx)
				suite.Require().Equal(test.expBlocks, stored)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteUserBlock() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		data             struct {
			blocker  string
			blocked  string
			subspace string
		}
		expError  bool
		expBlocks []types.UserBlock
	}{
		{
			name: "delete user block with len(stored) > 1",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked_2", "reason", "subspace"),
			},
			expError: false,
		},
		{
			name: "delete user block with len(stored) == 1",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expError: false,
		},
		{
			name:             "deleting a user block that does not exist returns an error",
			storedUserBlocks: nil,
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expError: true,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, block := range test.storedUserBlocks {
				err := suite.keeper.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			err := suite.keeper.DeleteUserBlock(suite.ctx, test.data.blocker, test.data.blocked, test.data.subspace)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				blocks := suite.keeper.GetAllUsersBlocks(suite.ctx)
				suite.Require().Equal(test.expBlocks, blocks)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserBlocks() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		user             string
		expUserBlocks    []types.UserBlock
	}{
		{
			name: "non empty slice is returned properly",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			user: "blocker",
			expUserBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
		},
		{
			name:             "empty slice is returned properly",
			storedUserBlocks: nil,
			user:             "blocker",
			expUserBlocks:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, block := range test.storedUserBlocks {
				err := suite.keeper.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			blocks := suite.keeper.GetUserBlocks(suite.ctx, test.user)
			suite.Require().Equal(test.expUserBlocks, blocks)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllUsersBlocks() {
	tests := []struct {
		name              string
		storedUsersBlocks []types.UserBlock
		expUsersBlocks    []types.UserBlock
	}{
		{
			name: "Returns a non-empty users blocks slice",
			storedUsersBlocks: []types.UserBlock{
				types.NewUserBlock("user_1", "user_2", "reason", "subspace_1"),
				types.NewUserBlock("user_1", "user_2", "reason", "subspace_2"),
				types.NewUserBlock("user_2", "user_1", "reason", "subspace_1"),
				types.NewUserBlock("user_2", "user_1", "reason", "subspace_2"),
			},
			expUsersBlocks: []types.UserBlock{
				types.NewUserBlock("user_1", "user_2", "reason", "subspace_1"),
				types.NewUserBlock("user_1", "user_2", "reason", "subspace_2"),
				types.NewUserBlock("user_2", "user_1", "reason", "subspace_1"),
				types.NewUserBlock("user_2", "user_1", "reason", "subspace_2"),
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

			actualBlocks := suite.keeper.GetAllUsersBlocks(suite.ctx)

			suite.Require().Len(actualBlocks, len(test.expUsersBlocks))
			for _, block := range test.expUsersBlocks {
				suite.Require().Contains(actualBlocks, block)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_HasUserBlocked() {
	tests := []struct {
		name         string
		storedBlocks []types.UserBlock
		data         struct {
			blocker  string
			blocked  string
			subspace string
		}
		expBlocked bool
	}{
		{
			name: "blocked user found returns true",
			storedBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace",
			},
			expBlocked: true,
		},
		{
			name: "blocked user not found returns false",
			storedBlocks: []types.UserBlock{
				types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
			},
			data: struct {
				blocker  string
				blocked  string
				subspace string
			}{
				blocker:  "blocker",
				blocked:  "blocked",
				subspace: "subspace_2",
			},
			expBlocked: false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, block := range test.storedBlocks {
				err := suite.keeper.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			blocked := suite.keeper.HasUserBlocked(suite.ctx, test.data.blocker, test.data.blocked, test.data.subspace)
			suite.Equal(test.expBlocked, blocked)
		})
	}
}
