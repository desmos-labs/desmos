package types

import (
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
)

type (
	MigrationCallback func(appMap genutiltypes.AppMap, args ...interface{}) genutiltypes.AppMap
)
