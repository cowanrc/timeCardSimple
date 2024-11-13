package payperiod

import (
	"time"
	"timeCardSimple/app/domain/id"
)

type PayPeriod struct {
	options Options
}

func New(employeeID id.ID, startDate time.Time) (*PayPeriod, error) {
	id, err := id.New()
	if err != nil {
		return nil, err
	}

	return NewWithOptions(Options{
		ID:         id,
		EmployeeID: employeeID,
		StartDate:  startDate,
		DaysWorked: 0,
		TotalHours: 0,
	})
}

func NewWithOptions(options Options) (*PayPeriod, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	return &PayPeriod{
		options: options,
	}, nil
}
