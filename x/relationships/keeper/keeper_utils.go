package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CheckForBlockedUser checks if the given user address is present inside the blocked users array
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked sdk.AccAddress) bool {
	blockedUsers := k.GetUserBlocks(ctx, blocker)
	for _, user := range blockedUsers {
		if user.Blocked.Equals(blocked) {
			return true
		}
	}
	return false
}
