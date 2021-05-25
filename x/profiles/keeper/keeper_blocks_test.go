package keeper_test

import "github.com/desmos-labs/desmos/x/profiles/types"

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
			if test.userBlocks != nil {
				_ = suite.k.SaveUserBlock(suite.ctx, test.userBlocks[0])
			}
			res := suite.k.IsUserBlocked(suite.ctx, test.blocker, test.blocked)
			suite.Equal(test.expBool, res)
		})
	}
}

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
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			err := suite.k.SaveUserBlock(suite.ctx, test.userBlock)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				stored := suite.k.GetAllUsersBlocks(suite.ctx)
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
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			err := suite.k.DeleteUserBlock(suite.ctx, test.data.blocker, test.data.blocked, test.data.subspace)

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				blocks := suite.k.GetAllUsersBlocks(suite.ctx)
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
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			blocks := suite.k.GetUserBlocks(suite.ctx, test.user)
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
				err := suite.k.SaveUserBlock(suite.ctx, userBlock)
				suite.Require().NoError(err)
			}

			actualBlocks := suite.k.GetAllUsersBlocks(suite.ctx)

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
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			blocked := suite.k.HasUserBlocked(suite.ctx, test.data.blocker, test.data.blocked, test.data.subspace)
			suite.Equal(test.expBlocked, blocked)
		})
	}
}
