package keeper

import (
	"regexp"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryMarshaler
	paramSubspace paramstypes.Subspace
}

// NewKeeper creates new instances of the subspaces keeper
func NewKeeper(storeKey sdk.StoreKey, cdc codec.BinaryMarshaler, paramSpace paramstypes.Subspace) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramSubspace: paramSpace,
	}
}

// SaveSubspace saves the given subspace inside the current context.
// It assumes that the subspace has been validated already.
func (k Keeper) SaveSubspace(ctx sdk.Context, subspace types.Subspace) {
	store := ctx.KVStore(k.storeKey)
	key := types.SubspaceStoreKey(subspace.ID)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&subspace))
}

// DoesSubspaceExist checks if the subspace with the given id exists.
func (k Keeper) DoesSubspaceExist(ctx sdk.Context, subspaceID string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SubspaceStoreKey(subspaceID))
}

// GetSubspace returns the subspace associated with the given ID.
// If there is no subspace associated with the given ID the function will return an empty subspace and false.
func (k Keeper) GetSubspace(ctx sdk.Context, subspaceID string) (subspace types.Subspace, found bool) {
	if !k.DoesSubspaceExist(ctx, subspaceID) {
		return types.Subspace{}, false
	}

	store := ctx.KVStore(k.storeKey)
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.SubspaceStoreKey(subspaceID)), &subspace)
	return subspace, true
}

// GetAllSubspaces returns a list of all the subspaces that have been store inside the given context
func (k Keeper) GetAllSubspaces(ctx sdk.Context) []types.Subspace {
	var subspaces []types.Subspace
	k.IterateSubspaces(ctx, func(_ int64, subspace types.Subspace) (stop bool) {
		subspaces = append(subspaces, subspace)
		return false
	})

	return subspaces
}

// ValidateSubspace check if the given subspace is valid according to the current module params and
func (k Keeper) ValidateSubspace(ctx sdk.Context, subspace types.Subspace) error {
	params := k.GetParams(ctx)

	nameRegEx := regexp.MustCompile(params.NameParams.RegEx)
	minNameLen := params.NameParams.MinNameLength.Int64()
	maxNameLen := params.NameParams.MaxNameLength.Int64()

	nameLen := int64(len(subspace.Name))
	if !nameRegEx.MatchString(subspace.Name) {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceName, "invalid subspace name, it should match the following regEx %s", nameRegEx)
	}
	if nameLen < minNameLen {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceNameLength, "subspace name cannot be less than %d characters", minNameLen)
	}
	if nameLen > maxNameLen {
		return sdkerrors.Wrapf(types.ErrInvalidSubspaceNameLength, "subspace name cannot exceed %d characters", maxNameLen)
	}

	return subspace.Validate()
}

// AddAdminToSubspace insert the user inside the admins array of the given subspace if his not present.
// Returns an error if the user is already an admin.
func (k Keeper) AddAdminToSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	// check if the subspace exists and the admin is the actual admin of it
	subspace, err := k.checkSubspaceAndOwner(ctx, subspaceID, owner)
	if err != nil {
		return err
	}

	// check if the user we want to set as admin is already an admin
	if types.IsPresent(subspace.Admins, user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address : %s is already an admin", user)
	}

	subspace.Admins = append(subspace.Admins, user)
	k.SaveSubspace(ctx, subspace)
	return nil
}

// RemoveAdminFromSubspace remove the given admin from the given subspace.
// It returns error when the admin is not present inside the subspace.
func (k Keeper) RemoveAdminFromSubspace(ctx sdk.Context, subspaceID, user, owner string) error {
	// check if the subspace exists and the admin is the actual admin of it
	subspace, err := k.checkSubspaceAndOwner(ctx, subspaceID, owner)
	if err != nil {
		return err
	}

	// check if the user is not an admin
	if !types.IsPresent(subspace.Admins, user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address : %s is not an admin", user)
	}

	subspace.Admins = types.RemoveUser(subspace.Admins, user)
	k.SaveSubspace(ctx, subspace)
	return nil
}

// RegisterUserInSubspace register the user in the subspace with the given ID.
// It returns error when the user is already registered.
func (k Keeper) RegisterUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if types.IsPresent(subspace.RegisteredUsers, user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is already registered inside the subspace: %s", user, subspaceID)
	}

	subspace.RegisteredUsers = append(subspace.RegisteredUsers, user)
	k.SaveSubspace(ctx, subspace)
	return nil
}

// UnregisterUserFromSubspace unregister the user from the subspace with the given ID.
// It returns error when the user is not registered.
func (k Keeper) UnregisterUserFromSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if !types.IsPresent(subspace.RegisteredUsers, user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is not registered inside the subspace: %s", user, subspaceID)
	}

	subspace.RegisteredUsers = types.RemoveUser(subspace.RegisteredUsers, user)
	k.SaveSubspace(ctx, subspace)
	return nil
}

// BanUserInSubspace block the given user inside the given subspace.
// It returns error if the user is already blocked inside the subspace.
func (k Keeper) BanUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if types.IsPresent(subspace.BannedUsers, user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is already blocked inside the subspace: %s", user, subspaceID)
	}

	subspace.BannedUsers = append(subspace.BannedUsers, user)
	k.SaveSubspace(ctx, subspace)
	return nil
}

// UnbanUserInSubspace unblock the given user inside the given subspace.
// It returns error if the user is not blocked inside the subspace.
func (k Keeper) UnbanUserInSubspace(ctx sdk.Context, subspaceID, user, admin string) error {
	// check if the subspace exists and the admin is an actual admin
	subspace, err := k.checkSubspaceAndAdmin(ctx, subspaceID, admin)
	if err != nil {
		return err
	}

	// check if the user is already registered inside the subspace
	if !types.IsPresent(subspace.BannedUsers, user) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user with address : %s is not blocked inside the subspace: %s", user, subspaceID)
	}

	subspace.BannedUsers = types.RemoveUser(subspace.BannedUsers, user)
	k.SaveSubspace(ctx, subspace)
	return nil
}
