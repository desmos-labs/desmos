package models

import (
	"fmt"
)

// Pictures contains the data of a user profile's related pictures
type Pictures struct {
	Profile *string `json:"profile,omitempty" yaml:"profile,omitempty"`
	Cover   *string `json:"cover,omitempty" yaml:"cover,omitempty"`
}

// NewPictures is a constructor function for Pictures
func NewPictures(profile, cover *string) *Pictures {
	if profile == nil && cover == nil {
		return nil
	}
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

// Validate check the validity of the Pictures
func (pic Pictures) Validate() error {

	if pic.Profile != nil {
		if valid := URIRegEx.MatchString(*pic.Profile); !valid {
			return fmt.Errorf("invalid profile picture uri provided")
		}
	}

	if pic.Cover != nil {
		if valid := URIRegEx.MatchString(*pic.Cover); !valid {
			return fmt.Errorf("invalid profile cover uri provided")
		}
	}
	return nil
}
