package app

import (
	v620 "github.com/desmos-labs/desmos/v7/app/upgrades/v620"
	v630 "github.com/desmos-labs/desmos/v7/app/upgrades/v630"
	v640 "github.com/desmos-labs/desmos/v7/app/upgrades/v640"
)

// registerUpgradeHandlers registers all the upgrade handlers that are supported by the app
func (app *DesmosApp) registerUpgradeHandlers() {
	app.registerUpgrade(v620.NewUpgrade(app.ModuleManager, app.Configurator()))
	app.registerUpgrade(v630.NewUpgrade(app.ModuleManager, app.Configurator()))
	app.registerUpgrade(v640.NewUpgrade(app.ModuleManager, app.Configurator()))
}
