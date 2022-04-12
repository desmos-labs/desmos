package types

import (
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DONTCOVER

// NewQueryCirculatingSupplyRequest returns a new QueryCirculatingSupplyRequest instance
func NewQueryCirculatingSupplyRequest(denom string, dividerExponent uint64) *QueryCirculatingSupplyRequest {
	return &QueryCirculatingSupplyRequest{
		Denom:           denom,
		DividerExponent: dividerExponent,
	}
}

// NewQueryTotalSupplyRequest returns a QueryTotalSupplyRequest instance
func NewQueryTotalSupplyRequest(denom string, dividerExponent uint64) *QueryTotalSupplyRequest {
	return &QueryTotalSupplyRequest{
		Denom:           denom,
		DividerExponent: dividerExponent,
	}
}

// NewDividerPoweredByExponent takes the given exponent using it to power 10 to calculate the correct
// divider
func NewDividerPoweredByExponent(exponent uint64) sdk.Int {
	return sdk.NewInt(int64(math.Pow10(int(exponent))))
}
