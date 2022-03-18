package types

// DONTCOVER

// NewQueryCirculatingSupplyRequest returns a new QueryCirculatingSupplyRequest instance
func NewQueryCirculatingSupplyRequest(denom string) *QueryCirculatingSupplyRequest {
	return &QueryCirculatingSupplyRequest{Denom: denom}
}

// NewQueryTotalSupplyRequest returns a QueryTotalSupplyRequest instance
func NewQueryTotalSupplyRequest(denom string) *QueryTotalSupplyRequest {
	return &QueryTotalSupplyRequest{Denom: denom}
}
