package types

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type MediaPost struct {
	Post   `json:"post"`
	Medias []PostMedia `json:"medias"`
}

func (mp MediaPost) Validate() error {
	if err := mp.Post.Validate(); err != nil {
		return err
	}

	for _, post := range mp.Medias {
		if err := post.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (mp MediaPost) Equals(other MediaPost) bool {
	if !mp.Post.Equals(other.Post) {
		return false
	}

	if len(mp.Medias) != len(other.Medias) {
		return false
	}

	for index, media := range mp.Medias {
		if media != other.Medias[index] {
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
