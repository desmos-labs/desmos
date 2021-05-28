package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrInvalidSubspaceID is returned if a subspace id is not valid
	ErrInvalidSubspaceID = sdkerrors.Register(ModuleName, 1, "invalid subspace id")

	// ErrInvalidSubspaceName is returned if a subspace name is empty or blank
	ErrInvalidSubspaceName = sdkerrors.Register(ModuleName, 2, "invalid subspace name")

	// ErrInvalidSubspaceNameLength is returned if a subspace name doesn't match the subspaces name params criteria
	ErrInvalidSubspaceNameLength = sdkerrors.Register(ModuleName, 3, "invalid subspace name length")
)
