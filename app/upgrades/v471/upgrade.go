package v471

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/desmos-labs/desmos/v5/app/upgrades"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

// Upgrade represents the v4.6.0 upgrade
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
	return "v4.7.1"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// Set the coin metadata
		u.bk.SetDenomMetaData(ctx, banktypes.Metadata{
			Description: "The token of Morpheus Apollo",
			DenomUnits: []*banktypes.DenomUnit{
				{
					Denom:    "udaric",
					Exponent: 0,
					Aliases:  nil,
				},
				{
					Denom:    "daric",
					Exponent: 6,
					Aliases:  []string{"Daric"},
				},
			},
			Base:    "udaric",
			Display: "daric",
			Name:    "Daric",
			Symbol:  "",
		})

		// Do nothing here as we don't have anything particular in this update
		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
