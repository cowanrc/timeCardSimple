package employeesvc

import (
	"context"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/password"
)

var _ employee.Service = &Service{}

type Service struct {
	employeeRepo employee.Repo
	passworder   password.Passworder
}

func New(
	employeeRepo employee.Repo,
	passworder password.Passworder,
) *Service {
	return &Service{
		employeeRepo: employeeRepo,
		passworder:   passworder,
	}
}

func (s *Service) CreateEmployee(ctx context.Context, createParams employee.CreateParamsWithPasswordStrings) (*employee.Employee, error) {
	createEmployeePasswordHash, err := s.passworder.EvaluateAndHash(createParams.Password, createParams.ConfirmPassword)
	if err != nil {
		return nil, err
	}

	e, err := employee.New(employee.CreateParamsWithPasswordHash{
		CreateParams: createParams.CreateParams,
		PasswordHash: createEmployeePasswordHash,
	})
	if err != nil {
		return nil, err
	}

	if err := s.employeeRepo.AddEmployee(ctx, e); err != nil {
		return nil, err
	}
}

func (s *Service) DeleteEmployee(ctx context.Context, employeeID id.ID) error {
	if err := s.employeeRepo.RemoveEmployee(ctx, employeeID); err != nil {
		return err
	}
}
