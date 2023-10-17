package v620

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	clientkeeper "github.com/cosmos/ibc-go/v7/modules/core/02-client/keeper"

	"github.com/desmos-labs/desmos/v6/app/upgrades"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

// Upgrade represents the v6 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator

	sk           *stakingkeeper.Keeper
	clientKeeper clientkeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, sk *stakingkeeper.Keeper, clientKeeper clientkeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		sk:           sk,
		clientKeeper: clientKeeper,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v6.2.0"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// This upgrade does not require any migration, so we can simply return the current version map
		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
