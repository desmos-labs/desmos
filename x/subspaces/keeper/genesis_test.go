package keeper_test

import (
	"time"

	types2 "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestsuite) TestKeeper_ExportGenesis() {
	tests := []struct {
		name     string
		store    func(ctx sdk.Context)
		expected *types2.GenesisState
	}{
		{
			name:     "Default expected state",
			expected: &types2.GenesisState{Subspaces: nil},
		},
		{
			name: "Genesis exported successfully",
			store: func(ctx sdk.Context) {

				subspace := types2.NewSubspace(
					"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5", subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", subspace.Owner)
				suite.Require().NoError(err)
			},
			expected: types2.NewGenesisState(
				[]types2.Subspace{
					types2.NewSubspace(
						"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
						"test",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types2.SubspaceTypeOpen,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
						[]string{"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm"},
					),
				},
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
						[]string{"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5"},
					),
				},
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
						[]string{"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0"},
					),
				},
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

			exported := suite.k.ExportGenesis(suite.ctx)
			suite.Equal(test.expected, exported)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_InitGenesis() {
	tests := []struct {
		name     string
		genesis  *types2.GenesisState
		expError bool
		check    func(ctx sdk.Context)
	}{
		{
			name:     "Default genesis is initialized properly",
			genesis:  types2.DefaultGenesisState(),
			expError: false,
		},
		{
			name: "Invalid subspace panics",
			genesis: types2.NewGenesisState(
				[]types2.Subspace{
					types2.NewSubspace(
						"",
						"test",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types2.SubspaceTypeOpen,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				nil,
				nil,
				nil,
			),
			expError: true,
		},
		{
			name: "Admins entry for non existing subspace returns error",
			genesis: types2.NewGenesisState(
				nil,
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
					),
				},
				nil,
				nil,
			),
			expError: true,
		},
		{
			name: "Duplicated admins returns error",
			genesis: types2.NewGenesisState(
				[]types2.Subspace{
					types2.NewSubspace(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						"test",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types2.SubspaceTypeOpen,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						},
					),
				},
				nil,
				nil,
			),
			expError: true,
		},
		{
			name: "Registered users entry for non existing subspace returns error",
			genesis: types2.NewGenesisState(
				nil,
				nil,
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
					),
				},
				nil,
			),
			expError: true,
		},
		{
			name: "Duplicated registered users returns error",
			genesis: types2.NewGenesisState(
				[]types2.Subspace{
					types2.NewSubspace(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						"test",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types2.SubspaceTypeOpen,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				nil,
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						},
					),
				},
				nil,
			),
			expError: true,
		},
		{
			name: "Banned users entry for non existing subspace returns error",
			genesis: types2.NewGenesisState(
				nil,
				nil,
				nil,
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
					),
				},
			),
			expError: true,
		},
		{
			name: "Duplicated banned users returns error",
			genesis: types2.NewGenesisState(
				[]types2.Subspace{
					types2.NewSubspace(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						"test",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types2.SubspaceTypeOpen,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				nil,
				nil,
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						},
					),
				},
			),
			expError: true,
		},
		{
			name: "Valid genesis initialized correctly",
			genesis: types2.NewGenesisState(
				[]types2.Subspace{
					types2.NewSubspace(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						"test",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types2.SubspaceTypeOpen,
						time.Date(2020, 1, 1, 0, 00, 00, 000, time.UTC),
					),
				},
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5"},
					),
				},
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm"},
					),
				},
				[]types2.UsersEntry{
					types2.NewUsersEntry(
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
						[]string{"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0"},
					),
				},
			),
			check: func(ctx sdk.Context) {
				subspace, found := suite.k.GetSubspace(ctx, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
				suite.Require().True(found)
				suite.Require().True(subspace.Equal(types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 0, 00, 00, 000, time.UTC),
				)))

				suite.Require().Equal(
					[]types2.UsersEntry{
						types2.NewUsersEntry(
							"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
							[]string{"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5"},
						),
					},
					suite.k.GetAllAdmins(ctx),
				)

				suite.Require().Equal(
					[]types2.UsersEntry{
						types2.NewUsersEntry(
							"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
							[]string{"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm"},
						),
					},
					suite.k.GetAllRegisteredUsers(ctx),
				)

				suite.Require().Equal(
					[]types2.UsersEntry{
						types2.NewUsersEntry(
							"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
							[]string{"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0"},
						),
					},
					suite.k.GetAllBannedUsers(ctx),
				)
			},
			expError: false,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.expError {
				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, test.genesis) })
			} else {
				suite.k.InitGenesis(suite.ctx, test.genesis)
				if test.check != nil {
					test.check(suite.ctx)
				}
			}
		})
	}
}
