package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryReportsRequest) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var target ReportTarget
	return unpacker.UnpackAny(r.Target, &target)
}

// NewQueryReportsRequest returns a new QueryReportsRequest instance
func NewQueryReportsRequest(subspaceID uint64, target ReportTarget, pagination *query.PageRequest) *QueryReportsRequest {
	var targetAny *codectypes.Any
	if target != nil {
		any, err := codectypes.NewAnyWithValue(target)
		if err != nil {
			panic("failed to pack target to any type")
		}
		targetAny = any
	}

	return &QueryReportsRequest{
		SubspaceId: subspaceID,
		Target:     targetAny,
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
