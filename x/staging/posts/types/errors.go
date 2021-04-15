package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// RootCodespace is the codespace for all errors defined in this package
const RootCodespace = "posts"

var (
	// ErrInvalidPostID is returned if we cannot parse a post id
	ErrInvalidPostID = sdkerrors.Register(RootCodespace, 1, "invalid post id")

	// ErrInvalidSubspace is returned if a post subspace is not valid
	ErrInvalidSubspace = sdkerrors.Register(RootCodespace, 2, "invalid subspace")

	// ErrInvalidReactionCode is returned if we cannot validate a reaction short code
	ErrInvalidReactionCode = sdkerrors.Register(RootCodespace, 3,
		"invalid reaction shortcode (it must only contains a-z, 0-9, - and _ and must start and end with a ':')")
)
