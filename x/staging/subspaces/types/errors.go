package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// RootCodespace is the codespace for all errors defined in this package
const RootCodespace = "subspaces"

var (
	// ErrInvalidSubspaceName is returned if a post subspace name is not valid
	ErrInvalidSubspaceName = sdkerrors.Register(RootCodespace, 1, "invalid subspace name")

	// ErrInvalidSubspaceId is returned if a post subspace id is not valid
	ErrInvalidSubspaceId = sdkerrors.Register(RootCodespace, 2, "invalid subspace id")
)
