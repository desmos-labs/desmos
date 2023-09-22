package upgrades

import (
	storetypes "cosmossdk.io/store/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// Upgrade represents a generic on-chain upgrade
type Upgrade interface {
	Name() string
	Handler() upgradetypes.UpgradeHandler
	StoreUpgrades() *storetypes.StoreUpgrades
}
