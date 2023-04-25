package types

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ParsePostID parses the given value as a post id, returning an error if it's invalid
func ParsePostID(value string) (uint64, error) {
	if value == "" {
		return 0, nil
	}

	postID, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid post id: %s", err)
	}
	return postID, nil
}

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
	owner string,
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
		Owner:           owner,
	}
}

// Validate implements fmt.Validator
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

	for _, tag := range p.Tags {
		if strings.TrimSpace(tag) == "" {
			return fmt.Errorf("invalid post tag: %s", tag)
		}
	}

	_, err := sdk.AccAddressFromBech32(p.Author)
	if err != nil {
		return fmt.Errorf("invalid author address: %s", err)
	}

	if p.ConversationID >= p.ID {
		return fmt.Errorf("invalid conversation id: %d", p.ConversationID)
	}

	for _, reference := range p.ReferencedPosts {
		err = reference.Validate()
		if err != nil {
			return fmt.Errorf("invalid post reference: %s", err)
		}

		if reference.PostID >= p.ID {
			return fmt.Errorf("invalid referenced post id: %d", reference.PostID)
		}
	}

	if p.ReplySettings == REPLY_SETTING_UNSPECIFIED {
		return fmt.Errorf("invalid reply setting: %s", p.ReplySettings)
	}

	if p.CreationDate.IsZero() {
		return fmt.Errorf("invalid post creation date: %s", err)
	}

	if p.LastEditedDate != nil {
		// Instead of zero, we should use nil
		if p.LastEditedDate.IsZero() {
			return fmt.Errorf("invalid post last edited date: %s", err)
		}

		// Make sure the creation date is always before the last edit date
		if p.LastEditedDate.Before(p.CreationDate) {
			return fmt.Errorf("last edited date cannot be before the creation date")
		}
	}

	_, err = sdk.AccAddressFromBech32(p.Owner)
	if err != nil {
		return fmt.Errorf("invalid owner address: %s", err)
	}

	return nil
}

// IsUserMentioned tells whether the given user is mentioned inside the post or not
func (p Post) IsUserMentioned(user string) bool {
	for _, mention := range p.GetMentionedUsers() {
		if mention == user {
			return true
		}
	}
	return false
}

// GetMentionedUsers returns all the mentioned users
func (p Post) GetMentionedUsers() []string {
	if p.Entities == nil {
		return nil
	}

	mentions := make([]string, len(p.Entities.Mentions))
	for i, mention := range p.Entities.Mentions {
		mentions[i] = mention.Tag
	}
	return mentions
}

// NewPostReference returns a new PostReference instance
func NewPostReference(referenceType PostReferenceType, postID uint64, position uint64) PostReference {
	return PostReference{
		Type:     referenceType,
		PostID:   postID,
		Position: position,
	}
}

// Validate implements fmt.Validator
func (r PostReference) Validate() error {
	if r.Type == POST_REFERENCE_TYPE_UNSPECIFIED {
		return fmt.Errorf("invalid reference type: %s", r.Type)
	}

	if r.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", r.PostID)
	}

	if r.Type != POST_REFERENCE_TYPE_QUOTE && r.Position > 0 {
		return fmt.Errorf("reference position should be set only with POST_REFERENCE_TYPE_QUOTE")
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// PostUpdate contains all the data that can be updated about a post.
// When performing an update, if a text field should not be edited then it must be set to types.DoNotModify.
type PostUpdate struct {
	// If it shouldn't replace the current text, it must be set to types.DoNotModify
	Text string

	// Update's entities will always replace the existing ones
	Entities *Entities

	// Update's tags will always replace the existing ones
	Tags []string

	UpdateTime time.Time
}

// NewPostUpdate returns a new PostUpdate instance
func NewPostUpdate(text string, entities *Entities, tags []string, updateTime time.Time) PostUpdate {
	return PostUpdate{
		Text:       text,
		Entities:   entities,
		Tags:       tags,
		UpdateTime: updateTime,
	}
}

// Update updates the fields of a given post without validating it.
// Before storing the updated post, a validation with keeper.ValidatePost should
// be performed.
func (p Post) Update(update PostUpdate) Post {
	if update.Text == DoNotModify {
		update.Text = p.Text
	}

	return NewPost(
		p.SubspaceID,
		p.SectionID,
		p.ID,
		p.ExternalID,
		update.Text,
		p.Author,
		p.ConversationID,
		update.Entities,
		update.Tags,
		p.ReferencedPosts,
		p.ReplySettings,
		p.CreationDate,
		&update.UpdateTime,
		p.Owner,
	)
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

// Validate implements fmt.Validator
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

	// --- Make sure there are no overlapping entities based on (start, end) ---

	// Map all entities into segments
	segments := e.getSegments()

	if len(segments) < 1 {
		// We cannot have an empty entities here
		return fmt.Errorf("entities must have at least one entity inside")
	}

	if len(segments) < 2 {
		// With less than 2 segments there cannot be any overlap
		return nil
	}

	// Sort the segments
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].start < segments[j].start
	})

	// Verify there are no overlapping segments
	for index := 0; index < len(segments)-1; index++ {
		first, second := segments[index], segments[index+1]
		if first.end >= second.start {
			return fmt.Errorf("entities cannot overlap: start %d end %d", first.end, second.start)
		}
	}

	return nil
}

