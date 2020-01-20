package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MediaPost struct {
	TextPost
	Medias PostMedias `json:"medias"`
}

func NewMediaPost(post TextPost, medias PostMedias) MediaPost {
	return MediaPost{
		TextPost: post,
		Medias:   medias,
	}
}

// String implements fmt.Stringer
func (mp MediaPost) String() string {
	bytes, err := json.Marshal(&mp)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// GetID implements Post
func (mp MediaPost) GetID() PostID {
	return mp.PostID
}

// GetParentID implements Post
func (mp MediaPost) GetParentID() PostID {
	return mp.ParentID
}

// SetMessage implements Post
func (mp MediaPost) SetMessage(message string) Post {
	mp.Message = message
	return mp
}

// GetMessage implements Post
func (mp MediaPost) GetMessage() string {
	return mp.Message
}

// CreationTime implements Post
func (mp MediaPost) CreationTime() time.Time {
	return mp.Created
}

// SetEditTime implements Post
func (mp MediaPost) SetEditTime(time time.Time) Post {
	mp.LastEdited = time
	return mp
}

// GetEditTime implements Post
func (mp MediaPost) GetEditTime() time.Time {
	return mp.LastEdited
}

// CanComment implements Post
func (mp MediaPost) CanComment() bool {
	return mp.AllowsComments
}

// GetSubspace implements Post
func (mp MediaPost) GetSubspace() string {
	return mp.Subspace
}

// GetOptionalData implements Post
func (mp MediaPost) GetOptionalData() map[string]string {
	return mp.OptionalData
}

// Owner implements Post
func (mp MediaPost) Owner() sdk.AccAddress {
	return mp.Creator
}

// Validate implements Post Validate
func (mp MediaPost) Validate() error {
	if err := mp.TextPost.Validate(); err != nil {
		return err
	}

	for _, media := range mp.Medias {
		if err := media.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Equal implements Post
func (mp MediaPost) Equals(other Post) bool {
	// Cast and delegate
	if otherMp, ok := other.(MediaPost); ok {
		return checkMediaPostEquals(mp, otherMp)
	}

	return false
}

// checkMediaPostEquals checks if two MediaPost are equal
func checkMediaPostEquals(first MediaPost, second MediaPost) bool {
	return first.TextPost.Equals(second.TextPost) && first.Medias.Equals(second.Medias)
}

// MarshalJSON implements Marshaler
func (mp MediaPost) MarshalJSON() ([]byte, error) {
	type temp MediaPost
	return json.Marshal(temp(mp))
}

// UnmarshalJSON implements Unmarshaler
func (mp *MediaPost) UnmarshalJSON(data []byte) error {
	type mediaPostJSON MediaPost
	var temp mediaPostJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*mp = MediaPost(temp)
	return nil
}

// -------------
// --- MediaPosts
// -------------

// TextPosts represents a slice of TextPost objects
type MediaPosts []MediaPost

// checkPostsEqual returns true iff the p slice contains the same
// data in the same order of the other slice
func (mps MediaPosts) Equals(other MediaPosts) bool {
	if len(mps) != len(other) {
		return false
	}

	for index, post := range mps {
		if !post.Equals(other[index]) {
			return false
		}
	}

	return true
}

// String implements stringer interface
func (mps MediaPosts) String() string {
	bytes, err := json.Marshal(&mps)
	if err != nil {
		panic(err)
	}

	return string(bytes)
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

// ---------------
// --- PostMedia
// ---------------

type PostMedia struct {
	URI      string `json:"uri"`
	Provider string `json:"provider"`
	MimeType string `json:"mime_Type"`
}

func NewPostMedia(uri, provider, mimeType string) PostMedia {
	return PostMedia{
		URI:      uri,
		Provider: provider,
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
	if len(strings.TrimSpace(pm.Provider)) == 0 {
		return fmt.Errorf("media provider must be specified and cannot be empty")
	}

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
	return pm.URI == other.URI && pm.MimeType == other.MimeType && pm.Provider == other.Provider
}

func ParseURI(uri string) error {
	rEx := regexp.MustCompile(
		`^(?:https:\/\/)[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:\/?#[\]@!\$&'\(\)\*\+,;=.]+$`)

	if !rEx.MatchString(uri) {
		return fmt.Errorf("invalid uri provided")
	}

	return nil
}
