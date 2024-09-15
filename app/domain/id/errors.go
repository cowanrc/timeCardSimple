package id

import (
	"errors"
	"fmt"
)

// ScanIDError is an error denoting that the src interface{} for scanning an ID
// is not supported.
type ScanIDError struct {
	src interface{}
}

// Error is the error implementation.
func (e *ScanIDError) Error() string {
	return fmt.Sprintf("id: unable to scan type %T into ID", e.src)
}

// ErrCouldNotCreateID is returned from New if an error occurs creating an ID.
var ErrCouldNotCreateID = errors.New("id: could not create a new ID")

// ParsingIDError is a wrapper error denoting an invalid string was presented to
// parsing an ID and thus failed.
type ParsingIDError struct {
	Err error
}

// Error is the error implementation.
func (e *ParsingIDError) Error() string {
	return fmt.Sprintf("id: parsing ID error: %v", e.Err)
}

// ErrWrongVersion indicates a version error for a parsed id.
var ErrWrongVersion = errors.New("id: wrong version for parsed ID")

// ErrEmpty indicates an ID is empty.
var ErrEmpty = errors.New("id: empty id")
