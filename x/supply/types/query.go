package types

import (
	math "math"

	cosmosmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewQueryTotalRequest returns a QueryTotalRequest instance
func NewQueryTotalRequest(denom string, dividerExponent uint64) *QueryTotalRequest {
	return &QueryTotalRequest{
		Denom:           denom,
		DividerExponent: dividerExponent,
	}
}

// NewQueryCirculatingRequest returns a new QueryCirculatingRequest instance
func NewQueryCirculatingRequest(denom string, dividerExponent uint64) *QueryCirculatingRequest {
	return &QueryCirculatingRequest{
		Denom:           denom,
		DividerExponent: dividerExponent,
	}
}

// NewDividerPoweredByExponent takes the given exponent using it to power 10 to calculate the correct
// divider
func NewDividerPoweredByExponent(exponent uint64) cosmosmath.Int {
	return sdk.NewInt(int64(math.Pow10(int(exponent))))
}
