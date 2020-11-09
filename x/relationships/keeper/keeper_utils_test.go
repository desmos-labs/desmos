package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

func (suite *KeeperTestSuite) TestCheckForBlockedUser() {
	user, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	otherUser, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

	tests := []struct {
		name       string
		blocker    sdk.AccAddress
		blocked    sdk.AccAddress
		userBlocks []types.UserBlock
		expBool    bool
	}{
		{
			name:       "blocked user found returns true",
			blocker:    user,
			blocked:    otherUser,
			userBlocks: []types.UserBlock{types.NewUserBlock(user, otherUser, "test", "")},
			expBool:    true,
		},
		{
			name:       "non blocked user not found returns false",
			blocker:    user,
			blocked:    otherUser,
			userBlocks: nil,
			expBool:    false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if test.userBlocks != nil {
				store.Set(types.UsersBlocksStoreKey(test.blocker),
					suite.keeper.Cdc.MustMarshalBinaryBare(&test.userBlocks))
			}
			res := suite.keeper.IsUserBlocked(suite.ctx, test.blocker, test.blocked)
			suite.Equal(test.expBool, res)
		})
	}
}
