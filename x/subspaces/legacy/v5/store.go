package v4

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// MigrateStore migrates the store from version 4 to version 5.
// The migration includes the following:
//
// - create account for all the users inside groups if they don't have it;
func MigrateStore(ctx sdk.Context, key sdk.StoreKey, accountKeeper authkeeper.AccountKeeper) error {
	return migrateUserAccountsInUserGroups(ctx, key, accountKeeper)
}

func migrateUserAccountsInUserGroups(ctx sdk.Context, key sdk.StoreKey, accountKeeper authkeeper.AccountKeeper) error {
	groupsStore := prefix.NewStore(ctx.KVStore(key), types.GroupsMembersPrefix)
	iterator := groupsStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		_, _, user := types.SplitGroupMemberStoreKey(append(types.GroupsMembersPrefix, iterator.Key()...))
		userAcc, err := sdk.AccAddressFromBech32(user)
		if err != nil {
			return err
		}

		accExists := accountKeeper.HasAccount(ctx, userAcc)
		if !accExists {
			accountKeeper.SetAccount(ctx, accountKeeper.NewAccountWithAddress(ctx, userAcc))
		}
	}
	return nil
}
