package employee

import (
	"timeCardSimple/database"
	"timeCardSimple/errors"
)

var (
	EmployeeService employeesServiceInterface = &employeesService{}
)

type employeesService struct{}

type employeesServiceInterface interface {
	CreateEmployee(database.Employee) (*database.Employee, *errors.RestErr)
	// GetAllEmployees() (database.Employee, *errors.RestErr)
	// GetEmployee(int64) (*database.Employee, *errors.RestErr)
	// DeleteEmployee(int64) *errors.RestErr
}

func (s *employeesService) CreateEmployee(employee database.Employee) (*database.Employee, *errors.RestErr) {

	// employee.DateOfBirth = date_utils.GetNowDBFormat()

	if err := employee.Save(); err != nil {
		return nil, err
	}

	return &employee, nil
}
