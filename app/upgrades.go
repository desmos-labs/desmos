package app

import (
	v700 "github.com/desmos-labs/desmos/v7/app/upgrades/v700"
	v710 "github.com/desmos-labs/desmos/v7/app/upgrades/v710"
)

// registerUpgradeHandlers registers all the upgrade handlers that are supported by the app
func (app *DesmosApp) registerUpgradeHandlers() {
	app.registerUpgrade(v700.NewUpgrade(app.ModuleManager, app.Configurator()))
	app.registerUpgrade(v710.NewUpgrade(app.ModuleManager, app.Configurator()))
}
