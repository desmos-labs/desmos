package v4

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	subspacesauthz "github.com/desmos-labs/desmos/v4/x/subspaces/authz"
	v2 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v2"
)

// MigrateStore migrates the store from version 3 to version 4.
// The process performs the missing authz migration that was excluded from version 2 to version 3.
func MigrateStore(ctx sdk.Context, authzStoreKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	authzStore := ctx.KVStore(authzStoreKey)
	iterator := sdk.KVStorePrefixIterator(authzStore, authzkeeper.GrantKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var grant authz.Grant
		err := cdc.Unmarshal(iterator.Value(), &grant)
		if err != nil {
			return err
		}

		if auth, ok := grant.Authorization.GetCachedValue().(*v2.GenericSubspaceAuthorization); ok {
			// Convert the generic grant authorization
			v3Auth := subspacesauthz.NewGenericSubspaceAuthorization(auth.SubspacesIDs, auth.Msg)

			// Update the grant authorization value
			v3AuthAny, err := codectypes.NewAnyWithValue(v3Auth)
			if err != nil {
				return err
			}
			grant.Authorization = v3AuthAny

			// Store the new authorization
			authzStore.Set(iterator.Key(), cdc.MustMarshal(&grant))
		}
	}

	return nil
}
