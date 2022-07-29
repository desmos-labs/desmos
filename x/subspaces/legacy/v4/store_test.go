package v4_test

import (
	"testing"
	"time"

	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	subspacesauthz "github.com/desmos-labs/desmos/v4/x/subspaces/authz"
	v4 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v4"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/testutil/storetesting"
	v2 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v2"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(authzkeeper.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// Build the authz keeper
	authzKeeper := authzkeeper.NewKeeper(keys[authzkeeper.StoreKey], cdc, nil)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "generic subspace authorization is migrated properly",
			store: func(ctx sdk.Context) {
				authorization := v2.NewGenericSubspaceAuthorization([]uint64{1, 2}, sdk.MsgTypeURL(&types.MsgEditSubspace{}))
				expiration := time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC)

				granterAddr, err := sdk.AccAddressFromBech32("cosmos1d6j0q7akh0sg69qz7cz6y0e8kc6ljxz8p0ah2t")
				require.NoError(t, err)

				granteeAddr, err := sdk.AccAddressFromBech32("cosmos1w6hk7984dpgtgxwqwuv5645yeq02c4svknv5ua")
				require.NoError(t, err)

				err = authzKeeper.SaveGrant(ctx, granteeAddr, granterAddr, authorization, expiration)
				require.NoError(t, err)
			},
			check: func(ctx sdk.Context) {
				granterAddr, err := sdk.AccAddressFromBech32("cosmos1d6j0q7akh0sg69qz7cz6y0e8kc6ljxz8p0ah2t")
				require.NoError(t, err)

				granteeAddr, err := sdk.AccAddressFromBech32("cosmos1w6hk7984dpgtgxwqwuv5645yeq02c4svknv5ua")
				require.NoError(t, err)

				authorizations := authzKeeper.GetAuthorizations(ctx, granteeAddr, granterAddr)
				require.Len(t, authorizations, 1)

				require.Equal(t, subspacesauthz.NewGenericSubspaceAuthorization(
					[]uint64{1, 2},
					sdk.MsgTypeURL(&types.MsgEditSubspace{}),
				), authorizations[0])
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, tKeys, memKeys)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v4.MigrateStore(ctx, keys[authzkeeper.StoreKey], cdc)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
