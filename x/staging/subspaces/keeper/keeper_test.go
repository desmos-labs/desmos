package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_SaveSubspace() {
	tests := []struct {
		name            string
		storedSubspaces []types.Subspace
		subspaceToSave  types.Subspace
		expErr          bool
		expStored       []types.Subspace
	}{
		{
			name: "Already stored subspace with different owner",
			storedSubspaces: []types.Subspace{
				types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
			subspaceToSave: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			expErr: true,
		},
		{
			name: "New subspace saved correctly",
			subspaceToSave: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			expErr: false,
			expStored: []types.Subspace{
				types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "Old subspace edited correctly",
			storedSubspaces: []types.Subspace{
				types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
			subspaceToSave: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test-updated",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeClosed,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			expErr: false,
			expStored: []types.Subspace{
				types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test-updated",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeClosed,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			store := suite.ctx.KVStore(suite.storeKey)
			for _, subspace := range test.storedSubspaces {
				store.Set(types.SubspaceStoreKey(subspace.ID), suite.cdc.MustMarshal(&subspace))
			}

			err := suite.k.SaveSubspace(suite.ctx, test.subspaceToSave, test.subspaceToSave.Owner)
			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.Require().Equal(test.expStored, suite.k.GetAllSubspaces(suite.ctx))
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_DoesSubspaceExists() {
	tests := []struct {
		name            string
		subspaceID      string
		storedSubspaces []types.Subspace
		exists          bool
	}{
		{
			name:       "Returns true when the subspace exists",
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			storedSubspaces: []types.Subspace{
				types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				),
			},
			exists: true,
		},
		{
			name:       "Return false when the subspace doesn't exist",
			subspaceID: "123",
			exists:     false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			for _, subspace := range test.storedSubspaces {
				suite.Require().NoError(suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner))
			}

			exists := suite.k.DoesSubspaceExist(suite.ctx, test.subspaceID)
			suite.Equal(test.exists, exists)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_GetSubspace() {
	tests := []struct {
		name        string
		store       func(ctx sdk.Context)
		subspaceID  string
		expFound    bool
		expSubspace types.Subspace
	}{
		{
			name:       "Return false when not found",
			subspaceID: "123",
			expFound:   false,
		},
		{
			name: "Returns the subspace and true when found",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			expFound:   true,
			expSubspace: types.NewSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types.SubspaceTypeOpen,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			subspace, found := suite.k.GetSubspace(suite.ctx, test.subspaceID)
			suite.Require().Equal(test.expFound, found)

			if test.expFound {
				suite.Equal(test.expSubspace, subspace)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_AddAdminToSubspace() {
	tests := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID string
		user       string
		owner      string
		expError   bool
		expAdmins  []string
	}{
		{
			name:       "Non existent subspace returns error",
			subspaceID: "",
			expError:   true,
		},
		{
			name: "User already an admin returns error",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   true,
			expAdmins: []string{
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
		{
			name: "User added as admin correctly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   false,
			expAdmins: []string{
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			err := suite.k.AddAdminToSubspace(suite.ctx, test.subspaceID, test.user, test.owner)
			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			var admins []string
			suite.k.IterateSubspaceAdmins(suite.ctx, test.subspaceID, func(index int64, admin string) (stop bool) {
				admins = append(admins, admin)
				return false
			})
			suite.Require().Equal(test.expAdmins, admins)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_RemoveAdminFromSubspace() {
	tests := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID string
		user       string
		owner      string
		expAdmins  []string
		expError   bool
	}{
		{
			name:       "Non existent subspace returns error",
			subspaceID: "non-existing",
			expError:   true,
		},
		{
			name: "Invalid admin returns error",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   true,
			expAdmins: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
		{
			name: "One of many admins is removed correctly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   false,
			expAdmins: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
		{
			name: "Single admin is removed correctly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			owner:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   false,
			expAdmins:  nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			err := suite.k.RemoveAdminFromSubspace(suite.ctx, test.subspaceID, test.user, test.owner)
			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})

		var admins []string
		suite.k.IterateSubspaceAdmins(suite.ctx, test.subspaceID, func(index int64, admin string) (stop bool) {
			admins = append(admins, admin)
			return false
		})
		suite.Require().Equal(test.expAdmins, admins)
	}
}

func (suite *KeeperTestsuite) TestKeeper_RegisterUserInSubspace() {
	tests := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID string
		user       string
		admin      string
		expError   bool
		expUsers   []string
	}{
		{
			name:       "Non existent subspace returns error",
			subspaceID: "non-existing",
			expError:   true,
		},
		{
			name: "User already registered returns error",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   true,
			expUsers: []string{
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
		{
			name: "User registered correctly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   false,
			expUsers: []string{
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			err := suite.k.RegisterUserInSubspace(suite.ctx, test.subspaceID, test.user, test.admin)
			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			var users []string
			suite.k.IterateSubspaceRegisteredUsers(suite.ctx, test.subspaceID, func(index int64, user string) (stop bool) {
				users = append(users, user)
				return false
			})
			suite.Require().Equal(users, test.expUsers)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_UnregisterUserInSubspace() {
	tests := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID string
		user       string
		admin      string
		expError   bool
		expUsers   []string
	}{
		{
			name:       "Non existent subspace returns error",
			subspaceID: "non-existing",
			expError:   true,
		},
		{
			name: "Invalid user returns error",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   true,
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
		{
			name: "Valid user unregistered correctly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   false,
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			err := suite.k.UnregisterUserFromSubspace(suite.ctx, test.subspaceID, test.user, test.admin)
			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})

		var users []string
		suite.k.IterateSubspaceRegisteredUsers(suite.ctx, test.subspaceID, func(index int64, user string) (stop bool) {
			users = append(users, user)
			return false
		})
		suite.Require().Equal(users, test.expUsers)
	}
}

func (suite *KeeperTestsuite) TestKeeper_BanUser() {
	tests := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID string
		user       string
		admin      string
		expError   bool
		expUsers   []string
	}{
		{
			name:       "Non existent subspace returns error",
			subspaceID: "non-existing",
			expError:   true,
		},
		{
			name: "User already banned returns error",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   true,
			expUsers: []string{
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
		{
			name: "User banned correctly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   false,
			expUsers: []string{
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			err := suite.k.BanUserInSubspace(suite.ctx, test.subspaceID, test.user, test.admin)
			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})

		var users []string
		suite.k.IterateSubspaceBannedUsers(suite.ctx, test.subspaceID, func(index int64, user string) (stop bool) {
			users = append(users, user)
			return false
		})
		suite.Require().Equal(users, test.expUsers)
	}
}

func (suite *KeeperTestsuite) TestKeeper_UnbanUser() {
	tests := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID string
		user       string
		admin      string
		expUsers   []string
		expError   bool
	}{
		{
			name:       "Non existent subspace returns error",
			subspaceID: "non-existing",
			expError:   true,
		},
		{
			name: "Invalid user returns error",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   true,
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
		{
			name: "User unbanned correctly",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			admin:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			expError:   false,
			expUsers: []string{
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			err := suite.k.UnbanUserInSubspace(suite.ctx, test.subspaceID, test.user, test.admin)
			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}

			var users []string
			suite.k.IterateSubspaceBannedUsers(suite.ctx, test.subspaceID, func(index int64, user string) (stop bool) {
				users = append(users, user)
				return false
			})
			suite.Require().Equal(users, test.expUsers)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_CheckSubspaceUserPermission() {
	tests := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID string
		user       string
		expError   bool
	}{
		{
			name:       "Non existent subspace returns error",
			subspaceID: "non-existing",
			expError:   true,
		},
		{
			name: "Banned user returns error",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expError:   true,
		},
		{
			name: "Subspace types closed and not registered user returns error",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeClosed,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expError:   true,
		},
		{
			name: "No errors",
			store: func(ctx sdk.Context) {
				subspace := types.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types.SubspaceTypeOpen,
					time.Unix(1, 1),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", subspace.Owner)
				suite.Require().NoError(err)
			},
			subspaceID: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			user:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expError:   false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			err := suite.k.CheckSubspaceUserPermission(suite.ctx, test.subspaceID, test.user)
			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}
