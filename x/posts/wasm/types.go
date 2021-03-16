package wasm

import (
	"encoding/json"
	"strconv"

	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
)

type Post struct {
	PostID         string                  `json:"post_id"`
	ParentID       string                  `json:"parent_id"`
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
		PostID:         post.PostId,
		ParentID:       post.ParentId,
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

type PostsModuleQuery struct {
	Posts *PostsQuery `json:"posts"`
}

type PostsQuery struct{}

type Posts []Post

func (p Posts) MarshalJSON() ([]byte, error) {
	if len(p) == 0 {
		return []byte("[]"), nil
	}
	var posts []Post = p
	return json.Marshal(posts)
}

func (p *Posts) UnmarshalJSON(data []byte) error {
	// make sure we deserialize [] back to null
	if string(data) == "[]" || string(data) == "null" {
		return nil
	}
	var posts []Post
	if err := json.Unmarshal(data, &posts); err != nil {
		return err
	}
	*p = posts
	return nil
}

type PostsResponse struct {
	Posts Posts `json:"posts"`
}
