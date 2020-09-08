package common

// NewStrPtr allows to build a new string pointer starting
// from a simple string value to setup tests more easily
func NewStrPtr(value string) *string {
	return &value
}
