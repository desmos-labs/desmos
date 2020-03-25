package types

// Pictures contains the data of a user profile's related pictures
type Pictures struct {
	Profile string `json:"profile,omitempty"`
	Cover   string `json:"cover,omitempty"`
}

// Equals allows to check whether the contents of pic are the same of otherPic
func (pic Pictures) Equals(otherPic *Pictures) bool {
	return pic.Profile == otherPic.Profile &&
		pic.Cover == otherPic.Cover
}
