package v030

import "regexp"

var (
	subspaceRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")
)

func IsValidSubspace(value string) bool {
	return subspaceRegEx.MatchString(value)
}
