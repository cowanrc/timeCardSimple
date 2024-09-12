package sqlrepo

import (
	"fmt"

	"github.com/lib/pq"
)

const (
	PostgresCodeUniqueViolation   = "23505"
	PostgresCodeRestrictViolation = "23503"
)

var _ Dialect = Postgres{}

type Postgres struct{}

func (p Postgres) NormalizeSQLError(err error) error {
	errPQ, ok := err.(*pq.Error)
	if ok {
		switch errPQ.Code {
		case PostgresCodeUniqueViolation, PostgresCodeRestrictViolation:
			err = &ConstraintViolationError{
				Err:        errPQ,
				Constraint: errPQ.Constraint,
				Column:     errPQ.Column,
				Table:      errPQ.Table,
				Value:      []any{errPQ.Detail},
				Detail:     errPQ.Detail,
			}
		}
	}

	return err
}

func (p Postgres) Placeholder(index int) string {
	return fmt.Sprintf("$%d", index+1)
}
