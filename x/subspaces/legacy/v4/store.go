package v4

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	subspacesauthz "github.com/desmos-labs/desmos/v5/x/subspaces/authz"
	v2 "github.com/desmos-labs/desmos/v5/x/subspaces/legacy/v2"
)

// MigrateStore migrates the store from version 3 to version 4.
// The process performs the missing authz migration that was excluded from version 2 to version 3.
func MigrateStore(ctx sdk.Context, authzKeeper authzkeeper.Keeper, cdc codec.BinaryCodec) error {
	authzKeeper.IterateGrants(ctx, func(granterAddr sdk.AccAddress, granteeAddr sdk.AccAddress, grant authz.Grant) bool {
		if auth, ok := grant.Authorization.GetCachedValue().(*v2.GenericSubspaceAuthorization); ok {
			// Convert the generic grant authorization
			v3Auth := subspacesauthz.NewGenericSubspaceAuthorization(auth.SubspacesIDs, auth.Msg)

			// Store the new authorization
			err := authzKeeper.SaveGrant(ctx, granteeAddr, granterAddr, v3Auth, grant.Expiration)
			if err != nil {
				panic(err)
			}
		}

		return false
	})

	return nil
}
