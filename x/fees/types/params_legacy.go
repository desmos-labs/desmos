package types

import paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

// DONTCOVER

const (
	DefaultParamspace = ModuleName
)

var (
	DefaultMinFees  []MinFee
	MinFeesStoreKey = []byte("MinFees")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(MinFeesStoreKey, &params.MinFees, ValidateMinFeesParam),
	}
}
