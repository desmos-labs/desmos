package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// GrantAllowance creates a new grant
func (k Keeper) GrantAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress, subspaceId uint64, feeAllowance feegrant.FeeAllowanceI) error {
	// create the account if it is not in account state
	granteeAcc := k.ak.GetAccount(ctx, grantee)
	if granteeAcc == nil {
		granteeAcc = k.ak.NewAccountWithAddress(ctx, grantee)
		k.ak.SetAccount(ctx, granteeAcc)
	}

	store := ctx.KVStore(k.storeKey)
	key := types.FeeAllowanceKey(granter.String(), grantee.String(), subspaceId)

	grant, err := feegrant.NewGrant(granter, grantee, feeAllowance)
	bz, err := k.cdc.Marshal(&grant)
	if err != nil {
		return err
	}

	store.Set(key, bz)

	return nil
}

func (k Keeper) GetAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress, subspaceId uint64) (feegrant.FeeAllowanceI, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.FeeAllowanceKey(granter.String(), grantee.String(), subspaceId)
	if !store.Has(key) {
		return nil, nil
	}
	bz := store.Get(key)
	var feegrant feegrant.Grant
	if err := k.cdc.Unmarshal(bz, &feegrant); err != nil {
		return nil, err
	}
	return feegrant.GetGrant()
}

func (k Keeper) UseGrantedFees(ctx sdk.Context, granter, grantee sdk.AccAddress, subspaceId uint64, fee sdk.Coins, msgs []sdk.Msg) error {
	grant, err := k.GetAllowance(ctx, granter, grantee, subspaceId)
	if err != nil {
		return err
	}
	_, err = grant.Accept(ctx, fee, msgs)
	if err != nil {
		return err
	}
	return k.GrantAllowance(ctx, granter, grantee, subspaceId, grant)
}
