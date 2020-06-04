package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/desmos-labs/desmos/x/profile/internal/types"
)

// SetNameSurnameLenParams sets NameSurnameLenParams to the global param store
func (k Keeper) SetNameSurnameLenParams(ctx sdk.Context, nmParams params.NameSurnameLenParams) {
	k.paramSpace.Set(ctx, params.ParamStoreKeyNameSurnameLen, &nmParams)
}

// GetNameSurnameLenParams returns the current NameSurnameLenParams from the global param store
func (k Keeper) GetNameSurnameLenParams(ctx sdk.Context) (nmParams params.NameSurnameLenParams) {
	k.paramSpace.Get(ctx, params.ParamStoreKeyNameSurnameLen, &nmParams)
	return nmParams
}

// SetMonikerLenParams sets MonikerLenParams to the global param store
func (k Keeper) SetMonikerLenParams(ctx sdk.Context, monikerParams params.MonikerLenParams) {
	k.paramSpace.Set(ctx, params.ParamStoreKeyMonikerLen, &monikerParams)
}

// GetMonikerLenParams returns the current MonikerLenParams from the global param store
func (k Keeper) GetMonikerLenParams(ctx sdk.Context) (mParams params.MonikerLenParams) {
	k.paramSpace.Get(ctx, params.ParamStoreKeyMonikerLen, &mParams)
	return mParams
}

// SetBioLenParams sets the BioLenParams to the global param store
func (k Keeper) SetBioLenParams(ctx sdk.Context, bioParams params.BioLenParams) {
	k.paramSpace.Set(ctx, params.ParamStoreKeyMaxBioLen, &bioParams)
}

// GetBioLenParams returns the current BioLenParams from the global param store
func (k Keeper) GetBioLenParams(ctx sdk.Context) (bioParams params.BioLenParams) {
	k.paramSpace.Get(ctx, params.ParamStoreKeyMaxBioLen, &bioParams)
	return bioParams
}
