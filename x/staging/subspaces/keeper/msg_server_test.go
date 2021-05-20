package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/keeper"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"time"
)

func (suite *KeeperTestsuite) TestMsgServer_CreateSubspace() {
	tests := []struct {
		name           string
		storedSubspace *types.Subspace
		msg            *types.MsgCreateSubspace
		expErr         bool
		expEvent       sdk.Event
	}{
		{
			name: "subspace already exists returns error",
			storedSubspace: &types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
			},
			msg: types.NewMsgCreateSubspace(
				"123",
				"test2",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				true,
			),
			expErr: true,
		},
		{
			name:           "subspace saved successfully",
			storedSubspace: nil,
			msg: types.NewMsgCreateSubspace(
				"123",
				"test2",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				true,
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeCreateSubspace,
				sdk.NewAttribute(types.AttributeKeySubspaceID, "123"),
				sdk.NewAttribute(types.AttributeKeySubspaceName, "test2"),
				sdk.NewAttribute(types.AttributeKeySubspaceCreator, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			),
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
	tests := []struct {
		name           string
		storedSubspace *types.Subspace
		msg            *types.MsgEditSubspace
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
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
			name: "subspace edited successfully",
			storedSubspace: &types.Subspace{
				ID:           "1234",
				Name:         "test",
				Owner:        "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				CreationTime: time.Time{},
			},
			msg: types.NewMsgEditSubspace(
				"1234",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"edited",
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeEditSubspace,
				sdk.NewAttribute(types.AttributeKeySubspaceID, "1234"),
				sdk.NewAttribute(types.AttributeKeyNewOwner, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				sdk.NewAttribute(types.AttributeKeySubspaceName, "edited"),
			),
			expSubspace: types.Subspace{
				ID:           "1234",
				Name:         "edited",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
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
				Admins:       types.Users{Users: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}},
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
				Admins:       types.Users{Users: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}},
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
				RegisteredUsers: types.Users{Users: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}},
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
				RegisteredUsers: types.Users{Users: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}},
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
		msg            *types.MsgBlockUser
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			storedSubspace: nil,
			msg: types.NewMsgBlockUser(
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
			msg: types.NewMsgBlockUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeBlockUser,
				sdk.NewAttribute(types.AttributeKeyBlockedUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				sdk.NewAttribute(types.AttributeKeySubspaceID, "123"),
			),
			expSubspace: types.Subspace{
				ID:           "123",
				Name:         "test",
				Owner:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Creator:      "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				CreationTime: time.Time{},
				Open:         false,
				BlockedUsers: types.Users{Users: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}},
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
			_, err := handler.BlockUser(sdk.WrapSDKContext(suite.ctx), test.msg)

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
		msg            *types.MsgUnblockUser
		expErr         bool
		expEvent       sdk.Event
		expSubspace    types.Subspace
	}{
		{
			name:           "subspace doesn't exists returns error",
			storedSubspace: nil,
			msg: types.NewMsgUnblockUser(
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
				BlockedUsers: types.Users{Users: []string{"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"}},
			},
			msg: types.NewMsgUnblockUser(
				"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				"123",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expErr: false,
			expEvent: sdk.NewEvent(
				types.EventTypeUnblockUser,
				sdk.NewAttribute(types.AttributeKeyUnblockedUser, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
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
			_, err := handler.UnblockUser(sdk.WrapSDKContext(suite.ctx), test.msg)

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
