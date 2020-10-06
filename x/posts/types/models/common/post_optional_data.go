package common

import "fmt"

// OptionalData represents a Posts' optional data and allows for custom
// Amino and JSON serialization and deserialization.
type OptionalData struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
}

// NewOptionalData returns a new OptionalData object
func NewOptionalData(key, value string) OptionalData {
	return OptionalData{
		Key:   key,
		Value: value,
	}
}

// Equals allows to check if the two optionalData objects are equal or not
func (op OptionalData) Equals(other OptionalData) bool {
	return op.Key == other.Key && op.Value == other.Value
}

// String implement fmt.Stringer
func (op OptionalData) String() string {
	out := "[Key] [Value]\n"
	out += fmt.Sprintf("[%s] [%s]", op.Key, op.Value)
	return out
}
