package weeklysummary

import (
	"time"
	"timeCardSimple/app/domain/id"
)

type Options struct {
	ID         id.ID
	EmployeeID id.ID
	StartDate  time.Time
	DaysWorked int64
	TotalHours int64
}

func (o Options) validate() error {
	if o.ID == (id.ID{}) {
		return ErrUnsetID
	}

	if o.EmployeeID == (id.ID{}) {
		return ErrUnsetEmployeeID
	}

	return nil
}

func (o Options) DeepClone() Options {
	clonedOptions := o

	return clonedOptions
}
