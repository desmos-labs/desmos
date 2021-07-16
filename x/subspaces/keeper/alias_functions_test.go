package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestKeeper_IterateSubspace() {
	date, err := time.Parse(time.RFC3339, "2010-10-02T12:10:00.000Z")
	suite.NoError(err)

	subspaces := []*types.Subspace{
		{
			ID:           "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Name:         "mooncake",
			Description:  "descr",
			Logo:         "https://shorturl.at/adnX3",
			Owner:        "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			Creator:      "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			CreationTime: date,
			Type:         types.SubspaceTypeOpen,
		},
		{
			ID:           "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
			Name:         "star",
			Description:  "descr",
			Logo:         "https://shorturl.at/adnX3",
			Owner:        "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			CreationTime: date,
			Type:         types.SubspaceTypeOpen,
		},
		{
			ID:           "e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
			Name:         "bad",
			Description:  "descr",
			Logo:         "https://shorturl.at/adnX3",
			Owner:        "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			CreationTime: date,
			Type:         types.SubspaceTypeOpen,
		},
	}

	expSubspaces := []*types.Subspace{
		subspaces[0],
		subspaces[1],
	}

	for _, subspace := range subspaces {
		_ = suite.k.SaveSubspace(suite.ctx, *subspace, subspace.Owner)
	}

	var validSubspaces []*types.Subspace
	suite.k.IterateSubspaces(suite.ctx, func(index int64, subspace types.Subspace) (stop bool) {
		if index == 2 {
			return false
		}
		validSubspaces = append(validSubspaces, &subspace)
		return false
	})

	suite.Len(expSubspaces, len(validSubspaces))
	suite.Equal(expSubspaces, validSubspaces)
}

