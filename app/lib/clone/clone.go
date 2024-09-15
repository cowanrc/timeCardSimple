package clone

// Pointer returns nil if v is nil.
// Otherwise, it create a new pointer to T, assigns the value *v, and returns
// the new pointer.
func Pointer[T any](v *T) *T {
	if v == nil {
		return nil
	}
	result := new(T)
	*result = *v
	return result
}

// Slice returns a deep copy of the
// provided slice.
//
// The function will return nil if
// nil is provided.
func Slice[T any](v []T) []T {
	if v == nil {
		return nil
	}

	vCopy := make([]T, len(v))
	copy(vCopy, v)

	return vCopy
}
