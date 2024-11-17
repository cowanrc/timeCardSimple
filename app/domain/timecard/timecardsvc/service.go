package timecardsvc

import (
	"context"
	"fmt"
	"strings"
	"time"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/payperiod"
	"timeCardSimple/app/domain/timecard"
	"timeCardSimple/app/domain/weeklysummary"
)

var _ timecard.Service = &Service{}

type Service struct {
	timecardRepo      timecard.Repo
	employeeRepo      employee.Repo
	weeklySummaryRepo weeklysummary.Repo
	payPeriodRepo     payperiod.Repo
}

func New(
	timecardRepo timecard.Repo,
	employeeRepo employee.Repo,
	weeklySummaryRepo weeklysummary.Repo,
	payPeriodRepo payperiod.Repo,
) *Service {
	return &Service{
		timecardRepo:      timecardRepo,
		employeeRepo:      employeeRepo,
		weeklySummaryRepo: weeklySummaryRepo,
		payPeriodRepo:     payPeriodRepo,
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

	err = s.getOrCreateWeeklySummary(ctx, employeeID, weekStartDate)
	if err != nil {
		return nil, fmt.Errorf("error handling weekly summary: %w", err)
	}

	err = s.getOrCreatePayPeriod(ctx, employeeID, biWeeklyStartDate)
	if err != nil {
		return nil, fmt.Errorf("error handling pay period: %w", err)
	}

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
	if weekday == 0 {
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

func (s *Service) getOrCreateWeeklySummary(ctx context.Context, employeeID id.ID, weekStartDate time.Time) error {
	weeklySummary, err := s.weeklySummaryRepo.GetWeeklySummaryByEmployeeID(ctx, employeeID)
	if err != nil && strings.Contains(err.Error(), "not found") {
		weeklySummary, err = weeklysummary.New(employeeID, weekStartDate)
		if err != nil {
			return fmt.Errorf("error creating new weekly summary: %w", err)
		}
		err = s.weeklySummaryRepo.CreateWeeklySummary(ctx, weeklySummary)
		if err != nil {
			return fmt.Errorf("failed to create weekly summary: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("error retrieving weekly summary: %w", err)
	}
	return nil
}

// Helper function to retrieve or create a pay period
func (s *Service) getOrCreatePayPeriod(ctx context.Context, employeeID id.ID, biWeeklyStartDate time.Time) error {
	payPeriod, err := s.payPeriodRepo.GetPayPeriodByEmployeeID(ctx, employeeID)
	if err != nil && strings.Contains(err.Error(), "not found") {
		payPeriod, err = payperiod.New(employeeID, biWeeklyStartDate)
		if err != nil {
			return fmt.Errorf("error creating new pay period: %w", err)
		}
		err = s.payPeriodRepo.CreatePayPeriod(ctx, payPeriod)
		if err != nil {
			return fmt.Errorf("failed to create pay period: %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("error retrieving pay period: %w", err)
	}
	return nil
}
