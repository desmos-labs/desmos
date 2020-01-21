package types

import (
	"fmt"
	"regexp"
	"strings"
)

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
	out := "Media - "

	out += fmt.Sprintf(" URI - [%s] ; Mime-Type - [%s] \n", pm.URI, pm.MimeType)

	return strings.TrimSpace(out)
}

// Validate implements validator
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

// Equals allows to check whether the contents of pm are the same of other
func (pm PostMedia) Equals(other PostMedia) bool {
	return pm.URI == other.URI && pm.MimeType == other.MimeType
}

// ParseURI checks if the given uri string is well-formed according to the regExp and return and error otherwise
func ParseURI(uri string) error {
	rEx := regexp.MustCompile(
		`^(?:https:\/\/)[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:\/?#[\]@!\$&'\(\)\*\+,;=.]+$`)

	if !rEx.MatchString(uri) {
		return fmt.Errorf("invalid uri provided")
	}

	return nil
}

// ---------------
// --- PostMedias
// ---------------

type PostMedias []PostMedia

// String implements stringer interface
func (pms PostMedias) String() string {
	out := "medias - [URI] [Mime-Type]\n"
	for _, post := range pms {
		out += fmt.Sprintf("[%s] %s \n", post.URI, post.MimeType)
	}

	return strings.TrimSpace(out)
}

// Equals returns true iff the pms slice contains the same
// data in the same order of the other slice
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

// AppendIfMissing appends the given otherMedia to the pms slice if it does not exist inside it yet.
// It returns a new slice of PostMedias containing such otherMedia and a boolean indicating whether or not the original
// slice has been modified.
func (pms PostMedias) AppendIfMissing(otherMedia PostMedia) (PostMedias, bool) {
	for _, media := range pms {
		if media.Equals(otherMedia) {
			return pms, false
		}
	}
	return append(pms, otherMedia), true
}

// Validate implements validator
func (pms PostMedias) Validate() error {
	for _, media := range pms {
		if err := media.Validate(); err != nil {
			return err
		}
	}
	return nil
}
