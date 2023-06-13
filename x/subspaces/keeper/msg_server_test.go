package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v5/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestMsgServer_CreateSubspace() {
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
					types.GetTreasuryAddress(1).String(),
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
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
					types.GetTreasuryAddress(1).String(),
					"cosmos17qcf9sv5yk0ly5vt3ztev70nwf6c5sprkwfh8t",
					"cosmos18atyyv6zycryhvnhpr2mjxgusdcah6kdpkffq0",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
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
				))
				suite.Require().NoError(err)
			},
			msg: types.NewMsgCreateSubspace(
				"Another subspace",
				"This is a second test subspace",
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
					types.GetTreasuryAddress(2).String(),
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
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

func (suite *KeeperTestSuite) TestMsgServer_EditSubspace() {
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
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
					nil,
				))
			},
			msg: types.NewMsgEditSubspace(
				1,
				types.DoNotModify,
				"This is a new description",
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
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
					nil,
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
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
					nil,
				))
			},
			msg: types.NewMsgEditSubspace(
				1,
				"",
				"This is a new description",
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
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
					nil,
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
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
					nil,
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					types.NewPermissions(types.PermissionEditSubspace),
				)
			},
			msg: types.NewMsgEditSubspace(
				1,
				"This is a new name",
				"This is a new description",
				"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
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
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					blockTime,
					nil,
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

func (suite *KeeperTestSuite) TestMsgServer_DeleteSubspace() {
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
					nil,
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
					nil,
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
					nil,
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					types.NewPermissions(types.PermissionDeleteSubspace),
				)
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

func (suite *KeeperTestSuite) TestMsgServer_CreateSection() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		msg         *types.MsgCreateSection
		shouldErr   bool
		expResponse *types.MsgCreateSectionResponse
		expEvent    sdk.Events
		check       func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgCreateSection(
				1,
				"Test section",
				"This is a test section",
				0,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "missing parent section returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgCreateSection(
				1,
				"Test section",
				"This is a test section",
				1,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Child section",
					"",
				))
			},
			msg: types.NewMsgCreateSection(
				1,
				"Test section",
				"This is a test section",
				1,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "missing next section id returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.DeleteNextSectionID(ctx, 1)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Child section",
					"",
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionManageSections),
				)
			},
			msg: types.NewMsgCreateSection(
				1,
				"Test section",
				"This is a test section",
				1,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "invalid data returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextSectionID(ctx, 1, 2)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Child section",
					"",
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionManageSections),
				)
			},
			msg: types.NewMsgCreateSection(
				1,
				"",
				"This is a test section",
				1,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "section is created properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextSectionID(ctx, 1, 2)

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Child section",
					"",
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionManageSections),
				)
			},
			msg: types.NewMsgCreateSection(
				1,
				"Test section",
				"This is a test section",
				0,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr:   false,
			expResponse: &types.MsgCreateSectionResponse{SectionID: 2},
			expEvent: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgCreateSection{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4"),
				),
				sdk.NewEvent(
					types.EventTypeCreateSection,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeySectionID, "2"),
				),
			},
			check: func(ctx sdk.Context) {
				// Check the next section id
				storedID, err := suite.k.GetNextSectionID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(3), storedID)

				// Check the section data
				stored, found := suite.k.GetSection(ctx, 1, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSection(
					1,
					2,
					0,
					"Test section",
					"This is a test section",
				), stored)
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

			msgServer := keeper.NewMsgServerImpl(suite.k)
			res, err := msgServer.CreateSection(sdk.WrapSDKContext(ctx), tc.msg)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
				suite.Require().Equal(tc.expEvent, ctx.EventManager().Events())
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_EditSection() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgEditSection
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgEditSection(
				1,
				1,
				"Edited section",
				"This is an edited section",
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "non existing section returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgEditSection(
				1,
				1,
				"Edited section",
				"This is an edited section",
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Child section",
					"",
				))
			},
			msg: types.NewMsgEditSection(
				1,
				1,
				"Edited section",
				"This is an edited section",
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "invalid update data returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Child section",
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionManageSections),
				)
			},
			msg: types.NewMsgEditSection(
				1,
				1,
				"",
				"This is an edited section",
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "section is updated properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Child section",
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionManageSections),
				)
			},
			msg: types.NewMsgEditSection(
				1,
				1,
				"Edited section",
				"This is an edited section",
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgEditSection{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4"),
				),
				sdk.NewEvent(
					types.EventTypeEditSection,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeySectionID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetSection(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSection(
					1,
					1,
					0,
					"Edited section",
					"This is an edited section",
				), stored)
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

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.EditSection(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_MoveSection() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgMoveSection
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgMoveSection(
				1,
				1,
				0,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "non existing section returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgMoveSection(
				1,
				1,
				0,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "non existing destination section returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					2,
					1,
					"Child section",
					"",
				))
			},
			msg: types.NewMsgMoveSection(
				1,
				2,
				3,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveSection(ctx, types.NewSection(1, 1, 0, "Child section", ""))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					2,
					1,
					"Child section",
					"",
				))
			},
			msg: types.NewMsgMoveSection(
				1,
				2,
				0,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "moving section to be its own parent returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					2,
					1,
					"Child section",
					"",
				))
			},
			msg: types.NewMsgMoveSection(
				1,
				2,
				2,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "moving section to create a circular path returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				// Create the following subspaces sections
				//  A
				//  |
				//  B - C
				//
				// We will then move A to be a child of C to create the following path
				//   A
				//  /  \
				//  B - C
				suite.k.SaveSection(ctx, types.NewSection(1, 1, types.RootSectionID, "A", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 2, 1, "B", ""))
				suite.k.SaveSection(ctx, types.NewSection(1, 3, 2, "C", ""))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionManageSections),
				)
			},
			msg: types.NewMsgMoveSection(
				1,
				1,
				3,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "section is moved properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveSection(ctx, types.NewSection(1, 1, 0, "Child section", ""))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					2,
					1,
					"Child section",
					"This is child section",
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionManageSections),
				)
			},
			msg: types.NewMsgMoveSection(
				1,
				2,
				0,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgMoveSection{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4"),
				),
				sdk.NewEvent(
					types.EventTypeMoveSection,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeySectionID, "2"),
				),
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetSection(ctx, 1, 2)
				suite.Require().True(found)
				suite.Require().Equal(types.NewSection(
					1,
					2,
					0,
					"Child section",
					"This is child section",
				), stored)
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

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.MoveSection(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_DeleteSection() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgDeleteSection
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgDeleteSection(
				1,
				1,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "non existing section returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgDeleteSection(
				1,
				1,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "user without permission returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				))
			},
			msg: types.NewMsgDeleteSection(
				1,
				1,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: true,
		},
		{
			name: "section is deleted properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace with another name and owner",
					"This is a test subspace with a changed description",
					"cosmos1fgppppwfjszpts4shpsfv7n2xtchcdwhycuvvm",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
					types.NewPermissions(types.PermissionManageSections),
				)
			},
			msg: types.NewMsgDeleteSection(
				1,
				1,
				"cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgDeleteSection{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1wq7mruftxd03qrrf9f7xnnzyqda9rkq5sshnr4"),
				),
				sdk.NewEvent(
					types.EventTypeDeleteSection,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeySectionID, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				result := suite.k.HasSection(ctx, 1, 1)
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

			msgServer := keeper.NewMsgServerImpl(suite.k)
			_, err := msgServer.DeleteSection(sdk.WrapSDKContext(ctx), tc.msg)
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

func (suite *KeeperTestSuite) TestMsgServer_CreateUserGroup() {
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
				0,
				"group",
				"description",
				types.NewPermissions(types.PermissionEditSubspace),
				nil,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "non existing section returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				1,
				"group",
				"description",
				types.NewPermissions(types.PermissionEditSubspace),
				nil,
				"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
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
					nil,
				))
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				0,
				"group",
				"description",
				types.NewPermissions(types.PermissionEditSubspace),
				nil,
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
					nil,
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1y4emx0mm4ncva9mnv9yvjrm7nrq3psvmwhk9ll",
					types.NewPermissions(types.PermissionManageGroups),
				)
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				0,
				"group",
				"description",
				types.NewPermissions(types.PermissionEditSubspace),
				nil,
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
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SetNextGroupID(ctx, 1, 2)

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.CombinePermissions(types.PermissionManageGroups, types.PermissionSetPermissions),
				)
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				0,
				"another group",
				"another description",
				types.NewPermissions("invalid"),
				nil,
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
					nil,
				))

				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
				suite.k.SetNextGroupID(ctx, 1, 2)

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.CombinePermissions(types.PermissionManageGroups, types.PermissionSetPermissions),
				)
			},
			msg: types.NewMsgCreateUserGroup(
				1,
				0,
				"another group",
				"another description",
				types.NewPermissions(types.PermissionEditSubspace),
				[]string{"cosmos16yhs7fgqnf6fjm4tftv66g2smtmee62wyg780l"},
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
				sdk.NewEvent(
					types.EventTypeAddUserToGroup,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUserGroupID, "2"),
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos16yhs7fgqnf6fjm4tftv66g2smtmee62wyg780l"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasUserGroup(ctx, 1, 1))
				suite.Require().True(suite.k.HasUserGroup(ctx, 1, 2))
				suite.Require().True(suite.k.IsMemberOfGroup(ctx, 1, 2, "cosmos16yhs7fgqnf6fjm4tftv66g2smtmee62wyg780l"))
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

func (suite *KeeperTestSuite) TestMsgServer_EditUserGroup() {
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
					nil,
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionManageGroups),
				)
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionManageGroups),
				)
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
					0,
					1,
					"Admins",
					"Group of the admins of th subspace",
					types.NewPermissions(types.PermissionEditSubspace),
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

