package timecard

import (
	"time"
	"timeCardSimple/app/domain/id"
)

type Timecard struct {
	options Options
}

type UpdateParams struct {
	StartTime           *time.Time
	EndTime             *time.Time
	Duration            *float64
	WeekStartDate       *time.Time
	BiWeeklyPeriodStart *time.Time
	UpdatedAt           time.Time
}

func New(employeeID id.ID) (*Timecard, error) {
	id, err := id.New()
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return NewWithOptions(Options{
		ID:         id,
		EmployeeID: employeeID,
		CreatedAt:  now,
		UpdatedAt:  now,
	})
}

func NewWithOptions(options Options) (*Timecard, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	return &Timecard{
		options: options,
	}, nil
}

func (t *Timecard) UpdateOptions(params UpdateParams) {
	if params.StartTime != nil {
		t.options.StartTime = params.StartTime
	}
	if params.EndTime != nil {
		t.options.EndTime = params.EndTime
	}
	if params.Duration != nil {
		t.options.Duration = params.Duration
	}
	if params.WeekStartDate != nil {
		t.options.WeekStartDate = params.WeekStartDate
	}
	if params.BiWeeklyPeriodStart != nil {
		t.options.BiWeeklyPeriodStart = params.BiWeeklyPeriodStart
	}

	t.options.UpdatedAt = params.UpdatedAt

}

// Options returns the options for the role
func (t *Timecard) Options() Options {
	return t.options.deepClone()
}

// Clone returns a copy of the role
func (t *Timecard) Clone() Timecard {
	return Timecard{
		options: t.options.deepClone(),
	}
}
