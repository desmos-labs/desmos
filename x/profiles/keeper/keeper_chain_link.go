package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// StoreChainLink stores the given chain link inside the current context.
// It assumes that the given chain link has already been validated.
func (k Keeper) StoreChainLink(ctx sdk.Context, user string, link types.ChainLink) error {

	srcAddrData, err := types.UnpackAddress(k.cdc, link.Address)
	if err != nil {
		return err
	}

	target := srcAddrData.GetAddressString()
	if _, found := k.GetAccountByChainLink(ctx, link.ChainConfig.Name, target); found {
		return fmt.Errorf("chain link already exists")
	}

	// check target address has a profile or not
	if link.ChainConfig.Name == "desmos" {
		if _, err := sdk.AccAddressFromBech32(target); err != nil {
			return err
		}
		_, found, err := k.GetProfile(ctx, target)
		if err != nil {
			return err
		}
		if found {
			return fmt.Errorf("the target address has profiles already")
		}
	}

	// Check if address has the profile
	profile, found, err := k.GetProfile(ctx, user)
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("address does not have any profile")
	}
	// Store chain link to the profile
	profile.ChainsLinks = append(profile.ChainsLinks, link)
	if err := k.StoreProfile(ctx, profile); err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.ChainsLinksStoreKey(link.ChainConfig.Name, target)
	store.Set(key, profile.GetAddress())
	return nil
}

// GetAccountByChainLink returns the account corresponding to the given address and the given chain name inside the current context.
func (k Keeper) GetAccountByChainLink(ctx sdk.Context, chainName string, address string) (account sdk.AccAddress, found bool) {
	store := ctx.KVStore((k.storeKey))

	bz := store.Get(types.ChainsLinksStoreKey(chainName, address))
	if bz == nil {
		return nil, false
	}
	return account, true
}

// DeleteChainLink allows to delete a link associated with the given address and chain name inside the current context.
// It assumes that the related link exists.
func (k Keeper) DeleteChainLink(ctx sdk.Context, owner, chainName, target string) error {
	// Check if address has the profile and get the profile
	profile, found, err := k.GetProfile(ctx, owner)
	if err != nil {
		return err
	}
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ("non existent profile on owner address"))
	}

	doesLinkExists := false
	// Try to find the target link
	for index, link := range profile.ChainsLinks {

		var address types.AddressData
		if err := k.cdc.UnpackAny(link.Address, &address); err != nil {
			return err
		}
		if link.ChainConfig.Name == chainName && address.GetAddressString() == target {
			doesLinkExists = true
			newChainsLinks := append(profile.ChainsLinks[:index], profile.ChainsLinks[index+1:]...)
			profile.ChainsLinks = newChainsLinks
			// Update profile status
			if err = k.StoreProfile(ctx, profile); err != nil {
				return err
			}
			break
		}
	}

	if !doesLinkExists {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, ("non existent target chain link in the profile"))
	}

	store := ctx.KVStore(k.storeKey)
	key := types.ChainsLinksStoreKey(chainName, target)
	store.Delete(key)
	return nil
}

// GetAllAccountsByChainsLink returns a list of all the accounts in chain link store inside the given context.
func (k Keeper) GetAllAccountsByChainLink(ctx sdk.Context) []sdk.AccAddress {
	accounts := []sdk.AccAddress{}
	k.IterateAccountsByChainLink(ctx, func(index int64, account sdk.AccAddress) (stop bool) {
		accounts = append(accounts, account)
		return false
	})
	return accounts
}
