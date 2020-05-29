package models

import (
	"github.com/desmos-labs/desmos/x/posts"
)

// ReportsStoreKey turn an id to a key used to store a reports inside the reports store
func ReportStoreKey(id posts.PostID) []byte {
	return append(ReportsStorePrefix, []byte(id)...)
}
