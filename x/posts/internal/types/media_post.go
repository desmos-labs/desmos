package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MediaPost struct {
	TextPost `json:"post"`
	Medias   PostMedias `json:"medias"`
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

// GetID implements Post GetID
func (mp MediaPost) GetID() PostID {
	return mp.PostID
}

// GetParentID implements Post GetParentID
func (mp MediaPost) GetParentID() PostID {
	return mp.ParentID
}

func (mp MediaPost) SetMessage(message string) Post {
	mp.Message = message
	return mp
}

func (mp MediaPost) GetMessage() string {
	return mp.Message
}

func (mp MediaPost) CreationTime() sdk.Int {
	return mp.Created
}

func (mp MediaPost) SetEditTime(time sdk.Int) Post {
	mp.LastEdited = time
	return mp
}

func (mp MediaPost) EditTime() sdk.Int {
	return mp.LastEdited
}

func (mp MediaPost) CanComment() bool {
	return mp.AllowsComments
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

func (mp MediaPost) Equals(other Post) bool {
	// Cast and delegate
	if otherMp, ok := other.(MediaPost); ok {
		return checkMediaPostEquals(mp, otherMp)
	}

	return false
}

// Equals implements Post Equals
func checkMediaPostEquals(first MediaPost, second MediaPost) bool {
	if !first.TextPost.Equals(second.TextPost) {
		return false
	}

	if len(first.Medias) != len(second.Medias) {
		return false
	}

	for index, media := range first.Medias {
		if media != second.Medias[index] {
			return false
		}
	}
	return true
}

// MarshalJSON implements Marshaler
func (mp MediaPost) MarshalJSON() ([]byte, error) {
	type mediaPostJSON MediaPost
	return json.Marshal(mediaPostJSON(mp))
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

// ---------------
// --- PostMedia
// ---------------

type PostMedias []PostMedia

type PostMedia struct {
	Provider string `json:"provider"`
	URI      string `json:"uri"`
	MimeType string `json:"mime_Type"`
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
		return fmt.Errorf("mime type must be specified and cannot be")
	}

	return nil
}

func (pm PostMedia) Equals(other PostMedia) bool {
	return pm.URI == other.URI && pm.MimeType == other.MimeType && pm.Provider == other.Provider
}

//todo test this properly
func ParseURI(uri string) error {
	rEx := regexp.MustCompile(
		`^(?:https:\/\/)[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:\/?#[\]@!\$&'\(\)\*\+,;=.]+$`)

	if !rEx.MatchString(uri) {
		return fmt.Errorf("invalid uri provided")
	}

	return nil
}
