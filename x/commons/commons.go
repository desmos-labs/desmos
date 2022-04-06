package commons

import (
	"net/url"
)

// Unique returns the given input slice without any duplicated value inside it
func Unique(input []string) []string {
	unique := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, val := range input {
		if _, ok := m[val]; !ok {
			m[val] = true
			unique = append(unique, val)
		}
	}
	return unique
}

// IsURIValid tells whether the given URI is valid or not
func IsURIValid(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}

	u, err := url.Parse(uri)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}
