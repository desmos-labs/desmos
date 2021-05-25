package keeper_test

import "github.com/desmos-labs/desmos/x/profiles/types"

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
				err := suite.k.SaveRelationship(suite.ctx, relationship)
				suite.Require().NoError(err)
			}

			err := suite.k.SaveRelationship(suite.ctx, test.relationship)

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
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			relationships := suite.k.GetAllRelationships(suite.ctx)

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
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			relationships := suite.k.GetUserRelationships(suite.ctx, test.user)
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
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			err := suite.k.RemoveRelationship(suite.ctx, test.relationshipToDelete)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				rel := suite.k.GetAllRelationships(suite.ctx)
				suite.Require().Equal(test.expRelationships, rel)
			}
		})
	}
}
