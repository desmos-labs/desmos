package v400

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"
	reactionstypes "github.com/desmos-labs/desmos/v4/x/reactions/types"
	reportstypes "github.com/desmos-labs/desmos/v4/x/reports/types"

	"github.com/desmos-labs/desmos/v4/app/upgrades"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

// Upgrade represents the v4.0.0 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v4.0.0"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// We set the modules initial versions to 1 because we need to run the migrations
		fromVM[poststypes.ModuleName] = 1
		fromVM[reportstypes.ModuleName] = 1
		fromVM[reactionstypes.ModuleName] = 1

		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{
			poststypes.StoreKey,
			reportstypes.StoreKey,
			reactionstypes.StoreKey,
		},
	}
}
