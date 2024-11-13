package timecard

import (
	"time"
	"timeCardSimple/app/domain/id"
)

func (tc *Timecard) ID() id.ID {
	return tc.options.ID
}

func (tc *Timecard) EmployeeID() id.ID {
	return tc.options.EmployeeID
}

func (tc *Timecard) StartTime() *time.Time {
	return tc.options.StartTime
}

func (tc *Timecard) EndTime() *time.Time {
	return tc.options.EndTime
}

func (tc *Timecard) CreatedAt() time.Time {
	return tc.options.CreatedAt
}

// UpdatedAt returns the User updatedAt field.
func (tc *Timecard) UpdatedAt() time.Time {
	return tc.options.UpdatedAt
}
