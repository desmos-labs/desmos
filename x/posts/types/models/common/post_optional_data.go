package common

import (
	"encoding/json"
	"sort"
)

// OptionalData represents a Posts' optional data and allows for custom
// Amino and JSON serialization and deserialization.
type OptionalData map[string]string

// KeyValue is a simple key/value representation of one field of a OptionalData.
type KeyValue struct {
	Key   string
	Value string
}

// MarshalAmino transforms the OptionalData to an array of key/value.
func (m OptionalData) MarshalAmino() ([]KeyValue, error) {
	fieldKeys := make([]string, len(m))
	i := 0
	for key := range m {
		fieldKeys[i] = key
		i++
	}

	sort.Stable(sort.StringSlice(fieldKeys))

	p := make([]KeyValue, len(m))
	for i, key := range fieldKeys {
		p[i] = KeyValue{
			Key:   key,
			Value: m[key],
		}
	}

	return p, nil
}

// UnmarshalAmino transforms the key/value array to a OptionalData.
func (m *OptionalData) UnmarshalAmino(keyValues []KeyValue) error {
	tempMap := make(map[string]string, len(keyValues))
	for _, p := range keyValues {
		tempMap[p.Key] = p.Value
	}

	*m = tempMap

	return nil
}

// MarshalJSON implements encode.Marshaler
func (m OptionalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string(m))
}

// UnmarshalJSON implements decode.Unmarshaler
func (m *OptionalData) UnmarshalJSON(data []byte) error {
	var value map[string]string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = value
	return nil
}
