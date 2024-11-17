package weeklysummary

//go:generate go run github.com/golang/mock/mockgen -package weeklysummarytest -destination weeklysummarytest/mock_test_repo.go timeCardSimple/app/domain/weeklysummary Repo

import (
	"context"
	"timeCardSimple/app/domain/id"
)

type QueryRepo interface {
	GetWeeklySummaryByEmployeeID(ctx context.Context, employeeID id.ID) (*WeeklySummary, error)
}

type Repo interface {
	QueryRepo

	CreateWeeklySummary(ctx context.Context, weeklySummary *WeeklySummary) error
}
