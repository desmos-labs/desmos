package wasm

import (
	reportsTypes "github.com/desmos-labs/desmos/x/reports/types"
)

type ReportsModuleQuery struct {
	Reports *ReportsQuery `json:"reports"`
}

type ReportsQuery struct {
	PostID string `json:"post_id"`
}

type ReportsResponse struct {
	Reports []reportsTypes.Report `json:"reports"`
}