func (suite *KeeperTestsuite) TestKeeper_GetAllSubspaces() {
	tests := []struct {
		name         string
		store        func(ctx sdk.Context)
		expSubspaces []types.Subspace
	}{
		{
			name: "Returns all the subspaces",
			store: func(ctx sdk.Context) {
				sub1 := types.NewSubspace(
					"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
					"test",
					"",
					"https://shorturl.at/adnX3",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err := suite.k.SaveSubspace(ctx, sub1, sub1.Owner)
				suite.Require().NoError(err)

				sub2 := types.NewSubspace(
					"C213BBE9641343190E4AAF12D882CD2F91EA588A9D4C18A20B0087C730DBA6CD",
					"mooncake",
					"",
					"https://shorturl.at/adnX3",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				err = suite.k.SaveSubspace(ctx, sub2, sub2.Owner)
				suite.Require().NoError(err)
			},
			expSubspaces: []types.Subspace{
				types.NewSubspace(
					"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
					"test",
					"",
					"https://shorturl.at/adnX3",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
				types.NewSubspace(
					"C213BBE9641343190E4AAF12D882CD2F91EA588A9D4C18A20B0087C730DBA6CD",
					"mooncake",
					"",
					"https://shorturl.at/adnX3",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name:         "Returns an empty slice with no subspaces",
			expSubspaces: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			subspaces := suite.k.GetAllSubspaces(suite.ctx)
			suite.Equal(test.expSubspaces, subspaces)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_IterateSubspaceAdmins() {
	subspace := types.NewSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"",
		"https://shorturl.at/adnX3",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		types.SubspaceTypeOpen,
		time.Now(),
	)
	suite.Require().NoError(suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner))

	admins := []string{
		"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm",
		"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
		"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
	}
	for _, admin := range admins {
		suite.Require().NoError(suite.k.AddAdminToSubspace(suite.ctx, subspace.ID, admin, subspace.Owner))
	}

	var iterated = 0
	suite.k.IterateSubspaceAdmins(suite.ctx, subspace.ID, func(index int64, admin string) (stop bool) {
		iterated++
		suite.Require().Contains(admins, admin)
		return false
	})
}

func (suite *KeeperTestsuite) TestKeeper_GetAllAdmins() {
	data := map[types.Subspace][]string{
		types.NewSubspace(
			"5A1B59C7DCB504039C04BCBC8C1C629CF5482374CF85ACD829C18102B301E299",
			"mooncake1",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {
			"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm",
		},
		types.NewSubspace(
			"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
			"mooncake2",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {
			"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
			"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
		},
		types.NewSubspace(
			"C213BBE9641343190E4AAF12D882CD2F91EA588A9D4C18A20B0087C730DBA6CD",
			"mooncake3",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {},
	}

	for subspace, admins := range data {
		err := suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
		suite.Require().NoError(err)

		for _, admin := range admins {
			err = suite.k.AddAdminToSubspace(suite.ctx, subspace.ID, admin, subspace.Owner)
			suite.Require().NoError(err)
		}
	}

	expected := []types.UsersEntry{
		types.NewUsersEntry(
			"5A1B59C7DCB504039C04BCBC8C1C629CF5482374CF85ACD829C18102B301E299",
			[]string{"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm"},
		),
		types.NewUsersEntry(
			"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
			[]string{
				"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
				"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
			},
		),
		types.NewUsersEntry(
			"C213BBE9641343190E4AAF12D882CD2F91EA588A9D4C18A20B0087C730DBA6CD",
			nil,
		),
	}

	stored := suite.k.GetAllAdmins(suite.ctx)
	suite.Require().Equal(expected, stored)
}

func (suite *KeeperTestsuite) TestKeeper_IterateSubspaceRegisteredUsers() {
	subspace := types.NewSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"",
		"https://shorturl.at/adnX3",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		types.SubspaceTypeOpen,
		time.Now(),
	)
	suite.Require().NoError(suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner))

	admins := []string{
		"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm",
		"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
		"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
	}
	for _, admin := range admins {
		suite.Require().NoError(suite.k.RegisterUserInSubspace(suite.ctx, subspace.ID, admin, subspace.Owner))
	}

	var iterated = 0
	suite.k.IterateSubspaceRegisteredUsers(suite.ctx, subspace.ID, func(index int64, admin string) (stop bool) {
		iterated++
		suite.Require().Contains(admins, admin)
		return false
	})
}

func (suite *KeeperTestsuite) TestKeeper_GetAllRegisteredUsers() {
	data := map[types.Subspace][]string{
		types.NewSubspace(
			"5A1B59C7DCB504039C04BCBC8C1C629CF5482374CF85ACD829C18102B301E299",
			"mooncake1",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {
			"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm",
		},
		types.NewSubspace(
			"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
			"mooncake2",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {
			"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
			"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
		},
		types.NewSubspace(
			"C213BBE9641343190E4AAF12D882CD2F91EA588A9D4C18A20B0087C730DBA6CD",
			"mooncake3",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {},
	}

	for subspace, users := range data {
		err := suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
		suite.Require().NoError(err)

		for _, admin := range users {
			err = suite.k.RegisterUserInSubspace(suite.ctx, subspace.ID, admin, subspace.Owner)
			suite.Require().NoError(err)
		}
	}

	expected := []types.UsersEntry{
		types.NewUsersEntry(
			"5A1B59C7DCB504039C04BCBC8C1C629CF5482374CF85ACD829C18102B301E299",
			[]string{"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm"},
		),
		types.NewUsersEntry(
			"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
			[]string{
				"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
				"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
			},
		),
		types.NewUsersEntry(
			"C213BBE9641343190E4AAF12D882CD2F91EA588A9D4C18A20B0087C730DBA6CD",
			nil,
		),
	}

	stored := suite.k.GetAllRegisteredUsers(suite.ctx)
	suite.Require().Equal(expected, stored)
}

func (suite *KeeperTestsuite) TestKeeper_IterateSubspaceBannedUsers() {
	subspace := types.NewSubspace(
		"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		"mooncake",
		"",
		"https://shorturl.at/adnX3",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
		types.SubspaceTypeOpen,
		time.Now(),
	)
	suite.Require().NoError(suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner))

	admins := []string{
		"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm",
		"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
		"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
	}
	for _, admin := range admins {
		suite.Require().NoError(suite.k.BanUserInSubspace(suite.ctx, subspace.ID, admin, subspace.Owner))
	}

	var iterated = 0
	suite.k.IterateSubspaceBannedUsers(suite.ctx, subspace.ID, func(index int64, admin string) (stop bool) {
		iterated++
		suite.Require().Contains(admins, admin)
		return false
	})
}

func (suite *KeeperTestsuite) TestKeeper_GetAllBannedUsers() {
	data := map[types.Subspace][]string{
		types.NewSubspace(
			"5A1B59C7DCB504039C04BCBC8C1C629CF5482374CF85ACD829C18102B301E299",
			"mooncake1",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {
			"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm",
		},
		types.NewSubspace(
			"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
			"mooncake2",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {
			"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
			"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
		},
		types.NewSubspace(
			"C213BBE9641343190E4AAF12D882CD2F91EA588A9D4C18A20B0087C730DBA6CD",
			"mooncake3",
			"",
			"https://shorturl.at/adnX3",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h",
			types.SubspaceTypeOpen,
			time.Now(),
		): {},
	}

	for subspace, users := range data {
		err := suite.k.SaveSubspace(suite.ctx, subspace, subspace.Owner)
		suite.Require().NoError(err)

		for _, admin := range users {
			err = suite.k.BanUserInSubspace(suite.ctx, subspace.ID, admin, subspace.Owner)
			suite.Require().NoError(err)
		}
	}

	expected := []types.UsersEntry{
		types.NewUsersEntry(
			"5A1B59C7DCB504039C04BCBC8C1C629CF5482374CF85ACD829C18102B301E299",
			[]string{"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm"},
		),
		types.NewUsersEntry(
			"A3C6CA0A7141715A61DFD73AB682C8E6B59C6D8C40F0231C2CFC7D21CF968476",
			[]string{
				"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
				"cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0",
			},
		),
		types.NewUsersEntry(
			"C213BBE9641343190E4AAF12D882CD2F91EA588A9D4C18A20B0087C730DBA6CD",
			nil,
		),
	}

	stored := suite.k.GetAllBannedUsers(suite.ctx)
	suite.Require().Equal(expected, stored)
}
