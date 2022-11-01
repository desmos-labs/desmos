package v4

import (
	"github.com/CosmWasm/wasmd/x/wasm"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	feestypes "github.com/desmos-labs/desmos/v4/x/fees/types"
	reactionstypes "github.com/desmos-labs/desmos/v4/x/reactions/types"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"
	reportstypes "github.com/desmos-labs/desmos/v4/x/reports/types"

	"github.com/desmos-labs/desmos/v4/app/upgrades"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

// Upgrade represents the v4.3.0 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator
	bk           bankkeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, bk bankkeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		bk:           bk,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v4"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// We set the modules initial versions to 1 because we need to run the migrations
		fromVM[relationshipstypes.ModuleName] = 1

		// Set the coin metadata
		u.bk.SetDenomMetaData(ctx, banktypes.Metadata{
			Description: "The token of Desmos",
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    "udsm",
					Exponent: 0,
					Aliases:  nil,
				},
				{
					Denom:    "DSM",
					Exponent: 6,
					Aliases:  nil,
				},
			},
			Base:    "udsm",
			Display: "DSM",
			Name:    "Desmos DSM",
			Symbol:  "DSM",
		})

		// Do nothing here as we don't have anything particular in this update
		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{
		// x/posts and x/subspaces keys where already present (by mistake) inside Desmos v2.3.1.
		// For this reason we don't add them again here as this would result in the following error:
		// failed to load latest version: failed to load store: initial version set to XX, but found earlier version 1
		Added: []string{
			wasm.StoreKey,
			feestypes.StoreKey,
			relationshipstypes.StoreKey,
			reportstypes.StoreKey,
			reactionstypes.StoreKey,
		},
		Deleted: []string{
			"supply", // Remove the supply store key
		},
	}
}
