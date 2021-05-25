package keeper_test

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"time"

	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SaveSubspace() {
	tests := []struct {
		name           string
		subspaceToSave types.Subspace
		expErr         bool
	}{
		{
			name: "Subspace saved correctly",
			subspaceToSave: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
			},
			expErr: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.k.SaveSubspace(suite.ctx, test.subspaceToSave)

			store := suite.ctx.KVStore(suite.storeKey)
			key := types.SubspaceStoreKey(test.subspaceToSave.ID)

			subspaceBytes := store.Get(key)
			var subspace types.Subspace
			suite.cdc.MustUnmarshalBinaryBare(subspaceBytes, &subspace)

			suite.Equal(test.subspaceToSave, subspace)
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
			name:       "Returns true when the subspace exists",
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
			name:       "Return false when the subspace doesn't exist",
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
				suite.k.SaveSubspace(suite.ctx, *test.subspace)
			}

			exists := suite.k.DoesSubspaceExist(suite.ctx, test.subspaceID)
			suite.Equal(test.exists, exists)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetSubspace() {
	tests := []struct {
		name       string
		subspaceID string
		subspace   types.Subspace
		found      bool
	}{
		{
			name:       "Return the subspace and true when found",
			subspaceID: "123",
			subspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
			},
			found: true,
		},
		{
			name:       "Return empty subspace and false when not found",
			subspaceID: "123",
			subspace:   types.Subspace{},
			found:      false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.found {
				suite.k.SaveSubspace(suite.ctx, test.subspace)
			}

			subspace, found := suite.k.GetSubspace(suite.ctx, test.subspaceID)
			if test.found {
				suite.True(found)
				suite.Equal(test.subspace, subspace)
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
				suite.k.SaveSubspace(suite.ctx, el)
			}

			subspaces := suite.k.GetAllSubspaces(suite.ctx)
			suite.Equal(test.subspaces, subspaces)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_ValidateSubspace() {
	date, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.NoError(err)

	tests := []struct {
		name     string
		subspace types.Subspace
		expError error
	}{
		{
			name:     "Subspace name not matching the regEx returns error",
			subspace: types.NewSubspace("", ".!#", "", "", true, time.Time{}),
			expError: sdkerrors.Wrapf(types.ErrInvalidSubspaceName, "invalid subspace name, it should match the following regEx ^[A-Za-z0-9_]+$"),
		},
		{
			name:     "Subspace name not reaching the min length returns error",
			subspace: types.NewSubspace("", "na", "", "", true, time.Time{}),
			expError: sdkerrors.Wrapf(types.ErrInvalidSubspaceNameLength, "subspace name cannot be less than 3 characters"),
		},
		{
			name:     "Subspace name exceeding the max length returns error",
			subspace: types.NewSubspace("", "nametoolongtobeaccepted", "", "", true, time.Time{}),
			expError: sdkerrors.Wrapf(types.ErrInvalidSubspaceNameLength, "subspace name cannot exceed 10 characters"),
		},
		{
			name: "Valid subspace returns no error",
			subspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"mooncake",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				true,
				date,
			),
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.k.SetParams(suite.ctx, types.DefaultParams())

			err := suite.k.ValidateSubspace(suite.ctx, test.subspace)
			if test.expError != nil {
				suite.Equal(test.expError.Error(), err.Error())
			} else {
				suite.NoError(err)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_AddAdminToSubspace() {
	tests := []struct {
		name             string
		existentSubspace *types.Subspace
		subspaceID       string
		user             string
		owner            string
		expAdmins        []string
		expError         bool
	}{
		{
			name:             "Non existent subspace returns error",
			existentSubspace: nil,
			subspaceID:       "",
			user:             "",
			owner:            "",
			expAdmins:        nil,
			expError:         true,
		},
		{
			name: "User already an admin returns error",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				Admins: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expAdmins: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: true,
		},
		{
			name: "User added as admin correctly",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				Admins: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expAdmins: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.existentSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.existentSubspace)
			}

			err := suite.k.AddAdminToSubspace(suite.ctx, test.subspaceID, test.user, test.owner)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
				subspace, found := suite.k.GetSubspace(suite.ctx, test.existentSubspace.ID)
				suite.True(found)
				suite.Equal(test.expAdmins, subspace.Admins)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RemoveAdminFromSubspace() {
	tests := []struct {
		name             string
		existentSubspace *types.Subspace
		subspaceID       string
		user             string
		owner            string
		expAdmins        []string
		expError         bool
	}{
		{
			name:             "Non existent subspace returns error",
			existentSubspace: nil,
			subspaceID:       "",
			user:             "",
			owner:            "",
			expAdmins:        nil,
			expError:         true,
		},
		{
			name: "User already not an returns error",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				Admins: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expAdmins: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			expError: true,
		},
		{
			name: "Admin removed correctly",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				Admins: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expAdmins: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.existentSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.existentSubspace)
			}

			err := suite.k.RemoveAdminFromSubspace(suite.ctx, test.subspaceID, test.user, test.owner)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
				subspace, found := suite.k.GetSubspace(suite.ctx, test.subspaceID)
				suite.True(found)
				suite.Equal(test.expAdmins, subspace.Admins)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RegisterUserInSubspace() {
	tests := []struct {
		name             string
		existentSubspace *types.Subspace
		subspaceID       string
		user             string
		admin            string
		expUsers         []string
		expError         bool
	}{
		{
			name:             "Non existent subspace returns error",
			existentSubspace: nil,
			subspaceID:       "",
			user:             "",
			admin:            "",
			expUsers:         nil,
			expError:         true,
		},
		{
			name: "User already registered returns error",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				RegisteredUsers: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: true,
		},
		{
			name: "User registered correctly",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				RegisteredUsers: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.existentSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.existentSubspace)
			}

			err := suite.k.RegisterUserInSubspace(suite.ctx, test.subspaceID, test.user, test.admin)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
				subspace, found := suite.k.GetSubspace(suite.ctx, test.subspaceID)
				suite.True(found)
				suite.Equal(test.expUsers, subspace.RegisteredUsers)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_UnregisterUserInSubspace() {
	tests := []struct {
		name             string
		existentSubspace *types.Subspace
		subspaceID       string
		user             string
		admin            string
		expUsers         []string
		expError         bool
	}{
		{
			name:             "Non existent subspace returns error",
			existentSubspace: nil,
			subspaceID:       "",
			user:             "",
			admin:            "",
			expUsers:         nil,
			expError:         true,
		},
		{
			name: "User already unregistered returns error",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				RegisteredUsers: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			expError: true,
		},
		{
			name: "User unregistered correctly",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				RegisteredUsers: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.existentSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.existentSubspace)
			}

			err := suite.k.UnregisterUserFromSubspace(suite.ctx, test.subspaceID, test.user, test.admin)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
				subspace, found := suite.k.GetSubspace(suite.ctx, test.subspaceID)
				suite.True(found)
				suite.Equal(test.expUsers, subspace.RegisteredUsers)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_BanUser() {
	tests := []struct {
		name             string
		existentSubspace *types.Subspace
		subspaceID       string
		user             string
		admin            string
		expUsers         []string
		expError         bool
	}{
		{
			name:             "Non existent subspace returns error",
			existentSubspace: nil,
			subspaceID:       "",
			user:             "",
			admin:            "",
			expUsers:         nil,
			expError:         true,
		},
		{
			name: "User already blocked returns error",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				BannedUsers: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: true,
		},
		{
			name: "User blocked correctly",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				BannedUsers: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.existentSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.existentSubspace)
			}

			err := suite.k.BanUserInSubspace(suite.ctx, test.subspaceID, test.user, test.admin)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
				subspace, found := suite.k.GetSubspace(suite.ctx, test.subspaceID)
				suite.True(found)
				suite.Equal(test.expUsers, subspace.BannedUsers)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_UnbanUser() {
	tests := []struct {
		name             string
		existentSubspace *types.Subspace
		subspaceID       string
		user             string
		admin            string
		expUsers         []string
		expError         bool
	}{
		{
			name:             "Non existent subspace returns error",
			existentSubspace: nil,
			subspaceID:       "",
			user:             "",
			admin:            "",
			expUsers:         nil,
			expError:         true,
		},
		{
			name: "User already unblocked returns error",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				BannedUsers: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			expError: true,
		},
		{
			name: "User unblocked correctly",
			existentSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Unix(1, 1),
				BannedUsers: []string{
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				},
			},
			subspaceID: "123",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.existentSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.existentSubspace)
			}

			err := suite.k.UnbanUserInSubspace(suite.ctx, test.subspaceID, test.user, test.admin)
			if test.expError {
				suite.NotNil(err)
				suite.Error(err)
			} else {
				suite.Nil(err)
				subspace, found := suite.k.GetSubspace(suite.ctx, test.subspaceID)
				suite.True(found)
				suite.Equal(test.expUsers, subspace.BannedUsers)
			}
		})
	}
}
