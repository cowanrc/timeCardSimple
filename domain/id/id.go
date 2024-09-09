package id

import (
	"bytes"
	"crypto/sha256"
	"encoding"
	"sort"

	"github.com/google/uuid"
)

var _ encoding.TextMarshaler = ID(uuid.UUID([16]byte{}))
var _ encoding.TextUnmarshaler = new(ID)

// ID is the domain type for identifiers throughout the application.
type ID uuid.UUID

// New creates a new ID. It is a version 4 random UUID.
// An error will be non-nil if the random bytes needed could not be read.
func New() (ID, error) {
	id, err := fromUUID(uuid.NewRandom())
	if err != nil {
		err = ErrCouldNotCreateID
	}
	return id, err
}

// ParseString attempts to parse s into an ID. A non-nil error is returned if s is not
// a properly formatted id.
func ParseString(s string) (ID, error) {
	return parseResult(uuid.Parse(s))
}

// ParseBytes attempts to parse b into an ID. A non-nil error is returned if b
// is not a properly formatted id.
func ParseBytes(b []byte) (ID, error) {
	return parseResult(uuid.ParseBytes(b))
}

func parseResult(uuid uuid.UUID, err error) (ID, error) {
	if err == nil {
		if uuid.Version() != 4 {
			err = ErrWrongVersion
		}
	}
	if err != nil {
		err = &ParsingIDError{Err: err}
	}
	return fromUUID(uuid, err)
}

func FromDeterministicHash(value []byte) (ID, error) {
	// THE HASH ALGORITHM CANNOT CHANGE.
	// USING A NEW ALGORITHM WILL REQUIRE A NEW FUNCTION.
	sha256Hash := sha256.New()
	sha256Hash.Write(value)
	uuid, err := uuid.NewRandomFromReader(bytes.NewReader(sha256Hash.Sum(nil)))
	return fromUUID(uuid, err)
}

func fromUUID(uuid uuid.UUID, err error) (ID, error) {
	return ID(uuid), err
}

func (id ID) Validate() error {
	if id.Equal(Empty()) {
		return ErrEmpty
	}
	return nil
}

// Empty returns an empty (and invalid) ID.
// This function is provided strictly for easily testing against IsEmpty().
func Empty() ID {
	return ID(uuid.UUID([16]byte{}))
}

// Equal determines whether id and other represent the same ID.
func (id ID) Equal(other ID) bool {
	return bytes.Equal([]byte(id[:]), []byte(other[:]))
}

// GoString returns the Go string representation of id.
func (id ID) GoString() string {
	return id.String()
}

// String returns the normal string representation of id.
func (id ID) String() string {
	return uuid.UUID(id).String()
}

// MarshalText is the encoding.TextMarshaler implementation.
//
// We use it for compatibility with other marshaling packages
// instead of adding custom marshaling and unmarshaling for each encoding package.
// E.g. encoding/json will use this method.
func (id ID) MarshalText() ([]byte, error) {
	return []byte(id.String()), nil
}

// UnmarshalText is the encoding.TextUnmarshaler implementation.
//
// We use it for compatibility with other marshaling packages
// instead of adding custom marshaling and unmarshaling for each encoding package.
// E.g. encoding/json will use this method.
func (id *ID) UnmarshalText(p []byte) error {
	var err error
	*id, err = ParseBytes(p)
	return err
}

// Less returns if id is "less than" other using bytes.Compare < 0.
func (id ID) Less(other ID) bool {
	return bytes.Compare(id[:], other[:]) < 0
}

// SortSlices sorts ids according to bytes.Compare on the internal byte array of the ids.
func SortSlices(slices ...[]ID) {
	for _, ids := range slices {
		sort.Slice(ids, func(i, j int) bool {
			return ids[i].Less(ids[j])
		})
	}
}

// ContainsSorted determines whether id is present in sortedIDs.
// SortedIDs MUST be sorted for this function to work properly.
func ContainsSorted(sortedIDs []ID, id ID) bool {
	return IndexOfSorted(sortedIDs, id) >= 0
}

// IndexOfSorted returns the index of the provided id in the sortedIDs slice.
func IndexOfSorted(sortedIDs []ID, id ID) int {
	i := sort.Search(len(sortedIDs), func(i int) bool {
		return bytes.Compare(sortedIDs[i][:], id[:]) >= 0
	})
	if i < len(sortedIDs) && sortedIDs[i].Equal(id) {
		return i
	}
	return -1
}

// SlicesEqual returns whether ID slices a and b are equal.
func SlicesEqual(a, b []ID) bool {
	if len(a) != len(b) {
		return false
	}
	for i, id := range a {
		if !id.Equal(b[i]) {
			return false
		}
	}
	return true
}
