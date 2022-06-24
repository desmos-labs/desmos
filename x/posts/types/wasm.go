package types

import (
	"encoding/json"
)

type PostsMsg struct {
	CreatePost           *json.RawMessage `json:"create_post"`
	EditPost             *json.RawMessage `json:"edit_post"`
	DeletePost           *json.RawMessage `json:"delete_post"`
	AddPostAttachment    *json.RawMessage `json:"add_post_attachment"`
	RemovePostAttachment *json.RawMessage `json:"remove_post_attachment"`
	AnswerPoll           *json.RawMessage `json:"answer_poll"`
}

type PostsQuery struct {
	SubspacePosts   *json.RawMessage `json:"subspace_posts"`
	SectionPosts    *json.RawMessage `json:"section_posts"`
	Post            *json.RawMessage `json:"post"`
	PostAttachments *json.RawMessage `json:"post_attachments"`
	PollAnswers     *json.RawMessage `json:"poll_answers"`
}
