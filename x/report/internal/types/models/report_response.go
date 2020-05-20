package models

import (
	"encoding/json"
	"fmt"
	"github.com/desmos-labs/desmos/x/posts"
	"strings"
)

type ReportsQueryResponse struct {
	PostID posts.PostID `json:"post_id" yaml:"post_id"`
	Reports
}

func NewReportResponse(postID posts.PostID, reports Reports) ReportsQueryResponse {
	return ReportsQueryResponse{
		PostID:  postID,
		Reports: reports,
	}
}

// String implements fmt.Stringer
func (response ReportsQueryResponse) String() string {
	out := fmt.Sprintf("Post ID: %s\n Reports: %s", response.PostID, response.Reports)
	return strings.TrimSpace(out)
}

// MarshalJSON implements json.Marshaler as Amino does
// not respect default json composition
func (response ReportsQueryResponse) MarshalJSON() ([]byte, error) {
	type temp ReportsQueryResponse
	return json.Marshal(temp(response))
}

// UnmarshalJSON implements json.Unmarshaler as Amino does
// not respect default json composition
func (response *ReportsQueryResponse) UnmarshalJSON(data []byte) error {
	type reportResponse ReportsQueryResponse
	var temp reportResponse
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*response = ReportsQueryResponse(temp)
	return nil
}
