package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// SaveUserGrant saves the given user grant
func (k Keeper) SaveUserGrant(ctx sdk.Context, grant types.UserGrant) {
	granteeAddr, err := sdk.AccAddressFromBech32(grant.Grantee)
	if err != nil {
		panic(err)
	}
	if !k.ak.HasAccount(ctx, granteeAddr) {
		k.ak.SetAccount(ctx, authtypes.NewBaseAccountWithAddress(granteeAddr))
	}

	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(grant.SubspaceID, grant.Granter, grant.Grantee)
	store.Set(key, k.cdc.MustMarshal(&grant))
}

// HasUserGrant tells whether the user grant having the given granter and grantee exists inside the provided subspace
func (k Keeper) HasUserGrant(ctx sdk.Context, subspaceID uint64, granter string, grantee string) bool {
	return ctx.KVStore(k.storeKey).Has(types.UserAllowanceKey(subspaceID, granter, grantee))
}

// DeleteSection deletes the grant having the given granter and grantee from the subspace with the provided id
func (k Keeper) DeleteUserGrant(ctx sdk.Context, subspaceID uint64, granter string, grantee string) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, granter, grantee)
	store.Delete(key)
}

// GetUserGrant returns the grant having the granter and grantee from the provided subspace.
// If there is no grant associated with the info the function will return an error.
func (k Keeper) GetUserGrant(ctx sdk.Context, subspaceID uint64, granter, grantee string) (types.UserGrant, error) {
	if !k.HasUserGrant(ctx, subspaceID, granter, grantee) {
		return types.UserGrant{}, fmt.Errorf("user grant does not exist: subspace id %d, granter %s, grantee %s", subspaceID, granter, grantee)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.UserAllowanceKey(subspaceID, granter, grantee))
	var grant types.UserGrant
	k.cdc.MustUnmarshal(bz, &grant)
	return grant, nil
}

// --------------------------------------------------------------------------------------------------------------------

// SaveGroupGrant saves the given group grant
func (k Keeper) SaveGroupGrant(ctx sdk.Context, grant types.GroupGrant) {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(grant.SubspaceID, grant.Granter, grant.GroupID)
	store.Set(key, k.cdc.MustMarshal(&grant))
}

// HasUserGrant tells whether the group grant having the given granter and group id exists inside the provided subspace
func (k Keeper) HasGroupGrant(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32) bool {
	return ctx.KVStore(k.storeKey).Has(types.GroupAllowanceKey(subspaceID, granter, groupID))
}

// DeleteSection deletes the grant having the given granter and group id from the subspace with the provided id
func (k Keeper) DeleteGroupGrant(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32) {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(subspaceID, granter, groupID)
	store.Delete(key)
}

// GetGroupGrant returns the grant having the granter and group id from the provided subspace.
// If there is no grant associated with the info the function will return an error.
func (k Keeper) GetGroupGrant(ctx sdk.Context, subspaceID uint64, granter string, groupID uint32) (types.GroupGrant, error) {
	if !k.HasGroupGrant(ctx, subspaceID, granter, groupID) {
		return types.GroupGrant{}, fmt.Errorf("group grant does not exist: subspace id %d, granter %s, group id %d", subspaceID, granter, groupID)
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GroupAllowanceKey(subspaceID, granter, groupID))
	var grant types.GroupGrant
	k.cdc.MustUnmarshal(bz, &grant)
	return grant, nil
}

// --------------------------------------------------------------------------------------------------------------------

// UseGrantedFees will try to pay the given fee from the granter's account as requested by the grantee
// if no valid allowance exists, then return false to show the fee will not be paid in this phase.
func (k Keeper) UseGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) bool {
	used := k.UseUserGrantedFees(ctx, subspaceID, granter, grantee, fee, msgs)
	if used {
		return used
	}
	used = k.UseGroupGrantedFees(ctx, subspaceID, granter, grantee, fee, msgs)
	return used
}

// UseUserGrantedFees will try to use the user grant to pay the given fee from the granter's account as requested by the grantee.
// if no valid allowance exists, then return false to show the fee will not be paid in this phase.
func (k Keeper) UseUserGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) (used bool) {
	grant, err := k.GetUserGrant(ctx, subspaceID, granter.String(), grantee.String())
	if err != nil {
		return false
	}

	// update the allowance
	allowance, err := grant.GetUnpackedAllowance()
	if err != nil {
		return false
	}
	remove, err := allowance.Accept(ctx, fee, msgs)
	if remove {
		k.DeleteUserGrant(ctx, subspaceID, granter.String(), grantee.String())
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
		k.SaveUserGrant(ctx, grant)
	}
	return true
}

// UseGroupGrantedFees will try to use group grant to pay the given fee from the granter's account as requested by the grantee.
// if no valid allowance exists, then return false to show the fee will not be paid in this phase.
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
			k.DeleteGroupGrant(ctx, subspaceID, grant.Granter, grant.GroupID)
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
			k.SaveGroupGrant(ctx, grant)
		}

		used = true
		return true
	})
	return used
}
