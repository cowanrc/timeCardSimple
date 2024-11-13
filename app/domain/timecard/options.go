package timecard

import (
	"time"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/lib/clone"
)

type Options struct {
	ID                  id.ID
	EmployeeID          id.ID
	StartTime           *time.Time
	EndTime             *time.Time
	Duration            *float64
	WeekStartDate       *time.Time
	BiWeeklyPeriodStart *time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
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
	clonedOptions := o

	clonedOptions.StartTime = clone.Pointer(o.StartTime)
	clonedOptions.EndTime = clone.Pointer(o.EndTime)
	clonedOptions.Duration = clone.Pointer(o.Duration)
	clonedOptions.BiWeeklyPeriodStart = clone.Pointer(o.BiWeeklyPeriodStart)

	return clonedOptions
}
