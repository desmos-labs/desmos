package types

import "regexp"

var (
	subspaceRegEx = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
)

// IsValidSubspace tells whether the given value is a valid subspace or not
func IsValidSubspace(value string) bool {
	return subspaceRegEx.MatchString(value)
}
