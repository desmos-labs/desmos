package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidGenesis = sdkerrors.Register(ModuleName, 1, "invalid genesis state")
	ErrInvalidPost    = sdkerrors.Register(ModuleName, 2, "invalid post")
)
