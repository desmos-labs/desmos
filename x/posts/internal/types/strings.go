package types

// Strings represents a slice of strings.
type Strings []string

// AppendIfMissing returns a new slice containing the given element.
// If the element was already present, it won't be appended and false will be returned.
// If it wasn't present, the element will be appended to the list and true will be returned.
func (elements Strings) AppendIfMissing(element string) (Strings, bool) {
	if elements.Contains(element) {
		return elements, false
	}
	return append(elements, element), true
}

// Contains returns true iff the given element is present inside the elements slice
func (elements Strings) Contains(element string) bool {
	for _, ele := range elements {
		if ele == element {
			return true
		}
	}
	return false
}

// Equals returns true if elements contain the same data of the other slice, in the same order.
func (elements Strings) Equals(other Strings) bool {
	if len(elements) != len(other) {
		return false
	}

	for index, element := range elements {
		if element != other[index] {
			return false
		}
	}

	return true
}
