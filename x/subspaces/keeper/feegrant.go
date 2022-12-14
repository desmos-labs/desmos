package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// GrantUserAllowance creates a new grant
func (k Keeper) GrantUserAllowance(ctx sdk.Context, subspaceID uint64, granter, grantee string, allowance feegrant.FeeAllowanceI) error {
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

func (k Keeper) RevokeUserGrant(ctx sdk.Context, subspaceID uint64, granter string, grantee string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, granter, grantee)
	if !store.Has(key) {
		return types.ErrNoAllowance
	}
	store.Delete(key)
	return nil
}

func (k Keeper) GetUserAllowance(ctx sdk.Context, subspaceID uint64, granter, grantee string) (feegrant.FeeAllowanceI, bool, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, granter, grantee)
	if !store.Has(key) {
		return nil, false, nil
	}
	bz := store.Get(key)
	var grant types.UserGrant
	if err := k.cdc.Unmarshal(bz, &grant); err != nil {
		return nil, false, err
	}
	allowance, err := grant.GetUnpackedAllowance()
	return allowance, true, err
}

func (k Keeper) GrantGroupAllowance(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32, feeAllowance feegrant.FeeAllowanceI) error {
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

func (k Keeper) RevokeGroupGrant(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(subspaceID, granter, groupID)
	if !store.Has(key) {
		return types.ErrNoAllowance
	}
	store.Delete(key)
	return nil
}

func (k Keeper) GetGroupAllowance(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32) (feegrant.FeeAllowanceI, bool, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(subspaceID, granter, groupID)
	if !store.Has(key) {
		return nil, false, nil
	}
	bz := store.Get(key)
	var grant types.GroupGrant
	if err := k.cdc.Unmarshal(bz, &grant); err != nil {
		return nil, false, err
	}
	allowance, err := grant.GetUnpackedAllowance()
	return allowance, true, err
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
	allowance, found, err := k.GetUserAllowance(ctx, subspaceID, granter.String(), grantee.String())
	if err != nil || !found {
		return false
	}
	remove, err := allowance.Accept(ctx, fee, msgs)
	if remove {
		k.RevokeUserGrant(ctx, subspaceID, granter.String(), grantee.String())
	}
	if err != nil {
		return false
	}
	err = k.GrantUserAllowance(ctx, subspaceID, granter.String(), grantee.String(), allowance)
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
			k.RevokeGroupGrant(ctx, subspaceID, entry.Granter, entry.GroupID)
		}
		if err != nil {
			return false
		}
		err = k.GrantGroupAllowance(ctx, subspaceID, granter.String(), entry.GroupID, allowance)
		if err != nil {
			return false
		}
		used = true
		return true
	})
	return used
}
