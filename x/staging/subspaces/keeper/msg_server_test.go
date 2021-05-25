package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"time"
)

func (suite *KeeperTestsuite) TestMsgServer_CreateSubspace() {
	creationTime, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.NoError(err)

	tests := []struct {
		name           string
		blockTime      time.Time
		storedSubspace *types.Subspace
		msg            *types.MsgCreateSubspace
		expErr         bool
		expEvent       sdk.Event
	}{
		{
			name:      "subspace already exists returns error",
			blockTime: creationTime,
			storedSubspace: &types.Subspace{
				ID:           "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: creationTime,
			},
			msg: types.NewMsgCreateSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test2",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				true,
			),
			expErr: true,
		},
		{
			name:           "subspace saved successfully",
			blockTime:      creationTime,
			storedSubspace: nil,
			msg: types.NewMsgCreateSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"test2",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				true,
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeCreateSubspace,
				sdk.NewAttribute(types.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				sdk.NewAttribute(types.AttributeKeySubspaceName, "test2"),
				sdk.NewAttribute(types.AttributeKeySubspaceCreator, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, types.DefaultParams())
			suite.ctx = suite.ctx.WithBlockTime(test.blockTime)
			if test.storedSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.CreateSubspace(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				events := suite.ctx.EventManager().Events()
				suite.Len(events, 1)
				suite.Equal(test.expEvent, events[0])
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_EditSubspace() {
	creationTime, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.NoError(err)

	tests := []struct {
		name           string
		blockTime      time.Time
		storedSubspace *types.Subspace
		msg            *types.MsgEditSubspace
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			blockTime:      creationTime,
			storedSubspace: nil,
			msg: types.NewMsgEditSubspace(
				"1234",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"edited",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expErr: true,
		},
		{
			name:      "subspace edited successfully",
			blockTime: creationTime,
			storedSubspace: &types.Subspace{
				ID:           "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				Creator:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: creationTime,
			},
			msg: types.NewMsgEditSubspace(
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"edited",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeEditSubspace,
				sdk.NewAttribute(types.AttributeKeySubspaceID, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				sdk.NewAttribute(types.AttributeKeyNewOwner, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				sdk.NewAttribute(types.AttributeKeySubspaceName, "edited"),
			),
			expSubspace: types.Subspace{
				ID:           "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Name:         "edited",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: creationTime,
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, types.DefaultParams())
			suite.ctx = suite.ctx.WithBlockTime(test.blockTime)

			if test.storedSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.EditSubspace(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				events := suite.ctx.EventManager().Events()
				suite.Len(events, 1)
				suite.Equal(test.expEvent, events[0])

				editeSubspace, found := suite.k.GetSubspace(suite.ctx, test.expSubspace.ID)
				suite.True(found)
				suite.Equal(test.expSubspace, editeSubspace)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_AddAdmin() {
	tests := []struct {
		name           string
		storedSubspace *types.Subspace
		msg            *types.MsgAddAdmin
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			storedSubspace: nil,
			msg: types.NewMsgAddAdmin(
				"123",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "admin added successfully",
			storedSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Open:         false,
				CreationTime: time.Time{},
			},
			msg: types.NewMsgAddAdmin(
				"123",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeAddAdmin,
				sdk.NewAttribute(types.AttributeKeySubspaceID, "123"),
				sdk.NewAttribute(types.AttributeKeySubspaceNewAdmin, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			),
			expSubspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
				Open:         false,
				Admins:       []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.storedSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.AddAdmin(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				events := suite.ctx.EventManager().Events()
				suite.Len(events, 1)
				suite.Equal(test.expEvent, events[0])

				subspace, found := suite.k.GetSubspace(suite.ctx, test.expSubspace.ID)
				suite.True(found)
				suite.Equal(test.expSubspace, subspace)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_RemoveAdmin() {
	tests := []struct {
		name           string
		storedSubspace *types.Subspace
		msg            *types.MsgRemoveAdmin
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			storedSubspace: nil,
			msg: types.NewMsgRemoveAdmin(
				"123",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "admin removed successfully",
			storedSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
				Open:         false,
				Admins:       []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			},
			msg: types.NewMsgRemoveAdmin(
				"123",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeRemoveAdmin,
				sdk.NewAttribute(types.AttributeKeySubspaceID, "123"),
				sdk.NewAttribute(types.AttributeKeySubspaceRemovedAdmin, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			),
			expSubspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
				Open:         false,
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.storedSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.RemoveAdmin(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				events := suite.ctx.EventManager().Events()
				suite.Len(events, 1)
				suite.Equal(test.expEvent, events[0])

				subspace, found := suite.k.GetSubspace(suite.ctx, test.expSubspace.ID)
				suite.True(found)
				suite.Equal(test.expSubspace, subspace)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_RegisterUser() {
	tests := []struct {
		name           string
		storedSubspace *types.Subspace
		msg            *types.MsgRegisterUser
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			storedSubspace: nil,
			msg: types.NewMsgRegisterUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "user registered successfully",
			storedSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Open:         false,
				CreationTime: time.Time{},
			},
			msg: types.NewMsgRegisterUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeRegisterUser,
				sdk.NewAttribute(types.AttributeKeyRegisteredUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				sdk.NewAttribute(types.AttributeKeySubspaceID, "123"),
			),
			expSubspace: types.Subspace{
				ID:              "123",
				Name:            "test",
				Owner:           "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:         "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime:    time.Time{},
				Open:            false,
				RegisteredUsers: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.storedSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.RegisterUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				events := suite.ctx.EventManager().Events()
				suite.Len(events, 1)
				suite.Equal(test.expEvent, events[0])

				subspace, found := suite.k.GetSubspace(suite.ctx, test.expSubspace.ID)
				suite.True(found)
				suite.Equal(test.expSubspace, subspace)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_UnregisterUser() {
	tests := []struct {
		name           string
		storedSubspace *types.Subspace
		msg            *types.MsgUnregisterUser
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			storedSubspace: nil,
			msg: types.NewMsgUnregisterUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "user unregistered successfully",
			storedSubspace: &types.Subspace{
				ID:              "123",
				Name:            "test",
				Owner:           "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:         "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime:    time.Time{},
				Open:            false,
				RegisteredUsers: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			},
			msg: types.NewMsgUnregisterUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeUnregisterUser,
				sdk.NewAttribute(types.AttributeKeyUnregisteredUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				sdk.NewAttribute(types.AttributeKeySubspaceID, "123"),
			),
			expSubspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
				Open:         false,
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.storedSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.UnregisterUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				events := suite.ctx.EventManager().Events()
				suite.Len(events, 1)
				suite.Equal(test.expEvent, events[0])

				subspace, found := suite.k.GetSubspace(suite.ctx, test.expSubspace.ID)
				suite.True(found)
				suite.Equal(test.expSubspace, subspace)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_BlockUser() {
	tests := []struct {
		name           string
		storedSubspace *types.Subspace
		msg            *types.MsgBanUser
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			storedSubspace: nil,
			msg: types.NewMsgBanUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "user blocked successfully",
			storedSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Open:         false,
				CreationTime: time.Time{},
			},
			msg: types.NewMsgBanUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeBanUser,
				sdk.NewAttribute(types.AttributeKeyBanUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				sdk.NewAttribute(types.AttributeKeySubspaceID, "123"),
			),
			expSubspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
				Open:         false,
				BannedUsers:  []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.storedSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.BanUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				events := suite.ctx.EventManager().Events()
				suite.Len(events, 1)
				suite.Equal(test.expEvent, events[0])

				subspace, found := suite.k.GetSubspace(suite.ctx, test.expSubspace.ID)
				suite.True(found)
				suite.Equal(test.expSubspace, subspace)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_UnblockUser() {
	tests := []struct {
		name           string
		storedSubspace *types.Subspace
		msg            *types.MsgUnbanUser
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			storedSubspace: nil,
			msg: types.NewMsgUnbanUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: true,
		},
		{
			name: "user unblock successfully",
			storedSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
				Open:         false,
				BannedUsers:  []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"},
			},
			msg: types.NewMsgUnbanUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeUnbanUser,
				sdk.NewAttribute(types.AttributeKeyUnbannedUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				sdk.NewAttribute(types.AttributeKeySubspaceID, "123"),
			),
			expSubspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
				Open:         false,
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.storedSubspace != nil {
				suite.k.SaveSubspace(suite.ctx, *test.storedSubspace)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.UnbanUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
				events := suite.ctx.EventManager().Events()
				suite.Len(events, 1)
				suite.Equal(test.expEvent, events[0])

				subspace, found := suite.k.GetSubspace(suite.ctx, test.expSubspace.ID)
				suite.True(found)
				suite.Equal(test.expSubspace, subspace)
			}
		})
	}
}
