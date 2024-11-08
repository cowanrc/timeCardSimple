package employee

import (
	"context"
	"timeCardSimple/app/domain/id"
)

//go:generate go run github.com/golang/mock/mockgen -package employeetest -destination employeetest/mock_test_repo.go timeCardSimple/app/domain/employee Repo

// QueryRepo provides the methods for querying Users in the application.
type QueryRepo interface {
	GetEmployeeByID(ctx context.Context, employeeID id.ID) (*Employee, error)
	GetEmployeeByEmail(ctx context.Context, email string) (*Employee, error)
	GetEmployees(ctx context.Context) ([]*Employee, error)
}

type Repo interface {
	QueryRepo
	AddEmployee(ctx context.Context, employee *Employee) error
	RemoveEmployee(ctx context.Context, employeeID id.ID) error
}
