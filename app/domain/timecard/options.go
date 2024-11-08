package timecard

import (
	"time"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/lib/clone"
)

type Options struct {
	ID         id.ID
	EmployeeID id.ID
	StartTime  *time.Time
	EndTime    *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (o Options) validate() error {
	if o.ID == (id.ID{}) {
		return ErrUnsetID
	}

	if o.EmployeeID == (id.ID{}) {
		return ErrUnsetID
	}

	if o.CreatedAt.IsZero() {
		return ErrUnsetCreatedAt
	}

	if o.UpdatedAt.IsZero() {
		return ErrUnsetUpdatedAt
	}

	return nil
}

func (o Options) deepClone() Options {
	clonedRole := o

	clonedRole.StartTime = clone.Pointer(o.StartTime)
	clonedRole.EndTime = clone.Pointer(o.EndTime)

	return clonedRole
}
