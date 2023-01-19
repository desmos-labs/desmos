package keeper

import (
	"fmt"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants registers all subspaces invariants
func RegisterInvariants(ir sdk.InvariantRegistry, keeper Keeper) {
	ir.RegisterRoute(types.ModuleName, "valid-subspaces",
		ValidSubspacesInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-sections",
		ValidSectionsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-user-groups",
		ValidUserGroupsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-user-groups-members",
		ValidUserGroupMembersInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-user-permissions",
		ValidUserPermissionsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-user-grants",
		ValidUserGrantsInvariant(keeper))
	ir.RegisterRoute(types.ModuleName, "valid-group-grants",
		ValidGroupGrantsInvariant(keeper))
}

// --------------------------------------------------------------------------------------------------------------------

// ValidSubspacesInvariant checks that all the subspaces are valid
func ValidSubspacesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidSubspaces []types.Subspace
		k.IterateSubspaces(ctx, func(subspace types.Subspace) (stop bool) {
			invalid := false

			nextSubspaceID, err := k.GetSubspaceID(ctx)
			if err != nil {
				invalid = true
			}

			// Make sure the subspace id is never higher than the next one
			if subspace.ID >= nextSubspaceID {
				invalid = true
			}

			// Check next section id
			if !k.HasNextSectionID(ctx, subspace.ID) {
				invalid = true
			}

			// Check the next group id
			if !k.HasNextGroupID(ctx, subspace.ID) {
				invalid = true
			}

			// Validate the subspace
			err = subspace.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidSubspaces = append(invalidSubspaces, subspace)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid subspaces",
			fmt.Sprintf("the following subspaces are invalid:\n %s", FormatOutputSubspaces(invalidSubspaces)),
		), invalidSubspaces != nil
	}
}

// FormatOutputSubspaces concatenate the subspaces given into a unique string
func FormatOutputSubspaces(subspaces []types.Subspace) (outputSubspaces string) {
	for _, subspace := range subspaces {
		outputSubspaces += fmt.Sprintf("%d\n", subspace.ID)
	}
	return outputSubspaces
}

// --------------------------------------------------------------------------------------------------------------------

// ValidSectionsInvariant checks that all the sections are valid
func ValidSectionsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidSections []types.Section

		k.IterateSections(ctx, func(section types.Section) (stop bool) {
			invalid := false

			// Check the subspace existence
			if !k.HasSubspace(ctx, section.SubspaceID) {
				invalid = true
			}

			nextSectionID, err := k.GetNextSectionID(ctx, section.SubspaceID)
			if err != nil {
				invalid = true
			}

			// Make sure the section id is never equal or higher than the next section id
			if section.ID >= nextSectionID {
				invalid = true
			}

			// Check the parent section
			if !k.HasSection(ctx, section.SubspaceID, section.ParentID) {
				invalid = true
			}

			// Validate the section
			err = section.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidSections = append(invalidSections, section)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid sections",
			fmt.Sprintf("the following sections are invalid:\n%s", formatOutputSections(invalidSections)),
		), invalidSections != nil
	}
}

// formatOutputSections concatenates the given sections info into a string
func formatOutputSections(sections []types.Section) (output string) {
	for _, section := range sections {
		output += fmt.Sprintf("SubspaceID: %d, SectionID: %d\n", section.SubspaceID, section.ID)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidUserGroupsInvariant checks that all the subspaces are valid
func ValidUserGroupsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidUserGroups []types.UserGroup
		k.IterateUserGroups(ctx, func(group types.UserGroup) (stop bool) {
			invalid := false

			// Check subspace existence
			if !k.HasSubspace(ctx, group.SubspaceID) {
				invalid = true
			}

			// Check section existence
			if !k.HasSection(ctx, group.SubspaceID, group.SectionID) {
				invalid = true
			}

			nextGroupID, err := k.GetNextGroupID(ctx, group.SubspaceID)
			if err != nil {
				invalid = true
			}

			// Make sure the group id is always lower than the next one
			if group.ID >= nextGroupID {
				invalid = true
			}

			// Validate the group
			err = group.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidUserGroups = append(invalidUserGroups, group)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid user groups",
			fmt.Sprintf("the following user groups are invalid:\n %s", formatOutputUserGroups(invalidUserGroups)),
		), invalidUserGroups != nil
	}
}

// formatOutputUserGroups concatenates the given subspaces info given into a string
func formatOutputUserGroups(groups []types.UserGroup) (outputUserGroups string) {
	for _, group := range groups {
		outputUserGroups += fmt.Sprintf("SubspaceID: %d, GroupID: %d\n", group.SubspaceID, group.ID)
	}
	return outputUserGroups
}

// --------------------------------------------------------------------------------------------------------------------

// ValidUserGroupMembersInvariant checks that all the user group members are valid
func ValidUserGroupMembersInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidMembers []types.UserGroupMemberEntry
		k.IterateUserGroupsMembers(ctx, func(entry types.UserGroupMemberEntry) (stop bool) {
			invalid := false

			// Check subspace existence
			if !k.HasSubspace(ctx, entry.SubspaceID) {
				invalid = true
			}

			// Check the group existence
			if !k.HasUserGroup(ctx, entry.SubspaceID, entry.GroupID) {
				invalid = true
			}

			// Validate the entry only if the group id is not 0, as this will return an error
			err := entry.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidMembers = append(invalidMembers, entry)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid user group members",
			fmt.Sprintf("the following user group members entries are invalid:\n%s", formatOutputUserGroupsMembers(invalidMembers)),
		), invalidMembers != nil
	}
}

