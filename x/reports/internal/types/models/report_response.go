package models

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/posts"
	"strings"
)

type ReportsQueryResponse struct {
	PostID  posts.PostID `json:"post_id" yaml:"post_id"`
	Reports `json:"reports" yaml:"reports"`
}

func NewReportResponse(postID posts.PostID, reports Reports) ReportsQueryResponse {
	return ReportsQueryResponse{
		PostID:  postID,
		Reports: reports,
	}
}

// String implements fmt.Stringer
func (response ReportsQueryResponse) String() string {
	out := fmt.Sprintf("Post ID: %s\n Reports: %s\n", response.PostID, response.Reports)
	return strings.TrimSpace(out)
}
