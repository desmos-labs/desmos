package app

import (
	v620 "github.com/desmos-labs/desmos/v6/app/upgrades/v620"
)

// registerUpgradeHandlers registers all the upgrade handlers that are supported by the app
func (app *DesmosApp) registerUpgradeHandlers() {
	app.registerUpgrade(v620.NewUpgrade(app.ModuleManager, app.Configurator(), app.StakingKeeper, app.IBCKeeper.ClientKeeper))
}