// formatOutputUserGroupsMembers concatenates the given user group members data into a string
func formatOutputUserGroupsMembers(members []types.UserGroupMemberEntry) (output string) {
	for _, entry := range members {
		output += fmt.Sprintf("SubspaceID: %d, GroupID: %d, Member: %s\n", entry.SubspaceID, entry.GroupID, entry.User)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidUserPermissionsInvariant checks that all the user permission entries are valid
func ValidUserPermissionsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidPermissionsEntries []types.UserPermission
		k.IterateUserPermissions(ctx, func(entry types.UserPermission) (stop bool) {
			invalid := false

			// Check subspace existence
			if !k.HasSubspace(ctx, entry.SubspaceID) {
				invalid = true
			}

			// Check section existence
			if !k.HasSection(ctx, entry.SubspaceID, entry.SectionID) {
				invalid = true
			}

			// Validate the entry
			err := entry.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidPermissionsEntries = append(invalidPermissionsEntries, entry)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid user permissions",
			fmt.Sprintf("the following user permissions are invalid:\n%s", formatOutputUserPermissions(invalidPermissionsEntries)),
		), invalidPermissionsEntries != nil
	}
}

// formatOutputUserPermissions concatenates the given permission entries into a string
func formatOutputUserPermissions(entries []types.UserPermission) (output string) {
	for _, entry := range entries {
		output += fmt.Sprintf("SubspaceID: %d, SectionID: %d, User: %s, Permissions: %s\n", entry.SubspaceID, entry.SectionID, entry.User, entry.Permissions)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidUserGrantsInvariant checks that all the user grants are valid
func ValidUserGrantsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidEntries []types.Grant
		k.IterateUserGrants(ctx, func(entry types.Grant) (stop bool) {
			invalid := false

			// Check subspace existence
			if !k.HasSubspace(ctx, entry.SubspaceID) {
				invalid = true
			}

			// Validate the entry
			err := entry.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidEntries = append(invalidEntries, entry)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid user grants",
			fmt.Sprintf("the following user grants are invalid:\n%s", formatOutputUserGrants(invalidEntries)),
		), invalidEntries != nil
	}
}

// formatOutputUserGrants concatenates the given user grants into a string
func formatOutputUserGrants(entries []types.Grant) (output string) {
	for _, entry := range entries {
		output += fmt.Sprintf("SubspaceID: %d, Granter: %s, Grantee: %s\n", entry.SubspaceID, entry.Granter, entry.Grantee.GetCachedValue().(*types.UserGrantee).User)
	}
	return output
}

// --------------------------------------------------------------------------------------------------------------------

// ValidGroupGrantsInvariant checks that all the user grants are valid
func ValidGroupGrantsInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var invalidEntries []types.Grant
		k.IterateUserGroupsGrants(ctx, func(entry types.Grant) (stop bool) {
			invalid := false

			// Check subspace existence
			if !k.HasSubspace(ctx, entry.SubspaceID) {
				invalid = true
			}

			// Check group existence
			grantee := entry.Grantee.GetCachedValue().(*types.GroupGrantee)
			if !k.HasUserGroup(ctx, entry.SubspaceID, grantee.GroupID) {
				invalid = true
			}

			// Validate the entry
			err := entry.Validate()
			if err != nil {
				invalid = true
			}

			if invalid {
				invalidEntries = append(invalidEntries, entry)
			}

			return false
		})

		return sdk.FormatInvariant(types.ModuleName, "invalid group grants",
			fmt.Sprintf("the following group grants are invalid:\n%s", formatOutputGroupGrants(invalidEntries)),
		), invalidEntries != nil
	}
}

// formatOutputGroupGrants concatenates the given group grants into a string
func formatOutputGroupGrants(entries []types.Grant) (output string) {
	for _, entry := range entries {
		output += fmt.Sprintf("SubspaceID: %d, Granter: %s, GroupID: %d\n", entry.SubspaceID, entry.Granter, entry.Grantee.GetCachedValue().(*types.GroupGrantee).GroupID)
	}
	return output
}
