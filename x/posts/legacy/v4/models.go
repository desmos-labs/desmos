package v4

// DONTCOVER

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
)

// NewPost allows to build a new Post instance
func NewPost(
	subspaceID uint64,
	sectionID uint32,
	id uint64,
	externalID string,
	text string,
	author string,
	conversationID uint64,
	entities *Entities,
	tags []string,
	referencedPosts []PostReference,
	replySetting ReplySetting,
	creationDate time.Time,
	lastEditedDate *time.Time,
) Post {
	return Post{
		SubspaceID:      subspaceID,
		SectionID:       sectionID,
		ID:              id,
		ExternalID:      externalID,
		Text:            text,
		Entities:        entities,
		Tags:            tags,
		Author:          author,
		ConversationID:  conversationID,
		ReferencedPosts: referencedPosts,
		ReplySettings:   replySetting,
		CreationDate:    creationDate,
		LastEditedDate:  lastEditedDate,
	}
}

// NewPostReference returns a new PostReference instance
func NewPostReference(referenceType PostReferenceType, postID uint64, position uint64) PostReference {
	return PostReference{
		Type:     referenceType,
		PostID:   postID,
		Position: position,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// NewEntities returns a new Entities instance
func NewEntities(hashtags []TextTag, mentions []TextTag, urls []Url) *Entities {
	return &Entities{
		Hashtags: hashtags,
		Mentions: mentions,
		Urls:     urls,
	}
}

// NewTextTag returns a new TextTag instance
func NewTextTag(start, end uint64, tag string) TextTag {
	return TextTag{
		Start: start,
		End:   end,
		Tag:   tag,
	}
}

// NewURL returns a new Url instance
func NewURL(start, end uint64, url, displayURL string) Url {
	return Url{
		Start:      start,
		End:        end,
		Url:        url,
		DisplayUrl: displayURL,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// NewAttachment returns a new Attachment instance
func NewAttachment(subspaceID uint64, postID uint64, id uint32, content AttachmentContent) Attachment {
	contentAny, err := codectypes.NewAnyWithValue(content)
	if err != nil {
		panic("failed to pack content to any type")
	}

	return Attachment{
		SubspaceID: subspaceID,
		PostID:     postID,
		ID:         id,
		Content:    contentAny,
	}
}

// UnpackInterfaces implements codecUnpackInterfacesMessage
func (a *Attachment) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var content AttachmentContent
	return unpacker.UnpackAny(a.Content, &content)
}

// --------------------------------------------------------------------------------------------------------------------

// AttachmentContent represents an attachment content
type AttachmentContent interface {
	proto.Message
	isAttachmentContent()
}

// --------------------------------------------------------------------------------------------------------------------

var _ AttachmentContent = &Media{}

// NewMedia returns a new Media instance
func NewMedia(uri, mimeType string) *Media {
	return &Media{
		Uri:      uri,
		MimeType: mimeType,
	}
}

func (*Media) isAttachmentContent() {}

// --------------------------------------------------------------------------------------------------------------------

var _ AttachmentContent = &Poll{}

// NewPoll returns a new Poll instance
func NewPoll(
	question string,
	providedAnswers []Poll_ProvidedAnswer,
	endDate time.Time,
	allowsMultipleAnswers bool,
	allowsAnswerEdits bool,
	tallyResults *PollTallyResults,
) *Poll {
	return &Poll{
		Question:              question,
		ProvidedAnswers:       providedAnswers,
		EndDate:               endDate,
		AllowsMultipleAnswers: allowsMultipleAnswers,
		AllowsAnswerEdits:     allowsAnswerEdits,
		FinalTallyResults:     tallyResults,
	}
}

func (*Poll) isAttachmentContent() {}

// --------------------------------------------------------------------------------------------------------------------

// NewProvidedAnswer returns a new Poll_ProvidedAnswer instance
func NewProvidedAnswer(text string, attachments []Attachment) Poll_ProvidedAnswer {
	return Poll_ProvidedAnswer{
		Text:        text,
		Attachments: attachments,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// NewUserAnswer returns a new UserAnswer instance
func NewUserAnswer(subspaceID uint64, postID uint64, pollID uint32, answersIndexes []uint32, user string) UserAnswer {
	return UserAnswer{
		SubspaceID:     subspaceID,
		PostID:         postID,
		PollID:         pollID,
		AnswersIndexes: answersIndexes,
		User:           user,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// NewPollTallyResults returns a new PollTallyResults instance
func NewPollTallyResults(results []PollTallyResults_AnswerResult) *PollTallyResults {
	return &PollTallyResults{
		Results: results,
	}
}

// NewAnswerResult returns a new PollTallyResults_AnswerResult instance
func NewAnswerResult(answerIndex uint32, votes uint64) PollTallyResults_AnswerResult {
	return PollTallyResults_AnswerResult{
		AnswerIndex: answerIndex,
		Votes:       votes,
	}
}
