package v080

import "regexp"

var (
	uriRegEx = regexp.MustCompile(`^(?:http(s)?://)[\w.-]+(?:\.[\w.-]+)+[\w\-._~:/?#[\]@!$&'()*+,;=.]+$`)
)

// IsValidURI tells whether the given value is a valid URI or not
func IsValidURI(value string) bool {
	return uriRegEx.MatchString(value)
}
