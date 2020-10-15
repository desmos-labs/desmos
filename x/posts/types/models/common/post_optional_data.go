package common

import "fmt"

type OptionalData []OptionalDataEntry

// OptionalDataEntry represents a Posts' optional data entry and allows for custom
// Amino and JSON serialization and deserialization.
type OptionalDataEntry struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

// NewOptionalData returns a new OptionalDataEntry object
func NewOptionalData(key, value string) OptionalDataEntry {
	return OptionalDataEntry{
		Key:   key,
		Value: value,
	}
}

// Equals allows to check if the two optionalData objects are equal or not
func (op OptionalDataEntry) Equals(other OptionalDataEntry) bool {
	return op.Key == other.Key && op.Value == other.Value
}

// String implement fmt.Stringer
func (op OptionalDataEntry) String() string {
	out := "[Key] [Value]\n"
	out += fmt.Sprintf("[%s] [%s]", op.Key, op.Value)
	return out
}
