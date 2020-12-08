package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

// IterateRelationships iterates through the relationships and perform the provided function
func (k Keeper) IterateRelationships(ctx sdk.Context, fn func(index int64, relationship types.Relationship) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.RelationshipsStorePrefix)
	defer iterator.Close()

	i := int64(0)

	for ; iterator.Valid(); iterator.Next() {
		relationships := types.MustUnmarshalRelationships(k.cdc, iterator.Value())

		var stop = false
		for _, relationship := range relationships {
			stop = fn(i, relationship)

			if stop {
				break
			}

			i++
		}

		if stop {
			break
		}
	}
}
