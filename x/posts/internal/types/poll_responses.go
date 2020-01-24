package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"strings"
)

// PollUserAnswersQueryResponse represents the answers user made to a post's poll
// that is returned to user upon a query
type PollUserAnswersQueryResponse struct {
	PostID  PostID         `json:"post_id"`
	User    sdk.AccAddress `json:"user"`
	Answers []uint64       `json:"answers"`
}

// String implements fmt.Stringer
func (response PollUserAnswersQueryResponse) String() string {
	out := fmt.Sprintf("Post ID [%s] - User [%s] \n Answers: ", response.PostID, response.User)
	for _, answer := range response.Answers {
		out += fmt.Sprintf("answerID [%s] ", strconv.FormatUint(answer, 10))
	}
	return strings.TrimSpace(out)
}

// PollAnswerAmountResponse represents the amount of answers made of a post's poll
// that is returned to user upon a query
type PollAnswersAmountResponse struct {
	PostID        PostID  `json:"post_id"`
	AnswersAmount sdk.Int `json:"answers_amount"`
}

// String implements fmt.Stringer
func (response PollAnswersAmountResponse) String() string {
	out := fmt.Sprintf("Post ID [%s] - Answers Amount [%s]\n", response.PostID, response.AnswersAmount)
	return strings.TrimSpace(out)
}

type PollAnswerVotesResponse struct {
	PostID      PostID  `json:"post_id"`
	AnswerID    uint64  `json:"answer_id"`
	VotesAmount sdk.Int `json:"votes_amount"`
}

// String implements fmt.Stringer
func (response PollAnswerVotesResponse) String() string {
	out := fmt.Sprintf("Post ID [%s] - Answers ID [%s] Votes Amount [%s]", response.PostID,
		strconv.FormatUint(response.AnswerID, 10), response.VotesAmount)
	return strings.TrimSpace(out)
}
