package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

func (suite *KeeperTestSuite) Test_handleMsgCreateRelationship() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	tests := []struct {
		name                string
		msg                 types.MsgCreateRelationship
		storedRelationships types.Relationships
		expErr              error
		expEvent            sdk.Event
	}{
		{
			name:                "Relationship already created returns error",
			msg:                 types.NewMsgCreateRelationship(sender, receiver, subspace),
			storedRelationships: types.Relationships{types.NewRelationship(receiver, subspace)},
			expErr:              sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("relationship already exists with %s", receiver)),
		},
		{
			name:                "Relationship has been saved correctly",
			msg:                 types.NewMsgCreateRelationship(sender, receiver, subspace),
			storedRelationships: nil,
			expErr:              nil,
			expEvent: sdk.NewEvent(
				types.EventTypeRelationshipCreated,
				sdk.NewAttribute(types.AttributeRelationshipSender, sender.String()),
				sdk.NewAttribute(types.AttributeRelationshipReceiver, receiver.String()),
				sdk.NewAttribute(types.AttributeRelationshipSubspace, subspace),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationships != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.RelationshipsStoreKey(test.msg.Sender),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))
			}

			handler := keeper.NewHandler(suite.keeper)
			res, err := handler(suite.ctx, test.msg)

			if test.expErr != nil {
				suite.Error(err)
				suite.Equal(test.expErr.Error(), err.Error())
			}

			if test.expErr == nil {
				suite.NoError(err)

				// Check the events
				suite.Len(res.Events, 1)
				suite.Contains(res.Events, test.expEvent)
				suite.Len(suite.keeper.GetUserRelationships(suite.ctx, sender), 1)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgDeleteRelationship() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	subspace := "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	store.Set(types.RelationshipsStoreKey(suite.testData.user),
		suite.keeper.Cdc.MustMarshalBinaryBare(&types.Relationships{
			types.NewRelationship(addr1, subspace),
			types.NewRelationship(addr2, subspace),
		}))

	testMsg := types.NewMsgDeleteRelationship(suite.testData.user, addr1, subspace)

	handler := keeper.NewHandler(suite.keeper)
	res, err := handler(suite.ctx, testMsg)

	suite.NoError(err)

	suite.Equal(types.Relationships{types.NewRelationship(addr2, subspace)},
		suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user))

	// Check the events
	suite.Len(res.Events, 1)
	suite.Contains(res.Events, sdk.NewEvent(
		types.EventTypeRelationshipsDeleted,
		sdk.NewAttribute(types.AttributeRelationshipSender, suite.testData.user.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, addr1.String()),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, subspace),
	))
}
