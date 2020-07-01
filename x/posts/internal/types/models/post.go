package models

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/commons"
)

// ---------------
// --- Post id
// ---------------

// PostID represents a unique post id
type PostID string

// ComputeID returns a sha256 hash of the given data concatenated together
// nolint: interfacer
func ComputeID(creationDate time.Time, creator sdk.AccAddress, subspace string) PostID {
	hash := sha256.Sum256([]byte(creationDate.String() + creator.String() + subspace))
	return PostID(hex.EncodeToString(hash[:]))
}

// Valid tells if the id can be used safely
func (id PostID) Valid() bool {
	return strings.TrimSpace(id.String()) != "" && Sha256RegEx.MatchString(id.String())
}

// String implements fmt.Stringer
func (id PostID) String() string {
	return string(id)
}

// Equals compares two PostID instances
func (id PostID) Equals(other PostID) bool {
	return id == other
}

// ParsePostID takes the given value and returns a PostID from it.
// If the given value cannot be parse, an error is returned instead.
func ParsePostID(value string) (PostID, error) {
	if !Sha256RegEx.MatchString(value) {
		return "", fmt.Errorf("%s is not a valid post id", value)
	}
	return PostID(value), nil
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
		if !id.Equals(other[index]) {
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

// String implements fmt.Stringer
func (ids PostIDs) String() string {
	var stringIDs = make([]string, len(ids))
	for index, id := range ids {
		stringIDs[index] = id.String()
	}

	out := strings.Join(stringIDs, ", ")
	return fmt.Sprintf("[%s]", strings.TrimSpace(out))
}

// ---------------
// --- Post
// ---------------

// Post is a struct of a post
type Post struct {
	PostID         PostID         `json:"id" yaml:"id" `                                          // Unique id
	ParentID       PostID         `json:"parent_id" yaml:"parent_id"`                             // Post of which this one is a comment
	Message        string         `json:"message" yaml:"message"`                                 // Message contained inside the post
	Created        time.Time      `json:"created" yaml:"created"`                                 // RFC3339 date at which the post has been created
	LastEdited     time.Time      `json:"last_edited" yaml:"last_edited"`                         // RFC3339 date at which the post has been edited the last time
	AllowsComments bool           `json:"allows_comments" yaml:"allows_comments"`                 // Tells if users can reference this PostID as the parent
	Subspace       string         `json:"subspace" yaml:"subspace"`                               // Identifies the application that has posted the message
	OptionalData   OptionalData   `json:"optional_data,omitempty" yaml:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        sdk.AccAddress `json:"creator" yaml:"creator"`                                 // Creator of the Post
	Medias         PostMedias     `json:"medias,omitempty" yaml:"medias,omitempty"`               // Contains all the medias that are shared with the post
	PollData       *PollData      `json:"poll_data,omitempty" yaml:"poll_data"`                   // Contains the poll details, if existing
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
		return fmt.Errorf("invalid postID: %s", p.PostID)
	}

	if p.Creator == nil {
		return fmt.Errorf("invalid post owner: %s", p.Creator)
	}

	if len(strings.TrimSpace(p.Message)) == 0 && len(p.Medias) == 0 {
		return fmt.Errorf("post message or medias required, they cannot be both empty")
	}

	if !Sha256RegEx.MatchString(p.Subspace) {
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

	if err := p.Medias.Validate(); err != nil {
		return err
	}

	if p.PollData != nil {
		if err := p.PollData.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Equals allows to check whether the contents of p are the same of other
func (p Post) Equals(other Post) bool {
	return p.PostID.Equals(other.PostID) && p.ContentsEquals(other)
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
		out += fmt.Sprintf("%s - [%s] %s\n",
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
	return p[i].Created.Before(p[j].Created)
}
