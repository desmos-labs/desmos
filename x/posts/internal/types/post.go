package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/desmos-labs/desmos/x/commons"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Post id
// ---------------

// PostID represents a unique post id
type PostID uint64

// Valid tells if the id can be used safely
func (id PostID) Valid() bool {
	return id != 0
}

// Next returns the subsequent id to this one
func (id PostID) Next() PostID {
	return id + 1
}

// String implements fmt.Stringer
func (id PostID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// Equals compares two PostID instances
func (id PostID) Equals(other PostID) bool {
	return id == other
}

// MarshalJSON implements Marshaler
func (id PostID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON implements Unmarshaler
func (id *PostID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	postID, err := ParsePostID(s)
	if err != nil {
		return err
	}

	*id = postID
	return nil
}

// ParsePostID returns the PostID represented inside the provided
// value, or an error if no id could be parsed properly
func ParsePostID(value string) (PostID, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return PostID(0), err
	}

	return PostID(intVal), err
}

// ----------------
// --- Post IDs
// ----------------

// PostIDs represents a slice of PostID objects
type PostIDs []PostID

// Equals returns true iff the ids slice and the other
// one contain the same data in the same order
func (ids PostIDs) Equals(other PostIDs) bool {
	if len(ids) != len(other) {
		return false
	}

	for index, id := range ids {
		if id != other[index] {
			return false
		}
	}

	return true
}

// AppendIfMissing appends the given postID to the ids slice if it does not exist inside it yet.
// It returns a new slice of PostIDs containing such ID and a boolean indicating whether or not the original
// slice has been modified.
func (ids PostIDs) AppendIfMissing(id PostID) (PostIDs, bool) {
	for _, ele := range ids {
		if ele.Equals(id) {
			return ids, false
		}
	}
	return append(ids, id), true
}

// ---------------
// --- Post
// ---------------

// Post is a struct of a post
type Post struct {
	PostID         PostID         `json:"id"`                      // Unique id
	ParentID       PostID         `json:"parent_id"`               // Post of which this one is a comment
	Message        string         `json:"message"`                 // Message contained inside the post
	Created        time.Time      `json:"created"`                 // RFC3339 date at which the post has been created
	LastEdited     time.Time      `json:"last_edited"`             // RFC3339 date at which the post has been edited the last time
	AllowsComments bool           `json:"allows_comments"`         // Tells if users can reference this PostID as the parent
	Subspace       string         `json:"subspace"`                // Identifies the application that has posted the message
	OptionalData   OptionalData   `json:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        sdk.AccAddress `json:"creator"`                 // Creator of the Post
	Medias         PostMedias     `json:"medias,omitempty"`        // Contains all the medias that are shared with the post
	PollData       *PollData      `json:"poll_data,omitempty"`     // Contains the poll details, if existing
}

func NewPost(id, parentID PostID, message string, allowsComments bool, subspace string,
	optionalData map[string]string, created time.Time, creator sdk.AccAddress) Post {
	return Post{
		PostID:         id,
		ParentID:       parentID,
		Message:        message,
		Created:        created,
		LastEdited:     time.Time{},
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Creator:        creator,
	}
}

// WithMedias allows to easily set the given medias as the multimedia files associated with the p Post
func (p Post) WithMedias(medias PostMedias) Post {
	p.Medias = medias
	return p
}

// WithMedias allows to easily set the given data as the poll data files associated with the p Post
func (p Post) WithPollData(data PollData) Post {
	p.PollData = &data
	return p
}

// String implements fmt.Stringer
func (p Post) String() string {
	bytes, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate implements validator
func (p Post) Validate() error {
	if !p.PostID.Valid() {
		return fmt.Errorf("invalid post id: %s", p.PostID)
	}

	if p.Creator == nil {
		return fmt.Errorf("invalid post owner: %s", p.Creator)
	}

	if len(strings.TrimSpace(p.Message)) == 0 && len(p.Medias) == 0 {
		return fmt.Errorf("post message or medias required, they cannot be both empty")
	}

	if len(p.Message) > MaxPostMessageLength {
		return fmt.Errorf("post message cannot be longer than %d characters", MaxPostMessageLength)
	}

	if !SubspaceRegEx.MatchString(p.Subspace) {
		return fmt.Errorf("post subspace must be a valid sha-256 hash")
	}

	if p.Created.IsZero() {
		return fmt.Errorf("invalid post creation time: %s", p.Created)
	}

	if p.Created.After(time.Now().UTC()) {
		return fmt.Errorf("post creation date cannot be in the future")
	}

	if !p.LastEdited.IsZero() && p.LastEdited.Before(p.Created) {
		return fmt.Errorf("invalid post last edit time: %s", p.LastEdited)
	}

	if !p.LastEdited.IsZero() && p.LastEdited.After(time.Now().UTC()) {
		return fmt.Errorf("post last edit date cannot be in the future")
	}

	if len(p.OptionalData) > MaxOptionalDataFieldsNumber {
		return fmt.Errorf("post optional data cannot contain more than %d key-value pairs", MaxOptionalDataFieldsNumber)
	}

	for key, value := range p.OptionalData {
		if len(value) > MaxOptionalDataFieldValueLength {
			return fmt.Errorf(
				"post optional data values cannot exceed %d characters. %s of post with id %s is longer than this",
				MaxOptionalDataFieldValueLength, key, p.PostID,
			)
		}
	}

	if err := p.Medias.Validate(); err != nil {
		return err
	}

	if err := p.PollData.Validate(); err != nil {
		return err
	}

	return nil
}

// Equals allows to check whether the contents of p are the same of other
func (p Post) Equals(other Post) bool {
	return p.PostID.Equals(other.PostID) && p.ContentsEquals(other)
}

// IsConflictingWith returns true if other is somehow conflicting with this post.
// In order to not be conflicting, two posts should have a different ID and also either:
// - a different creation date
// - a different subspace
// - a different creator
//
// If they have the same creation date, subspace and creator they are considered conflicting.
func (p Post) IsConflictingWith(other Post) bool {
	return p.PostID.Equals(other.PostID) ||
		(p.Created.Equal(other.Created) && p.Subspace == other.Subspace && p.Creator.Equals(other.Creator))
}

// ContentsEquals returns true if and only if p and other contain the same data, without considering the ID
func (p Post) ContentsEquals(other Post) bool {
	equalsOptionalData := len(p.OptionalData) == len(other.OptionalData)
	if equalsOptionalData {
		for key := range p.OptionalData {
			equalsOptionalData = equalsOptionalData && p.OptionalData[key] == other.OptionalData[key]
		}
	}

	return p.ParentID.Equals(other.ParentID) &&
		p.Message == other.Message &&
		p.Created.Equal(other.Created) &&
		p.LastEdited.Equal(other.LastEdited) &&
		p.AllowsComments == other.AllowsComments &&
		p.Subspace == other.Subspace &&
		equalsOptionalData &&
		p.Creator.Equals(other.Creator) &&
		p.Medias.Equals(other.Medias) &&
		ArePollDataEquals(p.PollData, other.PollData)
}

// tagsSplitter returns true if the current rune is a tag ending
// Tags MUST end with whitespace, '.' ',' '!' or ')'
func tagsSplitter(c rune) bool {
	if unicode.IsSpace(c) {
		return true
	}
	switch c {
	case '.', ',', '!', ')':
		return true
	}
	return false
}

// getTags matches tags and returns them as an array of strings
//
// The hashtag itself is NOT included as part of the tag string
//
// The function should match the javascript regex: '/([^\S]|^)#([^\s#.,!)]+)(?![^\s.,!)])/g'.
// Since golang re2 engine does not have positive lookahead, the end of the tag is matched by splitting the input string.
// The 'tagsSplitter' function defines the end of a tag, and the 'matchTags' regex has a requirement that it must match the end of a string.
func getTags(s string) []string {
	res := make([]string, 0)
	fields := strings.FieldsFunc(s, tagsSplitter)
	for _, v := range fields {
		sub := HashtagRegEx.FindStringSubmatch(v)
		if len(sub) > 1 {
			res = append(res, sub[1])
		}
	}
	return res
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// GetPostHashtags returns all the post's hashtags without duplicates
func (p Post) GetPostHashtags() []string {
	hashtags := getTags(p.Message)

	uniqueHashtags := commons.Unique(hashtags)
	withoutHashtag := make([]string, len(uniqueHashtags))

	for index, hashtag := range uniqueHashtags {
		trimmed := strings.TrimLeft(strings.TrimSpace(hashtag), "#")
		if !isNumeric(trimmed) {
			withoutHashtag[index] = trimmed
		} else {
			withoutHashtag = []string{}
		}
	}

	if len(withoutHashtag) == 0 {
		return []string{}
	}

	return withoutHashtag
}

// -------------
// --- Posts
// -------------

// Posts represents a slice of Post objects
type Posts []Post

// String implements stringer interface
func (p Posts) String() string {
	out := "ID - [Creator] Message\n"
	for _, post := range p {
		out += fmt.Sprintf("%d - [%s] %s\n",
			post.PostID, post.Creator, post.Message)
	}
	return strings.TrimSpace(out)
}

// Len implements sort.Interface
func (p Posts) Len() int {
	return len(p)
}

// Swap implements sort.Interface
func (p Posts) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less implements sort.Interface
func (p Posts) Less(i, j int) bool {
	return p[i].PostID < p[j].PostID
}
