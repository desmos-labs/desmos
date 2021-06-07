package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// StoreChainLink stores the given chain link
func (k Keeper) StoreChainLink(ctx sdk.Context, user string, link types.ChainLink) error {
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
	if _, found := k.GetAccountByChainLink(ctx, link.ChainConfig.Name, target); found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "chain link already existing")
	}

	// Make sure the user has a profile
	profile, found, err := k.GetProfile(ctx, user)
	if err != nil {
		return err
	}

	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "target user does not have a profile")
	}

	// Update the profile chain links
	profile.ChainsLinks = append(profile.ChainsLinks, link)
	err = k.StoreProfile(ctx, profile)
	if err != nil {
		return err
	}

	// Set chain link -> address association
	store := ctx.KVStore(k.storeKey)
	key := types.ChainsLinksStoreKey(link.ChainConfig.Name, target)
	store.Set(key, profile.GetAddress())
	return nil
}

// DeleteChainLink deletes the link associated with the given address and chain name
func (k Keeper) DeleteChainLink(ctx sdk.Context, owner, chainName, target string) error {
	// Check if address has the profile and get the profile
	profile, found, err := k.GetProfile(ctx, owner)
	if err != nil {
		return err
	}

	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "user does not have a profile")
	}

	// Try to remove the target link
	newLinks, found := types.RemoveChainLinkIfPresent(profile.ChainsLinks, chainName, target)
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "chain link for given chain name and address not found")
	}

	// Update the profile
	profile.ChainsLinks = newLinks
	err = k.StoreProfile(ctx, profile)
	if err != nil {
		return err
	}

	// Update the chain link -> address association
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ChainsLinksStoreKey(chainName, target))
	return nil
}

// GetAccountByChainLink returns the account corresponding to the given address and the given chain name
func (k Keeper) GetAccountByChainLink(ctx sdk.Context, chainName string, address string) (sdk.AccAddress, bool) {
	key := types.ChainsLinksStoreKey(chainName, address)
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		return nil, false
	}

	return store.Get(key), true
}
