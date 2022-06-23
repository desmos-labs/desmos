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
func NewQueryReportsRequest(subspaceID uint64, target ReportTarget, reporter string, pagination *query.PageRequest) *QueryReportsRequest {
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
		Reporter:   reporter,
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

// NewQueryReportRequest returns a new QueryReportRequest instance
func NewQueryReportRequest(subspaceID uint64, reportID uint64) *QueryReportRequest {
	return &QueryReportRequest{
		SubspaceId: subspaceID,
		ReportId:   reportID,
	}
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (r *QueryReportResponse) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	return r.Report.UnpackInterfaces(unpacker)
}

// NewQueryReasonsRequest returns a new QueryReasonsRequest instance
func NewQueryReasonsRequest(subspaceID uint64, pagination *query.PageRequest) *QueryReasonsRequest {
	return &QueryReasonsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQueryReasonRequest returns a new QueryReasonRequest instance
func NewQueryReasonRequest(subspaceID uint64, reasonID uint32) *QueryReasonRequest {
	return &QueryReasonRequest{
		SubspaceId: subspaceID,
		ReasonId:   reasonID,
	}
}

// NewQueryParamsRequest returns a new QueryParamsRequest instance
func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
