package types

import (
	"fmt"
	"sort"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewPost allows to build a new Post instance
func NewPost(
	subspaceID uint64,
	id uint64,
	externalID string,
	text string,
	entities *Entities,
	author string,
	conversationID uint64,
	referencedPosts []PostReference,
	replySetting ReplySetting,
	creationDate time.Time,
	lastEditedDate *time.Time,
) Post {
	return Post{
		SubspaceID:      subspaceID,
		ID:              id,
		ExternalID:      externalID,
		Text:            text,
		Entities:        entities,
		Author:          author,
		ConversationId:  conversationID,
		ReferencedPosts: referencedPosts,
		ReplySettings:   replySetting,
		CreationDate:    creationDate,
		LastEditedDate:  lastEditedDate,
	}
}

func (p Post) Validate() error {
	if p.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", p.SubspaceID)
	}

	if p.ID == 0 {
		return fmt.Errorf("invalid post id: %d", p.ID)
	}

	if p.Entities != nil {
		err := p.Entities.Validate()
		if err != nil {
			return fmt.Errorf("invalid entities: %s", err)
		}
	}

	_, err := sdk.AccAddressFromBech32(p.Author)
	if err != nil {
		return fmt.Errorf("invalid author address: %s", err)
	}

	for _, reference := range p.ReferencedPosts {
		err = reference.Validate()
		if err != nil {
			return fmt.Errorf("invalid post reference: %s", err)
		}
	}

	if p.CreationDate.IsZero() {
		return fmt.Errorf("invalid post creation date: %s", err)
	}

	if p.LastEditedDate != nil && p.LastEditedDate.IsZero() {
		return fmt.Errorf("invalid post last edited date: %s", err)
	}

	return nil
}

// NewPostReference returns a new PostReference instance
func NewPostReference(referenceType PostReference_Type, postID uint64) PostReference {
	return PostReference{
		Type:   referenceType,
		PostID: postID,
	}
}

func (r PostReference) Validate() error {
	if r.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", r.PostID)
	}

	return nil
}

// NewEntities returns a new Entities instance
func NewEntities(hashtags []Tag, mentions []Tag, urls []Url) *Entities {
	return &Entities{
		Hashtags: hashtags,
		Mentions: mentions,
		Urls:     urls,
	}
}

func (e *Entities) Validate() error {
	for _, tag := range e.Hashtags {
		err := tag.Validate()
		if err != nil {
			return fmt.Errorf("invalid hashtag: %s", err)
		}
	}

	for _, tag := range e.Mentions {
		err := tag.Validate()
		if err != nil {
			return fmt.Errorf("invalid mention: %s", err)
		}
	}

	for _, url := range e.Urls {
		err := url.Validate()
		if err != nil {
			return fmt.Errorf("invalid url: %s", err)
		}
	}

	// TODO: Make sure there are no overlapping entities based on (start, end)

	// Map all entities into segments

	type entitySegment struct {
		start uint64
		end   uint64
	}

	segments := make([]entitySegment, len(e.Hashtags)+len(e.Mentions)+len(e.Urls))
	if len(segments) < 1 {
		// We cannot have an empty entities here
		return fmt.Errorf("entities must have at least one entity inside")
	}

	if len(segments) < 2 {
		// With less than 2 segments there cannot be any overlap
		return nil
	}

	i := 0
	for _, hashtag := range e.Hashtags {
		segments[i] = entitySegment{start: hashtag.Start, end: hashtag.End}
		i++
	}

	for _, mention := range e.Mentions {
		segments[i] = entitySegment{start: mention.Start, end: mention.End}
		i++
	}

	for _, url := range e.Urls {
		segments[i] = entitySegment{start: url.Start, end: url.End}
		i++
	}

	// Sort the segments
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].start < segments[j].end
	})

	for index := 0; index < len(segments)-1; index++ {
		first, second := segments[index], segments[index+1]
		if first.end >= second.start {
			return fmt.Errorf("entities cannot overlap: start %d end %d", first.end, second.start)
		}
	}

	return nil
}

// NewTag returns a new Tag instance
func NewTag(start, end uint64, tag string) Tag {
	return Tag{
		Start: start,
		End:   end,
		Tag:   tag,
	}
}

