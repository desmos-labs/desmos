package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/relationships/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// Hooks represents a wrapper struct
type Hooks struct {
	k Keeper
}

var _ subspacestypes.SubspacesHooks = Hooks{}

// Hooks creates new relationships hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

func (h Hooks) AfterSubspaceSaved(sdk.Context, uint64) {}

func (h Hooks) AfterSubspaceDeleted(ctx sdk.Context, subspaceID uint64) {
	var relationships []types.Relationship
	h.k.IterateSubspaceRelationships(ctx, subspaceID, func(index int64, relationship types.Relationship) (stop bool) {
		relationships = append(relationships, relationship)
		return false
	})

	for _, relationship := range relationships {
		h.k.DeleteRelationship(ctx, relationship.Creator, relationship.Counterparty, relationship.SubspaceID)
	}

	var userBlocks []types.UserBlock
	h.k.IterateSubspaceUsersBlocks(ctx, subspaceID, func(index int64, block types.UserBlock) (stop bool) {
		userBlocks = append(userBlocks, block)
		return false
	})

	for _, block := range userBlocks {
		h.k.DeleteUserBlock(ctx, block.Blocker, block.Blocked, block.SubspaceID)
	}
}

func (h Hooks) AfterSubspaceGroupSaved(sdk.Context, uint64, uint32)                         {}
func (h Hooks) AfterSubspaceGroupMemberAdded(sdk.Context, uint64, uint32, sdk.AccAddress)   {}
func (h Hooks) AfterSubspaceGroupMemberRemoved(sdk.Context, uint64, uint32, sdk.AccAddress) {}
func (h Hooks) AfterSubspaceGroupDeleted(sdk.Context, uint64, uint32)                       {}
func (h Hooks) AfterUserPermissionSet(sdk.Context, uint64, sdk.AccAddress, subspacestypes.Permissions) {
}
func (h Hooks) AfterUserPermissionRemoved(sdk.Context, uint64, sdk.AccAddress) {}
