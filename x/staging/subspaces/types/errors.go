package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrInvalidSubspaceID is returned if a subspace id is not valid
	ErrInvalidSubspaceID = sdkerrors.Register(ModuleName, 1, "invalid subspace id")

	// ErrInvalidSubspaceName is returned if a subspace name is empty or blank
	ErrInvalidSubspaceName = sdkerrors.Register(ModuleName, 2, "invalid subspace name")

	// ErrInvalidSubspaceAdmin is returned if a subspace admin is not valid
	ErrInvalidSubspaceAdmin = sdkerrors.Register(ModuleName, 3, "invalid subspace admin")

	// ErrInvalidSubspaceOwner is returned if a subspace owner is not valid
	ErrInvalidSubspaceOwner = sdkerrors.Register(ModuleName, 4, "invalid subspace owner")
)
