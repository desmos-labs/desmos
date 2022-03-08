package v2_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/store"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/app"
	v2 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v2"
	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

func buildContext(
	keys map[string]*sdk.KVStoreKey, tKeys map[string]*sdk.TransientStoreKey, memKeys map[string]*sdk.MemoryStoreKey,
) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	for _, key := range keys {
		cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	}
	for _, tKey := range tKeys {
		cms.MountStoreWithDB(tKey, sdk.StoreTypeTransient, db)
	}
	for _, memKey := range memKeys {
		cms.MountStoreWithDB(memKey, sdk.StoreTypeMemory, nil)
	}

	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
}

func TestMigrateStore(t *testing.T) {
	cdc, legacyAmino := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(authtypes.StoreKey, types.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// Build the x/auth keeper
	paramsKeeper := paramskeeper.NewKeeper(
		cdc,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tKeys[paramstypes.TStoreKey],
	)
	authKeeper := authkeeper.NewAccountKeeper(
		cdc,
		keys[authtypes.StoreKey],
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "valid data is migrated properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				blockBz := cdc.MustMarshal(&v2.UserBlock{
					Blocker:    "blocker",
					Blocked:    "blocked",
					Reason:     "reason",
					SubspaceID: "",
				})
				store.Set(v2.UserBlockStoreKey("blocker", "", "blocked"), blockBz)

				relBz := cdc.MustMarshal(&v2.Relationship{
					Creator:    "user",
					Recipient:  "recipient",
					SubspaceID: "1",
				})
				store.Set(v2.RelationshipsStoreKey("user", "1", "recipient"), relBz)

				relBz = cdc.MustMarshal(&v2.Relationship{
					Creator:    "user",
					Recipient:  "recipient",
					SubspaceID: "2",
				})
				store.Set(v2.RelationshipsStoreKey("user", "2", "recipient"), relBz)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				// Make sure all blocks are deleted
				require.False(t, store.Has(v2.UserBlockStoreKey("blocker", "", "blocked")))

				// Make sure all relationships are deleted
				require.False(t, store.Has(v2.RelationshipsStoreKey("user", "1", "recipient")))
				require.False(t, store.Has(v2.RelationshipsStoreKey("user", "1", "recipient")))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := buildContext(keys, tKeys, memKeys)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v2.MigrateStore(ctx, authKeeper, keys[types.StoreKey], legacyAmino, cdc)
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
