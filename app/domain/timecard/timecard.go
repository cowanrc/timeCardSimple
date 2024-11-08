package timecard

import (
	"time"
	"timeCardSimple/app/domain/id"
)

type Timecard struct {
	options Options
}

type UpdateParams struct {
	StartedTime *time.Time
	EndTime     *time.Time
	UpdatedAt   time.Time
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
