package types

// DONTCOVER

// NewQueryCirculatingSupplyRequest returns a new QueryCirculatingSupplyRequest instance
func NewQueryCirculatingSupplyRequest(denom string, divider int64) *QueryCirculatingSupplyRequest {
	return &QueryCirculatingSupplyRequest{
		Denom:   denom,
		Divider: divider,
	}
}

// NewQueryTotalSupplyRequest returns a QueryTotalSupplyRequest instance
func NewQueryTotalSupplyRequest(denom string, divider int64) *QueryTotalSupplyRequest {
	return &QueryTotalSupplyRequest{
		Denom:   denom,
		Divider: divider,
	}
}
