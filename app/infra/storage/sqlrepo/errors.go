package sqlrepo

import (
	"database/sql"
	"errors"
	"fmt"
)

const packagePrefix = "sqlrepo"

var ErrNoRows = sql.ErrNoRows

var ErrInvalidTxType = errors.New("sqlrepov: invalid Tx type")

type ConstraintViolationError struct {
	Err        error
	Constraint string
	Column     string
	Table      string
	Value      []any
	Detail     string
}

func (e *ConstraintViolationError) Error() string {
	return fmt.Sprintf("%s: ConstraintViolationError: %q %q", packagePrefix, e.Constraint, e.Detail)
}

func (e *ConstraintViolationError) Unwrap() error {
	return e.Err
}

func (e *ConstraintViolationError) SyCode() string {
	return "ConstraintViolation"
}
