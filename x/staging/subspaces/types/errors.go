package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	// ErrInvalidSubspaceID is returned if a subspace id is not valid
	ErrInvalidSubspaceID = sdkerrors.Register(ModuleName, 1, "subspace id must be a valid SHA-256 hash")

	// ErrSubspaceAlreadyExistent is returned if a subspace already exists
	ErrSubspaceAlreadyExistent = sdkerrors.Register(ModuleName, 2, "the subspaces already exists")

	// ErrSubspaceNotFound is returned if a subspace doesn't exist
	ErrSubspaceNotFound = sdkerrors.Register(ModuleName, 3, "subspace not found")

	// ErrEmptySubspaceName is returned if a subspace name is empty or blank
	ErrEmptySubspaceName = sdkerrors.Register(ModuleName, 4, "subspace name cannot be empty or blank")

	// ErrAlreadySubspaceAdmin is returned if a subspace user is already an admin of a subspace
	ErrAlreadySubspaceAdmin = sdkerrors.Register(ModuleName, 5, "the user is already an admin")

	// ErrNotSubspaceAdmin is returned if a user is not a subspace admin
	ErrNotSubspaceAdmin = sdkerrors.Register(ModuleName, 6, "invalid subspace admin")

	// ErrNotRegisteredUserInSubspace is returned if a user is not registered inside a subspace
	ErrNotRegisteredUserInSubspace = sdkerrors.Register(ModuleName, 7, "user is not registered inside the subspace")

	ErrNotBannedUserInSubspace = sdkerrors.Register(ModuleName, 8, "the user is not banned inside the subspace")

	// ErrInvalidSubspaceOwner is returned if a subspace owner is not valid
	ErrInvalidSubspaceOwner = sdkerrors.Register(ModuleName, 9, "invalid subspace owner")

	// ErrPermissionDenied is returned if a user is banned or not registered inside a given subspace
	ErrPermissionDenied = sdkerrors.Register(ModuleName, 10, "permission denied for user")
)
