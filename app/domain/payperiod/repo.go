package payperiod

//go:generate go run github.com/golang/mock/mockgen -package payperiodtest -destination payperiodtest/mock_test_repo.go timeCardSimple/app/domain/payperiod Repo

import (
	"context"
	"timeCardSimple/app/domain/id"
)

type QueryRepo interface {
	GetPayPeriodByEmployeeID(ctx context.Context, employeeID id.ID) (*PayPeriod, error)
}

type Repo interface {
	QueryRepo

	CreatePayPeriod(ctx context.Context, payPeriod *PayPeriod) error
}
