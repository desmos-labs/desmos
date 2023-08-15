package v530

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/desmos-labs/desmos/v5/app/upgrades"
)

var (
	_ upgrades.Upgrade = &Upgrade{}
)

// Upgrade represents the v5.3.0 upgrade
type Upgrade struct {
	mm           *module.Manager
	configurator module.Configurator

	sk *stakingkeeper.Keeper
}

// NewUpgrade returns a new Upgrade instance
func NewUpgrade(mm *module.Manager, configurator module.Configurator, sk *stakingkeeper.Keeper) *Upgrade {
	return &Upgrade{
		mm:           mm,
		configurator: configurator,
		sk:           sk,
	}
}

// Name implements upgrades.Upgrade
func (u *Upgrade) Name() string {
	return "v5.3.0"
}

// Handler implements upgrades.Upgrade
func (u *Upgrade) Handler() upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// Inside Desmos v5 we added the support for on-chain minimum validator commission.
		// However, even if a proposal was submitted on-chain for this upgrade, the validators have not updated properly.
		// This is due to this issue: https://github.com/cosmos/cosmos-sdk/issues/10540#issuecomment-1155390615.
		//
		// To fix this, we are going to run through all validators and make sure that they have their on-chain
		// commission properly set based on the current on-chain minimum value that might have been set by the
		// governance through a param change proposal.

		minCommission := u.sk.GetParams(ctx).MinCommissionRate
		u.sk.IterateValidators(ctx, func(_ int64, val stakingtypes.ValidatorI) (stop bool) {
			validator, ok := val.(stakingtypes.Validator)
			if !ok {
				return false
			}

			// Make sure the commission rate is at least minCommission.
			// Otherwise, set it to be that minimum
			if validator.Commission.Rate.LT(minCommission) {
				validator.Commission.Rate = minCommission
				u.sk.SetValidator(ctx, validator)
			}

			return false
		})

		// After properly setting all the validator commissions, we can proceed with the normal migration
		return u.mm.RunMigrations(ctx, u.configurator, fromVM)
	}
}

// StoreUpgrades implements upgrades.Upgrade
func (u *Upgrade) StoreUpgrades() *storetypes.StoreUpgrades {
	return &storetypes.StoreUpgrades{}
}
