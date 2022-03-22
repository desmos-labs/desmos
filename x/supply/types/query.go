package types

// DONTCOVER

// NewQueryCirculatingSupplyRequest returns a new QueryCirculatingSupplyRequest instance
func NewQueryCirculatingSupplyRequest(denom string, multiplier int64) *QueryCirculatingSupplyRequest {
	return &QueryCirculatingSupplyRequest{
		Denom:      denom,
		Multiplier: multiplier,
	}
}

// NewQueryTotalSupplyRequest returns a QueryTotalSupplyRequest instance
func NewQueryTotalSupplyRequest(denom string, multiplier int64) *QueryTotalSupplyRequest {
	return &QueryTotalSupplyRequest{
		Denom:      denom,
		Multiplier: multiplier,
	}
}
