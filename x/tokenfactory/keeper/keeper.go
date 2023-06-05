package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/v15/x/tokenfactory/keeper"

	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

type Keeper struct {
	sk   types.SubspacesKeeper
	tkfk types.TokenFactoryKeeper
	ak   types.AccountKeeper
	bk   types.BankKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper returns a new instance of the x/tokenfactory keeper
func NewKeeper(
	ak types.AccountKeeper,
	bk types.BankKeeper,

	sk types.SubspacesKeeper,
	tkfk keeper.Keeper,

	authority string,
) Keeper {
	return Keeper{
		ak:   ak,
		bk:   bk,
		sk:   sk,
		tkfk: tkfk,

		authority: authority,
	}
}

// Logger returns a logger for the x/tokenfactory module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
