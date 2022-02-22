package v2

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MigrateStore performs in-place store migrations from v4 to v5
// The migration includes:
//
// - delete all the relationships
// - delete all the user blocks
//
// NOTE: This method must be called AFTER the migration from v0 to v1 of the x/relationships module.
// 		 If this order is not preserved, all relationships and blocks WILL BE DELETED.
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	legacyKeeper := NewKeeper(storeKey, cdc)
	legacyKeeper.DeleteRelationships(ctx)
	legacyKeeper.DeleteBlocks(ctx)
	return nil
}
