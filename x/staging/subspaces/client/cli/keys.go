package cli

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// DONTCOVER

// Subspaces flag
const (
	flagLimit = "limit"
	flagPage  = "page"

	FlagSubspaceType = "type"
	FlagName         = "name"
	FlagOwner        = "owner"

	DoNotEdit = "do-not-edit"
)

func DefaultQuerySubspacesRequest(page, limit uint64) types.QuerySubspacesRequest {
	return types.QuerySubspacesRequest{
		Pagination: &query.PageRequest{
			Key:        nil,
			Offset:     (page - 1) * limit,
			Limit:      limit,
			CountTotal: false,
		},
	}
}
