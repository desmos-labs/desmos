package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryReportsRequest) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var data ReportTarget
	return unpacker.UnpackAny(r.Data, &data)
}

// NewQueryReportsRequest returns a new QueryReportsRequest instance
func NewQueryReportsRequest(subspaceID uint64, data ReportTarget, pagination *query.PageRequest) *QueryReportsRequest {
	var dataAny *codectypes.Any
	if data != nil {
		any, err := codectypes.NewAnyWithValue(data)
		if err != nil {
			panic("failed to pack data to any type")
		}
		dataAny = any
	}

	return &QueryReportsRequest{
		SubspaceId: subspaceID,
		Data:       dataAny,
		Pagination: pagination,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryReportsResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, report := range r.Reports {
		err := report.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewQueryReasonsRequest returns a new QueryReasonsRequest instance
func NewQueryReasonsRequest(subspaceID uint64, pagination *query.PageRequest) *QueryReasonsRequest {
	return &QueryReasonsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQueryParamsRequest returns a new QueryParamsRequest instance
func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
