package v300

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/desmos-labs/desmos/v4/app/upgrades"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

// Upgrade represents the v3.0.0 upgrade
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
	return "v3.0.0"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// Set the initial version of the x/relationships module to 1 so that we can migrate to 2
		fromVM[relationshipstypes.ModuleName] = 1

		// Nothing to do here for the x/subspaces module since the InitGenesis will be called
		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		Added: []string{
			wasm.StoreKey,
			relationshipstypes.StoreKey,
		},

		// The subspaces key is here because it was already registered (due to an error) inside v2.3.1
		// https://github.com/desmos-labs/desmos/blob/v2.3.1/app/app.go#L270
		Renamed: []storetypes.StoreRename{
			{
				OldKey: "subspaces",
				NewKey: subspacestypes.StoreKey,
			},
		},
	}
}
