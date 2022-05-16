package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// SetNextSectionID sets the next section id for the specific subspace
func (k Keeper) SetNextSectionID(ctx sdk.Context, subspaceID uint64, sectionID uint32) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NextSectionIDStoreKey(subspaceID), types.GetSectionIDBytes(sectionID))
}

// HasNextSectionID tells whether the next section id key exists for the given subspace
func (k Keeper) HasNextSectionID(ctx sdk.Context, subspaceID uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NextSectionIDStoreKey(subspaceID))
}

// GetNextSectionID gets the next section id for the subspace having the given id
func (k Keeper) GetNextSectionID(ctx sdk.Context, subspaceID uint64) (sectionID uint32, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NextSectionIDStoreKey(subspaceID))
	if bz == nil {
		return 0, sdkerrors.Wrapf(types.ErrInvalidGenesis, "initial section id hasn't been set for subspace %d", subspaceID)
	}

	sectionID = types.GetSectionIDFromBytes(bz)
	return sectionID, nil
}

// DeleteNextSectionID deletes the next section id key for the given subspace from the store
func (k Keeper) DeleteNextSectionID(ctx sdk.Context, subspaceID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.NextSectionIDStoreKey(subspaceID))
}

// --------------------------------------------------------------------------------------------------------------------

// SaveSection saves the given section inside the current context
func (k Keeper) SaveSection(ctx sdk.Context, section types.Section) {
	store := ctx.KVStore(k.storeKey)

	// Save the section
	store.Set(types.SectionStoreKey(section.SubspaceID, section.ID), k.cdc.MustMarshal(&section))

	k.Logger(ctx).Info("section saved", "subspace id", section.SubspaceID, "section id", section.ID)
	k.AfterSubspaceSectionSaved(ctx, section.SubspaceID, section.ID)
}

// HasSection tells whether the section having the given id exists inside the provided subspace
func (k Keeper) HasSection(ctx sdk.Context, subspaceID uint64, sectionID uint32) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SectionStoreKey(subspaceID, sectionID))
}

// GetSection returns the section having the given id from the subspace with the provided id.
// If there is no section associated with the given id the function will return an empty section and false.
func (k Keeper) GetSection(ctx sdk.Context, subspaceID uint64, sectionID uint32) (section types.Section, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.SectionStoreKey(subspaceID, sectionID))
	if bz == nil {
		return types.Section{}, false
	}

	k.cdc.MustUnmarshal(bz, &section)
	return section, true
}

// DeleteSection deletes the section having the given id from the subspace with the provided id
func (k Keeper) DeleteSection(ctx sdk.Context, subspaceID uint64, sectionID uint32) {
	store := ctx.KVStore(k.storeKey)

	// Remove all the groups within this section
	k.IterateSubspaceUserGroups(ctx, subspaceID, func(index int64, group types.UserGroup) (stop bool) {
		if group.SectionID == sectionID {
			k.DeleteUserGroup(ctx, group.SubspaceID, group.ID)
		}
		return false
	})

	// Remove all the permissions set inside the section
	k.IterateSectionUserPermissions(ctx, subspaceID, sectionID, func(index int64, user sdk.AccAddress, permission types.Permission) (stop bool) {
		k.RemoveUserPermissions(ctx, subspaceID, sectionID, user)
		return false
	})

	// Delete the section
	store.Delete(types.SectionStoreKey(subspaceID, sectionID))

	k.Logger(ctx).Info("section deleted", "subspace id", subspaceID, "section id", sectionID)
	k.AfterSubspaceSectionDeleted(ctx, subspaceID, sectionID)
}
