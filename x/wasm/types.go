package wasm

import (
	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
	"strconv"
)

type Post struct {
	PostId         string                  `json:"post_id"`
	ParentId       string                  `json:"parent_id"`
	Message        string                  `json:"message"`
	Created        string                  `json:"created"`
	LastEdited     string                  `json:"last_edited"`
	AllowsComments string                  `json:"allows_comments"`
	Subspace       string                  `json:"subspace"`
	OptionalData   postsTypes.OptionalData `json:"optional_data"`
	Creator        string                  `json:"creator"`
	Attachments    postsTypes.Attachments  `json:"attachments"`
	PollData       PollData                `json:"poll_data"`
}

type PollData struct {
	Question              string                  `json:"question"`
	ProvidedAnswers       []postsTypes.PollAnswer `json:"provided_answers"`
	EndDate               string                  `json:"end_date"`
	AllowsMultipleAnswers string                  `json:"allows_multiple_answers"`
	AllowsAnswerEdits     string                  `json:"allows_answer_edits"`
}

// convertPost convert the desmos post to a compatible type for cosmwasm contract
func convertPost(post postsTypes.Post) Post {
	return Post{
		PostId:         post.PostId,
		ParentId:       post.ParentId,
		Message:        post.Message,
		Created:        post.Created.String(),
		LastEdited:     post.LastEdited.String(),
		AllowsComments: strconv.FormatBool(post.AllowsComments),
		Subspace:       post.Subspace,
		OptionalData:   post.OptionalData,
		Creator:        post.Creator,
		Attachments:    post.Attachments,
		PollData: PollData{
			Question:              post.PollData.Question,
			ProvidedAnswers:       post.PollData.ProvidedAnswers,
			EndDate:               post.PollData.EndDate.String(),
			AllowsMultipleAnswers: strconv.FormatBool(post.PollData.AllowsMultipleAnswers),
			AllowsAnswerEdits:     strconv.FormatBool(post.PollData.AllowsAnswerEdits),
		},
	}
}
