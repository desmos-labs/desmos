package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// RootCodespace is the codespace for all errors defined in this package
const RootCodespace = "subspaces"

var (
	// ErrInvalidSubspace is returned if a subspace id is not valid
	ErrInvalidSubspace = sdkerrors.Register(RootCodespace, 1, "invalid subspace")

	// ErrInvalidSubspaceName is returned if a subspace name is empty or blank
	ErrInvalidSubspaceName = sdkerrors.Register(RootCodespace, 2, "invalid subspace name")
)
