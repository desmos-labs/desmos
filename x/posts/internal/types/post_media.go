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
	MimeType string `json:"mime_type"`
}

var rEx = regexp.MustCompile(
	`^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$`)

func NewPostMedia(uri, mimeType string) PostMedia {
	return PostMedia{
		URI:      uri,
		MimeType: mimeType,
	}
}

// Validate implements validator
func (pm PostMedia) Validate() error {
	if err := ValidateURI(pm.URI); err != nil {
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

// ValidateURI checks if the given uri string is well-formed according to the regExp and return and error otherwise
func ValidateURI(uri string) error {
	if !rEx.MatchString(uri) {
		return fmt.Errorf("invalid uri provided")
	}

	return nil
}

// ---------------
// --- PostMedias
// ---------------

type PostMedias []PostMedia

// NewPostMedias creates a new PostMedias object starting from the given medias
func NewPostMedias(medias ...PostMedia) PostMedias {
	return medias
}

// String implements fmt.Stringer
func (pms PostMedias) String() string {
	out := "URI [Mime-Type]\n"
	for _, media := range pms {
		out += fmt.Sprintf("[%s] %s \n", media.URI, media.MimeType)
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
// It returns a new slice of PostMedias containing such otherMedia.
func (pms PostMedias) AppendIfMissing(otherMedia PostMedia) PostMedias {
	for _, media := range pms {
		if media.Equals(otherMedia) {
			return pms
		}
	}
	return append(pms, otherMedia)
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
