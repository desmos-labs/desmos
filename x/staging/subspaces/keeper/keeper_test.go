package keeper_test

/*
import (
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"time"
)

func (suite *KeeperTestsuite) TestKeeper_SaveSubspace() {
	tests := []struct {
		name             string
		existentSubspace *types.Subspace
		subspaceToSave   types.Subspace
		expErr           bool
	}{
		{
			name: "Already present subspaces returns error",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Unix(1, 1),
			},
			subspaceToSave: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Unix(1, 1),
			},
			expErr: true,
		},
		{
			name:             "Subspace saved correctly",
			existentSubspace: nil,
			subspaceToSave: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Unix(1, 1),
			},
			expErr: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.existentSubspace != nil {
				store := suite.ctx.KVStore(suite.storeKey)
				store.Set(types.SubspaceStoreKey(test.existentSubspace.ID),
					suite.cdc.MustMarshalBinaryBare(test.existentSubspace),
				)
			}

			err := suite.k.SaveSubspace(suite.ctx, test.subspaceToSave)
			if test.expErr {
				suite.Error(err)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DoesSubspaceExists() {
	tests := []struct {
		name       string
		subspaceID string
		subspace   *types.Subspace
		exists     bool
	}{
		{
			name:       "Subspace exists",
			subspaceID: "123",
			subspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Unix(1, 1),
			},
			exists: true,
		},
		{
			name:       "Subspace saved correctly",
			subspaceID: "123",
			subspace:   nil,
			exists:     false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.subspace != nil {
				_ = suite.k.SaveSubspace(suite.ctx, *test.subspace)
			}

			exists := suite.k.DoesSubspaceExists(suite.ctx, test.subspaceID)
			suite.Equal(test.exists, exists)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetSubspace() {
	tests := []struct {
		name       string
		subspaceID string
		subspace   *types.Subspace
		found      bool
	}{
		{
			name:       "Found the subspaces",
			subspaceID: "123",
			subspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
			},
			found: true,
		},
		{
			name:       "Subspace not found",
			subspaceID: "123",
			subspace:   nil,
			found:      false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.subspace != nil {
				_ = suite.k.SaveSubspace(suite.ctx, *test.subspace)
			}

			subspace, found := suite.k.GetSubspace(suite.ctx, test.subspaceID)
			if test.found {
				suite.True(found)
				suite.Equal(*test.subspace, subspace)
			} else {
				suite.False(found)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetAllSubspaces() {
	tests := []struct {
		name       string
		subspaceID string
		subspaces  []types.Subspace
	}{
		{
			name:       "Return all the subspaces",
			subspaceID: "123",
			subspaces: []types.Subspace{
				{
					ID:           "123",
					Name:         "test",
					Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					CreationTime: time.Time{},
				},
				{
					ID:           "124",
					Name:         "test",
					Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					CreationTime: time.Time{},
				},
			},
		},
		{
			name:       "Return empty subspaces array",
			subspaceID: "123",
			subspaces:  nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			for _, el := range test.subspaces {
				_ = suite.k.SaveSubspace(suite.ctx, el)
			}

			subspaces := suite.k.GetAllSubspaces(suite.ctx)
			suite.Equal(test.subspaces, subspaces)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_TransferSubspaceOwnership() {
	tests := []struct {
		name        string
		newOwner    string
		subspace    types.Subspace
		expSubspace types.Subspace
	}{
		{
			name:     "Transfer ownership correctly",
			newOwner: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			subspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
			},
			expSubspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				CreationTime: time.Time{},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			_ = suite.k.SaveSubspace(suite.ctx, test.subspace)
			suite.k.TransferSubspaceOwnership(suite.ctx, test.subspace.ID, test.newOwner)
			subspace, _ := suite.k.GetSubspace(suite.ctx, test.subspace.ID)

			suite.Equal(test.expSubspace, subspace)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_AddAdminToSubspace() {
	tests := []struct {
		name       string
		users      *types.Users
		user       string
		subspaceID string
		expUsers   *types.Users
		expError   bool
	}{
		{
			name: "Already added admin returns error",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			expError:   true,
		},
		{
			name: "Admin added correctly",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			expError:   false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			key := types.AdminsStoreKey(test.subspaceID)
			store.Set(key, types.MustMarshalUsers(suite.cdc, *test.users))

			err := suite.k.AddAdminToSubspace(suite.ctx, test.subspaceID, test.user)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
			}

			users := types.MustUnmarshalUsers(suite.cdc, store.Get(key))
			suite.Equal(test.expUsers.Users, users.Users)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_IsAdmin() {
	tests := []struct {
		name       string
		admins     types.Users
		user       string
		subspaceID string
		expBool    bool
	}{
		{
			name: "user is an admin",
			admins: types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspaceID: "123",
			expBool:    true,
		},
		{
			name: "user is not an admin",
			admins: types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			subspaceID: "123",
			expBool:    false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			key := types.AdminsStoreKey(test.subspaceID)
			store.Set(key, types.MustMarshalUsers(suite.cdc, test.admins))

			found := suite.k.IsAdmin(suite.ctx, test.user, test.subspaceID)
			suite.Equal(test.expBool, found)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetAllSubspaceAdmins() {
	tests := []struct {
		name       string
		admins     types.Users
		subspaceID string
	}{
		{
			name: "Returns all the admins correctly",
			admins: types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
		},
		{
			name: "Returns empty admins",
			admins: types.Users{
				Users: nil,
			},
			subspaceID: "123",
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			store := suite.ctx.KVStore(suite.storeKey)
			key := types.AdminsStoreKey(test.subspaceID)
			store.Set(key, types.MustMarshalUsers(suite.cdc, test.admins))

			admins := suite.k.GetSubspaceAdmins(suite.ctx, test.subspaceID)
			suite.Equal(test.admins, admins)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RemoveAdminFromSubspace() {
	tests := []struct {
		name       string
		users      *types.Users
		user       string
		subspaceID string
		expUsers   *types.Users
		expError   bool
	}{
		{
			name: "Already removed admin returns error",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			expError:   true,
		},
		{
			name: "Admin removed correctly",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			expError:   false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			store := suite.ctx.KVStore(suite.storeKey)
			key := types.AdminsStoreKey(test.subspaceID)
			store.Set(key, types.MustMarshalUsers(suite.cdc, *test.users))

			err := suite.k.RemoveAdminFromSubspace(suite.ctx, test.subspaceID, test.user)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
			}

			users := types.MustUnmarshalUsers(suite.cdc, store.Get(key))
			suite.Equal(test.expUsers.Users, users.Users)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_UnblockPostsForUser() {
	tests := []struct {
		name       string
		users      *types.Users
		user       string
		subspaceID string
		expUsers   *types.Users
		expError   bool
	}{
		{
			name: "Already unblocked user returns error",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			expError:   true,
		},
		{
			name: "User unblock correctly",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			expError:   false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			store := suite.ctx.KVStore(suite.storeKey)
			key := types.BlockedUsersStoreKey(test.subspaceID)
			store.Set(key, types.MustMarshalUsers(suite.cdc, *test.users))

			err := suite.k.RegisterUserInSubspace(suite.ctx, test.user, test.subspaceID)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
			}

			users := types.MustUnmarshalUsers(suite.cdc, store.Get(key))
			suite.Equal(test.expUsers.Users, users.Users)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_BlockPostsForUser() {
	tests := []struct {
		name       string
		users      *types.Users
		user       string
		subspaceID string
		expUsers   *types.Users
		expError   bool
	}{
		{
			name: "Already added admin returns error",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			expError:   true,
		},
		{
			name: "Admin added correctly",
			users: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			user: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expUsers: &types.Users{
				Users: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			expError:   false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)
			key := types.AdminsStoreKey(test.subspaceID)
			store.Set(key, types.MustMarshalUsers(suite.cdc, *test.users))

			err := suite.k.AddAdminToSubspace(suite.ctx, test.subspaceID, test.user)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
			}

			users := types.MustUnmarshalUsers(suite.cdc, store.Get(key))
			suite.Equal(test.expUsers.Users, users.Users)
		})
	}
}*/
