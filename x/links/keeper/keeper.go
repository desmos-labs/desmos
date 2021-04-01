package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
}

func NewKeeper(
	cdc codec.Marshaler,
	storeKey,
	memKey sdk.StoreKey,
	channelKeeper types.ChannelKeeper,
	portKeeper types.PortKeeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		channelKeeper: channelKeeper,
		portKeeper:    portKeeper,
		scopedKeeper:  scopedKeeper,
	}
}

func (k Keeper) GetLinks(ctx sdk.Context) []types.Link {
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

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
