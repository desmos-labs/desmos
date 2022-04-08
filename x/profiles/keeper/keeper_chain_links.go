package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

// SaveChainLink stores the given chain link
// Chain links are stored using two keys:
// 1. ChainLinkStoreKey (user + chain name + target) -> types.ChainLink
// 2. ChainLinkOwnerKey (chain name + target + user) -> 0x01
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
	err = link.Proof.Verify(k.cdc, k.legacyAmino, link.User, srcAddrData)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvalidProof, err.Error())
	}

	target := srcAddrData.GetValue()
	if _, found := k.GetChainLink(ctx, link.User, link.ChainConfig.Name, target); found {
		return types.ErrDuplicatedChainLink
	}

	// Store the data
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, target), types.MustMarshalChainLink(k.cdc, link))
	store.Set(types.ChainLinkOwnerKey(link.ChainConfig.Name, target, link.User), []byte{0x01})

	k.AfterChainLinkSaved(ctx, link)

	return nil
}

// HasChainLink tells whether the given chain link exists or not
func (k Keeper) HasChainLink(ctx sdk.Context, owner, chainName, target string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.ChainLinksStoreKey(owner, chainName, target))
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
func (k Keeper) DeleteChainLink(ctx sdk.Context, link types.ChainLink) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()))
	store.Delete(types.ChainLinkOwnerKey(link.ChainConfig.Name, link.GetAddressData().GetValue(), link.User))

	k.AfterChainLinkDeleted(ctx, link)
}

// DeleteAllUserChainLinks deletes all the chain links associated with the given user
func (k Keeper) DeleteAllUserChainLinks(ctx sdk.Context, user string) {
	var links []types.ChainLink
	k.IterateUserChainLinks(ctx, user, func(index int64, link types.ChainLink) (stop bool) {
		links = append(links, link)
		return false
	})

	for _, link := range links {
		k.DeleteChainLink(ctx, link)
	}
}
