package models

import (
	"encoding/json"
	"fmt"
	"strings"

	posts "github.com/desmos-labs/desmos/x/posts/types"
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
	out := fmt.Sprintf("Post RelationshipID: %s\n Reports: %s\n", response.PostID, response.Reports)
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
	type postResponse ReportsQueryResponse
	var temp postResponse
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*response = ReportsQueryResponse(temp)
	return nil
}
