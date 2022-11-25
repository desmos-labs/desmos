package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// GrantUserAllowance creates a new grant
func (k Keeper) GrantUserAllowance(ctx sdk.Context, subspaceID uint64, granter, grantee string, feeAllowance feegrant.FeeAllowanceI) error {
	// create an account for the grantee if it is not in account state
	granteeAddr, err := sdk.AccAddressFromBech32(grantee)
	if err != nil {
		return err
	}
	granteeAcc := k.ak.GetAccount(ctx, granteeAddr)
	if granteeAcc == nil {
		granteeAcc = k.ak.NewAccountWithAddress(ctx, granteeAddr)
		k.ak.SetAccount(ctx, granteeAcc)
	}
	store := ctx.KVStore(k.storeKey)
	key := types.UserAllowanceKey(subspaceID, granter, grantee)
	grant, err := types.NewUserGrant(subspaceID, granter, grantee, feeAllowance)
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

func (k Keeper) UseGrantedFees(ctx sdk.Context, subspaceID uint64, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) error {
	grant, found, err := k.GetUserAllowance(ctx, subspaceID, granter.String(), grantee.String())
	if err != nil {
		return err
	}
	if !found {
		return types.ErrNoAllowance
	}
	_, err = grant.Accept(ctx, fee, msgs)
	if err != nil {
		return err
	}
	return k.GrantUserAllowance(ctx, subspaceID, granter.String(), grantee.String(), grant)
}
