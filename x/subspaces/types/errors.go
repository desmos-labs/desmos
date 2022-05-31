package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrPermissionDenied is returned if a user cannot perform a specific action inside a subspace
	ErrPermissionDenied = sdkerrors.Register(ModuleName, 1, "permission denied for user")
	ErrInvalidGenesis   = sdkerrors.Register(ModuleName, 2, "invalid genesis state")
	ErrInvalidSection   = sdkerrors.Register(ModuleName, 3, "invalid section")
)
