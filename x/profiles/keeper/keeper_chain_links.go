package keeper

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/profiles/types"
)

// SaveChainLink stores the given chain link
// The first chain link of each chain for the owner will be set as default external address
// Chain links are stored using two keys:
// 1. ChainLinkStoreKey (user + chain name + target) -> types.ChainLink
// 2. ChainLinkOwnerKey (chain name + target + user) -> 0x01
func (k Keeper) SaveChainLink(ctx sdk.Context, link types.ChainLink) error {
	// Validate the chain link
	err := link.Validate()
	if err != nil {
		return errors.Wrap(types.ErrInvalidChainLink, err.Error())
	}

	// Make sure the user has a profile
	if !k.HasProfile(ctx, link.User) {
		return errors.Wrap(types.ErrProfileNotFound, "a profile is required to link a chain")
	}

	// Validate the source address
	srcAddrData, err := types.UnpackAddressData(k.cdc, link.Address)
	if err != nil {
		return err
	}

	err = srcAddrData.Validate()
	if err != nil {
		return errors.Wrap(types.ErrInvalidAddressData, err.Error())
	}

	// Verify the proof
	err = link.Proof.Verify(k.cdc, k.legacyAmino, link.User, srcAddrData)
	if err != nil {
		return errors.Wrap(types.ErrInvalidProof, err.Error())
	}

	target := srcAddrData.GetValue()
	if _, found := k.GetChainLink(ctx, link.User, link.ChainConfig.Name, target); found {
		return types.ErrDuplicatedChainLink
	}

	// Store the data
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, target), types.MustMarshalChainLink(k.cdc, link))
	store.Set(types.ChainLinkOwnerKey(link.ChainConfig.Name, target, link.User), []byte{0x01})

	// Set the link as default external address if there is no default external address
	if !k.HasDefaultExternalAddress(ctx, link.User, link.ChainConfig.Name) {
		k.SaveDefaultExternalAddress(ctx, link.User, link.ChainConfig.Name, srcAddrData.GetValue())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSavedChainLink,
			sdk.NewAttribute(types.AttributeKeyChainLinkExternalAddress, link.GetAddressData().GetValue()),
			sdk.NewAttribute(types.AttributeKeyChainLinkChainName, link.ChainConfig.Name),
			sdk.NewAttribute(types.AttributeKeyChainLinkOwner, link.User),
		),
	)

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

	// Update the default external address to be the oldest link if the deleted link is default exnternal address
	if k.isDefaultExternalAddress(ctx, link) {
		k.updateOwnerDefaultExternalAddress(ctx, link.User, link.ChainConfig.Name)
	}

	k.AfterChainLinkDeleted(ctx, link)
}

// DeleteAllUserChainLinks deletes all the chain links associated with the given user
func (k Keeper) DeleteAllUserChainLinks(ctx sdk.Context, user string) {
	var links []types.ChainLink
	k.IterateUserChainLinks(ctx, user, func(link types.ChainLink) (stop bool) {
		links = append(links, link)
		return false
	})

	for _, link := range links {
		k.DeleteChainLink(ctx, link)
	}
}

// getOldestUserChainLink returns the oldest chain link of the given owner associated to the given chain name.
// If such chain link does not exist, returns false instead.
func (k Keeper) getOldestUserChainLink(ctx sdk.Context, owner, chainName string) (types.ChainLink, bool) {
	var oldestLink types.ChainLink
	found := false
	k.IterateUserChainLinksByChain(ctx, owner, chainName, func(link types.ChainLink) (stop bool) {
		if !found {
			oldestLink = link
			found = true
			return false
		}

		if link.CreationTime.Before(oldestLink.CreationTime) {
			oldestLink = link
		}
		return false
	})

	return oldestLink, found
}

// updateOwnerDefaultExternalAddress updates the default external address of the given chain name for the given owner
// It must be performed after deleting the default external address chain link
func (k Keeper) updateOwnerDefaultExternalAddress(ctx sdk.Context, owner, chainName string) {
	store := ctx.KVStore(k.storeKey)
	link, found := k.getOldestUserChainLink(ctx, owner, chainName)
	if !found {
		// If the owner has no chain link on the given chain name, then delete the key
		store.Delete(types.DefaultExternalAddressKey(owner, chainName))
		return
	}

	srcAddrData, err := types.UnpackAddressData(k.cdc, link.Address)
	if err != nil {
		panic(err)
	}
	k.SaveDefaultExternalAddress(ctx, owner, chainName, srcAddrData.GetValue())
}

// SaveDefaultExternalAddress stores the given address as a default external address
func (k Keeper) SaveDefaultExternalAddress(ctx sdk.Context, owner, chainName, target string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.DefaultExternalAddressKey(owner, chainName), []byte(target))
}

// HasDefaultExternalAddress tells whether the given owner has a default external address on the given chain name or not
func (k Keeper) HasDefaultExternalAddress(ctx sdk.Context, owner, chainName string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.DefaultExternalAddressKey(owner, chainName))
}

// isDefaultExternalAddress tells whether the given chain link is a default external address or not
func (k Keeper) isDefaultExternalAddress(ctx sdk.Context, link types.ChainLink) bool {
	store := ctx.KVStore(k.storeKey)
	addressBz := store.Get(types.DefaultExternalAddressKey(link.User, link.ChainConfig.Name))
	return string(addressBz) == link.GetAddressData().GetValue()
}
