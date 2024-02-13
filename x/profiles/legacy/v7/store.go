package v7

import (
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/profiles/types"
)

// MigrateStore performs in-place store migrations from v7 to v8.
// The migration consists of params migration.
func MigrateStore(ctx sdk.Context, paramsSubspace paramstypes.Subspace) error {
	// Set the missing app links params
	paramsSubspace.Set(ctx, types.AppLinksParamsKey, types.DefaultAppLinksParams())

	return nil
}
