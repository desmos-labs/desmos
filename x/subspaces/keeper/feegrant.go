package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// SaveUserGrant saves a new user grant
func (k Keeper) SaveUserGrant(ctx sdk.Context, grant types.UserGrant) error {
	granteeAddr, err := sdk.AccAddressFromBech32(grant.Grantee)
	if err != nil {
		return err
	}
	if !k.ak.HasAccount(ctx, granteeAddr) {
		k.ak.SetAccount(ctx, authtypes.NewBaseAccountWithAddress(granteeAddr))
	}

	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(grant.SubspaceID, grant.Granter, grant.Grantee)
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

func (k Keeper) SaveGroupGrant(ctx sdk.Context, grant types.GroupGrant) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(grant.SubspaceID, grant.Granter, grant.GroupID)
	store.Set(key, k.cdc.MustMarshal(&grant))
	return nil
}

func (k Keeper) RemoveGroupGrant(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32) error {
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
	// update the allowance
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
	// update grant if allowance accept properly and still valid after execution
	if !remove {
		grant, err = types.NewUserGrant(subspaceID, granter.String(), grantee.String(), allowance)
		if err != nil {
			return false
		}
		err = k.SaveUserGrant(ctx, grant)
		if err != nil {
			return false
		}
	}
	return true
}

func (k Keeper) UseGroupGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) (used bool) {
	k.IterateSubspaceGranterGroupGrants(ctx, subspaceID, granter.String(), func(grant types.GroupGrant) (stop bool) {
		if !k.IsMemberOfGroup(ctx, grant.SubspaceID, grant.GroupID, grantee.String()) {
			return false
		}
		// update the allowance
		allowance, err := grant.GetUnpackedAllowance()
		if err != nil {
			return false
		}
		remove, err := allowance.Accept(ctx, fee, msgs)
		if remove {
			k.RemoveGroupGrant(ctx, subspaceID, grant.Granter, grant.GroupID)
		}
		if err != nil {
			return false
		}
		// update grant if allowance accept properly and still valid after execution
		if !remove {
			grant, err = types.NewGroupGrant(subspaceID, granter.String(), grant.GroupID, allowance)
			if err != nil {
				return false
			}
			err = k.SaveGroupGrant(ctx, grant)
			if err != nil {
				return false
			}
		}

		used = true
		return true
	})
	return used
}
