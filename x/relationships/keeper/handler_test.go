package keeper_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/relationships"

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
		isBlocked           bool
		expErr              error
		expEvent            sdk.Event
	}{
		{
			name:      "Relationship sender been blocked from receiver returns error",
			msg:       types.NewMsgCreateRelationship(sender, receiver, subspace),
			isBlocked: true,
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("The user with address %s has blocked you", receiver)),
		},
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
			if test.isBlocked {
				userBlock := relationships.NewUserBlock(receiver, sender, "test",
					"")
				_ = suite.keeper.SaveUserBlock(suite.ctx, userBlock)
			}

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

func (suite *KeeperTestSuite) Test_handleMsgBlockUser() {
	blocker, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	blocked, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	tests := []struct {
		name             string
		msg              types.MsgBlockUser
		storedUserBlocks []types.UserBlock
		expErr           error
		expEvent         sdk.Event
	}{
		{
			name:             "Relationship already created returns error",
			msg:              types.NewMsgBlockUser(blocker, blocked, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			storedUserBlocks: []types.UserBlock{types.NewUserBlock(blocker, blocked, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
			expErr:           sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("the user with %s address has already been blocked", blocked)),
		},
		{
			name:             "Relationship has been saved correctly",
			msg:              types.NewMsgBlockUser(blocker, blocked, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			storedUserBlocks: nil,
			expErr:           nil,
			expEvent: sdk.NewEvent(
				types.EventTypeBlockUser,
				sdk.NewAttribute(types.AttributeUserBlockBlocker, blocker.String()),
				sdk.NewAttribute(types.AttributeUserBlockBlocked, blocked.String()),
				sdk.NewAttribute(types.AttributeSubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				sdk.NewAttribute(types.AttributeUserBlockReason, "reason"),
			),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedUserBlocks != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.UsersBlocksStoreKey(test.msg.Blocker),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedUserBlocks))
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
				suite.Len(suite.keeper.GetUserBlocks(suite.ctx, blocker), 1)
			}

		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgUnblockUser() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	store.Set(types.UsersBlocksStoreKey(suite.testData.user),
		suite.keeper.Cdc.MustMarshalBinaryBare(&[]types.UserBlock{
			types.NewUserBlock(suite.testData.user, addr1, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
		}))

	testMsg := types.NewMsgUnblockUser(suite.testData.user, addr1, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")

	handler := keeper.NewHandler(suite.keeper)
	res, err := handler(suite.ctx, testMsg)

	suite.NoError(err)

	suite.Equal([]types.UserBlock{types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
		suite.keeper.GetUserBlocks(suite.ctx, suite.testData.user))

	// Check the events
	suite.Len(res.Events, 1)
	suite.Contains(res.Events, sdk.NewEvent(
		types.EventTypeUnblockUser,
		sdk.NewAttribute(types.AttributeUserBlockBlocker, suite.testData.user.String()),
		sdk.NewAttribute(types.AttributeUserBlockBlocked, addr1.String()),
		sdk.NewAttribute(types.AttributeSubspace, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
	))
}
