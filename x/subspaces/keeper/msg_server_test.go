package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *KeeperTestsuite) TestMsgServer_CreateSubspace() {
	blockTime := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		msg         *types.MsgCreateSubspace
		shouldErr   bool
		expResponse *types.MsgCreateSubspaceResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "subspace id not set returns error",
			msg: types.NewMsgCreateSubspace(
				"Test subspace",
				"This is a test subspace",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			),
			shouldErr: true,
		},
		{
			name: "invalid subspace returns error",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.SubspaceIDKey, types.GetSubspaceIDBytes(1))
			},
			msg: types.NewMsgCreateSubspace(
				"",
				"This is a test subspace",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			),
			shouldErr: true,
		},
		{
			name: "first subspace is created properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.SubspaceIDKey, types.GetSubspaceIDBytes(1))
			},
			msg: types.NewMsgCreateSubspace(
				"Test subspace",
				"This is a test subspace",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			),
			shouldErr:   false,
			expResponse: &types.MsgCreateSubspaceResponse{SubspaceID: 1},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateSubspace{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
				),
				sdk.NewEvent(
					types.EventTypeCreateSubspace,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeySubspaceName, "Test subspace"),
					sdk.NewAttribute(types.AttributeKeySubspaceCreator, "cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69"),
					sdk.NewAttribute(types.AttributeKeyCreationTime, "2020-01-01T12:00:00Z"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the subspace is stored
				subspace, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), subspace)

				// Make sure the default group has been created
				group, found := suite.k.GetUserGroup(ctx, 1, 0)
				suite.Require().True(found)
				suite.Require().Equal(types.DefaultUserGroup(1), group)

				// Make sure the subspace id has increased
				store := ctx.KVStore(suite.storeKey)
				id := types.GetSubspaceIDFromBytes(store.Get(types.SubspaceIDKey))
				suite.Require().Equal(uint64(2), id)
			},
		},
		{
			name: "subspace with three different addresses is created properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.SubspaceIDKey, types.GetSubspaceIDBytes(1))
			},
			msg: types.NewMsgCreateSubspace(
				"Test subspace",
				"This is a test subspace",
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				"cosmos17qcf9sv5yk0ly5vt3ztev70nwf6c5sprkwfh8t",
				"cosmos18atyyv6zycryhvnhpr2mjxgusdcah6kdpkffq0",
			),
			shouldErr:   false,
			expResponse: &types.MsgCreateSubspaceResponse{SubspaceID: 1},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateSubspace{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos18atyyv6zycryhvnhpr2mjxgusdcah6kdpkffq0"),
				),
				sdk.NewEvent(
					types.EventTypeCreateSubspace,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeySubspaceName, "Test subspace"),
					sdk.NewAttribute(types.AttributeKeySubspaceCreator, "cosmos18atyyv6zycryhvnhpr2mjxgusdcah6kdpkffq0"),
					sdk.NewAttribute(types.AttributeKeyCreationTime, "2020-01-01T12:00:00Z"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the subspace is stored
				subspace, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos17qcf9sv5yk0ly5vt3ztev70nwf6c5sprkwfh8t",
					"cosmos18atyyv6zycryhvnhpr2mjxgusdcah6kdpkffq0",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), subspace)

				// Make sure the subspace id has increased
				store := ctx.KVStore(suite.storeKey)
				id := types.GetSubspaceIDFromBytes(store.Get(types.SubspaceIDKey))
				suite.Require().Equal(uint64(2), id)
			},
		},
		{
			name: "subspace has correct id when another subspace already exists",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				store.Set(types.SubspaceIDKey, types.GetSubspaceIDBytes(1))

				// Run the fist subspace creation message
				msgServer := keeper.NewMsgServerImpl(suite.k)
				_, err := msgServer.CreateSubspace(sdk.WrapSDKContext(ctx), types.NewMsgCreateSubspace(
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
				))
				suite.Require().NoError(err)
			},
			msg: types.NewMsgCreateSubspace(
				"Another subspace",
				"This is a second test subspace",
				"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
				"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
				"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
			),
			shouldErr:   false,
			expResponse: &types.MsgCreateSubspaceResponse{SubspaceID: 2},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateSubspace{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll"),
				),
				sdk.NewEvent(
					types.EventTypeCreateSubspace,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "2"),
					sdk.NewAttribute(types.AttributeKeySubspaceName, "Another subspace"),
					sdk.NewAttribute(types.AttributeKeySubspaceCreator, "cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll"),
					sdk.NewAttribute(types.AttributeKeyCreationTime, "2020-01-01T12:00:00Z"),
				),
			},
			check: func(ctx sdk.Context) {
				// Make sure the subspace is stored
				subspace, found := suite.k.GetSubspace(ctx, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					2,
					"Another subspace",
					"This is a second test subspace",
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), subspace)

				// Make sure the subspace id has increased
				store := ctx.KVStore(suite.storeKey)
				id := types.GetSubspaceIDFromBytes(store.Get(types.SubspaceIDKey))
				suite.Require().Equal(uint64(3), id)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			ctx = ctx.WithBlockTime(blockTime)

			if tc.store != nil {
				tc.store(ctx)
			}

			// Reset any event that might have been emitted during the setup
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			res, err := service.CreateSubspace(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_EditSubspace() {
	blockTime := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgEditSubspace
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgEditSubspace(
				1,
				types.DoNotModify,
				"This is a new description",
				types.DoNotModify,
				types.DoNotModify,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().False(found)
			},
		},
		{
			name: "missing permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
				))
			},
			msg: types.NewMsgEditSubspace(
				1,
				types.DoNotModify,
				"This is a new description",
				types.DoNotModify,
				types.DoNotModify,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
			check: func(ctx sdk.Context) {
				subspace, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
				), subspace)
			},
		},
		{
			name: "invalid update returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
				))
			},
			msg: types.NewMsgEditSubspace(
				1,
				"",
				"This is a new description",
				types.DoNotModify,
				types.DoNotModify,
				"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
			),
			shouldErr: true,
			check: func(ctx sdk.Context) {
				subspace, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
				), subspace)
			},
		},
		{
			name: "existing subspace is updated correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionChangeInfo)
			},
			msg: types.NewMsgEditSubspace(
				1,
				"This is a new name",
				"This is a new description",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgEditSubspace{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				),
				sdk.NewEvent(
					types.EventTypeEditSubspace,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				subspace, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"This is a new name",
					"This is a new description",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
				), subspace)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			ctx = ctx.WithBlockTime(blockTime)

			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.EditSubspace(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_DeleteSubspace() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgDeleteSubspace
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name:      "subspace not found returns error",
			msg:       types.NewMsgDeleteSubspace(1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			shouldErr: true,
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().False(found)
			},
		},
		{
			name: "missing permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg:       types.NewMsgDeleteSubspace(1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			shouldErr: true,
			check: func(ctx sdk.Context) {
				subspace, found := suite.k.GetSubspace(ctx, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				), subspace)
			},
		},
		{
			name: "existing subspace is deleted correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionDeleteSubspace)
			},
			msg:       types.NewMsgDeleteSubspace(1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgDeleteSubspace{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				),
				sdk.NewEvent(
					types.EventTypeDeleteSubspace,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				exists := suite.k.HasSubspace(ctx, 1)
				suite.Require().False(exists)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.DeleteSubspace(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_CreateUserGroup() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		msg         *types.MsgCreateUserGroup
		shouldErr   bool
		expResponse *types.MsgCreateUserGroupResponse
		expEvents   sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgCreateUserGroup(
				1,
				"group",
				"description",
				types.PermissionWrite,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "user without PermissionManageGroups returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				"group",
				"description",
				types.PermissionWrite,
				"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
			),
			shouldErr: true,
		},
		{
			name: "user without PermissionSetPermissions returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionManageGroups)
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				"group",
				"description",
				types.PermissionWrite,
				"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
			),
			shouldErr: true,
		},
		{
			name: "invalid permissions value returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
				suite.k.SetNextGroupID(ctx, 1, 2)

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.CombinePermissions(types.PermissionManageGroups, types.PermissionSetPermissions))
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				"another group",
				"another description",
				256,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "group is created properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
				suite.k.SetNextGroupID(ctx, 1, 2)

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.CombinePermissions(types.PermissionManageGroups, types.PermissionSetPermissions))
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				"another group",
				"another description",
				types.PermissionWrite,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr:   false,
			expResponse: &types.MsgCreateUserGroupResponse{GroupID: 2},
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateUserGroup{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeCreateUserGroup,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "2"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasUserGroup(ctx, 1, 1))
				suite.Require().True(suite.k.HasUserGroup(ctx, 1, 2))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			res, err := service.CreateUserGroup(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_EditUserGroup() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgEditUserGroup
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgEditUserGroup(
				1,
				1,
				"Test group",
				"This is a test group",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "group not found returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgEditUserGroup(
				1,
				1,
				"Test group",
				"This is a test group",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			msg: types.NewMsgEditUserGroup(
				1,
				1,
				"Test group new name",
				"This is a test group with a new name",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "invalid update returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionManageGroups)
			},
			msg: types.NewMsgEditUserGroup(
				1,
				1,
				"",
				"",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "existing group is edited properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionManageGroups)
			},
			msg: types.NewMsgEditUserGroup(
				1,
				1,
				"Admins",
				"Group of the admins of th subspace",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgEditUserGroup{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeEditUserGroup,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				group, found := suite.k.GetUserGroup(ctx, 1, 1)
				suite.Require().True(found)

				suite.Require().Equal(types.NewUserGroup(
					1,
					1,
					"Admins",
					"Group of the admins of th subspace",
					types.PermissionWrite,
				), group)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.EditUserGroup(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_SetUserGroupPermissions() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgSetUserGroupPermissions
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.PermissionSetPermissions,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "group not found returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.PermissionSetPermissions,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.PermissionSetPermissions,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "invalid permission value returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionSetPermissions)
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				256,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "setting the permissions for a group you are part of returns error if not owner",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionSetPermissions,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, sdkAddr)
				suite.Require().NoError(err)
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.PermissionEverything,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "setting the permissions for a group you are part of does not return error if owner",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionSetPermissions,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, sdkAddr)
				suite.Require().NoError(err)
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.PermissionEverything,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgSetUserGroupPermissions{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
				),
				sdk.NewEvent(
					types.EventTypeSetUserGroupPermissions,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				group, found := suite.k.GetUserGroup(ctx, 1, 1)
				suite.Require().True(found)

				suite.Require().Equal(types.PermissionEverything, group.Permissions)
			},
		},
		{
			name: "group permissions are updated correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionSetPermissions)
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.PermissionSetPermissions,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgSetUserGroupPermissions{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeSetUserGroupPermissions,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				group, found := suite.k.GetUserGroup(ctx, 1, 1)
				suite.Require().True(found)

				suite.Require().Equal(types.PermissionSetPermissions, group.Permissions)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.SetUserGroupPermissions(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_DeleteUserGroup() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgDeleteUserGroup
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgDeleteUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "group not found returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgDeleteUserGroup(
				1,
				1,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			msg: types.NewMsgDeleteUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "existing group is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionManageGroups)
			},
			msg: types.NewMsgDeleteUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgDeleteUserGroup{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeDeleteUserGroup,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				hasGroup := suite.k.HasUserGroup(ctx, 1, 1)
				suite.Require().False(hasGroup)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.DeleteUserGroup(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_AddUserToGroup() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgAddUserToUserGroup
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgAddUserToUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "group not found returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgAddUserToUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			msg: types.NewMsgAddUserToUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "user already part of group returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, sdkAddr)
				suite.Require().NoError(err)
			},
			msg: types.NewMsgAddUserToUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "user not part of group is added correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionSetPermissions)
			},
			msg: types.NewMsgAddUserToUserGroup(
				1,
				1,
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgAddUserToUserGroup{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeAddUserToGroup,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "1"),
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"),
				),
			},
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				result := suite.k.IsMemberOfGroup(ctx, 1, 1, sdkAddr)
				suite.Require().True(result)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.AddUserToUserGroup(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_RemoveUserFromGroup() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgRemoveUserFromUserGroup
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgRemoveUserFromUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "group not found returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgRemoveUserFromUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			msg: types.NewMsgRemoveUserFromUserGroup(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "user not part of group returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))
			},
			msg: types.NewMsgRemoveUserFromUserGroup(
				1,
				1,
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "user part of group is removed correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					1,
					"Test group",
					"This is a test group",
					types.PermissionWrite,
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionSetPermissions)

				sdkAddr, err = sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				err = suite.k.AddUserToGroup(ctx, 1, 1, sdkAddr)
				suite.Require().NoError(err)
			},
			msg: types.NewMsgRemoveUserFromUserGroup(
				1,
				1,
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgRemoveUserFromUserGroup{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EventTypeRemoveUserFromGroup,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "1"),
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"),
				),
			},
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				suite.Require().NoError(err)

				result := suite.k.IsMemberOfGroup(ctx, 1, 1, sdkAddr)
				suite.Require().False(result)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.RemoveUserFromUserGroup(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestsuite) TestMsgServer_SetUserPermissions() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgSetUserPermissions
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "subspace not found returns error",
			msg: types.NewMsgSetUserPermissions(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.PermissionWrite,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			msg: types.NewMsgSetUserPermissions(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.PermissionWrite,
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
			),
			shouldErr: true,
		},
		{
			name: "invalid permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionSetPermissions)
			},
			msg: types.NewMsgSetUserPermissions(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				256,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "permissions are set correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos17ua98rre5j9ce7hfude0y5y3rh4gtqkygm8hru")
				suite.Require().NoError(err)
				suite.k.SetUserPermissions(ctx, 1, sdkAddr, types.PermissionSetPermissions)
			},
			msg: types.NewMsgSetUserPermissions(
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.PermissionWrite,
				"cosmos17ua98rre5j9ce7hfude0y5y3rh4gtqkygm8hru",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgSetUserPermissions{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos17ua98rre5j9ce7hfude0y5y3rh4gtqkygm8hru"),
				),
				sdk.NewEvent(
					types.EventTypeSetUserPermissions,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
			},
			check: func(ctx sdk.Context) {
				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().NoError(err)
				permissions := suite.k.GetUserPermissions(ctx, 1, sdkAddr)
				suite.Require().Equal(types.PermissionWrite, permissions)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			// Run the message
			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.SetUserPermissions(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
