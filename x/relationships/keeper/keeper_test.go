package keeper_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

func (suite *KeeperTestSuite) TestKeeper_StoreRelationship() {
	tests := []struct {
		name                string
		storedRelationships []sdk.AccAddress
		user                sdk.AccAddress
		receiver            sdk.AccAddress
		expErr              error
	}{
		{
			name:                "already existent relationship returns error",
			storedRelationships: []sdk.AccAddress{suite.testData.otherUser},
			user:                suite.testData.user,
			receiver:            suite.testData.otherUser,
			expErr:              fmt.Errorf("relationship already exists with %s", suite.testData.otherUser),
		},
		{
			name:                "relationship added correctly",
			storedRelationships: nil,
			user:                suite.testData.user,
			receiver:            suite.testData.otherUser,
			expErr:              nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationships != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.RelationshipsStoreKey(test.user), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))
			}
			err := suite.keeper.StoreRelationship(suite.ctx, test.user, test.receiver)
			suite.Equal(test.expErr, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationshipsMap() {
	relationshipsMap := map[string][]sdk.AccAddress{
		suite.testData.user.String():      {suite.testData.otherUser},
		suite.testData.otherUser.String(): {suite.testData.user},
	}

	_ = suite.keeper.StoreRelationship(suite.ctx, suite.testData.user, suite.testData.otherUser)
	_ = suite.keeper.StoreRelationship(suite.ctx, suite.testData.otherUser, suite.testData.user)

	actualIDsMap := suite.keeper.GetUsersRelationships(suite.ctx)
	suite.Equal(relationshipsMap, actualIDsMap)
}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	store.Set(types.RelationshipsStoreKey(suite.testData.user),
		suite.keeper.Cdc.MustMarshalBinaryBare(&[]sdk.AccAddress{addr1, addr2, suite.testData.otherUser}))

	expRelationships := []sdk.AccAddress{addr1, addr2, suite.testData.otherUser}
	suite.Equal(expRelationships, suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user))
}

func (suite *KeeperTestSuite) TestKeeper_DeleteRelationship() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	addr3, err := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	suite.NoError(err)

	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	store.Set(types.RelationshipsStoreKey(suite.testData.user),
		suite.keeper.Cdc.MustMarshalBinaryBare(&[]sdk.AccAddress{addr1, addr2, addr3}))

	suite.keeper.DeleteRelationship(suite.ctx, suite.testData.user, addr2)

	expRelationships := []sdk.AccAddress{addr1, addr3}
	suite.Equal(expRelationships, suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user))

	tests := []struct {
		name                string
		storedRelationships []sdk.AccAddress
		expRelationships    []sdk.AccAddress
		userToDelete        sdk.AccAddress
	}{
		{
			name:                "Delete a relationship with len(relationships) > 1",
			storedRelationships: []sdk.AccAddress{addr1, addr2, addr3},
			expRelationships:    []sdk.AccAddress{addr1, addr3},
			userToDelete:        addr2,
		},
		{
			name:                "Delete a relationship with len(relationships) == 1",
			storedRelationships: []sdk.AccAddress{addr1},
			expRelationships:    nil,
			userToDelete:        addr1,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			store.Set(types.RelationshipsStoreKey(suite.testData.user),
				suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))
			suite.keeper.DeleteRelationship(suite.ctx, suite.testData.user, test.userToDelete)
			rel := suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user)
			suite.Equal(test.expRelationships, rel)
		})
	}
}
