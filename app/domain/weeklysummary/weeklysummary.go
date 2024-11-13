package weeklysummary

import (
	"time"
	"timeCardSimple/app/domain/id"
)

type WeeklySummary struct {
	options Options
}

func New(employeeID id.ID, startDate time.Time) (*WeeklySummary, error) {
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

func NewWithOptions(options Options) (*WeeklySummary, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	return &WeeklySummary{
		options: options,
	}, nil
}

func (ws *WeeklySummary) Options() Options {
	return ws.options.DeepClone()
}
