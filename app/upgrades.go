package app

import (
	v7 "github.com/desmos-labs/desmos/v6/app/upgrades/v7"
)

// registerUpgradeHandlers registers all the upgrade handlers that are supported by the app
func (app *DesmosApp) registerUpgradeHandlers() {
	app.registerUpgrade(v7.NewUpgrade(app.ModuleManager, app.Configurator()))
}
