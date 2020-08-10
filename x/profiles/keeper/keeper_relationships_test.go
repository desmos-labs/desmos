package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveRelationship() {
	monoRelationship := types.NewMonodirectionalRelationship(suite.testData.user, suite.testData.otherUser)
	biRelationship := types.NewBiDirectionalRelationship(suite.testData.user, suite.testData.otherUser, types.Sent)

	tests := []struct {
		name                string
		storedRelationships types.Relationships
		relationship        types.Relationship
		expErr              error
	}{
		{
			name:                "Storing the same mono relationship returns error",
			storedRelationships: types.Relationships{monoRelationship},
			relationship:        monoRelationship,
			expErr:              fmt.Errorf("relationship between %s and %s has already been done", suite.testData.user, suite.testData.otherUser),
		},
		{
			name:                "Storing the same bidirectional relationship returns error",
			storedRelationships: types.Relationships{biRelationship},
			relationship:        biRelationship,
			expErr:              fmt.Errorf("relationship between %s and %s has already been done", suite.testData.user, suite.testData.otherUser),
		},
		{
			name:                "Storing new relationship returns no error",
			storedRelationships: types.Relationships{monoRelationship},
			relationship:        biRelationship,
			expErr:              nil,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			store.Set(types.RelationshipsStoreKey(suite.testData.user), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))
			err := suite.keeper.SaveRelationship(suite.ctx, test.relationship)
			suite.Equal(test.expErr, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	relationships := types.Relationships{
		types.NewMonodirectionalRelationship(suite.testData.user, suite.testData.otherUser),
		types.NewBiDirectionalRelationship(suite.testData.user, suite.testData.otherUser, types.Sent),
	}

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	store.Set(types.RelationshipsStoreKey(suite.testData.user), suite.keeper.Cdc.MustMarshalBinaryBare(&relationships))

	actualRelationships := suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user)
	suite.Equal(relationships, actualRelationships)
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRelationship() {
	monoRelationship := types.NewMonodirectionalRelationship(suite.testData.user, suite.testData.otherUser)

	user, _ := sdk.AccAddressFromBech32("cosmos1m5gfj4t5ddksytl65mmv7lfg5nef3etmrnl8a0")
	anotherRelationship := types.NewMonodirectionalRelationship(suite.testData.user, user)

	tests := []struct {
		name                string
		storedRelationships types.Relationships
		user                sdk.AccAddress
		counterpart         sdk.AccAddress
		expErr              error
	}{
		{
			name:                "Non existent relationship returns error",
			storedRelationships: types.Relationships{monoRelationship},
			user:                user,
			counterpart:         suite.testData.user,
			expErr:              fmt.Errorf("no relationship found between %s and %s", user, suite.testData.user),
		},
		{
			name:                "Existent relationship deleted correctly",
			storedRelationships: types.Relationships{anotherRelationship, monoRelationship},
			user:                suite.testData.user,
			counterpart:         user,
			expErr:              nil,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			store.Set(types.RelationshipsStoreKey(suite.testData.user), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))

			err := suite.keeper.DeleteRelationship(suite.ctx, test.user, test.counterpart)
			suite.Equal(test.expErr, err)

			if test.expErr == nil {
				relationships := suite.keeper.GetUserRelationships(suite.ctx, test.user)
				suite.Len(relationships, 1)
				suite.Equal(types.Relationships{monoRelationship}, relationships)
			}
		})
	}
}
