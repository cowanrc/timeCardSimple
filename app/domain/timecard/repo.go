package timecard

import (
	"context"
	"time"
	"timeCardSimple/app/domain/id"
)

//go:generate go run github.com/golang/mock/mockgen -package timecardtest -destination timecardtest/mock_test_repo.go timeCardSimple/app/domain/timecard Repo

type QueryRepo interface {
	GetTimecardByEmployeeID(ctx context.Context, employeeID id.ID) (*Timecard, error)
}

type Repo interface {
	QueryRepo

	CreateEmployeeTimecard(ctx context.Context, timecard *Timecard) error
	ClockInEmployee(ctx context.Context, employeeID id.ID, startTime time.Time, weekStartDate *time.Time, BiWeekStartDate *time.Time) error
}
