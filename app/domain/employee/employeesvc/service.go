package employeesvc

import (
	"context"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
)

var _ employee.Service = &Service{}

type Service struct {
	employeeRepo employee.Repo
	// passworder   password.Passworder
}

func New(
	employeeRepo employee.Repo,
	// passworder password.Passworder,
) *Service {
	return &Service{
		employeeRepo: employeeRepo,
		// passworder:   passworder,
	}
}

func (s *Service) CreateEmployee(ctx context.Context, createParams employee.CreateParams) (*employee.Employee, error) {
	// createEmployeePasswordHash, err := s.passworder.EvaluateAndHash(createParams.Password, createParams.ConfirmPassword)
	// if err != nil {
	// 	return nil, err
	// }

	e, err := employee.New(employee.CreateParams{
		FirstName: createParams.FirstName,
		LastName:  createParams.LastName,
		Email:     createParams.Email,
		// PasswordHash: createEmployeePasswordHash,
	})
	if err != nil {
		return nil, err
	}

	if err := s.employeeRepo.AddEmployee(ctx, e); err != nil {
		return nil, err
	}

	return e, nil
}

func (s *Service) GetEmployees(ctx context.Context) (*[]employee.Employee, error) {
	return nil, nil
}

func (s *Service) GetEmployeeByID(ctx context.Context, employeeID id.ID) (*employee.Employee, error) {
	e, err := s.employeeRepo.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (s *Service) DeleteEmployee(ctx context.Context, employeeID id.ID) error {
	if err := s.employeeRepo.RemoveEmployee(ctx, employeeID); err != nil {
		return err
	}

	return nil
}
