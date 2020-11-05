package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// PostQueryResponse represents the data of a post
// that is returned to user upon a query
type PostQueryResponse struct {
	Post
	PollAnswers []UserAnswer   `json:"poll_answers,omitempty" yaml:"poll_answers,omitempty"`
	Reactions   []PostReaction `json:"reactions" yaml:"reactions,omitempty"`
	Children    []string       `json:"children" yaml:"children"`
}

// String implements fmt.Stringer
func (response PostQueryResponse) String() string {
	out := fmt.Sprintf(`
ID: %s
Reactions: %s
Children: %s
`, response.Post.PostID, response.Reactions, response.Children)
	return strings.TrimSpace(out)
}

func NewPostResponse(
	post Post, pollAnswers []UserAnswer, reactions []PostReaction, children []string,
) PostQueryResponse {
	return PostQueryResponse{
		Post:        post,
		PollAnswers: pollAnswers,
		Reactions:   reactions,
		Children:    children,
	}
}

// MarshalJSON implements json.Marshaler as Amino does
// not respect default json composition
func (response PostQueryResponse) MarshalJSON() ([]byte, error) {
	type temp PostQueryResponse
	return json.Marshal(temp(response))
}

// UnmarshalJSON implements json.Unmarshaler as Amino does
// not respect default json composition
func (response *PostQueryResponse) UnmarshalJSON(data []byte) error {
	type postResponse PostQueryResponse
	var temp postResponse
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*response = PostQueryResponse(temp)
	return nil
}
