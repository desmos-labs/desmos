package app

import (
	v500 "github.com/desmos-labs/desmos/v5/app/upgrades/v500"
)

// registerUpgradeHandlers registers all the upgrade handlers that are supported by the app
func (app *DesmosApp) registerUpgradeHandlers() {
	app.registerUpgrade(v500.NewUpgrade(app.ModuleManager, app.Configurator(), app.ParamsKeeper, app.ConsensusParamsKeeper))
}
