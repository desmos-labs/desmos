package commons

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

// StringPtrsEqual returns true iff the given first and second string pointers
// are equals. This is true only if both are nil, or if they point to the same value.
func StringPtrsEqual(first, second *string) bool {
	if first == nil || second == nil {
		return first == second
	}

	return *first == *second
}
