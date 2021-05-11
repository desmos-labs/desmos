package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	"github.com/desmos-labs/desmos/x/links/types"
)

type Keeper struct {
	cdc      codec.BinaryMarshaler
	storeKey sdk.StoreKey

	channelKeeper ibctypes.ChannelKeeper
	portKeeper    ibctypes.PortKeeper
	scopedKeeper  capabilitykeeper.ScopedKeeper
	accountKeeper authkeeper.AccountKeeper
}

func NewKeeper(
	cdc codec.BinaryMarshaler,
	storeKey sdk.StoreKey,
	channelKeeper ibctypes.ChannelKeeper,
	portKeeper ibctypes.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	accountKeeper authkeeper.AccountKeeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
		accountKeeper: accountKeeper,
	}
}

// StoreLink stores the given link inside the current context.
// It assumes that the given link has already been validated.
// If the source address has already been inserted, nothing will be changed.
func (k Keeper) StoreLink(ctx sdk.Context, link types.Link) error {
	store := ctx.KVStore(k.storeKey)
	key := types.LinkStoreKey(link.SourceAddress)
	if store.Has(key) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "source address already exists")
	}
	store.Set(key, k.cdc.MustMarshalBinaryBare(&link))
	return nil
}

// GetLink returns the link corresponding to the given address inside the current context.
func (k Keeper) GetLink(ctx sdk.Context, address string) (link types.Link, found bool) {
	store := ctx.KVStore((k.storeKey))

	bz := store.Get(types.LinkStoreKey(address))
	if bz != nil {
		k.cdc.MustUnmarshalBinaryBare(bz, &link)
		return link, true
	}
	return types.Link{}, false
}

// GetAllLinks returns the list of all the links that have been stored inside the given context
func (k Keeper) GetAllLinks(ctx sdk.Context) []types.Link {
	var links []types.Link

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LinkStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var link types.Link
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &link)
		links = append(links, link)
	}
	return links
}

// GetLinkPubkey returns the pubkey corresponding to the given account
func (k Keeper) GetLinkPubKey(ctx sdk.Context, acc sdk.AccAddress) (cryptotypes.PubKey, error) {
	return k.accountKeeper.GetPubKey(ctx, acc)
}
