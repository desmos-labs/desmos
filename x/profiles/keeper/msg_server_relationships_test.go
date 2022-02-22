package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	"github.com/desmos-labs/desmos/v2/testutil"

	"github.com/desmos-labs/desmos/v2/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func (suite *KeeperTestSuite) TestMsgServer_CreateRelationship() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgCreateRelationship
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "returns an error if the relationship sender is blocked by the receiver",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"tc",
					0,
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocked)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				0,
			),
			shouldErr: true,
		},
		{
			name: "existing relationship returns error",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Recipient)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				0,
			),
			shouldErr: true,
		},
		{
			name: "new relationship is stored correctly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					0,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")))
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")))
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				0,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipCreated,
					sdk.NewAttribute(types.AttributeRelationshipSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "0"),
				),
			},
			check: func(ctx sdk.Context) {
				expected := []types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						0,
					),
				}
				suite.Require().Equal(expected, suite.k.GetAllRelationships(ctx))
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

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.CreateRelationship(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_DeleteRelationship() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgDeleteRelationship
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing relationship returns error",
			msg: types.NewMsgDeleteRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				2,
			),
			shouldErr: true,
		},
		{
			name: "existing relationship is removed properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					1,
				)
				store.Set(
					types.RelationshipsStoreKey(relationship.Creator, relationship.SubspaceID, relationship.Recipient),
					suite.cdc.MustMarshal(&relationship),
				)
			},
			msg: types.NewMsgDeleteRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				1,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipsDeleted,
					sdk.NewAttribute(types.AttributeRelationshipSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "1"),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().Empty(suite.k.GetAllRelationships(ctx))
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

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.DeleteRelationship(sdk.WrapSDKContext(ctx), tc.msg)

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
