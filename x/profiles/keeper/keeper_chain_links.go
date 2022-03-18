package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

// SaveChainLink stores the given chain link
func (k Keeper) SaveChainLink(ctx sdk.Context, link types.ChainLink) error {
	// Validate the chain link
	err := link.Validate()
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidChainLink, err.Error())
	}

	// Make sure the user has a profile
	if !k.HasProfile(ctx, link.User) {
		return sdkerrors.Wrap(types.ErrProfileNotFound, "a profile is required to link a chain")
	}

	// Validate the source address
	srcAddrData, err := types.UnpackAddressData(k.cdc, link.Address)
	if err != nil {
		return err
	}

	err = srcAddrData.Validate()
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidAddressData, err.Error())
	}

	// Verify the proof
	err = link.Proof.Verify(k.cdc, srcAddrData)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidProof, err.Error())
	}

	target := srcAddrData.GetValue()
	if _, found := k.GetChainLink(ctx, link.User, link.ChainConfig.Name, target); found {
		return types.ErrDuplicatedChainLink
	}

	// Set chain link -> address association
	store := ctx.KVStore(k.storeKey)
	key := types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, target)
	store.Set(key, types.MustMarshalChainLink(k.cdc, link))
	return nil
}

// GetChainLink returns the chain link for the given owner, chain name and target.
// If such link does not exist, returns false instead.
func (k Keeper) GetChainLink(ctx sdk.Context, owner, chainName, target string) (types.ChainLink, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ChainLinksStoreKey(owner, chainName, target)

	if !store.Has(key) {
		return types.ChainLink{}, false
	}

	return types.MustUnmarshalChainLink(k.cdc, store.Get(key)), true
}

// DeleteChainLink deletes the link associated with the given address and chain name
func (k Keeper) DeleteChainLink(ctx sdk.Context, owner, chainName, target string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.ChainLinksStoreKey(owner, chainName, target)
	if !store.Has(key) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"chain link between %s and %s for chain name %s not found",
			owner, target, chainName,
		)
	}
	store.Delete(types.ChainLinksStoreKey(owner, chainName, target))
	return nil
}

// DeleteAllUserChainLinks deletes all the chain links associated with the given user
func (k Keeper) DeleteAllUserChainLinks(ctx sdk.Context, user string) {
	var links []types.ChainLink
	k.IterateUserChainLinks(ctx, user, func(index int64, link types.ChainLink) (stop bool) {
		links = append(links, link)
		return false
	})

	store := ctx.KVStore(k.storeKey)
	for _, link := range links {
		address := link.Address.GetCachedValue().(types.AddressData)
		store.Delete(types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, address.GetValue()))
	}
}
