package types

import (
	"fmt"
)

// Pictures contains the data of a user profile's related pictures
type Pictures struct {
	Profile *string `json:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty"`
}

// NewPictures is a constructor function for Pictures
func NewPictures(profile, cover *string) *Pictures {
	return &Pictures{
		Profile: profile,
		Cover:   cover,
	}
}

// Equals allows to check whether the contents of pic are the same of otherPic
func (pic Pictures) Equals(otherPic *Pictures) bool {
	return pic.Profile == otherPic.Profile &&
		pic.Cover == otherPic.Cover
}

// ValidateURI checks if the given uri string is well-formed according to the regExp and return and error otherwise
func ValidateURI(uri string) error {
	if !URIRegEx.MatchString(uri) {
		return fmt.Errorf("invalid uri provided")
	}

	return nil
}

// Validate check the validity of the Pictures
func (pic Pictures) Validate() error {

	if pic.Profile != nil {
		if err := ValidateURI(*pic.Profile); err != nil {
			return fmt.Errorf("invalid profile picture uri provided")
		}
	}

	if pic.Cover != nil {
		if err := ValidateURI(*pic.Cover); err != nil {
			return fmt.Errorf("invalid profile cover uri provided")
		}
	}
	return nil
}
