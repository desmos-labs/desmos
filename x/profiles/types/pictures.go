package types

import (
	"fmt"

	"github.com/desmos-labs/desmos/x/commons"
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

// Equals allows to check whether the contents of pic are the same of otherPics
func (pic Pictures) Equals(otherPic *Pictures) bool {
	if otherPic == nil {
		return false
	}

	return commons.StringPtrsEqual(pic.Profile, otherPic.Profile) &&
		commons.StringPtrsEqual(pic.Cover, otherPic.Cover)
}

// Validate check the validity of the Pictures
func (pic Pictures) Validate() error {

	if pic.Profile != nil {
		if valid := commons.IsURIValid(*pic.Profile); !valid {
			return fmt.Errorf("invalid profile picture uri provided")
		}
	}

	if pic.Cover != nil {
		if valid := commons.IsURIValid(*pic.Cover); !valid {
			return fmt.Errorf("invalid profile cover uri provided")
		}
	}
	return nil
}
