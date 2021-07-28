package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrInvalidSubspaceID is returned if a subspace id is not valid
	ErrInvalidSubspaceID = sdkerrors.Register(ModuleName, 1, "subspace id must be a valid SHA-256 hash")

	// ErrDuplicatedSubspace is returned if a subspace already exists
	ErrDuplicatedSubspace = sdkerrors.Register(ModuleName, 2, "the subspace already exists")

	// ErrSubspaceNotFound is returned if a subspace doesn't exist
	ErrSubspaceNotFound = sdkerrors.Register(ModuleName, 3, "subspace not found")

	// ErrInvalidSubspace is returned if a subspace is invalid
	ErrInvalidSubspace = sdkerrors.Register(ModuleName, 4, "subspace is invalid")

	// ErrDuplicatedAdmin is returned if a subspace user is already an admin of a subspace
	ErrDuplicatedAdmin = sdkerrors.Register(ModuleName, 5, "the user is already an admin")

	// ErrInvalidAdmin is returned if a user is not a subspace admin
	ErrInvalidAdmin = sdkerrors.Register(ModuleName, 6, "invalid subspace admin")

	// ErrUserNotFound is returned if a user is not registered inside a subspace
	ErrUserNotFound = sdkerrors.Register(ModuleName, 7, "user not registered inside the subspace")

	// ErrInvalidSubspaceOwner is returned if a subspace owner is not valid
	ErrInvalidSubspaceOwner = sdkerrors.Register(ModuleName, 9, "invalid subspace owner")

	// ErrPermissionDenied is returned if a user is banned or not registered inside a given subspace
	ErrPermissionDenied = sdkerrors.Register(ModuleName, 10, "permission denied for user")
)