func (suite *KeeperTestSuite) TestMsgServer_MoveUserGroup() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgMoveUserGroup
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error",
			msg: types.NewMsgMoveUserGroup(
				1,
				1,
				1,
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
					nil,
				))
			},
			msg: types.NewMsgMoveUserGroup(
				1,
				1,
				1,
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "non existing destination section returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			msg: types.NewMsgMoveUserGroup(
				1,
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error - PermissionManageGroups inside current section",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			msg: types.NewMsgMoveUserGroup(
				1,
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error - PermissionManageGroups inside destination section",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionManageGroups),
				)
			},
			msg: types.NewMsgMoveUserGroup(
				1,
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "no permission returns error - PermissionSetPermissions inside destination section",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionManageGroups),
				)

				suite.k.SetUserPermissions(ctx,
					1,
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionManageGroups),
				)
			},
			msg: types.NewMsgMoveUserGroup(
				1,
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "existing group is moved properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SaveSection(ctx, types.NewSection(
					1,
					1,
					0,
					"Test section",
					"This is a test section",
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionManageGroups),
				)
				suite.k.SetUserPermissions(ctx,
					1,
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionManageGroups),
				)
				suite.k.SetUserPermissions(ctx,
					1,
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionSetPermissions),
				)
			},
			msg: types.NewMsgMoveUserGroup(
				1,
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgMoveUserGroup{})),
					sdk.NewAttribute(sdk.AttributeKeySender, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53"),
				),
				sdk.NewEvent(
					types.EvenTypeMoveUserGroup,
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
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
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
			_, err := service.MoveUserGroup(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_SetUserGroupPermissions() {
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
				types.NewPermissions(types.PermissionSetPermissions),
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
					nil,
				))
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.NewPermissions(types.PermissionSetPermissions),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.NewPermissions(types.PermissionSetPermissions),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionSetPermissions),
				)
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.NewPermissions("invalid"),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionSetPermissions),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.NewPermissions(types.PermissionEverything),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionSetPermissions),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.NewPermissions(types.PermissionEverything),
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

				suite.Require().Equal(types.NewPermissions(types.PermissionEverything), group.Permissions)
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionSetPermissions),
				)
			},
			msg: types.NewMsgSetUserGroupPermissions(
				1,
				1,
				types.NewPermissions(types.PermissionSetPermissions),
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

				suite.Require().Equal(types.NewPermissions(types.PermissionSetPermissions), group.Permissions)
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

func (suite *KeeperTestSuite) TestMsgServer_DeleteUserGroup() {
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
					nil,
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionManageGroups),
				)
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

