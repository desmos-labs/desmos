package v500

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	consensusparamskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/desmos-labs/desmos/v5/app/upgrades"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

// Upgrade represents the v5.2.0 upgrade
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
	return "v5.2.0"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
