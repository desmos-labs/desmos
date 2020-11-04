package keeper

// NewWrappedUInt build a new WrappedUint value from the given uint64 value
func NewWrappedUInt(value uint64) WrappedUInt {
	return WrappedUInt{Value: value}
}

// Add returns a new WrappedUInt wrapping the sum between w and value
func (w WrappedUInt) Add(value uint64) WrappedUInt {
	return WrappedUInt{Value: w.Value + value}
}

// ___________________________________________________________________________________________________________________

// AppendIfMissing appends the given id to the IDs wrapped inside ids.
// It returns the result to the append and either true/false if the value was appended or not.
func (ids CommentIDs) AppendIfMissing(id string) (newIDs CommentIDs, appended bool) {
	for _, existing := range ids.Ids {
		if id == existing {
			return ids, false
		}
	}
	return CommentIDs{Ids: append(ids.Ids, id)}, true
}
