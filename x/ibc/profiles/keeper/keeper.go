package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	ibctypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
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

// GetLinkPubkey returns the pubkey corresponding to the given account
func (k Keeper) GetAccountPubKey(ctx sdk.Context, acc sdk.AccAddress) (cryptotypes.PubKey, error) {
	return k.accountKeeper.GetPubKey(ctx, acc)
}
