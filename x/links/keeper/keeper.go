package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	"github.com/desmos-labs/desmos/x/links/types"
)

type Keeper struct {
	cdc      codec.Marshaler
	storeKey sdk.StoreKey
	memKey   sdk.StoreKey

	channelKeeper types.ChannelKeeper
	portKeeper    types.PortKeeper
	scopedKeeper  capabilitykeeper.ScopedKeeper
	accountKeeper authkeeper.AccountKeeper
}

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	accountKeeper authkeeper.AccountKeeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
		accountKeeper: accountKeeper,
	}
}

// StoreLink sotres the given link inside the current context.
// It assumes that the given link has already been validated.
func (k Keeper) StoreLink(ctx sdk.Context, link types.Link) error {
	store := ctx.KVStore(k.storeKey)
	key := types.LinkStoreKey(link.SourceAddress)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&link))
	return nil
}

// GetAllLinks returns the list of all the links that have been stored inside the given context
func (k Keeper) GetAllLinks(ctx sdk.Context) []types.Link {
	var links []types.Link

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LinksStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var link types.Link
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &link)
		links = append(links, link)
	}
	return links
}

func (k Keeper) GetPubKey(ctx sdk.Context, acc sdk.AccAddress) (cryptotypes.PubKey, error) {
	return k.accountKeeper.GetPubKey(ctx, acc)
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
