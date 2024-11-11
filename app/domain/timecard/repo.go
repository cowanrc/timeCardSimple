package timecard

import (
	"context"
	"timeCardSimple/app/domain/id"
)

//go:generate go run github.com/golang/mock/mockgen -package timecardtest -destination timecardtest/mock_test_repo.go timeCardSimple/app/domain/timecard Repo

type Repo interface {
	CreateEmployeeTimecard(ctx context.Context, employeeID id.ID) (*Timecard, error)
}
