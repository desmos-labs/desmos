package types

// DONTCOVER

// NewQueryCirculatingSupplyRequest returns a new QueryCirculatingSupplyRequest instance
func NewQueryCirculatingSupplyRequest(denom string) *QueryCirculatingSupplyRequest {
	return &QueryCirculatingSupplyRequest{Denom: denom}
}
