package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

type Keeper struct {
	sk  types.SubspacesKeeper
	tfk types.TokenFactoryKeeper
	bk  types.BankKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper returns a new instance of the x/tokenfactory keeper
func NewKeeper(
	sk types.SubspacesKeeper,
	tfk types.TokenFactoryKeeper,
	bk types.BankKeeper,
	authority string,
) Keeper {
	return Keeper{
		bk:  bk,
		sk:  sk,
		tfk: tfk,

		authority: authority,
	}
}

// Logger returns a logger for the x/tokenfactory module
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
