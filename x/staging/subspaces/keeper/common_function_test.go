package keeper_test

import "github.com/desmos-labs/desmos/x/staging/subspaces/types"

func (suite *KeeperTestsuite) TestKeeper_AddUserToList() {
	tests := []struct {
		name       string
		users      *types.Users
		user       string
		error      string
		subspaceID string
		expError   bool
	}{
		{
			name: "Already present user returns error",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			error:      "the user: %s is already an admin of the subspaces: %s",
			subspaceID: "123",
			expError:   true,
		},
		{
			name:       "New user returns no error",
			users:      nil,
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			error:      "",
			subspaceID: "123",
			expError:   false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			key := types.AdminsStoreKey(test.subspaceID)
			if test.users != nil {
				store := suite.ctx.KVStore(suite.storeKey)
				store.Set(key, types.MustMarshalUsers(suite.cdc, *test.users))
			}

			err := suite.k.AddUserToList(suite.ctx, key, test.subspaceID, test.user, test.error)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RemoveUserFromList() {
	tests := []struct {
		name       string
		users      *types.Users
		user       string
		error      string
		subspaceID string
		expError   bool
		expUsers   *types.Users
	}{
		{
			name: "Not present user returns error",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			error:      "the user: %s already can't post inside the subspaces: %s",
			subspaceID: "123",
			expError:   true,
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
		},
		{
			name: "User removed correctly",
			users: &types.Users{
				Users: []string{
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			error:      "",
			subspaceID: "123",
			expError:   false,
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
		},
		{
			name: "User removed correctly (key deleted)",
			users: &types.Users{
				Users: []string{
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			error:      "",
			subspaceID: "123",
			expError:   false,
			expUsers:   &types.Users{Users: nil},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			key := types.AdminsStoreKey(test.subspaceID)
			store := suite.ctx.KVStore(suite.storeKey)
			if test.users != nil {
				store.Set(key, types.MustMarshalUsers(suite.cdc, *test.users))
			}

			err := suite.k.RemoveUserFromList(suite.ctx, key, test.subspaceID, test.user, test.error)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
			}

			bz := store.Get(key)
			users := types.MustUnmarshalUsers(suite.cdc, bz)
			suite.Equal(test.expUsers.Users, users.Users)
		})
	}
}
