package v500

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	consensusparamskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	consensustypes "github.com/cosmos/cosmos-sdk/x/consensus/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"

	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	feestypes "github.com/desmos-labs/desmos/v5/x/fees/types"
	poststypes "github.com/desmos-labs/desmos/v5/x/posts/types"
	profilestypes "github.com/desmos-labs/desmos/v5/x/profiles/types"
	reportstypes "github.com/desmos-labs/desmos/v5/x/reports/types"

	"github.com/desmos-labs/desmos/v5/app/upgrades"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

// Upgrade represents the v5.0.0 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator

	paramsKeeper          paramskeeper.Keeper
	consensusParamsKeeper consensusparamskeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, pk paramskeeper.Keeper, consensusParamsKeeper consensusparamskeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:                    mm,
		configurator:          configurator,
		paramsKeeper:          pk,
		consensusParamsKeeper: consensusParamsKeeper,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v5.0.0"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {

		// Set param key table for params module migration
		for _, subspace := range u.paramsKeeper.GetSubspaces() {
			subspace := subspace

			var keyTable paramstypes.KeyTable
			switch subspace.Name() {
			case authtypes.ModuleName:
				keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
			case banktypes.ModuleName:
				keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
			case stakingtypes.ModuleName:
				keyTable = stakingtypes.ParamKeyTable()
			case minttypes.ModuleName:
				keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
			case distrtypes.ModuleName:
				keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
			case slashingtypes.ModuleName:
				keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
			case govtypes.ModuleName:
				keyTable = govv1.ParamKeyTable() //nolint:staticcheck
			case crisistypes.ModuleName:
				keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck

				// ibc
			case ibctransfertypes.ModuleName:
				keyTable = ibctransfertypes.ParamKeyTable()
			case icahosttypes.SubModuleName:
				keyTable = icahosttypes.ParamKeyTable()
			case icacontrollertypes.SubModuleName:
				keyTable = icacontrollertypes.ParamKeyTable()

				// custom
			case feestypes.ModuleName:
				keyTable = feestypes.ParamKeyTable()
			case poststypes.ModuleName:
				keyTable = poststypes.ParamKeyTable()
			case profilestypes.ModuleName:
				keyTable = profilestypes.ParamKeyTable()
			case reportstypes.ModuleName:
				keyTable = reportstypes.ParamKeyTable()

				// wasm
			case wasmtypes.ModuleName:
				keyTable = wasmtypes.ParamKeyTable() //nolint:staticcheck

				// Skip if module is not migration target
			default:
				continue
			}

			if !subspace.HasKeyTable() {
				subspace.WithKeyTable(keyTable)
			}
		}

		// Migrate Tendermint consensus parameters from x/params module to a dedicated x/consensus module.
		baseAppLegacySS := u.paramsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		baseapp.MigrateParams(ctx, baseAppLegacySS, &u.consensusParamsKeeper)

		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{
			// New IBC features integration upgrades
			ibcfeetypes.StoreKey,
			icacontrollertypes.StoreKey,
			icahosttypes.StoreKey,

			// Cosmos SDK v0.47.x upgrades
			consensustypes.ModuleName,
			crisistypes.ModuleName,
		},
	}
}
