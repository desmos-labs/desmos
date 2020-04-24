package types

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

// PostQueryResponse represents the data of a post
// that is returned to user upon a query
type PostQueryResponse struct {
	Post
	PollAnswers []UserAnswer          `json:"poll_answers,omitempty"`
	Reactions   PostReactionsResponse `json:"reactions"`
	Children    PostIDs               `json:"children"`
}

type PostReactionsResponse struct {
	Value string         `json:"value"`
	Code  string         `json:"code"`
	User  sdk.AccAddress `json:"user"`
}

type PostReactionsResponses []PostReactionsResponse

// String implements fmt.Stringer
func (reactionResponse PostReactionsResponse) String() string {
	out := "[Value] [Code] [Creator]\n"
	out += fmt.Sprintf("[%s] [%s] [%s]\n", reactionResponse.Value, reactionResponse.Code, reactionResponse.User)
	return strings.TrimSpace(out)
}

// String implements fmt.Stringer
func (response PostQueryResponse) String() string {
	out := "ID - [PostReactions] [Children] \n"
	out += fmt.Sprintf("%s - [%s] [%s] \n", response.Post.PostID, response.Reactions, response.Children)
	return strings.TrimSpace(out)
}

func NewPostReactionsResponse(reaction PostReaction) PostReactionsResponse {
	return PostReactionsResponse{
		Value: reaction.Value,
		Code:  reaction.Owner,
	}
}

func NewPostResponse(post Post, pollAnswers []UserAnswer, reactions PostReactions, children PostIDs) PostQueryResponse {
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
