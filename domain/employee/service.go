package employee

//go:generate go run github.com/golang/mock/mockgen -package employeetest -destination employeetest/mock_test_service.go timeCardSimple/domain/employee Service

import (
	"context"
	"timeCardSimple/domain/id"
)

type Service interface {
	CreateEmployee(ctx context.Context, createParams CreateParamsWithPasswordStrings) (*Employee, error)
	DeleteEmployee(ctx context.Context, employeeID id.ID) error
}
