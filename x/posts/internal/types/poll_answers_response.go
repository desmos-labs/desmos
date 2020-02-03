package types

import (
	"fmt"
	"strings"
)

// PollAnswersQueryResponse represents the answers made to a post's poll
// that is returned to user upon a query
type PollAnswersQueryResponse struct {
	PostID         PostID              `json:"post_id"`
	AnswersDetails UsersAnswersDetails `json:"answers_details"`
}

// String implements fmt.Stringer
func (response PollAnswersQueryResponse) String() string {
	out := fmt.Sprintf("Post ID [%s] - Answers Details:\n", response.PostID)
	for _, answerDetails := range response.AnswersDetails {
		out += fmt.Sprintf("%s\n", answerDetails.String())
	}
	return strings.TrimSpace(out)
}
