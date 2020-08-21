package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/desmos-labs/desmos/x/profiles/types/models"
)

func (suite *KeeperTestSuite) TestKeeper_SaveUserRelationshipAssociation() {
	id := types.RelationshipID("12345")
	expIDs := []types.RelationshipID{id}

	suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{suite.testData.user}, id)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	var actualIDs []types.RelationshipID
	suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.UserRelationshipsStoreKey(suite.testData.user)), &actualIDs)

	suite.Equal(expIDs, actualIDs)
}

func (suite *KeeperTestSuite) TestKeeper_DoesRelationshipExist() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	relationship := models.NewMonodirectionalRelationship(sender, receiver)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	store.Set(types.RelationshipsStoreKey(relationship.ID), suite.keeper.Cdc.MustMarshalBinaryBare(&relationship))

	tests := []struct {
		name           string
		relationshipID types.RelationshipID
		expBool        bool
	}{
		{
			name:           "Found relationship returns true",
			relationshipID: relationship.ID,
			expBool:        true,
		},
		{
			name:           "Not found relationship returns false",
			relationshipID: types.RelationshipID("123"),
			expBool:        false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			res := suite.keeper.DoesRelationshipExist(suite.ctx, test.relationshipID)
			suite.Equal(test.expBool, res)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_StoreRelationship() {
	monoRelationship := types.NewMonodirectionalRelationship(suite.testData.user, suite.testData.otherUser)
	suite.keeper.StoreRelationship(suite.ctx, monoRelationship)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	var actualRel types.Relationship
	suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.RelationshipsStoreKey(monoRelationship.ID)), &actualRel)

	suite.Equal(monoRelationship, actualRel)
}

