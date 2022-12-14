package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// SaveUserFeeGrant saves a new user grant
func (k Keeper) SaveUserFeeGrant(ctx sdk.Context, subspaceID uint64, granter, grantee string, allowance feegrant.FeeAllowanceI) error {
	granteeAddr, err := sdk.AccAddressFromBech32(grantee)
	if err != nil {
		return err
	}
	if !k.ak.HasAccount(ctx, granteeAddr) {
		k.ak.SetAccount(ctx, authtypes.NewBaseAccountWithAddress(granteeAddr))
	}

	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, granter, grantee)
	grant, err := types.NewUserGrant(subspaceID, granter, grantee, allowance)
	if err != nil {
		return err
	}
	store.Set(key, k.cdc.MustMarshal(&grant))
	return nil
}

// RemoveUserAllowance remove a user grant
func (k Keeper) RemoveUserAllowance(ctx sdk.Context, subspaceID uint64, granter string, grantee string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, granter, grantee)
	if !store.Has(key) {
		return types.ErrNoAllowance
	}
	store.Delete(key)
	return nil
}

func (k Keeper) GetUserGrant(ctx sdk.Context, subspaceID uint64, granter, grantee string) (types.UserGrant, bool, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, granter, grantee)
	if !store.Has(key) {
		return types.UserGrant{}, false, nil
	}
	bz := store.Get(key)
	var grant types.UserGrant
	if err := k.cdc.Unmarshal(bz, &grant); err != nil {
		return types.UserGrant{}, false, err
	}
	return grant, true, nil
}

func (k Keeper) SaveGroupAllowance(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32, feeAllowance feegrant.FeeAllowanceI) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(subspaceID, granter, groupID)
	grant, err := types.NewGroupGrant(subspaceID, granter, groupID, feeAllowance)
	if err != nil {
		return err
	}
	bz, err := k.cdc.Marshal(&grant)
	if err != nil {
		return err
	}
	store.Set(key, bz)
	return nil
}

func (k Keeper) RemoveGroupAllowance(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(subspaceID, granter, groupID)
	if !store.Has(key) {
		return types.ErrNoAllowance
	}
	store.Delete(key)
	return nil
}

func (k Keeper) GetGroupGrant(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32) (types.GroupGrant, bool, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(subspaceID, granter, groupID)
	if !store.Has(key) {
		return types.GroupGrant{}, false, nil
	}
	bz := store.Get(key)
	var grant types.GroupGrant
	if err := k.cdc.Unmarshal(bz, &grant); err != nil {
		return types.GroupGrant{}, false, err
	}
	return grant, true, nil
}

func (k Keeper) UseGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) bool {
	used := k.UseUserGrantedFees(ctx, subspaceID, granter, grantee, fee, msgs)
	if used {
		return used
	}
	used = k.UseGroupGrantedFees(ctx, subspaceID, granter, grantee, fee, msgs)
	return used
}

func (k Keeper) UseUserGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) (used bool) {
	grant, found, err := k.GetUserGrant(ctx, subspaceID, granter.String(), grantee.String())
	if err != nil || !found {
		return false
	}
	allowance, err := grant.GetUnpackedAllowance()
	if err != nil {
		return false
	}
	remove, err := allowance.Accept(ctx, fee, msgs)
	if remove {
		k.RemoveUserAllowance(ctx, subspaceID, granter.String(), grantee.String())
	}
	if err != nil {
		return false
	}
	err = k.SaveUserFeeGrant(ctx, subspaceID, granter.String(), grantee.String(), allowance)
	if err != nil {
		return false
	}
	return true
}

func (k Keeper) UseGroupGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) (used bool) {
	k.IterateSubspaceGranterGroupGrants(ctx, subspaceID, granter.String(), func(entry types.GroupGrant) (stop bool) {
		if !k.IsMemberOfGroup(ctx, entry.SubspaceID, entry.GroupID, grantee.String()) {
			return false
		}

		allowance, err := entry.GetUnpackedAllowance()
		if err != nil {
			return false
		}
		remove, err := allowance.Accept(ctx, fee, msgs)
		if remove {
			k.RemoveGroupAllowance(ctx, subspaceID, entry.Granter, entry.GroupID)
		}
		if err != nil {
			return false
		}
		err = k.SaveGroupAllowance(ctx, subspaceID, granter.String(), entry.GroupID, allowance)
		if err != nil {
			return false
		}
		used = true
		return true
	})
	return used
}
