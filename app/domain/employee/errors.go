package employee

import "fmt"

const packagePrefix = "employee"

var (
	ErrUnsetID        = fmt.Errorf("%s: id is empty", packagePrefix)
	ErrUnsetFirstName = fmt.Errorf("%s: first name is empty", packagePrefix)
	ErrUnsetLastName  = fmt.Errorf("%s: last name is empty", packagePrefix)
	ErrUnsetEmail     = fmt.Errorf("%s: email is empty", packagePrefix)
	ErrUnsetCreatedAt = fmt.Errorf("%s: created at is zero", packagePrefix)
	ErrUnsetUpdatedAt = fmt.Errorf("%s: updated at is zero", packagePrefix)
)
