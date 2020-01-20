package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

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

// checkPostsEqual compares two PostID instances
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

// checkPostsEqual returns true iff the ids slice and the other
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
	Medias         PostMedias     `json:"medias"`
}

func NewPost(id, parentID PostID, message string, allowsComments bool, subspace string, optionalData map[string]string,
	created time.Time, creator sdk.AccAddress, medias PostMedias) Post {
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
		Medias:         medias,
	}
}

// String implements fmt.Stringer
func (p Post) String() string {
	bytes, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate implements Post Validate
func (p Post) Validate() error {
	if !p.PostID.Valid() {
		return fmt.Errorf("invalid post id: %s", p.PostID)
	}

	if p.Creator == nil {
		return fmt.Errorf("invalid post owner: %s", p.Creator)
	}

	if len(strings.TrimSpace(p.Message)) == 0 {
		return fmt.Errorf("post message must be non empty and non blank")
	}

	if len(p.Message) > MaxPostMessageLength {
		return fmt.Errorf("post message cannot be longer than %d characters", MaxPostMessageLength)
	}

	if len(strings.TrimSpace(p.Subspace)) == 0 {
		return fmt.Errorf("post subspace must be non empty and non blank")
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

	err := p.Medias.Validate()
	if err != nil {
		return err
	}

	return nil
}

// Equals returns true if the two post are equal false otherwise
func (p Post) Equals(second Post) bool {
	equalsOptionalData := len(p.OptionalData) == len(second.OptionalData)
	if equalsOptionalData {
		for key := range p.OptionalData {
			equalsOptionalData = equalsOptionalData && p.OptionalData[key] == second.OptionalData[key]
		}
	}

	return p.PostID.Equals(second.PostID) &&
		p.ParentID.Equals(second.ParentID) &&
		p.Message == second.Message &&
		p.Created.Equal(second.Created) &&
		p.LastEdited.Equal(second.LastEdited) &&
		p.AllowsComments == second.AllowsComments &&
		p.Subspace == second.Subspace &&
		equalsOptionalData &&
		p.Creator.Equals(second.Creator) &&
		p.Medias.Equals(second.Medias)
}

// -------------
// --- Posts
// -------------

// Posts represents a slice of Post objects
type Posts []Post

// checkPostsEqual returns true iff the p slice contains the same
// data in the same order of the other slice
func (p Posts) Equals(other Posts) bool {
	if len(p) != len(other) {
		return false
	}

	for index, post := range p {
		if !post.Equals(other[index]) {
			return false
		}
	}

	return true
}

// String implements stringer interface
func (p Posts) String() string {
	out := "ID - [Creator] Message\n"
	for _, post := range p {
		out += fmt.Sprintf("%d - [%s] %s\n",
			post.PostID, post.Creator, post.Message)
	}
	return strings.TrimSpace(out)
}

// ---------------
// --- PostMedias
// ---------------

type PostMedias []PostMedia

func (pms PostMedias) String() string {
	bytes, err := json.Marshal(&pms)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
func (pms PostMedias) Equals(other PostMedias) bool {
	if len(pms) != len(other) {
		return false
	}

	for index, postMedia := range pms {
		if !postMedia.Equals(other[index]) {
			return false
		}
	}

	return true
}

func (pms PostMedias) AppendIfMissing(otherMedia PostMedia) (PostMedias, bool) {
	for _, media := range pms {
		if media.Equals(otherMedia) {
			return pms, false
		}
	}
	return append(pms, otherMedia), true
}

func (pms PostMedias) Validate() error {
	for _, media := range pms {
		if err := media.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// ---------------
// --- PostMedia
// ---------------

type PostMedia struct {
	URI      string `json:"uri"`
	MimeType string `json:"mime_Type"`
}

func NewPostMedia(uri, mimeType string) PostMedia {
	return PostMedia{
		URI:      uri,
		MimeType: mimeType,
	}
}

// String implements fmt.Stringer
func (pm PostMedia) String() string {
	bytes, err := json.Marshal(&pm)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func (pm PostMedia) Validate() error {
	if len(strings.TrimSpace(pm.URI)) == 0 {
		return fmt.Errorf("uri must be specified and cannot be empty")
	}

	if err := ParseURI(pm.URI); err != nil {
		return err
	}

	if len(strings.TrimSpace(pm.MimeType)) == 0 {
		return fmt.Errorf("mime type must be specified and cannot be empty")
	}

	return nil
}

func (pm PostMedia) Equals(other PostMedia) bool {
	return pm.URI == other.URI && pm.MimeType == other.MimeType
}

func ParseURI(uri string) error {
	rEx := regexp.MustCompile(
		`^(?:https:\/\/)[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:\/?#[\]@!\$&'\(\)\*\+,;=.]+$`)

	if !rEx.MatchString(uri) {
		return fmt.Errorf("invalid uri provided")
	}

	return nil
}
