package employee

//go:generate go run github.com/golang/mock/mockgen -package employeetest -destination employeetest/mock_test_service.go timeCardSimple/app/domain/employee Service

import (
	"context"
	"timeCardSimple/app/domain/id"
)

type Service interface {
	CreateEmployee(ctx context.Context, createParams CreateParams) (*Employee, error)
	GetEmployeeByID(ctx context.Context, employeeID id.ID) (*Employee, error)
	DeleteEmployee(ctx context.Context, employeeID id.ID) error
}
