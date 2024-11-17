package payperiod

import (
	"time"
	"timeCardSimple/app/domain/id"
)

type Options struct {
	ID         id.ID
	EmployeeID id.ID
	StartDate  *time.Time
	EndDate    *time.Time
	DaysWorked uint64
	TotalHours uint64
}

func (o *Options) validate() error {
	return nil
}
