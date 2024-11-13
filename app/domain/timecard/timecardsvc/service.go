package timecardsvc

import (
	"context"
	"fmt"
	"time"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/timecard"
)

var _ timecard.Service = &Service{}

type Service struct {
	timecardRepo timecard.Repo
	employeeRepo employee.Repo
}

func New(
	timecardRepo timecard.Repo,
	employeeRepo employee.Repo,
) *Service {
	return &Service{
		timecardRepo: timecardRepo,
		employeeRepo: employeeRepo,
	}
}

func (s *Service) ClockIn(ctx context.Context, employeeID id.ID, startTime time.Time) (*timecard.Timecard, error) {
	now := time.Now()

	employeeTimecard, err := s.timecardRepo.GetTimecardByEmployeeID(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	if isTimecardActive(employeeTimecard) {
		return nil, fmt.Errorf("Timecard is active. Please clock out first before continuing")
	}

	weekStartDate := getWeekStartDate(now)
	biWeeklyStartDate, _ := getBiWeeklyPeriod(now)

	updateParams := timecard.UpdateParams{
		StartTime:           &now,
		WeekStartDate:       &weekStartDate,
		BiWeeklyPeriodStart: &biWeeklyStartDate,
		UpdatedAt:           now,
	}

	employeeTimecard.UpdateOptions(updateParams)

	err = s.timecardRepo.ClockInEmployee(ctx, employeeID, startTime, &weekStartDate, &biWeeklyStartDate)
	if err != nil {
		return nil, fmt.Errorf("error clocking in employee: %w", err)
	}

	return employeeTimecard, nil
}

func isTimecardActive(timecard *timecard.Timecard) bool {
	if timecard.StartTime() != nil && timecard.EndTime() == nil {
		return true
	}

	return false
}

func getWeekStartDate(date time.Time) time.Time {
	weekday := int(date.Weekday())
	if weekday == 0 { // Sunday
		weekday = 7
	}
	return date.AddDate(0, 0, -weekday+1)
}

func getBiWeeklyPeriod(date time.Time) (time.Time, time.Time) {
	day := date.Day()
	var startDate, endDate time.Time

	if day <= 15 {
		startDate = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
		endDate = time.Date(date.Year(), date.Month(), 15, 23, 59, 59, 999, date.Location())
	} else {
		startDate = time.Date(date.Year(), date.Month(), 16, 0, 0, 0, 0, date.Location())
		endDate = time.Date(date.Year(), date.Month()+1, 0, 23, 59, 59, 999, date.Location())
	}

	return startDate, endDate
}
