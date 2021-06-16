package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// StoreChainLink stores the given chain link
func (k Keeper) StoreChainLink(ctx sdk.Context, link types.ChainLink) error {

	// Validate the chain link
	err := link.Validate()
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidChainLink, err.Error())
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
	err = link.Proof.Verify(k.cdc)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidProof, err.Error())
	}

	target := srcAddrData.GetAddress()
	if _, found := k.GetChainLink(ctx, link.User, link.ChainConfig.Name, target); found {
		return types.ErrDuplicatedChainLink
	}

	// Make sure the user has a profile
	_, found, err := k.GetProfile(ctx, link.User)
	if err != nil {
		return err
	}

	if !found {
		return sdkerrors.Wrap(types.ErrProfileNotFound, "target user does not have a profile")
	}

	// Set chain link -> address association
	store := ctx.KVStore(k.storeKey)
	key := types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, target)
	store.Set(key, types.MustMarshalChainLink(k.cdc, link))
	return nil
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

func (k Keeper) GetChainLink(ctx sdk.Context, owner, chainName, target string) (types.ChainLink, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.ChainLinksStoreKey(owner, chainName, target)

	if !store.Has(key) {
		return types.ChainLink{}, false
	}

	return types.MustUnmarshalChainLink(k.cdc, store.Get(key)), true
}

// GetAllChainLinks allows to returns the list of all stored chain links
func (k Keeper) GetAllChainLinks(ctx sdk.Context) []types.ChainLink {
	var links []types.ChainLink
	k.IterateChainLinks(ctx, func(_ int64, link types.ChainLink) (stop bool) {
		links = append(links, link)
		return false
	})
	return links
}