func (suite *KeeperTestSuite) TestMsgServer_AddUserToGroup() {
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
					nil,
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionSetPermissions),
				)
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
				result := suite.k.IsMemberOfGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
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

func (suite *KeeperTestSuite) TestMsgServer_RemoveUserFromGroup() {
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
					nil,
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
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
					nil,
				))
				suite.k.SaveUserGroup(ctx, types.NewUserGroup(
					1,
					0,
					1,
					"Test group",
					"This is a test group",
					types.NewPermissions(types.PermissionEditSubspace),
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionSetPermissions),
				)
				suite.k.AddUserToGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
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
				result := suite.k.IsMemberOfGroup(ctx, 1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
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

func (suite *KeeperTestSuite) TestMsgServer_SetUserPermissions() {
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
				0,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
			),
			shouldErr: true,
		},
		{
			name: "section not found returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgSetUserPermissions(
				1,
				1,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewPermissions(types.PermissionEditSubspace),
				"cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm",
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
					nil,
				))
			},
			msg: types.NewMsgSetUserPermissions(
				1,
				0,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewPermissions(types.PermissionEditSubspace),
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
					nil,
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewPermissions(types.PermissionSetPermissions),
				)
			},
			msg: types.NewMsgSetUserPermissions(
				1,
				0,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewPermissions("invalid"),
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
					nil,
				))

				suite.k.SetUserPermissions(ctx,
					1,
					0,
					"cosmos17ua98rre5j9ce7hfude0y5y3rh4gtqkygm8hru",
					types.NewPermissions(types.PermissionSetPermissions),
				)
			},
			msg: types.NewMsgSetUserPermissions(
				1,
				0,
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
				types.NewPermissions(types.PermissionEditSubspace),
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
				permissions := suite.k.GetUserPermissions(ctx, 1, 0, "cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53")
				suite.Require().Equal(types.NewPermissions(types.PermissionEditSubspace), permissions)
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

func (suite *KeeperTestSuite) TestMsgServer_UpdateSubspaceFeeTokens() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgUpdateSubspaceFeeTokens
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "invalid authority return error",
			msg: types.NewMsgUpdateSubspaceFeeTokens(
				1,
				sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(10))),
				"invalid",
			),
			shouldErr: true,
		},
		{
			name: "subspace not found returns error",
			msg: types.NewMsgUpdateSubspaceFeeTokens(
				1,
				sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(10))),
				"authority",
			),
			shouldErr: true,
		},
		{
			name: "invalid allowed fee tokens returns error",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgUpdateSubspaceFeeTokens(
				1,
				sdk.Coins{{Denom: "minttoken", Amount: sdk.NewInt(-10)}},
				"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
			),
			shouldErr: true,
		},
		{
			name: "subspace is updated correctly",
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspace(ctx, types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			msg: types.NewMsgUpdateSubspaceFeeTokens(
				1,
				sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(10))),
				authtypes.NewModuleAddress("gov").String(),
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					sdk.EventTypeMessage,
					sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
					sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(&types.MsgUpdateSubspaceFeeTokens{})),
					sdk.NewAttribute(sdk.AttributeKeySender, authtypes.NewModuleAddress("gov").String()),
				),
				sdk.NewEvent(
					types.EventTypeUpdateSubspaceFeeToken,
					sdk.NewAttribute(types.AttributeKeySubspaceID, "1"),
					sdk.NewAttribute(types.AttributeKeyUser, authtypes.NewModuleAddress("gov").String()),
				),
			},
			check: func(ctx sdk.Context) {
				subspace, _ := suite.k.GetSubspace(ctx, 1)
				expected := types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					sdk.NewCoins(sdk.NewCoin("minttoken", sdk.NewInt(10))),
				)

				suite.Require().Equal(expected, subspace)
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
			_, err := service.UpdateSubspaceFeeTokens(sdk.WrapSDKContext(ctx), tc.msg)

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
