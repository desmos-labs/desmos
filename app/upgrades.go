package app

import (
	v4 "github.com/desmos-labs/desmos/v4/app/upgrades/v4"
	v471 "github.com/desmos-labs/desmos/v4/app/upgrades/v471"
	v480 "github.com/desmos-labs/desmos/v4/app/upgrades/v480"
	v500 "github.com/desmos-labs/desmos/v4/app/upgrades/v500"
	v510 "github.com/desmos-labs/desmos/v4/app/upgrades/v510"
)

// registerUpgradeHandlers registers all the upgrade handlers that are supported by the app
func (app *DesmosApp) registerUpgradeHandlers() {
	app.registerUpgrade(v500.NewUpgrade(app.ModuleManager, app.configurator, app.ParamsKeeper, app.ConsensusParamsKeeper))
	app.registerUpgrade(v510.NewUpgrade(app.ModuleManager, app.configurator))
}