func (suite *KeeperTestSuite) TestKeeper_GetRelationshipFromID() {
	var abstractRel types.Relationship
	abstractRel = types.NewMonodirectionalRelationship(suite.testData.user, suite.testData.otherUser)

	tests := []struct {
		name               string
		storedRelationship *types.Relationship
		relationshipID     types.RelationshipID
		expRelationship    types.Relationship
		expErr             error
	}{
		{
			name:               "non existent relationship returns error",
			storedRelationship: nil,
			relationshipID:     abstractRel.RelationshipID(),
			expRelationship:    nil,
			expErr:             fmt.Errorf("relationship with id %s doesn't exist", abstractRel.RelationshipID()),
		},
		{
			name:               "existent relationship is returned correctly",
			storedRelationship: &abstractRel,
			relationshipID:     abstractRel.RelationshipID(),
			expRelationship:    abstractRel,
			expErr:             nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		if test.storedRelationship != nil {
			suite.keeper.StoreRelationship(suite.ctx, *test.storedRelationship)
		}

		relationship, err := suite.keeper.GetRelationshipFromID(suite.ctx, test.relationshipID)
		if test.expErr != nil {
			suite.Error(err)
			suite.Equal(test.expErr, err)
		}
		if test.expErr == nil {
			suite.Equal(test.expRelationship, relationship)
		}
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetRelationships() {
	monoRelationship := types.NewMonodirectionalRelationship(suite.testData.user, suite.testData.otherUser)
	biRelationship := types.NewBiDirectionalRelationship(suite.testData.user, suite.testData.otherUser, types.Sent)

	relationships := types.Relationships{monoRelationship, biRelationship}

	for _, rel := range relationships {
		store := suite.ctx.KVStore(suite.keeper.StoreKey)
		store.Set(types.RelationshipsStoreKey(rel.RelationshipID()), suite.keeper.Cdc.MustMarshalBinaryBare(&rel))
	}

	actualRelationships := suite.keeper.GetRelationships(suite.ctx)

	suite.Equal(relationships, actualRelationships)
}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationshipsIDsMap() {
	monoRelationship := types.NewMonodirectionalRelationship(suite.testData.user, suite.testData.otherUser)
	biRelationship := types.NewBiDirectionalRelationship(suite.testData.user, suite.testData.otherUser, types.Sent)

	relationshipIDsMap := map[string]types.RelationshipIDs{
		suite.testData.user.String():      {monoRelationship.ID, biRelationship.ID},
		suite.testData.otherUser.String(): {biRelationship.ID},
	}

	suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{suite.testData.user}, monoRelationship.ID)
	suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{suite.testData.user, suite.testData.otherUser}, biRelationship.ID)

	actualIDsMap := suite.keeper.GetUsersRelationshipsIDMap(suite.ctx)

	suite.Equal(relationshipIDsMap, actualIDsMap)
}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	monoRelationship := models.NewMonodirectionalRelationship(sender, receiver)
	biRelationship := models.NewBiDirectionalRelationship(sender, receiver, types.Accepted)

	IDs := []types.RelationshipID{monoRelationship.ID, biRelationship.ID}

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	store.Set(types.UserRelationshipsStoreKey(sender), suite.keeper.Cdc.MustMarshalBinaryBare(&IDs))

	suite.keeper.StoreRelationship(suite.ctx, monoRelationship)
	suite.keeper.StoreRelationship(suite.ctx, biRelationship)

	expRelationships := types.Relationships{monoRelationship, biRelationship}

	suite.Equal(expRelationships, suite.keeper.GetUserRelationships(suite.ctx, sender))
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRelationship() {
	sender, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	receiver, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	anotherUser, err := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	suite.NoError(err)

	monoRelationship := models.NewMonodirectionalRelationship(sender, receiver)
	biRelationship := models.NewBiDirectionalRelationship(sender, receiver, types.Accepted)

	tests := []struct {
		name                     string
		storedRelationships      types.Relationships
		user                     sdk.AccAddress
		relationshipID           types.RelationshipID
		expErr                   error
		expRelationshipsSender   types.Relationships
		expRelationshipsReceiver types.Relationships
	}{
		{
			name:                "Unauthorized user tries to delete a relationship returns error",
			storedRelationships: types.Relationships{monoRelationship, biRelationship},
			user:                anotherUser,
			relationshipID:      monoRelationship.ID,
			expErr:              fmt.Errorf("user with address cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h isn't the relationship's creator"),
		},
		{
			name:                "Unauthorized user tries to delete a relationship(bidirectional) returns error",
			storedRelationships: types.Relationships{monoRelationship, biRelationship},
			user:                anotherUser,
			relationshipID:      biRelationship.ID,
			expErr:              fmt.Errorf("user with address cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h is neither the creator nor the recipient of the relationship"),
		},
		{
			name:                     "User delete a monodirectionalRelationship successfully",
			storedRelationships:      types.Relationships{monoRelationship, biRelationship},
			user:                     sender,
			relationshipID:           monoRelationship.ID,
			expErr:                   nil,
			expRelationshipsSender:   types.Relationships{biRelationship},
			expRelationshipsReceiver: types.Relationships{biRelationship},
		},
		{
			name:                     "User delete a bidirectionalRelationship successfully (user equals to the creator)",
			storedRelationships:      types.Relationships{monoRelationship, biRelationship},
			user:                     sender,
			relationshipID:           biRelationship.ID,
			expErr:                   nil,
			expRelationshipsSender:   types.Relationships{monoRelationship},
			expRelationshipsReceiver: nil,
		},
		{
			name:                     "User delete a bidirectionalRelationship successfully (user equals to the recipient)",
			storedRelationships:      types.Relationships{monoRelationship, biRelationship},
			user:                     receiver,
			relationshipID:           biRelationship.ID,
			expErr:                   nil,
			expRelationshipsSender:   types.Relationships{monoRelationship},
			expRelationshipsReceiver: nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest() // reset
		suite.Run(test.name, func() {
			suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{sender}, monoRelationship.ID)
			suite.keeper.SaveUserRelationshipAssociation(suite.ctx, []sdk.AccAddress{sender, receiver}, biRelationship.ID)
			for _, rel := range test.storedRelationships {
				suite.keeper.StoreRelationship(suite.ctx, rel)
			}

			actualErr := suite.keeper.DeleteRelationship(suite.ctx, test.relationshipID, test.user)
			suite.Equal(test.expErr, actualErr)

			if test.expErr == nil {
				actualSenderRels := suite.keeper.GetUserRelationships(suite.ctx, sender)
				actualReceiverRels := suite.keeper.GetUserRelationships(suite.ctx, receiver)

				suite.Equal(test.expRelationshipsSender, actualSenderRels)
				suite.Equal(test.expRelationshipsReceiver, actualReceiverRels)
			}
		})
	}
}
