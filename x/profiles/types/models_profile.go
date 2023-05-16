package types

import (
	"fmt"

	"github.com/desmos-labs/desmos/v5/x/commons"
)

// NewPictures is a constructor function for Pictures
func NewPictures(profile, cover string) Pictures {
	return Pictures{
		Profile: profile,
		Cover:   cover,
	}
}

// Validate check the validity of the Pictures
func (pic Pictures) Validate() error {
	if pic.Profile != "" {
		valid := commons.IsURIValid(pic.Profile)
		if !valid {
			return fmt.Errorf("invalid profile picture uri provided")
		}
	}

	if pic.Cover != "" {
		valid := commons.IsURIValid(pic.Cover)
		if !valid {
			return fmt.Errorf("invalid profile cover uri provided")
		}
	}

	return nil
}
