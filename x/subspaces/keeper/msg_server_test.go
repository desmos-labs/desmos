package keeper_test

import (
	keeper2 "github.com/desmos-labs/desmos/x/subspaces/keeper"
	types2 "github.com/desmos-labs/desmos/x/subspaces/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestsuite) TestMsgServer_CreateSubspace() {
	creationTime := time.Date(2050, 01, 01, 15, 15, 00, 000, time.UTC)

	tests := []struct {
		name      string
		blockTime time.Time
		store     func(ctx sdk.Context)
		msg       *types2.MsgCreateSubspace
		expErr    bool
		expEvents sdk.Events
	}{
		{
			name:      "Subspace already existing returns error",
			blockTime: creationTime,
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types2.SubspaceTypeOpen,
					ctx.BlockTime(),
				)
				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgCreateSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test2",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types2.SubspaceTypeOpen,
			),
			expErr: true,
		},
		{
			name:      "Non existing subspace is saved properly",
			blockTime: creationTime,
			msg: types2.NewMsgCreateSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test2",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				types2.SubspaceTypeOpen,
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeCreateSubspace,
					sdk.NewAttribute(types2.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types2.AttributeKeySubspaceName, "test2"),
					sdk.NewAttribute(types2.AttributeKeySubspaceCreator, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types2.AttributeKeyCreationTime, creationTime.Format(time.RFC3339)),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.ctx = suite.ctx.WithBlockTime(test.blockTime)
			if test.store != nil {
				test.store(suite.ctx)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err := handler.CreateSubspace(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				events := suite.ctx.EventManager().Events()
				suite.Equal(test.expEvents, events)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_EditSubspace() {
	creationTime := time.Date(2050, 01, 01, 15, 15, 00, 000, time.UTC)
	tests := []struct {
		name      string
		blockTime time.Time
		store     func(ctx sdk.Context)
		msg       *types2.MsgEditSubspace
		expErr    bool
		expEvents sdk.Events
		expStored []types2.Subspace
	}{
		{
			name:      "Non existing subspace returns error",
			blockTime: creationTime,
			msg: types2.NewMsgEditSubspace(
				"1234",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"edited",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types2.SubspaceTypeOpen,
			),
			expErr: true,
		},
		{
			name:      "Wrong editor returns error",
			blockTime: creationTime,
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types2.SubspaceTypeOpen,
					creationTime,
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgEditSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"ccosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"edited",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types2.SubspaceTypeClosed,
			),
			expErr: true,
		},
		{
			name:      "subspace edited successfully",
			blockTime: creationTime,
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types2.SubspaceTypeOpen,
					creationTime,
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgEditSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"edited",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				types2.SubspaceTypeClosed,
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeEditSubspace,
					sdk.NewAttribute(types2.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types2.AttributeKeyNewOwner, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types2.AttributeKeySubspaceName, "edited"),
				),
			},
			expStored: []types2.Subspace{
				types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"edited",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					types2.SubspaceTypeClosed,
					creationTime,
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.ctx = suite.ctx.WithBlockTime(test.blockTime)
			if test.store != nil {
				test.store(suite.ctx)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err := handler.EditSubspace(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Equal(test.expEvents, suite.ctx.EventManager().Events())

				var subspaces []types2.Subspace
				suite.k.IterateSubspaces(suite.ctx, func(index int64, subspace types2.Subspace) (stop bool) {
					subspaces = append(subspaces, subspace)
					return false
				})
				suite.Require().Equal(test.expStored, subspaces)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_AddAdmin() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types2.MsgAddAdmin
		expErr    bool
		expEvents sdk.Events
		expAdmins []string
	}{
		{
			name: "Non existing subspace returns error",
			msg: types2.NewMsgAddAdmin(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Duplicated admin returns error",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgAddAdmin(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Wrong owner returns error",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgAddAdmin(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expErr: true,
		},
		{
			name: "Valid request adds the admin properly",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgAddAdmin(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeAddAdmin,
					sdk.NewAttribute(types2.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types2.AttributeKeySubspaceNewAdmin, "cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5"),
				),
			},
			expAdmins: []string{
				"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err := handler.AddAdmin(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				suite.Equal(test.expEvents, suite.ctx.EventManager().Events())

				var admins []string
				suite.k.IterateSubspaceAdmins(suite.ctx, test.msg.SubspaceID, func(index int64, admin string) (stop bool) {
					admins = append(admins, admin)
					return false
				})
				suite.Require().Equal(test.expAdmins, admins)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_RemoveAdmin() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types2.MsgRemoveAdmin
		expErr    bool
		expEvents sdk.Events
		expAdmins []string
	}{
		{
			name: "Non existing subspace returns error",
			msg: types2.NewMsgRemoveAdmin(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Non existing admin returns error",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgRemoveAdmin(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Existing admin is removed successfully",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.AddAdminToSubspace(ctx, subspace.ID, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgRemoveAdmin(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeRemoveAdmin,
					sdk.NewAttribute(types2.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types2.AttributeKeySubspaceRemovedAdmin, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				),
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

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err := handler.RemoveAdmin(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				var admins []string
				suite.k.IterateSubspaceAdmins(suite.ctx, test.msg.SubspaceID, func(index int64, admin string) (stop bool) {
					admins = append(admins, admin)
					return false
				})
				suite.Require().Equal(test.expAdmins, admins)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_RegisterUser() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types2.MsgRegisterUser
		expErr    bool
		expEvents sdk.Events
		expUsers  []string
	}{
		{
			name: "Non existing subspace returns error",
			msg: types2.NewMsgRegisterUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Registered user returns error",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgRegisterUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
			expUsers: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			},
		},
		{
			name: "Non registered user is registered correctly",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgRegisterUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeRegisterUser,
					sdk.NewAttribute(types2.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types2.AttributeKeyRegisteredUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				),
			},
			expUsers: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err := handler.RegisterUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Equal(test.expEvents, suite.ctx.EventManager().Events())
			}

			var users []string
			suite.k.IterateSubspaceRegisteredUsers(suite.ctx, test.msg.SubspaceID, func(index int64, user string) (stop bool) {
				users = append(users, user)
				return false
			})
			suite.Require().Equal(test.expUsers, users)
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_UnregisterUser() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types2.MsgUnregisterUser
		expErr    bool
		expEvents sdk.Events
		expUsers  []string
	}{
		{
			name: "Non existing subspace returns error",
			msg: types2.NewMsgUnregisterUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Not found user returns error",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgUnregisterUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
			expUsers: []string{
				"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
			},
		},
		{
			name: "Valid user unregistered successfully",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.RegisterUserInSubspace(ctx, subspace.ID, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgUnregisterUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeUnregisterUser,
					sdk.NewAttribute(types2.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types2.AttributeKeyUnregisteredUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				),
			},
			expUsers: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err := handler.UnregisterUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}

			var users []string
			suite.k.IterateSubspaceRegisteredUsers(suite.ctx, test.msg.SubspaceID, func(index int64, user string) (stop bool) {
				users = append(users, user)
				return false
			})
			suite.Require().Equal(test.expUsers, users)
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_BlockUser() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types2.MsgBanUser
		expErr    bool
		expEvents sdk.Events
		expUsers  []string
	}{
		{
			name: "Non existing subspace returns error",
			msg: types2.NewMsgBanUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Duplicated ban returns error",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgBanUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
			expUsers: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			},
		},
		{
			name: "Valid ban request works properly",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgBanUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeBanUser,
					sdk.NewAttribute(types2.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types2.AttributeKeyBanUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				),
			},
			expUsers: []string{
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
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

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err := handler.BanUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}

			var users []string
			suite.k.IterateSubspaceBannedUsers(suite.ctx, test.msg.SubspaceID, func(index int64, user string) (stop bool) {
				users = append(users, user)
				return false
			})
			suite.Require().Equal(test.expUsers, users)
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_UnblockUser() {
	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types2.MsgUnbanUser
		expErr    bool
		expEvents sdk.Events
		expUsers  []string
	}{
		{
			name: "Non existing subspace returns error",
			msg: types2.NewMsgUnbanUser(
				"123",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Not found user returns error",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgUnbanUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "Valid block works properly",
			store: func(ctx sdk.Context) {
				subspace := types2.NewSubspace(
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					"test",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					types2.SubspaceTypeOpen,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				err := suite.k.SaveSubspace(ctx, subspace, subspace.Owner)
				suite.Require().NoError(err)

				err = suite.k.BanUserInSubspace(ctx, subspace.ID, "cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5", subspace.Owner)
				suite.Require().NoError(err)
			},
			msg: types2.NewMsgUnbanUser(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types2.EventTypeUnbanUser,
					sdk.NewAttribute(types2.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types2.AttributeKeyUnbannedUser, "cosmos1mtanzwyk5p23haky8r6n4gxu7ypv0tlx9dgnk5"),
				),
			},
			expUsers: nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store(suite.ctx)
			}

			handler := keeper2.NewMsgServerImpl(suite.k)
			_, err := handler.UnbanUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())
			}

			var users []string
			suite.k.IterateSubspaceBannedUsers(suite.ctx, test.msg.SubspaceID, func(index int64, user string) (stop bool) {
				users = append(users, user)
				return false
			})
			suite.Require().Equal(test.expUsers, users)
		})
	}
}
