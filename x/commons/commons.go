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
