package types

import (
	"github.com/cosmos/cosmos-sdk/x/genutil"
)

type (
	MigrationCallback func(appMap genutil.AppMap, args ...interface{}) genutil.AppMap
)