func (t Tag) Validate() error {
	if t.Start > t.End {
		return fmt.Errorf("invalid start and end indexes: %d %d", t.Start, t.End)
	}

	if strings.TrimSpace(t.Tag) == "" {
		return fmt.Errorf("invalid start and end indexes: %d %d", t.Start, t.End)
	}

	return nil
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

func (u Url) Validate() error {
	if u.Start > u.End {
		return fmt.Errorf("invalid start and end indexes: %d %d", u.Start, u.End)
	}

	if strings.TrimSpace(u.Url) == "" {
		return fmt.Errorf("invalid url value: %s", u.Url)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a Attachment) Validate() error {
	if a.Id == 0 {
		return fmt.Errorf("invalid attachment id: %d", a.Id)
	}

	if pollAttachment, ok := a.Sum.(*Attachment_Poll); ok {
		return pollAttachment.Poll.Validate()
	}

	if mediaAttachment, ok := a.Sum.(*Attachment_Media); ok {
		return mediaAttachment.Media.Validate()
	}

	return nil
}

type Attachments []Attachment

func (a Attachments) Validate() error {
	ids := map[uint32]int{}
	for _, attachment := range a {
		if _, ok := ids[attachment.Id]; ok {
			return fmt.Errorf("duplicated attachment id: %d", attachment.Id)
		}
		ids[attachment.Id] = 1

		err := attachment.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// NewPollAttachment returns a new Attachment instance containing the given poll
func NewPollAttachment(id uint32, poll Poll) Attachment {
	return Attachment{
		Id: id,
		Sum: &Attachment_Poll{
			Poll: &poll,
		},
	}
}

// NewMediaAttachment returns a new Attachment instance containing the given media
func NewMediaAttachment(id uint32, media Media) Attachment {
	return Attachment{
		Id: id,
		Sum: &Attachment_Media{
			Media: &media,
		},
	}
}

// NewMedia returns a new Media instance
func NewMedia(uri, mimeType string) Media {
	return Media{
		Uri:      uri,
		MimeType: mimeType,
	}
}

func (m Media) Validate() error {
	if strings.TrimSpace(m.Uri) == "" {
		return fmt.Errorf("invalid uri: %s", m.Uri)
	}

	if strings.TrimSpace(m.MimeType) == "" {
		return fmt.Errorf("invalid mime type: %s", m.MimeType)
	}

	return nil
}

// NewPoll returns a new Poll instance
func NewPoll(
	question string,
	providedAnswers []Poll_ProvidedAnswer,
	endDate time.Time,
	allowsMultipleAnswers bool,
	allowsAnswerEdits bool,
) Poll {
	return Poll{
		Question:              question,
		ProvidedAnswers:       providedAnswers,
		EndDate:               endDate,
		AllowsMultipleAnswers: allowsMultipleAnswers,
		AllowsAnswerEdits:     allowsAnswerEdits,
	}
}

func (p Poll) Validate() error {
	if strings.TrimSpace(p.Question) == "" {
		return fmt.Errorf("invalid question: %s", p.Question)
	}

	if len(p.ProvidedAnswers) < 2 {
		return fmt.Errorf("insufficient amount of provided answers: %d", len(p.ProvidedAnswers))
	}

	answers := map[string]int{}
	for _, answer := range p.ProvidedAnswers {
		err := answer.Validate()
		if err != nil {
			return err
		}

		if _, ok := answers[answer.Text]; ok {
			return fmt.Errorf("duplicated provided answer: %s", answer.Text)
		}
		answers[answer.Text] = 1
	}

	if p.EndDate.IsZero() {
		return fmt.Errorf("invalid end date: %s", p.EndDate)
	}

	return nil
}

// NewProvidedAnswer returns a new Poll_ProvidedAnswer instance
func NewProvidedAnswer(text string, attachments []Attachment) Poll_ProvidedAnswer {
	return Poll_ProvidedAnswer{
		Text:        text,
		Attachments: attachments,
	}
}

func (a Poll_ProvidedAnswer) Validate() error {
	if strings.TrimSpace(a.Text) == "" {
		return fmt.Errorf("invalid text: %s", a.Text)
	}

	return Attachments(a.Attachments).Validate()
}

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

func (a UserAnswer) Validate() error {
	if a.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", a.SubspaceID)
	}

	if a.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", a.PostID)
	}

	if a.PollID == 0 {
		return fmt.Errorf("invalid poll id: %d", a.PollID)
	}

	if len(a.AnswersIndexes) == 0 {
		return fmt.Errorf("answer indexes cannot be empty")
	}

	indexes := map[uint32]int{}
	for _, answer := range a.AnswersIndexes {
		if _, ok := indexes[answer]; ok {
			return fmt.Errorf("duplicated answer index: %d", answer)
		}
		indexes[answer] = 1
	}

	_, err := sdk.AccAddressFromBech32(a.User)
	if err != nil {
		return fmt.Errorf("invalid user address: %s", err)
	}

	return nil
}

// NewPollTallyResults returns a new PollTallyResults instance
func NewPollTallyResults(
	subspaceID uint64,
	postID uint64,
	pollId uint32,
	results []PollTallyResults_AnswerResult,
) PollTallyResults {
	return PollTallyResults{
		SubspaceID: subspaceID,
		PostID:     postID,
		PollID:     pollId,
		Results:    results,
	}
}

func (r PollTallyResults) Validate() error {
	if r.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", r.SubspaceID)
	}

	if r.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", r.PostID)
	}

	if r.PollID == 0 {
		return fmt.Errorf("invalid poll id: %d", r.PollID)
	}

	if len(r.Results) == 0 {
		return fmt.Errorf("empty answer results")
	}

	ids := map[uint32]int{}
	for _, answerResult := range r.Results {
		if _, ok := ids[answerResult.AnswerIndex]; ok {
			return fmt.Errorf("duplicated result for answer %d", answerResult.AnswerIndex)
		}
		ids[answerResult.AnswerIndex] = 1
	}

	return nil
}

// NewAnswerResult returns a new PollTallyResults_AnswerResult instance
func NewAnswerResult(answerIndex uint32, votes uint64) PollTallyResults_AnswerResult {
	return PollTallyResults_AnswerResult{
		AnswerIndex: answerIndex,
		Votes:       votes,
	}
}
