package dto

import (
	"time"
	"timeCardSimple/app/domain/timecard"
)

func Timecard(v interface{}) interface{} {
	t := v.(*timecard.Timecard)

	id, _ := Transform(t.ID())
	employeeID, _ := Transform(t.EmployeeID())

	return &struct {
		ID                  interface{} `json:"id"`
		EmployeeID          interface{} `json:"employee_id"`
		StartTime           *time.Time  `json:"start_time"`
		EndTime             *time.Time  `json:"end_time"`
		Duration            *float64    `json:"duration"`
		WeeklyStartDate     *time.Time  `json:"weekly_start_date"`
		BiWeeklyPeriodStart *time.Time  `json:"biweekly_period_start"`
		CreatedAt           time.Time   `json:"created_at"`
		UpdatedAt           time.Time   `json:"updated_at"`
	}{
		id,
		employeeID,
		t.StartTime(),
		t.EndTime(),
		t.Duration(),
		t.WeeklyStartDate(),
		t.BiWeeklyPeriodStart(),
		t.CreatedAt(),
		t.UpdatedAt(),
	}
}
