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

func (suite *KeeperTestSuite) TestKeeper_GetUsersRelationships() {
	tests := []struct {
		name                string
		storedRelationships []sdk.AccAddress
		expMap              map[string][]sdk.AccAddress
	}{
		{
			name:                "Return a non-empty address -> userBlocks map",
			storedRelationships: []sdk.AccAddress{suite.testData.user, suite.testData.otherUser},
			expMap: map[string][]sdk.AccAddress{
				suite.testData.user.String():      {suite.testData.otherUser},
				suite.testData.otherUser.String(): {suite.testData.user},
			},
		},
		{
			name:                "Return an empty address -> userBlocks map",
			storedRelationships: nil,
			expMap:              map[string][]sdk.AccAddress{},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationships != nil {
				_ = suite.keeper.StoreRelationship(suite.ctx, test.storedRelationships[0], test.storedRelationships[1])
				_ = suite.keeper.StoreRelationship(suite.ctx, test.storedRelationships[1], test.storedRelationships[0])
			}

			actualIDsMap := suite.keeper.GetUsersRelationships(suite.ctx)
			suite.Equal(test.expMap, actualIDsMap)
		})
	}

}

func (suite *KeeperTestSuite) TestKeeper_GetUserRelationships() {
	addr1, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	suite.NoError(err)
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	tests := []struct {
		name                string
		storedRelationships []sdk.AccAddress
		expRelationships    []sdk.AccAddress
	}{
		{
			name:                "Returns non empty userBlocks slice",
			storedRelationships: []sdk.AccAddress{addr1, addr2},
			expRelationships:    []sdk.AccAddress{addr1, addr2},
		},
		{
			name:                "Returns empty userBlocks slice",
			storedRelationships: nil,
			expRelationships:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedRelationships != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.RelationshipsStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&[]sdk.AccAddress{addr1, addr2}))
			}

			suite.Equal(test.expRelationships, suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user))
		})
	}
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
			name:                "Delete a relationship with len(userBlocks) > 1",
			storedRelationships: []sdk.AccAddress{addr1, addr2, addr3},
			expRelationships:    []sdk.AccAddress{addr1, addr3},
			userToDelete:        addr2,
		},
		{
			name:                "Delete a relationship with len(userBlocks) == 1",
			storedRelationships: []sdk.AccAddress{addr1},
			expRelationships:    nil,
			userToDelete:        addr1,
		},
		{
			name:                "Delete a relationship with len(userBlocks) == 0",
			storedRelationships: nil,
			expRelationships:    nil,
			userToDelete:        addr1,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedRelationships != nil {
				store.Set(types.RelationshipsStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedRelationships))
			}

			suite.keeper.DeleteRelationship(suite.ctx, suite.testData.user, test.userToDelete)
			rel := suite.keeper.GetUserRelationships(suite.ctx, suite.testData.user)
			suite.Equal(test.expRelationships, rel)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveUserBlock() {
	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		userBlock        types.UserBlock
		expErr           error
	}{
		{
			name: "already blocked user returns error",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, suite.testData.otherUser, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			userBlock: types.NewUserBlock(suite.testData.user, suite.testData.otherUser, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expErr:    fmt.Errorf("the user with %s address has already been blocked", suite.testData.otherUser),
		},
		{
			name:             "user block added correctly",
			storedUserBlocks: nil,
			userBlock:        types.NewUserBlock(suite.testData.user, suite.testData.otherUser, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			expErr:           nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedUserBlocks != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.UsersBlocksStoreKey(suite.testData.user), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedUserBlocks))
			}
			err := suite.keeper.SaveUserBlock(suite.ctx, test.userBlock)
			suite.Equal(test.expErr, err)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_UnblockUser() {
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	addr3, err := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	suite.NoError(err)

	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		expBlocks        []types.UserBlock
		expError         error
		userToUnblock    sdk.AccAddress
	}{
		{
			name: "Unblock user with len(storedUserBlocks) > 1",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewUserBlock(suite.testData.user, addr3, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr3, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			userToUnblock: addr2,
			expError:      nil,
		},
		{
			name:             "Unblock user with len(storedUserBlocks) == 1",
			storedUserBlocks: []types.UserBlock{types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")},
			expBlocks:        nil,
			userToUnblock:    addr2,
			expError:         nil,
		},
		{
			name:             "Delete a relationship with len(userBlocks) == 0",
			storedUserBlocks: nil,
			expBlocks:        nil,
			userToUnblock:    addr2,
			expError:         fmt.Errorf("blocked user with address %s not found", addr2),
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.storedUserBlocks != nil {
				store.Set(types.UsersBlocksStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedUserBlocks))
			}

			err := suite.keeper.UnblockUser(suite.ctx, suite.testData.user, test.userToUnblock, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e")
			suite.Equal(test.expError, err)
			rel := suite.keeper.GetUserBlocks(suite.ctx, suite.testData.user)
			suite.Equal(test.expBlocks, rel)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUserBlocks() {
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)

	tests := []struct {
		name             string
		storedUserBlocks []types.UserBlock
		expUserBlocks    []types.UserBlock
	}{
		{
			name: "Returns non empty user blocks slice",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expUserBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
		},
		{
			name:             "Returns empty user blocks slice",
			storedUserBlocks: nil,
			expUserBlocks:    nil,
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.storedUserBlocks != nil {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.UsersBlocksStoreKey(suite.testData.user),
					suite.keeper.Cdc.MustMarshalBinaryBare(&[]types.UserBlock{
						types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
					}))
			}

			suite.Equal(test.expUserBlocks, suite.keeper.GetUserBlocks(suite.ctx, suite.testData.user))
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetUsersBlocks() {
	addr2, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	suite.NoError(err)
	addr3, err := sdk.AccAddressFromBech32("cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h")
	suite.NoError(err)

	tests := []struct {
		name              string
		storedUsersBlocks []types.UserBlock
		expUsersBlocks    []types.UserBlock
	}{
		{
			name: "Returns a non-empty users blocks slice",
			storedUsersBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewUserBlock(suite.testData.otherUser, addr3, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
			expUsersBlocks: []types.UserBlock{
				types.NewUserBlock(suite.testData.user, addr2, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
				types.NewUserBlock(suite.testData.otherUser, addr3, "reason", "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"),
			},
		},
	}

	for _, test := range tests {
		suite.SetupTest()
		suite.Run(test.name, func() {
			for _, userBlock := range test.storedUsersBlocks {
				err := suite.keeper.SaveUserBlock(suite.ctx, userBlock)
				suite.NoError(err)
			}

			actualBlocks := suite.keeper.GetUsersBlocks(suite.ctx)
			suite.Equal(test.expUsersBlocks, actualBlocks)
		})
	}
}
