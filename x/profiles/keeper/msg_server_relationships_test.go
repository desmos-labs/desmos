package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgCreateRelationship() {
	tests := []struct {
		name                string
		storedBlock         []types.UserBlock
		storedRelationships []types.Relationship
		msg                 *types.MsgCreateRelationship
		expErr              bool
		expEvents           sdk.Events
		expRelationships    []types.Relationship
	}{
		{
			name: "Relationship sender blocked by receiver returns error",
			storedBlock: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"test",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: true,
		},
		{
			name: "Existing relationship returns error",
			storedRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: true,
		},
		{
			name: "Relationship has been saved correctly",
			msg: types.NewMsgCreateRelationship(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipCreated,
					sdk.NewAttribute(types.AttributeRelationshipSender, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				),
			},
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, rel := range test.storedRelationships {
				err := suite.k.SaveRelationship(suite.ctx, rel)
				suite.Require().NoError(err)
			}

			for _, block := range test.storedBlock {
				err := suite.k.SaveUserBlock(suite.ctx, block)
				suite.Require().NoError(err)
			}

			handler := keeper.NewMsgServerImpl(suite.k)
			_, err := handler.CreateRelationship(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				stored := suite.k.GetAllRelationships(suite.ctx)
				suite.Require().Equal(test.expRelationships, stored)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteRelationship() {
	tests := []struct {
		name             string
		stored           []types.Relationship
		msg              *types.MsgDeleteRelationship
		expErr           bool
		expEvents        sdk.Events
		expRelationships []types.Relationship
	}{
		{
			name: "Relationship not found returns error",
			stored: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
			},
			msg:    types.NewMsgDeleteRelationship("creator", "recipient", "other_subspace"),
			expErr: true,
		},
		{
			name: "Existing relationship is removed properly and leaves empty array",
			stored: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
			},
			msg:    types.NewMsgDeleteRelationship("creator", "recipient", "subspace"),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipsDeleted,
					sdk.NewAttribute(types.AttributeRelationshipSender, "creator"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "recipient"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "subspace"),
				),
			},
		},
		{
			name: "Existing relationship is removed properly and leaves not empty array",
			stored: []types.Relationship{
				types.NewRelationship("creator", "recipient", "subspace"),
				types.NewRelationship("creator", "recipient", "other_subspace"),
			},
			msg:    types.NewMsgDeleteRelationship("creator", "recipient", "subspace"),
			expErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeRelationshipsDeleted,
					sdk.NewAttribute(types.AttributeRelationshipSender, "creator"),
					sdk.NewAttribute(types.AttributeRelationshipReceiver, "recipient"),
					sdk.NewAttribute(types.AttributeRelationshipSubspace, "subspace"),
				),
			},
			expRelationships: []types.Relationship{
				types.NewRelationship("creator", "recipient", "other_subspace"),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, relationship := range test.stored {
				err := suite.k.SaveRelationship(suite.ctx, relationship)
				suite.Require().NoError(err)
			}

			service := keeper.NewMsgServerImpl(suite.k)
			_, err := service.DeleteRelationship(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				left := suite.k.GetAllRelationships(suite.ctx)
				suite.Require().Equal(test.expRelationships, left)
			}
		})
	}
}
