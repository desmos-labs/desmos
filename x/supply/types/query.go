package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// DONTCOVER

// NewQueryCirculatingSupplyRequest returns a new QueryCirculatingSupplyRequest instance
func NewQueryCirculatingSupplyRequest(denom string, divider uint64) *QueryCirculatingSupplyRequest {
	return &QueryCirculatingSupplyRequest{
		Denom:   denom,
		Divider: divider,
	}
}

// NewQueryTotalSupplyRequest returns a QueryTotalSupplyRequest instance
func NewQueryTotalSupplyRequest(denom string, divider uint64) *QueryTotalSupplyRequest {
	return &QueryTotalSupplyRequest{
		Denom:   denom,
		Divider: divider,
	}
}

// NewDividerFromRawInt returns the divider wrapped in sdk.Int making sure it's never equal to 0
func NewDividerFromRawInt(dividerRaw uint64) sdk.Int {
	switch dividerRaw {
	case 0:
		return sdk.NewInt(1)
	default:
		return sdk.NewIntFromUint64(dividerRaw)
	}
}
