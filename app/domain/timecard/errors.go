package timecard

import "fmt"

const packagePrefix = "timecard"

var (
	ErrUnsetID        = fmt.Errorf("%s: id is empty", packagePrefix)
	ErrUnsetCreatedAt = fmt.Errorf("%s: created at is zero", packagePrefix)
	ErrUnsetUpdatedAt = fmt.Errorf("%s: updated at is zero", packagePrefix)
)
