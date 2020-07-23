package common

import (
	"fmt"
	"strings"

	"github.com/desmos-labs/desmos/x/commons"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Attachment
// ---------------

// Attachment contains the information representing any type of file provided with a post.
// This file can be an image or a multimedia file (vocals, video, documents, etc.).
type Attachment struct {
	URI      string           `json:"uri" yaml:"uri"`
	MimeType string           `json:"mime_type" yaml:"mime_type"`
	Tags     []sdk.AccAddress `json:"tags,omitempty" yaml:"tags,omitempty"`
}

func NewAttachment(uri, mimeType string, tags []sdk.AccAddress) Attachment {
	return Attachment{
		URI:      uri,
		MimeType: mimeType,
		Tags:     tags,
	}
}

// Validate implements validator
func (att Attachment) Validate() error {
	if !commons.IsURIValid(att.URI) {
		return fmt.Errorf("invalid uri provided")
	}

	if len(strings.TrimSpace(att.MimeType)) == 0 {
		return fmt.Errorf("mime type must be specified and cannot be empty")
	}

	for _, address := range att.Tags {
		if address.Empty() {
			return fmt.Errorf("invalid empty tag address: %s", address)
		}
	}

	return nil
}

// Equals allows to check whether the contents of att are the same of other
func (att Attachment) Equals(other Attachment) bool {
	if len(att.Tags) != len(other.Tags) {
		return false
	}

	for index, address := range att.Tags {
		if !address.Equals(other.Tags[index]) {
			return false
		}
	}

	return att.URI == other.URI && att.MimeType == other.MimeType
}

// ---------------
// --- Attachments
// ---------------

type Attachments []Attachment

// NewAttachments creates a new Attachments object starting from the given attachments
func NewAttachments(attachments ...Attachment) Attachments {
	return attachments
}

func formatTagsOutput(tags []sdk.AccAddress) (outputTags string) {
	for _, addr := range tags {
		outputTags += fmt.Sprintf("%s,\n", addr.String())
	}
	return outputTags
}

// String implements fmt.Stringer
func (atts Attachments) String() string {
	out := "[URI] [Mime-Type] [Tags]\n"
	for _, att := range atts {
		tags := formatTagsOutput(att.Tags)
		out += fmt.Sprintf("[%s] [%s] [%s] \n", att.URI, att.MimeType, tags)
	}

	return strings.TrimSpace(out)
}

// Equals returns true iff the atts slice contains the same
// data in the same order of the other slice
func (atts Attachments) Equals(other Attachments) bool {
	if len(atts) != len(other) {
		return false
	}

	for index, attachment := range atts {
		if !attachment.Equals(other[index]) {
			return false
		}
	}

	return true
}

// AppendIfMissing appends the given otherAttachment to the atts slice if it does not exist inside it yet.
// It returns a new slice of Attachments containing such otherAttachment.
func (atts Attachments) AppendIfMissing(otherAttachment Attachment) Attachments {
	for _, att := range atts {
		if att.Equals(otherAttachment) {
			return atts
		}
	}
	return append(atts, otherAttachment)
}

// Validate implements validator
func (atts Attachments) Validate() error {
	for _, att := range atts {
		if err := att.Validate(); err != nil {
			return err
		}
	}
	return nil
}
