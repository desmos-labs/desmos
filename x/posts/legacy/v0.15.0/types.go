package v0150

import (
	"strconv"
	"time"

	v0100 "github.com/desmos-labs/desmos/x/posts/legacy/v0.10.0"
	v0120 "github.com/desmos-labs/desmos/x/posts/legacy/v0.12.0"
	v0130 "github.com/desmos-labs/desmos/x/posts/legacy/v0.13.0"
	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060 "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	v080posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
)

// GenesisState contains the data of a v0.15.0 genesis state for the posts module
type GenesisState struct {
	Posts               []Post               `json:"posts"`
	UsersPollAnswers    []UserAnswersEntry   `json:"users_poll_answers"`
	PostReactions       []PostReactionsEntry `json:"post_reactions"`
	RegisteredReactions []RegisteredReaction `json:"registered_reactions"`
	Params              v080posts.Params     `json:"params"`
}

// UserPollAnswerEntry represents an entry containing all the answers to a poll
type UserAnswersEntry struct {
	PostID      string       `json:"post_id,omitempty" yaml:"post_id,omitempty"`
	UserAnswers []UserAnswer `json:"user_answers"`
}

// UserAnswer contains the data of a user's answer to a poll
type UserAnswer struct {
	User    string   `json:"user,omitempty" yaml:"user,omitempty"`
	Answers []string `json:"answers,omitempty" yaml:"answers,omitempty"`
}

// newUserAnswerEntry create a new userAnswerEntry from the old data from genesis
func newUserAnswerEntry(postID string, oldUsersAnswers []v040posts.UserAnswer) UserAnswersEntry {
	userAnswers := make([]UserAnswer, len(oldUsersAnswers))

	for index, oldUserAnswers := range oldUsersAnswers {
		answers := make([]string, len(oldUserAnswers.Answers))
		for index, id := range oldUserAnswers.Answers {
			answers[index] = strconv.FormatUint(uint64(id), 10)
		}
		userAnswers[index] = UserAnswer{
			User:    oldUserAnswers.User.String(),
			Answers: answers,
		}
	}

	return UserAnswersEntry{
		PostID:      postID,
		UserAnswers: userAnswers,
	}
}

// PostReactionEntry represents an entry containing all the reactions to a post
type PostReactionsEntry struct {
	PostID    string         `json:"post_id,omitempty" yaml:"post_id,omitempty"`
	Reactions []PostReaction `json:"reactions" yaml:"reactions"`
}

// PostReaction is a struct of a user reaction to a post
type PostReaction struct {
	ShortCode string `json:"short_code" yaml:"short_code"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
	Owner     string `json:"owner,omitempty" yaml:"owner,omitempty"`
}

// newPostReactionEntry create a new PostReactionEntry from the old data from genesis
func newPostReactionEntry(postID string, oldPostReactions []v060.PostReaction) PostReactionsEntry {
	reactions := make([]PostReaction, len(oldPostReactions))

	for index, oldPostReaction := range oldPostReactions {
		reactions[index] = PostReaction{
			ShortCode: oldPostReaction.Shortcode,
			Value:     oldPostReaction.Value,
			Owner:     oldPostReaction.Owner.String(),
		}
	}

	return PostReactionsEntry{
		PostID:    postID,
		Reactions: reactions,
	}
}

// RegisteredReaction represents a registered reaction that can be referenced
// by its shortCode inside post reactions
type RegisteredReaction struct {
	ShortCode string `json:"short_code" yaml:"short_code"`
	Value     string `json:"value,omitempty" yaml:"value,omitempty"`
	Subspace  string `json:"subspace,omitempty" yaml:"subspace,omitempty"`
	Creator   string `json:"creator,omitempty" yaml:"creator,omitempty"`
}

type Post struct {
	PostID         string                    `json:"id" yaml:"id" `                                          // Unique id
	ParentID       string                    `json:"parent_id" yaml:"parent_id"`                             // Post of which this one is a comment
	Message        string                    `json:"message" yaml:"message"`                                 // Message contained inside the post
	Created        time.Time                 `json:"created" yaml:"created"`                                 // RFC3339 date at which the post has been created
	LastEdited     time.Time                 `json:"last_edited" yaml:"last_edited"`                         // RFC3339 date at which the post has been edited the last time
	AllowsComments bool                      `json:"allows_comments" yaml:"allows_comments"`                 // Tells if users can reference this PostID as the parent
	Subspace       string                    `json:"subspace" yaml:"subspace"`                               // Identifies the application that has posted the message
	OptionalData   []v0130.OptionalDataEntry `json:"optional_data,omitempty" yaml:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        string                    `json:"creator" yaml:"creator"`                                 // Creator of the Post
	Attachments    []v0100.Attachment        `json:"attachments,omitempty" yaml:"attachments,omitempty"`     // Contains all the attachments that are shared with the post
	PollData       *v0120.PollData           `json:"poll_data,omitempty" yaml:"poll_data,omitempty"`         // Contains the poll details, if existing
}
