package wasm

import (
	"time"

	postsTypes "github.com/desmos-labs/desmos/x/staging/posts/types"
)

type Post struct {
	PostID         string                  `json:"post_id"`
	ParentID       string                  `json:"parent_id"`
	Message        string                  `json:"message"`
	Created        string                  `json:"created"`
	LastEdited     string                  `json:"last_edited"`
	AllowsComments bool                    `json:"allows_comments"`
	Subspace       string                  `json:"subspace"`
	OptionalData   postsTypes.OptionalData `json:"optional_data"`
	Creator        string                  `json:"creator"`
	Attachments    postsTypes.Attachments  `json:"attachments"`
	PollData       *PollData               `json:"poll_data"`
}

type PollData struct {
	Question              string                  `json:"question"`
	ProvidedAnswers       []postsTypes.PollAnswer `json:"provided_answers"`
	EndDate               string                  `json:"end_date"`
	AllowsMultipleAnswers bool                    `json:"allows_multiple_answers"`
	AllowsAnswerEdits     bool                    `json:"allows_answer_edits"`
}

// convertPosts convert the posts array into a decodable one for cosmwasm contract
func convertPosts(posts []postsTypes.Post) []Post {
	convertedPosts := make([]Post, len(posts))
	for index, post := range posts {
		convertedPosts[index] = convertPost(post)
	}
	return convertedPosts
}

// convertPost convert the desmos post to a decodable type for cosmwasm contract
func convertPost(post postsTypes.Post) Post {
	converted := Post{
		PostID:         post.PostId,
		ParentID:       post.ParentId,
		Message:        post.Message,
		Created:        post.Created.Format(time.RFC3339),
		LastEdited:     post.LastEdited.Format(time.RFC3339),
		AllowsComments: post.AllowsComments,
		Subspace:       post.Subspace,
		OptionalData:   postsTypes.OptionalData{},
		Creator:        post.Creator,
		Attachments:    postsTypes.Attachments{},
	}

	if post.OptionalData != nil {
		converted.OptionalData = post.OptionalData
	}

	if post.Attachments != nil {
		converted.Attachments = post.Attachments
	}

	if post.PollData != nil {
		converted.PollData = &PollData{
			Question:              post.PollData.Question,
			ProvidedAnswers:       post.PollData.ProvidedAnswers,
			EndDate:               post.PollData.EndDate.Format(time.RFC3339),
			AllowsMultipleAnswers: post.PollData.AllowsMultipleAnswers,
			AllowsAnswerEdits:     post.PollData.AllowsAnswerEdits,
		}
	}
	return converted
}

type PostsModuleQuery struct {
	Posts *PostsQuery `json:"posts"`
}

type PostsQuery struct{}

type PostsResponse struct {
	Posts []Post `json:"posts"`
}
