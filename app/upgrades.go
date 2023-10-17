package app

import (
	v500 "github.com/desmos-labs/desmos/v6/app/upgrades/v500"
	v520 "github.com/desmos-labs/desmos/v6/app/upgrades/v520"
	v6 "github.com/desmos-labs/desmos/v6/app/upgrades/v6"
	v620 "github.com/desmos-labs/desmos/v6/app/upgrades/v620"
)

// registerUpgradeHandlers registers all the upgrade handlers that are supported by the app
func (app *DesmosApp) registerUpgradeHandlers() {
	app.registerUpgrade(v500.NewUpgrade(app.ModuleManager, app.Configurator(), app.ParamsKeeper, app.ConsensusParamsKeeper))
	app.registerUpgrade(v520.NewUpgrade(app.ModuleManager, app.Configurator(), app.ParamsKeeper, app.ConsensusParamsKeeper))
	app.registerUpgrade(v6.NewUpgrade(app.ModuleManager, app.Configurator(), app.StakingKeeper, app.IBCKeeper.ClientKeeper))
	app.registerUpgrade(v620.NewUpgrade(app.ModuleManager, app.Configurator(), app.StakingKeeper, app.IBCKeeper.ClientKeeper))
}
