package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgBlockUser() {
	tests := []struct {
		name      string
		msg       *types.MsgBlockUser
		stored    []types.UserBlock
		expErr    bool
		expEvents sdk.Events
		expBlocks []types.UserBlock
	}{
		{
			name: "Existing block returns error",
			stored: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					suite.testData.otherUser,
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types.NewMsgBlockUser(
				suite.testData.user,
				suite.testData.otherUser,
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: true,
		},
		{
			name:   "Block has been saved correctly",
			stored: nil,
			msg: types.NewMsgBlockUser(
				suite.testData.user,
				suite.testData.otherUser,
				"reason",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeBlockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, suite.testData.user),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, suite.testData.otherUser),
					sdk.NewAttribute(types.AttributeKeySubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					sdk.NewAttribute(types.AttributeKeyUserBlockReason, "reason"),
				),
			},
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					suite.testData.user,
					suite.testData.otherUser,
					"reason",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {

			profile := suite.CreateProfileFromAddress(suite.testData.user)
			otherProfile := suite.CreateProfileFromAddress(suite.testData.otherUser)

			err := suite.k.StoreProfile(suite.ctx, profile)
			suite.Require().NoError(err)

			err = suite.k.StoreProfile(suite.ctx, otherProfile)
			suite.Require().NoError(err)

			for _, block := range test.stored {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err = service.BlockUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				blocks := suite.k.GetUserBlocks(suite.ctx, test.msg.Blocker)
				suite.Require().Equal(test.expBlocks, blocks)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgUnblockUser() {
	tests := []struct {
		name        string
		storedBlock []types.UserBlock
		msg         *types.MsgUnblockUser
		expErr      bool
		expEvents   sdk.Events
		expBlocks   []types.UserBlock
	}{
		{
			name:        "Invalid block returns error",
			storedBlock: []types.UserBlock{},
			msg:         types.NewMsgUnblockUser(suite.testData.user, "blocked", "subspace"),
			expErr:      true,
		},
		{
			name: "Existing block is removed and leaves empty array",
			storedBlock: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, "blocked", "reason", "subspace"),
			},
			msg:    types.NewMsgUnblockUser(suite.testData.user, "blocked", "subspace"),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnblockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, suite.testData.user),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "blocked"),
					sdk.NewAttribute(types.AttributeKeySubspace, "subspace"),
				),
			},
			expBlocks: nil,
		},
		{
			name: "Existing block is removed and leaves non empty array",
			storedBlock: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, "blocked", "reason", "subspace"),
				types.NewUserBlock(suite.testData.otherUser, "blocked", "reason", "other_subspace"),
			},
			msg:    types.NewMsgUnblockUser(suite.testData.user, "blocked", "subspace"),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnblockUser,
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, suite.testData.user),
					sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, "blocked"),
					sdk.NewAttribute(types.AttributeKeySubspace, "subspace"),
				),
			},
			expBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.otherUser, "blocked", "reason", "other_subspace"),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest()
		suite.Run(test.name, func() {

			profile := suite.CreateProfileFromAddress(suite.testData.user)
			otherProfile := suite.CreateProfileFromAddress(suite.testData.otherUser)

			err := suite.k.StoreProfile(suite.ctx, profile)
			suite.Require().NoError(err)

			err = suite.k.StoreProfile(suite.ctx, otherProfile)
			suite.Require().NoError(err)

			for _, block := range test.storedBlock {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err = service.UnblockUser(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				stored := suite.k.GetAllUsersBlocks(suite.ctx)
				suite.Require().Equal(test.expBlocks, stored)
			}
		})
	}
}
