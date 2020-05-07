package common

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- PostMedia
// ---------------

type PostMedia struct {
	URI      string           `json:"uri" yaml:"uri"`
	MimeType string           `json:"mime_type" yaml:"mime_type"`
	Tags     []sdk.AccAddress `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func NewPostMedia(uri, mimeType string, tags []sdk.AccAddress) PostMedia {
	return PostMedia{
		URI:      uri,
		MimeType: mimeType,
		Tags:     tags,
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

	for _, address := range pm.Tags {
		if address.Empty() {
			return fmt.Errorf("invalid empty tag address: %s", address)
		}
	}

	return nil
}

// Equals allows to check whether the contents of pm are the same of other
func (pm PostMedia) Equals(other PostMedia) bool {
	if len(pm.Tags) != len(other.Tags) {
		return false
	}

	for index, address := range pm.Tags {
		if !address.Equals(other.Tags[index]) {
			return false
		}
	}

	return pm.URI == other.URI && pm.MimeType == other.MimeType
}

// ValidateURI checks if the given uri string is well-formed according to the regExp and return and error otherwise
func ValidateURI(uri string) error {
	if !URIRegEx.MatchString(uri) {
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

func formatTagsOutput(tags []sdk.AccAddress) (outputTags string) {
	for _, addr := range tags {
		outputTags += fmt.Sprintf("%s,\n", addr.String())
	}
	return outputTags
}

// String implements fmt.Stringer
func (pms PostMedias) String() string {
	out := "[URI] [Mime-Type] [Tags]\n"
	for _, media := range pms {
		tags := formatTagsOutput(media.Tags)
		out += fmt.Sprintf("[%s] [%s] [%s] \n", media.URI, media.MimeType, tags)
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
