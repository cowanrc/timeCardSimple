package timecard

import (
	"context"
	"time"
	"timeCardSimple/app/domain/id"
)

//go:generate go run github.com/golang/mock/mockgen -package timecardtest -destination timecardtest/mock_test_service.go timeCardSimple/app/domain/timecard Service

type Service interface {
	ClockIn(ctx context.Context, employeeID id.ID, startTime time.Time) (*Timecard, error)
}
