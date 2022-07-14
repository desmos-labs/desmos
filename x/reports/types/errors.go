package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidGenesis = sdkerrors.Register(ModuleName, 1, "invalid genesis state")
)
