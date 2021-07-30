package wasm

import (
	"time"

	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
)

type PostsModuleQueryRoutes struct {
	Posts     *PostsQuery     `json:"posts"`
	Reports   *ReportsQuery   `json:"reports"`
	Reactions *ReactionsQuery `json:"reactions"`
}

//////////////Posts///////////////////

type Post struct {
	PostID               string                 `json:"post_id"`
	ParentID             string                 `json:"parent_id"`
	Message              string                 `json:"message"`
	Created              string                 `json:"created"`
	LastEdited           string                 `json:"last_edited"`
	CommentsState        string                 `json:"comments_state"`
	Subspace             string                 `json:"subspace"`
	AdditionalAttributes []postsTypes.Attribute `json:"additional_attributes"`
	Creator              string                 `json:"creator"`
	Attachments          postsTypes.Attachments `json:"attachments"`
	Poll                 *Poll                  `json:"poll_data"`
}

type Poll struct {
	Question              string                      `json:"question"`
	ProvidedAnswers       []postsTypes.ProvidedAnswer `json:"provided_answers"`
	EndDate               string                      `json:"end_date"`
	AllowsMultipleAnswers bool                        `json:"allows_multiple_answers"`
	AllowsAnswerEdits     bool                        `json:"allows_answer_edits"`
}

// convertPosts convert the posts array to make it easier for wasm contracts to deserialize them
func convertPosts(posts []postsTypes.Post) []Post {
	convertedPosts := make([]Post, len(posts))
	for index, post := range posts {
		convertedPosts[index] = convertPost(post)
	}
	return convertedPosts
}

// convertPost convert the desmos post to make it easier for wasm contracts to deserialize it
func convertPost(post postsTypes.Post) Post {
	converted := Post{
		PostID:               post.PostID,
		ParentID:             post.ParentID,
		Message:              post.Message,
		Created:              post.Created.Format(time.RFC3339),
		LastEdited:           post.LastEdited.Format(time.RFC3339),
		CommentsState:        post.CommentsState.String(),
		Subspace:             post.Subspace,
		AdditionalAttributes: []postsTypes.Attribute{},
		Creator:              post.Creator,
		Attachments:          postsTypes.Attachments{},
	}

	if post.AdditionalAttributes != nil {
		converted.AdditionalAttributes = post.AdditionalAttributes
	}

	if post.Attachments != nil {
		converted.Attachments = post.Attachments
	}

	if post.Poll != nil {
		converted.Poll = &Poll{
			Question:              post.Poll.Question,
			ProvidedAnswers:       post.Poll.ProvidedAnswers,
			EndDate:               post.Poll.EndDate.Format(time.RFC3339),
			AllowsMultipleAnswers: post.Poll.AllowsMultipleAnswers,
			AllowsAnswerEdits:     post.Poll.AllowsAnswerEdits,
		}
	}
	return converted
}

type PostsQuery struct{}

type PostsResponse struct {
	Posts []Post `json:"posts"`
}

//////////////Reactions///////////////////

type ReactionsQuery struct {
	PostID string `json:"post_id"`
}

type ReactionsResponse struct {
	Reactions []postsTypes.PostReaction `json:"reactions"`
}

//////////////Reports///////////////////

type ReportsQuery struct {
	PostID string `json:"post_id"`
}

type Report struct {
	PostID  string `json:"post_id"`
	Kind    string `json:"kind"`
	Message string `json:"message"`
	User    string `json:"user"`
}

func convertReports(reports []postsTypes.Report) []Report {
	convertedReports := make([]Report, len(reports))
	for index, report := range reports {
		convertedReports[index] = Report{
			PostID:  report.PostID,
			Kind:    report.Type,
			Message: report.Message,
			User:    report.User,
		}
	}
	return convertedReports
}

type ReportsResponse struct {
	Reports []Report `json:"reports"`
}
