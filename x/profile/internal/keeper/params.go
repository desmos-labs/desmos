package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
)

// SetNameSurnameLenParams sets NameSurnameLenParams to the global param store
func (k Keeper) SetNameSurnameLenParams(ctx sdk.Context, nmParams models.NameSurnameLenParams) {
	k.paramSpace.Set(ctx, models.ParamStoreKeyNameSurnameLen, &nmParams)
}

// GetNameSurnameLenParams returns the current NameSurnameLenParams from the global param store
func (k Keeper) GetNameSurnameLenParams(ctx sdk.Context) (nmParams models.NameSurnameLenParams) {
	k.paramSpace.Get(ctx, models.ParamStoreKeyNameSurnameLen, &nmParams)
	return nmParams
}

// SetMonikerLenParams sets MonikerLenParams to the global param store
func (k Keeper) SetMonikerLenParams(ctx sdk.Context, monikerParams models.MonikerLenParams) {
	k.paramSpace.Set(ctx, models.ParamStoreKeyMonikerLen, &monikerParams)
}

// GetMonikerLenParams returns the current MonikerLenParams from the global param store
func (k Keeper) GetMonikerLenParams(ctx sdk.Context) (mParams models.MonikerLenParams) {
	k.paramSpace.Get(ctx, models.ParamStoreKeyMonikerLen, &mParams)
	return mParams
}

// SetBioLenParams sets the BioLenParams to the global param store
func (k Keeper) SetBioLenParams(ctx sdk.Context, bioParams models.BioLenParams) {
	k.paramSpace.Set(ctx, models.ParamStoreKeyMaxBioLen, &bioParams)
}

// GetBioLenParams returns the current BioLenParams from the global param store
func (k Keeper) GetBioLenParams(ctx sdk.Context) (bioParams models.BioLenParams) {
	k.paramSpace.Get(ctx, models.ParamStoreKeyMaxBioLen, &bioParams)
	return bioParams
}