// entitySegment represents an entity as a segment within a post text.
// This is used to easily determine whether there are overlapping entities specified within the post.
// Two entities are said to overlap if the end of the first is after the start of the second.
type entitySegment struct {
	// start represents the index inside the post text at which this segment begins
	start uint64

	// end represents the index inside the post text at which this segment ends
	end uint64
}

// getSegments maps all the provided Entities into segments
func (e *Entities) getSegments() []entitySegment {
	segments := make([]entitySegment, len(e.Hashtags)+len(e.Mentions)+len(e.Urls))
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

	return segments
}

// NewTextTag returns a new TextTag instance
func NewTextTag(start, end uint64, tag string) TextTag {
	return TextTag{
		Start: start,
		End:   end,
		Tag:   tag,
	}
}

// Validate implements fmt.Validator
func (t TextTag) Validate() error {
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

// Validate implements fmt.Validator
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

// ParseAttachmentID parses the given value as an attachment id, returning an error if it's invalid
func ParseAttachmentID(value string) (uint32, error) {
	if value == "" {
		return 0, nil
	}

	attachmentID, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid attachment id: %s", err)
	}
	return uint32(attachmentID), nil
}

// Attachments represents a slice of Attachment objects
type Attachments []Attachment

// Validate implements fmt.Validators fmt.Validator
func (a Attachments) Validate() error {
	ids := map[uint32]bool{}
	for _, attachment := range a {
		if _, ok := ids[attachment.ID]; ok {
			return fmt.Errorf("duplicated attachment id: %d", attachment.ID)
		}
		ids[attachment.ID] = true

		err := attachment.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

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

// Validate implements fmt.Validator
func (a Attachment) Validate() error {
	if a.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", a.SubspaceID)
	}

	if a.PostID == 0 {
		return fmt.Errorf("invalid post id: %d", a.PostID)
	}

	if a.ID == 0 {
		return fmt.Errorf("invalid attachment id: %d", a.ID)
	}

	if a.Content == nil {
		return fmt.Errorf("invalid attachment content")
	}

	return a.Content.GetCachedValue().(AttachmentContent).Validate()
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (a *Attachment) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	var content AttachmentContent
	return unpacker.UnpackAny(a.Content, &content)
}

// --------------------------------------------------------------------------------------------------------------------

// AttachmentContent represents an attachment content
type AttachmentContent interface {
	proto.Message
	isAttachmentContent()
	Validate() error
	Equal(other interface{}) bool
}

// PackAttachments packs the given AttachmentContent instances as Any instances
func PackAttachments(attachments []AttachmentContent) ([]*codectypes.Any, error) {
	if attachments == nil {
		// Avoid allocating a new array
		return nil, nil
	}

	attachmentAnys := make([]*codectypes.Any, len(attachments))
	for i := range attachments {
		attachmentAny, err := codectypes.NewAnyWithValue(attachments[i])
		if err != nil {
			return nil, err
		}
		attachmentAnys[i] = attachmentAny
	}
	return attachmentAnys, nil
}

// UnpackAttachments unpacks the given Any instances as AttachmentContent
func UnpackAttachments(cdc codectypes.AnyUnpacker, attachmentAnys []*codectypes.Any) ([]AttachmentContent, error) {
	attachments := make([]AttachmentContent, len(attachmentAnys))
	for i, any := range attachmentAnys {
		var content AttachmentContent
		err := cdc.UnpackAny(any, &content)
		if err != nil {
			return nil, err
		}
		attachments[i] = content
	}
	return attachments, nil
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

// Validate implements fmt.Validator
func (m *Media) Validate() error {
	if strings.TrimSpace(m.Uri) == "" {
		return fmt.Errorf("invalid uri: %s", m.Uri)
	}

	if strings.TrimSpace(m.MimeType) == "" {
		return fmt.Errorf("invalid mime type: %s", m.MimeType)
	}

	return nil
}

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

// Validate implements fmt.Validator
func (p *Poll) Validate() error {
	if strings.TrimSpace(p.Question) == "" {
		return fmt.Errorf("invalid question: %s", p.Question)
	}

	if len(p.ProvidedAnswers) < 2 {
		return fmt.Errorf("insufficient amount of provided answers: %d", len(p.ProvidedAnswers))
	}

	answers := map[string]bool{}
	for _, answer := range p.ProvidedAnswers {
		err := answer.Validate()
		if err != nil {
			return err
		}

		if _, ok := answers[answer.Text]; ok {
			return fmt.Errorf("duplicated provided answer: %s", answer.Text)
		}
		answers[answer.Text] = true
	}

	if p.EndDate.IsZero() {
		return fmt.Errorf("invalid end date: %s", p.EndDate)
	}

	if p.FinalTallyResults != nil {
		err := p.FinalTallyResults.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// IsPoll tells whether the given attachment represents a poll or not
func IsPoll(attachment Attachment) bool {
	_, ok := attachment.Content.GetCachedValue().(*Poll)
	return ok
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (a *Poll) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, answer := range a.ProvidedAnswers {
		err := answer.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewProvidedAnswer returns a new Poll_ProvidedAnswer instance
func NewProvidedAnswer(text string, attachments []AttachmentContent) Poll_ProvidedAnswer {
	attachmentAnys, err := PackAttachments(attachments)
	if err != nil {
		panic(err)
	}

	return Poll_ProvidedAnswer{
		Text:        text,
		Attachments: attachmentAnys,
	}
}

// Validate implements fmt.Validator
func (a Poll_ProvidedAnswer) Validate() error {
	if strings.TrimSpace(a.Text) == "" {
		return fmt.Errorf("invalid text: %s", a.Text)
	}

	// Unpack the attachments
	attachments := make([]AttachmentContent, len(a.Attachments))
	for i, attachmentAny := range a.Attachments {
		attachments[i] = attachmentAny.GetCachedValue().(AttachmentContent)
	}

	// Validate the attachments
	for _, attachment := range attachments {
		if _, isPoll := attachment.(*Poll); isPoll {
			return fmt.Errorf("cannot have a poll as an attachment of a poll's provided answer")
		}

		if containsDuplicatedAttachments(attachments, attachment) {
			return fmt.Errorf("duplicated attachment")
		}

		err := attachment.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// containsDuplicatedAttachments tells whether the given attachments array contains a duplicate content
func containsDuplicatedAttachments(attachments []AttachmentContent, content AttachmentContent) bool {
	var found = 0
	for _, attachment := range attachments {
		if attachment.Equal(content) {
			found++
		}
	}
	return found > 1
}

// UnpackInterfaces implements codectypes.UnpackInterfacesMessage
func (a *Poll_ProvidedAnswer) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	for _, attachment := range a.Attachments {
		var content AttachmentContent
		err := unpacker.UnpackAny(attachment, &content)
		if err != nil {
			return err
		}
	}
	return nil
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

// Validate implements fmt.Validator
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

	indexes := map[uint32]bool{}
	for _, answer := range a.AnswersIndexes {
		if _, ok := indexes[answer]; ok {
			return fmt.Errorf("duplicated answer index: %d", answer)
		}
		indexes[answer] = true
	}

	_, err := sdk.AccAddressFromBech32(a.User)
	if err != nil {
		return fmt.Errorf("invalid user address: %s", err)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

// NewPollTallyResults returns a new PollTallyResults instance
func NewPollTallyResults(results []PollTallyResults_AnswerResult) *PollTallyResults {
	return &PollTallyResults{
		Results: results,
	}
}

// Validate implements fmt.Validator
func (r *PollTallyResults) Validate() error {
	if len(r.Results) == 0 {
		return fmt.Errorf("empty answer results")
	}

	ids := map[uint32]bool{}
	for _, answerResult := range r.Results {
		if _, ok := ids[answerResult.AnswerIndex]; ok {
			return fmt.Errorf("duplicated result for answer %d", answerResult.AnswerIndex)
		}
		ids[answerResult.AnswerIndex] = true
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
