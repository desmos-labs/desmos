package storetesting

import (
	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func BuildContext(
	keys map[string]*storetypes.KVStoreKey, tKeys map[string]*storetypes.TransientStoreKey, memKeys map[string]*storetypes.MemoryStoreKey,
) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range keys {
		cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	}
	for _, tKey := range tKeys {
		cms.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, db)
	}
	for _, memKey := range memKeys {
		cms.MountStoreWithDB(memKey, storetypes.StoreTypeMemory, nil)
	}

	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
}
