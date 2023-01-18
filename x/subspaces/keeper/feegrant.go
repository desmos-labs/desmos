package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// SaveGrant saves the given grant inside the current context
func (k Keeper) SaveGrant(ctx sdk.Context, grant types.Grant) {
	store := ctx.KVStore(k.storeKey)
	store.Set(getGrantKey(grant), k.cdc.MustMarshal(&grant))
	if grantee, ok := grant.Grantee.GetCachedValue().(*types.UserGrantee); ok {
		k.createAccountIfNotExists(ctx, grantee.User)
	}
}

// getGrantKey returns the store key used to save the grant reference based on its grantee type
func getGrantKey(grant types.Grant) []byte {
	switch grantee := grant.Grantee.GetCachedValue().(type) {
	case *types.UserGrantee:
		return types.UserAllowanceKey(grant.SubspaceID, grantee.User)
	case *types.GroupGrantee:
		return types.GroupAllowanceKey(grant.SubspaceID, grantee.GroupID)
	default:
		panic(fmt.Errorf("unsupported content type: %T", grantee))
	}
}

// HasUserGrant tells whether the user grant having the given granter and grantee exists inside the provided subspace
func (k Keeper) HasUserGrant(ctx sdk.Context, subspaceID uint64, grantee string) bool {
	return ctx.KVStore(k.storeKey).Has(types.UserAllowanceKey(subspaceID, grantee))
}

// DeleteSection deletes the grant having the given granter and grantee from the subspace with the provided id
func (k Keeper) DeleteUserGrant(ctx sdk.Context, subspaceID uint64, grantee string) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, grantee)
	store.Delete(key)
}

// GetUserGrant returns the grant having the granter and grantee from the provided subspace.
// If there is no grant associated with the info the function will return false.
func (k Keeper) GetUserGrant(ctx sdk.Context, subspaceID uint64, grantee string) (types.Grant, bool) {
	if !k.HasUserGrant(ctx, subspaceID, grantee) {
		return types.Grant{}, false
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.UserAllowanceKey(subspaceID, grantee))
	var grant types.Grant
	k.cdc.MustUnmarshal(bz, &grant)
	return grant, true
}

// --------------------------------------------------------------------------------------------------------------------

// HasUserGrant tells whether the group grant having the given granter and group id exists inside the provided subspace
func (k Keeper) HasGroupGrant(ctx sdk.Context, subspaceID uint64, groupID uint32) bool {
	return ctx.KVStore(k.storeKey).Has(types.GroupAllowanceKey(subspaceID, groupID))
}

// DeleteSection deletes the grant having the given granter and group id from the subspace with the provided id
func (k Keeper) DeleteGroupGrant(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(subspaceID, groupID)
	store.Delete(key)
}

// GetGroupGrant returns the grant having the granter and group id from the provided subspace.
// If there is no grant associated with the info the function will return an error.
func (k Keeper) GetGroupGrant(ctx sdk.Context, subspaceID uint64, groupID uint32) (types.Grant, bool) {
	if !k.HasGroupGrant(ctx, subspaceID, groupID) {
		return types.Grant{}, false
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GroupAllowanceKey(subspaceID, groupID))
	var grant types.Grant
	k.cdc.MustUnmarshal(bz, &grant)
	return grant, true
}

// --------------------------------------------------------------------------------------------------------------------

// UseGrantedFees will try to pay the given fee from the granter's account as requested by the grantee
// if no valid allowance exists, then return false to show the fee will not be paid in this phase.
func (k Keeper) UseGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) bool {
	used := k.UseUserGrantedFees(ctx, subspaceID, grantee, fee, msgs)
	if used {
		return used
	}
	used = k.UseGroupGrantedFees(ctx, subspaceID, grantee, fee, msgs)
	return used
}

// UseUserGrantedFees will try to use the user grant to pay the given fee from the granter's account as requested by the grantee.
// if no valid allowance exists, then return false to show the fee will not be paid in this phase.
func (k Keeper) UseUserGrantedFees(ctx sdk.Context, subspaceID uint64, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) (used bool) {
	grant, found := k.GetUserGrant(ctx, subspaceID, grantee.String())
	if !found {
		return false
	}

	// update the allowance
	allowance, err := grant.GetUnpackedAllowance()
	if err != nil {
		return false
	}
	remove, err := allowance.Accept(ctx, fee, msgs)
	if remove {
		k.DeleteUserGrant(ctx, subspaceID, grantee.String())
	}
	if err != nil {
		return false
	}

	// update grant if allowance accept properly and still valid after execution
	if !remove {
		k.SaveGrant(ctx, types.NewGrant(subspaceID,
			grant.Granter,
			grant.Grantee.GetCachedValue().(types.Grantee),
			allowance))
	}
	return true
}

// UseGroupGrantedFees will try to use group grant to pay the given fee from the granter's account as requested by the grantee.
// if no valid allowance exists, then return false to show the fee will not be paid in this phase.
func (k Keeper) UseGroupGrantedFees(ctx sdk.Context, subspaceID uint64, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) (used bool) {
	k.IterateSubspaceUserGroupGrants(ctx, subspaceID, func(grant types.Grant) (stop bool) {
		groupGrantee := grant.Grantee.GetCachedValue().(*types.GroupGrantee)
		if !k.IsMemberOfGroup(ctx, grant.SubspaceID, groupGrantee.GroupID, grantee.String()) {
			return false
		}

		// update the allowance
		allowance, err := grant.GetUnpackedAllowance()
		if err != nil {
			return false
		}
		remove, err := allowance.Accept(ctx, fee, msgs)
		if remove {
			k.DeleteGroupGrant(ctx, subspaceID, grant.Grantee.GetCachedValue().(*types.GroupGrantee).GroupID)
		}
		if err != nil {
			return false
		}

		// update grant if allowance accept properly and still valid after execution
		if !remove {
			k.SaveGrant(ctx, types.NewGrant(subspaceID, grant.Granter, groupGrantee, allowance))
		}

		used = true
		return true
	})
	return used
}
