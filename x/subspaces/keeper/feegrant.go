package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

// SaveGrant saves the given grant inside the current context
func (k Keeper) SaveGrant(ctx sdk.Context, grant types.Grant) {
	store := ctx.KVStore(k.storeKey)
	store.Set(getGrantKey(grant), k.cdc.MustMarshal(&grant))

	if grantee, isUserGrantee := grant.Grantee.GetCachedValue().(*types.UserGrantee); isUserGrantee {
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
		panic(fmt.Errorf("unsupported grantee type: %T", grantee))
	}
}

// HasUserGrant tells whether the user grant associated to the given user exists inside the provided subspace
func (k Keeper) HasUserGrant(ctx sdk.Context, subspaceID uint64, grantee string) bool {
	return ctx.KVStore(k.storeKey).Has(types.UserAllowanceKey(subspaceID, grantee))
}

// DeleteUserGrant deletes the grant associated the given user grantee from the subspace with the provided id
func (k Keeper) DeleteUserGrant(ctx sdk.Context, subspaceID uint64, grantee string) {
	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, grantee)
	store.Delete(key)
}

// GetUserGrant returns the grant associated to the given user from the provided subspace.
// If there is no grant associated with the info the function will return false.
func (k Keeper) GetUserGrant(ctx sdk.Context, subspaceID uint64, grantee string) (grant types.Grant, found bool) {
	if !k.HasUserGrant(ctx, subspaceID, grantee) {
		return types.Grant{}, false
	}

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.UserAllowanceKey(subspaceID, grantee))
	k.cdc.MustUnmarshal(bz, &grant)
	return grant, true
}

// --------------------------------------------------------------------------------------------------------------------

// HasGroupGrant tells whether the group grant associated to the given group id exists inside the provided subspace
func (k Keeper) HasGroupGrant(ctx sdk.Context, subspaceID uint64, groupID uint32) bool {
	return ctx.KVStore(k.storeKey).Has(types.GroupAllowanceKey(subspaceID, groupID))
}

// DeleteGroupGrant deletes the grant having associated to the given group id from the subspace with the provided id
func (k Keeper) DeleteGroupGrant(ctx sdk.Context, subspaceID uint64, groupID uint32) {
	store := ctx.KVStore(k.storeKey)
	key := types.GroupAllowanceKey(subspaceID, groupID)
	store.Delete(key)
}

// GetGroupGrant returns the grant associated to the given group id from the provided subspace.
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

// UseGrantedFees will try to pay the given fees from the treasury account as requested by the grantee.
// If no valid allowance exists, then returns false to show the fees will not be paid in this phase.
func (k Keeper) UseGrantedFees(ctx sdk.Context, subspaceID uint64, grantee sdk.AccAddress, fees sdk.Coins, msgs []sdk.Msg) bool {
	return k.UseUserGrantedFees(ctx, subspaceID, grantee, fees, msgs) || k.UseGroupGrantedFees(ctx, subspaceID, grantee, fees, msgs)
}

// UseUserGrantedFees will try to use the user grant to pay the given fees from treasury as requested by the grantee.
// If no valid allowance exists, then returns false to show the fees will not be paid in this phase.
func (k Keeper) UseUserGrantedFees(ctx sdk.Context, subspaceID uint64, grantee sdk.AccAddress, fees sdk.Coins, msgs []sdk.Msg) (used bool) {
	grant, found := k.GetUserGrant(ctx, subspaceID, grantee.String())
	if !found {
		return false
	}

	// Get and accept the allowance
	allowance, err := grant.GetUnpackedAllowance()
	if err != nil {
		return false
	}

	remove, err := allowance.Accept(ctx, fees, msgs)
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
			allowance,
		))
	}

	return true
}

// UseGroupGrantedFees will try to use group grant to pay the given fee from the granter's account as requested by the grantee.
// If no valid allowance exists, then returns false to show the fees will not be paid in this phase.
func (k Keeper) UseGroupGrantedFees(ctx sdk.Context, subspaceID uint64, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) (used bool) {
	k.IterateSubspaceUserGroupGrants(ctx, subspaceID, func(grant types.Grant) (stop bool) {
		groupGrantee := grant.Grantee.GetCachedValue().(*types.GroupGrantee)
		if !k.IsMemberOfGroup(ctx, grant.SubspaceID, groupGrantee.GroupID, grantee.String()) {
			return false
		}

		// Get and update the allowance
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

		// Update the grant if the allowance was accepted properly and is still valid after execution
		if !remove {
			k.SaveGrant(ctx, types.NewGrant(subspaceID, grant.Granter, groupGrantee, allowance))
		}

		used = true
		return true
	})

	return used
}
